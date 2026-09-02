package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Bitfield ------------------------------------------------------------

var (
	ubfm32 = form("Ubfm32")
	ubfm64 = form("Ubfm64")
	sbfm32 = form("Sbfm32")
	sbfm64 = form("Sbfm64")
	bfm32  = form("Bfm32")
	bfm64  = form("Bfm64")
	extr32 = form("Extr32")
	extr64 = form("Extr64")
)

// Ubfm32 emits UBFM Wd, Wn, #immr, #imms: an unsigned bitfield move, the
// general encoding LSL, LSR and UBFX are all narrower spellings of.
func (s *Section) Ubfm32(rd, rn reg.W, immr, imms uint8) {
	s.inst(ubfm32, rd, rn, uint64(immr), uint64(imms))
}
func (s *Section) Ubfm64(rd, rn reg.X, immr, imms uint8) {
	s.inst(ubfm64, rd, rn, uint64(immr), uint64(imms))
}

// Sbfm32 emits SBFM Wd, Wn, #immr, #imms: a signed bitfield move, ASR and
// SBFX's general encoding.
func (s *Section) Sbfm32(rd, rn reg.W, immr, imms uint8) {
	s.inst(sbfm32, rd, rn, uint64(immr), uint64(imms))
}
func (s *Section) Sbfm64(rd, rn reg.X, immr, imms uint8) {
	s.inst(sbfm64, rd, rn, uint64(immr), uint64(imms))
}

// Bfm32 emits BFM Wd, Wn, #immr, #imms: a bitfield move into Wd that leaves
// the untouched bits of Wd as they were, BFI and BFXIL's general encoding.
func (s *Section) Bfm32(rd, rn reg.W, immr, imms uint8) {
	s.inst(bfm32, rd, rn, uint64(immr), uint64(imms))
}
func (s *Section) Bfm64(rd, rn reg.X, immr, imms uint8) {
	s.inst(bfm64, rd, rn, uint64(immr), uint64(imms))
}

// Extr32 emits EXTR Wd, Wn, Wm, #lsb: the 32-bit word formed by Wn:Wm,
// shifted right by lsb.
func (s *Section) Extr32(rd, rn, rm reg.W, lsb uint8) { s.inst(extr32, rd, rn, rm, uint64(lsb)) }

// Extr64 emits EXTR Xd, Xn, Xm, #lsb.
func (s *Section) Extr64(rd, rn, rm reg.X, lsb uint8) { s.inst(extr64, rd, rn, rm, uint64(lsb)) }

// ---- Aliases: shift by immediate ------------------------------------------

var (
	lslImm64 = form("LslImm64")
	lsrImm64 = form("LsrImm64")
	asrImm64 = form("AsrImm64")
)

// LslImm64 emits LSL Xd, Xn, #shift — UBFM with immr and imms computed from
// shift.
func (s *Section) LslImm64(rd, rn reg.X, shift uint8) { s.inst(lslImm64, rd, rn, uint64(shift)) }

// LsrImm64 emits LSR Xd, Xn, #shift — UBFM with immr = shift, imms = 63.
func (s *Section) LsrImm64(rd, rn reg.X, shift uint8) { s.inst(lsrImm64, rd, rn, uint64(shift)) }

// AsrImm64 emits ASR Xd, Xn, #shift — SBFM with immr = shift, imms = 63.
func (s *Section) AsrImm64(rd, rn reg.X, shift uint8) { s.inst(asrImm64, rd, rn, uint64(shift)) }

var (
	lslImm32 = form("LslImm32")
	lsrImm32 = form("LsrImm32")
	asrImm32 = form("AsrImm32")
)

// LslImm32 emits LSL Wd, Wn, #shift. The immr:imms pair is computed against
// this form's own width, so the same rule serves both.
func (s *Section) LslImm32(rd, rn reg.W, shift uint8) { s.inst(lslImm32, rd, rn, uint64(shift)) }

// LsrImm32 emits LSR Wd, Wn, #shift — UBFM with immr = shift, imms = 31.
func (s *Section) LsrImm32(rd, rn reg.W, shift uint8) { s.inst(lsrImm32, rd, rn, uint64(shift)) }

// AsrImm32 emits ASR Wd, Wn, #shift — SBFM with immr = shift, imms = 31.
func (s *Section) AsrImm32(rd, rn reg.W, shift uint8) { s.inst(asrImm32, rd, rn, uint64(shift)) }

// ---- Aliases: sign and zero extension --------------------------------------
//
// SBFM and UBFM with both immediates fixed, so the width is in the mnemonic
// and there is no immediate to pass. The 64-bit destinations take a W source:
// only the low bits are read, and the architecture spells that in the operand.

var (
	sxtb32 = form("Sxtb32")
	sxtb64 = form("Sxtb64")
	sxth32 = form("Sxth32")
	sxth64 = form("Sxth64")
	sxtw64 = form("Sxtw64")
	uxtb32 = form("Uxtb32")
	uxth32 = form("Uxth32")
)

// Sxtb32 emits SXTB Wd, Wn: the low byte of Wn, sign-extended to 32 bits.
func (s *Section) Sxtb32(rd, rn reg.W) { s.inst(sxtb32, rd, rn) }

// Sxtb64 emits SXTB Xd, Wn.
func (s *Section) Sxtb64(rd reg.X, rn reg.W) { s.inst(sxtb64, rd, rn) }

// Sxth32 emits SXTH Wd, Wn: the low halfword, sign-extended to 32 bits.
func (s *Section) Sxth32(rd, rn reg.W) { s.inst(sxth32, rd, rn) }

// Sxth64 emits SXTH Xd, Wn.
func (s *Section) Sxth64(rd reg.X, rn reg.W) { s.inst(sxth64, rd, rn) }

// Sxtw64 emits SXTW Xd, Wn: a 32-bit word sign-extended to 64.
func (s *Section) Sxtw64(rd reg.X, rn reg.W) { s.inst(sxtw64, rd, rn) }

// Uxtb32 emits UXTB Wd, Wn: the low byte, zero-extended.
func (s *Section) Uxtb32(rd, rn reg.W) { s.inst(uxtb32, rd, rn) }

// Uxth32 emits UXTH Wd, Wn: the low halfword, zero-extended.
func (s *Section) Uxth32(rd, rn reg.W) { s.inst(uxth32, rd, rn) }

// ---- Aliases: bitfield extract ---------------------------------------------

var (
	ubfx32   = form("Ubfx32")
	ubfx64   = form("Ubfx64")
	sbfx32   = form("Sbfx32")
	sbfx64   = form("Sbfx64")
	rorImm32 = form("RorImm32")
	rorImm64 = form("RorImm64")
)

// Ubfx32 emits UBFX Wd, Wn, #lsb, #width: the width bits of Wn starting at
// lsb, zero-extended. UBFM with imms computed as lsb + width - 1.
func (s *Section) Ubfx32(rd, rn reg.W, lsb, width uint8) {
	s.inst(ubfx32, rd, rn, uint64(lsb), uint64(width))
}

// Ubfx64 emits UBFX Xd, Xn, #lsb, #width.
func (s *Section) Ubfx64(rd, rn reg.X, lsb, width uint8) {
	s.inst(ubfx64, rd, rn, uint64(lsb), uint64(width))
}

// Sbfx32 emits SBFX Wd, Wn, #lsb, #width — the same field, sign-extended.
func (s *Section) Sbfx32(rd, rn reg.W, lsb, width uint8) {
	s.inst(sbfx32, rd, rn, uint64(lsb), uint64(width))
}

// Sbfx64 emits SBFX Xd, Xn, #lsb, #width.
func (s *Section) Sbfx64(rd, rn reg.X, lsb, width uint8) {
	s.inst(sbfx64, rd, rn, uint64(lsb), uint64(width))
}

// RorImm32 emits ROR Wd, Ws, #shift — EXTR with Ws as both sources, which the
// row supplies rather than this function.
func (s *Section) RorImm32(rd, rs reg.W, shift uint8) { s.inst(rorImm32, rd, rs, uint64(shift)) }

// RorImm64 emits ROR Xd, Xs, #shift.
func (s *Section) RorImm64(rd, rs reg.X, shift uint8) { s.inst(rorImm64, rd, rs, uint64(shift)) }

// ---- Aliases: bitfield insert ----------------------------------------------
//
// The other direction from extract: a field is placed *at* a position rather
// than taken *from* one, so immr rotates the source into place and imms is
// the width alone. BFXIL is the odd one — it extracts, like UBFX, but into
// the low bits of a destination it leaves otherwise intact, which is why it
// aliases BFM and not UBFM.

var (
	bfi32   = form("Bfi32")
	bfi64   = form("Bfi64")
	bfxil32 = form("Bfxil32")
	bfxil64 = form("Bfxil64")
	ubfiz32 = form("Ubfiz32")
	ubfiz64 = form("Ubfiz64")
	sbfiz32 = form("Sbfiz32")
	sbfiz64 = form("Sbfiz64")
)

// Bfi64 emits BFI Xd, Xn, #lsb, #width: the low width bits of Xn placed at
// lsb in Xd, leaving the rest of Xd alone.
func (s *Section) Bfi32(rd, rn reg.W, lsb, width uint8) {
	s.inst(bfi32, rd, rn, uint64(lsb), uint64(width))
}

func (s *Section) Bfi64(rd, rn reg.X, lsb, width uint8) {
	s.inst(bfi64, rd, rn, uint64(lsb), uint64(width))
}

// Bfxil64 emits BFXIL Xd, Xn, #lsb, #width: the field at lsb in Xn copied
// into the low bits of Xd, leaving the rest of Xd alone.
func (s *Section) Bfxil32(rd, rn reg.W, lsb, width uint8) {
	s.inst(bfxil32, rd, rn, uint64(lsb), uint64(width))
}

func (s *Section) Bfxil64(rd, rn reg.X, lsb, width uint8) {
	s.inst(bfxil64, rd, rn, uint64(lsb), uint64(width))
}

// Ubfiz64 emits UBFIZ Xd, Xn, #lsb, #width: the low width bits of Xn placed
// at lsb, zeroing everything else — a shift and a mask in one instruction.
func (s *Section) Ubfiz32(rd, rn reg.W, lsb, width uint8) {
	s.inst(ubfiz32, rd, rn, uint64(lsb), uint64(width))
}

func (s *Section) Ubfiz64(rd, rn reg.X, lsb, width uint8) {
	s.inst(ubfiz64, rd, rn, uint64(lsb), uint64(width))
}

// Sbfiz64 is the same, sign-extending above the field.
func (s *Section) Sbfiz32(rd, rn reg.W, lsb, width uint8) {
	s.inst(sbfiz32, rd, rn, uint64(lsb), uint64(width))
}

func (s *Section) Sbfiz64(rd, rn reg.X, lsb, width uint8) {
	s.inst(sbfiz64, rd, rn, uint64(lsb), uint64(width))
}
