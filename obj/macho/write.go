// Package macho writes an *obj.Object as a Mach-O relocatable object.
//
// One function, macho.Write. The CPU is derived from the object's Arch —
// ArchARM64 is CPU_TYPE_ARM64 — and the header follows from it.
//
// Emission goes through github.com/vertex-language/macho, the same module
// that reads these objects back and that Apple's own toolchain's arm64
// backend already carries the relocation vocabulary for, imported under an
// alias so the package a caller names is macho and the call site stays
// macho.Write.
package macho

import (
	"fmt"
	"io"

	machocore "github.com/vertex-language/macho"
	machoobj "github.com/vertex-language/macho/obj"

	"github.com/vertex-language/arm64/obj"
)

// SegSect is a Mach-O section identity: the segment it lives in and the name
// it carries there.
type SegSect struct {
	Segment string
	Section string
}

// Options are the things the assembler has no opinion about.
type Options struct {
	// Platform is required. An object with no platform makes every linker
	// guess, and ld64 warns about exactly that.
	Platform machocore.Platform

	// MinOS is the deployment target, e.g. "11.0".
	MinOS string

	// SDK is the SDK version the object was built against. Optional.
	SDK string

	// Subsections sets MH_SUBSECTIONS_VIA_SYMBOLS, which says the sections
	// may be cut at their symbols.
	//
	// Set it for anything with unwind tables in it. Without it a linker
	// cannot see where one function ends and the next begins, and
	// __unwind_info's sentinel — the entry that says how far the last
	// function reaches — lands at the start of the last function instead
	// of past it. An exception thrown through that function then finds no
	// unwind information and the process terminates.
	Subsections bool

	// Sections gives a segment and a name to a section this package has no
	// mapping for.
	Sections map[string]SegSect
}

// Write emits o to w as an MH_OBJECT file.
func Write(w io.Writer, o *obj.Object, opts ...Options) error {
	if o == nil {
		return fmt.Errorf("macho: nil object")
	}
	if len(opts) > 1 {
		return fmt.Errorf("macho: Write takes at most one Options, got %d", len(opts))
	}
	var opt Options
	if len(opts) == 1 {
		opt = opts[0]
	}
	if o.Arch() != obj.ArchARM64 {
		return fmt.Errorf("macho: object is %s, not %s", o.Arch(), obj.ArchARM64)
	}

	target, err := targetFor(opt)
	if err != nil {
		return err
	}

	var flags machocore.Flags
	if opt.Subsections {
		flags |= machocore.MH_SUBSECTIONS_VIA_SYMBOLS
	}

	wr := machoobj.NewWriter(w, machoobj.Options{
		Target: target,
		Flags:  flags,
		Build:  target.Build(),
	})

	secs := o.Sections()
	builders := make([]*machoobj.SectionBuilder, len(secs))
	for i, s := range secs {
		b, err := newSection(wr, s, opt)
		if err != nil {
			return err
		}
		builders[i] = b
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

// targetFor builds the Mach-O target from the options and this package's one
// architecture.
func targetFor(opt Options) (machocore.Target, error) {
	if opt.Platform == machocore.PlatformUnknown {
		return machocore.Target{}, fmt.Errorf("macho: Options.Platform is required")
	}
	if opt.MinOS == "" {
		return machocore.Target{}, fmt.Errorf("macho: Options.MinOS is required, e.g. \"11.0\"")
	}
	minOS, err := machocore.ParseVersion(opt.MinOS)
	if err != nil {
		return machocore.Target{}, fmt.Errorf("macho: Options.MinOS: %w", err)
	}
	var sdk machocore.Version
	if opt.SDK != "" {
		if sdk, err = machocore.ParseVersion(opt.SDK); err != nil {
			return machocore.Target{}, fmt.Errorf("macho: Options.SDK: %w", err)
		}
	}
	return machocore.Target{
		CPU:      machocore.CPU_TYPE_ARM64,
		SubCPU:   machocore.CPU_SUBTYPE_ARM64_ALL,
		Platform: opt.Platform,
		MinOS:    minOS,
		SDK:      sdk,
		Endian:   machocore.LittleEndian,
	}, nil
}
