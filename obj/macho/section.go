package macho

import (
	"fmt"
	"strings"

	machocore "github.com/vertex-language/macho"
	machoobj "github.com/vertex-language/macho/obj"

	"github.com/vertex-language/arm64/obj"
)

// The four kinds and their conventional pairs. ROData goes to (__TEXT,__const)
// rather than a __DATA section because read-only is a __TEXT property here.
func kindPlacement(k obj.SectionKind) (SegSect, machocore.SecType, machocore.SecAttrs) {
	switch k {
	case obj.Text:
		return SegSect{machocore.SEG_TEXT, machocore.SECT_TEXT}, machocore.S_REGULAR,
			machocore.S_ATTR_PURE_INSTRUCTIONS | machocore.S_ATTR_SOME_INSTRUCTIONS
	case obj.Data:
		return SegSect{machocore.SEG_DATA, machocore.SECT_DATA}, machocore.S_REGULAR, 0
	case obj.ROData:
		return SegSect{machocore.SEG_TEXT, machocore.SECT_CONST}, machocore.S_REGULAR, 0
	case obj.BSS:
		return SegSect{machocore.SEG_DATA, machocore.SECT_BSS}, machocore.S_ZEROFILL, 0
	}
	return SegSect{machocore.SEG_DATA, machocore.SECT_DATA}, machocore.S_REGULAR, 0
}

// tlsPlacement is where the thread-local sections go, and what type
// each carries.
//
// The type is the point. __thread_data holds the templates every
// thread's copy is made from, __thread_bss the ones that start zeroed,
// and __thread_vars the three-word descriptors dyld fills in — and a
// linker and a loader tell them apart by S_THREAD_LOCAL_*, not by name.
// Written as S_REGULAR they would be ordinary data, copied once and
// shared by every thread.
//
// The names are spelled the ELF way, as every other section here is:
// .tdata and .tbss are ELF's own, and .tlvdesc has no ELF counterpart
// because ELF has no descriptors.
var tlsPlacement = map[string]struct {
	ss  SegSect
	typ machocore.SecType
}{
	".tdata":   {SegSect{machocore.SEG_DATA, "__thread_data"}, machocore.S_THREAD_LOCAL_REGULAR},
	".tbss":    {SegSect{machocore.SEG_DATA, "__thread_bss"}, machocore.S_THREAD_LOCAL_ZEROFILL},
	".tlvdesc": {SegSect{machocore.SEG_DATA, "__thread_vars"}, machocore.S_THREAD_LOCAL_VARIABLES},
}

// dwarfNames is the ELF-to-Mach-O spelling of the DWARF sections.
var dwarfNames = map[string]string{
	".debug_abbrev":      "__debug_abbrev",
	".debug_addr":        "__debug_addr",
	".debug_aranges":     "__debug_aranges",
	".debug_cu_index":    "__debug_cu_index",
	".debug_frame":       "__debug_frame",
	".debug_info":        "__debug_info",
	".debug_line":        "__debug_line",
	".debug_line_str":    "__debug_line_str",
	".debug_loc":         "__debug_loc",
	".debug_loclists":    "__debug_loclists",
	".debug_macinfo":     "__debug_macinfo",
	".debug_macro":       "__debug_macro",
	".debug_names":       "__debug_names",
	".debug_pubnames":    "__debug_pubnames",
	".debug_pubtypes":    "__debug_pubtypes",
	".debug_ranges":      "__debug_ranges",
	".debug_rnglists":    "__debug_rnglists",
	".debug_str":         "__debug_str",
	".debug_str_offsets": "__debug_str_offs",
	".debug_tu_index":    "__debug_tu_index",
	".debug_types":       "__debug_types",
}

// ErrSectionName is a custom section name with no segment to put it in.
var ErrSectionName = obj.ErrSectionName

// placement resolves a section to a segment, a name, a type and its
// attributes.
func placement(s *obj.Section, opt Options) (SegSect, machocore.SecType, machocore.SecAttrs, error) {
	if ss, ok := opt.Sections[s.Name()]; ok {
		typ, attrs := shapeFor(s.Kind())
		return ss, typ, attrs, nil
	}

	if s.Name() == s.Kind().String() {
		ss, typ, attrs := kindPlacement(s.Kind())
		return ss, typ, attrs, nil
	}

	if tls, ok := tlsPlacement[s.Name()]; ok {
		return tls.ss, tls.typ, 0, nil
	}

	if name, ok := dwarfNames[s.Name()]; ok {
		return SegSect{machocore.SEG_DWARF, name}, machocore.S_REGULAR,
			machocore.S_ATTR_DEBUG, nil
	}
	if strings.HasPrefix(s.Name(), ".debug") {
		return SegSect{}, 0, 0, &obj.Error{
			Sentinel: ErrSectionName,
			Arch:     obj.ArchARM64,
			Section:  s.Name(),
			Context:  fmt.Sprintf("no __DWARF spelling for %q", s.Name()),
			Notes: []string{
				"give it one with Options.Sections: {\"" + s.Name() + "\": {\"__DWARF\", \"__...\"}}",
			},
		}
	}

	return SegSect{}, 0, 0, &obj.Error{
		Sentinel: ErrSectionName,
		Arch:     obj.ArchARM64,
		Section:  s.Name(),
		Context:  fmt.Sprintf("no segment for %q", s.Name()),
		Notes: []string{
			"a segment is a load-time protection decision and this container will not guess one",
			"name it with Options.Sections: {\"" + s.Name() + "\": {\"__DATA\", \"__" + strings.TrimPrefix(s.Name(), ".") + "\"}}",
		},
	}
}

func shapeFor(k obj.SectionKind) (machocore.SecType, machocore.SecAttrs) {
	_, typ, attrs := kindPlacement(k)
	return typ, attrs
}

func newSection(wr *machoobj.Writer, s *obj.Section, opt Options) (*machoobj.SectionBuilder, error) {
	ss, typ, attrs, err := placement(s, opt)
	if err != nil {
		return nil, err
	}
	return wr.Section(machoobj.SectionHeader{
		Segment: ss.Segment,
		Name:    ss.Section,
		Type:    typ,
		Attrs:   attrs,
		Align:   uint32(s.Align()),
	}), nil
}

// writeContents deposits every addend and hands the section its bytes.
func writeContents(wr *machoobj.Writer, b *machoobj.SectionBuilder, s *obj.Section, syms map[string]machoobj.SymRef) error {
	if b.Zerofill() {
		if len(s.Refs()) > 0 {
			return fmt.Errorf("macho: %s is zerofill and has no bytes to hold %d relocation addends",
				s.Name(), len(s.Refs()))
		}
		b.Grow(uint64(s.Size()))
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
