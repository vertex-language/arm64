package arm64_test

// Differential testing against a reference assembler: every case below emits
// the same instruction two ways — through this package's typed helpers, and
// through clang's integrated ARM64 assembler — and the two are required to
// produce the identical instruction word. A table row that decoded correctly
// but encoded wrong would pass every self-consistency check internal/isa runs
// at init and still be wrong; this is what catches that.
//
// It is skipped rather than failed when clang is not on PATH, because a
// missing reference assembler is an environment fact, not a defect in this
// package.

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/vertex-language/arm64"
)

// diffCase is one instruction: the GNU-syntax source line clang assembles,
// and the typed-helper call that should produce the identical word. build
// receives a fresh Section per case's word — cases share one Section so a
// same-section branch target ("1:") resolves the same way on both sides.
type diffCase struct {
	name string
	asm  string
	// nwords is how many instruction words this case's asm line assembles to.
	// Every case here is one, except none currently — kept for a future case
	// that needs to state a literal pool entry alongside its load.
	build func(t *arm64.Section)
}

func diffCases() []diffCase {
	x0, x1, x2, x3 := arm64.X0, arm64.X1, arm64.X2, arm64.X3
	w0, w1, w2, w3 := arm64.W0, arm64.W1, arm64.W2, arm64.W3
	s0, s1, s2, s3 := arm64.S0, arm64.S1, arm64.S2, arm64.S3
	d0, d1, d2, d3 := arm64.D0, arm64.D1, arm64.D2, arm64.D3
	q0 := arm64.Q0
	return []diffCase{
		{"add x0,x1,x2", "add x0, x1, x2", func(t *arm64.Section) { t.AddShifted64(x0, x1, x2) }},
		{"add w0,w1,w2,lsl3", "add w0, w1, w2, lsl #3", func(t *arm64.Section) {
			t.AddShifted32(w0, w1, w2, arm64.Shifted(arm64.LSL, 3))
		}},
		{"add x0,x1,#16", "add x0, x1, #16", func(t *arm64.Section) { t.AddImm64(x0, x1, 16) }},
		{"add x0,x1,#16,lsl12", "add x0, x1, #16, lsl #12", func(t *arm64.Section) {
			t.AddImm64(x0, x1, 16, arm64.Shifted(arm64.LSL, 12))
		}},
		{"sub sp,sp,#32", "sub sp, sp, #32", func(t *arm64.Section) { t.SubImm64(arm64.SP, arm64.SP, 32) }},
		{"adds w3,w4,w5", "adds w3, w4, w5", func(t *arm64.Section) { t.AddsShifted32(w3, arm64.W4, arm64.W5) }},
		{"subs x3,x4,x5", "subs x3, x4, x5", func(t *arm64.Section) { t.SubsShifted64(x3, arm64.X4, arm64.X5) }},
		{"and x6,x7,#0xff", "and x6, x7, #0xff", func(t *arm64.Section) { t.AndImm64(arm64.X6, arm64.X7, 0xff) }},
		{"ands w8,w9,#0xf0", "ands w8, w9, #0xf0", func(t *arm64.Section) { t.AndsImm32(arm64.W8, arm64.W9, 0xf0) }},
		{"orr x8,xzr,x9", "orr x8, xzr, x9", func(t *arm64.Section) { t.OrrShifted64(arm64.X8, arm64.XZR, arm64.X9) }},
		{"eor w1,w2,w3", "eor w1, w2, w3", func(t *arm64.Section) { t.EorShifted32(w1, w2, w3) }},
		{"movz x0,#42", "movz x0, #42", func(t *arm64.Section) { t.MovzImm64(x0, 42) }},
		{"movz w0,#42", "movz w0, #42", func(t *arm64.Section) { t.MovzImm32(w0, 42) }},
		{"movn x1,#5", "movn x1, #5", func(t *arm64.Section) { t.MovnImm64(x1, 5) }},
		{"movk x1,#5,lsl16", "movk x1, #5, lsl #16", func(t *arm64.Section) {
			t.MovkImm64(x1, 5, arm64.Shifted(arm64.LSL, 16))
		}},
		{"bl memcpy", "bl memcpy", func(t *arm64.Section) { t.Bl(arm64.Ref("memcpy")) }},
		{"ret", "ret", func(t *arm64.Section) { t.Ret() }},
		{"br x5", "br x5", func(t *arm64.Section) { t.Br(arm64.X5) }},
		{"blr x6", "blr x6", func(t *arm64.Section) { t.Blr(arm64.X6) }},
		{"stp x29,x30,[sp,-16]!", "stp x29, x30, [sp, #-16]!", func(t *arm64.Section) {
			t.StpPre64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Pre(-16))
		}},
		{"ldp x29,x30,[sp],16", "ldp x29, x30, [sp], #16", func(t *arm64.Section) {
			t.LdpPost64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Post(16))
		}},
		{"stp w0,w1,[sp,16]", "stp w0, w1, [sp, #16]", func(t *arm64.Section) {
			t.Stp32(w0, w1, arm64.Mem32(arm64.SP).Off(16))
		}},
		{"csel x0,x1,x2,gt", "csel x0, x1, x2, gt", func(t *arm64.Section) { t.Csel64(x0, x1, x2, arm64.GT) }},
		{"csinc w0,w1,w2,le", "csinc w0, w1, w2, le", func(t *arm64.Section) { t.Csinc32(w0, w1, w2, arm64.LE) }},
		{"csinv x0,x1,x2,ne", "csinv x0, x1, x2, ne", func(t *arm64.Section) { t.Csinv64(x0, x1, x2, arm64.NE) }},
		{"csneg x0,x1,x2,cc", "csneg x0, x1, x2, cc", func(t *arm64.Section) { t.Csneg64(x0, x1, x2, arm64.CC) }},
		{"cset x3,eq", "cset x3, eq", func(t *arm64.Section) { t.Cset64(x3, arm64.EQ) }},
		{"mul x0,x1,x2", "mul x0, x1, x2", func(t *arm64.Section) { t.Mul64(x0, x1, x2) }},
		{"madd x0,x1,x2,x3", "madd x0, x1, x2, x3", func(t *arm64.Section) { t.Madd64(x0, x1, x2, x3) }},
		{"msub w0,w1,w2,w3", "msub w0, w1, w2, w3", func(t *arm64.Section) { t.Msub32(w0, w1, w2, w3) }},
		{"udiv x0,x1,x2", "udiv x0, x1, x2", func(t *arm64.Section) { t.Udiv64(x0, x1, x2) }},
		{"sdiv w0,w1,w2", "sdiv w0, w1, w2", func(t *arm64.Section) { t.Sdiv32(w0, w1, w2) }},
		{"lsl x0,x1,#4", "lsl x0, x1, #4", func(t *arm64.Section) { t.LslImm64(x0, x1, 4) }},
		{"lsr x0,x1,#4", "lsr x0, x1, #4", func(t *arm64.Section) { t.LsrImm64(x0, x1, 4) }},
		{"asr x0,x1,#4", "asr x0, x1, #4", func(t *arm64.Section) { t.AsrImm64(x0, x1, 4) }},
		{"lslv x0,x1,x2", "lsl x0, x1, x2", func(t *arm64.Section) { t.Lslv64(x0, x1, x2) }},
		{"lsrv w0,w1,w2", "lsr w0, w1, w2", func(t *arm64.Section) { t.Lsrv32(w0, w1, w2) }},
		{"ubfm x0,x1,4,10", "ubfm x0, x1, #4, #10", func(t *arm64.Section) { t.Ubfm64(x0, x1, 4, 10) }},
		{"sbfm w0,w1,2,6", "sbfm w0, w1, #2, #6", func(t *arm64.Section) { t.Sbfm32(w0, w1, 2, 6) }},
		{"bfm x0,x1,1,3", "bfm x0, x1, #1, #3", func(t *arm64.Section) { t.Bfm64(x0, x1, 1, 3) }},
		{"extr x0,x1,x2,5", "extr x0, x1, x2, #5", func(t *arm64.Section) { t.Extr64(x0, x1, x2, 5) }},
		{"rbit x0,x1", "rbit x0, x1", func(t *arm64.Section) { t.Rbit64(x0, x1) }},
		{"rev x0,x1", "rev x0, x1", func(t *arm64.Section) { t.Rev64(x0, x1) }},
		{"rev16 w0,w1", "rev16 w0, w1", func(t *arm64.Section) { t.Rev16_32(w0, w1) }},
		{"clz x0,x1", "clz x0, x1", func(t *arm64.Section) { t.Clz64(x0, x1) }},
		{"cls w0,w1", "cls w0, w1", func(t *arm64.Section) { t.Cls32(w0, w1) }},
		{"strb w0,[x1,4]", "strb w0, [x1, #4]", func(t *arm64.Section) { t.StrbImm(w0, arm64.Mem8(x1).Off(4)) }},
		{"ldrb w0,[x1,4]", "ldrb w0, [x1, #4]", func(t *arm64.Section) { t.LdrbImm(w0, arm64.Mem8(x1).Off(4)) }},
		{"strh w0,[x1,8]", "strh w0, [x1, #8]", func(t *arm64.Section) { t.StrhImm(w0, arm64.Mem16(x1).Off(8)) }},
		{"ldrh w0,[x1,8]", "ldrh w0, [x1, #8]", func(t *arm64.Section) { t.LdrhImm(w0, arm64.Mem16(x1).Off(8)) }},
		{"str w0,[x1,12]", "str w0, [x1, #12]", func(t *arm64.Section) { t.StrImm32(w0, arm64.Mem32(x1).Off(12)) }},
		{"ldr w0,[x1,12]", "ldr w0, [x1, #12]", func(t *arm64.Section) { t.LdrImm32(w0, arm64.Mem32(x1).Off(12)) }},
		{"str x0,[x1,16]", "str x0, [x1, #16]", func(t *arm64.Section) { t.StrImm64(x0, arm64.Mem64(x1).Off(16)) }},
		{"ldr x0,[x1,16]", "ldr x0, [x1, #16]", func(t *arm64.Section) { t.LdrImm64(x0, arm64.Mem64(x1).Off(16)) }},
		{"ldrsw x0,[x1,4]", "ldrsw x0, [x1, #4]", func(t *arm64.Section) { t.LdrswImm(x0, arm64.Mem32(x1).Off(4)) }},
		{"stur w0,[x1,-8]", "stur w0, [x1, #-8]", func(t *arm64.Section) { t.SturImm32(w0, arm64.Mem32(x1).Off(-8)) }},
		{"ldur x0,[x1,-8]", "ldur x0, [x1, #-8]", func(t *arm64.Section) { t.LdurImm64(x0, arm64.Mem64(x1).Off(-8)) }},
		{"svc #0", "svc #0", func(t *arm64.Section) { t.Svc(0) }},
		{"brk #1", "brk #1", func(t *arm64.Section) { t.Brk(1) }},
		{"hlt #2", "hlt #2", func(t *arm64.Section) { t.Hlt(2) }},
		{"nop", "nop", func(t *arm64.Section) { t.Nop() }},
		{"dsb sy", "dsb sy", func(t *arm64.Section) { t.Dsb(arm64.SY) }},
		{"dmb ish", "dmb ish", func(t *arm64.Section) { t.Dmb(arm64.ISH) }},
		{"isb", "isb", func(t *arm64.Section) { t.Isb() }},
		{"clrex", "clrex", func(t *arm64.Section) { t.Clrex() }},
		{"sub sp,sp,x3", "sub sp, sp, x3", func(t *arm64.Section) { t.SubExt64(arm64.SP, arm64.SP, x3) }},
		{"sub x0,sp,x3", "sub x0, sp, x3", func(t *arm64.Section) { t.SubExt64(x0, arm64.SP, x3) }},
		{"mrs x0,nzcv", "mrs x0, nzcv", func(t *arm64.Section) { t.Mrs(x0, arm64.NZCV) }},
		{"msr nzcv,x0", "msr nzcv, x0", func(t *arm64.Section) { t.MsrReg(arm64.NZCV, x0) }},
		{"tst x0,x1", "tst x0, x1", func(t *arm64.Section) { t.TstShifted64(x0, x1) }},
		{"cmp x0,x1", "cmp x0, x1", func(t *arm64.Section) { t.CmpShifted64(x0, x1) }},
		{"cmp x0,#4", "cmp x0, #4", func(t *arm64.Section) { t.CmpImm64(x0, 4) }},
		{"cmn x0,x1", "cmn x0, x1", func(t *arm64.Section) { t.CmnShifted64(x0, x1) }},
		{"neg x0,x1", "neg x0, x1", func(t *arm64.Section) { t.NegShifted64(x0, x1) }},

		// The 32-bit halves of the same aliases, and the two the table had
		// no row for at either width.
		{"tst w0,w1", "tst w0, w1", func(t *arm64.Section) { t.TstShifted32(w0, w1) }},
		{"tst w0,#0xff", "tst w0, #0xff", func(t *arm64.Section) { t.TstImm32(w0, 0xff) }},
		{"cmp w0,w1", "cmp w0, w1", func(t *arm64.Section) { t.CmpShifted32(w0, w1) }},
		{"cmp w0,#4", "cmp w0, #4", func(t *arm64.Section) { t.CmpImm32(w0, 4) }},
		{"cmn w0,w1", "cmn w0, w1", func(t *arm64.Section) { t.CmnShifted32(w0, w1) }},
		{"cmn x0,#4", "cmn x0, #4", func(t *arm64.Section) { t.CmnImm64(x0, 4) }},
		{"cmn w0,#4", "cmn w0, #4", func(t *arm64.Section) { t.CmnImm32(w0, 4) }},
		{"neg w0,w1", "neg w0, w1", func(t *arm64.Section) { t.NegShifted32(w0, w1) }},
		{"mvn x0,x1", "mvn x0, x1", func(t *arm64.Section) { t.MvnShifted64(x0, x1) }},
		{"mvn w0,w1", "mvn w0, w1", func(t *arm64.Section) { t.MvnShifted32(w0, w1) }},
		{"mvn x0,x1,lsl2", "mvn x0, x1, lsl #2", func(t *arm64.Section) {
			t.MvnShifted64(x0, x1, arm64.Shifted(arm64.LSL, 2))
		}},

		// The shifts by an immediate at 32 bits, where the computed immr:imms
		// pair is against a width of 32 and not 64.
		{"lsl w0,w1,#3", "lsl w0, w1, #3", func(t *arm64.Section) { t.LslImm32(w0, w1, 3) }},
		{"lsr w0,w1,#7", "lsr w0, w1, #7", func(t *arm64.Section) { t.LsrImm32(w0, w1, 7) }},
		{"asr w0,w1,#9", "asr w0, w1, #9", func(t *arm64.Section) { t.AsrImm32(w0, w1, 9) }},

		// The shifts by a register under their assembly-language names.
		{"lsl x0,x1,x2 (alias)", "lsl x0, x1, x2", func(t *arm64.Section) { t.LslReg64(x0, x1, x2) }},
		{"lsr w0,w1,w2 (alias)", "lsr w0, w1, w2", func(t *arm64.Section) { t.LsrReg32(w0, w1, w2) }},
		{"asr x0,x1,x2 (alias)", "asr x0, x1, x2", func(t *arm64.Section) { t.AsrReg64(x0, x1, x2) }},
		{"ror w0,w1,w2 (alias)", "ror w0, w1, w2", func(t *arm64.Section) { t.RorReg32(w0, w1, w2) }},

		// The extend aliases, whose 64-bit destinations take a W source.
		{"sxtb w0,w1", "sxtb w0, w1", func(t *arm64.Section) { t.Sxtb32(w0, w1) }},
		{"sxtb x0,w1", "sxtb x0, w1", func(t *arm64.Section) { t.Sxtb64(x0, w1) }},
		{"sxth w0,w1", "sxth w0, w1", func(t *arm64.Section) { t.Sxth32(w0, w1) }},
		{"sxth x0,w1", "sxth x0, w1", func(t *arm64.Section) { t.Sxth64(x0, w1) }},
		{"sxtw x0,w1", "sxtw x0, w1", func(t *arm64.Section) { t.Sxtw64(x0, w1) }},
		{"uxtb w0,w1", "uxtb w0, w1", func(t *arm64.Section) { t.Uxtb32(w0, w1) }},
		{"uxth w0,w1", "uxth w0, w1", func(t *arm64.Section) { t.Uxth32(w0, w1) }},

		// A prefetch, whose first operand is a hint and not a register.
		{"prfm pldl1keep,[x0]", "prfm pldl1keep, [x0]", func(t *arm64.Section) {
			t.PrfmImm(arm64.PLDL1KEEP, arm64.Mem64(x0))
		}},
		{"prfm pstl2strm,[x1,#16]", "prfm pstl2strm, [x1, #16]", func(t *arm64.Section) {
			t.PrfmImm(arm64.PSTL2STRM, arm64.Mem64(x1).Off(16))
		}},

		// Bitfield extraction, whose second immediate is computed from the
		// first, and the rotate that names one source twice.
		{"ubfx x0,x1,#4,#8", "ubfx x0, x1, #4, #8", func(t *arm64.Section) { t.Ubfx64(x0, x1, 4, 8) }},
		{"ubfx w0,w1,#2,#3", "ubfx w0, w1, #2, #3", func(t *arm64.Section) { t.Ubfx32(w0, w1, 2, 3) }},
		{"sbfx x0,x1,#4,#8", "sbfx x0, x1, #4, #8", func(t *arm64.Section) { t.Sbfx64(x0, x1, 4, 8) }},
		{"sbfx w0,w1,#1,#5", "sbfx w0, w1, #1, #5", func(t *arm64.Section) { t.Sbfx32(w0, w1, 1, 5) }},
		{"ror x0,x1,#5", "ror x0, x1, #5", func(t *arm64.Section) { t.RorImm64(x0, x1, 5) }},
		{"ror w0,w1,#3", "ror w0, w1, #3", func(t *arm64.Section) { t.RorImm32(w0, w1, 3) }},

		// Register-offset addressing, built with Mem.Indexed. The extend
		// says which width the index is read at, and mismatching the two
		// is what Mem.Validate refuses.
		{"ldr x0,[x1,x2]", "ldr x0, [x1, x2]", func(t *arm64.Section) {
			t.LdrReg64(x0, arm64.Mem64(x1).Indexed(x2, arm64.ExtLSL, 0))
		}},
		{"ldr x0,[x1,x2,lsl3]", "ldr x0, [x1, x2, lsl #3]", func(t *arm64.Section) {
			t.LdrReg64(x0, arm64.Mem64(x1).Indexed(x2, arm64.ExtLSL, 3))
		}},
		{"ldr w0,[x1,x2,lsl2]", "ldr w0, [x1, x2, lsl #2]", func(t *arm64.Section) {
			t.LdrReg32(w0, arm64.Mem32(x1).Indexed(x2, arm64.ExtLSL, 2))
		}},
		{"str x0,[x1,x2]", "str x0, [x1, x2]", func(t *arm64.Section) {
			t.StrReg64(x0, arm64.Mem64(x1).Indexed(x2, arm64.ExtLSL, 0))
		}},
		{"ldr x0,[x1,w2,uxtw3]", "ldr x0, [x1, w2, uxtw #3]", func(t *arm64.Section) {
			t.LdrReg64(x0, arm64.Mem64(x1).Indexed(w2, arm64.UXTW, 3))
		}},
		{"ldr x0,[x1,w2,sxtw]", "ldr x0, [x1, w2, sxtw]", func(t *arm64.Section) {
			t.LdrReg64(x0, arm64.Mem64(x1).Indexed(w2, arm64.SXTW, 0))
		}},
		{"ldrb w0,[x1,x2]", "ldrb w0, [x1, x2]", func(t *arm64.Section) {
			t.LdrbReg(w0, arm64.Mem8(x1).Indexed(x2, arm64.ExtLSL, 0))
		}},
		{"ldrsw x0,[x1,x2]", "ldrsw x0, [x1, x2]", func(t *arm64.Section) {
			t.LdrswReg(x0, arm64.Mem32(x1).Indexed(x2, arm64.ExtLSL, 0))
		}},
		{"prfm pldl1keep,[x1,x2]", "prfm pldl1keep, [x1, x2]", func(t *arm64.Section) {
			t.PrfmReg(arm64.PLDL1KEEP, arm64.Mem64(x1).Indexed(x2, arm64.ExtLSL, 0))
		}},

		// Add and subtract with carry, and the aliases on them.
		{"adc x0,x1,x2", "adc x0, x1, x2", func(t *arm64.Section) { t.Adc64(x0, x1, x2) }},
		{"adc w0,w1,w2", "adc w0, w1, w2", func(t *arm64.Section) { t.Adc32(w0, w1, w2) }},
		{"adcs x0,x1,x2", "adcs x0, x1, x2", func(t *arm64.Section) { t.Adcs64(x0, x1, x2) }},
		{"adcs w0,w1,w2", "adcs w0, w1, w2", func(t *arm64.Section) { t.Adcs32(w0, w1, w2) }},
		{"sbc x0,x1,x2", "sbc x0, x1, x2", func(t *arm64.Section) { t.Sbc64(x0, x1, x2) }},
		{"sbc w0,w1,w2", "sbc w0, w1, w2", func(t *arm64.Section) { t.Sbc32(w0, w1, w2) }},
		{"sbcs x0,x1,x2", "sbcs x0, x1, x2", func(t *arm64.Section) { t.Sbcs64(x0, x1, x2) }},
		{"sbcs w0,w1,w2", "sbcs w0, w1, w2", func(t *arm64.Section) { t.Sbcs32(w0, w1, w2) }},
		{"negs x0,x1", "negs x0, x1", func(t *arm64.Section) { t.Negs64(x0, x1) }},
		{"negs w0,w1", "negs w0, w1", func(t *arm64.Section) { t.Negs32(w0, w1) }},
		{"ngc x0,x1", "ngc x0, x1", func(t *arm64.Section) { t.Ngc64(x0, x1) }},
		{"ngc w0,w1", "ngc w0, w1", func(t *arm64.Section) { t.Ngc32(w0, w1) }},
		{"ngcs x0,x1", "ngcs x0, x1", func(t *arm64.Section) { t.Ngcs64(x0, x1) }},
		{"ngcs w0,w1", "ngcs w0, w1", func(t *arm64.Section) { t.Ngcs32(w0, w1) }},

		// The widening multiply: W sources, an X destination.
		{"smaddl x0,w1,w2,x3", "smaddl x0, w1, w2, x3", func(t *arm64.Section) {
			t.Smaddl(x0, w1, w2, x3)
		}},
		{"umaddl x0,w1,w2,x3", "umaddl x0, w1, w2, x3", func(t *arm64.Section) {
			t.Umaddl(x0, w1, w2, x3)
		}},
		{"smsubl x0,w1,w2,x3", "smsubl x0, w1, w2, x3", func(t *arm64.Section) {
			t.Smsubl(x0, w1, w2, x3)
		}},
		{"umsubl x0,w1,w2,x3", "umsubl x0, w1, w2, x3", func(t *arm64.Section) {
			t.Umsubl(x0, w1, w2, x3)
		}},
		{"smull x0,w1,w2", "smull x0, w1, w2", func(t *arm64.Section) { t.Smull(x0, w1, w2) }},
		{"umull x0,w1,w2", "umull x0, w1, w2", func(t *arm64.Section) { t.Umull(x0, w1, w2) }},
		{"smnegl x0,w1,w2", "smnegl x0, w1, w2", func(t *arm64.Section) { t.Smnegl(x0, w1, w2) }},
		{"umnegl x0,w1,w2", "umnegl x0, w1, w2", func(t *arm64.Section) { t.Umnegl(x0, w1, w2) }},

		// The conditional aliases, which name one source twice and invert
		// their condition — both on the row, so this is what checks that
		// the row does it and not the helper.
		{"cset w0,ne", "cset w0, ne", func(t *arm64.Section) { t.Cset32(w0, arm64.NE) }},
		{"csetm x0,ne", "csetm x0, ne", func(t *arm64.Section) { t.Csetm64(x0, arm64.NE) }},
		{"csetm w0,ne", "csetm w0, ne", func(t *arm64.Section) { t.Csetm32(w0, arm64.NE) }},
		{"cinc x0,x1,ne", "cinc x0, x1, ne", func(t *arm64.Section) { t.Cinc64(x0, x1, arm64.NE) }},
		{"cinc w0,w1,ne", "cinc w0, w1, ne", func(t *arm64.Section) { t.Cinc32(w0, w1, arm64.NE) }},
		{"cinv x0,x1,ne", "cinv x0, x1, ne", func(t *arm64.Section) { t.Cinv64(x0, x1, arm64.NE) }},
		{"cinv w0,w1,ne", "cinv w0, w1, ne", func(t *arm64.Section) { t.Cinv32(w0, w1, arm64.NE) }},
		{"cneg x0,x1,ne", "cneg x0, x1, ne", func(t *arm64.Section) { t.Cneg64(x0, x1, arm64.NE) }},
		{"cneg w0,w1,ne", "cneg w0, w1, ne", func(t *arm64.Section) { t.Cneg32(w0, w1, arm64.NE) }},

		// Bitfield insert, where immr rotates the source into position.
		{"bfi x0,x1,#4,#8", "bfi x0, x1, #4, #8", func(t *arm64.Section) { t.Bfi64(x0, x1, 4, 8) }},
		{"bfi w0,w1,#4,#8", "bfi w0, w1, #4, #8", func(t *arm64.Section) { t.Bfi32(w0, w1, 4, 8) }},
		{"bfxil x0,x1,#4,#8", "bfxil x0, x1, #4, #8", func(t *arm64.Section) {
			t.Bfxil64(x0, x1, 4, 8)
		}},
		{"bfxil w0,w1,#4,#8", "bfxil w0, w1, #4, #8", func(t *arm64.Section) {
			t.Bfxil32(w0, w1, 4, 8)
		}},
		{"ubfiz x0,x1,#4,#8", "ubfiz x0, x1, #4, #8", func(t *arm64.Section) {
			t.Ubfiz64(x0, x1, 4, 8)
		}},
		{"ubfiz w0,w1,#4,#8", "ubfiz w0, w1, #4, #8", func(t *arm64.Section) {
			t.Ubfiz32(w0, w1, 4, 8)
		}},
		{"sbfiz x0,x1,#4,#8", "sbfiz x0, x1, #4, #8", func(t *arm64.Section) {
			t.Sbfiz64(x0, x1, 4, 8)
		}},
		{"sbfiz w0,w1,#4,#8", "sbfiz w0, w1, #4, #8", func(t *arm64.Section) {
			t.Sbfiz32(w0, w1, 4, 8)
		}},

		// The hint space, the exception return, and the two byte reverses
		// that are not REV.
		{"hint #20", "hint #0x14", func(t *arm64.Section) { t.Hint(0x14) }},
		{"paciasp", "paciasp", func(t *arm64.Section) { t.Paciasp() }},
		{"autiasp", "autiasp", func(t *arm64.Section) { t.Autiasp() }},
		{"eret", "eret", func(t *arm64.Section) { t.Eret() }},
		{"rev64 x0,x1", "rev64 x0, x1", func(t *arm64.Section) { t.Rev64Alias(x0, x1) }},
		{"rev32 x0,x1", "rev32 x0, x1", func(t *arm64.Section) { t.Rev32In64(x0, x1) }},

		// Conditional compare, at both widths and with both operand shapes.
		{"ccmp x0,x1,#0,eq", "ccmp x0, x1, #0, eq", func(t *arm64.Section) {
			t.CcmpReg64(x0, x1, 0, arm64.EQ)
		}},
		{"ccmp w0,w1,#3,ne", "ccmp w0, w1, #3, ne", func(t *arm64.Section) {
			t.CcmpReg32(w0, w1, 3, arm64.NE)
		}},
		{"ccmp x0,#5,#0,eq", "ccmp x0, #5, #0, eq", func(t *arm64.Section) {
			t.CcmpImm64(x0, 5, 0, arm64.EQ)
		}},
		{"ccmp w0,#5,#7,lt", "ccmp w0, #5, #7, lt", func(t *arm64.Section) {
			t.CcmpImm32(w0, 5, 7, arm64.LT)
		}},
		{"ccmn x0,x1,#0,eq", "ccmn x0, x1, #0, eq", func(t *arm64.Section) {
			t.CcmnReg64(x0, x1, 0, arm64.EQ)
		}},
		{"ccmn w0,#5,#7,lt", "ccmn w0, #5, #7, lt", func(t *arm64.Section) {
			t.CcmnImm32(w0, 5, 7, arm64.LT)
		}},
		{"mov x0,x1", "mov x0, x1", func(t *arm64.Section) { t.MovReg64(x0, x1) }},
		{"mov w0,w1", "mov w0, w1", func(t *arm64.Section) { t.MovReg32(w0, w1) }},
		{"mov x0,sp", "mov x0, sp", func(t *arm64.Section) { t.MovSp64(x0, arm64.SP) }},
		{"mov sp,x0", "mov sp, x0", func(t *arm64.Section) { t.MovSp64(arm64.SP, x0) }},
		{"mov x0,#100", "mov x0, #100", func(t *arm64.Section) { t.MovWide64(x0, 100) }},
		// AddExt64 with a 32-bit source (UXTW/SXTW reading Wm) is not covered
		// here: the table declares AddExt64's Rm as ClassX only, so a Wm
		// source is refused today. See Known limitations in the README.
		{"add sp,x1,x2", "add sp, x1, x2", func(t *arm64.Section) {
			t.AddExt64(arm64.SP, x1, x2, arm64.Extended(arm64.ExtLSL, 0))
		}},
		{"smulh x0,x1,x2", "smulh x0, x1, x2", func(t *arm64.Section) { t.Smulh(x0, x1, x2) }},
		{"umulh x0,x1,x2", "umulh x0, x1, x2", func(t *arm64.Section) { t.Umulh(x0, x1, x2) }},

		// ---- Scalar floating point ----
		{"fadd s0,s1,s2", "fadd s0, s1, s2", func(t *arm64.Section) { t.FaddS(s0, s1, s2) }},
		{"fadd d0,d1,d2", "fadd d0, d1, d2", func(t *arm64.Section) { t.FaddD(d0, d1, d2) }},
		{"fsub s0,s1,s2", "fsub s0, s1, s2", func(t *arm64.Section) { t.FsubS(s0, s1, s2) }},
		{"fsub d0,d1,d2", "fsub d0, d1, d2", func(t *arm64.Section) { t.FsubD(d0, d1, d2) }},
		{"fmul s0,s1,s2", "fmul s0, s1, s2", func(t *arm64.Section) { t.FmulS(s0, s1, s2) }},
		{"fmul d0,d1,d2", "fmul d0, d1, d2", func(t *arm64.Section) { t.FmulD(d0, d1, d2) }},
		{"fdiv s0,s1,s2", "fdiv s0, s1, s2", func(t *arm64.Section) { t.FdivS(s0, s1, s2) }},
		{"fdiv d0,d1,d2", "fdiv d0, d1, d2", func(t *arm64.Section) { t.FdivD(d0, d1, d2) }},
		{"fmax s0,s1,s2", "fmax s0, s1, s2", func(t *arm64.Section) { t.FmaxS(s0, s1, s2) }},
		{"fmax d0,d1,d2", "fmax d0, d1, d2", func(t *arm64.Section) { t.FmaxD(d0, d1, d2) }},
		{"fmin s0,s1,s2", "fmin s0, s1, s2", func(t *arm64.Section) { t.FminS(s0, s1, s2) }},
		{"fmin d0,d1,d2", "fmin d0, d1, d2", func(t *arm64.Section) { t.FminD(d0, d1, d2) }},
		{"fmaxnm s0,s1,s2", "fmaxnm s0, s1, s2", func(t *arm64.Section) { t.FmaxnmS(s0, s1, s2) }},
		{"fmaxnm d0,d1,d2", "fmaxnm d0, d1, d2", func(t *arm64.Section) { t.FmaxnmD(d0, d1, d2) }},
		{"fminnm s0,s1,s2", "fminnm s0, s1, s2", func(t *arm64.Section) { t.FminnmS(s0, s1, s2) }},
		{"fminnm d0,d1,d2", "fminnm d0, d1, d2", func(t *arm64.Section) { t.FminnmD(d0, d1, d2) }},

		{"fmov s0,s1", "fmov s0, s1", func(t *arm64.Section) { t.FmovS(s0, s1) }},
		{"fmov d0,d1", "fmov d0, d1", func(t *arm64.Section) { t.FmovD(d0, d1) }},
		{"fabs s0,s1", "fabs s0, s1", func(t *arm64.Section) { t.FabsS(s0, s1) }},
		{"fabs d0,d1", "fabs d0, d1", func(t *arm64.Section) { t.FabsD(d0, d1) }},
		{"fneg s0,s1", "fneg s0, s1", func(t *arm64.Section) { t.FnegS(s0, s1) }},
		{"fneg d0,d1", "fneg d0, d1", func(t *arm64.Section) { t.FnegD(d0, d1) }},
		{"fsqrt s0,s1", "fsqrt s0, s1", func(t *arm64.Section) { t.FsqrtS(s0, s1) }},
		{"fsqrt d0,d1", "fsqrt d0, d1", func(t *arm64.Section) { t.FsqrtD(d0, d1) }},
		{"frintn s0,s1", "frintn s0, s1", func(t *arm64.Section) { t.FrintnS(s0, s1) }},
		{"frintn d0,d1", "frintn d0, d1", func(t *arm64.Section) { t.FrintnD(d0, d1) }},
		{"frintp s0,s1", "frintp s0, s1", func(t *arm64.Section) { t.FrintpS(s0, s1) }},
		{"frintp d0,d1", "frintp d0, d1", func(t *arm64.Section) { t.FrintpD(d0, d1) }},
		{"frintm s0,s1", "frintm s0, s1", func(t *arm64.Section) { t.FrintmS(s0, s1) }},
		{"frintm d0,d1", "frintm d0, d1", func(t *arm64.Section) { t.FrintmD(d0, d1) }},
		{"frintz s0,s1", "frintz s0, s1", func(t *arm64.Section) { t.FrintzS(s0, s1) }},
		{"frintz d0,d1", "frintz d0, d1", func(t *arm64.Section) { t.FrintzD(d0, d1) }},

		{"fcvt d0,s1", "fcvt d0, s1", func(t *arm64.Section) { t.FcvtSToD(d0, s1) }},
		{"fcvt s0,d1", "fcvt s0, d1", func(t *arm64.Section) { t.FcvtDToS(s0, d1) }},

		{"fmadd s0,s1,s2,s3", "fmadd s0, s1, s2, s3", func(t *arm64.Section) { t.FmaddS(s0, s1, s2, s3) }},
		{"fmadd d0,d1,d2,d3", "fmadd d0, d1, d2, d3", func(t *arm64.Section) { t.FmaddD(d0, d1, d2, d3) }},
		{"fmsub s0,s1,s2,s3", "fmsub s0, s1, s2, s3", func(t *arm64.Section) { t.FmsubS(s0, s1, s2, s3) }},
		{"fmsub d0,d1,d2,d3", "fmsub d0, d1, d2, d3", func(t *arm64.Section) { t.FmsubD(d0, d1, d2, d3) }},

		{"fcmp s0,s1", "fcmp s0, s1", func(t *arm64.Section) { t.FcmpS(s0, s1) }},
		{"fcmp d0,d1", "fcmp d0, d1", func(t *arm64.Section) { t.FcmpD(d0, d1) }},
		{"fcmp s0,#0.0", "fcmp s0, #0.0", func(t *arm64.Section) { t.FcmpZeroS(s0) }},
		{"fcmp d0,#0.0", "fcmp d0, #0.0", func(t *arm64.Section) { t.FcmpZeroD(d0) }},
		{"fcsel s0,s1,s2,eq", "fcsel s0, s1, s2, eq", func(t *arm64.Section) { t.FcselS(s0, s1, s2, arm64.EQ) }},
		{"fcsel d0,d1,d2,mi", "fcsel d0, d1, d2, mi", func(t *arm64.Section) { t.FcselD(d0, d1, d2, arm64.MI) }},

		{"scvtf s0,w1", "scvtf s0, w1", func(t *arm64.Section) { t.ScvtfWToS(s0, w1) }},
		{"scvtf d0,w1", "scvtf d0, w1", func(t *arm64.Section) { t.ScvtfWToD(d0, w1) }},
		{"scvtf s0,x1", "scvtf s0, x1", func(t *arm64.Section) { t.ScvtfXToS(s0, x1) }},
		{"scvtf d0,x1", "scvtf d0, x1", func(t *arm64.Section) { t.ScvtfXToD(d0, x1) }},
		{"ucvtf s0,w1", "ucvtf s0, w1", func(t *arm64.Section) { t.UcvtfWToS(s0, w1) }},
		{"ucvtf d0,w1", "ucvtf d0, w1", func(t *arm64.Section) { t.UcvtfWToD(d0, w1) }},
		{"ucvtf s0,x1", "ucvtf s0, x1", func(t *arm64.Section) { t.UcvtfXToS(s0, x1) }},
		{"ucvtf d0,x1", "ucvtf d0, x1", func(t *arm64.Section) { t.UcvtfXToD(d0, x1) }},

		{"fcvtzs w0,s1", "fcvtzs w0, s1", func(t *arm64.Section) { t.FcvtzsSToW(w0, s1) }},
		{"fcvtzs w0,d1", "fcvtzs w0, d1", func(t *arm64.Section) { t.FcvtzsDToW(w0, d1) }},
		{"fcvtzs x0,s1", "fcvtzs x0, s1", func(t *arm64.Section) { t.FcvtzsSToX(x0, s1) }},
		{"fcvtzs x0,d1", "fcvtzs x0, d1", func(t *arm64.Section) { t.FcvtzsDToX(x0, d1) }},
		{"fcvtzu w0,s1", "fcvtzu w0, s1", func(t *arm64.Section) { t.FcvtzuSToW(w0, s1) }},
		{"fcvtzu w0,d1", "fcvtzu w0, d1", func(t *arm64.Section) { t.FcvtzuDToW(w0, d1) }},
		{"fcvtzu x0,s1", "fcvtzu x0, s1", func(t *arm64.Section) { t.FcvtzuSToX(x0, s1) }},
		{"fcvtzu x0,d1", "fcvtzu x0, d1", func(t *arm64.Section) { t.FcvtzuDToX(x0, d1) }},

		{"fmov w0,s1", "fmov w0, s1", func(t *arm64.Section) { t.FmovSToW(w0, s1) }},
		{"fmov s0,w1", "fmov s0, w1", func(t *arm64.Section) { t.FmovWToS(s0, w1) }},
		{"fmov x0,d1", "fmov x0, d1", func(t *arm64.Section) { t.FmovDToX(x0, d1) }},
		{"fmov d0,x1", "fmov d0, x1", func(t *arm64.Section) { t.FmovXToD(d0, x1) }},

		{"ldr s0,[x1,16]", "ldr s0, [x1, #16]", func(t *arm64.Section) { t.LdrImmS(s0, arm64.Mem32(x1).Off(16)) }},
		{"ldr d0,[x1,16]", "ldr d0, [x1, #16]", func(t *arm64.Section) { t.LdrImmD(d0, arm64.Mem64(x1).Off(16)) }},
		{"ldr q0,[x1,16]", "ldr q0, [x1, #16]", func(t *arm64.Section) { t.LdrImmQ(q0, arm64.Mem128(x1).Off(16)) }},
		{"str s0,[x1,16]", "str s0, [x1, #16]", func(t *arm64.Section) { t.StrImmS(s0, arm64.Mem32(x1).Off(16)) }},
		{"str d0,[x1,16]", "str d0, [x1, #16]", func(t *arm64.Section) { t.StrImmD(d0, arm64.Mem64(x1).Off(16)) }},
		{"str q0,[x1,16]", "str q0, [x1, #16]", func(t *arm64.Section) { t.StrImmQ(q0, arm64.Mem128(x1).Off(16)) }},
		{"ldur s0,[x1,-4]", "ldur s0, [x1, #-4]", func(t *arm64.Section) { t.LdurImmS(s0, arm64.Mem32(x1).Off(-4)) }},
		{"ldur d0,[x1,-4]", "ldur d0, [x1, #-4]", func(t *arm64.Section) { t.LdurImmD(d0, arm64.Mem64(x1).Off(-4)) }},
		{"stur s0,[x1,-4]", "stur s0, [x1, #-4]", func(t *arm64.Section) { t.SturImmS(s0, arm64.Mem32(x1).Off(-4)) }},
		{"stur d0,[x1,-4]", "stur d0, [x1, #-4]", func(t *arm64.Section) { t.SturImmD(d0, arm64.Mem64(x1).Off(-4)) }},

		// ---- Atomics ----
		{"ldar w0,[x1]", "ldar w0, [x1]", func(t *arm64.Section) { t.Ldar32(w0, arm64.Mem32(x1)) }},
		{"ldar x0,[x1]", "ldar x0, [x1]", func(t *arm64.Section) { t.Ldar64(x0, arm64.Mem64(x1)) }},
		{"ldarb w0,[x1]", "ldarb w0, [x1]", func(t *arm64.Section) { t.Ldarb(w0, arm64.Mem8(x1)) }},
		{"ldarh w0,[x1]", "ldarh w0, [x1]", func(t *arm64.Section) { t.Ldarh(w0, arm64.Mem16(x1)) }},
		{"stlr w0,[x1]", "stlr w0, [x1]", func(t *arm64.Section) { t.Stlr32(w0, arm64.Mem32(x1)) }},
		{"stlr x0,[x1]", "stlr x0, [x1]", func(t *arm64.Section) { t.Stlr64(x0, arm64.Mem64(x1)) }},
		{"stlrb w0,[x1]", "stlrb w0, [x1]", func(t *arm64.Section) { t.Stlrb(w0, arm64.Mem8(x1)) }},
		{"stlrh w0,[x1]", "stlrh w0, [x1]", func(t *arm64.Section) { t.Stlrh(w0, arm64.Mem16(x1)) }},

		{"ldxr w0,[x1]", "ldxr w0, [x1]", func(t *arm64.Section) { t.Ldxr32(w0, arm64.Mem32(x1)) }},
		{"ldxr x0,[x1]", "ldxr x0, [x1]", func(t *arm64.Section) { t.Ldxr64(x0, arm64.Mem64(x1)) }},
		{"ldxrb w0,[x1]", "ldxrb w0, [x1]", func(t *arm64.Section) { t.Ldxrb(w0, arm64.Mem8(x1)) }},
		{"ldxrh w0,[x1]", "ldxrh w0, [x1]", func(t *arm64.Section) { t.Ldxrh(w0, arm64.Mem16(x1)) }},
		{"ldaxr w0,[x1]", "ldaxr w0, [x1]", func(t *arm64.Section) { t.Ldaxr32(w0, arm64.Mem32(x1)) }},
		{"ldaxr x0,[x1]", "ldaxr x0, [x1]", func(t *arm64.Section) { t.Ldaxr64(x0, arm64.Mem64(x1)) }},
		{"ldaxrb w0,[x1]", "ldaxrb w0, [x1]", func(t *arm64.Section) { t.Ldaxrb(w0, arm64.Mem8(x1)) }},
		{"ldaxrh w0,[x1]", "ldaxrh w0, [x1]", func(t *arm64.Section) { t.Ldaxrh(w0, arm64.Mem16(x1)) }},

		{"stxr w2,w0,[x1]", "stxr w2, w0, [x1]", func(t *arm64.Section) { t.Stxr32(w2, w0, arm64.Mem32(x1)) }},
		{"stxr w2,x0,[x1]", "stxr w2, x0, [x1]", func(t *arm64.Section) { t.Stxr64(w2, x0, arm64.Mem64(x1)) }},
		{"stxrb w2,w0,[x1]", "stxrb w2, w0, [x1]", func(t *arm64.Section) { t.Stxrb(w2, w0, arm64.Mem8(x1)) }},
		{"stxrh w2,w0,[x1]", "stxrh w2, w0, [x1]", func(t *arm64.Section) { t.Stxrh(w2, w0, arm64.Mem16(x1)) }},
		{"stlxr w2,w0,[x1]", "stlxr w2, w0, [x1]", func(t *arm64.Section) { t.Stlxr32(w2, w0, arm64.Mem32(x1)) }},
		{"stlxr w2,x0,[x1]", "stlxr w2, x0, [x1]", func(t *arm64.Section) { t.Stlxr64(w2, x0, arm64.Mem64(x1)) }},
		{"stlxrb w2,w0,[x1]", "stlxrb w2, w0, [x1]", func(t *arm64.Section) { t.Stlxrb(w2, w0, arm64.Mem8(x1)) }},
		{"stlxrh w2,w0,[x1]", "stlxrh w2, w0, [x1]", func(t *arm64.Section) { t.Stlxrh(w2, w0, arm64.Mem16(x1)) }},

		{"ldaddal w0,w2,[x1]", "ldaddal w0, w2, [x1]", func(t *arm64.Section) { t.Ldaddal32(w0, w2, arm64.Mem32(x1)) }},
		{"ldaddal x0,x2,[x1]", "ldaddal x0, x2, [x1]", func(t *arm64.Section) { t.Ldaddal64(x0, x2, arm64.Mem64(x1)) }},
		{"ldaddalb w0,w2,[x1]", "ldaddalb w0, w2, [x1]", func(t *arm64.Section) { t.Ldaddalb(w0, w2, arm64.Mem8(x1)) }},
		{"ldaddalh w0,w2,[x1]", "ldaddalh w0, w2, [x1]", func(t *arm64.Section) { t.Ldaddalh(w0, w2, arm64.Mem16(x1)) }},
		{"ldclral w0,w2,[x1]", "ldclral w0, w2, [x1]", func(t *arm64.Section) { t.Ldclral32(w0, w2, arm64.Mem32(x1)) }},
		{"ldclral x0,x2,[x1]", "ldclral x0, x2, [x1]", func(t *arm64.Section) { t.Ldclral64(x0, x2, arm64.Mem64(x1)) }},
		{"ldeoral w0,w2,[x1]", "ldeoral w0, w2, [x1]", func(t *arm64.Section) { t.Ldeoral32(w0, w2, arm64.Mem32(x1)) }},
		{"ldeoral x0,x2,[x1]", "ldeoral x0, x2, [x1]", func(t *arm64.Section) { t.Ldeoral64(x0, x2, arm64.Mem64(x1)) }},
		{"ldsetal w0,w2,[x1]", "ldsetal w0, w2, [x1]", func(t *arm64.Section) { t.Ldsetal32(w0, w2, arm64.Mem32(x1)) }},
		{"ldsetal x0,x2,[x1]", "ldsetal x0, x2, [x1]", func(t *arm64.Section) { t.Ldsetal64(x0, x2, arm64.Mem64(x1)) }},
		{"swpal w0,w2,[x1]", "swpal w0, w2, [x1]", func(t *arm64.Section) { t.Swpal32(w0, w2, arm64.Mem32(x1)) }},
		{"swpal x0,x2,[x1]", "swpal x0, x2, [x1]", func(t *arm64.Section) { t.Swpal64(x0, x2, arm64.Mem64(x1)) }},
		{"casal w0,w2,[x1]", "casal w0, w2, [x1]", func(t *arm64.Section) { t.Casal32(w0, w2, arm64.Mem32(x1)) }},
		{"casal x0,x2,[x1]", "casal x0, x2, [x1]", func(t *arm64.Section) { t.Casal64(x0, x2, arm64.Mem64(x1)) }},
		{"casalb w0,w2,[x1]", "casalb w0, w2, [x1]", func(t *arm64.Section) { t.Casalb(w0, w2, arm64.Mem8(x1)) }},
		{"casalh w0,w2,[x1]", "casalh w0, w2, [x1]", func(t *arm64.Section) { t.Casalh(w0, w2, arm64.Mem16(x1)) }},

		{"mul w0,w1,w2", "mul w0, w1, w2", func(t *arm64.Section) { t.Mul32(w0, w1, w2) }},
		{"mov w0,#100", "mov w0, #100", func(t *arm64.Section) { t.MovWide32(w0, 100) }},
		{"orn w0,w1,w2", "orn w0, w1, w2", func(t *arm64.Section) { t.OrnShifted32(w0, w1, w2) }},
		{"orn x0,x1,x2", "orn x0, x1, x2", func(t *arm64.Section) { t.OrnShifted64(x0, x1, x2) }},
		{"ldrsb w0,[x1,4]", "ldrsb w0, [x1, #4]", func(t *arm64.Section) { t.LdrsbImm32(w0, arm64.Mem8(x1).Off(4)) }},
		{"ldrsb x0,[x1,4]", "ldrsb x0, [x1, #4]", func(t *arm64.Section) { t.LdrsbImm64(x0, arm64.Mem8(x1).Off(4)) }},
		{"ldrsh w0,[x1,4]", "ldrsh w0, [x1, #4]", func(t *arm64.Section) { t.LdrshImm32(w0, arm64.Mem16(x1).Off(4)) }},
		{"ldrsh x0,[x1,4]", "ldrsh x0, [x1, #4]", func(t *arm64.Section) { t.LdrshImm64(x0, arm64.Mem16(x1).Off(4)) }},
	}
}

func TestDifferentialAgainstClang(t *testing.T) {
	clangPath, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not on PATH; skipping differential test against the reference assembler")
	}

	cases := diffCases()

	// Assemble every case's line as one source file, in order, so the .text
	// offsets line up one-to-one with the typed-helper words emitted below.
	// The LSE atomics are Armv8.1-A, so the reference assembler is told
	// which architecture it is assembling for rather than being handed
	// instructions it will refuse. The typed-helper side is told the same
	// thing through the module's feature set, below.
	var src strings.Builder
	src.WriteString(".arch armv8.1-a\n")
	for _, c := range cases {
		src.WriteString(c.asm)
		src.WriteByte('\n')
	}

	dir := t.TempDir()
	srcPath := filepath.Join(dir, "cases.s")
	objPath := filepath.Join(dir, "cases.o")
	if err := os.WriteFile(srcPath, []byte(src.String()), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(clangPath, "-target", "aarch64-linux-gnu", "-c", "-o", objPath, srcPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("clang: %v\n%s", err, stderr.String())
	}

	want, err := textWords(objPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(want) != len(cases) {
		t.Fatalf("clang produced %d words for %d cases; a source line assembled to something other than one instruction", len(want), len(cases))
	}

	m := arm64.NewModule(arm64.WithFeatures(arm64.Armv8_1A.Set()))
	m.Extern("memcpy")
	sec := m.Section(arm64.Text)
	for _, c := range cases {
		c.build(sec)
	}
	if err := m.Err(); err != nil {
		t.Fatalf("building the typed-helper sequence: %v", err)
	}
	obj, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	got := obj.SectionAt(0).Bytes()
	if len(got) != len(cases)*4 {
		t.Fatalf("got %d bytes for %d cases, want %d", len(got), len(cases), len(cases)*4)
	}

	for i, c := range cases {
		g := binary.LittleEndian.Uint32(got[i*4:])
		w := want[i]
		if g != w {
			t.Errorf("%-24s got %#08x, want %#08x (clang)", c.name, g, w)
		}
	}
}

// textWords disassembles the ELF object's .text section with objdump and
// returns each instruction word in address order.
//
// Disassembly rather than a raw section dump because a section dump would
// still need this same address bookkeeping to slice it into words, and
// disassembly output states the address per line for free, which is what
// catches a source line that silently assembled to more or fewer than one
// instruction (a macro, a `.word` directive after it) before that
// desynchronizes every case that follows it.
func textWords(objPath string) ([]uint32, error) {
	out, err := exec.Command("objdump", "-d", "-j", ".text", objPath).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("objdump: %w\n%s", err, out)
	}
	line := regexp.MustCompile(`^\s*[0-9a-f]+:\s+([0-9a-f]{8})\b`)
	var words []uint32
	for _, l := range strings.Split(string(out), "\n") {
		m := line.FindStringSubmatch(l)
		if m == nil {
			continue
		}
		v, err := strconv.ParseUint(m[1], 16, 32)
		if err != nil {
			return nil, fmt.Errorf("objdump: unparsable word %q: %w", m[1], err)
		}
		words = append(words, uint32(v))
	}
	return words, nil
}
