package operand

import (
	"strconv"

	"github.com/vertex-language/arm64/obj"
)

// RefKind and the reference kinds are obj's, aliased here so a lowering that
// builds operands and never touches obj directly still spells the same
// constant: arm64.RefAdrPage21, operand.RefAdrPage21 and obj.RefAdrPage21 are
// one value, and no conversion exists anywhere in the tree.
type RefKind = obj.RefKind

const RefNone = obj.RefNone

// Target is what an address operand points at before it is a number.
type Target interface {
	isTarget()
	String() string
}

// Label is a name defined somewhere in this object.
//
// A Label carries no relocation kind, which is what makes it foldable: a
// pc-relative reference to a local label in the same section resolves at
// Serialize and leaves no record behind. Ref is the spelling for everything
// else.
type Label string

func (Label) isTarget()        {}
func (l Label) String() string { return string(l) }

// SymRef is a reference to a symbol, with an addend and optionally the
// relocation kind the caller insists on.
//
// Stating a kind is a request, not a hint. Writing Ref("puts", R_AARCH64_CALL26)
// asks for a branch relocation, and folding it into a direct branch — even to a
// symbol two lines above — would answer a different question than the one
// asked.
type SymRef struct {
	Name   string
	Addend int64
	Kind   RefKind
}

func (SymRef) isTarget() {}

func (s SymRef) String() string {
	out := s.Name
	if s.Addend > 0 {
		out += "+" + strconv.FormatInt(s.Addend, 10)
	} else if s.Addend < 0 {
		out += strconv.FormatInt(s.Addend, 10)
	}
	return out
}

// Sym builds a symbol reference. The arch package re-exports it as Ref, which
// is the spelling the README and the builder examples use.
func Sym(name string, kind ...RefKind) SymRef {
	s := SymRef{Name: name}
	if len(kind) > 0 {
		s.Kind = kind[0]
	}
	return s
}

// Plus is a reference with an addend: Sym("puts").Plus(8) is bl puts+8.
func (s SymRef) Plus(addend int64) SymRef { s.Addend += addend; return s }

// AddrRole is which part of an address a reference names.
//
// Materializing an address on this architecture usually takes two instructions
// and therefore two references — adrp for the page, add or a load for the
// offset within it — and each needs its own record. The role is the portable
// part: GNU as spells the pair :pg_hi21: and :lo12:, the Darwin assembler
// spells it @PAGE and @PAGEOFF, and the kind is per format. The caller states
// the role and never the kind.
type AddrRole uint8

const (
	// RoleDirect is the address itself: a branch target, or a literal load.
	RoleDirect AddrRole = iota

	RolePage       // :pg_hi21: / @PAGE
	RolePageOff    // :lo12:    / @PAGEOFF
	RoleGotPage    // :got:     / @GOTPAGE
	RoleGotPageOff // :got_lo12:/ @GOTPAGEOFF

	// Mach-O's thread-local pair. The same two instructions as the GOT
	// pair against a thread-local's descriptor instead of a GOT entry,
	// and spelled @TLVPPAGE / @TLVPPAGEOFF in Apple's assembler.
	RoleTlvPage    // @TLVPPAGE
	RoleTlvPageOff // @TLVPPAGEOFF
)

func (r AddrRole) String() string {
	switch r {
	case RolePage:
		return "page"
	case RolePageOff:
		return "pageoff"
	case RoleGotPage:
		return "gotpage"
	case RoleTlvPage:
		return "tlvpage"
	case RoleTlvPageOff:
		return "tlvpageoff"
	case RoleGotPageOff:
		return "gotpageoff"
	}
	return "direct"
}

// GOT reports whether the role goes through the global offset table, which is
// what makes it a reference to a slot holding the address rather than to the
// address.
func (r AddrRole) GOT() bool { return r == RoleGotPage || r == RoleGotPageOff }

// AddrRef is a target together with the role naming which half of its address
// this operand wants.
type AddrRef struct {
	T    Target
	Role AddrRole
}

// Page, PageOff, GotPage and GotPageOff are the four roles.
func Page(t Target) AddrRef       { return AddrRef{T: t, Role: RolePage} }
func PageOff(t Target) AddrRef    { return AddrRef{T: t, Role: RolePageOff} }
func GotPage(t Target) AddrRef    { return AddrRef{T: t, Role: RoleGotPage} }
func GotPageOff(t Target) AddrRef { return AddrRef{T: t, Role: RoleGotPageOff} }
func TlvPage(t Target) AddrRef    { return AddrRef{T: t, Role: RoleTlvPage} }
func TlvPageOff(t Target) AddrRef { return AddrRef{T: t, Role: RoleTlvPageOff} }

// Direct wraps a bare target, so a caller lowering operands has one type to
// handle rather than two.
func Direct(t Target) AddrRef { return AddrRef{T: t, Role: RoleDirect} }

// Kind is the relocation kind the caller named on the underlying symbol, or
// RefNone.
func (a AddrRef) Kind() RefKind {
	if s, ok := a.T.(SymRef); ok {
		return s.Kind
	}
	return RefNone
}

// Addend is the offset from the symbol the reference means.
func (a AddrRef) Addend() int64 {
	if s, ok := a.T.(SymRef); ok {
		return s.Addend
	}
	return 0
}

func (a AddrRef) String() string {
	if a.T == nil {
		return "<nil>"
	}
	switch a.Role {
	case RolePage:
		return ":pg_hi21:" + a.T.String()
	case RolePageOff:
		return ":lo12:" + a.T.String()
	case RoleGotPage:
		return ":got:" + a.T.String()
	case RoleTlvPage:
		return "tlvpage"
	case RoleTlvPageOff:
		return "tlvpageoff"
	case RoleGotPageOff:
		return ":got_lo12:" + a.T.String()
	}
	return a.T.String()
}
