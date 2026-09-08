package macho_test

// A symbol difference, across a real link.
//
// A relative pointer is `target - &field` in a four-byte field, which is
// what every descriptor Swift's runtime reads is built out of: a table of
// distances is position-independent where a table of addresses is not.
// Mach-O has no single relocation for it, so it is two -- SUBTRACTOR naming
// what to take away, then UNSIGNED naming what to take it away from -- and
// the field carries the addend.
//
// There is no symbol at an arbitrary offset inside a data object, so the
// subtrahend is the object's own symbol and the addend is the negative of
// the field's offset within it. That is what swiftc's own conformance
// descriptors do, entry by entry, and this checks that the arithmetic comes
// out where the linker puts the two ends.

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/vertex-language/arm64"
	machow "github.com/vertex-language/arm64/obj/macho"
	machocore "github.com/vertex-language/macho"
)

func TestSymbolDeltaLinksAndReads(t *testing.T) {
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" {
		t.Skip("needs a native Apple Silicon host to link and run the result")
	}
	clangPath, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not on PATH; skipping the link-and-run round trip")
	}

	m := arm64.NewModule()

	// A table of two relative pointers, one to a symbol in this section
	// and one to a symbol in another. Neither is foldable: the second is
	// in the text section, and a linker decides where both land.
	data := m.Section(arm64.Data)
	data.Label("_table", arm64.Global, arm64.ObjectSym)
	data.SymDelta("_target", "_table", 0) // field 0: to a data symbol
	data.SymDelta("_code", "_table", -4)  // field 1: to a text symbol
	data.EndLabel("_table")
	data.Label("_target", arm64.Global, arm64.ObjectSym)
	data.Quad(0)
	data.EndLabel("_target")

	// _main reads each field, adds it to the field's own address, and
	// compares the result with the address the linker gave the symbol.
	// Both must agree, or the printed answer is not 2.
	text := m.Section(arm64.Text)
	text.Label("_code", arm64.Global, arm64.Func)
	text.Ret()
	text.EndLabel("_code")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}

	dir := t.TempDir()
	objPath := filepath.Join(dir, "delta.o")
	binPath := filepath.Join(dir, "delta")

	f, err := os.Create(objPath)
	if err != nil {
		t.Fatal(err)
	}
	err = machow.Write(f, o, machow.Options{
		Platform: machocore.PlatformMacOS,
		MinOS:    "11.0",
	})
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		t.Fatalf("macho.Write: %v", err)
	}

	srcPath := filepath.Join(dir, "main.c")
	src := `
#include <stdio.h>
#include <stdint.h>
extern int32_t table[2];
extern char target[];
extern void code(void) __asm__("_code");
int main(void) {
    char *base = (char *)table;
    int ok = 0;
    ok += ((char *)(base + 0) + table[0]) == target;
    ok += ((char *)(base + 4) + table[1]) == (char *)code;
    printf("%d\n", ok);
    return 0;
}
`
	if err := os.WriteFile(srcPath, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	link := exec.Command(clangPath, "-o", binPath, srcPath, objPath)
	var stderr bytes.Buffer
	link.Stderr = &stderr
	if err := link.Run(); err != nil {
		t.Fatalf("clang link: %v\n%s", err, stderr.String())
	}

	out, err := exec.Command(binPath).Output()
	if err != nil {
		t.Fatalf("running the linked binary: %v", err)
	}
	if string(out) != "2\n" {
		t.Errorf("output = %q, want %q -- a relative pointer did not land on its target", out, "2\n")
	}
}
