package arm64

import "github.com/vertex-language/arm64/obj"

// symbol is one entry in the module-level symbol table. Identity is the name,
// module-wide.
type symbol struct {
	name string
	sec  int // -1 when undefined
	off  int

	size       int
	sizeClosed bool

	binding Binding
	typ     SymbolType
	vis     Visibility

	defined bool
}

type aliasReq struct{ name, of string }
type visReq struct {
	name string
	vis  Visibility
}

// define promotes a label into the symbol table.
//
// The attributes are variadic any because a Binding, a SymbolType and a
// Visibility are three different types and a caller writes whichever they
// mean, in whatever order. It is the one place in this package where a type
// switch stands in for a signature, and the alternative — three optional
// parameters, or three methods — reads worse at every call site.
func (m *Module) define(s *Section, name string, off int, attrs ...any) {
	if i, ok := m.symAt[name]; ok && m.symbols[i].defined {
		prev := m.symbols[i]
		m.fail(s.errorAt(obj.ErrDuplicate,
			"symbol "+name+" is already defined",
			"the first definition is at "+m.sections[prev.sec].name+"+"+hex(prev.off)))
		return
	}

	sym := symbol{name: name, sec: s.index, off: off, defined: true, binding: Local}
	for _, a := range attrs {
		switch v := a.(type) {
		case Binding:
			sym.binding = v
		case SymbolType:
			sym.typ = v
		case Visibility:
			sym.vis = v
		default:
			m.fail(s.errorAt(obj.ErrOperand,
				"Label("+name+") was given something that is not a symbol attribute",
				"attributes are a Binding, a SymbolType or a Visibility"))
			return
		}
	}

	// Extern followed by a definition: the definition wins and fills in the
	// entry that was already there, so the symbol keeps its table position.
	if i, ok := m.symAt[name]; ok {
		m.symbols[i] = sym
		return
	}
	m.symAt[name] = len(m.symbols)
	m.symbols = append(m.symbols, sym)
}

// closeSizes fills in every size that EndLabel did not.
func (m *Module) closeSizes() error {
	for i := range m.symbols {
		sym := &m.symbols[i]
		if !sym.defined || sym.sizeClosed {
			continue
		}
		end := len(m.sections[sym.sec].buf)
		for j := range m.symbols {
			o := m.symbols[j]
			if !o.defined || o.sec != sym.sec || o.off <= sym.off {
				continue
			}
			if o.off < end {
				end = o.off
			}
		}
		sym.size = end - sym.off
	}
	return nil
}

// resolveAliases gives a symbol a second name at the same offset.
func (m *Module) resolveAliases() error {
	for _, a := range m.aliases {
		i, ok := m.symAt[a.of]
		if !ok || !m.symbols[i].defined {
			return &obj.Error{
				Arch: obj.ArchARM64, Sentinel: obj.ErrUndefined,
				Context: "Alias(" + a.name + ", " + a.of + ") names no defined symbol",
			}
		}
		if j, dup := m.symAt[a.name]; dup && m.symbols[j].defined {
			return &obj.Error{
				Arch: obj.ArchARM64, Sentinel: obj.ErrDuplicate,
				Context: "Alias(" + a.name + ") names a symbol that is already defined",
			}
		}
		src := m.symbols[i]
		src.name = a.name
		m.symAt[a.name] = len(m.symbols)
		m.symbols = append(m.symbols, src)
	}
	return nil
}

func (m *Module) resolveVisibility() error {
	for _, v := range m.visReqs {
		i, ok := m.symAt[v.name]
		if !ok {
			return &obj.Error{
				Arch: obj.ArchARM64, Sentinel: obj.ErrUndefined,
				Context: "SetVisibility(" + v.name + ") names no symbol",
			}
		}
		m.symbols[i].vis = v.vis
	}
	return nil
}

// verifyRefs checks that every surviving reference names something this
// module knows about.
func (m *Module) verifyRefs() error {
	for _, s := range m.sections {
		for _, r := range s.refs {
			if _, ok := m.symAt[r.Sym]; !ok {
				return &obj.Error{
					Arch: obj.ArchARM64, Section: s.name, Offset: r.Offset,
					Sentinel: obj.ErrUndefined,
					Context:  "reference to " + r.Sym + " names nothing",
					Notes: []string{
						"define it with Label, or declare it with Extern",
					},
				}
			}
		}
	}
	return nil
}
