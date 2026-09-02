package asm

import (
	"strconv"
	"strings"

	"github.com/vertex-language/arm64/operand"
	"github.com/vertex-language/arm64/reg"
	"github.com/vertex-language/gas"
)

// parseOperand reads one comma-separated operand and returns it in whatever
// type the encoder accepts for it.
//
// The order the cases are tried in is the whole of the ambiguity resolution,
// and it is: registers first, then the closed vocabularies (conditions,
// shifts, extends, barriers, prefetch hints), then anything else as an
// expression. A symbol named `eq` is therefore a condition and not a symbol,
// which is what GNU as does too — the alternative is deciding by the mnemonic,
// which means the operand parser has to know the instruction set, which is the
// thing this design is arranged to avoid.
func (t *target) parseOperand(p *gas.Parser) (any, error) {
	tok := p.Peek()

	// `[` opens an address.
	if tok.IsPunct("[") {
		return t.parseMem(p)
	}

	// `#` opens an immediate. It is optional in UAL for a few forms, so a
	// bare expression reaching parseExprOperand is not an error.
	if tok.IsPunct("#") {
		p.Take()
		return t.parseImm(p)
	}

	// `:lo12:sym` and friends, in the operand position of an add or a load.
	if tok.IsPunct(":") {
		return t.parseAddrRef(p)
	}

	if tok.Kind == gas.Ident {
		if v, ok := t.namedOperand(p, tok); ok {
			return v, nil
		}
	}

	return t.parseExprOperand(p)
}

// namedOperand resolves the identifier vocabularies. It reports false when the
// name is none of them, leaving the token unconsumed.
func (t *target) namedOperand(p *gas.Parser, tok gas.Token) (any, bool) {
	name := tok.Text

	// A vector register with an arrangement or a lane lexes as one
	// identifier, because `.` continues an identifier and `v0.16b` has one
	// in the middle. Splitting it here is cheaper than teaching the lexer
	// about registers.
	if r, ok := parseVecReg(p, tok); ok {
		return r, true
	}

	if r, ok := reg.Lookup(name); ok {
		p.Take()
		return r, true
	}

	lower := strings.ToLower(name)

	if s, ok := lookupShift(lower); ok {
		p.Take()
		// `lsl #3`; a shift with no amount is `lsl #0`, which some forms
		// admit and the table decides.
		n, err := t.optAmount(p)
		if err != nil {
			return nil, false
		}
		return operand.Shifted(s, n), true
	}
	if e, ok := lookupExtend(lower); ok {
		p.Take()
		n, err := t.optAmount(p)
		if err != nil {
			return nil, false
		}
		return operand.Extended(e, n), true
	}
	if c, ok := operand.LookupCond(lower); ok {
		p.Take()
		return c, true
	}
	if b, ok := operand.LookupBarrier(lower); ok {
		p.Take()
		return b, true
	}
	if pf, ok := operand.LookupPrfOp(lower); ok {
		p.Take()
		return pf, true
	}
	return nil, false
}

// optAmount reads the `#n` a shift or extend may carry.
func (t *target) optAmount(p *gas.Parser) (uint8, error) {
	if !p.Peek().IsPunct("#") {
		return 0, nil
	}
	p.Take()
	at := p.Peek()
	v, err := p.Expr()
	if err != nil {
		return 0, err
	}
	if !v.IsAbs() || v.Off < 0 || v.Off > 63 {
		return 0, p.Errorf(at, "shift amount %s is not between 0 and 63", v)
	}
	return uint8(v.Off), nil
}

var shiftNames = map[string]operand.Shift{
	"lsl": operand.LSL, "lsr": operand.LSR,
	"asr": operand.ASR, "ror": operand.ROR,
}

func lookupShift(s string) (operand.Shift, bool) { v, ok := shiftNames[s]; return v, ok }

var extendNames = map[string]operand.Extend{
	"uxtb": operand.UXTB, "uxth": operand.UXTH, "uxtw": operand.UXTW, "uxtx": operand.UXTX,
	"sxtb": operand.SXTB, "sxth": operand.SXTH, "sxtw": operand.SXTW, "sxtx": operand.SXTX,
}

func lookupExtend(s string) (operand.Extend, bool) { v, ok := extendNames[s]; return v, ok }

// parseVecReg handles `v0.16b`, `v0.4s` and `v2.s[1]`.
//
// The arrangement and the lane are the only place this architecture puts a
// `.` inside an operand, which is why they are the only reason this function
// exists.
func parseVecReg(p *gas.Parser, tok gas.Token) (reg.Reg, bool) {
	name := strings.ToLower(tok.Text)
	dot := strings.IndexByte(name, '.')
	if dot <= 0 || len(name) < 2 || name[0] != 'v' {
		return nil, false
	}
	num, err := strconv.Atoi(name[1:dot])
	if err != nil || num < 0 || num > 31 {
		return nil, false
	}
	v := reg.V(num)
	spec := name[dot+1:]

	// `v2.s[1]`: the element width alone, with the index in brackets.
	if e, ok := elemNames[spec]; ok && p.PeekAt(1).IsPunct("[") {
		p.Take() // the register
		p.Take() // [
		idx := p.Peek()
		if idx.Kind != gas.Int || idx.Num < 0 || idx.Num > 15 {
			return nil, false
		}
		p.Take()
		if !p.AcceptPunct("]") {
			return nil, false
		}
		return v.Lane(e, uint8(idx.Num)), true
	}

	if a, ok := arrNames[spec]; ok {
		p.Take()
		return v.Arr(a), true
	}
	return nil, false
}

var elemNames = map[string]reg.Elem{
	"b": reg.ElemB, "h": reg.ElemH, "s": reg.ElemS, "d": reg.ElemD,
}

var arrNames = map[string]reg.Arrangement{
	"8b": reg.V8B, "16b": reg.V16B,
	"4h": reg.V4H, "8h": reg.V8H,
	"2s": reg.V2S, "4s": reg.V4S,
	"1d": reg.V1D, "2d": reg.V2D,
}

// parseImm reads what follows a `#`.
func (t *target) parseImm(p *gas.Parser) (any, error) {
	// `#:lo12:sym` — the immediate of an add, or a load's offset written
	// outside the brackets, which never happens but costs nothing to admit.
	if p.Peek().IsPunct(":") {
		return t.parseAddrRef(p)
	}
	at := p.Peek()
	v, err := p.Expr()
	if err != nil {
		return nil, err
	}
	if !v.IsAbs() {
		return nil, p.Errorf(at, "an immediate must be a constant, and %s is an address; "+
			"write a relocation modifier such as :lo12: to name part of one", v)
	}
	return operand.Imm(v.Off), nil
}

// modifiers maps GNU as's relocation modifier spellings onto the roles the
// operand package names.
//
// This is the mapping gas cannot make and the reason Reloc is the
// architecture's business: `:lo12:` is meaningless on x86 and `@GOTPCREL` is
// meaningless here, so a shared layer holding a table of both would be holding
// a table of two unrelated things.
var modifiers = map[string]func(operand.Target) operand.AddrRef{
	"lo12":     operand.PageOff,
	"got":      operand.GotPage,
	"got_lo12": operand.GotPageOff,
}

// parseAddrRef reads `:mod:symbol`.
func (t *target) parseAddrRef(p *gas.Parser) (any, error) {
	open := p.Take() // ':'
	name := p.Peek()
	if name.Kind != gas.Ident {
		return nil, p.Errorf(name, "want a relocation modifier after ':', got %s", name)
	}
	p.Take()
	if err := p.ExpectPunct(":"); err != nil {
		return nil, err
	}
	mk, ok := modifiers[strings.ToLower(name.Text)]
	if !ok {
		return nil, p.Errorf(name, "unknown relocation modifier :%s:", name.Text)
	}
	at := p.Peek()
	v, err := p.Expr()
	if err != nil {
		return nil, err
	}
	if v.Sym == "" || v.Sub != "" {
		return nil, p.Errorf(at, ":%s: wants a symbol, got %s", name.Text, v)
	}
	_ = open
	t.e.refer(v.Sym)
	return mk(operand.Sym(v.Sym).Plus(v.Off)), nil
}

// parseExprOperand reads an operand that is neither a register nor a
// decorated form: a branch target, a bare immediate, or an adrp's page.
func (t *target) parseExprOperand(p *gas.Parser) (any, error) {
	at := p.Peek()
	v, err := p.Expr()
	if err != nil {
		return nil, err
	}
	switch {
	case v.IsAbs():
		return operand.Imm(v.Off), nil
	case v.IsDiff():
		return nil, p.Errorf(at, "%s: a label difference is not an operand", v)
	}
	// A bare symbol in an operand slot is a branch or address target. Which
	// relocation it becomes is the mnemonic's business, and the parent
	// package decides it from the fixup the encoder leaves behind.
	t.e.refer(v.Sym)
	return operand.Direct(operand.Sym(v.Sym).Plus(v.Off)), nil
}

// parseMem reads a bracketed address.
func (t *target) parseMem(p *gas.Parser) (any, error) {
	open := p.Take() // '['

	base := p.Peek()
	if base.Kind != gas.Ident {
		return nil, p.Errorf(base, "want a base register after '[', got %s", base)
	}
	r, ok := reg.Lookup(base.Text)
	if !ok {
		return nil, p.Errorf(base, "%s is not a register", base.Text)
	}
	p.Take()

	var m operand.Mem
	switch b := r.(type) {
	case reg.X:
		m = operand.MemOf(b)
	case reg.Xsp:
		m = operand.MemOf(b)
	default:
		return nil, p.Errorf(base, "%s is not a 64-bit base register", base.Text)
	}

	// [x0]
	if p.AcceptPunct("]") {
		return t.memPost(p, m)
	}
	if err := p.ExpectPunct(","); err != nil {
		return nil, err
	}

	// [x0, #imm] or [x0, #:lo12:sym]
	if p.Peek().IsPunct("#") {
		p.Take()
		if p.Peek().IsPunct(":") {
			ref, err := t.parseAddrRef(p)
			if err != nil {
				return nil, err
			}
			m = m.Off(ref)
		} else {
			at := p.Peek()
			v, err := p.Expr()
			if err != nil {
				return nil, err
			}
			if !v.IsAbs() {
				return nil, p.Errorf(at, "an address offset must be a constant, got %s", v)
			}
			m = m.Off(v.Off)
		}
		if !p.AcceptPunct("]") {
			return nil, p.Errorf(p.Peek(), "want ']' to close the address, got %s", p.Peek())
		}
		// [x0, #8]!
		if p.AcceptPunct("!") {
			return m.Pre(offsetOf(m)), nil
		}
		return m, nil
	}

	// [x0, x1{, lsl #3}] or [x0, w1, sxtw #2]
	idx := p.Peek()
	if idx.Kind != gas.Ident {
		return nil, p.Errorf(idx, "want an index register or an offset, got %s", idx)
	}
	ir, ok := reg.Lookup(idx.Text)
	if !ok {
		return nil, p.Errorf(idx, "%s is not a register", idx.Text)
	}
	p.Take()

	ext, amount := operand.ExtLSL, uint8(0)
	stated := false
	if p.AcceptPunct(",") {
		mod := p.Peek()
		if mod.Kind != gas.Ident {
			return nil, p.Errorf(mod, "want a shift or extend, got %s", mod)
		}
		lower := strings.ToLower(mod.Text)
		switch {
		case lower == "lsl":
			p.Take()
			n, err := t.optAmount(p)
			if err != nil {
				return nil, err
			}
			ext, amount, stated = operand.ExtLSL, n, true
		default:
			e, ok := lookupExtend(lower)
			if !ok {
				return nil, p.Errorf(mod, "%s is not a shift or an extend", mod.Text)
			}
			p.Take()
			n, err := t.optAmount(p)
			if err != nil {
				return nil, err
			}
			ext, amount, stated = e, n, true
		}
	}
	_ = stated
	if !p.AcceptPunct("]") {
		return nil, p.Errorf(p.Peek(), "want ']' to close the address, got %s", p.Peek())
	}
	_ = open
	return m.Indexed(ir, ext, amount), nil
}

// memPost handles the post-index form, whose offset sits outside the
// brackets: [x0], #8.
func (t *target) memPost(p *gas.Parser, m operand.Mem) (any, error) {
	if !p.Peek().IsPunct(",") || !p.PeekAt(1).IsPunct("#") {
		return m, nil
	}
	// Only consume the comma once it is known to introduce a post-index and
	// not the next operand of the instruction.
	p.Take()
	p.Take()
	at := p.Peek()
	v, err := p.Expr()
	if err != nil {
		return nil, err
	}
	if !v.IsAbs() {
		return nil, p.Errorf(at, "a post-index offset must be a constant, got %s", v)
	}
	return m.Post(v.Off), nil
}

// offsetOf reads back the constant displacement a Mem was built with, for the
// pre-index rebuild.
func offsetOf(m operand.Mem) int64 {
	if m.Disp.Sym {
		return 0
	}
	return m.Disp.Const
}
