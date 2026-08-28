package pe

import (
	"fmt"

	pecore "github.com/vertex-language/pe"
	"github.com/vertex-language/pe/coff"

	"github.com/vertex-language/arm64/obj"
)

// COFF spells linkage as a storage class; see amd64/obj/pe's symbol.go for
// the full reasoning, which is identical here.
func classOf(sym obj.Symbol) (pecore.StorageClass, error) {
	switch sym.Binding {
	case obj.Global:
		return pecore.ClassExternal, nil
	case obj.Local:
		return pecore.ClassStatic, nil
	case obj.Weak:
		return 0, fmt.Errorf("%w: %q", ErrWeak, sym.Name)
	}
	return pecore.ClassExternal, nil
}

func symTypeOf(t obj.SymbolType) pecore.SymType {
	if t == obj.Func {
		return pecore.PackSymType(pecore.BaseNull, pecore.DerivedFunction)
	}
	return pecore.PackSymType(pecore.BaseNull, pecore.DerivedNull)
}

func writeSymbols(wr *coff.Writer, o *obj.Object, builders []*coff.SectionBuilder) (map[string]*coff.SymbolRef, error) {
	syms := o.Symbols()
	out := make(map[string]*coff.SymbolRef, len(syms))

	for _, sym := range syms {
		class, err := classOf(sym)
		if err != nil {
			return nil, err
		}

		def := coff.SymbolDef{
			Name:  sym.Name,
			Class: class,
			Type:  symTypeOf(sym.Type),
		}
		if sym.Defined() {
			if sym.Section < 0 || sym.Section >= len(builders) {
				return nil, fmt.Errorf("pe: symbol %q names section %d of %d",
					sym.Name, sym.Section, len(builders))
			}
			def.Section = builders[sym.Section]
			def.Value = uint32(sym.Offset)
		}

		out[sym.Name] = wr.Symbol(def)
	}
	return out, wr.Err()
}
