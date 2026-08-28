package pe

import (
	"fmt"

	pecore "github.com/vertex-language/pe"
	"github.com/vertex-language/pe/coff"

	"github.com/vertex-language/arm64/obj"
)

// ELF and COFF disagree about what read-only initialized data is called:
// .rodata on one side, .rdata on the other, and link.exe's default merge
// rules are written against .rdata. The rename applies only to a section
// still carrying the conventional name for its kind.
func coffName(s *obj.Section) string {
	if s.Name() != s.Kind().String() {
		return s.Name()
	}
	if s.Kind() == obj.ROData {
		return ".rdata"
	}
	return s.Name()
}

func sectionShape(k obj.SectionKind) (pecore.SecKind, pecore.SecProt) {
	switch k {
	case obj.Text:
		return pecore.SecCode, pecore.SecExecute | pecore.SecRead
	case obj.Data:
		return pecore.SecInitData, pecore.SecRead | pecore.SecWrite
	case obj.ROData:
		return pecore.SecInitData, pecore.SecRead
	case obj.BSS:
		return pecore.SecUninitData, pecore.SecRead | pecore.SecWrite
	}
	return pecore.SecInitData, pecore.SecRead
}

func newSection(wr *coff.Writer, s *obj.Section) *coff.SectionBuilder {
	kind, prot := sectionShape(s.Kind())
	return wr.Section(coff.SectionHeader{
		Name:  coffName(s),
		Kind:  kind,
		Prot:  prot,
		Align: s.Align(),
	})
}

// writeContents deposits every addend and hands the section its bytes.
func writeContents(wr *coff.Writer, b *coff.SectionBuilder, s *obj.Section, syms map[string]*coff.SymbolRef) error {
	if s.Kind() == obj.BSS {
		if len(s.Refs()) > 0 {
			return fmt.Errorf("pe: %s is uninitialized data and has no bytes to hold %d relocation addends",
				s.Name(), len(s.Refs()))
		}
		b.Reserve(uint32(s.Size()))
		return wr.Err()
	}

	content := s.Bytes()
	if err := writeRelocs(wr, b, s, content, syms); err != nil {
		return err
	}
	if _, err := b.Write(content); err != nil {
		return err
	}
	return wr.Err()
}
