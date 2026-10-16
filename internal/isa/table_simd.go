package isa

import (
	"github.com/vertex-language/arm64/feature"
	"github.com/vertex-language/arm64/reg"
)

// Advanced SIMD's integer instructions and the cryptographic extensions.
// Generated from one list with inst_simd.go and simd_difftest_test.go, so
// a row, its typed helper and the clang case that checks its word cannot
// disagree; edit the list in gen_simd.py, not the output, and regenerate
// with "python3 internal/isa/gen_simd.py /tmp/out" and copying the three
// files into place.
//
// A vector operand names its arrangement: the size and Q bits each row
// pins are what tell v0.4s from v0.8h, so a slot takes the one its row
// encodes. Advanced SIMD is in the baseline; the cryptographic rows are
// gated on their features, PMULL's 64-bit form on AES as the architecture
// has it.
func init() {
	register(

		// ---- Integer arithmetic, three registers of one arrangement ----
		L("add", 0x0e208400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("AddV8B"),
		L("add", 0x4e208400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("AddV16B"),
		L("add", 0x0e608400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("AddV4H"),
		L("add", 0x4e608400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("AddV8H"),
		L("add", 0x0ea08400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("AddV2S"),
		L("add", 0x4ea08400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("AddV4S"),
		L("add", 0x4ee08400, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("AddV2D"),
		L("sub", 0x2e208400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SubV8B"),
		L("sub", 0x6e208400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SubV16B"),
		L("sub", 0x2e608400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SubV4H"),
		L("sub", 0x6e608400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SubV8H"),
		L("sub", 0x2ea08400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SubV2S"),
		L("sub", 0x6ea08400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SubV4S"),
		L("sub", 0x6ee08400, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("SubV2D"),
		L("mul", 0x0e209c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("MulV8B"),
		L("mul", 0x4e209c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("MulV16B"),
		L("mul", 0x0e609c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("MulV4H"),
		L("mul", 0x4e609c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("MulV8H"),
		L("mul", 0x0ea09c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("MulV2S"),
		L("mul", 0x4ea09c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("MulV4S"),
		L("mla", 0x0e209400, 0xffe0fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("MlaV8B"),
		L("mla", 0x4e209400, 0xffe0fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("MlaV16B"),
		L("mla", 0x0e609400, 0xffe0fc00).
			SrcDstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("MlaV4H"),
		L("mla", 0x4e609400, 0xffe0fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("MlaV8H"),
		L("mla", 0x0ea09400, 0xffe0fc00).
			SrcDstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("MlaV2S"),
		L("mla", 0x4ea09400, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("MlaV4S"),
		L("mls", 0x2e209400, 0xffe0fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("MlsV8B"),
		L("mls", 0x6e209400, 0xffe0fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("MlsV16B"),
		L("mls", 0x2e609400, 0xffe0fc00).
			SrcDstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("MlsV4H"),
		L("mls", 0x6e609400, 0xffe0fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("MlsV8H"),
		L("mls", 0x2ea09400, 0xffe0fc00).
			SrcDstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("MlsV2S"),
		L("mls", 0x6ea09400, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("MlsV4S"),
		L("cmeq", 0x2e208c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("CmeqV8B"),
		L("cmeq", 0x6e208c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("CmeqV16B"),
		L("cmeq", 0x2e608c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("CmeqV4H"),
		L("cmeq", 0x6e608c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("CmeqV8H"),
		L("cmeq", 0x2ea08c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("CmeqV2S"),
		L("cmeq", 0x6ea08c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("CmeqV4S"),
		L("cmeq", 0x6ee08c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("CmeqV2D"),
		L("cmtst", 0x0e208c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("CmtstV8B"),
		L("cmtst", 0x4e208c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("CmtstV16B"),
		L("cmtst", 0x0e608c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("CmtstV4H"),
		L("cmtst", 0x4e608c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("CmtstV8H"),
		L("cmtst", 0x0ea08c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("CmtstV2S"),
		L("cmtst", 0x4ea08c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("CmtstV4S"),
		L("cmtst", 0x4ee08c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("CmtstV2D"),
		L("cmgt", 0x0e203400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("CmgtV8B"),
		L("cmgt", 0x4e203400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("CmgtV16B"),
		L("cmgt", 0x0e603400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("CmgtV4H"),
		L("cmgt", 0x4e603400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("CmgtV8H"),
		L("cmgt", 0x0ea03400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("CmgtV2S"),
		L("cmgt", 0x4ea03400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("CmgtV4S"),
		L("cmgt", 0x4ee03400, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("CmgtV2D"),
		L("cmge", 0x0e203c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("CmgeV8B"),
		L("cmge", 0x4e203c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("CmgeV16B"),
		L("cmge", 0x0e603c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("CmgeV4H"),
		L("cmge", 0x4e603c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("CmgeV8H"),
		L("cmge", 0x0ea03c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("CmgeV2S"),
		L("cmge", 0x4ea03c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("CmgeV4S"),
		L("cmge", 0x4ee03c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("CmgeV2D"),
		L("cmhi", 0x2e203400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("CmhiV8B"),
		L("cmhi", 0x6e203400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("CmhiV16B"),
		L("cmhi", 0x2e603400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("CmhiV4H"),
		L("cmhi", 0x6e603400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("CmhiV8H"),
		L("cmhi", 0x2ea03400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("CmhiV2S"),
		L("cmhi", 0x6ea03400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("CmhiV4S"),
		L("cmhi", 0x6ee03400, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("CmhiV2D"),
		L("cmhs", 0x2e203c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("CmhsV8B"),
		L("cmhs", 0x6e203c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("CmhsV16B"),
		L("cmhs", 0x2e603c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("CmhsV4H"),
		L("cmhs", 0x6e603c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("CmhsV8H"),
		L("cmhs", 0x2ea03c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("CmhsV2S"),
		L("cmhs", 0x6ea03c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("CmhsV4S"),
		L("cmhs", 0x6ee03c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("CmhsV2D"),
		L("sqadd", 0x0e200c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SqaddV8B"),
		L("sqadd", 0x4e200c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SqaddV16B"),
		L("sqadd", 0x0e600c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SqaddV4H"),
		L("sqadd", 0x4e600c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SqaddV8H"),
		L("sqadd", 0x0ea00c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SqaddV2S"),
		L("sqadd", 0x4ea00c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SqaddV4S"),
		L("sqadd", 0x4ee00c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("SqaddV2D"),
		L("uqadd", 0x2e200c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UqaddV8B"),
		L("uqadd", 0x6e200c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UqaddV16B"),
		L("uqadd", 0x2e600c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UqaddV4H"),
		L("uqadd", 0x6e600c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UqaddV8H"),
		L("uqadd", 0x2ea00c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UqaddV2S"),
		L("uqadd", 0x6ea00c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UqaddV4S"),
		L("uqadd", 0x6ee00c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("UqaddV2D"),
		L("sqsub", 0x0e202c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SqsubV8B"),
		L("sqsub", 0x4e202c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SqsubV16B"),
		L("sqsub", 0x0e602c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SqsubV4H"),
		L("sqsub", 0x4e602c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SqsubV8H"),
		L("sqsub", 0x0ea02c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SqsubV2S"),
		L("sqsub", 0x4ea02c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SqsubV4S"),
		L("sqsub", 0x4ee02c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("SqsubV2D"),
		L("uqsub", 0x2e202c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UqsubV8B"),
		L("uqsub", 0x6e202c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UqsubV16B"),
		L("uqsub", 0x2e602c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UqsubV4H"),
		L("uqsub", 0x6e602c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UqsubV8H"),
		L("uqsub", 0x2ea02c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UqsubV2S"),
		L("uqsub", 0x6ea02c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UqsubV4S"),
		L("uqsub", 0x6ee02c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("UqsubV2D"),
		L("smax", 0x0e206400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SmaxV8B"),
		L("smax", 0x4e206400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SmaxV16B"),
		L("smax", 0x0e606400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SmaxV4H"),
		L("smax", 0x4e606400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SmaxV8H"),
		L("smax", 0x0ea06400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SmaxV2S"),
		L("smax", 0x4ea06400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SmaxV4S"),
		L("umax", 0x2e206400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UmaxV8B"),
		L("umax", 0x6e206400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UmaxV16B"),
		L("umax", 0x2e606400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UmaxV4H"),
		L("umax", 0x6e606400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UmaxV8H"),
		L("umax", 0x2ea06400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UmaxV2S"),
		L("umax", 0x6ea06400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UmaxV4S"),
		L("smin", 0x0e206c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SminV8B"),
		L("smin", 0x4e206c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SminV16B"),
		L("smin", 0x0e606c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SminV4H"),
		L("smin", 0x4e606c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SminV8H"),
		L("smin", 0x0ea06c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SminV2S"),
		L("smin", 0x4ea06c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SminV4S"),
		L("umin", 0x2e206c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UminV8B"),
		L("umin", 0x6e206c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UminV16B"),
		L("umin", 0x2e606c00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UminV4H"),
		L("umin", 0x6e606c00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UminV8H"),
		L("umin", 0x2ea06c00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UminV2S"),
		L("umin", 0x6ea06c00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UminV4S"),
		L("srhadd", 0x0e201400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SrhaddV8B"),
		L("srhadd", 0x4e201400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SrhaddV16B"),
		L("srhadd", 0x0e601400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SrhaddV4H"),
		L("srhadd", 0x4e601400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SrhaddV8H"),
		L("srhadd", 0x0ea01400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SrhaddV2S"),
		L("srhadd", 0x4ea01400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SrhaddV4S"),
		L("urhadd", 0x2e201400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UrhaddV8B"),
		L("urhadd", 0x6e201400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UrhaddV16B"),
		L("urhadd", 0x2e601400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UrhaddV4H"),
		L("urhadd", 0x6e601400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UrhaddV8H"),
		L("urhadd", 0x2ea01400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UrhaddV2S"),
		L("urhadd", 0x6ea01400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UrhaddV4S"),
		L("sabd", 0x0e207400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SabdV8B"),
		L("sabd", 0x4e207400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SabdV16B"),
		L("sabd", 0x0e607400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SabdV4H"),
		L("sabd", 0x4e607400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SabdV8H"),
		L("sabd", 0x0ea07400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SabdV2S"),
		L("sabd", 0x4ea07400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SabdV4S"),
		L("uabd", 0x2e207400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UabdV8B"),
		L("uabd", 0x6e207400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UabdV16B"),
		L("uabd", 0x2e607400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UabdV4H"),
		L("uabd", 0x6e607400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UabdV8H"),
		L("uabd", 0x2ea07400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UabdV2S"),
		L("uabd", 0x6ea07400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UabdV4S"),
		L("sshl", 0x0e204400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SshlV8B"),
		L("sshl", 0x4e204400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("SshlV16B"),
		L("sshl", 0x0e604400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SshlV4H"),
		L("sshl", 0x4e604400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SshlV8H"),
		L("sshl", 0x0ea04400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SshlV2S"),
		L("sshl", 0x4ea04400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SshlV4S"),
		L("sshl", 0x4ee04400, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("SshlV2D"),
		L("ushl", 0x2e204400, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UshlV8B"),
		L("ushl", 0x6e204400, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("UshlV16B"),
		L("ushl", 0x2e604400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UshlV4H"),
		L("ushl", 0x6e604400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("UshlV8H"),
		L("ushl", 0x2ea04400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UshlV2S"),
		L("ushl", 0x6ea04400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("UshlV4S"),
		L("ushl", 0x6ee04400, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("UshlV2D"),
		L("addp", 0x0e20bc00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("AddpV8B"),
		L("addp", 0x4e20bc00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("AddpV16B"),
		L("addp", 0x0e60bc00, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("AddpV4H"),
		L("addp", 0x4e60bc00, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("AddpV8H"),
		L("addp", 0x0ea0bc00, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("AddpV2S"),
		L("addp", 0x4ea0bc00, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("AddpV4S"),
		L("addp", 0x4ee0bc00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("AddpV2D"),
		L("sqdmulh", 0x0e60b400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SqdmulhV4H"),
		L("sqdmulh", 0x4e60b400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SqdmulhV8H"),
		L("sqdmulh", 0x0ea0b400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SqdmulhV2S"),
		L("sqdmulh", 0x4ea0b400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SqdmulhV4S"),
		L("sqrdmulh", 0x2e60b400, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SqrdmulhV4H"),
		L("sqrdmulh", 0x6e60b400, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("SqrdmulhV8H"),
		L("sqrdmulh", 0x2ea0b400, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SqrdmulhV2S"),
		L("sqrdmulh", 0x6ea0b400, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("SqrdmulhV4S"),

		// ---- Bitwise, three registers ----
		L("and", 0x0e201c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("AndV8B"),
		L("and", 0x4e201c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("AndV16B"),
		L("bic", 0x0e601c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("BicV8B"),
		L("bic", 0x4e601c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("BicV16B"),
		L("orr", 0x0ea01c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("OrrV8B"),
		L("orr", 0x4ea01c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("OrrV16B"),
		L("orn", 0x0ee01c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("OrnV8B"),
		L("orn", 0x4ee01c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("OrnV16B"),
		L("eor", 0x2e201c00, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("EorV8B"),
		L("eor", 0x6e201c00, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("EorV16B"),
		L("bsl", 0x2e601c00, 0xffe0fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("BslV8B"),
		L("bsl", 0x6e601c00, 0xffe0fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("BslV16B"),
		L("bit", 0x2ea01c00, 0xffe0fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("BitV8B"),
		L("bit", 0x6ea01c00, 0xffe0fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("BitV16B"),
		L("bif", 0x2ee01c00, 0xffe0fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("BifV8B"),
		L("bif", 0x6ee01c00, 0xffe0fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("BifV16B"),

		// ---- One source register ----
		L("rev64", 0x0e200800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("Rev64V8B"),
		L("rev64", 0x4e200800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("Rev64V16B"),
		L("rev64", 0x0e600800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Name("Rev64V4H"),
		L("rev64", 0x4e600800, 0xfffffc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Name("Rev64V8H"),
		L("rev64", 0x0ea00800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Name("Rev64V2S"),
		L("rev64", 0x4ea00800, 0xfffffc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Name("Rev64V4S"),
		L("rev32", 0x2e200800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("Rev32V8B"),
		L("rev32", 0x6e200800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("Rev32V16B"),
		L("rev32", 0x2e600800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Name("Rev32V4H"),
		L("rev32", 0x6e600800, 0xfffffc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Name("Rev32V8H"),
		L("rev16", 0x0e201800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("Rev16V8B"),
		L("rev16", 0x4e201800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("Rev16V16B"),
		L("cnt", 0x0e205800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("CntV8B"),
		L("cnt", 0x4e205800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("CntV16B"),
		L("not", 0x2e205800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("NotV8B"),
		L("not", 0x6e205800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("NotV16B"),
		L("rbit", 0x2e605800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("RbitV8B"),
		L("rbit", 0x6e605800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("RbitV16B"),
		L("neg", 0x2e20b800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("NegV8B"),
		L("neg", 0x6e20b800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("NegV16B"),
		L("neg", 0x2e60b800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Name("NegV4H"),
		L("neg", 0x6e60b800, 0xfffffc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Name("NegV8H"),
		L("neg", 0x2ea0b800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Name("NegV2S"),
		L("neg", 0x6ea0b800, 0xfffffc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Name("NegV4S"),
		L("neg", 0x6ee0b800, 0xfffffc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Name("NegV2D"),
		L("abs", 0x0e20b800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Name("AbsV8B"),
		L("abs", 0x4e20b800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Name("AbsV16B"),
		L("abs", 0x0e60b800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Name("AbsV4H"),
		L("abs", 0x4e60b800, 0xfffffc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Name("AbsV8H"),
		L("abs", 0x0ea0b800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Name("AbsV2S"),
		L("abs", 0x4ea0b800, 0xfffffc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Name("AbsV4S"),
		L("abs", 0x4ee0b800, 0xfffffc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Name("AbsV2D"),

		// ---- Narrowing and pairwise widening ----
		L("xtn", 0x0e212800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Name("XtnV8B"),
		L("xtn2", 0x4e212800, 0xfffffc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Name("Xtn2V16B"),
		L("xtn", 0x0e612800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Name("XtnV4H"),
		L("xtn2", 0x4e612800, 0xfffffc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Name("Xtn2V8H"),
		L("xtn", 0x0ea12800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Name("XtnV2S"),
		L("xtn2", 0x4ea12800, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Name("Xtn2V4S"),
		L("sqxtn", 0x0e214800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Name("SqxtnV8B"),
		L("sqxtn2", 0x4e214800, 0xfffffc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Name("Sqxtn2V16B"),
		L("sqxtn", 0x0e614800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Name("SqxtnV4H"),
		L("sqxtn2", 0x4e614800, 0xfffffc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Name("Sqxtn2V8H"),
		L("sqxtn", 0x0ea14800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Name("SqxtnV2S"),
		L("sqxtn2", 0x4ea14800, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Name("Sqxtn2V4S"),
		L("uqxtn", 0x2e214800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Name("UqxtnV8B"),
		L("uqxtn2", 0x6e214800, 0xfffffc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Name("Uqxtn2V16B"),
		L("uqxtn", 0x2e614800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Name("UqxtnV4H"),
		L("uqxtn2", 0x6e614800, 0xfffffc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Name("Uqxtn2V8H"),
		L("uqxtn", 0x2ea14800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Name("UqxtnV2S"),
		L("uqxtn2", 0x6ea14800, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Name("Uqxtn2V4S"),
		L("sqxtun", 0x2e212800, 0xfffffc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Name("SqxtunV8B"),
		L("sqxtun2", 0x6e212800, 0xfffffc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Name("Sqxtun2V16B"),
		L("sqxtun", 0x2e612800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Name("SqxtunV4H"),
		L("sqxtun2", 0x6e612800, 0xfffffc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Name("Sqxtun2V8H"),
		L("sqxtun", 0x2ea12800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Name("SqxtunV2S"),
		L("sqxtun2", 0x6ea12800, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Name("Sqxtun2V4S"),
		L("uaddlp", 0x2e202800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V8B, Rn).Name("UaddlpV4H"),
		L("uaddlp", 0x6e202800, 0xfffffc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).Name("UaddlpV8H"),
		L("uaddlp", 0x2e602800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V4H, Rn).Name("UaddlpV2S"),
		L("uaddlp", 0x6e602800, 0xfffffc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).Name("UaddlpV4S"),
		L("uaddlp", 0x2ea02800, 0xfffffc00).
			DstV(reg.V1D, Rd).SrcV(reg.V2S, Rn).Name("UaddlpV1D"),
		L("uaddlp", 0x6ea02800, 0xfffffc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).Name("UaddlpV2D"),
		L("saddlp", 0x0e202800, 0xfffffc00).
			DstV(reg.V4H, Rd).SrcV(reg.V8B, Rn).Name("SaddlpV4H"),
		L("saddlp", 0x4e202800, 0xfffffc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).Name("SaddlpV8H"),
		L("saddlp", 0x0e602800, 0xfffffc00).
			DstV(reg.V2S, Rd).SrcV(reg.V4H, Rn).Name("SaddlpV2S"),
		L("saddlp", 0x4e602800, 0xfffffc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).Name("SaddlpV4S"),
		L("saddlp", 0x0ea02800, 0xfffffc00).
			DstV(reg.V1D, Rd).SrcV(reg.V2S, Rn).Name("SaddlpV1D"),
		L("saddlp", 0x4ea02800, 0xfffffc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).Name("SaddlpV2D"),

		// ---- Across all lanes, into a scalar ----
		L("addv", 0x0e31b800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V8B, Rn).Name("AddvV8B"),
		L("addv", 0x4e31b800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V16B, Rn).Name("AddvV16B"),
		L("addv", 0x0e71b800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V4H, Rn).Name("AddvV4H"),
		L("addv", 0x4e71b800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8H, Rn).Name("AddvV8H"),
		L("addv", 0x4eb1b800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4S, Rn).Name("AddvV4S"),
		L("umaxv", 0x2e30a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V8B, Rn).Name("UmaxvV8B"),
		L("umaxv", 0x6e30a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V16B, Rn).Name("UmaxvV16B"),
		L("umaxv", 0x2e70a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V4H, Rn).Name("UmaxvV4H"),
		L("umaxv", 0x6e70a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8H, Rn).Name("UmaxvV8H"),
		L("umaxv", 0x6eb0a800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4S, Rn).Name("UmaxvV4S"),
		L("uminv", 0x2e31a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V8B, Rn).Name("UminvV8B"),
		L("uminv", 0x6e31a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V16B, Rn).Name("UminvV16B"),
		L("uminv", 0x2e71a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V4H, Rn).Name("UminvV4H"),
		L("uminv", 0x6e71a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8H, Rn).Name("UminvV8H"),
		L("uminv", 0x6eb1a800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4S, Rn).Name("UminvV4S"),
		L("smaxv", 0x0e30a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V8B, Rn).Name("SmaxvV8B"),
		L("smaxv", 0x4e30a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V16B, Rn).Name("SmaxvV16B"),
		L("smaxv", 0x0e70a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V4H, Rn).Name("SmaxvV4H"),
		L("smaxv", 0x4e70a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8H, Rn).Name("SmaxvV8H"),
		L("smaxv", 0x4eb0a800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4S, Rn).Name("SmaxvV4S"),
		L("sminv", 0x0e31a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V8B, Rn).Name("SminvV8B"),
		L("sminv", 0x4e31a800, 0xfffffc00).
			Dst(ClassB, Rd).SrcV(reg.V16B, Rn).Name("SminvV16B"),
		L("sminv", 0x0e71a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V4H, Rn).Name("SminvV4H"),
		L("sminv", 0x4e71a800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8H, Rn).Name("SminvV8H"),
		L("sminv", 0x4eb1a800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4S, Rn).Name("SminvV4S"),
		L("uaddlv", 0x2e303800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8B, Rn).Name("UaddlvV8B"),
		L("uaddlv", 0x6e303800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V16B, Rn).Name("UaddlvV16B"),
		L("uaddlv", 0x2e703800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4H, Rn).Name("UaddlvV4H"),
		L("uaddlv", 0x6e703800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V8H, Rn).Name("UaddlvV8H"),
		L("uaddlv", 0x6eb03800, 0xfffffc00).
			Dst(ClassD, Rd).SrcV(reg.V4S, Rn).Name("UaddlvV4S"),
		L("saddlv", 0x0e303800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V8B, Rn).Name("SaddlvV8B"),
		L("saddlv", 0x4e303800, 0xfffffc00).
			Dst(ClassH, Rd).SrcV(reg.V16B, Rn).Name("SaddlvV16B"),
		L("saddlv", 0x0e703800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V4H, Rn).Name("SaddlvV4H"),
		L("saddlv", 0x4e703800, 0xfffffc00).
			Dst(ClassS, Rd).SrcV(reg.V8H, Rn).Name("SaddlvV8H"),
		L("saddlv", 0x4eb03800, 0xfffffc00).
			Dst(ClassD, Rd).SrcV(reg.V4S, Rn).Name("SaddlvV4S"),

		// ---- Shifts by an immediate ----
		L("shl", 0x0f005400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV8B"),
		L("shl", 0x4f005400, 0xff80fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV16B"),
		L("shl", 0x0f005400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV4H"),
		L("shl", 0x4f005400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV8H"),
		L("shl", 0x0f005400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV2S"),
		L("shl", 0x4f005400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV4S"),
		L("shl", 0x4f005400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("ShlV2D"),
		L("sli", 0x2f005400, 0xff80fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV8B"),
		L("sli", 0x6f005400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV16B"),
		L("sli", 0x2f005400, 0xff80fc00).
			SrcDstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV4H"),
		L("sli", 0x6f005400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV8H"),
		L("sli", 0x2f005400, 0xff80fc00).
			SrcDstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV2S"),
		L("sli", 0x6f005400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV4S"),
		L("sli", 0x6f005400, 0xff80fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SliV2D"),
		L("ushr", 0x2f000400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV8B"),
		L("ushr", 0x6f000400, 0xff80fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV16B"),
		L("ushr", 0x2f000400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV4H"),
		L("ushr", 0x6f000400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV8H"),
		L("ushr", 0x2f000400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV2S"),
		L("ushr", 0x6f000400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV4S"),
		L("ushr", 0x6f000400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UshrV2D"),
		L("sshr", 0x0f000400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV8B"),
		L("sshr", 0x4f000400, 0xff80fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV16B"),
		L("sshr", 0x0f000400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV4H"),
		L("sshr", 0x4f000400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV8H"),
		L("sshr", 0x0f000400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV2S"),
		L("sshr", 0x4f000400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV4S"),
		L("sshr", 0x4f000400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SshrV2D"),
		L("usra", 0x2f001400, 0xff80fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV8B"),
		L("usra", 0x6f001400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV16B"),
		L("usra", 0x2f001400, 0xff80fc00).
			SrcDstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV4H"),
		L("usra", 0x6f001400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV8H"),
		L("usra", 0x2f001400, 0xff80fc00).
			SrcDstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV2S"),
		L("usra", 0x6f001400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV4S"),
		L("usra", 0x6f001400, 0xff80fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UsraV2D"),
		L("ssra", 0x0f001400, 0xff80fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV8B"),
		L("ssra", 0x4f001400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV16B"),
		L("ssra", 0x0f001400, 0xff80fc00).
			SrcDstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV4H"),
		L("ssra", 0x4f001400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV8H"),
		L("ssra", 0x0f001400, 0xff80fc00).
			SrcDstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV2S"),
		L("ssra", 0x4f001400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV4S"),
		L("ssra", 0x4f001400, 0xff80fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SsraV2D"),
		L("urshr", 0x2f002400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV8B"),
		L("urshr", 0x6f002400, 0xff80fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV16B"),
		L("urshr", 0x2f002400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV4H"),
		L("urshr", 0x6f002400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV8H"),
		L("urshr", 0x2f002400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV2S"),
		L("urshr", 0x6f002400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV4S"),
		L("urshr", 0x6f002400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UrshrV2D"),
		L("srshr", 0x0f002400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV8B"),
		L("srshr", 0x4f002400, 0xff80fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV16B"),
		L("srshr", 0x0f002400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV4H"),
		L("srshr", 0x4f002400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV8H"),
		L("srshr", 0x0f002400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV2S"),
		L("srshr", 0x4f002400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV4S"),
		L("srshr", 0x4f002400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SrshrV2D"),
		L("sri", 0x2f004400, 0xff80fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV8B"),
		L("sri", 0x6f004400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV16B"),
		L("sri", 0x2f004400, 0xff80fc00).
			SrcDstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV4H"),
		L("sri", 0x6f004400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV8H"),
		L("sri", 0x2f004400, 0xff80fc00).
			SrcDstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV2S"),
		L("sri", 0x6f004400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV4S"),
		L("sri", 0x6f004400, 0xff80fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SriV2D"),
		L("shrn", 0x0f008400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("ShrnV8B"),
		L("shrn2", 0x4f008400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Shrn2V16B"),
		L("shrn", 0x0f008400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("ShrnV4H"),
		L("shrn2", 0x4f008400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Shrn2V8H"),
		L("shrn", 0x0f008400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("ShrnV2S"),
		L("shrn2", 0x4f008400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Shrn2V4S"),
		L("rshrn", 0x0f008c00, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("RshrnV8B"),
		L("rshrn2", 0x4f008c00, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Rshrn2V16B"),
		L("rshrn", 0x0f008c00, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("RshrnV4H"),
		L("rshrn2", 0x4f008c00, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Rshrn2V8H"),
		L("rshrn", 0x0f008c00, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("RshrnV2S"),
		L("rshrn2", 0x4f008c00, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Rshrn2V4S"),
		L("sqshrn", 0x0f009400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SqshrnV8B"),
		L("sqshrn2", 0x4f009400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Sqshrn2V16B"),
		L("sqshrn", 0x0f009400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SqshrnV4H"),
		L("sqshrn2", 0x4f009400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Sqshrn2V8H"),
		L("sqshrn", 0x0f009400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SqshrnV2S"),
		L("sqshrn2", 0x4f009400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Sqshrn2V4S"),
		L("uqshrn", 0x2f009400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UqshrnV8B"),
		L("uqshrn2", 0x6f009400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Uqshrn2V16B"),
		L("uqshrn", 0x2f009400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UqshrnV4H"),
		L("uqshrn2", 0x6f009400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Uqshrn2V8H"),
		L("uqshrn", 0x2f009400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("UqshrnV2S"),
		L("uqshrn2", 0x6f009400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Uqshrn2V4S"),
		L("sqshrun", 0x2f008400, 0xff80fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SqshrunV8B"),
		L("sqshrun2", 0x6f008400, 0xff80fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Sqshrun2V16B"),
		L("sqshrun", 0x2f008400, 0xff80fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SqshrunV4H"),
		L("sqshrun2", 0x6f008400, 0xff80fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Sqshrun2V8H"),
		L("sqshrun", 0x2f008400, 0xff80fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("SqshrunV2S"),
		L("sqshrun2", 0x6f008400, 0xff80fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V2D, Rn).Imm(F(16, 7)).Kind(ImmVecShr).Name("Sqshrun2V4S"),
		L("ushll", 0x2f00a400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("UshllV8H"),
		L("ushll2", 0x6f00a400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("Ushll2V8H"),
		L("ushll", 0x2f00a400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("UshllV4S"),
		L("ushll2", 0x6f00a400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("Ushll2V4S"),
		L("ushll", 0x2f00a400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("UshllV2D"),
		L("ushll2", 0x6f00a400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("Ushll2V2D"),
		L("sshll", 0x0f00a400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SshllV8H"),
		L("sshll2", 0x4f00a400, 0xff80fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("Sshll2V8H"),
		L("sshll", 0x0f00a400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SshllV4S"),
		L("sshll2", 0x4f00a400, 0xff80fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("Sshll2V4S"),
		L("sshll", 0x0f00a400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("SshllV2D"),
		L("sshll2", 0x4f00a400, 0xff80fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).Imm(F(16, 7)).Kind(ImmVecShl).Name("Sshll2V2D"),

		// ---- Widening, three registers ----
		L("umull", 0x2e20c000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UmullV8H"),
		L("umull2", 0x6e20c000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Umull2V8H"),
		L("umull", 0x2e60c000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UmullV4S"),
		L("umull2", 0x6e60c000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Umull2V4S"),
		L("umull", 0x2ea0c000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UmullV2D"),
		L("umull2", 0x6ea0c000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Umull2V2D"),
		L("smull", 0x0e20c000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SmullV8H"),
		L("smull2", 0x4e20c000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Smull2V8H"),
		L("smull", 0x0e60c000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SmullV4S"),
		L("smull2", 0x4e60c000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Smull2V4S"),
		L("smull", 0x0ea0c000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SmullV2D"),
		L("smull2", 0x4ea0c000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Smull2V2D"),
		L("umlal", 0x2e208000, 0xffe0fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UmlalV8H"),
		L("umlal2", 0x6e208000, 0xffe0fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Umlal2V8H"),
		L("umlal", 0x2e608000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UmlalV4S"),
		L("umlal2", 0x6e608000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Umlal2V4S"),
		L("umlal", 0x2ea08000, 0xffe0fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UmlalV2D"),
		L("umlal2", 0x6ea08000, 0xffe0fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Umlal2V2D"),
		L("smlal", 0x0e208000, 0xffe0fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SmlalV8H"),
		L("smlal2", 0x4e208000, 0xffe0fc00).
			SrcDstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Smlal2V8H"),
		L("smlal", 0x0e608000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SmlalV4S"),
		L("smlal2", 0x4e608000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Smlal2V4S"),
		L("smlal", 0x0ea08000, 0xffe0fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SmlalV2D"),
		L("smlal2", 0x4ea08000, 0xffe0fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Smlal2V2D"),
		L("uaddl", 0x2e200000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UaddlV8H"),
		L("uaddl2", 0x6e200000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Uaddl2V8H"),
		L("uaddl", 0x2e600000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UaddlV4S"),
		L("uaddl2", 0x6e600000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Uaddl2V4S"),
		L("uaddl", 0x2ea00000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UaddlV2D"),
		L("uaddl2", 0x6ea00000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Uaddl2V2D"),
		L("saddl", 0x0e200000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SaddlV8H"),
		L("saddl2", 0x4e200000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Saddl2V8H"),
		L("saddl", 0x0e600000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SaddlV4S"),
		L("saddl2", 0x4e600000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Saddl2V4S"),
		L("saddl", 0x0ea00000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SaddlV2D"),
		L("saddl2", 0x4ea00000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Saddl2V2D"),
		L("usubl", 0x2e202000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("UsublV8H"),
		L("usubl2", 0x6e202000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Usubl2V8H"),
		L("usubl", 0x2e602000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("UsublV4S"),
		L("usubl2", 0x6e602000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Usubl2V4S"),
		L("usubl", 0x2ea02000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("UsublV2D"),
		L("usubl2", 0x6ea02000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Usubl2V2D"),
		L("ssubl", 0x0e202000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("SsublV8H"),
		L("ssubl2", 0x4e202000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Ssubl2V8H"),
		L("ssubl", 0x0e602000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("SsublV4S"),
		L("ssubl2", 0x4e602000, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Ssubl2V4S"),
		L("ssubl", 0x0ea02000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("SsublV2D"),
		L("ssubl2", 0x4ea02000, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Ssubl2V2D"),
		L("pmull", 0x0e20e000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("PmullV8H"),
		L("pmull2", 0x4e20e000, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Pmull2V8H"),
		L("pmull", 0x0ee0e000, 0xffe0fc00).
			DstV(reg.V1Q, Rd).SrcV(reg.V1D, Rn).SrcV(reg.V1D, Rm).Gate2(feature.AES).Name("PmullV1Q"),
		L("pmull2", 0x4ee0e000, 0xffe0fc00).
			DstV(reg.V1Q, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Gate2(feature.AES).Name("Pmull2V1Q"),

		// ---- Permutes ----
		L("uzp1", 0x0e001800, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("Uzp1V8B"),
		L("uzp1", 0x4e001800, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Uzp1V16B"),
		L("uzp1", 0x0e401800, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("Uzp1V4H"),
		L("uzp1", 0x4e401800, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Uzp1V8H"),
		L("uzp1", 0x0e801800, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("Uzp1V2S"),
		L("uzp1", 0x4e801800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Uzp1V4S"),
		L("uzp1", 0x4ec01800, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("Uzp1V2D"),
		L("trn1", 0x0e002800, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("Trn1V8B"),
		L("trn1", 0x4e002800, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Trn1V16B"),
		L("trn1", 0x0e402800, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("Trn1V4H"),
		L("trn1", 0x4e402800, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Trn1V8H"),
		L("trn1", 0x0e802800, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("Trn1V2S"),
		L("trn1", 0x4e802800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Trn1V4S"),
		L("trn1", 0x4ec02800, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("Trn1V2D"),
		L("zip1", 0x0e003800, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("Zip1V8B"),
		L("zip1", 0x4e003800, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Zip1V16B"),
		L("zip1", 0x0e403800, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("Zip1V4H"),
		L("zip1", 0x4e403800, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Zip1V8H"),
		L("zip1", 0x0e803800, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("Zip1V2S"),
		L("zip1", 0x4e803800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Zip1V4S"),
		L("zip1", 0x4ec03800, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("Zip1V2D"),
		L("uzp2", 0x0e005800, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("Uzp2V8B"),
		L("uzp2", 0x4e005800, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Uzp2V16B"),
		L("uzp2", 0x0e405800, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("Uzp2V4H"),
		L("uzp2", 0x4e405800, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Uzp2V8H"),
		L("uzp2", 0x0e805800, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("Uzp2V2S"),
		L("uzp2", 0x4e805800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Uzp2V4S"),
		L("uzp2", 0x4ec05800, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("Uzp2V2D"),
		L("trn2", 0x0e006800, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("Trn2V8B"),
		L("trn2", 0x4e006800, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Trn2V16B"),
		L("trn2", 0x0e406800, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("Trn2V4H"),
		L("trn2", 0x4e406800, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Trn2V8H"),
		L("trn2", 0x0e806800, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("Trn2V2S"),
		L("trn2", 0x4e806800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Trn2V4S"),
		L("trn2", 0x4ec06800, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("Trn2V2D"),
		L("zip2", 0x0e007800, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Name("Zip2V8B"),
		L("zip2", 0x4e007800, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("Zip2V16B"),
		L("zip2", 0x0e407800, 0xffe0fc00).
			DstV(reg.V4H, Rd).SrcV(reg.V4H, Rn).SrcV(reg.V4H, Rm).Name("Zip2V4H"),
		L("zip2", 0x4e407800, 0xffe0fc00).
			DstV(reg.V8H, Rd).SrcV(reg.V8H, Rn).SrcV(reg.V8H, Rm).Name("Zip2V8H"),
		L("zip2", 0x0e807800, 0xffe0fc00).
			DstV(reg.V2S, Rd).SrcV(reg.V2S, Rn).SrcV(reg.V2S, Rm).Name("Zip2V2S"),
		L("zip2", 0x4e807800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Name("Zip2V4S"),
		L("zip2", 0x4ec07800, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Name("Zip2V2D"),
		L("ext", 0x2e000000, 0xffe08400).
			DstV(reg.V8B, Rd).SrcV(reg.V8B, Rn).SrcV(reg.V8B, Rm).Imm(F(11, 4)).Name("ExtV8B"),
		L("ext", 0x6e000000, 0xffe08400).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Imm(F(11, 4)).Name("ExtV16B"),
		L("tbl", 0x0e000000, 0xffe0fc00).
			DstV(reg.V8B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V8B, Rm).Name("TblV8B"),
		L("tbl", 0x4e000000, 0xffe0fc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("TblV16B"),
		L("tbx", 0x0e001000, 0xffe0fc00).
			SrcDstV(reg.V8B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V8B, Rm).Name("TbxV8B"),
		L("tbx", 0x4e001000, 0xffe0fc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).Name("TbxV16B"),

		// ---- Lanes and general registers ----
		L("dup", 0x0e010c00, 0xfffffc00).
			DstV(reg.V8B, Rd).Src(ClassW, Rn).Name("DupV8BGen"),
		L("dup", 0x0e010400, 0xffe1fc00).
			DstV(reg.V8B, Rd).SrcLane(reg.ElemB, Rn, F(17, 4)).Name("DupV8BLane"),
		L("dup", 0x4e010c00, 0xfffffc00).
			DstV(reg.V16B, Rd).Src(ClassW, Rn).Name("DupV16BGen"),
		L("dup", 0x4e010400, 0xffe1fc00).
			DstV(reg.V16B, Rd).SrcLane(reg.ElemB, Rn, F(17, 4)).Name("DupV16BLane"),
		L("umov", 0x0e013c00, 0xffe1fc00).
			Dst(ClassW, Rd).SrcLane(reg.ElemB, Rn, F(17, 4)).Name("UmovB"),
		L("ins", 0x4e011c00, 0xffe1fc00).
			DstLane(reg.ElemB, Rd, F(17, 4)).Src(ClassW, Rn).Name("InsBGen"),
		L("dup", 0x0e020c00, 0xfffffc00).
			DstV(reg.V4H, Rd).Src(ClassW, Rn).Name("DupV4HGen"),
		L("dup", 0x0e020400, 0xffe3fc00).
			DstV(reg.V4H, Rd).SrcLane(reg.ElemH, Rn, F(18, 3)).Name("DupV4HLane"),
		L("dup", 0x4e020c00, 0xfffffc00).
			DstV(reg.V8H, Rd).Src(ClassW, Rn).Name("DupV8HGen"),
		L("dup", 0x4e020400, 0xffe3fc00).
			DstV(reg.V8H, Rd).SrcLane(reg.ElemH, Rn, F(18, 3)).Name("DupV8HLane"),
		L("umov", 0x0e023c00, 0xffe3fc00).
			Dst(ClassW, Rd).SrcLane(reg.ElemH, Rn, F(18, 3)).Name("UmovH"),
		L("ins", 0x4e021c00, 0xffe3fc00).
			DstLane(reg.ElemH, Rd, F(18, 3)).Src(ClassW, Rn).Name("InsHGen"),
		L("dup", 0x0e040c00, 0xfffffc00).
			DstV(reg.V2S, Rd).Src(ClassW, Rn).Name("DupV2SGen"),
		L("dup", 0x0e040400, 0xffe7fc00).
			DstV(reg.V2S, Rd).SrcLane(reg.ElemS, Rn, F(19, 2)).Name("DupV2SLane"),
		L("dup", 0x4e040c00, 0xfffffc00).
			DstV(reg.V4S, Rd).Src(ClassW, Rn).Name("DupV4SGen"),
		L("dup", 0x4e040400, 0xffe7fc00).
			DstV(reg.V4S, Rd).SrcLane(reg.ElemS, Rn, F(19, 2)).Name("DupV4SLane"),
		L("umov", 0x0e043c00, 0xffe7fc00).
			Dst(ClassW, Rd).SrcLane(reg.ElemS, Rn, F(19, 2)).Name("UmovS"),
		L("ins", 0x4e041c00, 0xffe7fc00).
			DstLane(reg.ElemS, Rd, F(19, 2)).Src(ClassW, Rn).Name("InsSGen"),
		L("dup", 0x4e080c00, 0xfffffc00).
			DstV(reg.V2D, Rd).Src(ClassX, Rn).Name("DupV2DGen"),
		L("dup", 0x4e080400, 0xffeffc00).
			DstV(reg.V2D, Rd).SrcLane(reg.ElemD, Rn, F(20, 1)).Name("DupV2DLane"),
		L("umov", 0x4e083c00, 0xffeffc00).
			Dst(ClassX, Rd).SrcLane(reg.ElemD, Rn, F(20, 1)).Name("UmovD"),
		L("ins", 0x4e081c00, 0xffeffc00).
			DstLane(reg.ElemD, Rd, F(20, 1)).Src(ClassX, Rn).Name("InsDGen"),
		L("smov", 0x0e012c00, 0xffe1fc00).
			Dst(ClassW, Rd).SrcLane(reg.ElemB, Rn, F(17, 4)).Name("SmovB32"),
		L("smov", 0x0e022c00, 0xffe3fc00).
			Dst(ClassW, Rd).SrcLane(reg.ElemH, Rn, F(18, 3)).Name("SmovH32"),
		L("smov", 0x4e012c00, 0xffe1fc00).
			Dst(ClassX, Rd).SrcLane(reg.ElemB, Rn, F(17, 4)).Name("SmovB64"),
		L("smov", 0x4e022c00, 0xffe3fc00).
			Dst(ClassX, Rd).SrcLane(reg.ElemH, Rn, F(18, 3)).Name("SmovH64"),
		L("smov", 0x4e042c00, 0xffe7fc00).
			Dst(ClassX, Rd).SrcLane(reg.ElemS, Rn, F(19, 2)).Name("SmovS64"),

		// ---- Immediates into a vector ----
		L("movi", 0x6f00e400, 0xfff8fc00).
			DstV(reg.V2D, Rd).Imm(FSplit(5, 5, 16, 3)).Kind(ImmVecByteMask).Name("MoviV2D"),
		L("movi", 0x0f00e400, 0xfff8fc00).
			DstV(reg.V8B, Rd).Imm(FSplit(5, 5, 16, 3)).Name("MoviV8B"),
		L("movi", 0x4f00e400, 0xfff8fc00).
			DstV(reg.V16B, Rd).Imm(FSplit(5, 5, 16, 3)).Name("MoviV16B"),

		// ---- AES ----
		L("aese", 0x4e284800, 0xfffffc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Gate2(feature.AES).Name("Aese"),
		L("aesd", 0x4e285800, 0xfffffc00).
			SrcDstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Gate2(feature.AES).Name("Aesd"),
		L("aesmc", 0x4e286800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Gate2(feature.AES).Name("Aesmc"),
		L("aesimc", 0x4e287800, 0xfffffc00).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).Gate2(feature.AES).Name("Aesimc"),

		// ---- SHA-1 and SHA-256 ----
		L("sha1c", 0x5e000000, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassS, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha1c"),
		L("sha1p", 0x5e001000, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassS, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha1p"),
		L("sha1m", 0x5e002000, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassS, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha1m"),
		L("sha1su0", 0x5e003000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha1su0"),
		L("sha256h", 0x5e004000, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassQ, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha256h"),
		L("sha256h2", 0x5e005000, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassQ, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha256h2"),
		L("sha256su1", 0x5e006000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SHA2).Name("Sha256su1"),
		L("sha1h", 0x5e280800, 0xfffffc00).
			Dst(ClassS, Rd).Src(ClassS, Rn).Gate2(feature.SHA2).Name("Sha1h"),
		L("sha1su1", 0x5e281800, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Gate2(feature.SHA2).Name("Sha1su1"),
		L("sha256su0", 0x5e282800, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Gate2(feature.SHA2).Name("Sha256su0"),

		// ---- SHA-512 and SHA-3 ----
		L("sha512h", 0xce608000, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassQ, Rn).SrcV(reg.V2D, Rm).Gate2(feature.SHA3).Name("Sha512h"),
		L("sha512h2", 0xce608400, 0xffe0fc00).
			SrcDst(ClassQ, Rd).Src(ClassQ, Rn).SrcV(reg.V2D, Rm).Gate2(feature.SHA3).Name("Sha512h2"),
		L("sha512su1", 0xce608800, 0xffe0fc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Gate2(feature.SHA3).Name("Sha512su1"),
		L("sha512su0", 0xcec08000, 0xfffffc00).
			SrcDstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).Gate2(feature.SHA3).Name("Sha512su0"),
		L("eor3", 0xce000000, 0xffe08000).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).SrcV(reg.V16B, Ra).Gate2(feature.SHA3).Name("Eor3"),
		L("bcax", 0xce200000, 0xffe08000).
			DstV(reg.V16B, Rd).SrcV(reg.V16B, Rn).SrcV(reg.V16B, Rm).SrcV(reg.V16B, Ra).Gate2(feature.SHA3).Name("Bcax"),
		L("rax1", 0xce608c00, 0xffe0fc00).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Gate2(feature.SHA3).Name("Rax1"),
		L("xar", 0xce800000, 0xffe00000).
			DstV(reg.V2D, Rd).SrcV(reg.V2D, Rn).SrcV(reg.V2D, Rm).Imm(F(10, 6)).Gate2(feature.SHA3).Name("Xar"),

		// ---- SM3 and SM4 ----
		L("sm3ss1", 0xce400000, 0xffe08000).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).SrcV(reg.V4S, Ra).Gate2(feature.SM4).Name("Sm3ss1"),
		L("sm3tt1a", 0xce408000, 0xffe0cc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcLane(reg.ElemS, Rm, F(12, 2)).Gate2(feature.SM4).Name("Sm3tt1a"),
		L("sm3tt1b", 0xce408400, 0xffe0cc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcLane(reg.ElemS, Rm, F(12, 2)).Gate2(feature.SM4).Name("Sm3tt1b"),
		L("sm3tt2a", 0xce408800, 0xffe0cc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcLane(reg.ElemS, Rm, F(12, 2)).Gate2(feature.SM4).Name("Sm3tt2a"),
		L("sm3tt2b", 0xce408c00, 0xffe0cc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcLane(reg.ElemS, Rm, F(12, 2)).Gate2(feature.SM4).Name("Sm3tt2b"),
		L("sm3partw1", 0xce60c000, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SM4).Name("Sm3partw1"),
		L("sm3partw2", 0xce60c400, 0xffe0fc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SM4).Name("Sm3partw2"),
		L("sm4ekey", 0xce60c800, 0xffe0fc00).
			DstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).SrcV(reg.V4S, Rm).Gate2(feature.SM4).Name("Sm4ekey"),
		L("sm4e", 0xcec08400, 0xfffffc00).
			SrcDstV(reg.V4S, Rd).SrcV(reg.V4S, Rn).Gate2(feature.SM4).Name("Sm4e"),

		// ---- CRC-32 and CRC-32C ----
		L("crc32b", 0x1ac04000, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Gate2(feature.CRC).Name("Crc32b"),
		L("crc32h", 0x1ac04400, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Gate2(feature.CRC).Name("Crc32h"),
		L("crc32w", 0x1ac04800, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Gate2(feature.CRC).Name("Crc32w"),
		L("crc32x", 0x9ac04c00, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassX, Rm).Gate2(feature.CRC).Name("Crc32x"),
		L("crc32cb", 0x1ac05000, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Gate2(feature.CRC).Name("Crc32cb"),
		L("crc32ch", 0x1ac05400, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Gate2(feature.CRC).Name("Crc32ch"),
		L("crc32cw", 0x1ac05800, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassW, Rm).Gate2(feature.CRC).Name("Crc32cw"),
		L("crc32cx", 0x9ac05c00, 0xffe0fc00).
			Dst(ClassW, Rd).Src(ClassW, Rn).Src(ClassX, Rm).Gate2(feature.CRC).Name("Crc32cx"),
	)
}
