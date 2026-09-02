package asm

import "testing"

func TestSmoke(t *testing.T) {
	src := `
	.text
	.globl main
	.type main, @function
main:
	stp x29, x30, [sp, #-16]!
	mov x29, sp
	add x0, x0, #1
	adrp x1, msg
	add x1, x1, :lo12:msg
	bl puts
	mov w0, wzr
	ldp x29, x30, [sp], #16
	ret
	.size main, . - main

	.section .rodata
msg:
	.asciz "hello"
`
	o, err := Assemble(src, Options{File: "smoke.s"})
	if err != nil {
		t.Fatalf("assemble: %v", err)
	}
	for _, s := range o.Sections() {
		t.Logf("section %s (%s): % x", s.Name(), s.Kind(), s.Bytes())
	}
	for _, sy := range o.Symbols() {
		t.Logf("symbol %+v", sy)
	}
}
