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
