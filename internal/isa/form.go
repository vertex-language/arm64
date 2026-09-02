package isa

import (
	"fmt"
	"strings"

	"github.com/vertex-language/arm64/feature"
)

// Attr is a property of a form that its consumers need and that the slots do
// not carry.
type Attr uint16

const (
	// AttrSetsFlags marks a form that writes NZCV.
	AttrSetsFlags Attr = 1 << iota

	// AttrBranch marks a form whose target field takes a branch relocation.
	AttrBranch

	// AttrScaled marks a load or store whose immediate offset is scaled by the
	// access width, which is what makes the ELF LDST relocation kind depend on
	// the class rather than on the mnemonic.
	AttrScaled

	// AttrPreIndex and AttrPostIndex mark the writeback addressing forms.
	AttrPreIndex
	AttrPostIndex

	// AttrAlias marks a form that is an alias of another: one mnemonic that is
	// exactly another with some field fixed. See alias.go.
	AttrAlias

	// AttrNoZR marks a form in which register 31 in a slot means something
	// other than the zero register even though the slot's class says X or W.
	AttrNoZR

	// AttrInvertCond marks a form whose condition field holds the inverse of
	// the condition written in source.
	//
	// CSET is the reason. It is CSINC with both sources pinned to the zero
	// register, and CSINC increments when its condition is *false* — so
	// "cset x0, eq" has to encode NE, or it would set x0 when the comparison
	// was not equal. That is a fact about the encoding, so it belongs on the
	// row, where both doors onto the encoder pass through it. It lived in the
	// typed helper until an assembler reached the same row without going
	// through that helper and encoded the condition uninverted — the two
	// doors producing different instructions for one mnemonic, which is the
	// one thing sharing an encoder is supposed to make impossible.
	AttrInvertCond
)

// Form is one declared encoding of one mnemonic.
type Form struct {
	// Mnem is the mnemonic, lower case, as the architecture spells it.
	Mnem string

	// Word is the base encoding with every operand field zero.
	Word uint32

	// Mask is the bits of Word that are fixed. A word w decodes to this form
	// when w&Mask == Word.
	Mask uint32

	// Slots are the operands in assembly-source order, which for A64 is
	// destination first.
	Slots []Slot

	// Gate is the extension required to encode this form, or feature.None for
	// the base instruction set.
	Gate feature.Feature

	// Attrs are the form's properties.
	Attrs Attr

	// name is the Go identifier of the generated helper, filled by finish.
	name string

	// alias links an alias form to what it is an alias of. See alias.go.
	alias *aliasOf
}

// GoName is the identifier of this form's generated helper method: AddImm64,
// AddShifted64, LdrImm64. It is derived once at table build time and stored,
// because a name computed at call sites could differ between the generator and
// a diagnostic.
func (f *Form) GoName() string { return f.name }

// Enabled reports whether a feature set permits this form.
func (f *Form) Enabled(s feature.Set) bool {
	return f.Gate == feature.None || s.Has(f.Gate)
}

// Arity is the number of operands, including optional ones.
func (f *Form) Arity() int { return len(f.Slots) }

// Required is the number of operands that must be written.
func (f *Form) Required() int {
	n := 0
	for _, s := range f.Slots {
		if !s.Optional {
			n++
		}
	}
	return n
}

// AccessBits is the width of the memory access, or 0 for a form that touches
// no memory. The encoder records it on the Fixup so the platform writer can
// pick between LDST8 and LDST128 without the caller ever naming a width.
func (f *Form) AccessBits() uint16 {
	for _, s := range f.Slots {
		if s.Class.Mem() {
			return s.Class.AccessBits()
		}
	}
	return 0
}

// Signature is the form's operand classes, which is what finish checks for
// collisions.
//
// The writeback attrs are appended because they are the one property that
// changes what a mnemonic's operands mean without changing a single Slot's
// class: Stp64, StpPre64 and LdpPost64 all take (Xt, Xt2, Mem64) and would
// collide here otherwise. Resolve does not yet make this same distinction —
// see the note on AddrForm in class.go — so it stays a table-build-time check
// that guards against two forms the typed helpers could never tell apart,
// pending Resolve gaining the same discrimination for Emit's sake.
func (f *Form) Signature() string {
	var b strings.Builder
	b.WriteString(f.Mnem)
	for _, s := range f.Slots {
		b.WriteByte(' ')
		b.WriteString(s.Class.String())
		if s.Optional {
			b.WriteByte('?')
		}
	}
	switch {
	case f.Attrs&AttrPreIndex != 0:
		b.WriteString(" pre")
	case f.Attrs&AttrPostIndex != 0:
		b.WriteString(" post")
	}
	return b.String()
}

func (f *Form) String() string { return f.Signature() }

// finish runs the table-time checks and derives the Go name. It panics rather
// than returning an error: the table is a package-level literal, a broken row
// is a build-time bug, and a bad encoder is worse than a failed build.
func (f *Form) finish() {
	if f.Mnem == "" {
		panic("isa: form with no mnemonic")
	}

	// The base word must not have bits set outside the mask, and every operand
	// field must lie outside the mask. A field overlapping a fixed bit means
	// the row disagrees with itself about whether that bit is an opcode.
	if f.Word&^f.Mask != 0 {
		panic(fmt.Sprintf("isa: %s: base word %#08x sets bits outside mask %#08x",
			f.Mnem, f.Word, f.Mask))
	}
	var covered uint32
	for i, s := range f.Slots {
		m := s.Field.Mask()
		if m == 0 && s.Class != ClassNone {
			continue // a slot with no field is fixed by the base word
		}
		if m&f.Mask != 0 {
			panic(fmt.Sprintf("isa: %s: operand %d field %s overlaps fixed bits",
				f.Mnem, i, s.Field))
		}
		if m&covered != 0 {
			panic(fmt.Sprintf("isa: %s: operand %d field %s overlaps another operand",
				f.Mnem, i, s.Field))
		}
		covered |= m
	}

	// An optional slot may not precede a required one; the syntax has no way
	// to omit the first of two operands and supply the second.
	seenOptional := false
	for i, s := range f.Slots {
		if s.Optional {
			seenOptional = true
			continue
		}
		if seenOptional {
			panic(fmt.Sprintf("isa: %s: required operand %d follows an optional one", f.Mnem, i))
		}
	}

	if f.name == "" {
		f.name = goFormName(f)
	}
}
