package operand

import "strings"

// PState is the field MSR (immediate) writes: not a system register, but a
// named piece of process state that takes a four-bit immediate.
//
// The encoding is the op1 and op2 pair of the MSR (immediate) form, packed
// here the way Sys packs its five, so the value is the fields' own and not an
// arbitrary enumeration.
//
// It is a type of its own rather than a Sys because the two are different
// operands of different instructions: `msr daifset, #2` sets bits in DAIF and
// `msr daif, x0` writes the whole of it, and only the second names a register
// the architecture will also let you read.
type PState uint8

// NewPState composes a field from its op1 and op2.
func NewPState(op1, op2 uint8) PState { return PState(op1&7<<3 | op2&7) }

func (p PState) Op1() uint8 { return uint8(p) >> 3 & 7 }
func (p PState) Op2() uint8 { return uint8(p) & 7 }

// The fields with a name. The list is short because the architecture's is:
// most process state is reached through a system register, and these are the
// pieces that are set and cleared often enough to have an instruction that
// does it in one.
const (
	SPSelField = PState(0<<3 | 5)
	DAIFSet    = PState(3<<3 | 6)
	DAIFClr    = PState(3<<3 | 7)
	PStateNone = PState(0xff)
)

var pstateName = map[PState]string{
	SPSelField: "spsel", DAIFSet: "daifset", DAIFClr: "daifclr",
}

func (p PState) String() string {
	if n, ok := pstateName[p]; ok {
		return n
	}
	return "pstate?"
}

// Valid reports whether p is one of the named fields. Unlike a system
// register there is no generic spelling to fall back on: the op1/op2 pair is
// not written out in assembly, so a field this package does not name is a
// field no source can reach.
func (p PState) Valid() bool { _, ok := pstateName[p]; return ok }

// LookupPState resolves a field name, case-insensitively.
func LookupPState(name string) (PState, bool) {
	s := strings.ToLower(strings.TrimSpace(name))
	for p, n := range pstateName {
		if n == s {
			return p, true
		}
	}
	return PStateNone, false
}
