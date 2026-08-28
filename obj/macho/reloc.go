package macho

import (
	"fmt"

	machocore "github.com/vertex-language/macho"
	machoobj "github.com/vertex-language/macho/obj"

	"github.com/vertex-language/arm64/obj"
)

// The intent→ARM64_RELOC_* table.
//
// Unlike x86-64, most of AArch64's Mach-O relocations carry no addend in
// their own entry or in the field: the field the ADRP/ADD/LDR pair writes is
// the page or the page offset, and a linker's answer to "which page" already
// accounts for wherever the two ends of the reference land. A nonzero
// Addend on one of these is only expressible through a preceding
// ARM64_RELOC_ADDEND entry, which this writer does not emit yet — see
// Known limitations in the README — so it is refused here by name rather
// than silently dropped.
//
// RefAbs64/32/16 are the exception: ARM64_RELOC_UNSIGNED is exactly
// X86_64_RELOC_UNSIGNED's model, an implicit addend folded into the field,
// and this writer folds it the same way.
//
// RefPrel64/32/16 have no row: Mach-O's arm64 backend has no plain
// PC-relative data relocation, only the page/branch families below. RefTLV
// is the one thread-local kind Mach-O answers for, through the same
// GOT-style ADRP/LDR pair TLVP_LOAD names rather than the ELF descriptor
// model's seven kinds — this writer does not reach it either, because
// nothing above the reference layer builds the TLV descriptor sequence yet.
type relocForm struct {
	typ       machocore.ARM64Reloc
	pcrel     bool
	pairsOnly bool // true if a nonzero Addend needs ARM64_RELOC_ADDEND
}

var relocTypes = map[obj.RefKind]relocForm{
	obj.RefAbs64: {machocore.ARM64_RELOC_UNSIGNED, false, false},
	obj.RefAbs32: {machocore.ARM64_RELOC_UNSIGNED, false, false},
	obj.RefAbs16: {machocore.ARM64_RELOC_UNSIGNED, false, false},

	obj.RefCall26: {machocore.ARM64_RELOC_BRANCH26, true, true},
	obj.RefJump26: {machocore.ARM64_RELOC_BRANCH26, true, true},

	obj.RefAdrPage21:      {machocore.ARM64_RELOC_PAGE21, true, true},
	obj.RefAddAbsLo12:     {machocore.ARM64_RELOC_PAGEOFF12, false, true},
	obj.RefLdSt8AbsLo12:   {machocore.ARM64_RELOC_PAGEOFF12, false, true},
	obj.RefLdSt16AbsLo12:  {machocore.ARM64_RELOC_PAGEOFF12, false, true},
	obj.RefLdSt32AbsLo12:  {machocore.ARM64_RELOC_PAGEOFF12, false, true},
	obj.RefLdSt64AbsLo12:  {machocore.ARM64_RELOC_PAGEOFF12, false, true},
	obj.RefLdSt128AbsLo12: {machocore.ARM64_RELOC_PAGEOFF12, false, true},

	obj.RefAdrGotPage21: {machocore.ARM64_RELOC_GOT_LOAD_PAGE21, true, true},
	obj.RefLd64GotLo12:  {machocore.ARM64_RELOC_GOT_LOAD_PAGEOFF12, false, true},
}

// writeRelocs translates one section's holes into relocation entries.
func writeRelocs(wr *machoobj.Writer, b *machoobj.SectionBuilder, s *obj.Section, content []byte, syms map[string]machoobj.SymRef) error {
	for _, r := range s.Refs() {
		form, ok := relocTypes[r.Kind]
		if !ok {
			return refKindError(s, r, "Mach-O has no relocation for this kind")
		}
		if want := r.Kind.Size(); want != r.Size {
			return fmt.Errorf("macho: %s+%#x: %s reference to %q is a %d-byte field; %v writes %d",
				s.Name(), r.Offset, r.Kind, r.Sym, r.Size, r.Kind, want)
		}
		length, ok := lengthOf(r.Size)
		if !ok {
			return refKindError(s, r, fmt.Sprintf("a %d-byte field has no r_length", r.Size))
		}

		sym, ok := syms[r.Sym]
		if !ok {
			return fmt.Errorf("macho: %s+%#x: reference to %q, which is not in the symbol table",
				s.Name(), r.Offset, r.Sym)
		}

		if form.pairsOnly && r.Addend != 0 {
			return refKindError(s, r, fmt.Sprintf(
				"a nonzero addend (%d) on %v needs a preceding ARM64_RELOC_ADDEND entry, which this writer does not emit yet",
				r.Addend, form.typ))
		}

		if !form.pairsOnly {
			// RefAbs64/32/16: an implicit addend, folded into the field the
			// same way the amd64 writer folds one into an UNSIGNED field.
			if err := deposit(content, r.Offset, r.Size, r.Addend); err != nil {
				return fmt.Errorf("macho: %s+%#x: %w", s.Name(), r.Offset, err)
			}
		}

		wr.Reloc(b, machoobj.RelocSpec{
			Address: uint64(r.Offset),
			Sym:     sym,
			Type:    uint8(form.typ),
			PCRel:   form.pcrel,
			Length:  length,
		})
	}
	return wr.Err()
}

func lengthOf(size int) (machocore.RelocLength, bool) {
	switch size {
	case 1:
		return machocore.RelocByte, true
	case 2:
		return machocore.RelocWord, true
	case 4:
		return machocore.RelocLong, true
	case 8:
		return machocore.RelocQuad, true
	}
	return 0, false
}

// deposit adds an implicit addend into the section's bytes, for the one
// kind here that carries one: RefAbs64/32/16, whose field is the value and
// nothing else.
func deposit(content []byte, off, size int, addend int64) error {
	if addend == 0 {
		return nil
	}
	if off < 0 || size < 0 || off+size > len(content) {
		return fmt.Errorf("a %d-byte field at %#x does not fit %d bytes of section",
			size, off, len(content))
	}
	var cur uint64
	for i := 0; i < size; i++ {
		cur |= uint64(content[off+i]) << (8 * i)
	}
	v := uint64(int64(cur) + addend)
	for i := 0; i < size; i++ {
		content[off+i] = byte(v >> (8 * i))
	}
	return nil
}

func refKindError(s *obj.Section, r obj.Reference, note string) error {
	return &obj.Error{
		Sentinel: obj.ErrRefKind,
		Arch:     obj.ArchARM64,
		Section:  s.Name(),
		Offset:   r.Offset,
		Context:  fmt.Sprintf("%s reference to %q", r.Kind, r.Sym),
		Notes:    []string{note, "the object is still legal for a container that does answer for this"},
	}
}
