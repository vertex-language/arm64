package elf

import (
	"fmt"

	elfcore "github.com/vertex-language/elf"
	elfobj "github.com/vertex-language/elf/obj"

	"github.com/vertex-language/arm64/obj"
)

// The intent→R_AARCH64_* table. This is the only place in the tree that
// knows what an ELF relocation number is.
//
// Every entry is a link-semantics decision the lowering already made — which
// half of an address, direct or through the GOT, which TLS model — so
// nothing here chooses between them; it spells what was chosen.
//
// RefTLV has no row: it is Mach-O's descriptor-call model and ELF has no
// relocation for it. RefSecRel32 and RefSecIdx are COFF's. RefSize32 and
// RefSize64 have no row either — the AArch64 ELF psABI does not declare a
// symbol-size relocation the way x86-64's R_X86_64_SIZE32/64 do, so a
// lowering that wants one has nothing to ask ELF for. Each refusal is
// ErrRefKind at emission rather than at construction, because the same
// object is legal for a container that does answer for the kind.
var relocTypes = map[obj.RefKind]elfcore.RelocAArch64{
	obj.RefAbs64:  elfcore.R_AARCH64_ABS64,
	obj.RefAbs32:  elfcore.R_AARCH64_ABS32,
	obj.RefAbs16:  elfcore.R_AARCH64_ABS16,
	obj.RefPrel64: elfcore.R_AARCH64_PREL64,
	obj.RefPrel32: elfcore.R_AARCH64_PREL32,
	obj.RefPrel16: elfcore.R_AARCH64_PREL16,

	obj.RefCall26:   elfcore.R_AARCH64_CALL26,
	obj.RefJump26:   elfcore.R_AARCH64_JUMP26,
	obj.RefCondBr19: elfcore.R_AARCH64_CONDBR19,
	obj.RefTstBr14:  elfcore.R_AARCH64_TSTBR14,

	obj.RefAdrPrel21: elfcore.R_AARCH64_ADR_PREL_LO21,
	obj.RefAdrPage21: elfcore.R_AARCH64_ADR_PREL_PG_HI21,

	obj.RefAddAbsLo12:     elfcore.R_AARCH64_ADD_ABS_LO12_NC,
	obj.RefLdSt8AbsLo12:   elfcore.R_AARCH64_LDST8_ABS_LO12_NC,
	obj.RefLdSt16AbsLo12:  elfcore.R_AARCH64_LDST16_ABS_LO12_NC,
	obj.RefLdSt32AbsLo12:  elfcore.R_AARCH64_LDST32_ABS_LO12_NC,
	obj.RefLdSt64AbsLo12:  elfcore.R_AARCH64_LDST64_ABS_LO12_NC,
	obj.RefLdSt128AbsLo12: elfcore.R_AARCH64_LDST128_ABS_LO12_NC,

	obj.RefAdrGotPage21: elfcore.R_AARCH64_ADR_GOT_PAGE,
	obj.RefLd64GotLo12:  elfcore.R_AARCH64_LD64_GOT_LO12_NC,

	obj.RefTlsGdAdrPage21:         elfcore.R_AARCH64_TLSGD_ADR_PAGE21,
	obj.RefTlsGdAddLo12:           elfcore.R_AARCH64_TLSGD_ADD_LO12_NC,
	obj.RefTlsIeAdrGottprelPage21: elfcore.R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21,
	obj.RefTlsIeLd64GottprelLo12:  elfcore.R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC,
	obj.RefTlsLeAddTprelHi12:      elfcore.R_AARCH64_TLSLE_ADD_TPREL_HI12,
	obj.RefTlsLeAddTprelLo12:      elfcore.R_AARCH64_TLSLE_ADD_TPREL_LO12_NC,
}

// writeRelocs translates one section's holes into RELA entries.
//
// The addend written is Addend + Adjust, which on this architecture is always
// just Addend: Adjust is 0 for every Reference obj/ref.go builds, because
// every PC-relative field here resolves against its own instruction's
// address rather than against the end of it. Nothing is deposited into the
// section bytes — that is the RELA half of the bargain, and the reason this
// writer can hand an object's own bytes straight through while a REL
// container would have to copy.
func writeRelocs(wr *elfobj.Writer, b *elfobj.SectionBuilder, s *obj.Section, syms map[string]elfobj.SymRef) error {
	for _, r := range s.Refs() {
		typ, ok := relocTypes[r.Kind]
		if !ok {
			return &obj.Error{
				Sentinel: obj.ErrRefKind,
				Arch:     obj.ArchARM64,
				Section:  s.Name(),
				Offset:   r.Offset,
				Context:  fmt.Sprintf("%s reference to %q", r.Kind, r.Sym),
				Notes:    []string{"AArch64 ELF has no relocation for this kind"},
			}
		}

		sym, ok := syms[r.Sym]
		if !ok {
			return fmt.Errorf("elf: %s+%#x: reference to %q, which is not in the symbol table",
				s.Name(), r.Offset, r.Sym)
		}

		wr.Reloc(b, elfobj.RelocSpec{
			Offset: uint64(r.Offset),
			Sym:    sym,
			Type:   uint32(typ),
			Addend: r.Addend + r.Adjust,
		})
	}
	return wr.Err()
}
