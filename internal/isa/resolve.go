package isa

import (
	"fmt"
	"strings"

	"github.com/vertex-language/arm64/feature"
)

// UnknownError is a mnemonic no form declares.
type UnknownError struct{ Mnem string }

func (e *UnknownError) Error() string {
	return fmt.Sprintf("unknown instruction %q", e.Mnem)
}

// FormError is a mnemonic that exists with no form accepting these operands.
type FormError struct {
	Mnem string
	Args []Arg
	// Near lists the forms that were closest, for the diagnostic.
	Near []*Form
}

func (e *FormError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "no form of %s takes ", e.Mnem)
	if len(e.Args) == 0 {
		b.WriteString("no operands")
	} else {
		for i, a := range e.Args {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(a.Class.String())
		}
	}
	if len(e.Near) > 0 {
		b.WriteString("\n  candidates:")
		for _, f := range e.Near {
			b.WriteString("\n    ")
			b.WriteString(f.Signature())
		}
	}
	return b.String()
}

// GateError is a form that matches but that the active feature set holds back.
type GateError struct {
	Form   *Form
	Active feature.Set
}

func (e *GateError) Error() string {
	return fmt.Sprintf("%s requires %s, not in the active feature set\n  active: %s\n  note: aarch64.WithFeatures(%s)",
		e.Form.Mnem, e.Form.Gate, e.Active,
		e.Active.Plus(e.Form.Gate).GoExpr())
}

// Resolve finds the one form of a mnemonic that accepts these operands.
//
// There is no shortest-form search and no preference order. Every A64
// instruction is one word, so two forms of a mnemonic accepting the same
// operand classes would be an ambiguity with no tiebreak that means anything —
// and the table refuses to build if two such forms exist, so this function
// never has to choose.
//
// A form that matches but is gated returns a GateError rather than falling
// through to a FormError. Being told an instruction does not exist when it
// exists and is disabled sends a reader looking for a typo that is not there.
func Resolve(mnem string, args []Arg, set feature.Set) (*Form, error) {
	mnem = strings.ToLower(mnem)
	forms := byMnem[mnem]
	if len(forms) == 0 {
		return nil, &UnknownError{Mnem: mnem}
	}

	var gated []*Form
	var near []*Form

	for _, f := range forms {
		if !accepts(f, args) {
			near = append(near, f)
			continue
		}
		if !f.Enabled(set) {
			gated = append(gated, f)
			continue
		}
		return f, nil
	}

	if len(gated) > 0 {
		return nil, &GateError{Form: gated[0], Active: set}
	}
	if len(near) == 0 {
		near = forms
	}
	if len(near) > 4 {
		near = near[:4]
	}
	return nil, &FormError{Mnem: mnem, Args: args, Near: near}
}

// accepts reports whether a form takes these arguments.
//
// Slots and arguments are not in step, which is the whole reason this is not a
// zip. A memory operand is one argument filling two slots — a base and the
// displacement beside it — because an address is one thing to whoever wrote it
// and two fields to the encoder. encodeForm walks the two indices separately
// for the same reason; this function has to walk them the same way, or a form
// would resolve that then failed to encode.
func accepts(f *Form, args []Arg) bool {
	si := 0
	for _, a := range args {
		if si >= len(f.Slots) {
			return false
		}
		s := f.Slots[si]
		if !s.Class.Match(a) {
			return false
		}
		si++
		if s.Class.Mem() {
			if !addrFits(f, s, a) {
				return false
			}
			if si < len(f.Slots) && f.Slots[si].Role == RoleOffset {
				si++
			}
		}
	}
	// Anything left has to be a slot the caller was allowed to omit.
	for ; si < len(f.Slots); si++ {
		if !f.Slots[si].Optional {
			return false
		}
	}
	return true
}

// addrFits reports whether a form's encoding expresses the addressing mode an
// argument asks for.
//
// The writeback modes are separate encodings, so they must agree in both
// directions: a pre-indexed address needs a form that writes back, and an
// ordinary one must not resolve to a form that does. Only the first direction
// is checked at encode time, because the typed surface names the form and its
// caller cannot get them crossed. Text can — `stp x0, x1, [sp, #16]` and
// `stp x0, x1, [sp, #16]!` differ by one character — so both directions are
// checked here.
func addrFits(f *Form, base Slot, a Arg) bool {
	writeback := f.Attrs&(AttrPreIndex|AttrPostIndex) != 0
	switch a.Addr {
	case AddrPreIndex:
		return f.Attrs&AttrPreIndex != 0
	case AddrPostIndex:
		return f.Attrs&AttrPostIndex != 0
	case AddrRegOffset:
		return f.Attrs&AttrRegOffset != 0
	case AddrBase, AddrOffset, AddrNone:
		// A register-offset row is a different encoding rather than a
		// different operand, so it does not also accept [Xn] — which is
		// what keeps `ldr x0, [x1]` on the scaled-immediate row.
		return !writeback && f.Attrs&AttrRegOffset == 0
	}
	return false
}

// ResolveWord finds the form a word decodes to, by linear scan over the table,
// skipping aliases. It is the reference answer the table's own checks compare
// against; anything that needs to decode at speed builds its own lookup from
// All().
func ResolveWord(word uint32) (*Form, bool) {
	for _, f := range all {
		if f.Attrs&AttrAlias != 0 {
			continue
		}
		if word&f.Mask == f.Word {
			return f, true
		}
	}
	return nil, false
}
