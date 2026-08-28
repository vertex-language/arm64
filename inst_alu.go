package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Logical, shifted register -----------------------------------------------

var (
	andShifted32  = form("AndShifted32")
	andShifted64  = form("AndShifted64")
	orrShifted32  = form("OrrShifted32")
	orrShifted64  = form("OrrShifted64")
	eorShifted32  = form("EorShifted32")
	eorShifted64  = form("EorShifted64")
	andsShifted32 = form("AndsShifted32")
	andsShifted64 = form("AndsShifted64")
)

func (s *Section) AndShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(andShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) AndShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(andShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) OrrShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(orrShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) OrrShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(orrShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) EorShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(eorShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) EorShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(eorShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) AndsShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(andsShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}

// AndsShifted64 emits ANDS Xd, Xn, Xm{, shift}, setting NZCV. Tst is this
// table's TST: ANDS with Xd pinned to XZR.
func (s *Section) AndsShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(andsShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}

// ---- Logical, bitmask immediate ----------------------------------------------

var (
	andImm32 = form("AndImm32")
	andImm64 = form("AndImm64")
	orrImm32 = form("OrrImm32")
	orrImm64 = form("OrrImm64")
	eorImm32 = form("EorImm32")
	eorImm64 = form("EorImm64")
)

// AndImm32 emits AND Wd|WSP, Wn, #imm. imm must be a valid 32-bit logical
// immediate — a rotated run of ones replicated to fill the register — or the
// call fails with ErrBitmask; there is no encoding for an arbitrary constant
// here; materialize it into a register instead.
func (s *Section) AndImm32(rd RegSP32, rn reg.W, imm uint32) {
	s.inst(andImm32, rd, rn, uint64(imm))
}

// AndImm64 emits AND Xd|SP, Xn, #imm.
func (s *Section) AndImm64(rd RegSP64, rn reg.X, imm uint64) {
	s.inst(andImm64, rd, rn, imm)
}

func (s *Section) OrrImm32(rd RegSP32, rn reg.W, imm uint32) {
	s.inst(orrImm32, rd, rn, uint64(imm))
}
func (s *Section) OrrImm64(rd RegSP64, rn reg.X, imm uint64) {
	s.inst(orrImm64, rd, rn, imm)
}
func (s *Section) EorImm32(rd RegSP32, rn reg.W, imm uint32) {
	s.inst(eorImm32, rd, rn, uint64(imm))
}
func (s *Section) EorImm64(rd RegSP64, rn reg.X, imm uint64) {
	s.inst(eorImm64, rd, rn, imm)
}

var (
	andsImm32 = form("AndsImm32")
	andsImm64 = form("AndsImm64")
)

// AndsImm32 emits ANDS Wd, Wn, #imm, setting NZCV. Unlike AndImm32, Wd is a
// plain register and never SP: a comparison result has nowhere useful to go
// but a register or the flags.
func (s *Section) AndsImm32(rd, rn reg.W, imm uint32) { s.inst(andsImm32, rd, rn, uint64(imm)) }
func (s *Section) AndsImm64(rd, rn reg.X, imm uint64) { s.inst(andsImm64, rd, rn, imm) }

// ---- Aliases: test bits ------------------------------------------------------

var (
	tstShifted64 = form("TstShifted64")
	tstImm64     = form("TstImm64")
)

// TstShifted64 emits TST Xn, Xm{, shift} — ANDS with the result discarded.
func (s *Section) TstShifted64(rn, rm reg.X, shift ...ShiftOp) {
	s.inst(tstShifted64, append([]any{rn, rm}, opt(shift)...)...)
}

// TstImm64 emits TST Xn, #imm — ANDS (immediate) with the result discarded.
func (s *Section) TstImm64(rn reg.X, imm uint64) { s.inst(tstImm64, rn, imm) }
