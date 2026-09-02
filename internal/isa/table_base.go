package isa

// The base integer, branch and load/store instruction set.
//
// Base words and masks are the ARM ARM's, one row per declared encoding per
// width. The 32- and 64-bit variants are separate rows rather than one row with
// a varying sf bit, because they are separate forms to a caller: AddImm32 and
// AddImm64 take different register types and a width mismatch should be a
// compile error at the typed call rather than a runtime check.
//
// This table is the part of the tree that must be differential-tested rather
// than reviewed. Every row here is checked against GNU as by internal/difftest.
func init() {
	register(
		// ---- Arithmetic, immediate ----
		L("add", 0x11000000, 0x7f800000).
			Dst(ClassWsp, Rd).Src(ClassWsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Name("AddImm32"),
		L("add", 0x91000000, 0xff800000).
			Dst(ClassXsp, Rd).Src(ClassXsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Name("AddImm64"),
		L("adds", 0x31000000, 0x7f800000).
			Dst(ClassW, Rd).Src(ClassWsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().Name("AddsImm32"),
		L("adds", 0xb1000000, 0xff800000).
			Dst(ClassX, Rd).Src(ClassXsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().Name("AddsImm64"),
		L("sub", 0x51000000, 0x7f800000).
			Dst(ClassWsp, Rd).Src(ClassWsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Name("SubImm32"),
		L("sub", 0xd1000000, 0xff800000).
			Dst(ClassXsp, Rd).Src(ClassXsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Name("SubImm64"),
		L("subs", 0x71000000, 0x7f800000).
			Dst(ClassW, Rd).Src(ClassWsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().Name("SubsImm32"),
		L("subs", 0xf1000000, 0xff800000).
			Dst(ClassX, Rd).Src(ClassXsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().Name("SubsImm64"),

		// ---- Arithmetic, shifted register ----
		L("add", 0x0b000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Name("AddShifted32"),
		L("add", 0x8b000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Name("AddShifted64"),
		L("adds", 0x2b000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().Name("AddsShifted32"),
		L("adds", 0xab000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().Name("AddsShifted64"),
		L("sub", 0x4b000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Name("SubShifted32"),
		L("sub", 0xcb000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Name("SubShifted64"),
		L("subs", 0x6b000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().Name("SubsShifted32"),
		L("subs", 0xeb000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().Name("SubsShifted64"),

		// ---- Arithmetic, extended register ----
		L("add", 0x0b200000, 0x7fe00000).
			Dst(ClassWsp, Rd).Src(ClassWsp, Rn).Src(ClassW, Rm).Opt(ClassExtend, Opt, 2).
			Name("AddExt32"),
		L("add", 0x8b200000, 0xffe00000).
			Dst(ClassXsp, Rd).Src(ClassXsp, Rn).Src(ClassX, Rm).Opt(ClassExtend, Opt, 3).
			Name("AddExt64"),
		L("sub", 0xcb200000, 0xffe00000).
			Dst(ClassXsp, Rd).Src(ClassXsp, Rn).Src(ClassX, Rm).Opt(ClassExtend, Opt, 3).
			Name("SubExt64"),
		L("subs", 0xeb200000, 0xffe00000).
			Dst(ClassX, Rd).Src(ClassXsp, Rn).Src(ClassX, Rm).Opt(ClassExtend, Opt, 3).
			Flags().Name("SubsExt64"),

		// ---- Logical, shifted register ----
		L("and", 0x0a000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Name("AndShifted32"),
		L("and", 0x8a000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Name("AndShifted64"),
		L("orr", 0x2a000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Name("OrrShifted32"),
		L("orr", 0xaa000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Name("OrrShifted64"),
		L("eor", 0x4a000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Name("EorShifted32"),
		L("eor", 0xca000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Name("EorShifted64"),
		L("ands", 0x6a000000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().Name("AndsShifted32"),
		L("ands", 0xea000000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().Name("AndsShifted64"),

		// ---- Logical, bitmask immediate ----
		// The immediate is the N:immr:imms triple operand/bitmask.go computes.
		// Whether a constant is expressible at all is that encoder's question,
		// asked before this form is ever reached.
		L("and", 0x12000000, 0x7f800000).
			Dst(ClassWsp, Rd).Src(ClassW, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Name("AndImm32"),
		L("and", 0x92000000, 0xff800000).
			Dst(ClassXsp, Rd).Src(ClassX, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Name("AndImm64"),
		L("orr", 0x32000000, 0x7f800000).
			Dst(ClassWsp, Rd).Src(ClassW, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Name("OrrImm32"),
		L("orr", 0xb2000000, 0xff800000).
			Dst(ClassXsp, Rd).Src(ClassX, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Name("OrrImm64"),
		L("eor", 0x52000000, 0x7f800000).
			Dst(ClassWsp, Rd).Src(ClassW, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Name("EorImm32"),
		L("eor", 0xd2000000, 0xff800000).
			Dst(ClassXsp, Rd).Src(ClassX, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Name("EorImm64"),
		L("ands", 0x72000000, 0x7f800000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Flags().Name("AndsImm32"),
		L("ands", 0xf2000000, 0xff800000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Flags().Name("AndsImm64"),

		// ---- Move wide ----
		L("movz", 0x52800000, 0x7f800000).
			Dst(ClassW, Rd).Imm(Imm16).Kind(ImmMoveWide).Opt(ClassShift, Hw, 0).
			Name("MovzImm32"),
		L("movz", 0xd2800000, 0xff800000).
			Dst(ClassX, Rd).Imm(Imm16).Kind(ImmMoveWide).Opt(ClassShift, Hw, 0).
			Name("MovzImm64"),
		L("movn", 0x12800000, 0x7f800000).
			Dst(ClassW, Rd).Imm(Imm16).Kind(ImmMoveWide).Opt(ClassShift, Hw, 0).
			Name("MovnImm32"),
		L("movn", 0x92800000, 0xff800000).
			Dst(ClassX, Rd).Imm(Imm16).Kind(ImmMoveWide).Opt(ClassShift, Hw, 0).
			Name("MovnImm64"),
		L("movk", 0x72800000, 0x7f800000).
			SrcDst(ClassW, Rd).Imm(Imm16).Kind(ImmMoveWide).Opt(ClassShift, Hw, 0).
			Name("MovkImm32"),
		L("movk", 0xf2800000, 0xff800000).
			SrcDst(ClassX, Rd).Imm(Imm16).Kind(ImmMoveWide).Opt(ClassShift, Hw, 0).
			Name("MovkImm64"),

		// ---- PC-relative address ----
		L("adr", 0x10000000, 0x9f000000).
			Dst(ClassX, Rd).Addr(RoleTarget, ImmPCRel).
			Name("Adr"),
		L("adrp", 0x90000000, 0x9f000000).
			Dst(ClassX, Rd).Addr(RolePage, ImmPCRel).Kind(ImmPage).
			Name("Adrp"),

		// ---- Bitfield ----
		L("ubfm", 0x53000000, 0x7f800000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(Immr).Imm(Imms).Name("Ubfm32"),
		L("ubfm", 0xd3400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(Immr).Imm(Imms).Name("Ubfm64"),
		L("sbfm", 0x13000000, 0x7f800000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(Immr).Imm(Imms).Name("Sbfm32"),
		L("sbfm", 0x93400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(Immr).Imm(Imms).Name("Sbfm64"),
		L("bfm", 0x33000000, 0x7f800000).
			SrcDst(ClassW, Rd).Src(ClassW, Rn).Imm(Immr).Imm(Imms).Name("Bfm32"),
		L("bfm", 0xb3400000, 0xffc00000).
			SrcDst(ClassX, Rd).Src(ClassX, Rn).Imm(Immr).Imm(Imms).Name("Bfm64"),
		L("extr", 0x13800000, 0x7fa00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Imm(Imms).Name("Extr32"),
		L("extr", 0x93c00000, 0xffe00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Imm(Imms).Name("Extr64"),

		// ---- Variable shift, multiply, divide ----
		L("lslv", 0x1ac02000, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Lslv32"),
		L("lslv", 0x9ac02000, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Lslv64"),
		L("lsrv", 0x1ac02400, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Lsrv32"),
		L("lsrv", 0x9ac02400, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Lsrv64"),
		L("asrv", 0x1ac02800, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Asrv32"),
		L("asrv", 0x9ac02800, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Asrv64"),
		L("rorv", 0x1ac02c00, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Rorv32"),
		L("rorv", 0x9ac02c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Rorv64"),
		L("udiv", 0x1ac00800, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Udiv32"),
		L("udiv", 0x9ac00800, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Udiv64"),
		L("sdiv", 0x1ac00c00, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Sdiv32"),
		L("sdiv", 0x9ac00c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Sdiv64"),
		L("madd", 0x1b000000, 0x7fe08000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Src(ClassW, Ra).Name("Madd32"),
		L("madd", 0x9b000000, 0xffe08000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Src(ClassX, Ra).Name("Madd64"),
		L("msub", 0x1b008000, 0x7fe08000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Src(ClassW, Ra).Name("Msub32"),
		L("msub", 0x9b008000, 0xffe08000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Src(ClassX, Ra).Name("Msub64"),

		// The high half of a 64x64 product, which no MADD can reach: the
		// multiply is 64x64 into 64 everywhere else, and this is the other
		// half of it. Ra is fixed to the zero register in the encoding rather
		// than being an operand, so there is no addend form to alias from.
		L("smulh", 0x9b407c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Smulh"),
		L("umulh", 0x9bc07c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Umulh"),

		// ---- Conditional select and compare ----
		// ---- Add and subtract with carry ----
		//
		// The instruction a multiword addition is made of, and absent until
		// an assembler asked for it: nothing here selects one, because this
		// IR's integers fit a register and its overflow predicates are
		// computed rather than carried.
		L("adc", 0x1a000000, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Adc32"),
		L("adc", 0x9a000000, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Adc64"),
		L("adcs", 0x3a000000, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Flags().Name("Adcs32"),
		L("adcs", 0xba000000, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Flags().Name("Adcs64"),
		L("sbc", 0x5a000000, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Name("Sbc32"),
		L("sbc", 0xda000000, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Name("Sbc64"),
		L("sbcs", 0x7a000000, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Flags().Name("Sbcs32"),
		L("sbcs", 0xfa000000, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Flags().Name("Sbcs64"),

		// ---- Widening multiply ----
		//
		// A 32-bit multiply with a 64-bit result, which is the one shape
		// this architecture cannot express by choosing register widths: the
		// sources are W and the destination and accumulator are X.
		L("smaddl", 0x9b200000, 0xffe08000).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Src(ClassX, Ra).Name("Smaddl"),
		L("umaddl", 0x9ba00000, 0xffe08000).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Src(ClassX, Ra).Name("Umaddl"),
		L("smsubl", 0x9b208000, 0xffe08000).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Src(ClassX, Ra).Name("Smsubl"),
		L("umsubl", 0x9ba08000, 0xffe08000).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Src(ClassX, Ra).Name("Umsubl"),

		// ---- Conditional compare ----
		//
		// A compare that happens only if the condition holds, and writes the
		// four immediate flags when it does not. That is what makes a chain
		// of C's && and || into straight-line code, and it is why the NZCV
		// operand is a value and not a modifier: it is the answer for the
		// half of the chain this instruction does not evaluate.
		L("ccmp", 0xfa400000, 0xffe00c10).
			Src(ClassX, Rn).Src(ClassX, Rm).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmpReg64"),
		L("ccmp", 0x7a400000, 0x7fe00c10).
			Src(ClassW, Rn).Src(ClassW, Rm).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmpReg32"),
		L("ccmp", 0xfa400800, 0xffe00c10).
			Src(ClassX, Rn).Imm(Imm5).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmpImm64"),
		L("ccmp", 0x7a400800, 0x7fe00c10).
			Src(ClassW, Rn).Imm(Imm5).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmpImm32"),
		L("ccmn", 0xba400000, 0xffe00c10).
			Src(ClassX, Rn).Src(ClassX, Rm).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmnReg64"),
		L("ccmn", 0x3a400000, 0x7fe00c10).
			Src(ClassW, Rn).Src(ClassW, Rm).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmnReg32"),
		L("ccmn", 0xba400800, 0xffe00c10).
			Src(ClassX, Rn).Imm(Imm5).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmnImm64"),
		L("ccmn", 0x3a400800, 0x7fe00c10).
			Src(ClassW, Rn).Imm(Imm5).Imm(Nzcv).Cnd(CondHi).
			Flags().Name("CcmnImm32"),

		L("csel", 0x1a800000, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Cnd(CondHi).Name("Csel32"),
		L("csel", 0x9a800000, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Cnd(CondHi).Name("Csel64"),
		L("csinc", 0x1a800400, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Cnd(CondHi).Name("Csinc32"),
		L("csinc", 0x9a800400, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Cnd(CondHi).Name("Csinc64"),
		L("csinv", 0x5a800000, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Cnd(CondHi).Name("Csinv32"),
		L("csinv", 0xda800000, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Cnd(CondHi).Name("Csinv64"),
		L("csneg", 0x5a800400, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Cnd(CondHi).Name("Csneg32"),
		L("csneg", 0xda800400, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Cnd(CondHi).Name("Csneg64"),

		// ---- Data processing, one source ----
		L("rbit", 0x5ac00000, 0x7ffffc00).Dst(ClassW, Rd).Src(ClassW, Rn).Name("Rbit32"),
		L("rbit", 0xdac00000, 0xfffffc00).Dst(ClassX, Rd).Src(ClassX, Rn).Name("Rbit64"),
		L("orn", 0x2a200000, 0x7f200000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Name("OrnShifted32"),
		L("orn", 0xaa200000, 0xff200000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Name("OrnShifted64"),

		L("rev16", 0x5ac00400, 0x7ffffc00).Dst(ClassW, Rd).Src(ClassW, Rn).Name("Rev16_32"),
		L("rev16", 0xdac00400, 0xfffffc00).Dst(ClassX, Rd).Src(ClassX, Rn).Name("Rev16_64"),
		L("rev", 0x5ac00800, 0x7ffffc00).Dst(ClassW, Rd).Src(ClassW, Rn).Name("Rev32"),
		L("rev", 0xdac00c00, 0xfffffc00).Dst(ClassX, Rd).Src(ClassX, Rn).Name("Rev64"),
		// REV64 is the same word under the name the ARM ARM prefers when
		// the register is an X; REV32 on an X register is a different
		// instruction, reversing the bytes within each word rather than
		// across the whole of it.
		L("rev64", 0xdac00c00, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).AliasOf("rev").Name("Rev64Alias"),
		L("rev32", 0xdac00800, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Name("Rev32In64"),
		L("clz", 0x5ac01000, 0x7ffffc00).Dst(ClassW, Rd).Src(ClassW, Rn).Name("Clz32"),
		L("clz", 0xdac01000, 0xfffffc00).Dst(ClassX, Rd).Src(ClassX, Rn).Name("Clz64"),
		L("cls", 0x5ac01400, 0x7ffffc00).Dst(ClassW, Rd).Src(ClassW, Rn).Name("Cls32"),
		L("cls", 0xdac01400, 0xfffffc00).Dst(ClassX, Rd).Src(ClassX, Rn).Name("Cls64"),

		// ---- Branches ----
		L("b", 0x14000000, 0xfc000000).Target(Imm26).Kind(ImmBranch).Name("B"),
		L("bl", 0x94000000, 0xfc000000).Target(Imm26).Kind(ImmBranch).Name("Bl"),
		L("b.cond", 0x54000000, 0xff000010).Cnd(Cond).Target(Imm19).Kind(ImmBranch).Name("BCond"),
		L("cbz", 0x34000000, 0x7f000000).Src(ClassW, Rt).Target(Imm19).Kind(ImmBranch).Name("Cbz32"),
		L("cbz", 0xb4000000, 0xff000000).Src(ClassX, Rt).Target(Imm19).Kind(ImmBranch).Name("Cbz64"),
		L("cbnz", 0x35000000, 0x7f000000).Src(ClassW, Rt).Target(Imm19).Kind(ImmBranch).Name("Cbnz32"),
		L("cbnz", 0xb5000000, 0xff000000).Src(ClassX, Rt).Target(Imm19).Kind(ImmBranch).Name("Cbnz64"),
		L("tbz", 0x36000000, 0x7f000000).
			Src(ClassX, Rt).Imm(BitPos).Kind(ImmBitPos).Target(Imm14).Kind(ImmBranch).Name("Tbz64"),
		L("tbnz", 0x37000000, 0x7f000000).
			Src(ClassX, Rt).Imm(BitPos).Kind(ImmBitPos).Target(Imm14).Kind(ImmBranch).Name("Tbnz64"),
		// The 32-bit forms are the same words with b5 necessarily zero, and
		// they are separate rows because the bit number's range is what
		// differs: a W register has no bit 32 to test.
		L("tbz", 0x36000000, 0xff000000).
			Src(ClassW, Rt).Imm(BitPos5).Kind(ImmBitPos).Target(Imm14).Kind(ImmBranch).Name("Tbz32"),
		L("tbnz", 0x37000000, 0xff000000).
			Src(ClassW, Rt).Imm(BitPos5).Kind(ImmBitPos).Target(Imm14).Kind(ImmBranch).Name("Tbnz32"),
		// ERET returns from an exception, and takes no operand because the
		// address it returns to is in ELR_ELn rather than in a register.
		L("eret", 0xd69f03e0, 0xffffffff).Name("Eret"),
		L("br", 0xd61f0000, 0xfffffc1f).Src(ClassX, Rn).Name("Br"),
		L("blr", 0xd63f0000, 0xfffffc1f).Src(ClassX, Rn).Name("Blr"),
		// RET's operand is optional and defaults to X30. That default is the
		// architecture's, stated in the encoding, not an alias.
		L("ret", 0xd65f0000, 0xfffffc1f).OptReg(ClassX, Rn, 30).Name("Ret"),

		// ---- Load and store, unsigned scaled offset ----
		L("strb", 0x39000000, 0xffc00000).
			Src(ClassW, Rt).Mem(8, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrbImm"),
		L("ldrb", 0x39400000, 0xffc00000).
			Dst(ClassW, Rt).Mem(8, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrbImm"),
		L("strh", 0x79000000, 0xffc00000).
			Src(ClassW, Rt).Mem(16, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrhImm"),
		L("ldrh", 0x79400000, 0xffc00000).
			Dst(ClassW, Rt).Mem(16, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrhImm"),
		L("str", 0xb9000000, 0xffc00000).
			Src(ClassW, Rt).Mem(32, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrImm32"),
		L("ldr", 0xb9400000, 0xffc00000).
			Dst(ClassW, Rt).Mem(32, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrImm32"),
		L("str", 0xf9000000, 0xffc00000).
			Src(ClassX, Rt).Mem(64, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("StrImm64"),
		L("ldr", 0xf9400000, 0xffc00000).
			Dst(ClassX, Rt).Mem(64, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrImm64"),
		L("ldrsb", 0x39c00000, 0xffc00000).
			Dst(ClassW, Rt).Mem(8, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrsbImm32"),
		L("ldrsb", 0x39800000, 0xffc00000).
			Dst(ClassX, Rt).Mem(8, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrsbImm64"),
		L("ldrsh", 0x79c00000, 0xffc00000).
			Dst(ClassW, Rt).Mem(16, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrshImm32"),
		L("ldrsh", 0x79800000, 0xffc00000).
			Dst(ClassX, Rt).Mem(16, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrshImm64"),
		// PRFM's destination is a hint rather than a register: the operand
		// says what to prefetch and into which cache, and occupies the field
		// a load's Rt would.
		L("prfm", 0xf9800000, 0xffc00000).
			Prf(Rt).Mem(64, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("PrfmImm"),
		L("ldrsw", 0xb9800000, 0xffc00000).
			Dst(ClassX, Rt).Mem(32, Rn, Imm12).Kind(ImmScaled).Attr(AttrScaled).Name("LdrswImm"),

		// ---- Load and store, register offset ----
		//
		// [Xn, Xm{, LSL #s}] and [Xn, Wm, SXTW #s]: the addressing mode a
		// subscripted array uses, and the one an assembler meets first.
		// Bit 21 and bits 11:10 make it a different encoding from the
		// scaled-immediate form rather than a different operand of it,
		// which is why these are rows and not an attribute on those.
		L("strb", 0x38200800, 0xffe00c00).
			Src(ClassW, Rt).MemIdx(8, Rn).Name("StrbReg"),
		L("ldrb", 0x38600800, 0xffe00c00).
			Dst(ClassW, Rt).MemIdx(8, Rn).Name("LdrbReg"),
		L("strh", 0x78200800, 0xffe00c00).
			Src(ClassW, Rt).MemIdx(16, Rn).Name("StrhReg"),
		L("ldrh", 0x78600800, 0xffe00c00).
			Dst(ClassW, Rt).MemIdx(16, Rn).Name("LdrhReg"),
		L("str", 0xb8200800, 0xffe00c00).
			Src(ClassW, Rt).MemIdx(32, Rn).Name("StrReg32"),
		L("ldr", 0xb8600800, 0xffe00c00).
			Dst(ClassW, Rt).MemIdx(32, Rn).Name("LdrReg32"),
		L("str", 0xf8200800, 0xffe00c00).
			Src(ClassX, Rt).MemIdx(64, Rn).Name("StrReg64"),
		L("ldr", 0xf8600800, 0xffe00c00).
			Dst(ClassX, Rt).MemIdx(64, Rn).Name("LdrReg64"),
		L("ldrsb", 0x38e00800, 0xffe00c00).
			Dst(ClassW, Rt).MemIdx(8, Rn).Name("LdrsbReg32"),
		L("ldrsb", 0x38a00800, 0xffe00c00).
			Dst(ClassX, Rt).MemIdx(8, Rn).Name("LdrsbReg64"),
		L("ldrsh", 0x78e00800, 0xffe00c00).
			Dst(ClassW, Rt).MemIdx(16, Rn).Name("LdrshReg32"),
		L("ldrsh", 0x78a00800, 0xffe00c00).
			Dst(ClassX, Rt).MemIdx(16, Rn).Name("LdrshReg64"),
		L("ldrsw", 0xb8a00800, 0xffe00c00).
			Dst(ClassX, Rt).MemIdx(32, Rn).Name("LdrswReg"),
		L("prfm", 0xf8a00800, 0xffe00c00).
			Prf(Rt).MemIdx(64, Rn).Name("PrfmReg"),

		// ---- Load and store, unscaled ----
		L("stur", 0xb8000000, 0xffe00c00).
			Src(ClassW, Rt).Mem(32, Rn, Imm9).Kind(ImmUnscaled).Name("SturImm32"),
		L("ldur", 0xb8400000, 0xffe00c00).
			Dst(ClassW, Rt).Mem(32, Rn, Imm9).Kind(ImmUnscaled).Name("LdurImm32"),
		L("stur", 0xf8000000, 0xffe00c00).
			Src(ClassX, Rt).Mem(64, Rn, Imm9).Kind(ImmUnscaled).Name("SturImm64"),
		L("ldur", 0xf8400000, 0xffe00c00).
			Dst(ClassX, Rt).Mem(64, Rn, Imm9).Kind(ImmUnscaled).Name("LdurImm64"),

		// ---- Load and store pair ----
		L("stp", 0x29000000, 0xffc00000).
			Src(ClassW, Rt).Src(ClassW, Rt2).Mem(32, Rn, Imm7).Kind(ImmScaled).
			Attr(AttrScaled).Name("Stp32"),
		L("ldp", 0x29400000, 0xffc00000).
			Dst(ClassW, Rt).Dst(ClassW, Rt2).Mem(32, Rn, Imm7).Kind(ImmScaled).
			Attr(AttrScaled).Name("Ldp32"),
		L("stp", 0xa9000000, 0xffc00000).
			Src(ClassX, Rt).Src(ClassX, Rt2).Mem(64, Rn, Imm7).Kind(ImmScaled).
			Attr(AttrScaled).Name("Stp64"),
		L("ldp", 0xa9400000, 0xffc00000).
			Dst(ClassX, Rt).Dst(ClassX, Rt2).Mem(64, Rn, Imm7).Kind(ImmScaled).
			Attr(AttrScaled).Name("Ldp64"),
		L("stp", 0xa9800000, 0xffc00000).
			Src(ClassX, Rt).Src(ClassX, Rt2).Mem(64, Rn, Imm7).Kind(ImmScaled).
			Attr(AttrScaled|AttrPreIndex).Name("StpPre64"),
		L("ldp", 0xa8c00000, 0xffc00000).
			Dst(ClassX, Rt).Dst(ClassX, Rt2).Mem(64, Rn, Imm7).Kind(ImmScaled).
			Attr(AttrScaled|AttrPostIndex).Name("LdpPost64"),

		// ---- Load literal ----
		L("ldr", 0x18000000, 0xff000000).
			Dst(ClassW, Rt).Target(Imm19).Kind(ImmBranch).Name("LdrLit32"),
		L("ldr", 0x58000000, 0xff000000).
			Dst(ClassX, Rt).Target(Imm19).Kind(ImmBranch).Name("LdrLit64"),

		// ---- Exceptions, hints, barriers ----
		L("svc", 0xd4000001, 0xffe0001f).Imm(Imm16).Name("Svc"),
		L("hvc", 0xd4000002, 0xffe0001f).Imm(Imm16).Name("Hvc"),
		L("smc", 0xd4000003, 0xffe0001f).Imm(Imm16).Name("Smc"),
		L("brk", 0xd4200000, 0xffe0001f).Imm(Imm16).Name("Brk"),
		L("hlt", 0xd4400000, 0xffe0001f).Imm(Imm16).Name("Hlt"),

		// The hint space. NOP, YIELD and the rest below are named hints, and
		// this is the rest of it: a bare number, which is what a barrier or
		// a pointer-authentication instruction is on a processor that does
		// not implement one.
		L("hint", 0xd503201f, 0xfffff01f).Imm(F(5, 7)).Name("Hint"),
		L("paciasp", 0xd503233f, 0xffffffff).AliasOf("hint").Name("Paciasp"),
		L("autiasp", 0xd50323bf, 0xffffffff).AliasOf("hint").Name("Autiasp"),
		L("nop", 0xd503201f, 0xffffffff).Name("Nop"),
		L("yield", 0xd503203f, 0xffffffff).Name("Yield"),
		L("wfe", 0xd503205f, 0xffffffff).Name("Wfe"),
		L("wfi", 0xd503207f, 0xffffffff).Name("Wfi"),
		L("sev", 0xd503209f, 0xffffffff).Name("Sev"),
		L("sevl", 0xd50320bf, 0xffffffff).Name("Sevl"),

		L("dsb", 0xd503309f, 0xfffff0ff).Opt(ClassBarrier, CRm, 15).Name("Dsb"),
		L("dmb", 0xd50330bf, 0xfffff0ff).Opt(ClassBarrier, CRm, 15).Name("Dmb"),
		L("isb", 0xd50330df, 0xfffff0ff).Opt(ClassBarrier, CRm, 15).Name("Isb"),

		// ---- System register move ----
		L("mrs", 0xd5300000, 0xfff00000).
			Dst(ClassX, Rt).SysReg(SysReg).Name("Mrs"),
		L("msr", 0xd5100000, 0xfff00000).
			SysReg(SysReg).Src(ClassX, Rt).Name("MsrReg"),
	)

	// ---- The architecture's aliases ----
	//
	// Each is one-to-one with a word of the instruction it aliases, and each
	// pins the field that makes it narrower. The preferred-disassembly
	// predicates are the ARM ARM's own, stated per alias.
	register(
		L("cmp", 0xeb00001f, 0xff20001f).
			Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("subs").Pins(Rd, 31).Name("CmpShifted64"),
		L("cmp", 0xf100001f, 0xff80001f).
			Src(ClassXsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().AliasOf("subs").Pins(Rd, 31).Name("CmpImm64"),
		L("cmp", 0x7100001f, 0x7f80001f).
			Src(ClassWsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().AliasOf("subs").Pins(Rd, 31).Name("CmpImm32"),
		L("cmp", 0x6b00001f, 0x7f20001f).
			Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("subs").Pins(Rd, 31).Name("CmpShifted32"),
		L("cmn", 0xab00001f, 0xff20001f).
			Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("adds").Pins(Rd, 31).Name("CmnShifted64"),
		L("cmn", 0x2b00001f, 0x7f20001f).
			Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("adds").Pins(Rd, 31).Name("CmnShifted32"),
		L("cmn", 0xb100001f, 0xff80001f).
			Src(ClassXsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().AliasOf("adds").Pins(Rd, 31).Name("CmnImm64"),
		L("cmn", 0x3100001f, 0x7f80001f).
			Src(ClassWsp, Rn).Imm(Imm12).Kind(ImmAddSub12).Opt(ClassShift, Sh, 0).
			Flags().AliasOf("adds").Pins(Rd, 31).Name("CmnImm32"),
		L("tst", 0xea00001f, 0xff20001f).
			Src(ClassX, Rn).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("ands").Pins(Rd, 31).Name("TstShifted64"),
		L("tst", 0xf200001f, 0xff80001f).
			Src(ClassX, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Flags().AliasOf("ands").Pins(Rd, 31).Name("TstImm64"),
		L("tst", 0x6a00001f, 0x7f20001f).
			Src(ClassW, Rn).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("ands").Pins(Rd, 31).Name("TstShifted32"),
		L("tst", 0x7200001f, 0x7f80001f).
			Src(ClassW, Rn).Imm(ImmLogical13).Kind(ImmLogical).
			Flags().AliasOf("ands").Pins(Rd, 31).Name("TstImm32"),
		L("neg", 0xcb0003e0, 0xff2003e0).
			Dst(ClassX, Rd).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			AliasOf("sub").Pins(Rn, 31).Name("NegShifted64"),
		L("neg", 0x4b0003e0, 0x7f2003e0).
			Dst(ClassW, Rd).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			AliasOf("sub").Pins(Rn, 31).Name("NegShifted32"),

		// MVN is ORN with the zero register as its first source, which is the
		// same relation MOV has to ORR one row further down.
		L("mvn", 0xaa2003e0, 0xff2003e0).
			Dst(ClassX, Rd).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			AliasOf("orn").Pins(Rn, 31).Name("MvnShifted64"),
		L("mvn", 0x2a2003e0, 0x7f2003e0).
			Dst(ClassW, Rd).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			AliasOf("orn").Pins(Rn, 31).Name("MvnShifted32"),

		// MOV (register) is ORR with the zero register as its first source.
		// It is preferred only when no shift is applied; ORR with a shift is
		// not a move and the ARM ARM says so.
		L("mov", 0xaa0003e0, 0xffe0ffe0).
			Dst(ClassX, Rd).Src(ClassX, Rm).
			AliasOf("orr").Pins(Rn, 31).Name("MovReg64"),
		L("mov", 0x2a0003e0, 0x7fe0ffe0).
			Dst(ClassW, Rd).Src(ClassW, Rm).
			AliasOf("orr").Pins(Rn, 31).Name("MovReg32"),

		// MOV (to/from SP) is ADD with a zero immediate, and is preferred only
		// when one of the two registers really is SP. Otherwise the word is a
		// plain add of zero and prints as one.
		L("mov", 0x91000000, 0xfffffc00).
			Dst(ClassXsp, Rd).Src(ClassXsp, Rn).
			AliasOf("add").Pins(Imm12, 0).
			PreferredWhen(func(w uint32) bool {
				return Rd.Get(w) == 31 || Rn.Get(w) == 31
			}).Name("MovSp64"),

		// MOV (wide immediate) is MOVZ, preferred unless the immediate is zero
		// and shifted — movz x0, #0, lsl #16 is not a move of anything.
		L("mov", 0x52800000, 0x7f800000).
			Dst(ClassW, Rd).Imm(Imm16).Kind(ImmMoveWide).
			AliasOf("movz").
			PreferredWhen(func(w uint32) bool {
				return !(Imm16.Get(w) == 0 && Hw.Get(w) != 0)
			}).Name("MovWide32"),
		L("mov", 0xd2800000, 0xff800000).
			Dst(ClassX, Rd).Imm(Imm16).Kind(ImmMoveWide).
			AliasOf("movz").
			PreferredWhen(func(w uint32) bool {
				return !(Imm16.Get(w) == 0 && Hw.Get(w) != 0)
			}).Name("MovWide64"),

		// The shift aliases are UBFM and SBFM with computed immediates. The
		// computation is encode/'s; the alias relation is the architecture's.
		L("lsl", 0xd3400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(Immr).Kind(ImmShiftLeft).
			AliasOf("ubfm").Name("LslImm64"),
		L("lsr", 0xd3400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(Immr).Kind(ImmShiftRight).
			AliasOf("ubfm").Name("LsrImm64"),
		L("asr", 0x93400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(Immr).Kind(ImmShiftRight).
			AliasOf("sbfm").Name("AsrImm64"),
		L("lsl", 0x53000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(Immr).Kind(ImmShiftLeft).
			AliasOf("ubfm").Name("LslImm32"),
		L("lsr", 0x53000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(Immr).Kind(ImmShiftRight).
			AliasOf("ubfm").Name("LsrImm32"),
		L("asr", 0x13000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(Immr).Kind(ImmShiftRight).
			AliasOf("sbfm").Name("AsrImm32"),

		// The register forms of the same four shifts are the variable-shift
		// instructions under their assembly-language names. Unlike the
		// immediate forms these pin nothing and compute nothing: LSL Wd, Wn,
		// Wm and LSLV Wd, Wn, Wm are one word spelled two ways.
		L("lsl", 0x9ac02000, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).
			AliasOf("lslv").Name("LslReg64"),
		L("lsl", 0x1ac02000, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("lslv").Name("LslReg32"),
		L("lsr", 0x9ac02400, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).
			AliasOf("lsrv").Name("LsrReg64"),
		L("lsr", 0x1ac02400, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("lsrv").Name("LsrReg32"),
		L("asr", 0x9ac02800, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).
			AliasOf("asrv").Name("AsrReg64"),
		L("asr", 0x1ac02800, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("asrv").Name("AsrReg32"),
		L("ror", 0x9ac02c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).
			AliasOf("rorv").Name("RorReg64"),
		L("ror", 0x1ac02c00, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("rorv").Name("RorReg32"),

		// The sign- and zero-extend aliases are SBFM and UBFM with both
		// immediates fixed, so they carry no immediate operand at all: the
		// width is in the mnemonic. SXTB Xd, Wn is the odd-looking one and is
		// the architecture's own spelling — the source is named as a W
		// register because only its low eight bits are read.
		L("sxtb", 0x13001c00, 0x7ffffc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			AliasOf("sbfm").Pins(Immr, 0).Pins(Imms, 7).Name("Sxtb32"),
		L("sxtb", 0x93401c00, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).
			AliasOf("sbfm").Pins(Immr, 0).Pins(Imms, 7).Name("Sxtb64"),
		L("sxth", 0x13003c00, 0x7ffffc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			AliasOf("sbfm").Pins(Immr, 0).Pins(Imms, 15).Name("Sxth32"),
		L("sxth", 0x93403c00, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).
			AliasOf("sbfm").Pins(Immr, 0).Pins(Imms, 15).Name("Sxth64"),
		L("sxtw", 0x93407c00, 0xfffffc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).
			AliasOf("sbfm").Pins(Immr, 0).Pins(Imms, 31).Name("Sxtw64"),
		L("uxtb", 0x53001c00, 0x7ffffc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			AliasOf("ubfm").Pins(Immr, 0).Pins(Imms, 7).Name("Uxtb32"),
		L("uxth", 0x53003c00, 0x7ffffc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			AliasOf("ubfm").Pins(Immr, 0).Pins(Imms, 15).Name("Uxth32"),

		// UBFX and SBFX are UBFM and SBFM under the operands a programmer
		// has: where the field starts and how wide it is, rather than where
		// it starts and where it ends.
		L("ubfx", 0x53000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			Imm(Immr).Kind(ImmBitfieldLsb).Imm(Imms).Kind(ImmBitfieldWidth).
			AliasOf("ubfm").Name("Ubfx32"),
		L("ubfx", 0xd3400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).
			Imm(Immr).Kind(ImmBitfieldLsb).Imm(Imms).Kind(ImmBitfieldWidth).
			AliasOf("ubfm").Name("Ubfx64"),
		L("sbfx", 0x13000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			Imm(Immr).Kind(ImmBitfieldLsb).Imm(Imms).Kind(ImmBitfieldWidth).
			AliasOf("sbfm").Name("Sbfx32"),
		L("sbfx", 0x93400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).
			Imm(Immr).Kind(ImmBitfieldLsb).Imm(Imms).Kind(ImmBitfieldWidth).
			AliasOf("sbfm").Name("Sbfx64"),

		// ROR (immediate) is EXTR with one source named twice: rotating a
		// register right is extracting from the pair it forms with itself.
		// The row copies Rn into Rm rather than the caller passing it, which
		// is what keeps the assembler and the typed helper on one word.
		L("ror", 0x13800000, 0x7fa00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).Imm(Imms).
			AliasOf("extr").Attr(AttrRnIntoRm).Name("RorImm32"),
		L("ror", 0x93c00000, 0xffe00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).Imm(Imms).
			AliasOf("extr").Attr(AttrRnIntoRm).Name("RorImm64"),

		L("cset", 0x9a9f07e0, 0xffff0fe0).
			Dst(ClassX, Rd).Cnd(CondHi).
			AliasOf("csinc").Pins(Rn, 31).Pins(Rm, 31).Attr(AttrInvertCond).Name("Cset64"),
		L("cset", 0x1a9f07e0, 0x7fff0fe0).
			Dst(ClassW, Rd).Cnd(CondHi).
			AliasOf("csinc").Pins(Rn, 31).Pins(Rm, 31).Attr(AttrInvertCond).Name("Cset32"),

		// CSETM is CSET's other half: all ones rather than one, which is a
		// mask and is what a branchless select is built from.
		L("csetm", 0xda9f03e0, 0xffff0fe0).
			Dst(ClassX, Rd).Cnd(CondHi).
			AliasOf("csinv").Pins(Rn, 31).Pins(Rm, 31).Attr(AttrInvertCond).Name("Csetm64"),
		L("csetm", 0x5a9f03e0, 0x7fff0fe0).
			Dst(ClassW, Rd).Cnd(CondHi).
			AliasOf("csinv").Pins(Rn, 31).Pins(Rm, 31).Attr(AttrInvertCond).Name("Csetm32"),

		// CINC, CINV and CNEG name one source twice, the way ROR does, and
		// invert their condition, the way CSET does. Both facts are on the
		// row for the same reason: an assembler reaching these through Emit
		// has to get the word the typed helper would.
		L("cinc", 0x1a800400, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Cnd(CondHi).
			AliasOf("csinc").Attr(AttrInvertCond).Attr(AttrRnIntoRm).Name("Cinc32"),
		L("cinc", 0x9a800400, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Cnd(CondHi).
			AliasOf("csinc").Attr(AttrInvertCond).Attr(AttrRnIntoRm).Name("Cinc64"),
		L("cinv", 0x5a800000, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Cnd(CondHi).
			AliasOf("csinv").Attr(AttrInvertCond).Attr(AttrRnIntoRm).Name("Cinv32"),
		L("cinv", 0xda800000, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Cnd(CondHi).
			AliasOf("csinv").Attr(AttrInvertCond).Attr(AttrRnIntoRm).Name("Cinv64"),
		L("cneg", 0x5a800400, 0x7fe00c00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Cnd(CondHi).
			AliasOf("csneg").Attr(AttrInvertCond).Attr(AttrRnIntoRm).Name("Cneg32"),
		L("cneg", 0xda800400, 0xffe00c00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Cnd(CondHi).
			AliasOf("csneg").Attr(AttrInvertCond).Attr(AttrRnIntoRm).Name("Cneg64"),

		// NEGS, NGC and NGCS are SUBS, SBC and SBCS with the zero register
		// as the first source, which is the same relation NEG has to SUB.
		L("negs", 0x6b0003e0, 0x7f2003e0).
			Dst(ClassW, Rd).Src(ClassW, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("subs").Pins(Rn, 31).Name("Negs32"),
		L("negs", 0xeb0003e0, 0xff2003e0).
			Dst(ClassX, Rd).Src(ClassX, Rm).Opt(ClassShift, Shift, 0).
			Flags().AliasOf("subs").Pins(Rn, 31).Name("Negs64"),
		L("ngc", 0x5a0003e0, 0x7fe0ffe0).
			Dst(ClassW, Rd).Src(ClassW, Rm).
			AliasOf("sbc").Pins(Rn, 31).Name("Ngc32"),
		L("ngc", 0xda0003e0, 0xffe0ffe0).
			Dst(ClassX, Rd).Src(ClassX, Rm).
			AliasOf("sbc").Pins(Rn, 31).Name("Ngc64"),
		L("ngcs", 0x7a0003e0, 0x7fe0ffe0).
			Dst(ClassW, Rd).Src(ClassW, Rm).
			Flags().AliasOf("sbcs").Pins(Rn, 31).Name("Ngcs32"),
		L("ngcs", 0xfa0003e0, 0xffe0ffe0).
			Dst(ClassX, Rd).Src(ClassX, Rm).
			Flags().AliasOf("sbcs").Pins(Rn, 31).Name("Ngcs64"),

		// The widening multiply's own aliases, with the accumulator pinned
		// to the zero register.
		L("smull", 0x9b207c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("smaddl").Pins(Ra, 31).Name("Smull"),
		L("umull", 0x9ba07c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("umaddl").Pins(Ra, 31).Name("Umull"),
		L("smnegl", 0x9b20fc00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("smsubl").Pins(Ra, 31).Name("Smnegl"),
		L("umnegl", 0x9ba0fc00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("umsubl").Pins(Ra, 31).Name("Umnegl"),

		// The insert half of the bitfield aliases. BFXIL extracts to the low
		// bits and shares UBFX's arithmetic; BFI, UBFIZ and SBFIZ place a
		// field at a position and rotate to get it there.
		L("bfxil", 0x33000000, 0x7fc00000).
			SrcDst(ClassW, Rd).Src(ClassW, Rn).
			Imm(Immr).Kind(ImmBitfieldLsb).Imm(Imms).Kind(ImmBitfieldWidth).
			AliasOf("bfm").Name("Bfxil32"),
		L("bfxil", 0xb3400000, 0xffc00000).
			SrcDst(ClassX, Rd).Src(ClassX, Rn).
			Imm(Immr).Kind(ImmBitfieldLsb).Imm(Imms).Kind(ImmBitfieldWidth).
			AliasOf("bfm").Name("Bfxil64"),
		L("bfi", 0x33000000, 0x7fc00000).
			SrcDst(ClassW, Rd).Src(ClassW, Rn).
			Imm(Immr).Kind(ImmBitfieldLsbNeg).Imm(Imms).Kind(ImmBitfieldWidthM1).
			AliasOf("bfm").Name("Bfi32"),
		L("bfi", 0xb3400000, 0xffc00000).
			SrcDst(ClassX, Rd).Src(ClassX, Rn).
			Imm(Immr).Kind(ImmBitfieldLsbNeg).Imm(Imms).Kind(ImmBitfieldWidthM1).
			AliasOf("bfm").Name("Bfi64"),
		L("ubfiz", 0x53000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			Imm(Immr).Kind(ImmBitfieldLsbNeg).Imm(Imms).Kind(ImmBitfieldWidthM1).
			AliasOf("ubfm").Name("Ubfiz32"),
		L("ubfiz", 0xd3400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).
			Imm(Immr).Kind(ImmBitfieldLsbNeg).Imm(Imms).Kind(ImmBitfieldWidthM1).
			AliasOf("ubfm").Name("Ubfiz64"),
		L("sbfiz", 0x13000000, 0x7fc00000).
			Dst(ClassW, Rd).Src(ClassW, Rn).
			Imm(Immr).Kind(ImmBitfieldLsbNeg).Imm(Imms).Kind(ImmBitfieldWidthM1).
			AliasOf("sbfm").Name("Sbfiz32"),
		L("sbfiz", 0x93400000, 0xffc00000).
			Dst(ClassX, Rd).Src(ClassX, Rn).
			Imm(Immr).Kind(ImmBitfieldLsbNeg).Imm(Imms).Kind(ImmBitfieldWidthM1).
			AliasOf("sbfm").Name("Sbfiz64"),
		L("mul", 0x1b007c00, 0x7fe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).
			AliasOf("madd").Pins(Ra, 31).Name("Mul32"),
		L("mul", 0x9b007c00, 0xffe0fc00).
			Dst(ClassX, Rd).Src(ClassX, Rn).Src(ClassX, Rm).
			AliasOf("madd").Pins(Ra, 31).Name("Mul64"),

		// .inst states a word rather than naming an instruction. It is the one
		// case where emitting bytes nobody selected is what was asked for.
		L(".inst", 0, 0).Imm(F(0, 32)).Name("Inst"),
	)

}
