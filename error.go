package arm64

import (
	"errors"

	"github.com/vertex-language/arm64/internal/encode"
	"github.com/vertex-language/arm64/internal/isa"
	"github.com/vertex-language/arm64/obj"
)

// fail records the first error and keeps it.
//
// Errors are sticky and first-wins: every builder call after a failure is a
// no-op, and Finalize surfaces the first one, positioned. That is what lets
// an instruction-selection loop run without a single error check in it.
func (m *Module) fail(err error) {
	if m.err == nil {
		m.err = err
	}
}

// errorAt builds a positioned diagnostic.
//
//	arm64 .text+0x14: ADRP page displacement does not fit R_AARCH64_ADR_PREL_PG_HI21
//	  note: the field is 21 bits after the 12-bit page shift; the range is -1GiB..1GiB-4KiB
func (s *Section) errorAt(sentinel error, context string, notes ...string) *obj.Error {
	return &obj.Error{
		Arch:     obj.ArchARM64,
		Section:  s.name,
		Offset:   len(s.buf),
		Context:  context,
		Sentinel: sentinel,
		Notes:    notes,
	}
}

// lift turns an error from the resolver or the encoder into a positioned one.
//
// A resolver or encoder failure reaches a caller through *obj.Error as a
// sentinel, a message and notes — data, not internal types. The internal
// error joins the chain as the cause, so Unwrap returns both it and the
// sentinel and errors.Is works against either, but nothing a caller needs is
// only reachable by asserting on it.
func (s *Section) lift(err error) *obj.Error {
	if err == nil {
		return nil
	}
	var oe *obj.Error
	if errors.As(err, &oe) {
		return oe
	}
	return &obj.Error{
		Arch:     obj.ArchARM64,
		Section:  s.name,
		Offset:   len(s.buf),
		Context:  err.Error(),
		Sentinel: sentinelFor(err),
		Cause:    err,
	}
}

// sentinelFor maps the resolver's and encoder's typed errors onto the
// sentinels errors.Is answers for.
func sentinelFor(err error) error {
	var (
		unknown    *isa.UnknownError
		formErr    *isa.FormError
		gate       *isa.GateError
		count      *encode.CountError
		operandErr *encode.OperandError
		regErr     *encode.RegisterError
		rangeErr   *encode.RangeError
		bitmask    *encode.BitmaskError
		addrErr    *encode.AddressError
		unsup      *encode.UnsupportedError
	)
	switch {
	case errors.As(err, &unknown), errors.As(err, &formErr), errors.As(err, &count):
		return obj.ErrForm
	case errors.As(err, &gate):
		return obj.ErrFeature
	case errors.As(err, &rangeErr):
		return obj.ErrRange
	case errors.As(err, &bitmask):
		return obj.ErrBitmask
	case errors.As(err, &operandErr), errors.As(err, &regErr), errors.As(err, &addrErr), errors.As(err, &unsup):
		return obj.ErrOperand
	}
	return obj.ErrForm
}

// Err returns the first error, or nil. It is the same error Finalize will
// return, and it is there so a long build can bail out early rather than
// running to completion over a module that is already spent.
func (m *Module) Err() error { return m.err }
