package asm

import (
	"strings"
	"testing"

	"github.com/vertex-language/arm64"
)

// A fragment names a symbol the module around it already defines.
//
// The implicit extern is a guess about a name nothing defines — the rule
// that makes `bl puts` work with no `.extern puts` in front of it — and it
// has to stop at the module's own definitions. An asm goto is where this
// shows up: the block it branches to is a label in the same function, and
// which of the two the emitter reaches first is a question about block
// layout that the assembler has no business depending on.
func TestFragmentDoesNotExternWhatTheModuleDefines(t *testing.T) {
	for _, tc := range []struct {
		name  string
		first func(*arm64.Section)
	}{
		{"defined before the fragment", func(text *arm64.Section) {
			text.Label("target", arm64.Local)
			text.Ret()
		}},
		{"defined after it", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := arm64.NewModule()
			text := m.Section(arm64.Text)
			text.Label("f", arm64.Global, arm64.Func)

			if tc.first != nil {
				tc.first(text)
			}
			if err := AssembleFragment(text, "b target", Options{File: "t.s"}); err != nil {
				t.Fatalf("AssembleFragment: %v", err)
			}
			if tc.first == nil {
				text.Label("target", arm64.Local)
				text.Ret()
			}
			text.EndLabel("f")

			o, err := m.Finalize()
			if err != nil {
				t.Fatalf("Finalize: %v", err)
			}
			for _, sy := range o.Symbols() {
				if sy.Name == "target" && sy.Section < 0 {
					t.Error("target is undefined in the object; the module defines it")
				}
			}
		})
	}
}

// The rule it must not break: a name nothing in the module defines is still
// an import, which is what a call to a libc function relies on.
func TestFragmentStillExternsWhatNothingDefines(t *testing.T) {
	m := arm64.NewModule()
	text := m.Section(arm64.Text)
	text.Label("f", arm64.Global, arm64.Func)
	if err := AssembleFragment(text, "bl puts", Options{File: "t.s"}); err != nil {
		t.Fatalf("AssembleFragment: %v", err)
	}
	text.Ret()
	text.EndLabel("f")

	o, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	var names []string
	for _, sy := range o.Symbols() {
		if sy.Section < 0 {
			names = append(names, sy.Name)
		}
	}
	if strings.Join(names, ",") != "puts" {
		t.Errorf("undefined symbols = %v, want just puts", names)
	}
}
