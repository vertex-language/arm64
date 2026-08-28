package arm64

import (
	"github.com/vertex-language/arm64/internal/encode"
	"github.com/vertex-language/arm64/obj"
	"github.com/vertex-language/arm64/operand"
)

// refKindFor decides which obj.RefKind a fixup that survives to Finalize
// becomes.
//
// A caller who named a kind on the symbol — Ref("x", obj.RefTlsIeAdrGottprelPage21)
// — has already answered this question, and that answer wins outright: it is
// the only path to the TLS general-dynamic, initial-exec and local-exec kinds,
// none of which operand.AddrRole has a role for. Absent that, the kind
// follows from what the field is for: mnem picks the branch family a field of
// AttrBranch belongs to, since CALL26, JUMP26, CONDBR19 and TSTBR14 are four
// different fields the table cannot tell apart by shape alone, and Role picks
// among ADR, ADRP and the page-offset and GOT families for everything else.
func refKindFor(mnem string, fx encode.Fixup) (obj.RefKind, error) {
	if fx.Kind.Valid() {
		return fx.Kind, nil
	}

	if fx.Branch {
		switch mnem {
		case "bl":
			return obj.RefCall26, nil
		case "b":
			return obj.RefJump26, nil
		case "b.cond", "cbz", "cbnz":
			return obj.RefCondBr19, nil
		case "tbz", "tbnz":
			return obj.RefTstBr14, nil
		}
		// ldr (literal) reaches here: its nineteen-bit field has no relocation
		// in any of the three container formats, because a literal pool is a
		// same-section construct in every ABI that defines one. A same-section
		// Label folds before this function is ever called; naming an external
		// symbol here is the caller asking for something no linker can do.
		return 0, &obj.Error{
			Sentinel: obj.ErrRefKind,
			Context:  mnem + " has no relocation for a symbol outside this section",
			Notes:    []string{"a literal load reaches only a label defined in the same section"},
		}
	}

	switch fx.Role {
	case operand.RoleDirect:
		return obj.RefAdrPrel21, nil
	case operand.RolePage:
		return obj.RefAdrPage21, nil
	case operand.RoleGotPage:
		return obj.RefAdrGotPage21, nil
	case operand.RoleGotPageOff:
		return obj.RefLd64GotLo12, nil
	case operand.RolePageOff:
		switch fx.Access {
		case operand.Width8:
			return obj.RefLdSt8AbsLo12, nil
		case operand.Width16:
			return obj.RefLdSt16AbsLo12, nil
		case operand.Width32:
			return obj.RefLdSt32AbsLo12, nil
		case operand.Width64:
			return obj.RefLdSt64AbsLo12, nil
		case operand.Width128:
			return obj.RefLdSt128AbsLo12, nil
		}
		// No memory access at this slot: the page offset is going into an ADD
		// immediate rather than a load or store's addressing mode.
		return obj.RefAddAbsLo12, nil
	}

	return 0, &obj.Error{
		Sentinel: obj.ErrRefKind,
		Context:  "reference has no relocation for its role",
	}
}
