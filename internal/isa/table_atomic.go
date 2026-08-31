package isa

// The atomic instruction set: the load/store-exclusive pairs, the
// acquire-release ordered accesses, and the LSE single-instruction family.
//
// The exclusive pairs and LDAR/STLR are Armv8-A baseline and are sufficient on
// their own — every atomic read-modify-write is a retry loop around LDAXR and
// STLXR. The LSE forms are one instruction for what that loop does and are
// gated on FEAT_LSE, mandatory from Armv8.1-A.
//
// None of these takes an offset. The exclusive and LSE encodings have no
// immediate field at all, so the address is [Xn|SP] and nothing else, which is
// why every Mem here states a zero offset field.
import "github.com/vertex-language/arm64/feature"

func init() {
	register(
		// ---- Ordered load and store ----
		//
		// LDAR is acquire and STLR is release. No exclusive monitor is
		// involved: these order an ordinary access, where the pairs below
		// make one atomic.
		L("ldarb", 0x08dffc00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(8, Rn, Field{}).Name("Ldarb"),
		L("ldarh", 0x48dffc00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(16, Rn, Field{}).Name("Ldarh"),
		L("ldar", 0x88dffc00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(32, Rn, Field{}).Name("Ldar32"),
		L("ldar", 0xc8dffc00, 0xfffffc00).
			Dst(ClassX, Rt).Mem(64, Rn, Field{}).Name("Ldar64"),
		L("stlrb", 0x089ffc00, 0xfffffc00).
			Src(ClassW, Rt).Mem(8, Rn, Field{}).Name("Stlrb"),
		L("stlrh", 0x489ffc00, 0xfffffc00).
			Src(ClassW, Rt).Mem(16, Rn, Field{}).Name("Stlrh"),
		L("stlr", 0x889ffc00, 0xfffffc00).
			Src(ClassW, Rt).Mem(32, Rn, Field{}).Name("Stlr32"),
		L("stlr", 0xc89ffc00, 0xfffffc00).
			Src(ClassX, Rt).Mem(64, Rn, Field{}).Name("Stlr64"),

		// ---- Load and store exclusive ----
		//
		// STXR's first destination is the status register: zero if the store
		// succeeded, one if the monitor was lost and the loop must retry.
		L("ldxrb", 0x085f7c00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(8, Rn, Field{}).Name("Ldxrb"),
		L("ldxrh", 0x485f7c00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(16, Rn, Field{}).Name("Ldxrh"),
		L("ldxr", 0x885f7c00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(32, Rn, Field{}).Name("Ldxr32"),
		L("ldxr", 0xc85f7c00, 0xfffffc00).
			Dst(ClassX, Rt).Mem(64, Rn, Field{}).Name("Ldxr64"),
		L("ldaxrb", 0x085ffc00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(8, Rn, Field{}).Name("Ldaxrb"),
		L("ldaxrh", 0x485ffc00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(16, Rn, Field{}).Name("Ldaxrh"),
		L("ldaxr", 0x885ffc00, 0xfffffc00).
			Dst(ClassW, Rt).Mem(32, Rn, Field{}).Name("Ldaxr32"),
		L("ldaxr", 0xc85ffc00, 0xfffffc00).
			Dst(ClassX, Rt).Mem(64, Rn, Field{}).Name("Ldaxr64"),

		L("stxrb", 0x08007c00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(8, Rn, Field{}).Name("Stxrb"),
		L("stxrh", 0x48007c00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(16, Rn, Field{}).Name("Stxrh"),
		L("stxr", 0x88007c00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(32, Rn, Field{}).Name("Stxr32"),
		L("stxr", 0xc8007c00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassX, Rt).Mem(64, Rn, Field{}).Name("Stxr64"),
		L("stlxrb", 0x0800fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(8, Rn, Field{}).Name("Stlxrb"),
		L("stlxrh", 0x4800fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(16, Rn, Field{}).Name("Stlxrh"),
		L("stlxr", 0x8800fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(32, Rn, Field{}).Name("Stlxr32"),
		L("stlxr", 0xc800fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassX, Rt).Mem(64, Rn, Field{}).Name("Stlxr64"),

		// ---- LSE: one instruction for the loop above ----
		//
		// The acquire-release variants only. A relaxed atomic is expressible
		// and is not here: what the IR above this needs is sequential
		// consistency, and a row nothing selects is a row nothing tests.
		//
		// Rs is the operand and Rt the old value the instruction returns,
		// which is why both are named on what reads as a store.
		L("ldaddalb", 0x38e00000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(8, Rn, Field{}).Gate2(feature.LSE).Name("Ldaddalb"),
		L("ldaddalh", 0x78e00000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(16, Rn, Field{}).Gate2(feature.LSE).Name("Ldaddalh"),
		L("ldaddal", 0xb8e00000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(32, Rn, Field{}).Gate2(feature.LSE).Name("Ldaddal32"),
		L("ldaddal", 0xf8e00000, 0xffe0fc00).
			Src(ClassX, Rs).Dst(ClassX, Rt).Mem(64, Rn, Field{}).Gate2(feature.LSE).Name("Ldaddal64"),

		// LDCLR clears the bits its operand sets, so an atomic AND is a
		// LDCLR of the complement rather than of the mask.
		L("ldclralb", 0x38e01000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(8, Rn, Field{}).Gate2(feature.LSE).Name("Ldclralb"),
		L("ldclralh", 0x78e01000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(16, Rn, Field{}).Gate2(feature.LSE).Name("Ldclralh"),
		L("ldclral", 0xb8e01000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(32, Rn, Field{}).Gate2(feature.LSE).Name("Ldclral32"),
		L("ldclral", 0xf8e01000, 0xffe0fc00).
			Src(ClassX, Rs).Dst(ClassX, Rt).Mem(64, Rn, Field{}).Gate2(feature.LSE).Name("Ldclral64"),

		L("ldeoralb", 0x38e02000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(8, Rn, Field{}).Gate2(feature.LSE).Name("Ldeoralb"),
		L("ldeoralh", 0x78e02000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(16, Rn, Field{}).Gate2(feature.LSE).Name("Ldeoralh"),
		L("ldeoral", 0xb8e02000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(32, Rn, Field{}).Gate2(feature.LSE).Name("Ldeoral32"),
		L("ldeoral", 0xf8e02000, 0xffe0fc00).
			Src(ClassX, Rs).Dst(ClassX, Rt).Mem(64, Rn, Field{}).Gate2(feature.LSE).Name("Ldeoral64"),

		L("ldsetalb", 0x38e03000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(8, Rn, Field{}).Gate2(feature.LSE).Name("Ldsetalb"),
		L("ldsetalh", 0x78e03000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(16, Rn, Field{}).Gate2(feature.LSE).Name("Ldsetalh"),
		L("ldsetal", 0xb8e03000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(32, Rn, Field{}).Gate2(feature.LSE).Name("Ldsetal32"),
		L("ldsetal", 0xf8e03000, 0xffe0fc00).
			Src(ClassX, Rs).Dst(ClassX, Rt).Mem(64, Rn, Field{}).Gate2(feature.LSE).Name("Ldsetal64"),

		L("swpalb", 0x38e08000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(8, Rn, Field{}).Gate2(feature.LSE).Name("Swpalb"),
		L("swpalh", 0x78e08000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(16, Rn, Field{}).Gate2(feature.LSE).Name("Swpalh"),
		L("swpal", 0xb8e08000, 0xffe0fc00).
			Src(ClassW, Rs).Dst(ClassW, Rt).Mem(32, Rn, Field{}).Gate2(feature.LSE).Name("Swpal32"),
		L("swpal", 0xf8e08000, 0xffe0fc00).
			Src(ClassX, Rs).Dst(ClassX, Rt).Mem(64, Rn, Field{}).Gate2(feature.LSE).Name("Swpal64"),

		// CAS reads its expected value out of Rs and writes the old value
		// back into it, so the operand is both. Rt is the new value.
		L("casalb", 0x08e0fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(8, Rn, Field{}).Gate2(feature.LSE).Name("Casalb"),
		L("casalh", 0x48e0fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(16, Rn, Field{}).Gate2(feature.LSE).Name("Casalh"),
		L("casal", 0x88e0fc00, 0xffe0fc00).
			Dst(ClassW, Rs).Src(ClassW, Rt).Mem(32, Rn, Field{}).Gate2(feature.LSE).Name("Casal32"),
		L("casal", 0xc8e0fc00, 0xffe0fc00).
			Dst(ClassX, Rs).Src(ClassX, Rt).Mem(64, Rn, Field{}).Gate2(feature.LSE).Name("Casal64"),
	)
}
