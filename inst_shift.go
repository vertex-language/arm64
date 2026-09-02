package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Variable shift, multiply, divide -----------------------------------

var (
	lslv32 = form("Lslv32")
	lslv64 = form("Lslv64")
	lsrv32 = form("Lsrv32")
	lsrv64 = form("Lsrv64")
	asrv32 = form("Asrv32")
	asrv64 = form("Asrv64")
	rorv32 = form("Rorv32")
	rorv64 = form("Rorv64")
	udiv32 = form("Udiv32")
	udiv64 = form("Udiv64")
	sdiv32 = form("Sdiv32")
	sdiv64 = form("Sdiv64")
	madd32 = form("Madd32")
	madd64 = form("Madd64")
	msub32 = form("Msub32")
	msub64 = form("Msub64")
)

func (s *Section) Lslv32(rd, rn, rm reg.W) { s.inst(lslv32, rd, rn, rm) }
func (s *Section) Lslv64(rd, rn, rm reg.X) { s.inst(lslv64, rd, rn, rm) }
func (s *Section) Lsrv32(rd, rn, rm reg.W) { s.inst(lsrv32, rd, rn, rm) }
func (s *Section) Lsrv64(rd, rn, rm reg.X) { s.inst(lsrv64, rd, rn, rm) }
func (s *Section) Asrv32(rd, rn, rm reg.W) { s.inst(asrv32, rd, rn, rm) }
func (s *Section) Asrv64(rd, rn, rm reg.X) { s.inst(asrv64, rd, rn, rm) }
func (s *Section) Rorv32(rd, rn, rm reg.W) { s.inst(rorv32, rd, rn, rm) }
func (s *Section) Rorv64(rd, rn, rm reg.X) { s.inst(rorv64, rd, rn, rm) }

func (s *Section) Udiv32(rd, rn, rm reg.W) { s.inst(udiv32, rd, rn, rm) }
func (s *Section) Udiv64(rd, rn, rm reg.X) { s.inst(udiv64, rd, rn, rm) }
func (s *Section) Sdiv32(rd, rn, rm reg.W) { s.inst(sdiv32, rd, rn, rm) }
func (s *Section) Sdiv64(rd, rn, rm reg.X) { s.inst(sdiv64, rd, rn, rm) }

// Madd32 emits MADD Wd, Wn, Wm, Wa: Wd = Wa + Wn*Wm.
func (s *Section) Madd32(rd, rn, rm, ra reg.W) { s.inst(madd32, rd, rn, rm, ra) }
func (s *Section) Madd64(rd, rn, rm, ra reg.X) { s.inst(madd64, rd, rn, rm, ra) }

// Msub32 emits MSUB Wd, Wn, Wm, Wa: Wd = Wa - Wn*Wm.
func (s *Section) Msub32(rd, rn, rm, ra reg.W) { s.inst(msub32, rd, rn, rm, ra) }
func (s *Section) Msub64(rd, rn, rm, ra reg.X) { s.inst(msub64, rd, rn, rm, ra) }

// ---- Aliases: shift by register --------------------------------------------
//
// The assembly-language names for the four variable shifts. These pin nothing
// and compute nothing — LSL Wd, Wn, Wm and LSLV Wd, Wn, Wm are one word under
// two spellings — and exist because that is the spelling assembly is written
// in.

var (
	lslReg32 = form("LslReg32")
	lslReg64 = form("LslReg64")
	lsrReg32 = form("LsrReg32")
	lsrReg64 = form("LsrReg64")
	asrReg32 = form("AsrReg32")
	asrReg64 = form("AsrReg64")
	rorReg32 = form("RorReg32")
	rorReg64 = form("RorReg64")
)

func (s *Section) LslReg32(rd, rn, rm reg.W) { s.inst(lslReg32, rd, rn, rm) }
func (s *Section) LslReg64(rd, rn, rm reg.X) { s.inst(lslReg64, rd, rn, rm) }
func (s *Section) LsrReg32(rd, rn, rm reg.W) { s.inst(lsrReg32, rd, rn, rm) }
func (s *Section) LsrReg64(rd, rn, rm reg.X) { s.inst(lsrReg64, rd, rn, rm) }
func (s *Section) AsrReg32(rd, rn, rm reg.W) { s.inst(asrReg32, rd, rn, rm) }
func (s *Section) AsrReg64(rd, rn, rm reg.X) { s.inst(asrReg64, rd, rn, rm) }
func (s *Section) RorReg32(rd, rn, rm reg.W) { s.inst(rorReg32, rd, rn, rm) }
func (s *Section) RorReg64(rd, rn, rm reg.X) { s.inst(rorReg64, rd, rn, rm) }

// ---- Alias: multiply ------------------------------------------------------

var mul64 = form("Mul64")

// Mul64 emits MUL Xd, Xn, Xm — MADD with Xa pinned to XZR.
func (s *Section) Mul64(rd, rn, rm reg.X) { s.inst(mul64, rd, rn, rm) }
