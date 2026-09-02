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
