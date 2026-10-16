package arm64

import "github.com/vertex-language/arm64/reg"

// Advanced SIMD and the cryptographic extensions, one helper per row of
// internal/isa/table_simd.go: the arrangement is in the name -- AddV4S adds
// four 32-bit lanes -- and the registers are plain V, decorated here. A
// lane is a register and an index. Generated with the table; see there.

// ---- Integer arithmetic, three registers of one arrangement ----

var addV8B = form("AddV8B")

func (s *Section) AddV8B(rd, rn, rm reg.V) {
	s.inst(addV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var addV16B = form("AddV16B")

func (s *Section) AddV16B(rd, rn, rm reg.V) {
	s.inst(addV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var addV4H = form("AddV4H")

func (s *Section) AddV4H(rd, rn, rm reg.V) {
	s.inst(addV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var addV8H = form("AddV8H")

func (s *Section) AddV8H(rd, rn, rm reg.V) {
	s.inst(addV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var addV2S = form("AddV2S")

func (s *Section) AddV2S(rd, rn, rm reg.V) {
	s.inst(addV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var addV4S = form("AddV4S")

func (s *Section) AddV4S(rd, rn, rm reg.V) {
	s.inst(addV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var addV2D = form("AddV2D")

func (s *Section) AddV2D(rd, rn, rm reg.V) {
	s.inst(addV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var subV8B = form("SubV8B")

func (s *Section) SubV8B(rd, rn, rm reg.V) {
	s.inst(subV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var subV16B = form("SubV16B")

func (s *Section) SubV16B(rd, rn, rm reg.V) {
	s.inst(subV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var subV4H = form("SubV4H")

func (s *Section) SubV4H(rd, rn, rm reg.V) {
	s.inst(subV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var subV8H = form("SubV8H")

func (s *Section) SubV8H(rd, rn, rm reg.V) {
	s.inst(subV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var subV2S = form("SubV2S")

func (s *Section) SubV2S(rd, rn, rm reg.V) {
	s.inst(subV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var subV4S = form("SubV4S")

func (s *Section) SubV4S(rd, rn, rm reg.V) {
	s.inst(subV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var subV2D = form("SubV2D")

func (s *Section) SubV2D(rd, rn, rm reg.V) {
	s.inst(subV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var mulV8B = form("MulV8B")

func (s *Section) MulV8B(rd, rn, rm reg.V) {
	s.inst(mulV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var mulV16B = form("MulV16B")

func (s *Section) MulV16B(rd, rn, rm reg.V) {
	s.inst(mulV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var mulV4H = form("MulV4H")

func (s *Section) MulV4H(rd, rn, rm reg.V) {
	s.inst(mulV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var mulV8H = form("MulV8H")

func (s *Section) MulV8H(rd, rn, rm reg.V) {
	s.inst(mulV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var mulV2S = form("MulV2S")

func (s *Section) MulV2S(rd, rn, rm reg.V) {
	s.inst(mulV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var mulV4S = form("MulV4S")

func (s *Section) MulV4S(rd, rn, rm reg.V) {
	s.inst(mulV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var mlaV8B = form("MlaV8B")

func (s *Section) MlaV8B(rd, rn, rm reg.V) {
	s.inst(mlaV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var mlaV16B = form("MlaV16B")

func (s *Section) MlaV16B(rd, rn, rm reg.V) {
	s.inst(mlaV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var mlaV4H = form("MlaV4H")

func (s *Section) MlaV4H(rd, rn, rm reg.V) {
	s.inst(mlaV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var mlaV8H = form("MlaV8H")

func (s *Section) MlaV8H(rd, rn, rm reg.V) {
	s.inst(mlaV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var mlaV2S = form("MlaV2S")

func (s *Section) MlaV2S(rd, rn, rm reg.V) {
	s.inst(mlaV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var mlaV4S = form("MlaV4S")

func (s *Section) MlaV4S(rd, rn, rm reg.V) {
	s.inst(mlaV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var mlsV8B = form("MlsV8B")

func (s *Section) MlsV8B(rd, rn, rm reg.V) {
	s.inst(mlsV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var mlsV16B = form("MlsV16B")

func (s *Section) MlsV16B(rd, rn, rm reg.V) {
	s.inst(mlsV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var mlsV4H = form("MlsV4H")

func (s *Section) MlsV4H(rd, rn, rm reg.V) {
	s.inst(mlsV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var mlsV8H = form("MlsV8H")

func (s *Section) MlsV8H(rd, rn, rm reg.V) {
	s.inst(mlsV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var mlsV2S = form("MlsV2S")

func (s *Section) MlsV2S(rd, rn, rm reg.V) {
	s.inst(mlsV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var mlsV4S = form("MlsV4S")

func (s *Section) MlsV4S(rd, rn, rm reg.V) {
	s.inst(mlsV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmeqV8B = form("CmeqV8B")

func (s *Section) CmeqV8B(rd, rn, rm reg.V) {
	s.inst(cmeqV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var cmeqV16B = form("CmeqV16B")

func (s *Section) CmeqV16B(rd, rn, rm reg.V) {
	s.inst(cmeqV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var cmeqV4H = form("CmeqV4H")

func (s *Section) CmeqV4H(rd, rn, rm reg.V) {
	s.inst(cmeqV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var cmeqV8H = form("CmeqV8H")

func (s *Section) CmeqV8H(rd, rn, rm reg.V) {
	s.inst(cmeqV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var cmeqV2S = form("CmeqV2S")

func (s *Section) CmeqV2S(rd, rn, rm reg.V) {
	s.inst(cmeqV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var cmeqV4S = form("CmeqV4S")

func (s *Section) CmeqV4S(rd, rn, rm reg.V) {
	s.inst(cmeqV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmeqV2D = form("CmeqV2D")

func (s *Section) CmeqV2D(rd, rn, rm reg.V) {
	s.inst(cmeqV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var cmtstV8B = form("CmtstV8B")

func (s *Section) CmtstV8B(rd, rn, rm reg.V) {
	s.inst(cmtstV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var cmtstV16B = form("CmtstV16B")

func (s *Section) CmtstV16B(rd, rn, rm reg.V) {
	s.inst(cmtstV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var cmtstV4H = form("CmtstV4H")

func (s *Section) CmtstV4H(rd, rn, rm reg.V) {
	s.inst(cmtstV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var cmtstV8H = form("CmtstV8H")

func (s *Section) CmtstV8H(rd, rn, rm reg.V) {
	s.inst(cmtstV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var cmtstV2S = form("CmtstV2S")

func (s *Section) CmtstV2S(rd, rn, rm reg.V) {
	s.inst(cmtstV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var cmtstV4S = form("CmtstV4S")

func (s *Section) CmtstV4S(rd, rn, rm reg.V) {
	s.inst(cmtstV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmtstV2D = form("CmtstV2D")

func (s *Section) CmtstV2D(rd, rn, rm reg.V) {
	s.inst(cmtstV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var cmgtV8B = form("CmgtV8B")

func (s *Section) CmgtV8B(rd, rn, rm reg.V) {
	s.inst(cmgtV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var cmgtV16B = form("CmgtV16B")

func (s *Section) CmgtV16B(rd, rn, rm reg.V) {
	s.inst(cmgtV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var cmgtV4H = form("CmgtV4H")

func (s *Section) CmgtV4H(rd, rn, rm reg.V) {
	s.inst(cmgtV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var cmgtV8H = form("CmgtV8H")

func (s *Section) CmgtV8H(rd, rn, rm reg.V) {
	s.inst(cmgtV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var cmgtV2S = form("CmgtV2S")

func (s *Section) CmgtV2S(rd, rn, rm reg.V) {
	s.inst(cmgtV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var cmgtV4S = form("CmgtV4S")

func (s *Section) CmgtV4S(rd, rn, rm reg.V) {
	s.inst(cmgtV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmgtV2D = form("CmgtV2D")

func (s *Section) CmgtV2D(rd, rn, rm reg.V) {
	s.inst(cmgtV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var cmgeV8B = form("CmgeV8B")

func (s *Section) CmgeV8B(rd, rn, rm reg.V) {
	s.inst(cmgeV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var cmgeV16B = form("CmgeV16B")

func (s *Section) CmgeV16B(rd, rn, rm reg.V) {
	s.inst(cmgeV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var cmgeV4H = form("CmgeV4H")

func (s *Section) CmgeV4H(rd, rn, rm reg.V) {
	s.inst(cmgeV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var cmgeV8H = form("CmgeV8H")

func (s *Section) CmgeV8H(rd, rn, rm reg.V) {
	s.inst(cmgeV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var cmgeV2S = form("CmgeV2S")

func (s *Section) CmgeV2S(rd, rn, rm reg.V) {
	s.inst(cmgeV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var cmgeV4S = form("CmgeV4S")

func (s *Section) CmgeV4S(rd, rn, rm reg.V) {
	s.inst(cmgeV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmgeV2D = form("CmgeV2D")

func (s *Section) CmgeV2D(rd, rn, rm reg.V) {
	s.inst(cmgeV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var cmhiV8B = form("CmhiV8B")

func (s *Section) CmhiV8B(rd, rn, rm reg.V) {
	s.inst(cmhiV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var cmhiV16B = form("CmhiV16B")

func (s *Section) CmhiV16B(rd, rn, rm reg.V) {
	s.inst(cmhiV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var cmhiV4H = form("CmhiV4H")

func (s *Section) CmhiV4H(rd, rn, rm reg.V) {
	s.inst(cmhiV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var cmhiV8H = form("CmhiV8H")

func (s *Section) CmhiV8H(rd, rn, rm reg.V) {
	s.inst(cmhiV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var cmhiV2S = form("CmhiV2S")

func (s *Section) CmhiV2S(rd, rn, rm reg.V) {
	s.inst(cmhiV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var cmhiV4S = form("CmhiV4S")

func (s *Section) CmhiV4S(rd, rn, rm reg.V) {
	s.inst(cmhiV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmhiV2D = form("CmhiV2D")

func (s *Section) CmhiV2D(rd, rn, rm reg.V) {
	s.inst(cmhiV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var cmhsV8B = form("CmhsV8B")

func (s *Section) CmhsV8B(rd, rn, rm reg.V) {
	s.inst(cmhsV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var cmhsV16B = form("CmhsV16B")

func (s *Section) CmhsV16B(rd, rn, rm reg.V) {
	s.inst(cmhsV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var cmhsV4H = form("CmhsV4H")

func (s *Section) CmhsV4H(rd, rn, rm reg.V) {
	s.inst(cmhsV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var cmhsV8H = form("CmhsV8H")

func (s *Section) CmhsV8H(rd, rn, rm reg.V) {
	s.inst(cmhsV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var cmhsV2S = form("CmhsV2S")

func (s *Section) CmhsV2S(rd, rn, rm reg.V) {
	s.inst(cmhsV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var cmhsV4S = form("CmhsV4S")

func (s *Section) CmhsV4S(rd, rn, rm reg.V) {
	s.inst(cmhsV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var cmhsV2D = form("CmhsV2D")

func (s *Section) CmhsV2D(rd, rn, rm reg.V) {
	s.inst(cmhsV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var sqaddV8B = form("SqaddV8B")

func (s *Section) SqaddV8B(rd, rn, rm reg.V) {
	s.inst(sqaddV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var sqaddV16B = form("SqaddV16B")

func (s *Section) SqaddV16B(rd, rn, rm reg.V) {
	s.inst(sqaddV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var sqaddV4H = form("SqaddV4H")

func (s *Section) SqaddV4H(rd, rn, rm reg.V) {
	s.inst(sqaddV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sqaddV8H = form("SqaddV8H")

func (s *Section) SqaddV8H(rd, rn, rm reg.V) {
	s.inst(sqaddV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sqaddV2S = form("SqaddV2S")

func (s *Section) SqaddV2S(rd, rn, rm reg.V) {
	s.inst(sqaddV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sqaddV4S = form("SqaddV4S")

func (s *Section) SqaddV4S(rd, rn, rm reg.V) {
	s.inst(sqaddV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sqaddV2D = form("SqaddV2D")

func (s *Section) SqaddV2D(rd, rn, rm reg.V) {
	s.inst(sqaddV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var uqaddV8B = form("UqaddV8B")

func (s *Section) UqaddV8B(rd, rn, rm reg.V) {
	s.inst(uqaddV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uqaddV16B = form("UqaddV16B")

func (s *Section) UqaddV16B(rd, rn, rm reg.V) {
	s.inst(uqaddV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uqaddV4H = form("UqaddV4H")

func (s *Section) UqaddV4H(rd, rn, rm reg.V) {
	s.inst(uqaddV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uqaddV8H = form("UqaddV8H")

func (s *Section) UqaddV8H(rd, rn, rm reg.V) {
	s.inst(uqaddV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uqaddV2S = form("UqaddV2S")

func (s *Section) UqaddV2S(rd, rn, rm reg.V) {
	s.inst(uqaddV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uqaddV4S = form("UqaddV4S")

func (s *Section) UqaddV4S(rd, rn, rm reg.V) {
	s.inst(uqaddV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uqaddV2D = form("UqaddV2D")

func (s *Section) UqaddV2D(rd, rn, rm reg.V) {
	s.inst(uqaddV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var sqsubV8B = form("SqsubV8B")

func (s *Section) SqsubV8B(rd, rn, rm reg.V) {
	s.inst(sqsubV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var sqsubV16B = form("SqsubV16B")

func (s *Section) SqsubV16B(rd, rn, rm reg.V) {
	s.inst(sqsubV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var sqsubV4H = form("SqsubV4H")

func (s *Section) SqsubV4H(rd, rn, rm reg.V) {
	s.inst(sqsubV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sqsubV8H = form("SqsubV8H")

func (s *Section) SqsubV8H(rd, rn, rm reg.V) {
	s.inst(sqsubV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sqsubV2S = form("SqsubV2S")

func (s *Section) SqsubV2S(rd, rn, rm reg.V) {
	s.inst(sqsubV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sqsubV4S = form("SqsubV4S")

func (s *Section) SqsubV4S(rd, rn, rm reg.V) {
	s.inst(sqsubV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sqsubV2D = form("SqsubV2D")

func (s *Section) SqsubV2D(rd, rn, rm reg.V) {
	s.inst(sqsubV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var uqsubV8B = form("UqsubV8B")

func (s *Section) UqsubV8B(rd, rn, rm reg.V) {
	s.inst(uqsubV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uqsubV16B = form("UqsubV16B")

func (s *Section) UqsubV16B(rd, rn, rm reg.V) {
	s.inst(uqsubV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uqsubV4H = form("UqsubV4H")

func (s *Section) UqsubV4H(rd, rn, rm reg.V) {
	s.inst(uqsubV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uqsubV8H = form("UqsubV8H")

func (s *Section) UqsubV8H(rd, rn, rm reg.V) {
	s.inst(uqsubV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uqsubV2S = form("UqsubV2S")

func (s *Section) UqsubV2S(rd, rn, rm reg.V) {
	s.inst(uqsubV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uqsubV4S = form("UqsubV4S")

func (s *Section) UqsubV4S(rd, rn, rm reg.V) {
	s.inst(uqsubV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uqsubV2D = form("UqsubV2D")

func (s *Section) UqsubV2D(rd, rn, rm reg.V) {
	s.inst(uqsubV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var smaxV8B = form("SmaxV8B")

func (s *Section) SmaxV8B(rd, rn, rm reg.V) {
	s.inst(smaxV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var smaxV16B = form("SmaxV16B")

func (s *Section) SmaxV16B(rd, rn, rm reg.V) {
	s.inst(smaxV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var smaxV4H = form("SmaxV4H")

func (s *Section) SmaxV4H(rd, rn, rm reg.V) {
	s.inst(smaxV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var smaxV8H = form("SmaxV8H")

func (s *Section) SmaxV8H(rd, rn, rm reg.V) {
	s.inst(smaxV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var smaxV2S = form("SmaxV2S")

func (s *Section) SmaxV2S(rd, rn, rm reg.V) {
	s.inst(smaxV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var smaxV4S = form("SmaxV4S")

func (s *Section) SmaxV4S(rd, rn, rm reg.V) {
	s.inst(smaxV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var umaxV8B = form("UmaxV8B")

func (s *Section) UmaxV8B(rd, rn, rm reg.V) {
	s.inst(umaxV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var umaxV16B = form("UmaxV16B")

func (s *Section) UmaxV16B(rd, rn, rm reg.V) {
	s.inst(umaxV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var umaxV4H = form("UmaxV4H")

func (s *Section) UmaxV4H(rd, rn, rm reg.V) {
	s.inst(umaxV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var umaxV8H = form("UmaxV8H")

func (s *Section) UmaxV8H(rd, rn, rm reg.V) {
	s.inst(umaxV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var umaxV2S = form("UmaxV2S")

func (s *Section) UmaxV2S(rd, rn, rm reg.V) {
	s.inst(umaxV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var umaxV4S = form("UmaxV4S")

func (s *Section) UmaxV4S(rd, rn, rm reg.V) {
	s.inst(umaxV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sminV8B = form("SminV8B")

func (s *Section) SminV8B(rd, rn, rm reg.V) {
	s.inst(sminV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var sminV16B = form("SminV16B")

func (s *Section) SminV16B(rd, rn, rm reg.V) {
	s.inst(sminV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var sminV4H = form("SminV4H")

func (s *Section) SminV4H(rd, rn, rm reg.V) {
	s.inst(sminV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sminV8H = form("SminV8H")

func (s *Section) SminV8H(rd, rn, rm reg.V) {
	s.inst(sminV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sminV2S = form("SminV2S")

func (s *Section) SminV2S(rd, rn, rm reg.V) {
	s.inst(sminV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sminV4S = form("SminV4S")

func (s *Section) SminV4S(rd, rn, rm reg.V) {
	s.inst(sminV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uminV8B = form("UminV8B")

func (s *Section) UminV8B(rd, rn, rm reg.V) {
	s.inst(uminV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uminV16B = form("UminV16B")

func (s *Section) UminV16B(rd, rn, rm reg.V) {
	s.inst(uminV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uminV4H = form("UminV4H")

func (s *Section) UminV4H(rd, rn, rm reg.V) {
	s.inst(uminV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uminV8H = form("UminV8H")

func (s *Section) UminV8H(rd, rn, rm reg.V) {
	s.inst(uminV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uminV2S = form("UminV2S")

func (s *Section) UminV2S(rd, rn, rm reg.V) {
	s.inst(uminV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uminV4S = form("UminV4S")

func (s *Section) UminV4S(rd, rn, rm reg.V) {
	s.inst(uminV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var srhaddV8B = form("SrhaddV8B")

func (s *Section) SrhaddV8B(rd, rn, rm reg.V) {
	s.inst(srhaddV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var srhaddV16B = form("SrhaddV16B")

func (s *Section) SrhaddV16B(rd, rn, rm reg.V) {
	s.inst(srhaddV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var srhaddV4H = form("SrhaddV4H")

func (s *Section) SrhaddV4H(rd, rn, rm reg.V) {
	s.inst(srhaddV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var srhaddV8H = form("SrhaddV8H")

func (s *Section) SrhaddV8H(rd, rn, rm reg.V) {
	s.inst(srhaddV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var srhaddV2S = form("SrhaddV2S")

func (s *Section) SrhaddV2S(rd, rn, rm reg.V) {
	s.inst(srhaddV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var srhaddV4S = form("SrhaddV4S")

func (s *Section) SrhaddV4S(rd, rn, rm reg.V) {
	s.inst(srhaddV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var urhaddV8B = form("UrhaddV8B")

func (s *Section) UrhaddV8B(rd, rn, rm reg.V) {
	s.inst(urhaddV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var urhaddV16B = form("UrhaddV16B")

func (s *Section) UrhaddV16B(rd, rn, rm reg.V) {
	s.inst(urhaddV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var urhaddV4H = form("UrhaddV4H")

func (s *Section) UrhaddV4H(rd, rn, rm reg.V) {
	s.inst(urhaddV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var urhaddV8H = form("UrhaddV8H")

func (s *Section) UrhaddV8H(rd, rn, rm reg.V) {
	s.inst(urhaddV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var urhaddV2S = form("UrhaddV2S")

func (s *Section) UrhaddV2S(rd, rn, rm reg.V) {
	s.inst(urhaddV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var urhaddV4S = form("UrhaddV4S")

func (s *Section) UrhaddV4S(rd, rn, rm reg.V) {
	s.inst(urhaddV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sabdV8B = form("SabdV8B")

func (s *Section) SabdV8B(rd, rn, rm reg.V) {
	s.inst(sabdV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var sabdV16B = form("SabdV16B")

func (s *Section) SabdV16B(rd, rn, rm reg.V) {
	s.inst(sabdV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var sabdV4H = form("SabdV4H")

func (s *Section) SabdV4H(rd, rn, rm reg.V) {
	s.inst(sabdV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sabdV8H = form("SabdV8H")

func (s *Section) SabdV8H(rd, rn, rm reg.V) {
	s.inst(sabdV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sabdV2S = form("SabdV2S")

func (s *Section) SabdV2S(rd, rn, rm reg.V) {
	s.inst(sabdV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sabdV4S = form("SabdV4S")

func (s *Section) SabdV4S(rd, rn, rm reg.V) {
	s.inst(sabdV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uabdV8B = form("UabdV8B")

func (s *Section) UabdV8B(rd, rn, rm reg.V) {
	s.inst(uabdV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uabdV16B = form("UabdV16B")

func (s *Section) UabdV16B(rd, rn, rm reg.V) {
	s.inst(uabdV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uabdV4H = form("UabdV4H")

func (s *Section) UabdV4H(rd, rn, rm reg.V) {
	s.inst(uabdV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uabdV8H = form("UabdV8H")

func (s *Section) UabdV8H(rd, rn, rm reg.V) {
	s.inst(uabdV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uabdV2S = form("UabdV2S")

func (s *Section) UabdV2S(rd, rn, rm reg.V) {
	s.inst(uabdV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uabdV4S = form("UabdV4S")

func (s *Section) UabdV4S(rd, rn, rm reg.V) {
	s.inst(uabdV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sshlV8B = form("SshlV8B")

func (s *Section) SshlV8B(rd, rn, rm reg.V) {
	s.inst(sshlV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var sshlV16B = form("SshlV16B")

func (s *Section) SshlV16B(rd, rn, rm reg.V) {
	s.inst(sshlV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var sshlV4H = form("SshlV4H")

func (s *Section) SshlV4H(rd, rn, rm reg.V) {
	s.inst(sshlV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sshlV8H = form("SshlV8H")

func (s *Section) SshlV8H(rd, rn, rm reg.V) {
	s.inst(sshlV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sshlV2S = form("SshlV2S")

func (s *Section) SshlV2S(rd, rn, rm reg.V) {
	s.inst(sshlV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sshlV4S = form("SshlV4S")

func (s *Section) SshlV4S(rd, rn, rm reg.V) {
	s.inst(sshlV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sshlV2D = form("SshlV2D")

func (s *Section) SshlV2D(rd, rn, rm reg.V) {
	s.inst(sshlV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var ushlV8B = form("UshlV8B")

func (s *Section) UshlV8B(rd, rn, rm reg.V) {
	s.inst(ushlV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var ushlV16B = form("UshlV16B")

func (s *Section) UshlV16B(rd, rn, rm reg.V) {
	s.inst(ushlV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var ushlV4H = form("UshlV4H")

func (s *Section) UshlV4H(rd, rn, rm reg.V) {
	s.inst(ushlV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var ushlV8H = form("UshlV8H")

func (s *Section) UshlV8H(rd, rn, rm reg.V) {
	s.inst(ushlV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var ushlV2S = form("UshlV2S")

func (s *Section) UshlV2S(rd, rn, rm reg.V) {
	s.inst(ushlV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var ushlV4S = form("UshlV4S")

func (s *Section) UshlV4S(rd, rn, rm reg.V) {
	s.inst(ushlV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var ushlV2D = form("UshlV2D")

func (s *Section) UshlV2D(rd, rn, rm reg.V) {
	s.inst(ushlV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var addpV8B = form("AddpV8B")

func (s *Section) AddpV8B(rd, rn, rm reg.V) {
	s.inst(addpV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var addpV16B = form("AddpV16B")

func (s *Section) AddpV16B(rd, rn, rm reg.V) {
	s.inst(addpV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var addpV4H = form("AddpV4H")

func (s *Section) AddpV4H(rd, rn, rm reg.V) {
	s.inst(addpV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var addpV8H = form("AddpV8H")

func (s *Section) AddpV8H(rd, rn, rm reg.V) {
	s.inst(addpV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var addpV2S = form("AddpV2S")

func (s *Section) AddpV2S(rd, rn, rm reg.V) {
	s.inst(addpV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var addpV4S = form("AddpV4S")

func (s *Section) AddpV4S(rd, rn, rm reg.V) {
	s.inst(addpV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var addpV2D = form("AddpV2D")

func (s *Section) AddpV2D(rd, rn, rm reg.V) {
	s.inst(addpV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var sqdmulhV4H = form("SqdmulhV4H")

func (s *Section) SqdmulhV4H(rd, rn, rm reg.V) {
	s.inst(sqdmulhV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sqdmulhV8H = form("SqdmulhV8H")

func (s *Section) SqdmulhV8H(rd, rn, rm reg.V) {
	s.inst(sqdmulhV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sqdmulhV2S = form("SqdmulhV2S")

func (s *Section) SqdmulhV2S(rd, rn, rm reg.V) {
	s.inst(sqdmulhV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sqdmulhV4S = form("SqdmulhV4S")

func (s *Section) SqdmulhV4S(rd, rn, rm reg.V) {
	s.inst(sqdmulhV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sqrdmulhV4H = form("SqrdmulhV4H")

func (s *Section) SqrdmulhV4H(rd, rn, rm reg.V) {
	s.inst(sqrdmulhV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var sqrdmulhV8H = form("SqrdmulhV8H")

func (s *Section) SqrdmulhV8H(rd, rn, rm reg.V) {
	s.inst(sqrdmulhV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var sqrdmulhV2S = form("SqrdmulhV2S")

func (s *Section) SqrdmulhV2S(rd, rn, rm reg.V) {
	s.inst(sqrdmulhV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var sqrdmulhV4S = form("SqrdmulhV4S")

func (s *Section) SqrdmulhV4S(rd, rn, rm reg.V) {
	s.inst(sqrdmulhV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

// ---- Bitwise, three registers ----

var andV8B = form("AndV8B")

func (s *Section) AndV8B(rd, rn, rm reg.V) {
	s.inst(andV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var andV16B = form("AndV16B")

func (s *Section) AndV16B(rd, rn, rm reg.V) {
	s.inst(andV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var bicV8B = form("BicV8B")

func (s *Section) BicV8B(rd, rn, rm reg.V) {
	s.inst(bicV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var bicV16B = form("BicV16B")

func (s *Section) BicV16B(rd, rn, rm reg.V) {
	s.inst(bicV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var orrV8B = form("OrrV8B")

func (s *Section) OrrV8B(rd, rn, rm reg.V) {
	s.inst(orrV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var orrV16B = form("OrrV16B")

func (s *Section) OrrV16B(rd, rn, rm reg.V) {
	s.inst(orrV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var ornV8B = form("OrnV8B")

func (s *Section) OrnV8B(rd, rn, rm reg.V) {
	s.inst(ornV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var ornV16B = form("OrnV16B")

func (s *Section) OrnV16B(rd, rn, rm reg.V) {
	s.inst(ornV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var eorV8B = form("EorV8B")

func (s *Section) EorV8B(rd, rn, rm reg.V) {
	s.inst(eorV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var eorV16B = form("EorV16B")

func (s *Section) EorV16B(rd, rn, rm reg.V) {
	s.inst(eorV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var bslV8B = form("BslV8B")

func (s *Section) BslV8B(rd, rn, rm reg.V) {
	s.inst(bslV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var bslV16B = form("BslV16B")

func (s *Section) BslV16B(rd, rn, rm reg.V) {
	s.inst(bslV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var bitV8B = form("BitV8B")

func (s *Section) BitV8B(rd, rn, rm reg.V) {
	s.inst(bitV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var bitV16B = form("BitV16B")

func (s *Section) BitV16B(rd, rn, rm reg.V) {
	s.inst(bitV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var bifV8B = form("BifV8B")

func (s *Section) BifV8B(rd, rn, rm reg.V) {
	s.inst(bifV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var bifV16B = form("BifV16B")

func (s *Section) BifV16B(rd, rn, rm reg.V) {
	s.inst(bifV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

// ---- One source register ----

var rev64V8B = form("Rev64V8B")

func (s *Section) Rev64V8B(rd, rn reg.V) { s.inst(rev64V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var rev64V16B = form("Rev64V16B")

func (s *Section) Rev64V16B(rd, rn reg.V) { s.inst(rev64V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var rev64V4H = form("Rev64V4H")

func (s *Section) Rev64V4H(rd, rn reg.V) { s.inst(rev64V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H)) }

var rev64V8H = form("Rev64V8H")

func (s *Section) Rev64V8H(rd, rn reg.V) { s.inst(rev64V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H)) }

var rev64V2S = form("Rev64V2S")

func (s *Section) Rev64V2S(rd, rn reg.V) { s.inst(rev64V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S)) }

var rev64V4S = form("Rev64V4S")

func (s *Section) Rev64V4S(rd, rn reg.V) { s.inst(rev64V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S)) }

var rev32V8B = form("Rev32V8B")

func (s *Section) Rev32V8B(rd, rn reg.V) { s.inst(rev32V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var rev32V16B = form("Rev32V16B")

func (s *Section) Rev32V16B(rd, rn reg.V) { s.inst(rev32V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var rev32V4H = form("Rev32V4H")

func (s *Section) Rev32V4H(rd, rn reg.V) { s.inst(rev32V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H)) }

var rev32V8H = form("Rev32V8H")

func (s *Section) Rev32V8H(rd, rn reg.V) { s.inst(rev32V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H)) }

var rev16V8B = form("Rev16V8B")

func (s *Section) Rev16V8B(rd, rn reg.V) { s.inst(rev16V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var rev16V16B = form("Rev16V16B")

func (s *Section) Rev16V16B(rd, rn reg.V) { s.inst(rev16V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var cntV8B = form("CntV8B")

func (s *Section) CntV8B(rd, rn reg.V) { s.inst(cntV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var cntV16B = form("CntV16B")

func (s *Section) CntV16B(rd, rn reg.V) { s.inst(cntV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var notV8B = form("NotV8B")

func (s *Section) NotV8B(rd, rn reg.V) { s.inst(notV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var notV16B = form("NotV16B")

func (s *Section) NotV16B(rd, rn reg.V) { s.inst(notV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var rbitV8B = form("RbitV8B")

func (s *Section) RbitV8B(rd, rn reg.V) { s.inst(rbitV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var rbitV16B = form("RbitV16B")

func (s *Section) RbitV16B(rd, rn reg.V) { s.inst(rbitV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var negV8B = form("NegV8B")

func (s *Section) NegV8B(rd, rn reg.V) { s.inst(negV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var negV16B = form("NegV16B")

func (s *Section) NegV16B(rd, rn reg.V) { s.inst(negV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var negV4H = form("NegV4H")

func (s *Section) NegV4H(rd, rn reg.V) { s.inst(negV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H)) }

var negV8H = form("NegV8H")

func (s *Section) NegV8H(rd, rn reg.V) { s.inst(negV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H)) }

var negV2S = form("NegV2S")

func (s *Section) NegV2S(rd, rn reg.V) { s.inst(negV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S)) }

var negV4S = form("NegV4S")

func (s *Section) NegV4S(rd, rn reg.V) { s.inst(negV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S)) }

var negV2D = form("NegV2D")

func (s *Section) NegV2D(rd, rn reg.V) { s.inst(negV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D)) }

var absV8B = form("AbsV8B")

func (s *Section) AbsV8B(rd, rn reg.V) { s.inst(absV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B)) }

var absV16B = form("AbsV16B")

func (s *Section) AbsV16B(rd, rn reg.V) { s.inst(absV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var absV4H = form("AbsV4H")

func (s *Section) AbsV4H(rd, rn reg.V) { s.inst(absV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H)) }

var absV8H = form("AbsV8H")

func (s *Section) AbsV8H(rd, rn reg.V) { s.inst(absV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H)) }

var absV2S = form("AbsV2S")

func (s *Section) AbsV2S(rd, rn reg.V) { s.inst(absV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S)) }

var absV4S = form("AbsV4S")

func (s *Section) AbsV4S(rd, rn reg.V) { s.inst(absV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S)) }

var absV2D = form("AbsV2D")

func (s *Section) AbsV2D(rd, rn reg.V) { s.inst(absV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D)) }

// ---- Narrowing and pairwise widening ----

var xtnV8B = form("XtnV8B")

func (s *Section) XtnV8B(rd, rn reg.V) { s.inst(xtnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H)) }

var xtn2V16B = form("Xtn2V16B")

func (s *Section) Xtn2V16B(rd, rn reg.V) { s.inst(xtn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H)) }

var xtnV4H = form("XtnV4H")

func (s *Section) XtnV4H(rd, rn reg.V) { s.inst(xtnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S)) }

var xtn2V8H = form("Xtn2V8H")

func (s *Section) Xtn2V8H(rd, rn reg.V) { s.inst(xtn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S)) }

var xtnV2S = form("XtnV2S")

func (s *Section) XtnV2S(rd, rn reg.V) { s.inst(xtnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D)) }

var xtn2V4S = form("Xtn2V4S")

func (s *Section) Xtn2V4S(rd, rn reg.V) { s.inst(xtn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D)) }

var sqxtnV8B = form("SqxtnV8B")

func (s *Section) SqxtnV8B(rd, rn reg.V) { s.inst(sqxtnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H)) }

var sqxtn2V16B = form("Sqxtn2V16B")

func (s *Section) Sqxtn2V16B(rd, rn reg.V) { s.inst(sqxtn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H)) }

var sqxtnV4H = form("SqxtnV4H")

func (s *Section) SqxtnV4H(rd, rn reg.V) { s.inst(sqxtnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S)) }

var sqxtn2V8H = form("Sqxtn2V8H")

func (s *Section) Sqxtn2V8H(rd, rn reg.V) { s.inst(sqxtn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S)) }

var sqxtnV2S = form("SqxtnV2S")

func (s *Section) SqxtnV2S(rd, rn reg.V) { s.inst(sqxtnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D)) }

var sqxtn2V4S = form("Sqxtn2V4S")

func (s *Section) Sqxtn2V4S(rd, rn reg.V) { s.inst(sqxtn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D)) }

var uqxtnV8B = form("UqxtnV8B")

func (s *Section) UqxtnV8B(rd, rn reg.V) { s.inst(uqxtnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H)) }

var uqxtn2V16B = form("Uqxtn2V16B")

func (s *Section) Uqxtn2V16B(rd, rn reg.V) { s.inst(uqxtn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H)) }

var uqxtnV4H = form("UqxtnV4H")

func (s *Section) UqxtnV4H(rd, rn reg.V) { s.inst(uqxtnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S)) }

var uqxtn2V8H = form("Uqxtn2V8H")

func (s *Section) Uqxtn2V8H(rd, rn reg.V) { s.inst(uqxtn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S)) }

var uqxtnV2S = form("UqxtnV2S")

func (s *Section) UqxtnV2S(rd, rn reg.V) { s.inst(uqxtnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D)) }

var uqxtn2V4S = form("Uqxtn2V4S")

func (s *Section) Uqxtn2V4S(rd, rn reg.V) { s.inst(uqxtn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D)) }

var sqxtunV8B = form("SqxtunV8B")

func (s *Section) SqxtunV8B(rd, rn reg.V) { s.inst(sqxtunV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H)) }

var sqxtun2V16B = form("Sqxtun2V16B")

func (s *Section) Sqxtun2V16B(rd, rn reg.V) { s.inst(sqxtun2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H)) }

var sqxtunV4H = form("SqxtunV4H")

func (s *Section) SqxtunV4H(rd, rn reg.V) { s.inst(sqxtunV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S)) }

var sqxtun2V8H = form("Sqxtun2V8H")

func (s *Section) Sqxtun2V8H(rd, rn reg.V) { s.inst(sqxtun2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S)) }

var sqxtunV2S = form("SqxtunV2S")

func (s *Section) SqxtunV2S(rd, rn reg.V) { s.inst(sqxtunV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D)) }

var sqxtun2V4S = form("Sqxtun2V4S")

func (s *Section) Sqxtun2V4S(rd, rn reg.V) { s.inst(sqxtun2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D)) }

var uaddlpV4H = form("UaddlpV4H")

func (s *Section) UaddlpV4H(rd, rn reg.V) { s.inst(uaddlpV4H, rd.Arr(reg.V4H), rn.Arr(reg.V8B)) }

var uaddlpV8H = form("UaddlpV8H")

func (s *Section) UaddlpV8H(rd, rn reg.V) { s.inst(uaddlpV8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B)) }

var uaddlpV2S = form("UaddlpV2S")

func (s *Section) UaddlpV2S(rd, rn reg.V) { s.inst(uaddlpV2S, rd.Arr(reg.V2S), rn.Arr(reg.V4H)) }

var uaddlpV4S = form("UaddlpV4S")

func (s *Section) UaddlpV4S(rd, rn reg.V) { s.inst(uaddlpV4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H)) }

var uaddlpV1D = form("UaddlpV1D")

func (s *Section) UaddlpV1D(rd, rn reg.V) { s.inst(uaddlpV1D, rd.Arr(reg.V1D), rn.Arr(reg.V2S)) }

var uaddlpV2D = form("UaddlpV2D")

func (s *Section) UaddlpV2D(rd, rn reg.V) { s.inst(uaddlpV2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S)) }

var saddlpV4H = form("SaddlpV4H")

func (s *Section) SaddlpV4H(rd, rn reg.V) { s.inst(saddlpV4H, rd.Arr(reg.V4H), rn.Arr(reg.V8B)) }

var saddlpV8H = form("SaddlpV8H")

func (s *Section) SaddlpV8H(rd, rn reg.V) { s.inst(saddlpV8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B)) }

var saddlpV2S = form("SaddlpV2S")

func (s *Section) SaddlpV2S(rd, rn reg.V) { s.inst(saddlpV2S, rd.Arr(reg.V2S), rn.Arr(reg.V4H)) }

var saddlpV4S = form("SaddlpV4S")

func (s *Section) SaddlpV4S(rd, rn reg.V) { s.inst(saddlpV4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H)) }

var saddlpV1D = form("SaddlpV1D")

func (s *Section) SaddlpV1D(rd, rn reg.V) { s.inst(saddlpV1D, rd.Arr(reg.V1D), rn.Arr(reg.V2S)) }

var saddlpV2D = form("SaddlpV2D")

func (s *Section) SaddlpV2D(rd, rn reg.V) { s.inst(saddlpV2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S)) }

// ---- Across all lanes, into a scalar ----

var addvV8B = form("AddvV8B")

func (s *Section) AddvV8B(rd reg.B, rn reg.V) { s.inst(addvV8B, rd, rn.Arr(reg.V8B)) }

var addvV16B = form("AddvV16B")

func (s *Section) AddvV16B(rd reg.B, rn reg.V) { s.inst(addvV16B, rd, rn.Arr(reg.V16B)) }

var addvV4H = form("AddvV4H")

func (s *Section) AddvV4H(rd reg.H, rn reg.V) { s.inst(addvV4H, rd, rn.Arr(reg.V4H)) }

var addvV8H = form("AddvV8H")

func (s *Section) AddvV8H(rd reg.H, rn reg.V) { s.inst(addvV8H, rd, rn.Arr(reg.V8H)) }

var addvV4S = form("AddvV4S")

func (s *Section) AddvV4S(rd reg.S, rn reg.V) { s.inst(addvV4S, rd, rn.Arr(reg.V4S)) }

var umaxvV8B = form("UmaxvV8B")

func (s *Section) UmaxvV8B(rd reg.B, rn reg.V) { s.inst(umaxvV8B, rd, rn.Arr(reg.V8B)) }

var umaxvV16B = form("UmaxvV16B")

func (s *Section) UmaxvV16B(rd reg.B, rn reg.V) { s.inst(umaxvV16B, rd, rn.Arr(reg.V16B)) }

var umaxvV4H = form("UmaxvV4H")

func (s *Section) UmaxvV4H(rd reg.H, rn reg.V) { s.inst(umaxvV4H, rd, rn.Arr(reg.V4H)) }

var umaxvV8H = form("UmaxvV8H")

func (s *Section) UmaxvV8H(rd reg.H, rn reg.V) { s.inst(umaxvV8H, rd, rn.Arr(reg.V8H)) }

var umaxvV4S = form("UmaxvV4S")

func (s *Section) UmaxvV4S(rd reg.S, rn reg.V) { s.inst(umaxvV4S, rd, rn.Arr(reg.V4S)) }

var uminvV8B = form("UminvV8B")

func (s *Section) UminvV8B(rd reg.B, rn reg.V) { s.inst(uminvV8B, rd, rn.Arr(reg.V8B)) }

var uminvV16B = form("UminvV16B")

func (s *Section) UminvV16B(rd reg.B, rn reg.V) { s.inst(uminvV16B, rd, rn.Arr(reg.V16B)) }

var uminvV4H = form("UminvV4H")

func (s *Section) UminvV4H(rd reg.H, rn reg.V) { s.inst(uminvV4H, rd, rn.Arr(reg.V4H)) }

var uminvV8H = form("UminvV8H")

func (s *Section) UminvV8H(rd reg.H, rn reg.V) { s.inst(uminvV8H, rd, rn.Arr(reg.V8H)) }

var uminvV4S = form("UminvV4S")

func (s *Section) UminvV4S(rd reg.S, rn reg.V) { s.inst(uminvV4S, rd, rn.Arr(reg.V4S)) }

var smaxvV8B = form("SmaxvV8B")

func (s *Section) SmaxvV8B(rd reg.B, rn reg.V) { s.inst(smaxvV8B, rd, rn.Arr(reg.V8B)) }

var smaxvV16B = form("SmaxvV16B")

func (s *Section) SmaxvV16B(rd reg.B, rn reg.V) { s.inst(smaxvV16B, rd, rn.Arr(reg.V16B)) }

var smaxvV4H = form("SmaxvV4H")

func (s *Section) SmaxvV4H(rd reg.H, rn reg.V) { s.inst(smaxvV4H, rd, rn.Arr(reg.V4H)) }

var smaxvV8H = form("SmaxvV8H")

func (s *Section) SmaxvV8H(rd reg.H, rn reg.V) { s.inst(smaxvV8H, rd, rn.Arr(reg.V8H)) }

var smaxvV4S = form("SmaxvV4S")

func (s *Section) SmaxvV4S(rd reg.S, rn reg.V) { s.inst(smaxvV4S, rd, rn.Arr(reg.V4S)) }

var sminvV8B = form("SminvV8B")

func (s *Section) SminvV8B(rd reg.B, rn reg.V) { s.inst(sminvV8B, rd, rn.Arr(reg.V8B)) }

var sminvV16B = form("SminvV16B")

func (s *Section) SminvV16B(rd reg.B, rn reg.V) { s.inst(sminvV16B, rd, rn.Arr(reg.V16B)) }

var sminvV4H = form("SminvV4H")

func (s *Section) SminvV4H(rd reg.H, rn reg.V) { s.inst(sminvV4H, rd, rn.Arr(reg.V4H)) }

var sminvV8H = form("SminvV8H")

func (s *Section) SminvV8H(rd reg.H, rn reg.V) { s.inst(sminvV8H, rd, rn.Arr(reg.V8H)) }

var sminvV4S = form("SminvV4S")

func (s *Section) SminvV4S(rd reg.S, rn reg.V) { s.inst(sminvV4S, rd, rn.Arr(reg.V4S)) }

var uaddlvV8B = form("UaddlvV8B")

func (s *Section) UaddlvV8B(rd reg.H, rn reg.V) { s.inst(uaddlvV8B, rd, rn.Arr(reg.V8B)) }

var uaddlvV16B = form("UaddlvV16B")

func (s *Section) UaddlvV16B(rd reg.H, rn reg.V) { s.inst(uaddlvV16B, rd, rn.Arr(reg.V16B)) }

var uaddlvV4H = form("UaddlvV4H")

func (s *Section) UaddlvV4H(rd reg.S, rn reg.V) { s.inst(uaddlvV4H, rd, rn.Arr(reg.V4H)) }

var uaddlvV8H = form("UaddlvV8H")

func (s *Section) UaddlvV8H(rd reg.S, rn reg.V) { s.inst(uaddlvV8H, rd, rn.Arr(reg.V8H)) }

var uaddlvV4S = form("UaddlvV4S")

func (s *Section) UaddlvV4S(rd reg.D, rn reg.V) { s.inst(uaddlvV4S, rd, rn.Arr(reg.V4S)) }

var saddlvV8B = form("SaddlvV8B")

func (s *Section) SaddlvV8B(rd reg.H, rn reg.V) { s.inst(saddlvV8B, rd, rn.Arr(reg.V8B)) }

var saddlvV16B = form("SaddlvV16B")

func (s *Section) SaddlvV16B(rd reg.H, rn reg.V) { s.inst(saddlvV16B, rd, rn.Arr(reg.V16B)) }

var saddlvV4H = form("SaddlvV4H")

func (s *Section) SaddlvV4H(rd reg.S, rn reg.V) { s.inst(saddlvV4H, rd, rn.Arr(reg.V4H)) }

var saddlvV8H = form("SaddlvV8H")

func (s *Section) SaddlvV8H(rd reg.S, rn reg.V) { s.inst(saddlvV8H, rd, rn.Arr(reg.V8H)) }

var saddlvV4S = form("SaddlvV4S")

func (s *Section) SaddlvV4S(rd reg.D, rn reg.V) { s.inst(saddlvV4S, rd, rn.Arr(reg.V4S)) }

// ---- Shifts by an immediate ----

var shlV8B = form("ShlV8B")

func (s *Section) ShlV8B(rd, rn reg.V, shift uint8) {
	s.inst(shlV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var shlV16B = form("ShlV16B")

func (s *Section) ShlV16B(rd, rn reg.V, shift uint8) {
	s.inst(shlV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var shlV4H = form("ShlV4H")

func (s *Section) ShlV4H(rd, rn reg.V, shift uint8) {
	s.inst(shlV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var shlV8H = form("ShlV8H")

func (s *Section) ShlV8H(rd, rn reg.V, shift uint8) {
	s.inst(shlV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var shlV2S = form("ShlV2S")

func (s *Section) ShlV2S(rd, rn reg.V, shift uint8) {
	s.inst(shlV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var shlV4S = form("ShlV4S")

func (s *Section) ShlV4S(rd, rn reg.V, shift uint8) {
	s.inst(shlV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var shlV2D = form("ShlV2D")

func (s *Section) ShlV2D(rd, rn reg.V, shift uint8) {
	s.inst(shlV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var sliV8B = form("SliV8B")

func (s *Section) SliV8B(rd, rn reg.V, shift uint8) {
	s.inst(sliV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var sliV16B = form("SliV16B")

func (s *Section) SliV16B(rd, rn reg.V, shift uint8) {
	s.inst(sliV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var sliV4H = form("SliV4H")

func (s *Section) SliV4H(rd, rn reg.V, shift uint8) {
	s.inst(sliV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var sliV8H = form("SliV8H")

func (s *Section) SliV8H(rd, rn reg.V, shift uint8) {
	s.inst(sliV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var sliV2S = form("SliV2S")

func (s *Section) SliV2S(rd, rn reg.V, shift uint8) {
	s.inst(sliV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var sliV4S = form("SliV4S")

func (s *Section) SliV4S(rd, rn reg.V, shift uint8) {
	s.inst(sliV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var sliV2D = form("SliV2D")

func (s *Section) SliV2D(rd, rn reg.V, shift uint8) {
	s.inst(sliV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var ushrV8B = form("UshrV8B")

func (s *Section) UshrV8B(rd, rn reg.V, shift uint8) {
	s.inst(ushrV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var ushrV16B = form("UshrV16B")

func (s *Section) UshrV16B(rd, rn reg.V, shift uint8) {
	s.inst(ushrV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var ushrV4H = form("UshrV4H")

func (s *Section) UshrV4H(rd, rn reg.V, shift uint8) {
	s.inst(ushrV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var ushrV8H = form("UshrV8H")

func (s *Section) UshrV8H(rd, rn reg.V, shift uint8) {
	s.inst(ushrV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var ushrV2S = form("UshrV2S")

func (s *Section) UshrV2S(rd, rn reg.V, shift uint8) {
	s.inst(ushrV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var ushrV4S = form("UshrV4S")

func (s *Section) UshrV4S(rd, rn reg.V, shift uint8) {
	s.inst(ushrV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var ushrV2D = form("UshrV2D")

func (s *Section) UshrV2D(rd, rn reg.V, shift uint8) {
	s.inst(ushrV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var sshrV8B = form("SshrV8B")

func (s *Section) SshrV8B(rd, rn reg.V, shift uint8) {
	s.inst(sshrV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var sshrV16B = form("SshrV16B")

func (s *Section) SshrV16B(rd, rn reg.V, shift uint8) {
	s.inst(sshrV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var sshrV4H = form("SshrV4H")

func (s *Section) SshrV4H(rd, rn reg.V, shift uint8) {
	s.inst(sshrV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var sshrV8H = form("SshrV8H")

func (s *Section) SshrV8H(rd, rn reg.V, shift uint8) {
	s.inst(sshrV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var sshrV2S = form("SshrV2S")

func (s *Section) SshrV2S(rd, rn reg.V, shift uint8) {
	s.inst(sshrV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var sshrV4S = form("SshrV4S")

func (s *Section) SshrV4S(rd, rn reg.V, shift uint8) {
	s.inst(sshrV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var sshrV2D = form("SshrV2D")

func (s *Section) SshrV2D(rd, rn reg.V, shift uint8) {
	s.inst(sshrV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var usraV8B = form("UsraV8B")

func (s *Section) UsraV8B(rd, rn reg.V, shift uint8) {
	s.inst(usraV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var usraV16B = form("UsraV16B")

func (s *Section) UsraV16B(rd, rn reg.V, shift uint8) {
	s.inst(usraV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var usraV4H = form("UsraV4H")

func (s *Section) UsraV4H(rd, rn reg.V, shift uint8) {
	s.inst(usraV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var usraV8H = form("UsraV8H")

func (s *Section) UsraV8H(rd, rn reg.V, shift uint8) {
	s.inst(usraV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var usraV2S = form("UsraV2S")

func (s *Section) UsraV2S(rd, rn reg.V, shift uint8) {
	s.inst(usraV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var usraV4S = form("UsraV4S")

func (s *Section) UsraV4S(rd, rn reg.V, shift uint8) {
	s.inst(usraV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var usraV2D = form("UsraV2D")

func (s *Section) UsraV2D(rd, rn reg.V, shift uint8) {
	s.inst(usraV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var ssraV8B = form("SsraV8B")

func (s *Section) SsraV8B(rd, rn reg.V, shift uint8) {
	s.inst(ssraV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var ssraV16B = form("SsraV16B")

func (s *Section) SsraV16B(rd, rn reg.V, shift uint8) {
	s.inst(ssraV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var ssraV4H = form("SsraV4H")

func (s *Section) SsraV4H(rd, rn reg.V, shift uint8) {
	s.inst(ssraV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var ssraV8H = form("SsraV8H")

func (s *Section) SsraV8H(rd, rn reg.V, shift uint8) {
	s.inst(ssraV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var ssraV2S = form("SsraV2S")

func (s *Section) SsraV2S(rd, rn reg.V, shift uint8) {
	s.inst(ssraV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var ssraV4S = form("SsraV4S")

func (s *Section) SsraV4S(rd, rn reg.V, shift uint8) {
	s.inst(ssraV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var ssraV2D = form("SsraV2D")

func (s *Section) SsraV2D(rd, rn reg.V, shift uint8) {
	s.inst(ssraV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var urshrV8B = form("UrshrV8B")

func (s *Section) UrshrV8B(rd, rn reg.V, shift uint8) {
	s.inst(urshrV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var urshrV16B = form("UrshrV16B")

func (s *Section) UrshrV16B(rd, rn reg.V, shift uint8) {
	s.inst(urshrV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var urshrV4H = form("UrshrV4H")

func (s *Section) UrshrV4H(rd, rn reg.V, shift uint8) {
	s.inst(urshrV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var urshrV8H = form("UrshrV8H")

func (s *Section) UrshrV8H(rd, rn reg.V, shift uint8) {
	s.inst(urshrV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var urshrV2S = form("UrshrV2S")

func (s *Section) UrshrV2S(rd, rn reg.V, shift uint8) {
	s.inst(urshrV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var urshrV4S = form("UrshrV4S")

func (s *Section) UrshrV4S(rd, rn reg.V, shift uint8) {
	s.inst(urshrV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var urshrV2D = form("UrshrV2D")

func (s *Section) UrshrV2D(rd, rn reg.V, shift uint8) {
	s.inst(urshrV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var srshrV8B = form("SrshrV8B")

func (s *Section) SrshrV8B(rd, rn reg.V, shift uint8) {
	s.inst(srshrV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var srshrV16B = form("SrshrV16B")

func (s *Section) SrshrV16B(rd, rn reg.V, shift uint8) {
	s.inst(srshrV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var srshrV4H = form("SrshrV4H")

func (s *Section) SrshrV4H(rd, rn reg.V, shift uint8) {
	s.inst(srshrV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var srshrV8H = form("SrshrV8H")

func (s *Section) SrshrV8H(rd, rn reg.V, shift uint8) {
	s.inst(srshrV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var srshrV2S = form("SrshrV2S")

func (s *Section) SrshrV2S(rd, rn reg.V, shift uint8) {
	s.inst(srshrV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var srshrV4S = form("SrshrV4S")

func (s *Section) SrshrV4S(rd, rn reg.V, shift uint8) {
	s.inst(srshrV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var srshrV2D = form("SrshrV2D")

func (s *Section) SrshrV2D(rd, rn reg.V, shift uint8) {
	s.inst(srshrV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var sriV8B = form("SriV8B")

func (s *Section) SriV8B(rd, rn reg.V, shift uint8) {
	s.inst(sriV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), uint64(shift))
}

var sriV16B = form("SriV16B")

func (s *Section) SriV16B(rd, rn reg.V, shift uint8) {
	s.inst(sriV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), uint64(shift))
}

var sriV4H = form("SriV4H")

func (s *Section) SriV4H(rd, rn reg.V, shift uint8) {
	s.inst(sriV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), uint64(shift))
}

var sriV8H = form("SriV8H")

func (s *Section) SriV8H(rd, rn reg.V, shift uint8) {
	s.inst(sriV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), uint64(shift))
}

var sriV2S = form("SriV2S")

func (s *Section) SriV2S(rd, rn reg.V, shift uint8) {
	s.inst(sriV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), uint64(shift))
}

var sriV4S = form("SriV4S")

func (s *Section) SriV4S(rd, rn reg.V, shift uint8) {
	s.inst(sriV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), uint64(shift))
}

var sriV2D = form("SriV2D")

func (s *Section) SriV2D(rd, rn reg.V, shift uint8) {
	s.inst(sriV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), uint64(shift))
}

var shrnV8B = form("ShrnV8B")

func (s *Section) ShrnV8B(rd, rn reg.V, shift uint8) {
	s.inst(shrnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H), uint64(shift))
}

var shrn2V16B = form("Shrn2V16B")

func (s *Section) Shrn2V16B(rd, rn reg.V, shift uint8) {
	s.inst(shrn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H), uint64(shift))
}

var shrnV4H = form("ShrnV4H")

func (s *Section) ShrnV4H(rd, rn reg.V, shift uint8) {
	s.inst(shrnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S), uint64(shift))
}

var shrn2V8H = form("Shrn2V8H")

func (s *Section) Shrn2V8H(rd, rn reg.V, shift uint8) {
	s.inst(shrn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S), uint64(shift))
}

var shrnV2S = form("ShrnV2S")

func (s *Section) ShrnV2S(rd, rn reg.V, shift uint8) {
	s.inst(shrnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D), uint64(shift))
}

var shrn2V4S = form("Shrn2V4S")

func (s *Section) Shrn2V4S(rd, rn reg.V, shift uint8) {
	s.inst(shrn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D), uint64(shift))
}

var rshrnV8B = form("RshrnV8B")

func (s *Section) RshrnV8B(rd, rn reg.V, shift uint8) {
	s.inst(rshrnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H), uint64(shift))
}

var rshrn2V16B = form("Rshrn2V16B")

func (s *Section) Rshrn2V16B(rd, rn reg.V, shift uint8) {
	s.inst(rshrn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H), uint64(shift))
}

var rshrnV4H = form("RshrnV4H")

func (s *Section) RshrnV4H(rd, rn reg.V, shift uint8) {
	s.inst(rshrnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S), uint64(shift))
}

var rshrn2V8H = form("Rshrn2V8H")

func (s *Section) Rshrn2V8H(rd, rn reg.V, shift uint8) {
	s.inst(rshrn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S), uint64(shift))
}

var rshrnV2S = form("RshrnV2S")

func (s *Section) RshrnV2S(rd, rn reg.V, shift uint8) {
	s.inst(rshrnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D), uint64(shift))
}

var rshrn2V4S = form("Rshrn2V4S")

func (s *Section) Rshrn2V4S(rd, rn reg.V, shift uint8) {
	s.inst(rshrn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D), uint64(shift))
}

var sqshrnV8B = form("SqshrnV8B")

func (s *Section) SqshrnV8B(rd, rn reg.V, shift uint8) {
	s.inst(sqshrnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H), uint64(shift))
}

var sqshrn2V16B = form("Sqshrn2V16B")

func (s *Section) Sqshrn2V16B(rd, rn reg.V, shift uint8) {
	s.inst(sqshrn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H), uint64(shift))
}

var sqshrnV4H = form("SqshrnV4H")

func (s *Section) SqshrnV4H(rd, rn reg.V, shift uint8) {
	s.inst(sqshrnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S), uint64(shift))
}

var sqshrn2V8H = form("Sqshrn2V8H")

func (s *Section) Sqshrn2V8H(rd, rn reg.V, shift uint8) {
	s.inst(sqshrn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S), uint64(shift))
}

var sqshrnV2S = form("SqshrnV2S")

func (s *Section) SqshrnV2S(rd, rn reg.V, shift uint8) {
	s.inst(sqshrnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D), uint64(shift))
}

var sqshrn2V4S = form("Sqshrn2V4S")

func (s *Section) Sqshrn2V4S(rd, rn reg.V, shift uint8) {
	s.inst(sqshrn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D), uint64(shift))
}

var uqshrnV8B = form("UqshrnV8B")

func (s *Section) UqshrnV8B(rd, rn reg.V, shift uint8) {
	s.inst(uqshrnV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H), uint64(shift))
}

var uqshrn2V16B = form("Uqshrn2V16B")

func (s *Section) Uqshrn2V16B(rd, rn reg.V, shift uint8) {
	s.inst(uqshrn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H), uint64(shift))
}

var uqshrnV4H = form("UqshrnV4H")

func (s *Section) UqshrnV4H(rd, rn reg.V, shift uint8) {
	s.inst(uqshrnV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S), uint64(shift))
}

var uqshrn2V8H = form("Uqshrn2V8H")

func (s *Section) Uqshrn2V8H(rd, rn reg.V, shift uint8) {
	s.inst(uqshrn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S), uint64(shift))
}

var uqshrnV2S = form("UqshrnV2S")

func (s *Section) UqshrnV2S(rd, rn reg.V, shift uint8) {
	s.inst(uqshrnV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D), uint64(shift))
}

var uqshrn2V4S = form("Uqshrn2V4S")

func (s *Section) Uqshrn2V4S(rd, rn reg.V, shift uint8) {
	s.inst(uqshrn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D), uint64(shift))
}

var sqshrunV8B = form("SqshrunV8B")

func (s *Section) SqshrunV8B(rd, rn reg.V, shift uint8) {
	s.inst(sqshrunV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8H), uint64(shift))
}

var sqshrun2V16B = form("Sqshrun2V16B")

func (s *Section) Sqshrun2V16B(rd, rn reg.V, shift uint8) {
	s.inst(sqshrun2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V8H), uint64(shift))
}

var sqshrunV4H = form("SqshrunV4H")

func (s *Section) SqshrunV4H(rd, rn reg.V, shift uint8) {
	s.inst(sqshrunV4H, rd.Arr(reg.V4H), rn.Arr(reg.V4S), uint64(shift))
}

var sqshrun2V8H = form("Sqshrun2V8H")

func (s *Section) Sqshrun2V8H(rd, rn reg.V, shift uint8) {
	s.inst(sqshrun2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V4S), uint64(shift))
}

var sqshrunV2S = form("SqshrunV2S")

func (s *Section) SqshrunV2S(rd, rn reg.V, shift uint8) {
	s.inst(sqshrunV2S, rd.Arr(reg.V2S), rn.Arr(reg.V2D), uint64(shift))
}

var sqshrun2V4S = form("Sqshrun2V4S")

func (s *Section) Sqshrun2V4S(rd, rn reg.V, shift uint8) {
	s.inst(sqshrun2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V2D), uint64(shift))
}

var ushllV8H = form("UshllV8H")

func (s *Section) UshllV8H(rd, rn reg.V, shift uint8) {
	s.inst(ushllV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), uint64(shift))
}

var ushll2V8H = form("Ushll2V8H")

func (s *Section) Ushll2V8H(rd, rn reg.V, shift uint8) {
	s.inst(ushll2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), uint64(shift))
}

var ushllV4S = form("UshllV4S")

func (s *Section) UshllV4S(rd, rn reg.V, shift uint8) {
	s.inst(ushllV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), uint64(shift))
}

var ushll2V4S = form("Ushll2V4S")

func (s *Section) Ushll2V4S(rd, rn reg.V, shift uint8) {
	s.inst(ushll2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), uint64(shift))
}

var ushllV2D = form("UshllV2D")

func (s *Section) UshllV2D(rd, rn reg.V, shift uint8) {
	s.inst(ushllV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), uint64(shift))
}

var ushll2V2D = form("Ushll2V2D")

func (s *Section) Ushll2V2D(rd, rn reg.V, shift uint8) {
	s.inst(ushll2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), uint64(shift))
}

var sshllV8H = form("SshllV8H")

func (s *Section) SshllV8H(rd, rn reg.V, shift uint8) {
	s.inst(sshllV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), uint64(shift))
}

var sshll2V8H = form("Sshll2V8H")

func (s *Section) Sshll2V8H(rd, rn reg.V, shift uint8) {
	s.inst(sshll2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), uint64(shift))
}

var sshllV4S = form("SshllV4S")

func (s *Section) SshllV4S(rd, rn reg.V, shift uint8) {
	s.inst(sshllV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), uint64(shift))
}

var sshll2V4S = form("Sshll2V4S")

func (s *Section) Sshll2V4S(rd, rn reg.V, shift uint8) {
	s.inst(sshll2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), uint64(shift))
}

var sshllV2D = form("SshllV2D")

func (s *Section) SshllV2D(rd, rn reg.V, shift uint8) {
	s.inst(sshllV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), uint64(shift))
}

var sshll2V2D = form("Sshll2V2D")

func (s *Section) Sshll2V2D(rd, rn reg.V, shift uint8) {
	s.inst(sshll2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), uint64(shift))
}

// ---- Widening, three registers ----

var umullV8H = form("UmullV8H")

func (s *Section) UmullV8H(rd, rn, rm reg.V) {
	s.inst(umullV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var umull2V8H = form("Umull2V8H")

func (s *Section) Umull2V8H(rd, rn, rm reg.V) {
	s.inst(umull2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var umullV4S = form("UmullV4S")

func (s *Section) UmullV4S(rd, rn, rm reg.V) {
	s.inst(umullV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var umull2V4S = form("Umull2V4S")

func (s *Section) Umull2V4S(rd, rn, rm reg.V) {
	s.inst(umull2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var umullV2D = form("UmullV2D")

func (s *Section) UmullV2D(rd, rn, rm reg.V) {
	s.inst(umullV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var umull2V2D = form("Umull2V2D")

func (s *Section) Umull2V2D(rd, rn, rm reg.V) {
	s.inst(umull2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var smullV8H = form("SmullV8H")

func (s *Section) SmullV8H(rd, rn, rm reg.V) {
	s.inst(smullV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var smull2V8H = form("Smull2V8H")

func (s *Section) Smull2V8H(rd, rn, rm reg.V) {
	s.inst(smull2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var smullV4S = form("SmullV4S")

func (s *Section) SmullV4S(rd, rn, rm reg.V) {
	s.inst(smullV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var smull2V4S = form("Smull2V4S")

func (s *Section) Smull2V4S(rd, rn, rm reg.V) {
	s.inst(smull2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var smullV2D = form("SmullV2D")

func (s *Section) SmullV2D(rd, rn, rm reg.V) {
	s.inst(smullV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var smull2V2D = form("Smull2V2D")

func (s *Section) Smull2V2D(rd, rn, rm reg.V) {
	s.inst(smull2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var umlalV8H = form("UmlalV8H")

func (s *Section) UmlalV8H(rd, rn, rm reg.V) {
	s.inst(umlalV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var umlal2V8H = form("Umlal2V8H")

func (s *Section) Umlal2V8H(rd, rn, rm reg.V) {
	s.inst(umlal2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var umlalV4S = form("UmlalV4S")

func (s *Section) UmlalV4S(rd, rn, rm reg.V) {
	s.inst(umlalV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var umlal2V4S = form("Umlal2V4S")

func (s *Section) Umlal2V4S(rd, rn, rm reg.V) {
	s.inst(umlal2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var umlalV2D = form("UmlalV2D")

func (s *Section) UmlalV2D(rd, rn, rm reg.V) {
	s.inst(umlalV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var umlal2V2D = form("Umlal2V2D")

func (s *Section) Umlal2V2D(rd, rn, rm reg.V) {
	s.inst(umlal2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var smlalV8H = form("SmlalV8H")

func (s *Section) SmlalV8H(rd, rn, rm reg.V) {
	s.inst(smlalV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var smlal2V8H = form("Smlal2V8H")

func (s *Section) Smlal2V8H(rd, rn, rm reg.V) {
	s.inst(smlal2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var smlalV4S = form("SmlalV4S")

func (s *Section) SmlalV4S(rd, rn, rm reg.V) {
	s.inst(smlalV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var smlal2V4S = form("Smlal2V4S")

func (s *Section) Smlal2V4S(rd, rn, rm reg.V) {
	s.inst(smlal2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var smlalV2D = form("SmlalV2D")

func (s *Section) SmlalV2D(rd, rn, rm reg.V) {
	s.inst(smlalV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var smlal2V2D = form("Smlal2V2D")

func (s *Section) Smlal2V2D(rd, rn, rm reg.V) {
	s.inst(smlal2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uaddlV8H = form("UaddlV8H")

func (s *Section) UaddlV8H(rd, rn, rm reg.V) {
	s.inst(uaddlV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uaddl2V8H = form("Uaddl2V8H")

func (s *Section) Uaddl2V8H(rd, rn, rm reg.V) {
	s.inst(uaddl2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uaddlV4S = form("UaddlV4S")

func (s *Section) UaddlV4S(rd, rn, rm reg.V) {
	s.inst(uaddlV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uaddl2V4S = form("Uaddl2V4S")

func (s *Section) Uaddl2V4S(rd, rn, rm reg.V) {
	s.inst(uaddl2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uaddlV2D = form("UaddlV2D")

func (s *Section) UaddlV2D(rd, rn, rm reg.V) {
	s.inst(uaddlV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uaddl2V2D = form("Uaddl2V2D")

func (s *Section) Uaddl2V2D(rd, rn, rm reg.V) {
	s.inst(uaddl2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var saddlV8H = form("SaddlV8H")

func (s *Section) SaddlV8H(rd, rn, rm reg.V) {
	s.inst(saddlV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var saddl2V8H = form("Saddl2V8H")

func (s *Section) Saddl2V8H(rd, rn, rm reg.V) {
	s.inst(saddl2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var saddlV4S = form("SaddlV4S")

func (s *Section) SaddlV4S(rd, rn, rm reg.V) {
	s.inst(saddlV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var saddl2V4S = form("Saddl2V4S")

func (s *Section) Saddl2V4S(rd, rn, rm reg.V) {
	s.inst(saddl2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var saddlV2D = form("SaddlV2D")

func (s *Section) SaddlV2D(rd, rn, rm reg.V) {
	s.inst(saddlV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var saddl2V2D = form("Saddl2V2D")

func (s *Section) Saddl2V2D(rd, rn, rm reg.V) {
	s.inst(saddl2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var usublV8H = form("UsublV8H")

func (s *Section) UsublV8H(rd, rn, rm reg.V) {
	s.inst(usublV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var usubl2V8H = form("Usubl2V8H")

func (s *Section) Usubl2V8H(rd, rn, rm reg.V) {
	s.inst(usubl2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var usublV4S = form("UsublV4S")

func (s *Section) UsublV4S(rd, rn, rm reg.V) {
	s.inst(usublV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var usubl2V4S = form("Usubl2V4S")

func (s *Section) Usubl2V4S(rd, rn, rm reg.V) {
	s.inst(usubl2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var usublV2D = form("UsublV2D")

func (s *Section) UsublV2D(rd, rn, rm reg.V) {
	s.inst(usublV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var usubl2V2D = form("Usubl2V2D")

func (s *Section) Usubl2V2D(rd, rn, rm reg.V) {
	s.inst(usubl2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var ssublV8H = form("SsublV8H")

func (s *Section) SsublV8H(rd, rn, rm reg.V) {
	s.inst(ssublV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var ssubl2V8H = form("Ssubl2V8H")

func (s *Section) Ssubl2V8H(rd, rn, rm reg.V) {
	s.inst(ssubl2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var ssublV4S = form("SsublV4S")

func (s *Section) SsublV4S(rd, rn, rm reg.V) {
	s.inst(ssublV4S, rd.Arr(reg.V4S), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var ssubl2V4S = form("Ssubl2V4S")

func (s *Section) Ssubl2V4S(rd, rn, rm reg.V) {
	s.inst(ssubl2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var ssublV2D = form("SsublV2D")

func (s *Section) SsublV2D(rd, rn, rm reg.V) {
	s.inst(ssublV2D, rd.Arr(reg.V2D), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var ssubl2V2D = form("Ssubl2V2D")

func (s *Section) Ssubl2V2D(rd, rn, rm reg.V) {
	s.inst(ssubl2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var pmullV8H = form("PmullV8H")

func (s *Section) PmullV8H(rd, rn, rm reg.V) {
	s.inst(pmullV8H, rd.Arr(reg.V8H), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var pmull2V8H = form("Pmull2V8H")

func (s *Section) Pmull2V8H(rd, rn, rm reg.V) {
	s.inst(pmull2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var pmullV1Q = form("PmullV1Q")

func (s *Section) PmullV1Q(rd, rn, rm reg.V) {
	s.inst(pmullV1Q, rd.Arr(reg.V1Q), rn.Arr(reg.V1D), rm.Arr(reg.V1D))
}

var pmull2V1Q = form("Pmull2V1Q")

func (s *Section) Pmull2V1Q(rd, rn, rm reg.V) {
	s.inst(pmull2V1Q, rd.Arr(reg.V1Q), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

// ---- Permutes ----

var uzp1V8B = form("Uzp1V8B")

func (s *Section) Uzp1V8B(rd, rn, rm reg.V) {
	s.inst(uzp1V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uzp1V16B = form("Uzp1V16B")

func (s *Section) Uzp1V16B(rd, rn, rm reg.V) {
	s.inst(uzp1V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uzp1V4H = form("Uzp1V4H")

func (s *Section) Uzp1V4H(rd, rn, rm reg.V) {
	s.inst(uzp1V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uzp1V8H = form("Uzp1V8H")

func (s *Section) Uzp1V8H(rd, rn, rm reg.V) {
	s.inst(uzp1V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uzp1V2S = form("Uzp1V2S")

func (s *Section) Uzp1V2S(rd, rn, rm reg.V) {
	s.inst(uzp1V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uzp1V4S = form("Uzp1V4S")

func (s *Section) Uzp1V4S(rd, rn, rm reg.V) {
	s.inst(uzp1V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uzp1V2D = form("Uzp1V2D")

func (s *Section) Uzp1V2D(rd, rn, rm reg.V) {
	s.inst(uzp1V2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var trn1V8B = form("Trn1V8B")

func (s *Section) Trn1V8B(rd, rn, rm reg.V) {
	s.inst(trn1V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var trn1V16B = form("Trn1V16B")

func (s *Section) Trn1V16B(rd, rn, rm reg.V) {
	s.inst(trn1V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var trn1V4H = form("Trn1V4H")

func (s *Section) Trn1V4H(rd, rn, rm reg.V) {
	s.inst(trn1V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var trn1V8H = form("Trn1V8H")

func (s *Section) Trn1V8H(rd, rn, rm reg.V) {
	s.inst(trn1V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var trn1V2S = form("Trn1V2S")

func (s *Section) Trn1V2S(rd, rn, rm reg.V) {
	s.inst(trn1V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var trn1V4S = form("Trn1V4S")

func (s *Section) Trn1V4S(rd, rn, rm reg.V) {
	s.inst(trn1V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var trn1V2D = form("Trn1V2D")

func (s *Section) Trn1V2D(rd, rn, rm reg.V) {
	s.inst(trn1V2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var zip1V8B = form("Zip1V8B")

func (s *Section) Zip1V8B(rd, rn, rm reg.V) {
	s.inst(zip1V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var zip1V16B = form("Zip1V16B")

func (s *Section) Zip1V16B(rd, rn, rm reg.V) {
	s.inst(zip1V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var zip1V4H = form("Zip1V4H")

func (s *Section) Zip1V4H(rd, rn, rm reg.V) {
	s.inst(zip1V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var zip1V8H = form("Zip1V8H")

func (s *Section) Zip1V8H(rd, rn, rm reg.V) {
	s.inst(zip1V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var zip1V2S = form("Zip1V2S")

func (s *Section) Zip1V2S(rd, rn, rm reg.V) {
	s.inst(zip1V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var zip1V4S = form("Zip1V4S")

func (s *Section) Zip1V4S(rd, rn, rm reg.V) {
	s.inst(zip1V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var zip1V2D = form("Zip1V2D")

func (s *Section) Zip1V2D(rd, rn, rm reg.V) {
	s.inst(zip1V2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var uzp2V8B = form("Uzp2V8B")

func (s *Section) Uzp2V8B(rd, rn, rm reg.V) {
	s.inst(uzp2V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var uzp2V16B = form("Uzp2V16B")

func (s *Section) Uzp2V16B(rd, rn, rm reg.V) {
	s.inst(uzp2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var uzp2V4H = form("Uzp2V4H")

func (s *Section) Uzp2V4H(rd, rn, rm reg.V) {
	s.inst(uzp2V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var uzp2V8H = form("Uzp2V8H")

func (s *Section) Uzp2V8H(rd, rn, rm reg.V) {
	s.inst(uzp2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var uzp2V2S = form("Uzp2V2S")

func (s *Section) Uzp2V2S(rd, rn, rm reg.V) {
	s.inst(uzp2V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var uzp2V4S = form("Uzp2V4S")

func (s *Section) Uzp2V4S(rd, rn, rm reg.V) {
	s.inst(uzp2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var uzp2V2D = form("Uzp2V2D")

func (s *Section) Uzp2V2D(rd, rn, rm reg.V) {
	s.inst(uzp2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var trn2V8B = form("Trn2V8B")

func (s *Section) Trn2V8B(rd, rn, rm reg.V) {
	s.inst(trn2V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var trn2V16B = form("Trn2V16B")

func (s *Section) Trn2V16B(rd, rn, rm reg.V) {
	s.inst(trn2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var trn2V4H = form("Trn2V4H")

func (s *Section) Trn2V4H(rd, rn, rm reg.V) {
	s.inst(trn2V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var trn2V8H = form("Trn2V8H")

func (s *Section) Trn2V8H(rd, rn, rm reg.V) {
	s.inst(trn2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var trn2V2S = form("Trn2V2S")

func (s *Section) Trn2V2S(rd, rn, rm reg.V) {
	s.inst(trn2V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var trn2V4S = form("Trn2V4S")

func (s *Section) Trn2V4S(rd, rn, rm reg.V) {
	s.inst(trn2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var trn2V2D = form("Trn2V2D")

func (s *Section) Trn2V2D(rd, rn, rm reg.V) {
	s.inst(trn2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var zip2V8B = form("Zip2V8B")

func (s *Section) Zip2V8B(rd, rn, rm reg.V) {
	s.inst(zip2V8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B))
}

var zip2V16B = form("Zip2V16B")

func (s *Section) Zip2V16B(rd, rn, rm reg.V) {
	s.inst(zip2V16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B))
}

var zip2V4H = form("Zip2V4H")

func (s *Section) Zip2V4H(rd, rn, rm reg.V) {
	s.inst(zip2V4H, rd.Arr(reg.V4H), rn.Arr(reg.V4H), rm.Arr(reg.V4H))
}

var zip2V8H = form("Zip2V8H")

func (s *Section) Zip2V8H(rd, rn, rm reg.V) {
	s.inst(zip2V8H, rd.Arr(reg.V8H), rn.Arr(reg.V8H), rm.Arr(reg.V8H))
}

var zip2V2S = form("Zip2V2S")

func (s *Section) Zip2V2S(rd, rn, rm reg.V) {
	s.inst(zip2V2S, rd.Arr(reg.V2S), rn.Arr(reg.V2S), rm.Arr(reg.V2S))
}

var zip2V4S = form("Zip2V4S")

func (s *Section) Zip2V4S(rd, rn, rm reg.V) {
	s.inst(zip2V4S, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var zip2V2D = form("Zip2V2D")

func (s *Section) Zip2V2D(rd, rn, rm reg.V) {
	s.inst(zip2V2D, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var extV8B = form("ExtV8B")

func (s *Section) ExtV8B(rd, rn, rm reg.V, index uint8) {
	s.inst(extV8B, rd.Arr(reg.V8B), rn.Arr(reg.V8B), rm.Arr(reg.V8B), uint64(index))
}

var extV16B = form("ExtV16B")

func (s *Section) ExtV16B(rd, rn, rm reg.V, index uint8) {
	s.inst(extV16B, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B), uint64(index))
}

var tblV8B = form("TblV8B")

func (s *Section) TblV8B(rd, table, index reg.V) {
	s.inst(tblV8B, rd.Arr(reg.V8B), table.Arr(reg.V16B), index.Arr(reg.V8B))
}

var tblV16B = form("TblV16B")

func (s *Section) TblV16B(rd, table, index reg.V) {
	s.inst(tblV16B, rd.Arr(reg.V16B), table.Arr(reg.V16B), index.Arr(reg.V16B))
}

var tbxV8B = form("TbxV8B")

func (s *Section) TbxV8B(rd, table, index reg.V) {
	s.inst(tbxV8B, rd.Arr(reg.V8B), table.Arr(reg.V16B), index.Arr(reg.V8B))
}

var tbxV16B = form("TbxV16B")

func (s *Section) TbxV16B(rd, table, index reg.V) {
	s.inst(tbxV16B, rd.Arr(reg.V16B), table.Arr(reg.V16B), index.Arr(reg.V16B))
}

// ---- Lanes and general registers ----

var dupV8BGen = form("DupV8BGen")

func (s *Section) DupV8BGen(rd reg.V, rn reg.W) { s.inst(dupV8BGen, rd.Arr(reg.V8B), rn) }

var dupV8BLane = form("DupV8BLane")

func (s *Section) DupV8BLane(rd, rn reg.V, index uint8) {
	s.inst(dupV8BLane, rd.Arr(reg.V8B), rn.Lane(reg.ElemB, index))
}

var dupV16BGen = form("DupV16BGen")

func (s *Section) DupV16BGen(rd reg.V, rn reg.W) { s.inst(dupV16BGen, rd.Arr(reg.V16B), rn) }

var dupV16BLane = form("DupV16BLane")

func (s *Section) DupV16BLane(rd, rn reg.V, index uint8) {
	s.inst(dupV16BLane, rd.Arr(reg.V16B), rn.Lane(reg.ElemB, index))
}

var umovB = form("UmovB")

func (s *Section) UmovB(rd reg.W, rn reg.V, index uint8) {
	s.inst(umovB, rd, rn.Lane(reg.ElemB, index))
}

var insBGen = form("InsBGen")

func (s *Section) InsBGen(rd reg.V, index uint8, rn reg.W) {
	s.inst(insBGen, rd.Lane(reg.ElemB, index), rn)
}

var dupV4HGen = form("DupV4HGen")

func (s *Section) DupV4HGen(rd reg.V, rn reg.W) { s.inst(dupV4HGen, rd.Arr(reg.V4H), rn) }

var dupV4HLane = form("DupV4HLane")

func (s *Section) DupV4HLane(rd, rn reg.V, index uint8) {
	s.inst(dupV4HLane, rd.Arr(reg.V4H), rn.Lane(reg.ElemH, index))
}

var dupV8HGen = form("DupV8HGen")

func (s *Section) DupV8HGen(rd reg.V, rn reg.W) { s.inst(dupV8HGen, rd.Arr(reg.V8H), rn) }

var dupV8HLane = form("DupV8HLane")

func (s *Section) DupV8HLane(rd, rn reg.V, index uint8) {
	s.inst(dupV8HLane, rd.Arr(reg.V8H), rn.Lane(reg.ElemH, index))
}

var umovH = form("UmovH")

func (s *Section) UmovH(rd reg.W, rn reg.V, index uint8) {
	s.inst(umovH, rd, rn.Lane(reg.ElemH, index))
}

var insHGen = form("InsHGen")

func (s *Section) InsHGen(rd reg.V, index uint8, rn reg.W) {
	s.inst(insHGen, rd.Lane(reg.ElemH, index), rn)
}

var dupV2SGen = form("DupV2SGen")

func (s *Section) DupV2SGen(rd reg.V, rn reg.W) { s.inst(dupV2SGen, rd.Arr(reg.V2S), rn) }

var dupV2SLane = form("DupV2SLane")

func (s *Section) DupV2SLane(rd, rn reg.V, index uint8) {
	s.inst(dupV2SLane, rd.Arr(reg.V2S), rn.Lane(reg.ElemS, index))
}

var dupV4SGen = form("DupV4SGen")

func (s *Section) DupV4SGen(rd reg.V, rn reg.W) { s.inst(dupV4SGen, rd.Arr(reg.V4S), rn) }

var dupV4SLane = form("DupV4SLane")

func (s *Section) DupV4SLane(rd, rn reg.V, index uint8) {
	s.inst(dupV4SLane, rd.Arr(reg.V4S), rn.Lane(reg.ElemS, index))
}

var umovS = form("UmovS")

func (s *Section) UmovS(rd reg.W, rn reg.V, index uint8) {
	s.inst(umovS, rd, rn.Lane(reg.ElemS, index))
}

var insSGen = form("InsSGen")

func (s *Section) InsSGen(rd reg.V, index uint8, rn reg.W) {
	s.inst(insSGen, rd.Lane(reg.ElemS, index), rn)
}

var dupV2DGen = form("DupV2DGen")

func (s *Section) DupV2DGen(rd reg.V, rn reg.X) { s.inst(dupV2DGen, rd.Arr(reg.V2D), rn) }

var dupV2DLane = form("DupV2DLane")

func (s *Section) DupV2DLane(rd, rn reg.V, index uint8) {
	s.inst(dupV2DLane, rd.Arr(reg.V2D), rn.Lane(reg.ElemD, index))
}

var umovD = form("UmovD")

func (s *Section) UmovD(rd reg.X, rn reg.V, index uint8) {
	s.inst(umovD, rd, rn.Lane(reg.ElemD, index))
}

var insDGen = form("InsDGen")

func (s *Section) InsDGen(rd reg.V, index uint8, rn reg.X) {
	s.inst(insDGen, rd.Lane(reg.ElemD, index), rn)
}

var smovB32 = form("SmovB32")

func (s *Section) SmovB32(rd reg.W, rn reg.V, index uint8) {
	s.inst(smovB32, rd, rn.Lane(reg.ElemB, index))
}

var smovH32 = form("SmovH32")

func (s *Section) SmovH32(rd reg.W, rn reg.V, index uint8) {
	s.inst(smovH32, rd, rn.Lane(reg.ElemH, index))
}

var smovB64 = form("SmovB64")

func (s *Section) SmovB64(rd reg.X, rn reg.V, index uint8) {
	s.inst(smovB64, rd, rn.Lane(reg.ElemB, index))
}

var smovH64 = form("SmovH64")

func (s *Section) SmovH64(rd reg.X, rn reg.V, index uint8) {
	s.inst(smovH64, rd, rn.Lane(reg.ElemH, index))
}

var smovS64 = form("SmovS64")

func (s *Section) SmovS64(rd reg.X, rn reg.V, index uint8) {
	s.inst(smovS64, rd, rn.Lane(reg.ElemS, index))
}

// ---- Immediates into a vector ----

var moviV2D = form("MoviV2D")

func (s *Section) MoviV2D(rd reg.V, mask uint64) { s.inst(moviV2D, rd.Arr(reg.V2D), mask) }

var moviV8B = form("MoviV8B")

func (s *Section) MoviV8B(rd reg.V, imm uint8) { s.inst(moviV8B, rd.Arr(reg.V8B), uint64(imm)) }

var moviV16B = form("MoviV16B")

func (s *Section) MoviV16B(rd reg.V, imm uint8) { s.inst(moviV16B, rd.Arr(reg.V16B), uint64(imm)) }

// ---- AES ----

var aese = form("Aese")

func (s *Section) Aese(rd, rn reg.V) { s.inst(aese, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var aesd = form("Aesd")

func (s *Section) Aesd(rd, rn reg.V) { s.inst(aesd, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var aesmc = form("Aesmc")

func (s *Section) Aesmc(rd, rn reg.V) { s.inst(aesmc, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

var aesimc = form("Aesimc")

func (s *Section) Aesimc(rd, rn reg.V) { s.inst(aesimc, rd.Arr(reg.V16B), rn.Arr(reg.V16B)) }

// ---- SHA-1 and SHA-256 ----

var sha1c = form("Sha1c")

func (s *Section) Sha1c(rd reg.Q, rn reg.S, rm reg.V) { s.inst(sha1c, rd, rn, rm.Arr(reg.V4S)) }

var sha1p = form("Sha1p")

func (s *Section) Sha1p(rd reg.Q, rn reg.S, rm reg.V) { s.inst(sha1p, rd, rn, rm.Arr(reg.V4S)) }

var sha1m = form("Sha1m")

func (s *Section) Sha1m(rd reg.Q, rn reg.S, rm reg.V) { s.inst(sha1m, rd, rn, rm.Arr(reg.V4S)) }

var sha1su0 = form("Sha1su0")

func (s *Section) Sha1su0(rd, rn, rm reg.V) {
	s.inst(sha1su0, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sha256h = form("Sha256h")

func (s *Section) Sha256h(rd, rn reg.Q, rm reg.V) { s.inst(sha256h, rd, rn, rm.Arr(reg.V4S)) }

var sha256h2 = form("Sha256h2")

func (s *Section) Sha256h2(rd, rn reg.Q, rm reg.V) { s.inst(sha256h2, rd, rn, rm.Arr(reg.V4S)) }

var sha256su1 = form("Sha256su1")

func (s *Section) Sha256su1(rd, rn, rm reg.V) {
	s.inst(sha256su1, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sha1h = form("Sha1h")

func (s *Section) Sha1h(rd, rn reg.S) { s.inst(sha1h, rd, rn) }

var sha1su1 = form("Sha1su1")

func (s *Section) Sha1su1(rd, rn reg.V) { s.inst(sha1su1, rd.Arr(reg.V4S), rn.Arr(reg.V4S)) }

var sha256su0 = form("Sha256su0")

func (s *Section) Sha256su0(rd, rn reg.V) { s.inst(sha256su0, rd.Arr(reg.V4S), rn.Arr(reg.V4S)) }

// ---- SHA-512 and SHA-3 ----

var sha512h = form("Sha512h")

func (s *Section) Sha512h(rd, rn reg.Q, rm reg.V) { s.inst(sha512h, rd, rn, rm.Arr(reg.V2D)) }

var sha512h2 = form("Sha512h2")

func (s *Section) Sha512h2(rd, rn reg.Q, rm reg.V) { s.inst(sha512h2, rd, rn, rm.Arr(reg.V2D)) }

var sha512su1 = form("Sha512su1")

func (s *Section) Sha512su1(rd, rn, rm reg.V) {
	s.inst(sha512su1, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var sha512su0 = form("Sha512su0")

func (s *Section) Sha512su0(rd, rn reg.V) { s.inst(sha512su0, rd.Arr(reg.V2D), rn.Arr(reg.V2D)) }

var eor3 = form("Eor3")

func (s *Section) Eor3(rd, rn, rm, ra reg.V) {
	s.inst(eor3, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B), ra.Arr(reg.V16B))
}

var bcax = form("Bcax")

func (s *Section) Bcax(rd, rn, rm, ra reg.V) {
	s.inst(bcax, rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B), ra.Arr(reg.V16B))
}

var rax1 = form("Rax1")

func (s *Section) Rax1(rd, rn, rm reg.V) {
	s.inst(rax1, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D))
}

var xar = form("Xar")

func (s *Section) Xar(rd, rn, rm reg.V, rotate uint8) {
	s.inst(xar, rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D), uint64(rotate))
}

// ---- SM3 and SM4 ----

var sm3ss1 = form("Sm3ss1")

func (s *Section) Sm3ss1(rd, rn, rm, ra reg.V) {
	s.inst(sm3ss1, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S), ra.Arr(reg.V4S))
}

var sm3tt1a = form("Sm3tt1a")

func (s *Section) Sm3tt1a(rd, rn, rm reg.V, index uint8) {
	s.inst(sm3tt1a, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Lane(reg.ElemS, index))
}

var sm3tt1b = form("Sm3tt1b")

func (s *Section) Sm3tt1b(rd, rn, rm reg.V, index uint8) {
	s.inst(sm3tt1b, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Lane(reg.ElemS, index))
}

var sm3tt2a = form("Sm3tt2a")

func (s *Section) Sm3tt2a(rd, rn, rm reg.V, index uint8) {
	s.inst(sm3tt2a, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Lane(reg.ElemS, index))
}

var sm3tt2b = form("Sm3tt2b")

func (s *Section) Sm3tt2b(rd, rn, rm reg.V, index uint8) {
	s.inst(sm3tt2b, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Lane(reg.ElemS, index))
}

var sm3partw1 = form("Sm3partw1")

func (s *Section) Sm3partw1(rd, rn, rm reg.V) {
	s.inst(sm3partw1, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sm3partw2 = form("Sm3partw2")

func (s *Section) Sm3partw2(rd, rn, rm reg.V) {
	s.inst(sm3partw2, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sm4ekey = form("Sm4ekey")

func (s *Section) Sm4ekey(rd, rn, rm reg.V) {
	s.inst(sm4ekey, rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S))
}

var sm4e = form("Sm4e")

func (s *Section) Sm4e(rd, rn reg.V) { s.inst(sm4e, rd.Arr(reg.V4S), rn.Arr(reg.V4S)) }

// ---- CRC-32 and CRC-32C ----

var crc32b = form("Crc32b")

func (s *Section) Crc32b(rd, rn reg.W, rm reg.W) { s.inst(crc32b, rd, rn, rm) }

var crc32h = form("Crc32h")

func (s *Section) Crc32h(rd, rn reg.W, rm reg.W) { s.inst(crc32h, rd, rn, rm) }

var crc32w = form("Crc32w")

func (s *Section) Crc32w(rd, rn reg.W, rm reg.W) { s.inst(crc32w, rd, rn, rm) }

var crc32x = form("Crc32x")

func (s *Section) Crc32x(rd, rn reg.W, rm reg.X) { s.inst(crc32x, rd, rn, rm) }

var crc32cb = form("Crc32cb")

func (s *Section) Crc32cb(rd, rn reg.W, rm reg.W) { s.inst(crc32cb, rd, rn, rm) }

var crc32ch = form("Crc32ch")

func (s *Section) Crc32ch(rd, rn reg.W, rm reg.W) { s.inst(crc32ch, rd, rn, rm) }

var crc32cw = form("Crc32cw")

func (s *Section) Crc32cw(rd, rn reg.W, rm reg.W) { s.inst(crc32cw, rd, rn, rm) }

var crc32cx = form("Crc32cx")

func (s *Section) Crc32cx(rd, rn reg.W, rm reg.X) { s.inst(crc32cx, rd, rn, rm) }
