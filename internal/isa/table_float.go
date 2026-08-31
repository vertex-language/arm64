package isa

// The scalar floating-point instruction set: FP arithmetic, the conversions
// between the two register files, the comparisons and the FP loads and stores.
//
// Single and double precision only. Half precision is a separate ftype and a
// separate feature, and the vector arrangements are a different operand shape
// entirely; neither is here because nothing yet asks for them.
//
// FP and SIMD are in the Armv8-A baseline, so no row here is gated.
func init() {
	register(
		// ---- Arithmetic, two source ----
		L("fadd", 0x1e202800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FaddS"),
		L("fadd", 0x1e602800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FaddD"),
		L("fsub", 0x1e203800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FsubS"),
		L("fsub", 0x1e603800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FsubD"),
		L("fmul", 0x1e200800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FmulS"),
		L("fmul", 0x1e600800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FmulD"),
		L("fdiv", 0x1e201800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FdivS"),
		L("fdiv", 0x1e601800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FdivD"),

		// FMAX and FMIN propagate a NaN operand; FMAXNM and FMINNM return the
		// other one. Both pairs exist in the silicon here, which is not true
		// everywhere: they are IEEE-754-2019 maximum/minimum and 2008
		// maxNum/minNum, and C's fmax and fmin want the second.
		L("fmax", 0x1e204800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FmaxS"),
		L("fmax", 0x1e604800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FmaxD"),
		L("fmin", 0x1e205800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FminS"),
		L("fmin", 0x1e605800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FminD"),
		L("fmaxnm", 0x1e206800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FmaxnmS"),
		L("fmaxnm", 0x1e606800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FmaxnmD"),
		L("fminnm", 0x1e207800, 0xffe0fc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Name("FminnmS"),
		L("fminnm", 0x1e607800, 0xffe0fc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Name("FminnmD"),

		// ---- Arithmetic, one source ----
		L("fmov", 0x1e204000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FmovS"),
		L("fmov", 0x1e604000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FmovD"),
		L("fabs", 0x1e20c000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FabsS"),
		L("fabs", 0x1e60c000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FabsD"),
		L("fneg", 0x1e214000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FnegS"),
		L("fneg", 0x1e614000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FnegD"),
		L("fsqrt", 0x1e21c000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FsqrtS"),
		L("fsqrt", 0x1e61c000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FsqrtD"),

		// The four rounding modes that are their own instruction rather than a
		// read of FPCR: nearest-even, +inf, -inf and zero.
		L("frintn", 0x1e244000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FrintnS"),
		L("frintn", 0x1e644000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FrintnD"),
		L("frintp", 0x1e24c000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FrintpS"),
		L("frintp", 0x1e64c000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FrintpD"),
		L("frintm", 0x1e254000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FrintmS"),
		L("frintm", 0x1e654000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FrintmD"),
		L("frintz", 0x1e25c000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Name("FrintzS"),
		L("frintz", 0x1e65c000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Name("FrintzD"),

		// ---- Width conversion, float to float ----
		L("fcvt", 0x1e22c000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassS, Rn).Name("FcvtSToD"),
		L("fcvt", 0x1e624000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassD, Rn).Name("FcvtDToS"),

		// ---- Arithmetic, three source ----
		//
		// One rounding for the whole a*b+c, which is what makes these the
		// fused multiply-add rather than a multiply and an add.
		L("fmadd", 0x1f000000, 0xffe08000).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Src(ClassS, Ra).Name("FmaddS"),
		L("fmadd", 0x1f400000, 0xffe08000).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Src(ClassD, Ra).Name("FmaddD"),
		L("fmsub", 0x1f008000, 0xffe08000).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Src(ClassS, Ra).Name("FmsubS"),
		L("fmsub", 0x1f408000, 0xffe08000).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Src(ClassD, Ra).Name("FmsubD"),

		// ---- Compare ----
		//
		// FCMP writes NZCV and has no destination. An unordered pair sets C
		// and V, which is what makes the unsigned conditions the ones that
		// answer a float comparison: MI is less-than and it is false for a
		// NaN, where LT would be true.
		L("fcmp", 0x1e202000, 0xffe0fc1f).
			Src(ClassS, Rn).Src(ClassS, Rm).Flags().Name("FcmpS"),
		L("fcmp", 0x1e602000, 0xffe0fc1f).
			Src(ClassD, Rn).Src(ClassD, Rm).Flags().Name("FcmpD"),
		L("fcmp", 0x1e202008, 0xffe0fc1f).
			Src(ClassS, Rn).Flags().Name("FcmpZeroS"),
		L("fcmp", 0x1e602008, 0xffe0fc1f).
			Src(ClassD, Rn).Flags().Name("FcmpZeroD"),

		// ---- Conditional select ----
		L("fcsel", 0x1e200c00, 0xffe00c00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Src(ClassS, Rm).Cnd(CondHi).Name("FcselS"),
		L("fcsel", 0x1e600c00, 0xffe00c00).
			Dst(ClassD, Rd).Src(ClassD, Rn).Src(ClassD, Rm).Cnd(CondHi).Name("FcselD"),

		// ---- Integer to float ----
		L("scvtf", 0x1e220000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassW, Rn).Name("ScvtfWToS"),
		L("scvtf", 0x1e620000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassW, Rn).Name("ScvtfWToD"),
		L("scvtf", 0x9e220000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassX, Rn).Name("ScvtfXToS"),
		L("scvtf", 0x9e620000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassX, Rn).Name("ScvtfXToD"),
		L("ucvtf", 0x1e230000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassW, Rn).Name("UcvtfWToS"),
		L("ucvtf", 0x1e630000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassW, Rn).Name("UcvtfWToD"),
		L("ucvtf", 0x9e230000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassX, Rn).Name("UcvtfXToS"),
		L("ucvtf", 0x9e630000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassX, Rn).Name("UcvtfXToD"),

		// ---- Float to integer, round toward zero ----
		//
		// Saturating, not trapping: out of range gives the nearest
		// representable integer and a NaN gives zero. A language wanting a
		// trap checks the range first.
		L("fcvtzs", 0x1e380000, 0xfffffc00).
			Dst(ClassW, Rd).Src(ClassS, Rn).Name("FcvtzsSToW"),
		L("fcvtzs", 0x1e780000, 0xfffffc00).
			Dst(ClassW, Rd).Src(ClassD, Rn).Name("FcvtzsDToW"),
		L("fcvtzs", 0x9e380000, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassS, Rn).Name("FcvtzsSToX"),
		L("fcvtzs", 0x9e780000, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassD, Rn).Name("FcvtzsDToX"),
		L("fcvtzu", 0x1e390000, 0xfffffc00).
			Dst(ClassW, Rd).Src(ClassS, Rn).Name("FcvtzuSToW"),
		L("fcvtzu", 0x1e790000, 0xfffffc00).
			Dst(ClassW, Rd).Src(ClassD, Rn).Name("FcvtzuDToW"),
		L("fcvtzu", 0x9e390000, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassS, Rn).Name("FcvtzuSToX"),
		L("fcvtzu", 0x9e790000, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassD, Rn).Name("FcvtzuDToX"),

		// ---- Bit pattern between the two register files ----
		//
		// FMOV and not a conversion: the bits are copied, not reinterpreted,
		// which is what a bitcast needs and what a conversion is not.
		L("fmov", 0x1e260000, 0xfffffc00).
			Dst(ClassW, Rd).Src(ClassS, Rn).Name("FmovSToW"),
		L("fmov", 0x1e270000, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassW, Rn).Name("FmovWToS"),
		L("fmov", 0x9e660000, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassD, Rn).Name("FmovDToX"),
		L("fmov", 0x9e670000, 0xfffffc00).
			Dst(ClassD, Rd).Src(ClassX, Rn).Name("FmovXToD"),

		// ---- Load and store, scaled unsigned offset ----
		L("ldr", 0xbd400000, 0xffc00000).
			Dst(ClassS, Rt).Mem(32, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrImmS"),
		L("ldr", 0xfd400000, 0xffc00000).
			Dst(ClassD, Rt).Mem(64, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrImmD"),
		L("ldr", 0x3dc00000, 0xffc00000).
			Dst(ClassQ, Rt).Mem(128, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrImmQ"),
		L("str", 0xbd000000, 0xffc00000).
			Src(ClassS, Rt).Mem(32, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrImmS"),
		L("str", 0xfd000000, 0xffc00000).
			Src(ClassD, Rt).Mem(64, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrImmD"),
		L("str", 0x3d800000, 0xffc00000).
			Src(ClassQ, Rt).Mem(128, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrImmQ"),

		// ---- Load and store, unscaled ----
		L("ldur", 0xbc400000, 0xffe00c00).
			Dst(ClassS, Rt).Mem(32, Rn, Imm9).Kind(ImmUnscaled).Name("LdurImmS"),
		L("ldur", 0xfc400000, 0xffe00c00).
			Dst(ClassD, Rt).Mem(64, Rn, Imm9).Kind(ImmUnscaled).Name("LdurImmD"),
		L("stur", 0xbc000000, 0xffe00c00).
			Src(ClassS, Rt).Mem(32, Rn, Imm9).Kind(ImmUnscaled).Name("SturImmS"),
		L("stur", 0xfc000000, 0xffe00c00).
			Src(ClassD, Rt).Mem(64, Rn, Imm9).Kind(ImmUnscaled).Name("SturImmD"),
	)
}
