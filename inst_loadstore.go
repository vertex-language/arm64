package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Load and store, unsigned scaled offset --------------------------------
//
// m is built with Mem8/16/32/64/128(base).Off(n) or .Off(PageOff(sym)): the
// scaled forms here take [Xn|SP, #imm], imm in units of the access width.
// Handing one of these a pre- or post-indexed Mem is refused at the call —
// StpPre64 and LdpPost64 are the separate, writeback-specific forms.

var (
	strbImm  = form("StrbImm")
	ldrbImm  = form("LdrbImm")
	strhImm  = form("StrhImm")
	ldrhImm  = form("LdrhImm")
	strImm32 = form("StrImm32")
	ldrImm32 = form("LdrImm32")
	strImm64 = form("StrImm64")
	ldrImm64 = form("LdrImm64")
	ldrswImm = form("LdrswImm")
)

// StrbImm emits STRB Wt, [Xn|SP{, #imm}].
func (s *Section) StrbImm(rt reg.W, m Mem) { s.inst(strbImm, rt, m) }

// LdrbImm emits LDRB Wt, [Xn|SP{, #imm}]: an eight-bit load, zero-extended.
func (s *Section) LdrbImm(rt reg.W, m Mem) { s.inst(ldrbImm, rt, m) }

// StrhImm emits STRH Wt, [Xn|SP{, #imm}].
func (s *Section) StrhImm(rt reg.W, m Mem) { s.inst(strhImm, rt, m) }

// LdrhImm emits LDRH Wt, [Xn|SP{, #imm}]: a sixteen-bit load, zero-extended.
func (s *Section) LdrhImm(rt reg.W, m Mem) { s.inst(ldrhImm, rt, m) }

// StrImm32 emits STR Wt, [Xn|SP{, #imm}].
func (s *Section) StrImm32(rt reg.W, m Mem) { s.inst(strImm32, rt, m) }
func (s *Section) LdrImm32(rt reg.W, m Mem) { s.inst(ldrImm32, rt, m) }

// StrImm64 emits STR Xt, [Xn|SP{, #imm}].
func (s *Section) StrImm64(rt reg.X, m Mem) { s.inst(strImm64, rt, m) }
func (s *Section) LdrImm64(rt reg.X, m Mem) { s.inst(ldrImm64, rt, m) }

// LdrswImm emits LDRSW Xt, [Xn|SP{, #imm}]: a thirty-two-bit load,
// sign-extended into Xt.
func (s *Section) LdrswImm(rt reg.X, m Mem) { s.inst(ldrswImm, rt, m) }

// ---- Load and store, unscaled -----------------------------------------------
//
// The STUR/LDUR family: the same [Xn|SP, #imm] shape with a signed
// nine-bit byte offset rather than a scaled twelve-bit one, for the address
// a scaled form's alignment requirement refuses.

var (
	sturImm32 = form("SturImm32")
	ldurImm32 = form("LdurImm32")
	sturImm64 = form("SturImm64")
	ldurImm64 = form("LdurImm64")
)

func (s *Section) SturImm32(rt reg.W, m Mem) { s.inst(sturImm32, rt, m) }
func (s *Section) LdurImm32(rt reg.W, m Mem) { s.inst(ldurImm32, rt, m) }
func (s *Section) SturImm64(rt reg.X, m Mem) { s.inst(sturImm64, rt, m) }
func (s *Section) LdurImm64(rt reg.X, m Mem) { s.inst(ldurImm64, rt, m) }

// ---- Load and store pair -----------------------------------------------------

var (
	stp32     = form("Stp32")
	ldp32     = form("Ldp32")
	stp64     = form("Stp64")
	ldp64     = form("Ldp64")
	stpPre64  = form("StpPre64")
	ldpPost64 = form("LdpPost64")
)

// Stp32 emits STP Wt, Wt2, [Xn|SP{, #imm}].
func (s *Section) Stp32(rt, rt2 reg.W, m Mem) { s.inst(stp32, rt, rt2, m) }
func (s *Section) Ldp32(rt, rt2 reg.W, m Mem) { s.inst(ldp32, rt, rt2, m) }

// Stp64 emits STP Xt, Xt2, [Xn|SP{, #imm}].
func (s *Section) Stp64(rt, rt2 reg.X, m Mem) { s.inst(stp64, rt, rt2, m) }
func (s *Section) Ldp64(rt, rt2 reg.X, m Mem) { s.inst(ldp64, rt, rt2, m) }

// StpPre64 emits STP Xt, Xt2, [Xn|SP, #imm]! — pre-indexed: the address
// computes first and writes back into the base. m must be built with .Pre.
func (s *Section) StpPre64(rt, rt2 reg.X, m Mem) { s.inst(stpPre64, rt, rt2, m) }

// LdpPost64 emits LDP Xt, Xt2, [Xn|SP], #imm — post-indexed: the base is
// used as written, then written back. m must be built with .Post.
func (s *Section) LdpPost64(rt, rt2 reg.X, m Mem) { s.inst(ldpPost64, rt, rt2, m) }

// ---- Load literal -------------------------------------------------------------

var (
	ldrLit32 = form("LdrLit32")
	ldrLit64 = form("LdrLit64")
)

// LdrLit32 emits LDR Wt, target: a PC-relative literal load, +/-1MiB. target
// must be a same-section Label — the field has no relocation in any of the
// three container formats a literal pool can cross a section boundary with.
func (s *Section) LdrLit32(rt reg.W, target Label) { s.inst(ldrLit32, rt, target) }

// LdrLit64 emits LDR Xt, target.
func (s *Section) LdrLit64(rt reg.X, target Label) { s.inst(ldrLit64, rt, target) }

// ---- Sign-extending sub-width loads ----------------------------------------
//
// One instruction rather than a load and a widen: the sign extension is the
// mnemonic, and the destination's width says how far.

var (
	ldrsbImm32 = form("LdrsbImm32")
	ldrsbImm64 = form("LdrsbImm64")
	ldrshImm32 = form("LdrshImm32")
	ldrshImm64 = form("LdrshImm64")
)

// LdrsbImm32 emits LDRSB Wt, [Xn|SP{, #imm}]: a byte, sign-extended to 32 bits.
func (s *Section) LdrsbImm32(rt reg.W, m Mem) { s.inst(ldrsbImm32, rt, m) }

// LdrsbImm64 emits LDRSB Xt, [Xn|SP{, #imm}], sign-extended to 64.
func (s *Section) LdrsbImm64(rt reg.X, m Mem) { s.inst(ldrsbImm64, rt, m) }

// LdrshImm32 emits LDRSH Wt, [Xn|SP{, #imm}].
func (s *Section) LdrshImm32(rt reg.W, m Mem) { s.inst(ldrshImm32, rt, m) }

// LdrshImm64 emits LDRSH Xt, [Xn|SP{, #imm}].
func (s *Section) LdrshImm64(rt reg.X, m Mem) { s.inst(ldrshImm64, rt, m) }
