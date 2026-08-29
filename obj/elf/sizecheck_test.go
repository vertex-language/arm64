package elf_test

// A regression test, not a round trip: it builds a *obj.Object directly
// rather than through the high-level Module, because the Module never
// constructs a Reference whose Size disagrees with its Kind. This is the one
// place that contradiction can be manufactured to check the writer actually
// refuses it, the way obj/pe and obj/macho already do for this architecture.

import (
	"bytes"
	"testing"

	"github.com/vertex-language/arm64/obj"
	elfw "github.com/vertex-language/arm64/obj/elf"
)

func TestWriteRejectsSizeMismatch(t *testing.T) {
	o := obj.New(obj.ArchARM64, []obj.SectionData{
		{
			Name: ".text",
			Kind: obj.Text,
			// RefAbs64 fills an 8-byte field; claiming 4 here is the
			// contradiction under test.
			Bytes: []byte{0, 0, 0, 0},
			Refs:  []obj.Reference{{Offset: 0, Size: 4, Sym: "x", Kind: obj.RefAbs64}},
		},
	}, []obj.Symbol{{Name: "x", Section: 0}})

	if err := elfw.Write(&bytes.Buffer{}, o); err == nil {
		t.Error("Write should refuse a Reference whose Size disagrees with its Kind's field width")
	}
}
