// Package pe writes an *obj.Object as a COFF relocatable object.
//
// One function, pe.Write. The machine is derived from the object's Arch —
// ArchARM64 is IMAGE_FILE_MACHINE_ARM64 — and everything else about the
// header follows from it.
//
// Emission goes through github.com/vertex-language/pe, the same module that
// reads these objects back and that already carries this architecture's
// relocation vocabulary in reloc_aarch64.go, imported under an alias so the
// package a caller names is pe and the call site stays pe.Write.
package pe

import (
	"fmt"
	"io"

	pecore "github.com/vertex-language/pe"
	"github.com/vertex-language/pe/coff"

	"github.com/vertex-language/arm64/obj"
)

// Options are the things the assembler has no opinion about.
type Options struct {
	// ABI selects the toolchain convention the object is written for. The
	// zero value is ABIMSVC.
	ABI pecore.ABI

	// TimeDateStamp is written verbatim. Zero is the deterministic choice.
	TimeDateStamp uint32

	// BigObj decides the header family.
	BigObj coff.BigObjMode

	// Characteristics is the COFF file header's flag field.
	Characteristics pecore.FileChar

	// File is recorded as a .file symbol when non-empty.
	File string

	// Directives are linker options for the .drectve section.
	Directives []Directive
}

// Directive is one .drectve option.
type Directive struct {
	Name  string
	Value string
}

// ErrWeak is a weak symbol reaching a COFF writer. A COFF weak external is a
// name plus the alternate definition to use when nothing else defines it,
// and an object stating Weak and nothing else has not said what that
// alternate is.
var ErrWeak = fmt.Errorf("pe: a COFF weak external needs an alternate symbol")

// Write emits o to w as a relocatable COFF object.
func Write(w io.Writer, o *obj.Object, opts ...Options) error {
	if o == nil {
		return fmt.Errorf("pe: nil object")
	}
	if len(opts) > 1 {
		return fmt.Errorf("pe: Write takes at most one Options, got %d", len(opts))
	}
	var opt Options
	if len(opts) == 1 {
		opt = opts[0]
	}
	if o.Arch() != obj.ArchARM64 {
		return fmt.Errorf("pe: object is %s, not %s", o.Arch(), obj.ArchARM64)
	}

	abi := opt.ABI
	if abi == pecore.ABIUnknown {
		abi = pecore.ABIMSVC
	}

	wr := coff.NewWriter(w, coff.Options{
		Target: pecore.Target{
			Machine: pecore.MachineARM64,
			ABI:     abi,
			OS:      pecore.OSWindows,
		},
		BigObj:          opt.BigObj,
		TimeDateStamp:   opt.TimeDateStamp,
		Characteristics: opt.Characteristics,
	})

	if opt.File != "" {
		wr.FileSymbol(opt.File)
	}
	for _, d := range opt.Directives {
		wr.Directive(d.Name, d.Value)
	}

	secs := o.Sections()
	builders := make([]*coff.SectionBuilder, len(secs))
	for i, s := range secs {
		builders[i] = newSection(wr, s)
	}

	syms, err := writeSymbols(wr, o, builders)
	if err != nil {
		return err
	}

	for i, s := range secs {
		if err := writeContents(wr, builders[i], s, syms); err != nil {
			return err
		}
	}

	return wr.Close()
}
