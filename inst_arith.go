package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Arithmetic, immediate --------------------------------------------------

var (
	addImm32  = form("AddImm32")
	addImm64  = form("AddImm64")
	addsImm32 = form("AddsImm32")
	addsImm64 = form("AddsImm64")
	subImm32  = form("SubImm32")
	subImm64  = form("SubImm64")
	subsImm32 = form("SubsImm32")
	subsImm64 = form("SubsImm64")
)

// AddImm32 emits ADD Wd|WSP, Wn|WSP, #imm{, shift}. imm may be a plain
// constant or the :lo12: half of an address — add w0, w0, :lo12:msg is the
// second instruction of an ADRP/ADD pair.
func (s *Section) AddImm32(rd, rn RegSP32, imm ImmOrRef, shift ...ShiftOp) {
	s.inst(addImm32, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// AddImm64 emits ADD Xd|SP, Xn|SP, #imm{, shift}.
func (s *Section) AddImm64(rd, rn RegSP64, imm ImmOrRef, shift ...ShiftOp) {
	s.inst(addImm64, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// AddsImm32 emits ADDS Wd, Wn|WSP, #imm{, shift}, setting NZCV.
func (s *Section) AddsImm32(rd reg.W, rn RegSP32, imm int64, shift ...ShiftOp) {
	s.inst(addsImm32, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// AddsImm64 emits ADDS Xd, Xn|SP, #imm{, shift}, setting NZCV.
func (s *Section) AddsImm64(rd reg.X, rn RegSP64, imm int64, shift ...ShiftOp) {
	s.inst(addsImm64, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// SubImm32 emits SUB Wd|WSP, Wn|WSP, #imm{, shift}.
func (s *Section) SubImm32(rd, rn RegSP32, imm int64, shift ...ShiftOp) {
	s.inst(subImm32, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// SubImm64 emits SUB Xd|SP, Xn|SP, #imm{, shift}.
func (s *Section) SubImm64(rd, rn RegSP64, imm int64, shift ...ShiftOp) {
	s.inst(subImm64, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// SubsImm32 emits SUBS Wd, Wn|WSP, #imm{, shift}, setting NZCV.
func (s *Section) SubsImm32(rd reg.W, rn RegSP32, imm int64, shift ...ShiftOp) {
	s.inst(subsImm32, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// SubsImm64 emits SUBS Xd, Xn|SP, #imm{, shift}, setting NZCV.
func (s *Section) SubsImm64(rd reg.X, rn RegSP64, imm int64, shift ...ShiftOp) {
	s.inst(subsImm64, append([]any{rd, rn, imm}, opt(shift)...)...)
}

// ---- Arithmetic, shifted register --------------------------------------------

var (
	addShifted32  = form("AddShifted32")
	addShifted64  = form("AddShifted64")
	addsShifted32 = form("AddsShifted32")
	addsShifted64 = form("AddsShifted64")
	subShifted32  = form("SubShifted32")
	subShifted64  = form("SubShifted64")
	subsShifted32 = form("SubsShifted32")
	subsShifted64 = form("SubsShifted64")
)

func (s *Section) AddShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(addShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) AddShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(addShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) AddsShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(addsShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) AddsShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(addsShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) SubShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(subShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) SubShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(subShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}
func (s *Section) SubsShifted32(rd, rn, rm reg.W, shift ...ShiftOp) {
	s.inst(subsShifted32, append([]any{rd, rn, rm}, opt(shift)...)...)
}

// SubsShifted64 emits SUBS Xd, Xn, Xm{, shift}. Cmp is this table's CMP: SUBS
// with Xd pinned to XZR.
func (s *Section) SubsShifted64(rd, rn, rm reg.X, shift ...ShiftOp) {
	s.inst(subsShifted64, append([]any{rd, rn, rm}, opt(shift)...)...)
}

// ---- Arithmetic, extended register -------------------------------------------

var (
	addExt32  = form("AddExt32")
	addExt64  = form("AddExt64")
	subExt64  = form("SubExt64")
	subsExt64 = form("SubsExt64")
)

// AddExt32 emits ADD Wd|WSP, Wn|WSP, Wm{, extend}.
func (s *Section) AddExt32(rd, rn RegSP32, rm reg.W, ext ...ExtendOp) {
	s.inst(addExt32, append([]any{rd, rn, rm}, opt(ext)...)...)
}

// AddExt64 emits ADD Xd|SP, Xn|SP, Rm{, extend}. Rm is a reg.X or a reg.W —
// UXTW and SXTW read the 32-bit half of the register named — so it is any,
// refused by name at the call if it holds anything else.
func (s *Section) AddExt64(rd, rn RegSP64, rm any, ext ...ExtendOp) {
	s.inst(addExt64, append([]any{rd, rn, rm}, opt(ext)...)...)
}

// SubExt64 emits SUB Xd|SP, Xn|SP, Rm{, extend}. The shifted-register form
// reads register 31 as ZR, so this is the only SUB that can write SP or read
// it — which is what moving the stack pointer by a value rather than a
// literal needs.
func (s *Section) SubExt64(rd, rn RegSP64, rm any, ext ...ExtendOp) {
	s.inst(subExt64, append([]any{rd, rn, rm}, opt(ext)...)...)
}

// SubsExt64 emits SUBS Xd, Xn|SP, Rm{, extend}, setting NZCV. Cmp on an
// extended register goes through this form rather than SubsShifted64.
func (s *Section) SubsExt64(rd reg.X, rn RegSP64, rm any, ext ...ExtendOp) {
	s.inst(subsExt64, append([]any{rd, rn, rm}, opt(ext)...)...)
}

// ---- Aliases: compare, compare-negative, negate ------------------------------

var (
	cmpShifted64 = form("CmpShifted64")
	cmpShifted32 = form("CmpShifted32")
	cmpImm64     = form("CmpImm64")
	cmpImm32     = form("CmpImm32")
	cmnShifted64 = form("CmnShifted64")
	cmnShifted32 = form("CmnShifted32")
	cmnImm64     = form("CmnImm64")
	cmnImm32     = form("CmnImm32")
	negShifted64 = form("NegShifted64")
	negShifted32 = form("NegShifted32")
)

// CmpShifted64 emits CMP Xn, Xm{, shift} — SUBS with the result discarded.
func (s *Section) CmpShifted64(rn, rm reg.X, shift ...ShiftOp) {
	s.inst(cmpShifted64, append([]any{rn, rm}, opt(shift)...)...)
}

// CmpImm64 emits CMP Xn|SP, #imm{, shift}.
func (s *Section) CmpImm64(rn RegSP64, imm int64, shift ...ShiftOp) {
	s.inst(cmpImm64, append([]any{rn, imm}, opt(shift)...)...)
}

// CmnShifted64 emits CMN Xn, Xm{, shift} — ADDS with the result discarded.
func (s *Section) CmnShifted64(rn, rm reg.X, shift ...ShiftOp) {
	s.inst(cmnShifted64, append([]any{rn, rm}, opt(shift)...)...)
}

// NegShifted64 emits NEG Xd, Xm{, shift} — SUB with Xn pinned to XZR.
func (s *Section) NegShifted64(rd, rm reg.X, shift ...ShiftOp) {
	s.inst(negShifted64, append([]any{rd, rm}, opt(shift)...)...)
}

// CmpShifted32 emits CMP Wn, Wm{, shift}.
func (s *Section) CmpShifted32(rn, rm reg.W, shift ...ShiftOp) {
	s.inst(cmpShifted32, append([]any{rn, rm}, opt(shift)...)...)
}

// CmpImm32 emits CMP Wn|WSP, #imm{, shift}.
func (s *Section) CmpImm32(rn RegSP32, imm int64, shift ...ShiftOp) {
	s.inst(cmpImm32, append([]any{rn, imm}, opt(shift)...)...)
}

// CmnShifted32 emits CMN Wn, Wm{, shift}.
func (s *Section) CmnShifted32(rn, rm reg.W, shift ...ShiftOp) {
	s.inst(cmnShifted32, append([]any{rn, rm}, opt(shift)...)...)
}

// CmnImm64 emits CMN Xn|SP, #imm{, shift} — ADDS with the result discarded.
func (s *Section) CmnImm64(rn RegSP64, imm int64, shift ...ShiftOp) {
	s.inst(cmnImm64, append([]any{rn, imm}, opt(shift)...)...)
}

// CmnImm32 emits CMN Wn|WSP, #imm{, shift}.
func (s *Section) CmnImm32(rn RegSP32, imm int64, shift ...ShiftOp) {
	s.inst(cmnImm32, append([]any{rn, imm}, opt(shift)...)...)
}

// NegShifted32 emits NEG Wd, Wm{, shift}.
func (s *Section) NegShifted32(rd, rm reg.W, shift ...ShiftOp) {
	s.inst(negShifted32, append([]any{rd, rm}, opt(shift)...)...)
}

// ---- Widening multiply -----------------------------------------------------

var (
	mul32 = form("Mul32")
	smulh = form("Smulh")
	umulh = form("Umulh")
)

// Smulh emits SMULH Xd, Xn, Xm: the high 64 bits of the signed product, which
// MUL throws away and no MADD form can reach.
func (s *Section) Smulh(rd, rn, rm reg.X) { s.inst(smulh, rd, rn, rm) }

// Umulh emits UMULH Xd, Xn, Xm, the unsigned half.
func (s *Section) Umulh(rd, rn, rm reg.X) { s.inst(umulh, rd, rn, rm) }

// Mul32 emits MUL Wd, Wn, Wm, which is MADD with the zero register.
func (s *Section) Mul32(rd, rn, rm reg.W) { s.inst(mul32, rd, rn, rm) }
