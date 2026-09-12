package elf_test

// TestComdat is the ELF spelling of a section the linker keeps once: a
// GRP_COMDAT group signed by the leader, holding the section and whatever
// is associated with it.

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/vertex-language/arm64"
	"github.com/vertex-language/arm64/obj/elf"
	elfobj "github.com/vertex-language/elf/obj"
)

func TestComdat(t *testing.T) {
	m := arm64.NewModule()
	m.Section(arm64.Text).Label("main", arm64.Global, arm64.Func)
	m.Section(arm64.Text).Ret()

	inl := m.ComdatSection(".text", arm64.Text, "inline_f")
	inl.Label("inline_f", arm64.Global, arm64.Func)
	inl.Ret()
	inl.EndLabel("inline_f")

	eh := m.AssociativeSection(".rodata", arm64.ROData, inl)
	eh.Label("$eh$inline_f", arm64.Local, arm64.ObjectSym)
	eh.Long(0)

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	var buf bytes.Buffer
	if err := elf.Write(&buf, o); err != nil {
		t.Fatalf("Write: %v", err)
	}
	path := filepath.Join(t.TempDir(), "t.o")
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}

	f, err := elfobj.Open(path)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer f.Close()

	groups, err := f.Groups()
	if err != nil {
		t.Fatalf("Groups: %v", err)
	}
	if len(groups) != 1 {
		t.Fatalf("wrote %d groups, want 1", len(groups))
	}
	g := groups[0]
	if !g.COMDAT() || g.Key() != "inline_f" {
		t.Errorf("group is COMDAT=%v key=%q, want a COMDAT group keyed on inline_f", g.COMDAT(), g.Key())
	}
	if len(g.Members) != 2 {
		t.Errorf("group has %d members, want the function's section and its associated one", len(g.Members))
	}
}

// A leader the section does not define is refused at Finalize.
func TestComdatLeaderMustBeDefined(t *testing.T) {
	m := arm64.NewModule()
	inl := m.ComdatSection(".text", arm64.Text, "elsewhere")
	inl.Label("inline_f", arm64.Global, arm64.Func)
	inl.Ret()
	inl.EndLabel("inline_f")
	if _, err := m.Finalize(); err == nil {
		t.Error("Finalize accepted a COMDAT section elected on a symbol it does not define")
	}
}
