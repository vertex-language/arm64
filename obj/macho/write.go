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
	places, err := placeSections(wr, secs, opt)
	if err != nil {
		return err
	}

	syms, err := writeSymbols(wr, o, places)
	if err != nil {
		return err
	}

	for i, s := range secs {
		if err := writeContents(wr, places[i], s, syms); err != nil {
			return err
		}
	}

	return wr.Close()
}

// A place is where one of the object's sections landed in the file: which
// Mach-O section, and at what offset into it.
//
// Mach-O has no COMDAT. What it has is a weak definition, which ld64
// coalesces by symbol, and a section is one per (segment, name) pair -- so
// a COMDAT section, and anything associated with it, is folded into the
// ordinary section of the same placement at the next aligned offset, and
// its symbols become weak. Every symbol and relocation of the folded
// section moves by that offset; nothing else changes, since references
// name symbols and not sections.
type place struct {
	b      *machoobj.SectionBuilder
	offset uint64
}

// placeSections creates one Mach-O section per placement and assigns every
// object section an offset in it.
func placeSections(wr *machoobj.Writer, secs []*obj.Section, opt Options) ([]place, error) {
	type slot struct {
		b        *machoobj.SectionBuilder
		size     uint64
		ordinary bool
	}
	byPlacement := map[SegSect]*slot{}
	places := make([]place, len(secs))

	for i, s := range secs {
		ss, _, _, err := placement(s, opt)
		if err != nil {
			return nil, err
		}
		folded := s.Comdat() != "" || s.Associated() != nil
		sl, exists := byPlacement[ss]
		if !exists {
			b, err := newSection(wr, s, opt)
			if err != nil {
				return nil, err
			}
			sl = &slot{b: b}
			byPlacement[ss] = sl
		} else if !folded && sl.ordinary {
			// Two ordinary sections with one placement is a caller's
			// mistake and was one before COMDAT existed here; only a
			// folded section joins another. Which of them came first is
			// not the question: a vtable's COMDAT .data may well precede
			// the unit's own .data.
			return nil, fmt.Errorf("macho: %s and an earlier section both place at %s", s.Name(), ss)
		}
		if !folded {
			sl.ordinary = true
		}
		off := sl.size
		if a := uint64(s.Align()); a > 1 && off%a != 0 {
			off += a - off%a
		}
		places[i] = place{b: sl.b, offset: off}
		sl.size = off + uint64(s.Size())
	}
	return places, nil
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
