package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Move wide ---------------------------------------------------------------

var (
	movzImm32 = form("MovzImm32")
	movzImm64 = form("MovzImm64")
	movnImm32 = form("MovnImm32")
	movnImm64 = form("MovnImm64")
	movkImm32 = form("MovkImm32")
	movkImm64 = form("MovkImm64")
)

// MovzImm32 emits MOVZ Wd, #imm{, LSL #shift}: imm at the named halfword,
// zero elsewhere.
func (s *Section) MovzImm32(rd reg.W, imm uint32, shift ...ShiftOp) {
	s.inst(movzImm32, append([]any{rd, uint64(imm)}, opt(shift)...)...)
}

// MovzImm64 emits MOVZ Xd, #imm{, LSL #shift}.
func (s *Section) MovzImm64(rd reg.X, imm uint64, shift ...ShiftOp) {
	s.inst(movzImm64, append([]any{rd, imm}, opt(shift)...)...)
}

// MovnImm32 emits MOVN Wd, #imm{, LSL #shift}: the bitwise NOT of imm at the
// named halfword, ones elsewhere — the way a 32-bit constant with its top
// bits set is loaded in one instruction rather than a MOVZ/MOVK chain.
func (s *Section) MovnImm32(rd reg.W, imm uint32, shift ...ShiftOp) {
	s.inst(movnImm32, append([]any{rd, uint64(imm)}, opt(shift)...)...)
}
func (s *Section) MovnImm64(rd reg.X, imm uint64, shift ...ShiftOp) {
	s.inst(movnImm64, append([]any{rd, imm}, opt(shift)...)...)
}

// MovkImm32 emits MOVK Wd, #imm{, LSL #shift}: imm at the named halfword,
// every other bit of Wd unchanged.
func (s *Section) MovkImm32(rd reg.W, imm uint32, shift ...ShiftOp) {
	s.inst(movkImm32, append([]any{rd, uint64(imm)}, opt(shift)...)...)
}
func (s *Section) MovkImm64(rd reg.X, imm uint64, shift ...ShiftOp) {
	s.inst(movkImm64, append([]any{rd, imm}, opt(shift)...)...)
}

// ---- PC-relative address ------------------------------------------------------

var (
	adr  = form("Adr")
	adrp = form("Adrp")
)

// Adr emits ADR Xd, target: the byte address of target, which must be within
// 1MiB. target is a Label for a same-section fold, or a SymRef (optionally
// wrapped in arm64.Direct) for one that survives to Finalize as RefAdrPrel21.
func (s *Section) Adr(rd reg.X, target TargetOp) {
	s.inst(adr, rd, target)
}

// Adrp emits ADRP Xd, target: the 4KiB page address of target. Follow with
// AddImm64(rd, rd, PageOff(target)) or a load through PageOff/GotPage to
// reach the byte within the page.
func (s *Section) Adrp(rd reg.X, target TargetOp) {
	s.inst(adrp, rd, target)
}

// ---- Aliases: move ------------------------------------------------------------

var (
	movReg64  = form("MovReg64")
	movReg32  = form("MovReg32")
	movSp64   = form("MovSp64")
	movWide64 = form("MovWide64")
)

// MovReg64 emits MOV Xd, Xm — ORR with XZR as the first source. Preferred
// only when no shift is applied; ORR with a shift is not a move.
func (s *Section) MovReg64(rd, rm reg.X) { s.inst(movReg64, rd, rm) }

// MovReg32 emits MOV Wd, Wm.
func (s *Section) MovReg32(rd, rm reg.W) { s.inst(movReg32, rd, rm) }

// MovSp64 emits MOV Xd|SP, Xn|SP — ADD with a zero immediate. Preferred only
// when one of the two registers really is SP; RegSP64 accepts either the
// numbered X or the SP alias for that reason.
func (s *Section) MovSp64(rd, rn RegSP64) { s.inst(movSp64, rd, rn) }

// MovWide64 emits MOV Xd, #imm — MOVZ, preferred unless imm is zero and
// shifted, which is not a move of anything.
func (s *Section) MovWide64(rd reg.X, imm uint64) { s.inst(movWide64, rd, imm) }
