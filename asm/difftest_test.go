package asm_test

// Differential testing against a reference assembler.
//
// Every line below is assembled twice: once by this package, and once by
// clang's integrated AArch64 assembler. The two must produce identical
// instruction words.
//
// This is a different claim from the parent package's difftest, which checks
// that the ISA table's encodings are right. Here the encodings are taken as
// given and what is on trial is the parse: that `[sp, #-16]!` reaches the
// encoder as a pre-indexed 64-bit address and not as something else that
// happens to encode. Those are the bugs a parser has, and no amount of
// self-consistency checking finds them — only a second assembler does.
//
// Skipped rather than failed when clang is absent: a missing reference
// assembler is a fact about the machine, not a defect here.

import (
	"bytes"
	"encoding/binary"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/vertex-language/arm64/asm"
)

// lines is one instruction per entry, each assembling to exactly one word.
var lines = []string{
	// Arithmetic, immediate and shifted register.
	"add x0, x1, x2",
	"add w0, w1, w2, lsl #3",
	"add x0, x1, #16",
	"add x0, x1, #16, lsl #12",
	"sub sp, sp, #32",
	"add x0, sp, #8",
	"adds w3, w4, w5",
	"subs x3, x4, x5",
	"sub x9, x10, x11, asr #7",
	"neg x0, x1",
	"neg w0, w1",
	"cmp x0, x1",
	"cmp w0, w1",
	"cmp w0, w1, lsl #2",
	"cmp w0, #4",
	"cmn x2, x3",
	"cmn w2, w3",
	"cmn x0, #4",
	"cmn w0, #4",

	// Logical.
	"and x6, x7, #0xff",
	"orr w1, w2, w3",
	"eor x4, x5, x6, ror #12",
	"and x0, x1, x2, lsl #1",
	"tst x5, x6",
	"tst w5, w6",
	"tst w0, #0xff",
	"mvn x3, x4",
	"mvn w3, w4",
	"mvn x3, x4, lsl #2",

	// Moves.
	"mov x0, x1",
	"mov w0, w1",
	"mov x0, #1",
	"movz x0, #0x1234",
	"movz x0, #0x1234, lsl #16",
	"movk w1, #0xbeef",
	"movn x2, #0",
	"mov x0, sp",
	"mov sp, x0",

	// Shifts and bitfield.
	"lsl x0, x1, #4",
	"lsl w0, w1, #3",
	"lsr w2, w3, #7",
	"asr w4, w5, #9",
	"lsl x0, x1, x2",
	"lsr w0, w1, w2",
	"asr x4, x5, x6",
	"ror w0, w1, w2",
	"sxtb w0, w1",
	"sxtb x0, w1",
	"sxth w0, w1",
	"sxth x0, w1",
	"sxtw x0, w1",
	"uxtb w2, w3",
	"uxth w2, w3",
	"ubfx x0, x1, #4, #8",
	"ubfx w0, w1, #2, #3",
	"sbfx w2, w3, #1, #5",
	"sbfx x0, x1, #4, #8",
	"ror w7, w8, #3",
	"ror x0, x1, #5",

	// Add and subtract with carry, and the negate aliases built on them.
	"adc x0, x1, x2",
	"adc w0, w1, w2",
	"adcs x0, x1, x2",
	"sbc x0, x1, x2",
	"sbc w0, w1, w2",
	"sbcs x0, x1, x2",
	"ngc x0, x1",
	"ngcs x0, x1",
	"negs x0, x1",
	"negs w0, w1",

	// Widening multiply: W sources, an X destination.
	"smull x0, w1, w2",
	"umull x0, w1, w2",
	"smaddl x0, w1, w2, x3",
	"umaddl x0, w1, w2, x3",
	"smsubl x0, w1, w2, x3",
	"umsubl x0, w1, w2, x3",
	"smnegl x0, w1, w2",

	// The conditional aliases, which name one source twice and invert
	// their condition.
	"cinc x0, x1, ne",
	"cinv x0, x1, ne",
	"cneg x0, x1, ne",
	"csetm x0, ne",
	"csetm w0, ne",
	"cset w0, ne",

	// Bitfield insert, the other direction from extract.
	"bfi x0, x1, #4, #8",
	"bfi w0, w1, #4, #8",
	"bfxil x0, x1, #4, #8",
	"ubfiz x0, x1, #4, #8",
	"sbfiz w0, w1, #4, #8",

	// The hint space, and the exception return.
	"hint #0x14",
	"paciasp",
	"autiasp",
	"eret",
	"rev64 x0, x1",
	"rev32 x0, x1",

	// Multiply and divide.
	"mul x0, x1, x2",
	"madd x0, x1, x2, x3",
	"msub w0, w1, w2, w3",
	"smulh x0, x1, x2",
	"umulh x0, x1, x2",
	"sdiv x3, x4, x5",
	"udiv w6, w7, w8",

	// Counts and reverses.
	"clz x0, x1",
	"rbit w2, w3",
	"rev x4, x5",
	"rev16 w6, w7",

	// Conditionals.
	"csel x0, x1, x2, eq",
	"csinc w3, w4, w5, ne",
	"cset x6, lt",
	"csneg x7, x8, x9, ge",
	"ccmp x0, x1, #0, eq",
	"ccmp w0, w1, #3, ne",
	"ccmp x0, #5, #0, eq",
	"ccmp w0, #5, #7, lt",
	"ccmn x0, x1, #0, eq",
	"ccmn w0, #5, #7, lt",

	// Addressing: the whole point of this file.
	"ldr x0, [x1]",
	"ldr x0, [x1, #8]",
	"ldr w0, [x1, #4]",
	"str x2, [sp, #16]",
	"strb w3, [x4, #1]",
	"ldrh w5, [x6, #2]",
	"ldrsw x7, [x8, #4]",
	"ldur x0, [x1, #-8]",
	"stur w2, [x3, #-4]",
	"stp x29, x30, [sp, #-16]!",
	"ldp x29, x30, [sp], #16",
	"stp w0, w1, [x2, #8]",
	"ldp x3, x4, [x5, #16]",
	"ldr x0, [sp]",
	"prfm pldl1keep, [x0]",
	"prfm pstl2strm, [x1, #16]",

	// Branches.
	"b 1f",
	"bl 1f",
	"br x0",
	"blr x1",
	"ret",
	"ret x30",
	"b.eq 1f",
	"b.ne 1f",
	"b.lt 1f",
	"cbz x0, 1f",
	"cbnz w1, 1f",
	"tbz x2, #3, 1f",
	"tbnz w4, #5, 1f",
	"tbz w4, #5, 1f",

	// Floating point.
	"fadd s0, s1, s2",
	"fsub d3, d4, d5",
	"fmul s6, s7, s8",
	"fdiv d0, d1, d2",
	"fmov s0, s1",
	"fmov x0, d1",
	"fcmp d2, d3",
	"fcvtzs w0, s1",
	"scvtf d0, x1",
	"fneg s4, s5",
	"fabs d6, d7",
	"fsqrt s8, s9",

	// Atomics and barriers.
	"ldxr x0, [x1]",
	"stxr w2, x3, [x4]",
	"ldaxr w5, [x6]",
	"stlxr w7, w8, [x9]",
	"dmb ish",
	"dsb sy",
	"isb",
	"clrex",

	// System and misc.
	"nop",
	"brk #0",
	"svc #0",
	"mrs x0, tpidr_el0",
	"msr tpidr_el0, x1",
	"adr x0, 1f",
}

// tableGaps are lines this package parses correctly and the parent package's
// ISA table cannot yet encode: mnemonics it declares no row for, and widths it
// declares only one of.
//
// The list is empty, and it is kept because it is how it got that way. It held
// twelve entries — the 32-bit halves of aliases declared only at 64, the
// bitfield and extend aliases, the register forms of the shifts, conditional
// compare, and a prefetch — each discovered by pointing a parser at ordinary
// assembly rather than by reading the ARM ARM front to back. The test below
// fails when an entry starts working, which is what kept the list honest as
// the rows landed; it will do the same for whatever the next parser finds.
var tableGaps = []string{}

// TestTableGapsStillGap keeps the list above honest: an entry that starts
// working is one to move up into lines, and a silent list would never say so.
func TestTableGapsStillGap(t *testing.T) {
	for _, l := range tableGaps {
		if _, err := asm.Assemble(l+"\n1:\n", asm.Options{File: "gap.s"}); err == nil {
			t.Errorf("%q now assembles; move it into lines", l)
		}
	}
}

func TestDifferentialAgainstClang(t *testing.T) {
	clangPath, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not on PATH; skipping differential test against the reference assembler")
	}

	// One source, both assemblers, so a same-section branch to `1:` resolves
	// identically on the two sides. The label goes last, after every line
	// that refers to it.
	body := strings.Join(lines, "\n") + "\n1:\n"

	dir := t.TempDir()
	srcPath := filepath.Join(dir, "cases.s")
	objPath := filepath.Join(dir, "cases.o")
	if err := os.WriteFile(srcPath, []byte(body), 0o644); err != nil {
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
	if len(want) != len(lines) {
		t.Fatalf("clang produced %d words for %d lines; a line assembled to other than one instruction",
			len(want), len(lines))
	}

	o, err := asm.Assemble(body, asm.Options{File: "cases.s"})
	if err != nil {
		t.Fatalf("assemble: %v", err)
	}
	got := o.SectionAt(0).Bytes()
	if len(got) != len(lines)*4 {
		t.Fatalf("got %d bytes for %d lines, want %d", len(got), len(lines), len(lines)*4)
	}

	bad := 0
	for i, line := range lines {
		g := binary.LittleEndian.Uint32(got[i*4:])
		if g != want[i] {
			t.Errorf("%-32s got %#08x, want %#08x (clang)", line, g, want[i])
			bad++
		}
	}
	if bad == 0 {
		t.Logf("%d lines agree with clang", len(lines))
	}
}

// textWords disassembles .text and returns each instruction word in order.
func textWords(objPath string) ([]uint32, error) {
	out, err := exec.Command("objdump", "-d", "-j", ".text", objPath).CombinedOutput()
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`(?m)^\s+[0-9a-f]+:\s+([0-9a-f]{8})\s`)
	var words []uint32
	for _, m := range re.FindAllStringSubmatch(string(out), -1) {
		v, err := strconv.ParseUint(m[1], 16, 32)
		if err != nil {
			return nil, err
		}
		words = append(words, uint32(v))
	}
	return words, nil
}
