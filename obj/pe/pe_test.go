package pe_test

// An external round trip, not a self-check: it writes a real COFF object and
// hands it to objdump, a tool this tree did not write. Skipped rather than
// failed when objdump is not on PATH — a missing reference tool is an
// environment fact, not a defect in this package. There is no link-and-run
// counterpart the way the Mach-O writer has one: this host has no ARM64
// Windows linker to hand the object to.

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vertex-language/arm64"
	pew "github.com/vertex-language/arm64/obj/pe"
)

func TestWriteReadWithObjdump(t *testing.T) {
	objdump, err := exec.LookPath("objdump")
	if err != nil {
		t.Skip("objdump not on PATH; skipping the external round trip")
	}

	m := arm64.NewModule()
	text := m.Section(arm64.Text)
	text.Label("main", arm64.Global, arm64.Func)
	text.StpPre64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Pre(-16))
	text.MovSp64(arm64.FP, arm64.SP)
	text.Adrp(arm64.X0, arm64.Ref("msg"))
	text.AddImm64(arm64.X0, arm64.X0, arm64.PageOff(arm64.Ref("msg")))
	text.Bl(arm64.Ref("puts"))
	text.MovzImm32(arm64.W0, 0)
	text.LdpPost64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Post(16))
	text.Ret()
	text.EndLabel("main")

	rodata := m.Section(arm64.ROData)
	rodata.Label("msg", arm64.Local, arm64.ObjectSym)
	rodata.Asciz("hello\n")

	m.Extern("puts")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "hello.obj")
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := pew.Write(f, o); err != nil {
		f.Close()
		t.Fatalf("pe.Write: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	relocs := runObjdump(t, objdump, "-r", path)
	for _, want := range []string{
		"IMAGE_REL_ARM64_PAGEBASE_REL21 msg",
		"IMAGE_REL_ARM64_PAGEOFFSET_12A msg",
		"IMAGE_REL_ARM64_BRANCH26",
		"puts",
	} {
		if !strings.Contains(relocs, want) {
			t.Errorf("objdump -r output missing %q\n%s", want, relocs)
		}
	}

	syms := runObjdump(t, objdump, "-t", path)
	for _, want := range []string{"main", "msg", "puts", ".rdata"} {
		if !strings.Contains(syms, want) {
			t.Errorf("objdump -t output missing %q\n%s", want, syms)
		}
	}

	disasm := runObjdump(t, objdump, "-d", path)
	for _, want := range []string{"stp", "adrp", "bl", "ret"} {
		if !strings.Contains(disasm, want) {
			t.Errorf("objdump -d output missing %q\n%s", want, disasm)
		}
	}
}

func runObjdump(t *testing.T, objdump string, args ...string) string {
	t.Helper()
	cmd := exec.Command(objdump, args...)
	var out, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("objdump %v: %v\n%s", args, err, stderr.String())
	}
	return out.String()
}
