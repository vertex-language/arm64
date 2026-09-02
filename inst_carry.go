package arm64

import "github.com/vertex-language/arm64/reg"

// Add and subtract with carry, the widening multiply, and the aliases built
// on both.
//
// None of this is selected by anything in this tree, which is why it was
// missing until an assembler asked: this IR's integers fit a register, so a
// multiword addition never appears, and its overflow predicates are computed
// rather than carried in a flag.

var (
	adc32  = form("Adc32")
	adc64  = form("Adc64")
	adcs32 = form("Adcs32")
	adcs64 = form("Adcs64")
	sbc32  = form("Sbc32")
	sbc64  = form("Sbc64")
	sbcs32 = form("Sbcs32")
	sbcs64 = form("Sbcs64")
)

// Adc64 emits ADC Xd, Xn, Xm: Xn + Xm + C. Adcs sets the flags, which is what
// carries into the next word of a multiword addition.
func (s *Section) Adc32(rd, rn, rm reg.W)  { s.inst(adc32, rd, rn, rm) }
func (s *Section) Adc64(rd, rn, rm reg.X)  { s.inst(adc64, rd, rn, rm) }
func (s *Section) Adcs32(rd, rn, rm reg.W) { s.inst(adcs32, rd, rn, rm) }
func (s *Section) Adcs64(rd, rn, rm reg.X) { s.inst(adcs64, rd, rn, rm) }

// Sbc64 emits SBC Xd, Xn, Xm: Xn - Xm - (1 - C), which is the borrow written
// the way the architecture spells it.
func (s *Section) Sbc32(rd, rn, rm reg.W)  { s.inst(sbc32, rd, rn, rm) }
func (s *Section) Sbc64(rd, rn, rm reg.X)  { s.inst(sbc64, rd, rn, rm) }
func (s *Section) Sbcs32(rd, rn, rm reg.W) { s.inst(sbcs32, rd, rn, rm) }
func (s *Section) Sbcs64(rd, rn, rm reg.X) { s.inst(sbcs64, rd, rn, rm) }

var (
	negs32 = form("Negs32")
	negs64 = form("Negs64")
	ngc32  = form("Ngc32")
	ngc64  = form("Ngc64")
	ngcs32 = form("Ngcs32")
	ngcs64 = form("Ngcs64")
)

// Negs64 emits NEGS Xd, Xm{, shift} — SUBS with the zero register as its
// first source, which is NEG's relation to SUB.
func (s *Section) Negs32(rd, rm reg.W, shift ...ShiftOp) {
	s.inst(negs32, append([]any{rd, rm}, opt(shift)...)...)
}

func (s *Section) Negs64(rd, rm reg.X, shift ...ShiftOp) {
	s.inst(negs64, append([]any{rd, rm}, opt(shift)...)...)
}

// Ngc64 emits NGC Xd, Xm — SBC with the zero register, which negates with a
// borrow. Ngcs sets the flags.
func (s *Section) Ngc32(rd, rm reg.W)  { s.inst(ngc32, rd, rm) }
func (s *Section) Ngc64(rd, rm reg.X)  { s.inst(ngc64, rd, rm) }
func (s *Section) Ngcs32(rd, rm reg.W) { s.inst(ngcs32, rd, rm) }
func (s *Section) Ngcs64(rd, rm reg.X) { s.inst(ngcs64, rd, rm) }

var (
	smaddl = form("Smaddl")
	umaddl = form("Umaddl")
	smsubl = form("Smsubl")
	umsubl = form("Umsubl")
	smull  = form("Smull")
	umull  = form("Umull")
	smnegl = form("Smnegl")
	umnegl = form("Umnegl")
)

// Smaddl emits SMADDL Xd, Wn, Wm, Xa: Xa + Wn*Wm, signed. The widening
// multiply is the one shape this architecture cannot express by choosing
// register widths, since its sources and its destination differ.
func (s *Section) Smaddl(rd reg.X, rn, rm reg.W, ra reg.X) { s.inst(smaddl, rd, rn, rm, ra) }
func (s *Section) Umaddl(rd reg.X, rn, rm reg.W, ra reg.X) { s.inst(umaddl, rd, rn, rm, ra) }
func (s *Section) Smsubl(rd reg.X, rn, rm reg.W, ra reg.X) { s.inst(smsubl, rd, rn, rm, ra) }
func (s *Section) Umsubl(rd reg.X, rn, rm reg.W, ra reg.X) { s.inst(umsubl, rd, rn, rm, ra) }

// Smull emits SMULL Xd, Wn, Wm — SMADDL with the accumulator pinned to the
// zero register.
func (s *Section) Smull(rd reg.X, rn, rm reg.W)  { s.inst(smull, rd, rn, rm) }
func (s *Section) Umull(rd reg.X, rn, rm reg.W)  { s.inst(umull, rd, rn, rm) }
func (s *Section) Smnegl(rd reg.X, rn, rm reg.W) { s.inst(smnegl, rd, rn, rm) }
func (s *Section) Umnegl(rd reg.X, rn, rm reg.W) { s.inst(umnegl, rd, rn, rm) }
