package macho

import (
	"testing"

	machocore "github.com/vertex-language/macho"

	"github.com/vertex-language/arm64/obj"
)

// TestRelROPlacement: a section holding relocated constants goes to
// (__DATA,__const), not (__TEXT,__const).
//
// __TEXT is the text segment, mapped read-only and executable, so a
// table of addresses placed there cannot be rebased -- a program that
// called through one died on SIGBUS. clang emits (__DATA,__const) for
// a `const` array of function pointers, and ld moves that into the
// __DATA_CONST segment when it builds the image.
func TestRelROPlacement(t *testing.T) {
	tests := []struct {
		kind obj.SectionKind
		seg  string
		sect string
	}{
		{obj.ROData, machocore.SEG_TEXT, machocore.SECT_CONST},
		{obj.RelROData, machocore.SEG_DATA, machocore.SECT_CONST},
		{obj.Data, machocore.SEG_DATA, machocore.SECT_DATA},
	}
	for _, tt := range tests {
		t.Run(tt.kind.String(), func(t *testing.T) {
			ss, typ, _ := kindPlacement(tt.kind)
			if ss.Segment != tt.seg || ss.Section != tt.sect {
				t.Errorf("kindPlacement(%v) = (%s,%s), want (%s,%s)",
					tt.kind, ss.Segment, ss.Section, tt.seg, tt.sect)
			}
			if typ != machocore.S_REGULAR {
				t.Errorf("type = %v, want S_REGULAR", typ)
			}
		})
	}
	// And it is writable, because the loader writes it before it is
	// read-only.
	if !obj.RelROData.Writable() {
		t.Error("RelROData is not writable; the loader could not relocate it")
	}
}
