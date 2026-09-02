package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Conditional select ----------------------------------------------------

var (
	csel32  = form("Csel32")
	csel64  = form("Csel64")
	csinc32 = form("Csinc32")
	csinc64 = form("Csinc64")
	csinv32 = form("Csinv32")
	csinv64 = form("Csinv64")
	csneg32 = form("Csneg32")
	csneg64 = form("Csneg64")
)

// Csel32 emits CSEL Wd, Wn, Wm, cond: Wd = cond ? Wn : Wm.
func (s *Section) Csel32(rd, rn, rm reg.W, c Cond) { s.inst(csel32, rd, rn, rm, c) }
func (s *Section) Csel64(rd, rn, rm reg.X, c Cond) { s.inst(csel64, rd, rn, rm, c) }

// Csinc32 emits CSINC Wd, Wn, Wm, cond: Wd = cond ? Wn : Wm+1.
func (s *Section) Csinc32(rd, rn, rm reg.W, c Cond) { s.inst(csinc32, rd, rn, rm, c) }
func (s *Section) Csinc64(rd, rn, rm reg.X, c Cond) { s.inst(csinc64, rd, rn, rm, c) }

// Csinv32 emits CSINV Wd, Wn, Wm, cond: Wd = cond ? Wn : ^Wm.
func (s *Section) Csinv32(rd, rn, rm reg.W, c Cond) { s.inst(csinv32, rd, rn, rm, c) }
func (s *Section) Csinv64(rd, rn, rm reg.X, c Cond) { s.inst(csinv64, rd, rn, rm, c) }

// Csneg32 emits CSNEG Wd, Wn, Wm, cond: Wd = cond ? Wn : -Wm.
func (s *Section) Csneg32(rd, rn, rm reg.W, c Cond) { s.inst(csneg32, rd, rn, rm, c) }
func (s *Section) Csneg64(rd, rn, rm reg.X, c Cond) { s.inst(csneg64, rd, rn, rm, c) }

// ---- Conditional compare ----------------------------------------------------
//
// A compare that happens only when its condition holds, and writes the flags
// it was given when it does not. The nzcv operand is that answer: four bits
// standing in for the comparison that was not made, which is how a chain of
// C's && and || becomes straight-line code.

var (
	ccmpReg32 = form("CcmpReg32")
	ccmpReg64 = form("CcmpReg64")
	ccmpImm32 = form("CcmpImm32")
	ccmpImm64 = form("CcmpImm64")
	ccmnReg32 = form("CcmnReg32")
	ccmnReg64 = form("CcmnReg64")
	ccmnImm32 = form("CcmnImm32")
	ccmnImm64 = form("CcmnImm64")
)

// CcmpReg32 emits CCMP Wn, Wm, #nzcv, cond.
func (s *Section) CcmpReg32(rn, rm reg.W, nzcv uint8, c Cond) {
	s.inst(ccmpReg32, rn, rm, uint64(nzcv), c)
}

// CcmpReg64 emits CCMP Xn, Xm, #nzcv, cond.
func (s *Section) CcmpReg64(rn, rm reg.X, nzcv uint8, c Cond) {
	s.inst(ccmpReg64, rn, rm, uint64(nzcv), c)
}

// CcmpImm32 emits CCMP Wn, #imm, #nzcv, cond. The immediate is five bits
// unsigned, which is the whole of what this form compares against.
func (s *Section) CcmpImm32(rn reg.W, imm uint8, nzcv uint8, c Cond) {
	s.inst(ccmpImm32, rn, uint64(imm), uint64(nzcv), c)
}

// CcmpImm64 emits CCMP Xn, #imm, #nzcv, cond.
func (s *Section) CcmpImm64(rn reg.X, imm uint8, nzcv uint8, c Cond) {
	s.inst(ccmpImm64, rn, uint64(imm), uint64(nzcv), c)
}

// CcmnReg32 emits CCMN Wn, Wm, #nzcv, cond — the same, comparing against the
// negation.
func (s *Section) CcmnReg32(rn, rm reg.W, nzcv uint8, c Cond) {
	s.inst(ccmnReg32, rn, rm, uint64(nzcv), c)
}

// CcmnReg64 emits CCMN Xn, Xm, #nzcv, cond.
func (s *Section) CcmnReg64(rn, rm reg.X, nzcv uint8, c Cond) {
	s.inst(ccmnReg64, rn, rm, uint64(nzcv), c)
}

// CcmnImm32 emits CCMN Wn, #imm, #nzcv, cond.
func (s *Section) CcmnImm32(rn reg.W, imm uint8, nzcv uint8, c Cond) {
	s.inst(ccmnImm32, rn, uint64(imm), uint64(nzcv), c)
}

// CcmnImm64 emits CCMN Xn, #imm, #nzcv, cond.
func (s *Section) CcmnImm64(rn reg.X, imm uint8, nzcv uint8, c Cond) {
	s.inst(ccmnImm64, rn, uint64(imm), uint64(nzcv), c)
}

// ---- Alias: set on condition ------------------------------------------------

var cset64 = form("Cset64")

// Cset64 emits CSET Xd, cond: Xd = cond ? 1 : 0.
//
// The condition is inverted on the way into the field, because CSINC
// increments when its own condition is false. That inversion is the row's
// (AttrInvertCond), not this function's: an assembler reaching the same row
// through Emit has to get the same word, and it did not while the rule lived
// here.
func (s *Section) Cset64(rd reg.X, c Cond) { s.inst(cset64, rd, c) }
