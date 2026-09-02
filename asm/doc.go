// Package asm assembles AArch64 source text.
//
// It is the second front door onto the encoder in the parent package. The
// first is the typed surface — Section.AddImm64 and its four hundred
// siblings — where the instruction is known when the Go is written and a
// width mismatch is a compile error. This one is for when the instruction is
// text: a .s file, or the template of an inline asm statement.
//
// Both doors reach the same table and the same encoder, which is the property
// worth having. An assembler that had its own encoding tables would be a
// second answer to "what are the bytes of this instruction", and a second
// answer is a thing that can disagree.
//
// The split with the gas package is that gas knows statements and this knows
// instructions. Everything up to and including a mnemonic — labels,
// directives, expressions, `1f`, `.pushsection` — is gas's, and none of it is
// specific to this architecture. Everything after a mnemonic is here, and none
// of it is shared with any other.
//
// # Syntax
//
// UAL, as GNU as spells it: `//` comments, `#` immediates, `[x0, #8]!`
// addressing, `:lo12:` relocation modifiers. There is no second dialect to
// select, because AArch64 does not have one — armasm's differences are all
// directives, which is to say gas's business rather than this package's.
package asm
