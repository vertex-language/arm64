package obj

// RefKind states how a linker should resolve a reference.
//
// Names follow the ELF psABI's own relocation names — R_AARCH64_ADR_PREL_PG_HI21
// becomes RefAdrPage21 — because that vocabulary is what every AArch64 ABI
// document, disassembler and linker error message already uses, and a second
// naming scheme for the same sixteen ideas would only cost a reader the
// translation. Mach-O and COFF get their own kind where their model has no ELF
// counterpart: the RefTlv pair for Mach-O's descriptor-call TLS, RefSecRel32 and
// RefSecIdx for COFF's section-relative debug forms.
//
// The set is the union of what the three containers can express, not the
// intersection. An intersection would drop the page/page-offset pairing that
// is how every position-independent AArch64 address is built, and make every
// interesting object unbuildable; a per-container set would mean a lowering
// picks its container before it picks its instructions. So it is the union,
// every writer states what it cannot do, and ErrRefKind names the kind and
// the offset when a kind meets a container with no answer for it.
//
// It is declared once, here, and aliased upward. arm64.RefAdrPage21,
// operand.RefAdrPage21 and obj.RefAdrPage21 are the same constant and no
// conversion exists anywhere in the tree.
type RefKind uint8

const (
	RefNone RefKind = iota

	// Absolute data, by field width — a pointer in a data section, never an
	// instruction's immediate. AArch64 has no instruction that takes a wide
	// absolute immediate directly; that is what the page/page-offset kinds
	// below are for.
	RefAbs64
	RefAbs32
	RefAbs16

	// PC-relative data, by field width. Uncommon — a same-section table of
	// offsets is ordinarily a LabelDiff, needing no relocation at all — but
	// declared for the case where the subtrahend is not known until link
	// time.
	RefPrel64
	RefPrel32
	RefPrel16

	// The branch and test fields. Each names a specific instruction's
	// immediate, because the field width and scale are fixed by which
	// instruction it is, not by anything a Reference can add.
	//
	// RefCall26 and RefJump26 share one bit layout — a 26-bit word-scaled
	// signed displacement — and differ only in what a linker is permitted to
	// do at the far end: JUMP26 (a plain B) may be redirected to a long-branch
	// thunk or interworking stub as a tail call, where CALL26 (BL) must return
	// through the same instruction that reached it. Naming the wrong one
	// tells a linker the wrong thing about a return address.
	RefCall26   // BL, or B to an external symbol via a PLT stub
	RefJump26   // B
	RefCondBr19 // B.cond, CBZ, CBNZ
	RefTstBr14  // TBZ, TBNZ

	// ADR and ADRP. ADR's immediate is a 21-bit byte displacement to an
	// address anywhere in the +/-1MiB window; ADRP's is the same 21 bits
	// shifted 12 to name a 4KiB page up to +/-4GiB away, which is why
	// position-independent code pairs an ADRP with one of the Lo12 kinds
	// below rather than materializing an address in one instruction.
	RefAdrPrel21
	RefAdrPage21

	// The Lo12 family completes an ADRP: the page from RefAdrPage21, the
	// offset within it from one of these, added by ADD or folded into an
	// addressing mode by LDR/STR. Each LDST kind is named for the access
	// width because the immediate is scaled by it — LDR (64-bit) shifts the
	// page offset right by 3 before checking it fits, LDR (8-bit) does not
	// shift at all — and a linker applying the wrong one silently computes
	// the wrong address rather than failing.
	RefAddAbsLo12     // ADD (immediate), unscaled
	RefLdSt8AbsLo12   // LDRB/STRB, unscaled
	RefLdSt16AbsLo12  // LDRH/STRH, scaled by 2
	RefLdSt32AbsLo12  // LDR/STR (32-bit), scaled by 4
	RefLdSt64AbsLo12  // LDR/STR (64-bit), scaled by 8
	RefLdSt128AbsLo12 // LDR/STR (128-bit, vector), scaled by 16

	// The GOT pair: an ADRP to the page holding this symbol's GOT entry, and
	// an LDR (always 64-bit — a GOT entry is a pointer) of the entry itself.
	// This is how a PIC/PIE object reaches a symbol that may resolve outside
	// the final image, and it is two relocations and two instructions where
	// x86-64 needs one RIP-relative load, because ADRP+LDR is how every
	// PC-relative address on this architecture is built.
	RefAdrGotPage21
	RefLd64GotLo12

	// Thread-local storage. AArch64 ELF states three models, each an ADRP
	// paired with a second instruction; which pair a lowering picks is a
	// decision about how the symbol will be accessed, and the two halves of
	// one access always carry the same model.
	//
	// General dynamic: call __tls_get_addr with the module and offset this
	// pair computes the address of.
	RefTlsGdAdrPage21
	RefTlsGdAddLo12

	// Initial exec: the offset from the thread pointer sits in the GOT, read
	// through this pair the same way RefAdrGotPage21/RefLd64GotLo12 read an
	// ordinary GOT entry.
	RefTlsIeAdrGottprelPage21
	RefTlsIeLd64GottprelLo12

	// Local exec: the offset from the thread pointer is a link-time constant,
	// so no GOT indirection is needed — just the constant's high and low
	// twelve bits, each added with its own ADD.
	RefTlsLeAddTprelHi12
	RefTlsLeAddTprelLo12

	// Mach-O's, and a pair rather than one kind.
	//
	// A thread-local there is reached through a descriptor and a call
	// rather than through a relocation model, so where ELF has seven kinds
	// above this has two — but it does need two. The address of the
	// descriptor is built the way every other PC-relative address on this
	// architecture is built, an ADRP and then the low twelve bits, and the
	// container names the halves separately: ARM64_RELOC_TLVP_LOAD_PAGE21
	// and ARM64_RELOC_TLVP_LOAD_PAGEOFF12. One kind could not say which
	// instruction it was on.
	//
	// The LOAD in those names is the unrelaxed form, where the low half is
	// an LDR through __thread_ptrs. A linker that finds the descriptor in
	// the image it is building rewrites the LDR as an ADD and addresses it
	// directly; the relocation is the same either way, and which one it
	// became is the linker's business rather than this layer's.
	RefAdrTlvPage21
	RefLdTlvLo12

	// Size and COFF's section-relative forms, carried for parity with the
	// other architectures in this tree; a lowering that never emits debug
	// info or symbol-size requests never produces one.
	RefSize32
	RefSize64
	RefSecRel32
	RefSecIdx

	numRefKinds
)

var refNames = [numRefKinds]string{
	RefNone: "none",

	RefAbs64: "abs64", RefAbs32: "abs32", RefAbs16: "abs16",
	RefPrel64: "prel64", RefPrel32: "prel32", RefPrel16: "prel16",

	RefCall26: "call26", RefJump26: "jump26",
	RefCondBr19: "condbr19", RefTstBr14: "tstbr14",

	RefAdrPrel21: "adr-prel21", RefAdrPage21: "adr-page21",

	RefAddAbsLo12:     "add-abs-lo12",
	RefLdSt8AbsLo12:   "ldst8-abs-lo12",
	RefLdSt16AbsLo12:  "ldst16-abs-lo12",
	RefLdSt32AbsLo12:  "ldst32-abs-lo12",
	RefLdSt64AbsLo12:  "ldst64-abs-lo12",
	RefLdSt128AbsLo12: "ldst128-abs-lo12",

	RefAdrGotPage21: "adr-got-page21", RefLd64GotLo12: "ld64-got-lo12",

	RefTlsGdAdrPage21: "tlsgd-adr-page21", RefTlsGdAddLo12: "tlsgd-add-lo12",
	RefTlsIeAdrGottprelPage21: "tlsie-adr-gottprel-page21",
	RefTlsIeLd64GottprelLo12:  "tlsie-ld64-gottprel-lo12",
	RefTlsLeAddTprelHi12:      "tlsle-add-tprel-hi12",
	RefTlsLeAddTprelLo12:      "tlsle-add-tprel-lo12",

	RefAdrTlvPage21: "tlv-adr-page21",
	RefLdTlvLo12:    "tlv-ld-lo12",

	RefSize32: "size32", RefSize64: "size64",
	RefSecRel32: "secrel32", RefSecIdx: "secidx",
}

// String is what ErrRefKind names when a writer refuses one.
func (k RefKind) String() string {
	if int(k) < len(refNames) {
		return refNames[k]
	}
	return "refkind?"
}

// Valid reports whether the value names a declared kind. RefNone is not one.
func (k RefKind) Valid() bool { return k > RefNone && k < numRefKinds }

// Size is the width in bytes of the hole the kind fills.
//
// For the data kinds this is the field itself: a RefAbs64 pointer is eight
// bytes wherever it appears. For every kind that names a bit-field inside an
// instruction — every branch, ADR/ADRP, and Lo12 or TLS kind — the hole is the
// one instruction word the field lives in, four bytes, and the bit position
// and width within that word are fixed facts about the kind that a writer
// looks up rather than a number this method could usefully return.
func (k RefKind) Size() int {
	switch k {
	case RefAbs64, RefPrel64, RefSize64:
		return 8
	case RefAbs32, RefPrel32, RefSize32, RefSecRel32:
		return 4
	case RefAbs16, RefPrel16, RefSecIdx:
		return 2
	case RefNone:
		return 0
	}
	// Every remaining kind names a field inside a single AArch64 instruction.
	return 4
}

// PCRel reports whether the kind resolves against the position of its own
// instruction rather than against an absolute or page-relative address.
//
// The ADRP-paired Lo12 and TLS-offset kinds are not PC-relative even though
// they only ever appear beside a PCRel ADRP: each names an offset within a
// page or a TLS block, a quantity with no dependency on where the field
// itself sits.
func (k RefKind) PCRel() bool {
	switch k {
	case RefPrel64, RefPrel32, RefPrel16,
		RefCall26, RefJump26, RefCondBr19, RefTstBr14,
		RefAdrPrel21, RefAdrPage21, RefAdrGotPage21,
		RefTlsGdAdrPage21, RefTlsIeAdrGottprelPage21:
		return true
	}
	return false
}

// TLS reports whether the kind names a thread-local storage model. The
// writers each accept a different subset and refuse the rest, so asking the
// question in one place keeps three answers from drifting.
func (k RefKind) TLS() bool {
	switch k {
	case RefTlsGdAdrPage21, RefTlsGdAddLo12,
		RefTlsIeAdrGottprelPage21, RefTlsIeLd64GottprelLo12,
		RefTlsLeAddTprelHi12, RefTlsLeAddTprelLo12,
		RefAdrTlvPage21, RefLdTlvLo12:
		return true
	}
	return false
}

// Reference is one hole a linker fills.
//
// The identity every consumer relies on is
//
//	value = target - (section offset of the field) + Adjust + Addend
//
// Adjust exists on the other architectures in this tree for a PC-relative
// field that does not resolve against its own start — x86-64's displacement
// resolves against the end of the instruction, wherever that falls. Nothing
// on AArch64 works that way: every PC-relative field here, ADRP's included,
// resolves against the address of its own instruction word, so a Reference
// built by this package always carries Adjust == 0. The field is kept rather
// than dropped so a writer shared textually with the other obj packages does
// not need a special case for the one architecture that never sets it.
//
// Size is 2, 4 or 8 for the data kinds and always 4 — one instruction word —
// for every kind that names a field inside an instruction; RefKind.Size
// answers this the same way a constructed Reference does.
type Reference struct {
	Offset int // where the hole starts, section-relative
	Size   int
	PCRel  bool
	Adjust int64 // always 0 on this architecture; see above
	Sym    string
	Kind   RefKind
	Addend int64 // logical addend, never adjusted for the field
}
