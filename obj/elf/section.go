package elf

import (
	"strings"

	elfcore "github.com/vertex-language/elf"
	elfobj "github.com/vertex-language/elf/obj"

	"github.com/vertex-language/arm64/obj"
)

// Section kinds map onto a type and a flag word. The kind is a load-time
// property, which is exactly what SHF_ALLOC, SHF_WRITE and SHF_EXECINSTR
// describe, so this is a translation and not a decision.
//
// One rule is not derivable from the kind, and it is the one heuristic in
// this package: a section whose name marks it as debug or comment data is not
// allocated at run time, whatever kind it was created with.
func sectionShape(s *obj.Section) (elfcore.SHType, uint64) {
	if s.Kind() == obj.BSS {
		if nonAlloc(s.Name()) {
			return elfcore.SHT_PROGBITS, 0
		}
		return elfcore.SHT_NOBITS, elfcore.SHF_ALLOC | elfcore.SHF_WRITE
	}

	if nonAlloc(s.Name()) {
		return elfcore.SHT_PROGBITS, 0
	}

	switch s.Kind() {
	case obj.Text:
		return elfcore.SHT_PROGBITS, elfcore.SHF_ALLOC | elfcore.SHF_EXECINSTR
	case obj.Data:
		return elfcore.SHT_PROGBITS, elfcore.SHF_ALLOC | elfcore.SHF_WRITE
	case obj.ROData:
		return elfcore.SHT_PROGBITS, elfcore.SHF_ALLOC
	case obj.RelROData:
		// Allocated and writable: the loader relocates it, then
		// mprotects it read-only. Without SHF_WRITE the relocation
		// would fault.
		return elfcore.SHT_PROGBITS, elfcore.SHF_ALLOC | elfcore.SHF_WRITE
	}
	return elfcore.SHT_PROGBITS, elfcore.SHF_ALLOC
}

// nonAlloc names the sections ELF itself defines as not present at run time.
func nonAlloc(name string) bool {
	switch {
	case strings.HasPrefix(name, ".debug"):
		return true
	case name == ".comment":
		return true
	}
	return false
}

// writeSection creates the section and fills it.
//
// A BSS section arrives with a run of zero bytes in it, because the builder's
// Zero appends real bytes and Offset has to keep meaning what it says. Here
// that run becomes sh_size and no file content, which is the whole point of
// SHT_NOBITS.
func writeSection(wr *elfobj.Writer, s *obj.Section) (*elfobj.SectionBuilder, error) {
	typ, flags := sectionShape(s)

	hdr := elfobj.SectionHeader{
		Name:      s.Name(),
		Type:      typ,
		Flags:     flags,
		Addralign: uint64(s.Align()),
	}
	if typ == elfcore.SHT_NOBITS {
		hdr.Size = uint64(s.Size())
	}

	b := wr.Section(hdr)
	if typ == elfcore.SHT_NOBITS {
		return b, nil
	}
	// The bytes go through untouched. AArch64 is RELA, so the addend rides
	// in the entry and there is nothing to fold in here.
	if _, err := b.Write(s.Bytes()); err != nil {
		return nil, err
	}
	return b, nil
}
