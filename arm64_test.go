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

// A value in a halfword above the first, through the typed API.
//
// The immediate encoder computes the halfword's index and writes it into Hw
// as a sibling of the value; the slot it lands in is an *optional* shift
// slot, and an optional slot with no operand left for it places its default.
// Both were true at once here, so the default landed on top of the index and
// `movz x0, #0x10000` assembled as `movz x0, #1` — a different instruction,
// silently, through every door this package has.
func TestMovzShiftedHalfword(t *testing.T) {
	for _, tc := range []struct {
		name string
		emit func(*Section)
		want uint32
	}{
		{"movz x0, #0x10000", func(s *Section) { s.MovzImm64(X0, 0x10000) }, 0xd2a00020},
		{"movz x0, #0x1000000000000", func(s *Section) { s.MovzImm64(X0, 0x1000000000000) }, 0xd2e00020},
		{"movz w0, #0x30000", func(s *Section) { s.MovzImm32(W0, 0x30000) }, 0x52a00060},
		{"movz x0, #1 keeps hw zero", func(s *Section) { s.MovzImm64(X0, 1) }, 0xd2800020},
		{"movk x0, #0x20000", func(s *Section) { s.MovkImm64(X0, 0x20000) }, 0xf2a00040},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := NewModule()
			text := m.Section(Text)
			text.Label("f", Global, Func)
			tc.emit(text)
			text.EndLabel("f")

			o, err := m.Finalize()
			if err != nil {
				t.Fatalf("Finalize: %v", err)
			}
			var b []byte
			for _, s := range o.Sections() {
				if len(s.Bytes()) >= 4 {
					b = s.Bytes()
				}
			}
			if got := binary.LittleEndian.Uint32(b[:4]); got != tc.want {
				t.Errorf("= %#08x, want %#08x", got, tc.want)
			}
		})
	}
}
