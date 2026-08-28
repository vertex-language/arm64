// Package arm64: this file is the root re-export. One declaration per
// concept, aliased upward. A value crossing a package line in this tree is
// never converted, only renamed: arm64.RefAdrPage21, operand.RefAdrPage21 and
// obj.RefAdrPage21 are the same constant and no conversion exists anywhere.
package arm64

import (
	"github.com/vertex-language/arm64/feature"
	"github.com/vertex-language/arm64/obj"
	"github.com/vertex-language/arm64/operand"
	"github.com/vertex-language/arm64/reg"
)

type (
	Object    = obj.Object
	Reference = obj.Reference
	Error     = obj.Error
	Symbol    = obj.Symbol

	Binding    = obj.Binding
	SymbolType = obj.SymbolType
	Visibility = obj.Visibility
	RefKind    = obj.RefKind
)

const (
	Local  = obj.Local
	Global = obj.Global
	Weak   = obj.Weak

	NoType      = obj.NoType
	Func        = obj.Func
	ObjectSym   = obj.ObjectSym
	ThreadLocal = obj.ThreadLocal

	Default   = obj.Default
	Hidden    = obj.Hidden
	Protected = obj.Protected
	Internal  = obj.Internal
)

// The relocation kinds, re-exported so a caller naming one on a Ref never
// imports obj for it.
const (
	RefAbs64 = obj.RefAbs64
	RefAbs32 = obj.RefAbs32
	RefAbs16 = obj.RefAbs16

	RefPrel64 = obj.RefPrel64
	RefPrel32 = obj.RefPrel32
	RefPrel16 = obj.RefPrel16

	RefCall26   = obj.RefCall26
	RefJump26   = obj.RefJump26
	RefCondBr19 = obj.RefCondBr19
	RefTstBr14  = obj.RefTstBr14

	RefAdrPrel21 = obj.RefAdrPrel21
	RefAdrPage21 = obj.RefAdrPage21

	RefAddAbsLo12     = obj.RefAddAbsLo12
	RefLdSt8AbsLo12   = obj.RefLdSt8AbsLo12
	RefLdSt16AbsLo12  = obj.RefLdSt16AbsLo12
	RefLdSt32AbsLo12  = obj.RefLdSt32AbsLo12
	RefLdSt64AbsLo12  = obj.RefLdSt64AbsLo12
	RefLdSt128AbsLo12 = obj.RefLdSt128AbsLo12

	RefAdrGotPage21 = obj.RefAdrGotPage21
	RefLd64GotLo12  = obj.RefLd64GotLo12

	RefTlsGdAdrPage21         = obj.RefTlsGdAdrPage21
	RefTlsGdAddLo12           = obj.RefTlsGdAddLo12
	RefTlsIeAdrGottprelPage21 = obj.RefTlsIeAdrGottprelPage21
	RefTlsIeLd64GottprelLo12  = obj.RefTlsIeLd64GottprelLo12
	RefTlsLeAddTprelHi12      = obj.RefTlsLeAddTprelHi12
	RefTlsLeAddTprelLo12      = obj.RefTlsLeAddTprelLo12

	RefTLV = obj.RefTLV

	RefSize32   = obj.RefSize32
	RefSize64   = obj.RefSize64
	RefSecRel32 = obj.RefSecRel32
	RefSecIdx   = obj.RefSecIdx
)

// The error sentinels. errors.Is works against every one of them, and the
// concrete type behind all of them is *obj.Error.
var (
	ErrFeature     = obj.ErrFeature
	ErrForm        = obj.ErrForm
	ErrOperand     = obj.ErrOperand
	ErrDuplicate   = obj.ErrDuplicate
	ErrUndefined   = obj.ErrUndefined
	ErrRange       = obj.ErrRange
	ErrBitmask     = obj.ErrBitmask
	ErrAlign       = obj.ErrAlign
	ErrFinalized   = obj.ErrFinalized
	ErrRefKind     = obj.ErrRefKind
	ErrSectionName = obj.ErrSectionName
)

// ---- registers -------------------------------------------------------------

const (
	X0, X1, X2, X3, X4, X5, X6, X7         = reg.X0, reg.X1, reg.X2, reg.X3, reg.X4, reg.X5, reg.X6, reg.X7
	X8, X9, X10, X11, X12, X13, X14, X15   = reg.X8, reg.X9, reg.X10, reg.X11, reg.X12, reg.X13, reg.X14, reg.X15
	X16, X17, X18, X19, X20, X21, X22, X23 = reg.X16, reg.X17, reg.X18, reg.X19, reg.X20, reg.X21, reg.X22, reg.X23
	X24, X25, X26, X27, X28, X29, X30, XZR = reg.X24, reg.X25, reg.X26, reg.X27, reg.X28, reg.X29, reg.X30, reg.XZR
)

const (
	W0, W1, W2, W3, W4, W5, W6, W7         = reg.W0, reg.W1, reg.W2, reg.W3, reg.W4, reg.W5, reg.W6, reg.W7
	W8, W9, W10, W11, W12, W13, W14, W15   = reg.W8, reg.W9, reg.W10, reg.W11, reg.W12, reg.W13, reg.W14, reg.W15
	W16, W17, W18, W19, W20, W21, W22, W23 = reg.W16, reg.W17, reg.W18, reg.W19, reg.W20, reg.W21, reg.W22, reg.W23
	W24, W25, W26, W27, W28, W29, W30, WZR = reg.W24, reg.W25, reg.W26, reg.W27, reg.W28, reg.W29, reg.W30, reg.WZR
)

// SP, WSP and the AAPCS64 role names. IP0, IP1 and LR drop their X prefix
// here because a caller writing Blr(LR) is stating a role, not a register
// number; X30 is exactly as valid and exactly less clear.
const (
	SP  = reg.SP
	WSP = reg.WSP
	IP0 = reg.IP0
	IP1 = reg.IP1
	FP  = reg.FP
	LR  = reg.LR
)

const (
	V0, V1, V2, V3, V4, V5, V6, V7         = reg.V0, reg.V1, reg.V2, reg.V3, reg.V4, reg.V5, reg.V6, reg.V7
	V8, V9, V10, V11, V12, V13, V14, V15   = reg.V8, reg.V9, reg.V10, reg.V11, reg.V12, reg.V13, reg.V14, reg.V15
	V16, V17, V18, V19, V20, V21, V22, V23 = reg.V16, reg.V17, reg.V18, reg.V19, reg.V20, reg.V21, reg.V22, reg.V23
	V24, V25, V26, V27, V28, V29, V30, V31 = reg.V24, reg.V25, reg.V26, reg.V27, reg.V28, reg.V29, reg.V30, reg.V31
)

const (
	Q0, Q1, Q2, Q3, Q4, Q5, Q6, Q7         = reg.Q0, reg.Q1, reg.Q2, reg.Q3, reg.Q4, reg.Q5, reg.Q6, reg.Q7
	Q8, Q9, Q10, Q11, Q12, Q13, Q14, Q15   = reg.Q8, reg.Q9, reg.Q10, reg.Q11, reg.Q12, reg.Q13, reg.Q14, reg.Q15
	Q16, Q17, Q18, Q19, Q20, Q21, Q22, Q23 = reg.Q16, reg.Q17, reg.Q18, reg.Q19, reg.Q20, reg.Q21, reg.Q22, reg.Q23
	Q24, Q25, Q26, Q27, Q28, Q29, Q30, Q31 = reg.Q24, reg.Q25, reg.Q26, reg.Q27, reg.Q28, reg.Q29, reg.Q30, reg.Q31
)

const (
	D0, D1, D2, D3, D4, D5, D6, D7         = reg.D0, reg.D1, reg.D2, reg.D3, reg.D4, reg.D5, reg.D6, reg.D7
	D8, D9, D10, D11, D12, D13, D14, D15   = reg.D8, reg.D9, reg.D10, reg.D11, reg.D12, reg.D13, reg.D14, reg.D15
	D16, D17, D18, D19, D20, D21, D22, D23 = reg.D16, reg.D17, reg.D18, reg.D19, reg.D20, reg.D21, reg.D22, reg.D23
	D24, D25, D26, D27, D28, D29, D30, D31 = reg.D24, reg.D25, reg.D26, reg.D27, reg.D28, reg.D29, reg.D30, reg.D31
)

const (
	S0, S1, S2, S3, S4, S5, S6, S7         = reg.S0, reg.S1, reg.S2, reg.S3, reg.S4, reg.S5, reg.S6, reg.S7
	S8, S9, S10, S11, S12, S13, S14, S15   = reg.S8, reg.S9, reg.S10, reg.S11, reg.S12, reg.S13, reg.S14, reg.S15
	S16, S17, S18, S19, S20, S21, S22, S23 = reg.S16, reg.S17, reg.S18, reg.S19, reg.S20, reg.S21, reg.S22, reg.S23
	S24, S25, S26, S27, S28, S29, S30, S31 = reg.S24, reg.S25, reg.S26, reg.S27, reg.S28, reg.S29, reg.S30, reg.S31
)

const (
	H0, H1, H2, H3, H4, H5, H6, H7         = reg.H0, reg.H1, reg.H2, reg.H3, reg.H4, reg.H5, reg.H6, reg.H7
	H8, H9, H10, H11, H12, H13, H14, H15   = reg.H8, reg.H9, reg.H10, reg.H11, reg.H12, reg.H13, reg.H14, reg.H15
	H16, H17, H18, H19, H20, H21, H22, H23 = reg.H16, reg.H17, reg.H18, reg.H19, reg.H20, reg.H21, reg.H22, reg.H23
	H24, H25, H26, H27, H28, H29, H30, H31 = reg.H24, reg.H25, reg.H26, reg.H27, reg.H28, reg.H29, reg.H30, reg.H31
)

const (
	B0, B1, B2, B3, B4, B5, B6, B7         = reg.B0, reg.B1, reg.B2, reg.B3, reg.B4, reg.B5, reg.B6, reg.B7
	B8, B9, B10, B11, B12, B13, B14, B15   = reg.B8, reg.B9, reg.B10, reg.B11, reg.B12, reg.B13, reg.B14, reg.B15
	B16, B17, B18, B19, B20, B21, B22, B23 = reg.B16, reg.B17, reg.B18, reg.B19, reg.B20, reg.B21, reg.B22, reg.B23
	B24, B25, B26, B27, B28, B29, B30, B31 = reg.B24, reg.B25, reg.B26, reg.B27, reg.B28, reg.B29, reg.B30, reg.B31
)

// The scalable vector and predicate registers. Nothing in the ISA table
// reaches these yet — SVE is a tranche of its own — but the register file is
// a fact about the architecture and not about the table's coverage, so they
// are declared here regardless, the same reasoning amd64 states for its
// AVX-512 registers 16 through 31.
const (
	Z0, Z1, Z2, Z3, Z4, Z5, Z6, Z7         = reg.Z0, reg.Z1, reg.Z2, reg.Z3, reg.Z4, reg.Z5, reg.Z6, reg.Z7
	Z8, Z9, Z10, Z11, Z12, Z13, Z14, Z15   = reg.Z8, reg.Z9, reg.Z10, reg.Z11, reg.Z12, reg.Z13, reg.Z14, reg.Z15
	Z16, Z17, Z18, Z19, Z20, Z21, Z22, Z23 = reg.Z16, reg.Z17, reg.Z18, reg.Z19, reg.Z20, reg.Z21, reg.Z22, reg.Z23
	Z24, Z25, Z26, Z27, Z28, Z29, Z30, Z31 = reg.Z24, reg.Z25, reg.Z26, reg.Z27, reg.Z28, reg.Z29, reg.Z30, reg.Z31
)

const (
	P0, P1, P2, P3, P4, P5, P6, P7       = reg.P0, reg.P1, reg.P2, reg.P3, reg.P4, reg.P5, reg.P6, reg.P7
	P8, P9, P10, P11, P12, P13, P14, P15 = reg.P8, reg.P9, reg.P10, reg.P11, reg.P12, reg.P13, reg.P14, reg.P15
)

var FFR = reg.FFR

// ---- operand vocabulary -----------------------------------------------------

type (
	Imm      = operand.Imm
	Mem      = operand.Mem
	Label    = operand.Label
	SymRef   = operand.SymRef
	AddrRef  = operand.AddrRef
	Cond     = operand.Cond
	Shift    = operand.Shift
	ShiftOp  = operand.ShiftOp
	Extend   = operand.Extend
	ExtendOp = operand.ExtendOp
	Barrier  = operand.Barrier
	PrfOp    = operand.PrfOp
	Sys      = reg.Sys
)

// RegSP64, RegSP32, ImmOrRef and TargetOp are the operand slots Go's type
// system cannot close.
//
// RegSP64 accepts SP or a numbered reg.X: a slot that reads register 31 as
// the stack pointer, which is a different type from reg.X and reg.Xsp both,
// and no single Go type names their union — the same problem x86_64's RM
// types solve the same way, a documented any refused by name at the call
// when it holds the wrong thing.
//
// ImmOrRef is an immediate slot that also takes the :lo12: half of an
// address — an int, an int64, or a PageOff/GotPageOff AddrRef.
//
// TargetOp is a branch or address destination: a SymRef, or an AddrRef
// wrapping one, for the Page/GotPage/GotPageOff modifiers.
type (
	RegSP64  = any
	RegSP32  = any
	ImmOrRef = any
	TargetOp = any
)

// Ref builds a symbol reference. Naming a kind is a request that blocks
// folding: Ref("puts", RefCall26) asks for a branch relocation even to a
// symbol two lines above.
func Ref(name string, kind ...RefKind) SymRef { return operand.Sym(name, kind...) }

var (
	Page       = operand.Page
	PageOff    = operand.PageOff
	GotPage    = operand.GotPage
	GotPageOff = operand.GotPageOff
	Direct     = operand.Direct
)

type baseReg interface{ reg.X | reg.Xsp }

func MemOf[T baseReg](base T) Mem  { return operand.MemOf(base) }
func Mem8[T baseReg](base T) Mem   { return operand.Mem8(base) }
func Mem16[T baseReg](base T) Mem  { return operand.Mem16(base) }
func Mem32[T baseReg](base T) Mem  { return operand.Mem32(base) }
func Mem64[T baseReg](base T) Mem  { return operand.Mem64(base) }
func Mem128[T baseReg](base T) Mem { return operand.Mem128(base) }

const (
	EQ, NE, CS, CC, MI, PL, VS, VC = operand.EQ, operand.NE, operand.CS, operand.CC, operand.MI, operand.PL, operand.VS, operand.VC
	HI, LS, GE, LT, GT, LE, AL, NV = operand.HI, operand.LS, operand.GE, operand.LT, operand.GT, operand.LE, operand.AL, operand.NV
	HS, LO                         = operand.HS, operand.LO
)

const (
	LSL, LSR, ASR, ROR = operand.LSL, operand.LSR, operand.ASR, operand.ROR
)

var Shifted = operand.Shifted

const (
	UXTB, UXTH, UXTW, UXTX = operand.UXTB, operand.UXTH, operand.UXTW, operand.UXTX
	SXTB, SXTH, SXTW, SXTX = operand.SXTB, operand.SXTH, operand.SXTW, operand.SXTX
	ExtLSL                 = operand.ExtLSL
)

var Extended = operand.Extended

const (
	SY, ISH, ISHLD, ISHST = operand.SY, operand.ISH, operand.ISHLD, operand.ISHST
	OSH, NSH, LD, ST      = operand.OSH, operand.NSH, operand.LD, operand.ST
)

const (
	PLDL1KEEP, PLDL1STRM = operand.PLDL1KEEP, operand.PLDL1STRM
	PLDL2KEEP, PLDL2STRM = operand.PLDL2KEEP, operand.PLDL2STRM
	PLDL3KEEP, PLDL3STRM = operand.PLDL3KEEP, operand.PLDL3STRM
	PLIL1KEEP, PLIL1STRM = operand.PLIL1KEEP, operand.PLIL1STRM
	PLIL2KEEP, PLIL2STRM = operand.PLIL2KEEP, operand.PLIL2STRM
	PLIL3KEEP, PLIL3STRM = operand.PLIL3KEEP, operand.PLIL3STRM
	PSTL1KEEP, PSTL1STRM = operand.PSTL1KEEP, operand.PSTL1STRM
	PSTL2KEEP, PSTL2STRM = operand.PSTL2KEEP, operand.PSTL2STRM
	PSTL3KEEP, PSTL3STRM = operand.PSTL3KEEP, operand.PSTL3STRM
)

// ---- system registers -------------------------------------------------------

func NewSys(op0, op1, crn, crm, op2 uint8) Sys { return reg.NewSys(op0, op1, crn, crm, op2) }

const (
	NZCV      = reg.NZCV
	DAIF      = reg.DAIF
	CurrentEL = reg.CurrentEL
	SPSel     = reg.SPSel

	FPCR = reg.FPCR
	FPSR = reg.FPSR

	TPIDR_EL0   = reg.TPIDR_EL0
	TPIDRRO_EL0 = reg.TPIDRRO_EL0
	TPIDR_EL1   = reg.TPIDR_EL1
	TPIDR_EL2   = reg.TPIDR_EL2
	TPIDR_EL3   = reg.TPIDR_EL3

	MIDR_EL1   = reg.MIDR_EL1
	MPIDR_EL1  = reg.MPIDR_EL1
	CTR_EL0    = reg.CTR_EL0
	DCZID_EL0  = reg.DCZID_EL0
	CNTVCT_EL0 = reg.CNTVCT_EL0

	SCTLR_EL1 = reg.SCTLR_EL1
	TTBR0_EL1 = reg.TTBR0_EL1
	TTBR1_EL1 = reg.TTBR1_EL1
	TCR_EL1   = reg.TCR_EL1
	ESR_EL1   = reg.ESR_EL1
	FAR_EL1   = reg.FAR_EL1
	VBAR_EL1  = reg.VBAR_EL1
	ELR_EL1   = reg.ELR_EL1
	SPSR_EL1  = reg.SPSR_EL1
)

// ---- features ---------------------------------------------------------------
//
// Every level and extension is re-exported at the root so the gating
// diagnostic's note line — aarch64.WithFeatures(...) — names something that
// compiles at the call site.

type (
	Feature    = feature.Feature
	FeatureSet = feature.Set
	Level      = feature.Level
)

var (
	Baseline          = feature.Baseline
	NewFeatureSet     = feature.NewSet
	ParseFeatures     = feature.ParseFeatures
	MustParseFeatures = feature.MustParseFeatures
	ParseLevel        = feature.ParseLevel
	Levels            = feature.Levels
	Decompose         = feature.Decompose
)

const (
	Armv8A, Armv8_1A, Armv8_2A, Armv8_3A, Armv8_4A           = feature.Armv8A, feature.Armv8_1A, feature.Armv8_2A, feature.Armv8_3A, feature.Armv8_4A
	Armv8_5A, Armv8_6A, Armv8_7A, Armv8_8A, Armv8_9A         = feature.Armv8_5A, feature.Armv8_6A, feature.Armv8_7A, feature.Armv8_8A, feature.Armv8_9A
	Armv9A, Armv9_1A, Armv9_2A, Armv9_3A, Armv9_4A, Armv9_5A = feature.Armv9A, feature.Armv9_1A, feature.Armv9_2A, feature.Armv9_3A, feature.Armv9_4A, feature.Armv9_5A
)

// FeatFP is feature.FP under a different name: reg.FP, the frame-pointer
// role of X29, already claims FP at this root.
const (
	FeatFP, SIMD, FP16, FP16FML, BF16, I8MM = feature.FP, feature.SIMD, feature.FP16, feature.FP16FML, feature.BF16, feature.I8MM
	DotProd, FCMA, JSCVT, FRIntTS           = feature.DotProd, feature.FCMA, feature.JSCVT, feature.FRIntTS
	FAMINMAX, LUT                           = feature.FAMINMAX, feature.LUT

	AES, SHA2, SHA3, SM4 = feature.AES, feature.SHA2, feature.SHA3, feature.SM4

	LSE, LSE128, D128, RCPC, RCPC2, RCPC3 = feature.LSE, feature.LSE128, feature.D128, feature.RCPC, feature.RCPC2, feature.RCPC3
	LS64, MOPS, XS, THE                   = feature.LS64, feature.MOPS, feature.XS, feature.THE

	CRC, RDMA, FlagM, FlagM2, CSSC, HBC, WFxT = feature.CRC, feature.RDMA, feature.FlagM, feature.FlagM2, feature.CSSC, feature.HBC, feature.WFxT

	PAuth, BTI, MemTag, SB, SSBS, PredRes, GCS, CPA, RNG, TME = feature.PAuth, feature.BTI, feature.MemTag, feature.SB, feature.SSBS, feature.PredRes, feature.GCS, feature.CPA, feature.RNG, feature.TME

	SVE, SVE2, SVE2AES, SVE2SM4, SVE2SHA3 = feature.SVE, feature.SVE2, feature.SVE2AES, feature.SVE2SM4, feature.SVE2SHA3
	SVE2BitPerm, SVE2p1, SVEB16B16        = feature.SVE2BitPerm, feature.SVE2p1, feature.SVEB16B16
	F32MM, F64MM                          = feature.F32MM, feature.F64MM

	SME, SMEI16I64, SMEF64F64, SME2, SME2p1, SMEB16B16, SMEF16F16 = feature.SME, feature.SMEI16I64, feature.SMEF64F64, feature.SME2, feature.SME2p1, feature.SMEB16B16, feature.SMEF16F16

	FP8, FP8FMA, FP8DOT4, FP8DOT2 = feature.FP8, feature.FP8FMA, feature.FP8DOT4, feature.FP8DOT2

	Profile = feature.Profile
)
