package arm64

import "github.com/vertex-language/arm64/reg"

// Scalar floating point, single and double precision.
//
// Two helpers per operation rather than one taking a width, for the reason the
// integer set has AddImm32 and AddImm64: the register types differ, and a width
// mismatch should be a compile error at the call rather than a runtime check.

var (
	faddS, faddD     = form("FaddS"), form("FaddD")
	fsubS, fsubD     = form("FsubS"), form("FsubD")
	fmulS, fmulD     = form("FmulS"), form("FmulD")
	fdivS, fdivD     = form("FdivS"), form("FdivD")
	fmaxS, fmaxD     = form("FmaxS"), form("FmaxD")
	fminS, fminD     = form("FminS"), form("FminD")
	fmaxnmS, fmaxnmD = form("FmaxnmS"), form("FmaxnmD")
	fminnmS, fminnmD = form("FminnmS"), form("FminnmD")
)

// FaddS emits FADD Sd, Sn, Sm.
func (s *Section) FaddS(rd, rn, rm reg.S) { s.inst(faddS, rd, rn, rm) }

// FaddD emits FADD Dd, Dn, Dm.
func (s *Section) FaddD(rd, rn, rm reg.D) { s.inst(faddD, rd, rn, rm) }

func (s *Section) FsubS(rd, rn, rm reg.S) { s.inst(fsubS, rd, rn, rm) }
func (s *Section) FsubD(rd, rn, rm reg.D) { s.inst(fsubD, rd, rn, rm) }
func (s *Section) FmulS(rd, rn, rm reg.S) { s.inst(fmulS, rd, rn, rm) }
func (s *Section) FmulD(rd, rn, rm reg.D) { s.inst(fmulD, rd, rn, rm) }
func (s *Section) FdivS(rd, rn, rm reg.S) { s.inst(fdivS, rd, rn, rm) }
func (s *Section) FdivD(rd, rn, rm reg.D) { s.inst(fdivD, rd, rn, rm) }

// FmaxS emits FMAX Sd, Sn, Sm, which propagates a NaN operand. FmaxnmS is the
// one that returns the other operand instead, and is what C's fmaxf wants.
func (s *Section) FmaxS(rd, rn, rm reg.S) { s.inst(fmaxS, rd, rn, rm) }
func (s *Section) FmaxD(rd, rn, rm reg.D) { s.inst(fmaxD, rd, rn, rm) }
func (s *Section) FminS(rd, rn, rm reg.S) { s.inst(fminS, rd, rn, rm) }
func (s *Section) FminD(rd, rn, rm reg.D) { s.inst(fminD, rd, rn, rm) }

// FmaxnmS emits FMAXNM Sd, Sn, Sm, which returns the non-NaN operand.
func (s *Section) FmaxnmS(rd, rn, rm reg.S) { s.inst(fmaxnmS, rd, rn, rm) }
func (s *Section) FmaxnmD(rd, rn, rm reg.D) { s.inst(fmaxnmD, rd, rn, rm) }
func (s *Section) FminnmS(rd, rn, rm reg.S) { s.inst(fminnmS, rd, rn, rm) }
func (s *Section) FminnmD(rd, rn, rm reg.D) { s.inst(fminnmD, rd, rn, rm) }

// ---- One source -----------------------------------------------------------

var (
	fmovS, fmovD     = form("FmovS"), form("FmovD")
	fabsS, fabsD     = form("FabsS"), form("FabsD")
	fnegS, fnegD     = form("FnegS"), form("FnegD")
	fsqrtS, fsqrtD   = form("FsqrtS"), form("FsqrtD")
	frintnS, frintnD = form("FrintnS"), form("FrintnD")
	frintpS, frintpD = form("FrintpS"), form("FrintpD")
	frintmS, frintmD = form("FrintmS"), form("FrintmD")
	frintzS, frintzD = form("FrintzS"), form("FrintzD")
)

// FmovS emits FMOV Sd, Sn: a register-to-register move within the vector file.
func (s *Section) FmovS(rd, rn reg.S) { s.inst(fmovS, rd, rn) }
func (s *Section) FmovD(rd, rn reg.D) { s.inst(fmovD, rd, rn) }

// FabsS emits FABS Sd, Sn: the sign bit cleared, NaN payload preserved.
func (s *Section) FabsS(rd, rn reg.S) { s.inst(fabsS, rd, rn) }
func (s *Section) FabsD(rd, rn reg.D) { s.inst(fabsD, rd, rn) }

// FnegS emits FNEG Sd, Sn: the sign bit flipped, which is not a subtraction
// from zero and differs from one on a zero and on a NaN.
func (s *Section) FnegS(rd, rn reg.S) { s.inst(fnegS, rd, rn) }
func (s *Section) FnegD(rd, rn reg.D) { s.inst(fnegD, rd, rn) }

func (s *Section) FsqrtS(rd, rn reg.S) { s.inst(fsqrtS, rd, rn) }
func (s *Section) FsqrtD(rd, rn reg.D) { s.inst(fsqrtD, rd, rn) }

// The four rounding modes that are their own instruction: nearest-even, toward
// +inf, toward -inf and toward zero.
func (s *Section) FrintnS(rd, rn reg.S) { s.inst(frintnS, rd, rn) }
func (s *Section) FrintnD(rd, rn reg.D) { s.inst(frintnD, rd, rn) }
func (s *Section) FrintpS(rd, rn reg.S) { s.inst(frintpS, rd, rn) }
func (s *Section) FrintpD(rd, rn reg.D) { s.inst(frintpD, rd, rn) }
func (s *Section) FrintmS(rd, rn reg.S) { s.inst(frintmS, rd, rn) }
func (s *Section) FrintmD(rd, rn reg.D) { s.inst(frintmD, rd, rn) }
func (s *Section) FrintzS(rd, rn reg.S) { s.inst(frintzS, rd, rn) }
func (s *Section) FrintzD(rd, rn reg.D) { s.inst(frintzD, rd, rn) }

// ---- Width conversion ------------------------------------------------------

var (
	fcvtSToD = form("FcvtSToD")
	fcvtDToS = form("FcvtDToS")
)

// FcvtSToD emits FCVT Dd, Sn: exact, every single being a double.
func (s *Section) FcvtSToD(rd reg.D, rn reg.S) { s.inst(fcvtSToD, rd, rn) }

// FcvtDToS emits FCVT Sd, Dn, which rounds by FPCR.
func (s *Section) FcvtDToS(rd reg.S, rn reg.D) { s.inst(fcvtDToS, rd, rn) }

// ---- Three source ----------------------------------------------------------

var (
	fmaddS, fmaddD = form("FmaddS"), form("FmaddD")
	fmsubS, fmsubD = form("FmsubS"), form("FmsubD")
)

// FmaddS emits FMADD Sd, Sn, Sm, Sa: Sd = Sa + Sn*Sm with one rounding.
func (s *Section) FmaddS(rd, rn, rm, ra reg.S) { s.inst(fmaddS, rd, rn, rm, ra) }
func (s *Section) FmaddD(rd, rn, rm, ra reg.D) { s.inst(fmaddD, rd, rn, rm, ra) }

// FmsubS emits FMSUB Sd, Sn, Sm, Sa: Sd = Sa - Sn*Sm with one rounding.
func (s *Section) FmsubS(rd, rn, rm, ra reg.S) { s.inst(fmsubS, rd, rn, rm, ra) }
func (s *Section) FmsubD(rd, rn, rm, ra reg.D) { s.inst(fmsubD, rd, rn, rm, ra) }

// ---- Compare and select ----------------------------------------------------

var (
	fcmpS, fcmpD         = form("FcmpS"), form("FcmpD")
	fcmpZeroS, fcmpZeroD = form("FcmpZeroS"), form("FcmpZeroD")
	fcselS, fcselD       = form("FcselS"), form("FcselD")
)

// FcmpS emits FCMP Sn, Sm, setting NZCV.
//
// An unordered pair sets C and V, which is what makes the unsigned conditions
// the ones that answer a float comparison: MI is less-than and reads false for
// a NaN, where LT would read true.
func (s *Section) FcmpS(rn, rm reg.S) { s.inst(fcmpS, rn, rm) }
func (s *Section) FcmpD(rn, rm reg.D) { s.inst(fcmpD, rn, rm) }

// FcmpZeroS emits FCMP Sn, #0.0, which is a separate encoding rather than a
// comparison against a register that happens to hold zero.
func (s *Section) FcmpZeroS(rn reg.S) { s.inst(fcmpZeroS, rn) }
func (s *Section) FcmpZeroD(rn reg.D) { s.inst(fcmpZeroD, rn) }

// FcselS emits FCSEL Sd, Sn, Sm, cond: Sn when cond holds, Sm otherwise.
func (s *Section) FcselS(rd, rn, rm reg.S, c Cond) { s.inst(fcselS, rd, rn, rm, c) }
func (s *Section) FcselD(rd, rn, rm reg.D, c Cond) { s.inst(fcselD, rd, rn, rm, c) }

// ---- Between the two register files ----------------------------------------

var (
	scvtfWToS, scvtfWToD = form("ScvtfWToS"), form("ScvtfWToD")
	scvtfXToS, scvtfXToD = form("ScvtfXToS"), form("ScvtfXToD")
	ucvtfWToS, ucvtfWToD = form("UcvtfWToS"), form("UcvtfWToD")
	ucvtfXToS, ucvtfXToD = form("UcvtfXToS"), form("UcvtfXToD")
)

// ScvtfWToS emits SCVTF Sd, Wn: a signed integer to a float, rounded by FPCR.
func (s *Section) ScvtfWToS(rd reg.S, rn reg.W) { s.inst(scvtfWToS, rd, rn) }
func (s *Section) ScvtfWToD(rd reg.D, rn reg.W) { s.inst(scvtfWToD, rd, rn) }
func (s *Section) ScvtfXToS(rd reg.S, rn reg.X) { s.inst(scvtfXToS, rd, rn) }
func (s *Section) ScvtfXToD(rd reg.D, rn reg.X) { s.inst(scvtfXToD, rd, rn) }

// UcvtfWToS emits UCVTF Sd, Wn: an unsigned integer to a float.
func (s *Section) UcvtfWToS(rd reg.S, rn reg.W) { s.inst(ucvtfWToS, rd, rn) }
func (s *Section) UcvtfWToD(rd reg.D, rn reg.W) { s.inst(ucvtfWToD, rd, rn) }
func (s *Section) UcvtfXToS(rd reg.S, rn reg.X) { s.inst(ucvtfXToS, rd, rn) }
func (s *Section) UcvtfXToD(rd reg.D, rn reg.X) { s.inst(ucvtfXToD, rd, rn) }

var (
	fcvtzsSToW, fcvtzsDToW = form("FcvtzsSToW"), form("FcvtzsDToW")
	fcvtzsSToX, fcvtzsDToX = form("FcvtzsSToX"), form("FcvtzsDToX")
	fcvtzuSToW, fcvtzuDToW = form("FcvtzuSToW"), form("FcvtzuDToW")
	fcvtzuSToX, fcvtzuDToX = form("FcvtzuSToX"), form("FcvtzuDToX")
)

// FcvtzsSToW emits FCVTZS Wd, Sn: a float to a signed integer, rounded toward
// zero, saturating rather than trapping — out of range gives the nearest
// representable integer and a NaN gives zero. A language wanting a trap checks
// the range first.
func (s *Section) FcvtzsSToW(rd reg.W, rn reg.S) { s.inst(fcvtzsSToW, rd, rn) }
func (s *Section) FcvtzsDToW(rd reg.W, rn reg.D) { s.inst(fcvtzsDToW, rd, rn) }
func (s *Section) FcvtzsSToX(rd reg.X, rn reg.S) { s.inst(fcvtzsSToX, rd, rn) }
func (s *Section) FcvtzsDToX(rd reg.X, rn reg.D) { s.inst(fcvtzsDToX, rd, rn) }

// FcvtzuSToW emits FCVTZU Wd, Sn, the unsigned form.
func (s *Section) FcvtzuSToW(rd reg.W, rn reg.S) { s.inst(fcvtzuSToW, rd, rn) }
func (s *Section) FcvtzuDToW(rd reg.W, rn reg.D) { s.inst(fcvtzuDToW, rd, rn) }
func (s *Section) FcvtzuSToX(rd reg.X, rn reg.S) { s.inst(fcvtzuSToX, rd, rn) }
func (s *Section) FcvtzuDToX(rd reg.X, rn reg.D) { s.inst(fcvtzuDToX, rd, rn) }

var (
	fmovSToW = form("FmovSToW")
	fmovWToS = form("FmovWToS")
	fmovDToX = form("FmovDToX")
	fmovXToD = form("FmovXToD")
)

// FmovSToW emits FMOV Wd, Sn: the bit pattern moved between the files, not
// converted. This is what a bitcast is; ScvtfWToS is what a conversion is.
func (s *Section) FmovSToW(rd reg.W, rn reg.S) { s.inst(fmovSToW, rd, rn) }
func (s *Section) FmovWToS(rd reg.S, rn reg.W) { s.inst(fmovWToS, rd, rn) }
func (s *Section) FmovDToX(rd reg.X, rn reg.D) { s.inst(fmovDToX, rd, rn) }
func (s *Section) FmovXToD(rd reg.D, rn reg.X) { s.inst(fmovXToD, rd, rn) }

// ---- Load and store --------------------------------------------------------

var (
	ldrImmS, ldrImmD, ldrImmQ = form("LdrImmS"), form("LdrImmD"), form("LdrImmQ")
	strImmS, strImmD, strImmQ = form("StrImmS"), form("StrImmD"), form("StrImmQ")
	ldurImmS, ldurImmD        = form("LdurImmS"), form("LdurImmD")
	sturImmS, sturImmD        = form("SturImmS"), form("SturImmD")
)

// LdrImmS emits LDR St, [Xn|SP{, #imm}], imm in units of the access width.
func (s *Section) LdrImmS(rt reg.S, m Mem) { s.inst(ldrImmS, rt, m) }
func (s *Section) LdrImmD(rt reg.D, m Mem) { s.inst(ldrImmD, rt, m) }
func (s *Section) LdrImmQ(rt reg.Q, m Mem) { s.inst(ldrImmQ, rt, m) }
func (s *Section) StrImmS(rt reg.S, m Mem) { s.inst(strImmS, rt, m) }
func (s *Section) StrImmD(rt reg.D, m Mem) { s.inst(strImmD, rt, m) }
func (s *Section) StrImmQ(rt reg.Q, m Mem) { s.inst(strImmQ, rt, m) }

// LdurImmS emits LDUR St, [Xn|SP{, #imm}], imm a signed byte count. Unscaled
// is the whole difference from LDR.
func (s *Section) LdurImmS(rt reg.S, m Mem) { s.inst(ldurImmS, rt, m) }
func (s *Section) LdurImmD(rt reg.D, m Mem) { s.inst(ldurImmD, rt, m) }
func (s *Section) SturImmS(rt reg.S, m Mem) { s.inst(sturImmS, rt, m) }
func (s *Section) SturImmD(rt reg.D, m Mem) { s.inst(sturImmD, rt, m) }
