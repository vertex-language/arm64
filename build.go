package arm64

import "github.com/vertex-language/arm64/obj"

// build turns the finished module into the artifact.
//
// It runs after every Finalize step has passed: labels patched, sizes
// closed, aliases and visibility resolved, every reference verified. So
// there is nothing to check here and nothing that can fail — this is a
// translation, and the only reason it is a separate function is that the
// builder's storage and the artifact's are different shapes.
func (m *Module) build() *obj.Object {
	secs := make([]obj.SectionData, 0, len(m.sections))
	for _, s := range m.sections {
		secs = append(secs, obj.SectionData{
			Name:  s.name,
			Kind:  s.kind,
			Align: s.align,
			Bytes: s.buf,
			Refs:  s.refs,
		})
	}

	syms := make([]obj.Symbol, 0, len(m.symbols))
	for _, sym := range m.symbols {
		syms = append(syms, obj.Symbol{
			Name:       sym.name,
			Section:    sym.sec,
			Offset:     sym.off,
			Size:       sym.size,
			Binding:    sym.binding,
			Type:       sym.typ,
			Visibility: sym.vis,
		})
	}

	return obj.New(obj.ArchARM64, secs, syms)
}
