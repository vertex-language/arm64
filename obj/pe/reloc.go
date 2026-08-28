package pe

import (
	"fmt"

	pecore "github.com/vertex-language/pe"
	"github.com/vertex-language/pe/coff"

	"github.com/vertex-language/arm64/obj"
)

// The intent→IMAGE_REL_ARM64_* table.
//
// There is no PAIR mechanism on this machine — RelocARM64.IsPair and
// TakesPair both always report false — which is simpler than x86-64's ladder
// and Mach-O's ARM64_RELOC_ADDEND: every row here is one relocation record,
// full stop.
//
// The Lo12 family folds to one type. Where ELF and Mach-O need a different
// number per access width — LDST8 through LDST128 — COFF has one,
// PAGEOFFSET_12L, and says so explicitly: "the scaling depends on the
// instruction's access size, so applying this one requires decoding the
// instruction," which is the linker's job and not this table's.
//
// A nonzero Addend on any bit-field kind — every row but the plain absolute
// ones — is refused. None of them stores an addend anywhere a linker reads
// one; the field is the value, and a linker computes it from the symbol and
// the field alone. RefPrel32 is the one PC-relative kind accepted for
// arbitrary data: IMAGE_REL_ARM64_REL32 resolves against the byte following
// the field, the same convention x86-64's REL32 uses, which is why it is the
// one row here needing a +4 correction this package's Adjust — always 0 on
// this architecture — does not carry.
//
// The GOT kinds and every TLS model are refused: there is no GOT on this
// machine, and COFF addresses thread-local data through SECREL against a
// .tls section rather than through any of the models obj declares.
type relocForm struct {
	typ     pecore.RelocARM64
	deposit bool // an implicit addend belongs in the field
	rel32   bool // IMAGE_REL_ARM64_REL32's own +4 bias
}

var relocTypes = map[obj.RefKind]relocForm{
	obj.RefAbs64: {pecore.IMAGE_REL_ARM64_ADDR64, true, false},
	obj.RefAbs32: {pecore.IMAGE_REL_ARM64_ADDR32, true, false},

	obj.RefPrel32: {pecore.IMAGE_REL_ARM64_REL32, true, true},

	obj.RefCall26:   {pecore.IMAGE_REL_ARM64_BRANCH26, false, false},
	obj.RefJump26:   {pecore.IMAGE_REL_ARM64_BRANCH26, false, false},
	obj.RefCondBr19: {pecore.IMAGE_REL_ARM64_BRANCH19, false, false},
	obj.RefTstBr14:  {pecore.IMAGE_REL_ARM64_BRANCH14, false, false},

	obj.RefAdrPrel21: {pecore.IMAGE_REL_ARM64_REL21, false, false},
	obj.RefAdrPage21: {pecore.IMAGE_REL_ARM64_PAGEBASE_REL21, false, false},

	obj.RefAddAbsLo12:     {pecore.IMAGE_REL_ARM64_PAGEOFFSET_12A, false, false},
	obj.RefLdSt8AbsLo12:   {pecore.IMAGE_REL_ARM64_PAGEOFFSET_12L, false, false},
	obj.RefLdSt16AbsLo12:  {pecore.IMAGE_REL_ARM64_PAGEOFFSET_12L, false, false},
	obj.RefLdSt32AbsLo12:  {pecore.IMAGE_REL_ARM64_PAGEOFFSET_12L, false, false},
	obj.RefLdSt64AbsLo12:  {pecore.IMAGE_REL_ARM64_PAGEOFFSET_12L, false, false},
	obj.RefLdSt128AbsLo12: {pecore.IMAGE_REL_ARM64_PAGEOFFSET_12L, false, false},

	obj.RefSecRel32: {pecore.IMAGE_REL_ARM64_SECREL, true, false},
	obj.RefSecIdx:   {pecore.IMAGE_REL_ARM64_SECTION, false, false},
}

// writeRelocs translates one section's holes into COFF relocation records.
func writeRelocs(wr *coff.Writer, b *coff.SectionBuilder, s *obj.Section, content []byte, syms map[string]*coff.SymbolRef) error {
	for _, r := range s.Refs() {
		form, ok := relocTypes[r.Kind]
		if !ok {
			return refKindError(s, r, "COFF has no relocation for this kind")
		}
		if want := r.Kind.Size(); want != r.Size {
			return fmt.Errorf("pe: %s+%#x: %s reference to %q is a %d-byte field; %v writes %d",
				s.Name(), r.Offset, r.Kind, r.Sym, r.Size, r.Kind, want)
		}

		sym, ok := syms[r.Sym]
		if !ok {
			return fmt.Errorf("pe: %s+%#x: reference to %q, which is not in the symbol table",
				s.Name(), r.Offset, r.Sym)
		}

		if !form.deposit && r.Addend != 0 {
			return refKindError(s, r, fmt.Sprintf(
				"%v has no field to hold a nonzero addend (%d); a linker computes this field from the symbol alone",
				form.typ, r.Addend))
		}
		if form.deposit {
			stored := r.Addend
			if form.rel32 {
				// REL32 resolves against the byte following the field; this
				// package's own Adjust is always 0, so the +4 IMAGE_REL_ARM64_REL32
				// bakes into the field is added here rather than carried on
				// the Reference.
				stored += 4
			}
			if err := deposit(content, r.Offset, r.Size, stored); err != nil {
				return fmt.Errorf("pe: %s+%#x: %w", s.Name(), r.Offset, err)
			}
		}

		wr.Reloc(b, coff.RelocSpec{
			Address: uint32(r.Offset),
			Sym:     sym,
			Type:    uint16(form.typ),
		})
	}
	return wr.Err()
}

// deposit adds an implicit addend into the section's bytes.
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
	switch size {
	case 1:
		cur = uint64(int64(int8(cur)))
	case 2:
		cur = uint64(int64(int16(cur)))
	case 4:
		cur = uint64(int64(int32(cur)))
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
