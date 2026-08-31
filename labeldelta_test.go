package arm64_test

import (
	"encoding/binary"
	"testing"

	"github.com/vertex-language/arm64"
)

// A table of distances, which is what a jump table is here: four bytes per
// entry, patched at Finalize, and no relocation left behind for a linker to
// refuse.
func TestLabelDelta(t *testing.T) {
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
	s.LabelDelta("table", "a")
	s.LabelDelta("table", "b")
	s.LabelDelta("table", "start")

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
func TestLabelDeltaUndefined(t *testing.T) {
	m := arm64.NewModule()
	s := m.Section(arm64.Text)
	s.Label("here", arm64.Local)
	s.LabelDelta("here", "nowhere")
	if _, err := m.Finalize(); err == nil {
		t.Error("Finalize should refuse a delta to an undefined label")
	}
}
