package encode

import (
	"strconv"

	"github.com/vertex-language/arm64/feature"
	"github.com/vertex-language/arm64/internal/isa"
	"github.com/vertex-language/arm64/operand"
)

// Opts is what an encode call needs that is not an operand.
//
// x86_64's Opts carries broadcast, zeroing, rounding and SAE — one bit each,
// with no register behind them, which is what keeps them out of the operand
// list. A64 has no such modifier: a predicate is a register and goes in a slot,
// and a shift is an operand. So this carries call context only, and exists so
// the signature does not have to change when something does turn up.
type Opts struct {
	// Offset is the byte offset of this instruction within its section. It is
	// copied onto every Fixup and is otherwise unused; nothing here computes an
	// address.
	Offset int
}

// Encode resolves a mnemonic against a feature set and encodes it.
func Encode(set feature.Set, mnem string, ops ...any) (uint32, []Fixup, error) {
	return EncodeWith(Opts{}, set, mnem, ops...)
}

// EncodeWith is Encode with call context.
func EncodeWith(opts Opts, set feature.Set, mnem string, ops ...any) (uint32, []Fixup, error) {
	vals := make([]val, len(ops))
	args := make([]isa.Arg, len(ops))
	for i, o := range ops {
		vals[i] = lower(o)
		args[i] = vals[i].arg()
	}
	f, err := isa.Resolve(mnem, args, set)
	if err != nil {
		return 0, nil, err
	}
	return encodeForm(f, vals, opts)
}

// EncodeForm encodes against a form the caller already resolved, which is what
// the generated helpers do.
//
// It does not consult a feature set. The gate was checked when the form was
// resolved, and re-checking here would let the two answers disagree.
func EncodeForm(f *isa.Form, ops []any, opts Opts) (uint32, []Fixup, error) {
	vals := make([]val, len(ops))
	for i, o := range ops {
		vals[i] = lower(o)
	}
	return encodeForm(f, vals, opts)
}

func encodeForm(f *isa.Form, ops []val, opts Opts) (uint32, []Fixup, error) {
	word := f.Word
	var fixups []Fixup

	// Slots and operands are not in step. A memory operand is one value
	// filling two slots — a base and an offset — so the two indices are walked
	// separately rather than zipped.
	// prevImm is the value of the last immediate operand encoded, which one
	// rule needs: UBFX's imms is computed from its own operand and the one
	// before it. Nothing else reads it.
	prevImm := int64(0)

	// filled is the field an immediate's sibling wrote into, if any. An
	// optional shift slot that goes unwritten normally places its default,
	// and the default would land on top of the sibling — `movz x0, #0x10000`
	// names a halfword the immediate encoder put in Hw, and no shift was
	// written, so the two slots are competing for one field. The immediate
	// wins: it is the one that knows the value.
	var filled *isa.Field

	oi := 0
	for si := 0; si < len(f.Slots); si++ {
		s := f.Slots[si]

		if oi >= len(ops) {
			if !s.Optional {
				return 0, nil, &CountError{Form: f, Got: len(ops)}
			}
			if filled == nil || *filled != s.Field {
				word = place(word, s.Field, s.Default)
			}
			continue
		}
		v := ops[oi]

		switch {
		case s.Class.Mem():
			w, fx, consumed, err := encodeMem(f, oi, si, v, opts)
			if err != nil {
				return 0, nil, err
			}
			word = w(word)
			fixups = append(fixups, fx...)
			si += consumed - 1
			oi++

		case s.Class.Reg():
			if v.kind != valReg && v.kind != valSys {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			n, why := regNum(s.Class, v.reg)
			if why != "" {
				return 0, nil, &RegisterError{f, oi, s.Class, v.reg.String(), why}
			}
			word = place(word, s.Field, n)
			oi++

		// ClassSys is not s.Class.Reg(): class.go gives it reg.FileNone on
		// purpose, so a plain GPR or vector register can never satisfy a
		// system-register slot by sharing a file check with them. That means
		// it needs its own case here rather than falling under the one above,
		// even though regNum's own switch already has the row that handles it.
		case s.Class == isa.ClassSys:
			if v.kind != valSys {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			n, why := regNum(s.Class, v.reg)
			if why != "" {
				return 0, nil, &RegisterError{f, oi, s.Class, v.reg.String(), why}
			}
			word = place(word, s.Field, n)
			oi++

		case s.Class == isa.ClassImm:
			// The low-twelve half of an address landing in an immediate
			// field: add x0, x0, :lo12:msg. The value is not a number yet, so
			// the field stays zero and a fixup carries the role. Only the
			// page-offset roles are accepted — a direct reference here would
			// be inventing the half of the address Mem.Off refuses to invent,
			// and the page half has its own field on ADRP.
			if v.kind == valRef {
				switch v.ref.Role {
				case operand.RolePageOff, operand.RoleGotPageOff, operand.RoleTlvPageOff:
				default:
					return 0, nil, &OperandError{f, oi, s.Class, v.raw}
				}
				fixups = append(fixups, Fixup{
					Field:  s.Field,
					Target: v.ref.T,
					Role:   v.ref.Role,
					Kind:   v.ref.Kind(),
					Addend: v.ref.Addend(),
					Access: operand.Width(f.AccessBits()),
					Bits:   s.Field.Width(),
					Scale:  scaleOf(f),
				})
				oi++
				continue
			}
			if v.kind != valImm {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			r, err := encodeImm(f, oi, s, v, hasShiftOperand(ops), prevImm)
			if err != nil {
				return 0, nil, err
			}
			prevImm = v.imm
			word = place(word, s.Field, r.value)
			if r.hasSibling {
				// ImmShiftLeft and ImmShiftRight write their sibling into
				// Imms directly: the row that carries them declares only the
				// one Imm slot the shift amount itself lands in, so there is
				// no second slot for siblingField to find by class.
				if s.Imm == isa.ImmShiftLeft || s.Imm == isa.ImmShiftRight {
					word = place(word, isa.Imms, r.sibling)
				} else if fld, ok := siblingField(f, r.siblingClass); ok {
					word = place(word, fld, r.sibling)
					sib := fld
					filled = &sib
				} else if r.sibling != 0 {
					// No field for it, and it is not zero: encoding this
					// form would drop part of the value and produce a
					// different instruction with no diagnostic. `mov x0,
					// #0x10000` is the case — the MOV alias row carries the
					// halfword and not the shift, and the row that can say
					// which halfword is MOVZ's.
					return 0, nil, &RangeError{f, oi, v.imm,
						"a value this form encodes in one field: it names a halfword and " +
							"has nowhere to put the shift, so a value in a higher halfword " +
							"needs the instruction the alias stands for"}
				}
			}
			oi++

		case s.Class == isa.ClassLabel:
			w, fx, err := encodeTarget(f, oi, s, v, opts)
			if err != nil {
				return 0, nil, err
			}
			word = w(word)
			fixups = append(fixups, fx...)
			oi++

		case s.Class == isa.ClassShift:
			if v.kind != valShift {
				// An unwritten optional shift; the slot keeps its default and
				// this operand belongs to a later slot.
				if s.Optional {
					if filled == nil || *filled != s.Field {
						word = place(word, s.Field, s.Default)
					}
					continue
				}
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			n, sibling, hasSibling, err := encodeShift(f, oi, s, v.shift)
			if err != nil {
				return 0, nil, err
			}
			word = place(word, s.Field, n)
			if hasSibling {
				word = place(word, isa.Imm6, sibling)
			}
			oi++

		case s.Class == isa.ClassExtend:
			if v.kind != valExtend {
				if s.Optional {
					word = place(word, s.Field, s.Default)
					continue
				}
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			if !v.ext.Valid() {
				return 0, nil, &RangeError{f, oi, int64(v.ext.Amount),
					"an extend shift of 0 to 4"}
			}
			word = place(word, s.Field, uint64(v.ext.Op))
			if fld, ok := siblingField(f, isa.ClassImm); ok && f.Attrs&isa.AttrScaled == 0 {
				word = place(word, fld, uint64(v.ext.Amount))
			}
			oi++

		case s.Class == isa.ClassCond:
			if v.kind != valCond {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			if !v.cond.Valid() {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			c := v.cond
			if f.Attrs&isa.AttrInvertCond != 0 {
				// AL and NV invert to nothing — both always execute — and
				// Invert says so. They pass through unchanged, which is the
				// architecture's own reading of "cset x0, al".
				if inv, ok := c.Invert(); ok {
					c = inv
				}
			}
			word = place(word, s.Field, uint64(c))
			oi++

		case s.Class == isa.ClassBarrier:
			if v.kind != valBarrier {
				if s.Optional {
					word = place(word, s.Field, s.Default)
					continue
				}
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			if !v.bar.Valid() {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			word = place(word, s.Field, uint64(v.bar))
			oi++

		case s.Class == isa.ClassPState:
			// The field is op1 and op2, either side of the immediate. Two
			// fields and one operand, which is why the slot names neither
			// and the form's attribute says what to do instead.
			if v.kind != valPState || !v.pst.Valid() {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			word = place(word, isa.Op1, uint64(v.pst.Op1()))
			word = place(word, isa.Op2, uint64(v.pst.Op2()))
			oi++

		case s.Class == isa.ClassPrfOp:
			if v.kind != valPrfOp {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			if !v.prf.Valid() {
				return 0, nil, &OperandError{f, oi, s.Class, v.raw}
			}
			word = place(word, s.Field, uint64(v.prf))
			oi++

		default:
			return 0, nil, &OperandError{f, oi, s.Class, v.raw}
		}
	}

	if oi != len(ops) {
		return 0, nil, &CountError{Form: f, Got: len(ops)}
	}

	// One source register named twice: ROR (immediate) is EXTR with Rm equal
	// to Rn, and the row says so rather than the caller.
	if f.Attrs&isa.AttrRnIntoRm != 0 {
		word = place(word, isa.Rm, isa.Rn.Get(word))
	}

	for i := range fixups {
		fixups[i].Offset = opts.Offset
	}
	return word, fixups, nil
}

// hasShiftOperand reports whether the caller wrote a shift, which changes the
// immediate rule rather than adding to it.
func hasShiftOperand(ops []val) bool {
	for _, v := range ops {
		if v.kind == valShift {
			return true
		}
	}
	return false
}

// encodeShift fills a shift slot.
//
// What the field holds depends on which immediate the shift belongs to, not on
// the field's own width, and the two happen to differ in width as well — which
// is why an earlier version of this function got it wrong in a way that still
// produced a plausible word. Hw and Shift are both two bits. Hw holds a
// halfword index and Shift holds a shift kind, and writing the kind into Hw
// encodes movz x0, #1, lsl #16 as lsl #0.
// encodeShift returns the value for the shift slot's own field and, for a
// shifted-register form, the amount as a sibling value the caller places into
// Imm6 — the field is fixed at bits 15:10 on every add/sub/logical
// shifted-register form the architecture has, which is why it is a literal
// here rather than something looked up on the form the way ImmAddSub12's and
// ImmMoveWide's sibling fields are: those two write back into a field the
// same slot list also declares for a different reason (Sh, Hw), and this one
// writes into a field the shifted-register forms never declare at all,
// because nothing about them needs to name it except this.
func encodeShift(f *isa.Form, i int, s isa.Slot, sh operand.ShiftOp) (value, sibling uint64, hasSibling bool, err error) {
	switch immKindOf(f) {
	case isa.ImmMoveWide:
		if sh.Op != operand.LSL {
			return 0, 0, false, &OperandError{f, i, isa.ClassShift, sh}
		}
		if sh.Amount%16 != 0 {
			return 0, 0, false, &RangeError{f, i, int64(sh.Amount),
				"a shift of 0, 16, 32 or 48: the field names a halfword, not a distance"}
		}
		hw := uint64(sh.Amount) / 16
		limit := uint64(4)
		if FormWidth(f) == 32 {
			limit = 2
		}
		if hw >= limit {
			return 0, 0, false, &RangeError{f, i, int64(sh.Amount),
				"a shift inside a " + operand.Width(FormWidth(f)).String()}
		}
		return hw, 0, false, nil

	case isa.ImmAddSub12:
		if !sh.ValidImm12() {
			return 0, 0, false, &RangeError{f, i, int64(sh.Amount),
				"lsl #0 or lsl #12: the field is one bit"}
		}
		if sh.Amount == 12 {
			return 1, 0, false, nil
		}
		return 0, 0, false, nil
	}

	// A shifted-register form: the field is the shift kind, and the amount is
	// six bits at Imm6 regardless of which shifted-register mnemonic this is.
	if !sh.Op.Valid() {
		return 0, 0, false, &OperandError{f, i, isa.ClassShift, sh}
	}
	if !fitsU(uint64(sh.Amount), isa.Imm6.Width()) {
		return 0, 0, false, &RangeError{f, i, int64(sh.Amount),
			plural(isa.Imm6.Width()) + " of shift amount"}
	}
	return uint64(sh.Op), uint64(sh.Amount), true, nil
}

// encodeMem fills a memory operand's slots and returns how many it consumed.
func encodeMem(f *isa.Form, oi, si int, v val, opts Opts) (func(uint32) uint32, []Fixup, int, error) {
	if v.kind != valMem {
		return nil, nil, 0, &OperandError{f, oi, f.Slots[si].Class, v.raw}
	}
	m := v.mem
	if err := m.Validate(); err != nil {
		return nil, nil, 0, &AddressError{f, oi, m, err.Error()}
	}

	base := f.Slots[si]
	var off *isa.Slot
	consumed := 1
	if si+1 < len(f.Slots) && f.Slots[si+1].Role == isa.RoleOffset {
		off = &f.Slots[si+1]
		consumed = 2
	}

	// The writeback forms are separate encodings, and a form that is not one
	// cannot express an address that asks for it.
	switch m.Form {
	case operand.AddrPreIndex:
		if f.Attrs&isa.AttrPreIndex == 0 {
			return nil, nil, 0, &AddressError{f, oi, m,
				"this form has no writeback; a pre-indexed address is a different encoding"}
		}
	case operand.AddrPostIndex:
		if f.Attrs&isa.AttrPostIndex == 0 {
			return nil, nil, 0, &AddressError{f, oi, m,
				"this form has no writeback; a post-indexed address is a different encoding"}
		}
	case operand.AddrRegOffset:
		if f.Attrs&isa.AttrRegOffset == 0 {
			return nil, nil, 0, &AddressError{f, oi, m,
				"this form takes an immediate offset; a register-offset address is a different encoding"}
		}
		return encodeRegOffset(f, oi, base, m, consumed)
	}

	n, why := regNum(isa.ClassXsp, m.Base)
	if why != "" {
		return nil, nil, 0, &RegisterError{f, oi, isa.ClassXsp, m.Base.String(), why}
	}

	var fixups []Fixup
	apply := func(w uint32) uint32 { return place(w, base.Field, n) }

	if off == nil {
		// This form has no offset field at all. An address that states one is
		// refused whatever its value: accepting a zero displacement here would
		// mean [x1, #0] assembles and [x1, #8] does not, which reads as a range
		// problem when it is a form problem, and sends the caller looking for
		// the wrong fix.
		if m.Form != operand.AddrBase {
			return nil, nil, 0, &AddressError{f, oi, m,
				"this form takes a base register only; the offset forms are separate encodings"}
		}
		return apply, nil, consumed, nil
	}

	if m.Symbolic() {
		fixups = append(fixups, Fixup{
			Field:  off.Field,
			Target: m.Disp.Ref.T,
			Role:   m.Disp.Ref.Role,
			Kind:   m.Disp.Ref.Kind(),
			Addend: m.Disp.Ref.Addend(),
			Access: operand.Width(f.AccessBits()),
			Bits:   off.Field.Width(),
			Scale:  scaleOf(f),
		})
		return apply, fixups, consumed, nil
	}

	// No previous immediate: a memory offset's rule never reads one.
	r, err := encodeImm(f, oi, *off, val{kind: valImm, imm: m.Disp.Const,
		uimm: uint64(m.Disp.Const)}, false, 0)
	if err != nil {
		return nil, nil, 0, err
	}
	inner := apply
	apply = func(w uint32) uint32 { return place(inner(w), off.Field, r.value) }
	return apply, fixups, consumed, nil
}

// encodeRegOffset fills the three fields of [Xn, Xm{, extend {amount}}].
//
// The extend's own value is the option field — operand.Extend is defined that
// way — and the shift is one bit, either no shift or exactly the log of the
// access width, which Mem.Validate has already checked. There is no
// displacement and so no fixup: an index register is a value at run time and
// a relocation has nothing to say about it.
func encodeRegOffset(f *isa.Form, oi int, base isa.Slot, m operand.Mem, consumed int) (func(uint32) uint32, []Fixup, int, error) {
	n, why := regNum(isa.ClassXsp, m.Base)
	if why != "" {
		return nil, nil, 0, &RegisterError{f, oi, isa.ClassXsp, m.Base.String(), why}
	}

	// The class the index is read at is the extend's business rather than
	// the slot's — Mem.Validate has already refused a W index under SXTX and
	// an X index under UXTW — so this asks only for the register number.
	cls := isa.ClassX
	if m.Ext.SourceIsW() {
		cls = isa.ClassW
	}
	idx, why := regNum(cls, m.Index)
	if why != "" {
		return nil, nil, 0, &RegisterError{f, oi, cls, m.Index.String(), why}
	}

	// The shift is one bit: no shift, or exactly the log of this form's
	// access width. The form is what knows that width — an address written
	// in assembly does not state one, because the mnemonic already did —
	// so the check belongs here rather than in Mem.Validate.
	sbit := uint64(0)
	if m.Amount != 0 {
		sc, ok := operand.Width(f.AccessBits()).Scale()
		if !ok || m.Amount != sc {
			return nil, nil, 0, &AddressError{f, oi, m,
				"index shift must be 0 or " + strconv.Itoa(int(sc)) + " for this access width"}
		}
		sbit = 1
	}
	opt := uint64(m.Ext)

	return func(w uint32) uint32 {
		w = place(w, base.Field, n)
		w = place(w, isa.Rm, idx)
		w = place(w, isa.Option, opt)
		return place(w, isa.SBit, sbit)
	}, nil, consumed, nil
}

func scaleOf(f *isa.Form) uint8 {
	if f.Attrs&isa.AttrScaled == 0 {
		return 0
	}
	sc, _ := operand.Width(f.AccessBits()).Scale()
	return sc
}

// encodeTarget fills a branch or address destination.
//
// A bare immediate is a displacement the caller computed and takes
// responsibility for. A reference is left blank and recorded, because the
// distance is not a number until something knows where both ends are.
func encodeTarget(f *isa.Form, i int, s isa.Slot, v val, opts Opts) (func(uint32) uint32, []Fixup, error) {
	switch v.kind {
	case valImm:
		r, err := encodeImm(f, i, s, v, false, 0)
		if err != nil {
			return nil, nil, err
		}
		return func(w uint32) uint32 { return place(w, s.Field, r.value) }, nil, nil

	case valRef:
		role := v.ref.Role
		if role == operand.RoleDirect && s.Role == isa.RolePage {
			// adrp x0, msg with no modifier written: the form is the page
			// form, so the reference is to the page. This is the one place a
			// role is inferred, and it is inferred from the form rather than
			// from the spelling.
			role = operand.RolePage
		}
		fx := Fixup{
			Field:  s.Field,
			Target: v.ref.T,
			Role:   role,
			Kind:   v.ref.Kind(),
			Addend: v.ref.Addend(),
			Access: operand.Width(f.AccessBits()),
			Branch: f.Attrs&isa.AttrBranch != 0,
			Bits:   s.Field.Width(),
			Scale:  targetScale(role, f),
		}
		return func(w uint32) uint32 { return w }, []Fixup{fx}, nil
	}
	return nil, nil, &OperandError{f, i, isa.ClassLabel, v.raw}
}

func targetScale(role operand.AddrRole, f *isa.Form) uint8 {
	switch role {
	case operand.RolePage, operand.RoleGotPage, operand.RoleTlvPage:
		return 12 // 4KiB pages
	case operand.RolePageOff, operand.RoleGotPageOff, operand.RoleTlvPageOff:
		return scaleOf(f)
	}
	if f.Attrs&isa.AttrBranch != 0 {
		return 2 // word-aligned
	}
	return 0
}
