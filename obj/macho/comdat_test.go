package macho_test

// Mach-O's spelling of a section the linker keeps once is no section at
// all: the bytes are folded into the ordinary section of the same placement
// and the leader is a weak definition, which ld64 coalesces by name. So
// there are two claims, and each is checked the way it can be -- the shape
// by reading the object back, the coalescing by handing ld64 two objects
// that both define the function and running what comes out.

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/vertex-language/arm64"
	"github.com/vertex-language/arm64/obj"
	machow "github.com/vertex-language/arm64/obj/macho"
	machocore "github.com/vertex-language/macho"
	machoobj "github.com/vertex-language/macho/obj"
)

// comdatModule is a unit that defines entry in .text and _inline_f in a
// COMDAT section of its own, returning 7.
func comdatModule(t *testing.T, entry string, callsInline bool) *obj.Object {
	t.Helper()
	m := arm64.NewModule()
	text := m.Section(arm64.Text)
	text.Label(entry, arm64.Global, arm64.Func)
	if callsInline {
		text.B(arm64.Ref("_inline_f"))
	} else {
		text.MovzImm32(arm64.W0, 0)
		text.Ret()
	}
	text.EndLabel(entry)

	inl := m.ComdatSection(".text", arm64.Text, "_inline_f")
	inl.Align(16)
	inl.Label("_inline_f", arm64.Global, arm64.Func)
	inl.MovzImm32(arm64.W0, 7)
	inl.Ret()
	inl.EndLabel("_inline_f")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	return o
}

func writeObject(t *testing.T, dir, name string, o *obj.Object) string {
	t.Helper()
	path := filepath.Join(dir, name)
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	err = machow.Write(f, o, machow.Options{
		Platform:    machocore.PlatformMacOS,
		MinOS:       "11.0",
		Subsections: true,
	})
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		t.Fatalf("macho.Write: %v", err)
	}
	return path
}

func TestComdat(t *testing.T) {
	dir := t.TempDir()
	path := writeObject(t, dir, "a.o", comdatModule(t, "_main", true))

	f, err := machoobj.Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer f.Close()

	var texts int
	for _, s := range f.Sections {
		if s.Name == "__text" {
			texts++
		}
	}
	if texts != 1 {
		t.Errorf("wrote %d __text sections, want the COMDAT one folded into 1", texts)
	}

	syms, err := f.Symbols()
	if err != nil {
		t.Fatalf("Symbols: %v", err)
	}
	for _, s := range syms {
		switch s.Name {
		case "_inline_f":
			if !s.WeakDef() {
				t.Error("_inline_f is the leader of a COMDAT section; want N_WEAK_DEF")
			}
			if s.Value != 16 {
				t.Errorf("_inline_f at %#x, want 16: folded after _main at its alignment", s.Value)
			}
		case "_main":
			if s.WeakDef() {
				t.Error("_main is ordinary; want no N_WEAK_DEF")
			}
		}
	}
}

// TestComdatCoalesces links two objects that both define _inline_f. Were
// either definition strong, ld64 would refuse the pair as a duplicate.
func TestComdatCoalesces(t *testing.T) {
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" {
		t.Skip("needs a native Apple Silicon host to link and run the result")
	}
	clangPath, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not on PATH")
	}

	dir := t.TempDir()
	a := writeObject(t, dir, "a.o", comdatModule(t, "_main", true))
	b := writeObject(t, dir, "b.o", comdatModule(t, "_other", false))
	bin := filepath.Join(dir, "prog")

	link := exec.Command(clangPath, "-o", bin, a, b)
	var stderr bytes.Buffer
	link.Stderr = &stderr
	if err := link.Run(); err != nil {
		t.Fatalf("clang link: %v\n%s", err, stderr.String())
	}

	err = exec.Command(bin).Run()
	code := 0
	if ee, ok := err.(*exec.ExitError); ok {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("running: %v", err)
	}
	if code != 7 {
		t.Errorf("exit %d, want 7 from the coalesced _inline_f", code)
	}
}

// A COMDAT section may come before the ordinary section it folds into --
// a class's static member ahead of the unit's own .data -- and the
// ordinary one still joins it rather than being taken for a second.
func TestComdatBeforeOrdinary(t *testing.T) {
	m := arm64.NewModule()
	c := m.ComdatSection(".data", arm64.Data, "_limit")
	c.Label("_limit", arm64.Global, arm64.ObjectSym)
	c.Long(3)

	d := m.Section(arm64.Data)
	d.Align(4)
	d.Label("_plain", arm64.Global, arm64.ObjectSym)
	d.Long(4)

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	path := writeObject(t, t.TempDir(), "d.o", o)

	f, err := machoobj.Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer f.Close()
	syms, err := f.Symbols()
	if err != nil {
		t.Fatalf("Symbols: %v", err)
	}
	for _, s := range syms {
		if s.Name == "_plain" && s.Value != 4 {
			t.Errorf("_plain at %#x, want 4: after the folded _limit", s.Value)
		}
	}
}
