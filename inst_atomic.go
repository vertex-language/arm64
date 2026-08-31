package arm64

import "github.com/vertex-language/arm64/reg"

// Atomics: the ordered accesses, the load/store-exclusive pairs, and the LSE
// single-instruction family.
//
// None takes an offset — the encodings have no immediate field — so every Mem
// here must be a bare base. Building one with .Off is refused at the call.

var (
	ldarb, ldarh   = form("Ldarb"), form("Ldarh")
	ldar32, ldar64 = form("Ldar32"), form("Ldar64")
	stlrb, stlrh   = form("Stlrb"), form("Stlrh")
	stlr32, stlr64 = form("Stlr32"), form("Stlr64")
)

// Ldar32 emits LDAR Wt, [Xn|SP]: an acquire load. No exclusive monitor is
// involved — this orders an ordinary access, where the pairs below make one
// atomic.
func (s *Section) Ldar32(rt reg.W, m Mem) { s.inst(ldar32, rt, m) }
func (s *Section) Ldar64(rt reg.X, m Mem) { s.inst(ldar64, rt, m) }
func (s *Section) Ldarb(rt reg.W, m Mem)  { s.inst(ldarb, rt, m) }
func (s *Section) Ldarh(rt reg.W, m Mem)  { s.inst(ldarh, rt, m) }

// Stlr32 emits STLR Wt, [Xn|SP]: a release store.
func (s *Section) Stlr32(rt reg.W, m Mem) { s.inst(stlr32, rt, m) }
func (s *Section) Stlr64(rt reg.X, m Mem) { s.inst(stlr64, rt, m) }
func (s *Section) Stlrb(rt reg.W, m Mem)  { s.inst(stlrb, rt, m) }
func (s *Section) Stlrh(rt reg.W, m Mem)  { s.inst(stlrh, rt, m) }

// ---- Load and store exclusive ----------------------------------------------

var (
	ldxrb, ldxrh     = form("Ldxrb"), form("Ldxrh")
	ldxr32, ldxr64   = form("Ldxr32"), form("Ldxr64")
	ldaxrb, ldaxrh   = form("Ldaxrb"), form("Ldaxrh")
	ldaxr32, ldaxr64 = form("Ldaxr32"), form("Ldaxr64")
	stxrb, stxrh     = form("Stxrb"), form("Stxrh")
	stxr32, stxr64   = form("Stxr32"), form("Stxr64")
	stlxrb, stlxrh   = form("Stlxrb"), form("Stlxrh")
	stlxr32, stlxr64 = form("Stlxr32"), form("Stlxr64")
)

// Ldxr32 emits LDXR Wt, [Xn|SP], taking the exclusive monitor.
func (s *Section) Ldxr32(rt reg.W, m Mem) { s.inst(ldxr32, rt, m) }
func (s *Section) Ldxr64(rt reg.X, m Mem) { s.inst(ldxr64, rt, m) }
func (s *Section) Ldxrb(rt reg.W, m Mem)  { s.inst(ldxrb, rt, m) }
func (s *Section) Ldxrh(rt reg.W, m Mem)  { s.inst(ldxrh, rt, m) }

// Ldaxr32 is Ldxr32 with acquire ordering, which is the half of a
// sequentially consistent read-modify-write that the load carries.
func (s *Section) Ldaxr32(rt reg.W, m Mem) { s.inst(ldaxr32, rt, m) }
func (s *Section) Ldaxr64(rt reg.X, m Mem) { s.inst(ldaxr64, rt, m) }
func (s *Section) Ldaxrb(rt reg.W, m Mem)  { s.inst(ldaxrb, rt, m) }
func (s *Section) Ldaxrh(rt reg.W, m Mem)  { s.inst(ldaxrh, rt, m) }

// Stxr32 emits STXR Ws, Wt, [Xn|SP]. Ws is the status: zero if the store
// succeeded, one if the monitor was lost and the loop has to retry.
func (s *Section) Stxr32(rs, rt reg.W, m Mem)       { s.inst(stxr32, rs, rt, m) }
func (s *Section) Stxr64(rs reg.W, rt reg.X, m Mem) { s.inst(stxr64, rs, rt, m) }
func (s *Section) Stxrb(rs, rt reg.W, m Mem)        { s.inst(stxrb, rs, rt, m) }
func (s *Section) Stxrh(rs, rt reg.W, m Mem)        { s.inst(stxrh, rs, rt, m) }

// Stlxr32 is Stxr32 with release ordering.
func (s *Section) Stlxr32(rs, rt reg.W, m Mem)       { s.inst(stlxr32, rs, rt, m) }
func (s *Section) Stlxr64(rs reg.W, rt reg.X, m Mem) { s.inst(stlxr64, rs, rt, m) }
func (s *Section) Stlxrb(rs, rt reg.W, m Mem)        { s.inst(stlxrb, rs, rt, m) }
func (s *Section) Stlxrh(rs, rt reg.W, m Mem)        { s.inst(stlxrh, rs, rt, m) }

var clrex = form("Clrex")

// Clrex emits CLREX: clears this processor's exclusive monitor. The failure
// path of a compare-and-swap loop, which took a reservation with LDXR and then
// did not store.
func (s *Section) Clrex() { s.inst(clrex) }

// ---- LSE -------------------------------------------------------------------
//
// One instruction for what the exclusive loop does, and one the loop cannot be
// interrupted out of. Gated on FEAT_LSE, mandatory from Armv8.1-A.
//
// Rs is the operand and Rt the old value returned, which is why both are named
// on what reads as a store.

var (
	ldaddalb, ldaddalh   = form("Ldaddalb"), form("Ldaddalh")
	ldaddal32, ldaddal64 = form("Ldaddal32"), form("Ldaddal64")
	ldclralb, ldclralh   = form("Ldclralb"), form("Ldclralh")
	ldclral32, ldclral64 = form("Ldclral32"), form("Ldclral64")
	ldeoralb, ldeoralh   = form("Ldeoralb"), form("Ldeoralh")
	ldeoral32, ldeoral64 = form("Ldeoral32"), form("Ldeoral64")
	ldsetalb, ldsetalh   = form("Ldsetalb"), form("Ldsetalh")
	ldsetal32, ldsetal64 = form("Ldsetal32"), form("Ldsetal64")
	swpalb, swpalh       = form("Swpalb"), form("Swpalh")
	swpal32, swpal64     = form("Swpal32"), form("Swpal64")
	casalb, casalh       = form("Casalb"), form("Casalh")
	casal32, casal64     = form("Casal32"), form("Casal64")
)

// Ldaddal32 emits LDADDAL Ws, Wt, [Xn|SP]: add Ws to memory, return the old
// value in Wt, with acquire-release ordering.
func (s *Section) Ldaddal32(rs, rt reg.W, m Mem) { s.inst(ldaddal32, rs, rt, m) }
func (s *Section) Ldaddal64(rs, rt reg.X, m Mem) { s.inst(ldaddal64, rs, rt, m) }
func (s *Section) Ldaddalb(rs, rt reg.W, m Mem)  { s.inst(ldaddalb, rs, rt, m) }
func (s *Section) Ldaddalh(rs, rt reg.W, m Mem)  { s.inst(ldaddalh, rs, rt, m) }

// Ldclral32 clears the bits Ws sets, so an atomic AND is a clear of the
// complement rather than of the mask.
func (s *Section) Ldclral32(rs, rt reg.W, m Mem) { s.inst(ldclral32, rs, rt, m) }
func (s *Section) Ldclral64(rs, rt reg.X, m Mem) { s.inst(ldclral64, rs, rt, m) }
func (s *Section) Ldclralb(rs, rt reg.W, m Mem)  { s.inst(ldclralb, rs, rt, m) }
func (s *Section) Ldclralh(rs, rt reg.W, m Mem)  { s.inst(ldclralh, rs, rt, m) }

// Ldeoral32 is the atomic exclusive-or.
func (s *Section) Ldeoral32(rs, rt reg.W, m Mem) { s.inst(ldeoral32, rs, rt, m) }
func (s *Section) Ldeoral64(rs, rt reg.X, m Mem) { s.inst(ldeoral64, rs, rt, m) }
func (s *Section) Ldeoralb(rs, rt reg.W, m Mem)  { s.inst(ldeoralb, rs, rt, m) }
func (s *Section) Ldeoralh(rs, rt reg.W, m Mem)  { s.inst(ldeoralh, rs, rt, m) }

// Ldsetal32 is the atomic inclusive-or: it sets the bits Ws sets.
func (s *Section) Ldsetal32(rs, rt reg.W, m Mem) { s.inst(ldsetal32, rs, rt, m) }
func (s *Section) Ldsetal64(rs, rt reg.X, m Mem) { s.inst(ldsetal64, rs, rt, m) }
func (s *Section) Ldsetalb(rs, rt reg.W, m Mem)  { s.inst(ldsetalb, rs, rt, m) }
func (s *Section) Ldsetalh(rs, rt reg.W, m Mem)  { s.inst(ldsetalh, rs, rt, m) }

// Swpal32 emits SWPAL Ws, Wt, [Xn|SP]: the atomic exchange.
func (s *Section) Swpal32(rs, rt reg.W, m Mem) { s.inst(swpal32, rs, rt, m) }
func (s *Section) Swpal64(rs, rt reg.X, m Mem) { s.inst(swpal64, rs, rt, m) }
func (s *Section) Swpalb(rs, rt reg.W, m Mem)  { s.inst(swpalb, rs, rt, m) }
func (s *Section) Swpalh(rs, rt reg.W, m Mem)  { s.inst(swpalh, rs, rt, m) }

// Casal32 emits CASAL Ws, Wt, [Xn|SP]: compare memory against Ws, store Wt if
// they matched, and write the old value back into Ws either way. Ws is both
// the expected value going in and the answer coming out, which is what makes
// the comparison that follows a compare of Ws against what it was.
func (s *Section) Casal32(rs, rt reg.W, m Mem) { s.inst(casal32, rs, rt, m) }
func (s *Section) Casal64(rs, rt reg.X, m Mem) { s.inst(casal64, rs, rt, m) }
func (s *Section) Casalb(rs, rt reg.W, m Mem)  { s.inst(casalb, rs, rt, m) }
func (s *Section) Casalh(rs, rt reg.W, m Mem)  { s.inst(casalh, rs, rt, m) }
