package macho

import (
	"fmt"

	machocore "github.com/vertex-language/macho"
	machoobj "github.com/vertex-language/macho/obj"

	"github.com/vertex-language/arm64/obj"
)

// Mach-O spells linkage as two bits on the type byte; see amd64/obj/macho's
// symbol.go for the full reasoning, which is identical here since none of it
// is architecture-specific.
func writeSymbols(wr *machoobj.Writer, o *obj.Object, builders []*machoobj.SectionBuilder) (map[string]machoobj.SymRef, error) {
	syms := o.Symbols()
	out := make(map[string]machoobj.SymRef, len(syms))

	for _, sym := range syms {
		def := machoobj.SymbolDef{
			Name: sym.Name,
			Type: machocore.N_UNDF,
			Ext:  sym.Binding != obj.Local,
			Pext: sym.Binding != obj.Local && sym.Visibility == obj.Hidden,
		}

		if sym.Defined() {
			if sym.Section < 0 || sym.Section >= len(builders) {
				return nil, fmt.Errorf("macho: symbol %q names section %d of %d",
					sym.Name, sym.Section, len(builders))
			}
			def.Type = machocore.N_SECT
			def.Section = builders[sym.Section]
			def.Value = uint64(sym.Offset)
		}

		if sym.Binding == obj.Weak {
			if sym.Defined() {
				def.WeakDef = true
			} else {
				def.WeakRef = true
			}
		}

		out[sym.Name] = wr.Symbol(def)
	}
	return out, wr.Err()
}
