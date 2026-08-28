package arm64

import (
	"encoding/binary"
	"errors"
	"testing"
)

// TestSmokeLeafFunction builds a small leaf function — the AArch64 hello:
// mov x0, #imm; ret — and checks the encoded words against the ARM ARM's own
// worked bit patterns, so a mistake in the ISA table or the typed helpers
// fails a build rather than surfacing as a wrong-looking disassembly.
func TestSmokeLeafFunction(t *testing.T) {
	m := NewModule()
	text := m.Section(Text)

	text.Label("main", Global, Func)
	text.MovzImm64(X0, 42)
	text.Ret()
	text.EndLabel("main")

	obj, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}

	sec := obj.SectionAt(0)
	b := sec.Bytes()
	if len(b) != 8 {
		t.Fatalf("got %d bytes, want 8", len(b))
	}

	movz := binary.LittleEndian.Uint32(b[0:4])
	ret := binary.LittleEndian.Uint32(b[4:8])

	// mov x0, #42 == movz x0, #42: 0xd2800540
	if want := uint32(0xd2800540); movz != want {
		t.Errorf("movz x0, #42 = %#08x, want %#08x", movz, want)
	}
	// ret == ret x30: 0xd65f03c0
	if want := uint32(0xd65f03c0); ret != want {
		t.Errorf("ret = %#08x, want %#08x", ret, want)
	}

	syms := obj.Symbols()
	if len(syms) != 1 || syms[0].Name != "main" || syms[0].Size != 8 {
		t.Errorf("symbols = %+v, want one \"main\" of size 8", syms)
	}
}

// TestSmokeBranchFold exercises a same-section branch: the label is defined
// after the branch that names it, so Finalize has to patch the displacement
// in rather than the typed helper computing it at the call.
func TestSmokeBranchFold(t *testing.T) {
	m := NewModule()
	text := m.Section(Text)

	text.Cbz64(X0, Label("done"))
	text.MovzImm64(X0, 1)
	text.Label("done")
	text.Ret()

	obj, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	b := obj.SectionAt(0).Bytes()
	if len(b) != 12 {
		t.Fatalf("got %d bytes, want 12", len(b))
	}
	cbz := binary.LittleEndian.Uint32(b[0:4])
	// cbz x0, +8 (two instructions ahead): imm19 = 2, encoded at bits 23:5.
	if want := uint32(0xb4000040); cbz != want {
		t.Errorf("cbz x0, done = %#08x, want %#08x", cbz, want)
	}
}

// TestSmokeUndefinedReference checks that a branch to a symbol this module
// never defines or declares extern fails Finalize with ErrUndefined, rather
// than silently emitting a zero displacement.
func TestSmokeUndefinedReference(t *testing.T) {
	m := NewModule()
	text := m.Section(Text)
	text.Bl(Ref("memcpy"))

	if _, err := m.Finalize(); err == nil {
		t.Fatal("Finalize succeeded on a reference to an undeclared symbol")
	} else if !errors.Is(err, ErrUndefined) {
		t.Errorf("error = %v, want ErrUndefined", err)
	}
}
