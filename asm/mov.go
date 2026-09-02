package asm

import (
	"github.com/vertex-language/arm64/operand"
	"github.com/vertex-language/arm64/reg"
)

// movPseudo rewrites `mov Rd, #imm` the way GNU as does.
//
// `mov` with an immediate is not one encoding. GNU as picks among three, in
// this order: MOVZ when the value is one halfword of the register, MOVN when
// the inverse of the value is, and ORR against the zero register when the
// value is a bitmask immediate. Only the first is in the table, and belongs
// there, because it is also the preferred *disassembly* of that MOVZ — the
// alias relation runs both ways. The other two are not aliases at all: a
// disassembler prints them as movn and orr, and giving them `mov` rows would
// claim a preference the architecture does not state.
//
// So the selection lives here, at the one layer whose job is to accept what
// GNU as accepts. `mov x0, #-1` is the case that matters — it is how every
// hand-written .s writes an all-ones mask, and it is MOVN with a zero field.
func movPseudo(name string, ops []any) (string, []any) {
	if name != "mov" || len(ops) != 2 {
		return name, ops
	}
	imm, ok := ops[1].(operand.Imm)
	if !ok {
		return name, ops
	}

	var w operand.Width
	var zero any
	switch ops[0].(type) {
	case reg.X:
		w, zero = operand.Width64, reg.XZR
	case reg.W:
		w, zero = operand.Width32, reg.WZR
	default:
		return name, ops
	}

	v := truncate(uint64(imm), w)
	if _, hw, ok := operand.FitsImm16Shifted(v, w); ok {
		if hw == 0 {
			return name, ops // the MOV alias row, which holds a low halfword
		}
		// A higher halfword needs the shift written, and MOV is the one
		// spelling that cannot say it: GNU as refuses `mov x0, #1, lsl #16`
		// for exactly that reason. MOVZ is where the field lives.
		return "movz", ops
	}
	inv := truncate(^v, w)
	if _, _, ok := operand.FitsImm16Shifted(inv, w); ok {
		// MOVN's field is the value to be inverted, so the operand is the
		// inverse and the instruction produces what was asked for.
		return "movn", []any{ops[0], operand.Imm(inv)}
	}
	if _, _, _, ok := operand.EncodeBitmask(v, w); ok {
		return "orr", []any{ops[0], zero, operand.Imm(v)}
	}

	// Out of reach of all three. Leave the mnemonic alone: the table's
	// diagnostic names the halfword rule, which is the first thing to check
	// and the reason a movz/movk pair is what a wider value needs.
	return name, ops
}

func truncate(v uint64, w operand.Width) uint64 {
	if w == operand.Width32 {
		return v & 0xffffffff
	}
	return v
}
