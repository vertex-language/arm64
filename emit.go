package arm64

import (
	"encoding/binary"

	"github.com/vertex-language/arm64/internal/encode"
)

// Emit resolves a form at run time from a mnemonic and operands.
//
// It is the escape hatch for table-driven emission, where the mnemonic is
// data. If you know the instruction at compile time, the typed helper is the
// surface: it pins its form, and a width or class mismatch is a compile
// error rather than an ErrForm at run time.
//
// Because the table refuses ambiguity when it is built, Emit matches exactly
// one form or fails, naming the near-miss candidates; a form that matches but
// is gated returns ErrFeature rather than "no such form" — being told an
// instruction does not exist when it exists and is disabled sends a reader
// looking for a typo that is not there.
//
// A label target is spelled arm64.Label("done"); a bare string is refused,
// because a bare name could be a label or a register.
func (s *Section) Emit(mnem string, ops ...any) {
	if !s.ok() {
		return
	}
	word, fixups, err := encode.EncodeWith(
		encode.Opts{Offset: len(s.buf)}, s.m.features, mnem, ops...)
	if err != nil {
		s.m.fail(s.lift(err))
		return
	}
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], word)
	s.buf = append(s.buf, b[:]...)
	for _, fx := range fixups {
		s.pending = append(s.pending, pending{fx: fx, mnem: mnem})
	}
}

// Inst states a whole word rather than naming an instruction — the .inst
// directive, for the encoding this table has not reached yet. It is routed
// through the table's own row so it goes through the same gate and the same
// section bookkeeping as every typed helper.
func (s *Section) Inst(word uint32) {
	s.inst(instForm, uint64(word))
}
