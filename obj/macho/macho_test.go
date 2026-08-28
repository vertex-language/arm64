package macho_test

// This is a full round trip, not a self-check: it writes a real Mach-O
// object, links it with the host's own linker, and runs the result. On an
// Apple Silicon host that is a native link — no cross-toolchain, no emulator
// — so a mistake anywhere from the ISA table down to the relocation mapping
// has nowhere to hide: the process either prints what it was told to and
// exits zero, or it does not. It is skipped rather than failed off Darwin or
// off arm64, and when clang is not on PATH, because none of those is a
// defect in this package.

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

func TestWriteLinkRun(t *testing.T) {
	if runtime.GOOS != "darwin" || runtime.GOARCH != "arm64" {
		t.Skip("needs a native Apple Silicon host to link and run the result")
	}
	clangPath, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not on PATH; skipping the link-and-run round trip")
	}

	// Asciz is the bare bytes puts(3) is handed; puts appends the trailing
	// newline itself.
	const message = "hello from arm64 macho"
	const greeting = message + "\n"

	m := arm64.NewModule()
	text := m.Section(arm64.Text)
	text.Label("_main", arm64.Global, arm64.Func)
	text.StpPre64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Pre(-16))
	text.MovSp64(arm64.FP, arm64.SP)
	text.Adrp(arm64.X0, arm64.Ref("_msg"))
	text.AddImm64(arm64.X0, arm64.X0, arm64.PageOff(arm64.Ref("_msg")))
	text.Bl(arm64.Ref("_puts"))
	text.MovzImm32(arm64.W0, 0)
	text.LdpPost64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Post(16))
	text.Ret()
	text.EndLabel("_main")

	rodata := m.Section(arm64.ROData)
	rodata.Label("_msg", arm64.Local, arm64.ObjectSym)
	rodata.Asciz(message)

	m.Extern("_puts")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}

	dir := t.TempDir()
	objPath := filepath.Join(dir, "hello.o")
	binPath := filepath.Join(dir, "hello")

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

	link := exec.Command(clangPath, "-o", binPath, objPath)
	var stderr bytes.Buffer
	link.Stderr = &stderr
	if err := link.Run(); err != nil {
		t.Fatalf("clang link: %v\n%s", err, stderr.String())
	}

	run := exec.Command(binPath)
	out, err := run.Output()
	if err != nil {
		t.Fatalf("running the linked binary: %v", err)
	}
	if string(out) != greeting {
		t.Errorf("output = %q, want %q", out, greeting)
	}
}
