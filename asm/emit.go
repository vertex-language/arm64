package asm

import (
	"fmt"
	"sort"

	"github.com/vertex-language/arm64"
	"github.com/vertex-language/arm64/obj"
	"github.com/vertex-language/asm/gas"
)

// emitter adapts a Module to gas.Emitter.
//
// It is small on purpose, and it is small because the builder API underneath
// it already answers every question a directive asks. `.quad` is Quad,
// `.pushsection` is SectionNamed, `.size` is EndLabel. The only real work is
// the two symbol-reference cases, where a width has to become a relocation
// kind — and that is precisely the mapping gas cannot make, because RefKind is
// declared once per architecture and gas is not one.
type emitter struct {
	m   *arm64.Module
	sec *arm64.Section

	// defined and referred are what makes an implicit Extern possible.
	//
	// GNU as treats a name it never sees defined as one this object imports,
	// and every hand-written .s relies on it: nothing writes `.extern puts`
	// before calling puts. The parent package refuses an undefined reference
	// on purpose — a typo in a symbol name is otherwise a link error a long
	// way from its cause — so the two are reconciled here, at the end of the
	// parse, when what was never defined is finally known.
	defined  map[string]bool
	referred map[string]bool
}

func newEmitter(m *arm64.Module) *emitter {
	return &emitter{m: m, defined: map[string]bool{}, referred: map[string]bool{}}
}

// refer records a symbol this assembly names.
func (e *emitter) refer(name string) { e.referred[name] = true }

// declareExterns declares everything referred to and never defined.
func (e *emitter) declareExterns() {
	names := make([]string, 0, len(e.referred))
	for n := range e.referred {
		if !e.defined[n] {
			names = append(names, n)
		}
	}
	sort.Strings(names) // a stable symbol table beats an incidental one
	for _, n := range names {
		e.m.Extern(n)
	}
}

func (e *emitter) Section(name string, k gas.SectionKind) error {
	e.sec = e.m.SectionNamed(name, sectionKind(k))
	return e.m.Err()
}

func sectionKind(k gas.SectionKind) arm64.SectionKind {
	switch k {
	case gas.Text:
		return arm64.Text
	case gas.ROData:
		return arm64.ROData
	case gas.BSS:
		return arm64.BSS
	}
	return arm64.Data
}

func (e *emitter) Offset() int { return e.sec.Offset() }

// Label names the current offset.
//
// Every label becomes a symbol, including the plain ones the parent package
// would otherwise fold and forget. That is not the parent package being wrong:
// a label nothing takes the address of has no reason to reach the symbol
// table, and folding it is the right default for a caller who is emitting
// instructions and knows which labels are targets.
//
// A parser does not know. `adrp x0, msg` names a page, and a page depends on
// where the section finally loads, so the reference has to survive Finalize as
// a Reference and a Reference needs a symbol. Whether a given label is ever
// used that way is not known when the label is emitted — it may be referenced
// hundreds of lines later — so the choice is to promote every label or to walk
// the token stream twice. Promoting costs a local symbol per label, which a
// linker is free to drop; guessing costs a build that fails at Finalize for a
// reason nothing in the source suggests.
func (e *emitter) Label(name string, a gas.Attrs) error {
	e.defined[name] = true
	attrs := []any{binding(a.Binding)}
	switch a.Type {
	case gas.TypeFunc:
		attrs = append(attrs, arm64.Func)
	case gas.TypeObject:
		attrs = append(attrs, arm64.ObjectSym)
	}
	switch a.Vis {
	case gas.VisHidden:
		attrs = append(attrs, arm64.Hidden)
	case gas.VisProtected:
		attrs = append(attrs, arm64.Protected)
	}
	e.sec.Label(name, attrs...)
	return e.m.Err()
}

func binding(b gas.Binding) arm64.Binding {
	switch b {
	case gas.BindGlobal:
		return arm64.Global
	case gas.BindWeak:
		return arm64.Weak
	}
	return arm64.Local
}

func (e *emitter) EndLabel(name string) error {
	e.sec.EndLabel(name)
	return e.m.Err()
}

func (e *emitter) Extern(name string) {
	e.defined[name] = true
	e.m.Extern(name)
}

func (e *emitter) Alias(name, of string) {
	e.defined[name] = true
	e.refer(of)
	e.m.Alias(name, of)
}

func (e *emitter) Data(b []byte) { e.sec.Data(b) }

func (e *emitter) Int(v int64, size int) error {
	switch size {
	case 1:
		e.sec.Byte(byte(v))
	case 2:
		e.sec.Byte(byte(v))
		e.sec.Byte(byte(v >> 8))
	case 4:
		e.sec.Long(uint32(v))
	case 8:
		e.sec.Quad(uint64(v))
	default:
		return fmt.Errorf("no %d-byte integer directive", size)
	}
	return e.m.Err()
}

func (e *emitter) Zero(n int) { e.sec.Zero(n) }

func (e *emitter) Align(n int) error {
	e.sec.Align(n)
	return e.m.Err()
}

// SymRef places an absolute reference to a symbol.
//
// A four- or eight-byte hole in data is the only place an absolute relocation
// belongs: in text, an address is built by adrp/add or read from the GOT, and
// a data-style reference there would be a word the CPU tries to execute.
func (e *emitter) SymRef(sym string, addend int64, size int) error {
	var kind obj.RefKind
	switch size {
	case 8:
		kind = arm64.RefAbs64
	case 4:
		kind = arm64.RefAbs32
	case 2:
		kind = arm64.RefAbs16
	default:
		return fmt.Errorf("a %d-byte symbol reference has no relocation", size)
	}
	e.refer(sym)
	e.sec.Ref(sym, kind)
	if addend != 0 {
		return fmt.Errorf("%s%+d: an addend on a data reference is not supported yet", sym, addend)
	}
	return e.m.Err()
}

// SymDiff places the distance between two labels, which needs no relocation
// at all — see Section.LabelDiff.
func (e *emitter) SymDiff(to, from string, addend int64, size int) error {
	if size != 4 {
		return fmt.Errorf("a label difference is four bytes, not %d", size)
	}
	if addend != 0 {
		return fmt.Errorf("%s-%s%+d: an addend on a label difference is not supported yet",
			to, from, addend)
	}
	e.refer(to)
	e.refer(from)
	e.sec.LabelDiff(to, from)
	return e.m.Err()
}
