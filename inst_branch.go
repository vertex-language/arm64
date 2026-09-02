package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Branches ---------------------------------------------------------------
//
// target is a Label for a same-section fold at Finalize, or a SymRef —
// optionally Ref("name", kind) to insist on a relocation kind — for one that
// survives as a Reference. B and Bl's target may also be wrapped in Page or
// GotPage, though a branch to a GOT slot's address rather than through it is
// unusual enough that most callers never will.

var (
	bForm  = form("B")
	blForm = form("Bl")
	bCond  = form("BCond")
	cbz32  = form("Cbz32")
	cbz64  = form("Cbz64")
	cbnz32 = form("Cbnz32")
	cbnz64 = form("Cbnz64")
	tbz64  = form("Tbz64")
	tbnz64 = form("Tbnz64")
	tbz32  = form("Tbz32")
	tbnz32 = form("Tbnz32")
	br     = form("Br")
	blr    = form("Blr")
	ret    = form("Ret")
)

// B emits B target: an unconditional branch, +/-128MiB.
func (s *Section) B(target TargetOp) { s.inst(bForm, target) }

// Bl emits BL target: a branch with link, X30 set to the return address.
func (s *Section) Bl(target TargetOp) { s.inst(blForm, target) }

// BCond emits B.cond target, +/-1MiB.
func (s *Section) BCond(c Cond, target TargetOp) { s.inst(bCond, c, target) }

// Cbz32 emits CBZ Wt, target: branch if Wt is zero, +/-1MiB.
func (s *Section) Cbz32(rt reg.W, target TargetOp) { s.inst(cbz32, rt, target) }
func (s *Section) Cbz64(rt reg.X, target TargetOp) { s.inst(cbz64, rt, target) }

// Cbnz32 emits CBNZ Wt, target: branch if Wt is nonzero.
func (s *Section) Cbnz32(rt reg.W, target TargetOp) { s.inst(cbnz32, rt, target) }
func (s *Section) Cbnz64(rt reg.X, target TargetOp) { s.inst(cbnz64, rt, target) }

// Tbz64 emits TBZ Xt, #bit, target: branch if bit bit of Xt is zero,
// +/-32KiB. bit ranges 0 to 63, and its top bit folds into the opcode itself.
func (s *Section) Tbz64(rt reg.X, bit uint8, target TargetOp) {
	s.inst(tbz64, rt, uint64(bit), target)
}

// Tbnz64 emits TBNZ Xt, #bit, target: branch if bit bit of Xt is one.
func (s *Section) Tbnz64(rt reg.X, bit uint8, target TargetOp) {
	s.inst(tbnz64, rt, uint64(bit), target)
}

// Tbz32 emits TBZ Wt, #bit, target. bit ranges 0 to 31, which is the whole
// difference from the 64-bit form: the word is the same one.
func (s *Section) Tbz32(rt reg.W, bit uint8, target TargetOp) {
	s.inst(tbz32, rt, uint64(bit), target)
}

// Tbnz32 emits TBNZ Wt, #bit, target.
func (s *Section) Tbnz32(rt reg.W, bit uint8, target TargetOp) {
	s.inst(tbnz32, rt, uint64(bit), target)
}

// Br emits BR Xn: an unconditional branch to a register, no relocation
// possible or needed.
func (s *Section) Br(rn reg.X) { s.inst(br, rn) }

// Blr emits BLR Xn: a branch with link to a register.
func (s *Section) Blr(rn reg.X) { s.inst(blr, rn) }

// Ret emits RET Xn, or RET with no operand for the architecture's own
// default of X30.
func (s *Section) Ret(rn ...reg.X) {
	if len(rn) > 0 {
		s.inst(ret, rn[0])
		return
	}
	s.inst(ret)
}
