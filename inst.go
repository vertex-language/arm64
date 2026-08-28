package arm64

// inst.go is the join between the ISA table and the typed helpers. The
// helpers themselves live one tranche per file — inst_arith.go opposite the
// arithmetic rows of internal/isa/table_base.go, inst_branch.go opposite the
// branch rows — and every one of them is three lines:
//
//	var addImm64 = form("AddImm64")
//
//	// AddImm64 emits ADD Xd|SP, Xn|SP, #imm{, shift}.
//	func (s *Section) AddImm64(rd, rn RegSP64, imm ImmOrRef, shift ...ShiftOp) {
//		s.inst(addImm64, rd, rn, imm, opt(shift)...)
//	}
//
// The binding is by name, not by table index. Appending rows breaks nothing;
// a removed or renamed row panics at program start naming the missing form,
// rather than silently binding to the wrong row or failing halfway through
// someone's code generation. Two helpers binding one name is a duplicate Go
// identifier and so a compile error, which is the earliest any of these can
// fail — and isa.checkTable checks the same thing from the table's side,
// because the table is what makes the promise.
import "github.com/vertex-language/arm64/internal/isa"

// form binds a helper to its row.
//
// It panics rather than returning an error because a missing row is a
// question about this tree's own data, not about anything a caller did. It
// should fail whoever the caller is and whatever they were about to do,
// which is what package initialization means.
func form(helper string) *isa.Form {
	f := isa.ByHelper(helper)
	if f == nil {
		panic("arm64: the ISA table declares no form for helper " + helper +
			"; the row was removed or renamed")
	}
	return f
}

// instForm is Inst's — the ".inst" row, one raw-word immediate and nothing
// else, in emit.go rather than beside the other forms because it is not an
// instruction so much as an escape from having declared one.
var instForm = form("Inst")

// opt turns a variadic modifier into the single trailing operand inst wants,
// or no operand at all when the caller left it at its default.
func opt[T any](v []T) []any {
	if len(v) == 0 {
		return nil
	}
	return []any{v[0]}
}
