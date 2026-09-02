package arm64_test

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/vertex-language/arm64"
)

// A table of distances, which is what a jump table is here: four bytes per
// entry, patched at Finalize, and no relocation left behind for a linker to
// refuse.
func TestLabelDiff(t *testing.T) {
	m := arm64.NewModule()
	s := m.Section(arm64.Text)
	s.Label("start", arm64.Global, arm64.Func)
	s.Nop() // +0
	s.Label("a")
	s.Nop() // +4
	s.Label("b")
	s.Nop() // +8
	s.EndLabel("start")

	s.Label("table", arm64.Local)
	s.LabelDiff("a", "table")
	s.LabelDiff("b", "table")
	s.LabelDiff("start", "table")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	text := o.SectionNamed(".text")
	if text == nil {
		t.Fatal("no .text")
	}
	data := text.Bytes()
	// The table begins at offset 12: three NOPs before it.
	got := []int32{
		int32(binary.LittleEndian.Uint32(data[12:])),
		int32(binary.LittleEndian.Uint32(data[16:])),
		int32(binary.LittleEndian.Uint32(data[20:])),
	}
	want := []int32{4 - 12, 8 - 12, 0 - 12}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("entry %d = %d, want %d", i, got[i], want[i])
		}
	}
	for _, r := range text.Refs() {
		t.Errorf("unexpected relocation %v", r)
	}
}

// An undefined label is refused by name rather than patched as zero.
func TestLabelDiffUndefined(t *testing.T) {
	m := arm64.NewModule()
	s := m.Section(arm64.Text)
	s.Label("here", arm64.Local)
	s.LabelDiff("nowhere", "here")
	if _, err := m.Finalize(); err == nil {
		t.Error("Finalize should refuse a delta to an undefined label")
	}
}

// TBZ and TBNZ at both widths, against the words clang produces for the same
// two lines with the label immediately after them.
//
// Here rather than in difftest_test.go because these take a branch target,
// and that harness assembles its cases as one straight run of instructions
// with no label for a target to name.
func TestTestBitBranchWidths(t *testing.T) {
	m := arm64.NewModule()
	s := m.Section(arm64.Text)
	s.Label("start", arm64.Global, arm64.Func)
	s.Tbnz32(arm64.W4, 5, arm64.Label("after"))
	s.Tbz64(arm64.X2, 40, arm64.Label("after"))
	s.Label("after")
	s.EndLabel("start")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	data := o.SectionNamed(".text").Bytes()

	// clang, for "tbnz w4, #5, 1f" and "tbz x2, #40, 1f" with 1: after both.
	want := []uint32{0x37280044, 0xb6400022}
	for i, w := range want {
		if got := binary.LittleEndian.Uint32(data[i*4:]); got != w {
			t.Errorf("word %d = %#08x, want %#08x", i, got, w)
		}
	}
}

// A bit number the register does not have is a range error, not a word that
// silently means the other width.
func TestTestBitBranchBitOutOfRange(t *testing.T) {
	m := arm64.NewModule()
	s := m.Section(arm64.Text)
	s.Label("start", arm64.Global, arm64.Func)
	s.Tbnz32(arm64.W4, 40, arm64.Label("start"))

	if err := m.Err(); err == nil {
		t.Fatal("bit 40 of a W register was accepted")
	}
}

// A bitfield that would run off the end of the register is a range error
// naming what would fit, not a truncated imms field.
func TestBitfieldExtractWidth(t *testing.T) {
	m := arm64.NewModule()
	s := m.Section(arm64.Text)
	s.Label("start", arm64.Global, arm64.Func)
	s.Ubfx32(arm64.W0, arm64.W1, 30, 5) // bits 30..34 of a 32-bit register

	err := m.Err()
	if err == nil {
		t.Fatal("a five-bit field starting at bit 30 of a W register was accepted")
	}
	if !strings.Contains(err.Error(), "starting at bit 30") {
		t.Errorf("error %q does not say where the field started", err)
	}
}
