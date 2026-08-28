// Package obj is the vocabulary the rest of this tree spells its types in,
// and the finished artifact the builder hands to a writer.
//
// It knows no instruction set and imports nothing from this tree. That is
// what puts it at the bottom of the import graph beside reg: operand needs
// RefKind and Error, internal/encode needs both, the root re-exports
// everything, and the three writers under obj/ take an *Object. A package
// this far down cannot import anything above it, and nothing here wants to.
//
// Everything in here is a constant, a plain data struct, or an accessor over
// one. Nothing acts. Past Finalize the artifact is inert: data with no
// methods that do anything but read, safe to write more than once, in more
// than one format, from more than one goroutine.
package obj

// Arch names the architecture an object was built for.
//
// It is a field of the object rather than of a writer's options because
// everything a container's header needs follows from it. A caller who could
// pass a different one could produce a file whose header disagrees with its
// bytes, which is why no writer here has a Target field.
type Arch uint8

const (
	ArchNone Arch = iota
	ArchARM64
)

var archNames = [...]string{
	ArchNone:  "",
	ArchARM64: "arm64",
}

// String is what a diagnostic leads with: "arm64 .text+0x11: ...".
func (a Arch) String() string {
	if int(a) < len(archNames) {
		return archNames[a]
	}
	return "arch?"
}

// Valid reports whether the value names a declared architecture. ArchNone is
// not one and reports false.
func (a Arch) Valid() bool { return a == ArchARM64 }

// Bits is the architecture's pointer width, or 0 for ArchNone. AArch64 has no
// 32-bit mode this tree builds for: ILP32 is a data-model variant of the same
// instruction set, not a different Arch, and nothing here emits it.
func (a Arch) Bits() int {
	switch a {
	case ArchARM64:
		return 64
	}
	return 0
}
