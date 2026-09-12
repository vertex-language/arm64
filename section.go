package arm64

import (
	"encoding/binary"

	"github.com/vertex-language/arm64/internal/encode"
	"github.com/vertex-language/arm64/internal/isa"
	"github.com/vertex-language/arm64/obj"
	"github.com/vertex-language/arm64/operand"
)

// Section is one section under construction. The typed helpers are methods
// on this type, which is what makes a width mismatch a compile error; an
// interface or a generic builder would erase exactly the checking the
// surface exists for.
type Section struct {
	m *Module

	kind  SectionKind
	name  string
	index int
	align int

	buf  []byte
	refs []obj.Reference

	labels map[string]int

	// comdat is the symbol this section is elected on, and associated
	// the section it follows out of the link; see Module.ComdatSection.
	comdat     string
	associated *Section

	// pending holds every fixup a typed helper or Emit left behind, in
	// placement order, until resolve folds or promotes each one at Finalize.
	pending []pending

	// deltas holds the label differences LabelDiff left behind. Kept
	// apart from pending because a difference names two labels and a
	// fixup names one, and because it is not a relocation candidate: a
	// distance between two offsets in one section is a constant whatever
	// the section's load address turns out to be.
	deltas []labelDelta

	// dead marks the spent handle. Every call on it returns immediately.
	dead bool
}

// pending is one fixup with the mnemonic that produced it. The mnemonic is
// kept beside the fixup rather than on it because encode/ knows nothing of
// object formats or relocations — refKindFor is this package's question, and
// it needs the mnemonic to tell BL's field from B's.
type pending struct {
	fx   encode.Fixup
	mnem string
}

// labelDelta is one four-byte hole holding the distance from one label in
// this section to another.
type labelDelta struct {
	at       int
	from, to string
}

// Module is the module this section belongs to.
//
// It is here for a caller holding a section and needing the module-level
// operations that go with it — Extern and Alias, which an assembler reaches
// for on behalf of text it is assembling into this very section.
func (s *Section) Module() *Module { return s.m }

func (s *Section) Kind() SectionKind { return s.kind }
func (s *Section) Name() string      { return s.name }
func (s *Section) Index() int        { return s.index }

// Comdat is the symbol this section is elected on, or empty for an
// ordinary section. See Module.ComdatSection.
func (s *Section) Comdat() string { return s.comdat }

// Offset is the current end of the section: the offset the next word will
// land at, and the value a Label placed now would name.
func (s *Section) Offset() int { return len(s.buf) }

// ok reports whether the section should do anything. It is the single gate
// every method passes through, which is what makes "sticky and first-wins"
// one rule rather than a rule repeated forty times.
func (s *Section) ok() bool {
	return !s.dead && s.m.err == nil && !s.m.done
}

// ---- instructions ----------------------------------------------------------

// inst is what every typed helper funnels into: encode the form against
// these operands, append the word, and remember any fixup for Finalize.
func (s *Section) inst(f *isa.Form, ops ...any) {
	if !s.ok() {
		return
	}
	if !f.Enabled(s.m.features) {
		s.m.fail(s.lift(&isa.GateError{Form: f, Active: s.m.features}))
		return
	}

	word, fixups, err := encode.EncodeForm(f, ops, encode.Opts{Offset: len(s.buf)})
	if err != nil {
		s.m.fail(s.lift(err))
		return
	}

	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], word)
	s.buf = append(s.buf, b[:]...)
	for _, fx := range fixups {
		s.pending = append(s.pending, pending{fx: fx, mnem: f.Mnem})
	}
}

// ---- labels and symbols -----------------------------------------------------

// Label names an offset in this section.
//
// A bare Label is not a symbol. It gets folded at Finalize wherever a
// same-section reference names it, and leaves no trace in the symbol table.
// Any attribute — a Binding, a SymbolType, a Visibility — promotes it into
// Symbols(), because a page or GOT reference to it needs a symbol to survive
// Finalize as a Reference: the page of an address depends on where the
// section finally loads, which nothing at this layer assigns.
func (s *Section) Label(name string, attrs ...any) {
	if !s.ok() {
		return
	}
	if prev, dup := s.labels[name]; dup {
		s.m.fail(s.errorAt(obj.ErrDuplicate,
			"label "+name+" is already defined in "+s.name,
			"the first definition is at "+s.name+"+"+hex(prev)))
		return
	}
	s.labels[name] = len(s.buf)

	if len(attrs) == 0 {
		return
	}
	s.m.define(s, name, len(s.buf), attrs...)
}

// EndLabel closes a symbol's size range at the current offset.
func (s *Section) EndLabel(name string) {
	if !s.ok() {
		return
	}
	i, ok := s.m.symAt[name]
	if !ok || !s.m.symbols[i].defined {
		s.m.fail(s.errorAt(obj.ErrUndefined,
			"EndLabel("+name+") names no symbol defined in this module"))
		return
	}
	sym := &s.m.symbols[i]
	sym.size = len(s.buf) - sym.off
	sym.sizeClosed = true
}

// ---- alignment and data -----------------------------------------------------

// Align pads to a power-of-two boundary: a code section with D503201F, the
// architecture's NOP, a data section with zeros.
//
// n must be a power of two, and in a code section a multiple of four: every
// instruction here is one word, and an alignment that would strand a partial
// one is ErrAlign rather than rounded.
func (s *Section) Align(n int) {
	if !s.ok() {
		return
	}
	if n <= 0 || n&(n-1) != 0 {
		s.m.fail(s.errorAt(obj.ErrAlign,
			"alignment must be a power of two", "got "+itoa(n)))
		return
	}
	if n > s.align {
		s.align = n
	}
	pad := (n - len(s.buf)%n) % n
	if pad == 0 {
		return
	}
	if s.kind == Text {
		if n%4 != 0 || pad%4 != 0 {
			s.m.fail(s.errorAt(obj.ErrAlign,
				"alignment of a code section must be whole instructions",
				"this alignment would strand a partial word"))
			return
		}
		s.buf = encode.Pad(s.buf, pad)
		return
	}
	s.buf = append(s.buf, make([]byte, pad)...)
}

func (s *Section) Byte(v byte) {
	if s.ok() {
		s.buf = append(s.buf, v)
	}
}

func (s *Section) Long(v uint32) { s.le(uint64(v), 4) }
func (s *Section) Quad(v uint64) { s.le(v, 8) }

func (s *Section) le(v uint64, n int) {
	if !s.ok() {
		return
	}
	for i := 0; i < n; i++ {
		s.buf = append(s.buf, byte(v>>(8*i)))
	}
}

func (s *Section) Ascii(str string) {
	if s.ok() {
		s.buf = append(s.buf, str...)
	}
}

func (s *Section) Asciz(str string) {
	if s.ok() {
		s.buf = append(s.buf, str...)
		s.buf = append(s.buf, 0)
	}
}

func (s *Section) Zero(n int) {
	if s.ok() && n > 0 {
		s.buf = append(s.buf, make([]byte, n)...)
	}
}

func (s *Section) Data(b []byte) {
	if s.ok() {
		s.buf = append(s.buf, b...)
	}
}

// Ref places a data-side hole and a relocation: a pointer in a vtable, a
// jump-table entry that survives to link time. It is the data-side twin of
// what a branch or address operand does with a SymRef, for the layouts this
// package refuses to build for you.
//
// The default kind is RefAbs64 — a plain pointer, the only reading a data
// hole with no instruction around it has. Naming a different kind is for a
// writer-specific layout, such as a COFF .pdata entry.
func (s *Section) Ref(sym string, kind ...obj.RefKind) { s.RefAt(sym, 0, kind...) }

// RefAt is Ref with an addend: the reference names sym plus a constant.
//
// Only the absolute data kinds carry one here. Every other kind names a field
// inside an instruction, where an addend has to travel in a relocation entry
// of its own — ARM64_RELOC_ADDEND on Mach-O — which the writers say they do
// not emit; a caller that asks for one gets the writer's refusal by name
// rather than a silently dropped offset. An absolute pointer is different:
// its addend is folded into the field, which is where every container's
// unsigned relocation reads one from.
func (s *Section) RefAt(sym string, addend int64, kind ...obj.RefKind) {
	if !s.ok() {
		return
	}
	k := obj.RefAbs64
	if len(kind) > 0 {
		k = kind[0]
	}
	size := k.Size()
	if size == 0 {
		size = 8
	}
	s.refs = append(s.refs, obj.Reference{
		Offset: len(s.buf), Size: size, PCRel: k.PCRel(),
		Sym: sym, Kind: k, Addend: addend,
	})
	s.buf = append(s.buf, make([]byte, size)...)
}

// LabelRef places an eight-byte hole patched at Finalize from a same-section
// label. No relocation.
func (s *Section) LabelRef(name string) {
	if !s.ok() {
		return
	}
	s.pending = append(s.pending, pending{fx: encode.Fixup{
		Offset: len(s.buf),
		Target: operand.Label(name),
	}, mnem: "<label-ref>"})
	s.buf = append(s.buf, make([]byte, 8)...)
}

// LabelDiff places a four-byte hole patched at Finalize with the signed
// distance to - from, both labels in this section. The arguments are in the
// order the source expression writes them, which is also the order amd64's
// LabelDiff takes: an assembler reading `.long a - b` passes a then b.
//
// No relocation, and none possible: the difference between two offsets in one
// section is the same number wherever the section loads, which is what makes
// a table of these position-independent where a table of addresses is not. A
// jump table is the reason it exists — Mach-O will not accept an absolute
// pointer into text from a read-only section in a position-independent image,
// and a table of distances needs no such pointer.
func (s *Section) LabelDiff(to, from string) {
	if !s.ok() {
		return
	}
	s.deltas = append(s.deltas, labelDelta{at: len(s.buf), from: from, to: to})
	s.buf = append(s.buf, make([]byte, 4)...)
}

// SymDelta places a four-byte hole holding to - from, where the two are
// symbols this section cannot fold: one of them is in another section, or in
// another object entirely.
//
// LabelDiff above is the foldable case and needs no relocation at all. This
// is the other one, and it is what every descriptor Swift's runtime reads is
// built out of: a relative pointer is `to - from` with from at the field's
// own address. There is no symbol at an arbitrary offset inside a data
// object, so the caller names the object and supplies the field's offset
// within it as a negative addend -- which is what swiftc's own conformance
// descriptors do, entry by entry.
//
// The addend is folded into the field, which is where Mach-O's SUBTRACTOR
// pair reads one from.
func (s *Section) SymDelta(to, from string, addend int64) {
	if !s.ok() {
		return
	}
	if !fitsSigned(addend, 32) {
		s.m.fail(s.errorAt(obj.ErrRange, "symbol delta: addend "+decimal(addend),
			"does not fit 32 signed bits"))
		return
	}
	s.refs = append(s.refs, obj.Reference{
		Offset:     len(s.buf),
		Size:       4,
		Sym:        to,
		Subtrahend: from,
		Kind:       obj.RefDelta32,
		Addend:     addend,
	})
	s.buf = append(s.buf, make([]byte, 4)...)
}

// ---- Finalize machinery -----------------------------------------------------

// resolve runs at Finalize: same-section direct references fold into the
// bytes, everything else that names a symbol this module knows about becomes
// a Reference, and everything else is refused by name.
func (s *Section) resolve() error {
	for _, p := range s.pending {
		switch t := p.fx.Target.(type) {
		case operand.Label:
			if err := s.foldLabel(p, string(t)); err != nil {
				return err
			}
		case operand.SymRef:
			if err := s.promote(p, t); err != nil {
				return err
			}
		default:
			return s.errorAt(obj.ErrUndefined, p.mnem+": reference with no target")
		}
	}
	s.pending = nil

	for _, d := range s.deltas {
		from, ok := s.labels[d.from]
		if !ok {
			return s.errorAt(obj.ErrUndefined,
				"label delta: "+d.from+" is not defined in "+s.name)
		}
		to, ok := s.labels[d.to]
		if !ok {
			return s.errorAt(obj.ErrUndefined,
				"label delta: "+d.to+" is not defined in "+s.name)
		}
		v := int64(to) - int64(from)
		if !fitsSigned(v, 32) {
			return s.errorAt(obj.ErrRange,
				"label delta: "+d.from+" to "+d.to,
				"distance "+decimal(v)+" does not fit 32 signed bits")
		}
		binary.LittleEndian.PutUint32(s.buf[d.at:], uint32(int32(v)))
	}
	s.deltas = nil
	return nil
}

// foldLabel patches a same-section, direct, pc-relative reference into the
// word it belongs to.
func (s *Section) foldLabel(p pending, name string) error {
	target, ok := s.labels[name]
	if !ok {
		return s.errorAt(obj.ErrUndefined,
			p.mnem+": label "+name+" is not defined in "+s.name,
			"labels are per-section; a cross-section or external target needs a symbol and a Ref")
	}

	// LabelRef's marker mnemonic: a raw eight-byte same-section pointer, not
	// a bit-field inside an instruction word.
	if p.mnem == "<label-ref>" {
		binary.LittleEndian.PutUint64(s.buf[p.fx.Offset:], uint64(target))
		return nil
	}

	if p.fx.Role != operand.RoleDirect {
		return s.errorAt(obj.ErrUndefined,
			p.mnem+": label "+name+" needs a "+p.fx.Role.String()+" relocation, and a relocation needs a symbol",
			"promote the label with an attribute: Label(\""+name+"\", arm64.Global)")
	}

	fx := p.fx
	delta := int64(target) - int64(fx.Offset) + fx.Addend
	if fx.Scale > 0 && delta%(1<<fx.Scale) != 0 {
		return s.errorAt(obj.ErrRange,
			p.mnem+": branch to "+name,
			"displacement "+decimal(delta)+" is not a multiple of "+itoa(1<<fx.Scale))
	}
	v := delta >> fx.Scale
	if !fitsSigned(v, fx.Bits) {
		return s.errorAt(obj.ErrRange,
			p.mnem+": branch to "+name,
			"displacement "+decimal(delta)+" does not fit "+itoa(int(fx.Bits))+" signed bits")
	}
	word := binary.LittleEndian.Uint32(s.buf[fx.Offset:])
	word = fx.Field.Put(word, uint64(v)&maskBits(fx.Bits))
	binary.LittleEndian.PutUint32(s.buf[fx.Offset:], word)
	return nil
}

// promote turns a fixup naming a symbol into a Reference, once refKindFor has
// said which relocation the field wants.
func (s *Section) promote(p pending, t operand.SymRef) error {
	i, ok := s.m.symAt[t.Name]
	if !ok {
		return s.errorAt(obj.ErrUndefined,
			p.mnem+": reference to "+t.Name,
			"define the symbol in this module or declare it with Extern")
	}

	// A direct reference to a local label in this very section is a distance
	// this section already knows, and a relocation for it would ask a linker
	// to compute a number that cannot change. Fold it.
	//
	// Local is the whole of the condition, and it is not a nicety: a branch to
	// a *global* in this section must stay a relocation, because another
	// object may interpose a different definition and the linker has to be
	// left able to redirect it. GNU as draws the line in exactly this place.
	if p.fx.Role == operand.RoleDirect && s.m.symbols[i].defined &&
		s.m.symbols[i].binding == Local && s.m.symbols[i].sec == s.index {
		if _, here := s.labels[t.Name]; here {
			q := p
			q.fx.Addend += t.Addend
			return s.foldLabel(q, t.Name)
		}
	}
	kind, err := refKindFor(p.mnem, p.fx)
	if err != nil {
		return s.lift(err)
	}
	s.refs = append(s.refs, obj.Reference{
		Offset: p.fx.Offset, Size: kind.Size(), PCRel: kind.PCRel(),
		Sym: t.Name, Kind: kind, Addend: p.fx.Addend + t.Addend,
	})
	return nil
}

func fitsSigned(v int64, bits uint8) bool {
	if bits == 0 || bits > 63 {
		return false
	}
	lo := -(int64(1) << (bits - 1))
	hi := (int64(1) << (bits - 1)) - 1
	return v >= lo && v <= hi
}

func maskBits(bits uint8) uint64 {
	if bits >= 64 {
		return ^uint64(0)
	}
	return (uint64(1) << bits) - 1
}
