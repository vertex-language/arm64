package asm

import (
	"strings"

	"github.com/vertex-language/arm64"
	"github.com/vertex-language/arm64/obj"
	"github.com/vertex-language/arm64/operand"
	"github.com/vertex-language/gas"
)

// target is the gas.Target half: the dialect, and everything after a mnemonic.
type target struct{ e *emitter }

func (t *target) Syntax() gas.Syntax { return gas.UAL }

// Inst parses one instruction's operands and emits it.
//
// The whole of the encoding decision is Section.Emit's: this function's only
// job is to turn text into the operand types the table matches against, and
// then to let the table pick the form. That is why a width mismatch is
// diagnosed by the parent package and not here — `add x0, w1, x2` is not a
// syntax error, it is an instruction that does not exist, and the table is
// what knows the difference.
func (t *target) Inst(p *gas.Parser, mnem string) error {
	at := p.Peek()
	name, lead := splitMnemonic(mnem)

	ops := lead
	for !p.AtEnd() {
		o, err := t.parseOperand(p)
		if err != nil {
			return err
		}
		ops = append(ops, o)
		if !p.AcceptPunct(",") {
			break
		}
	}

	// The module records failures rather than returning them, so a
	// diagnostic is a difference: an error that was not there before this
	// instruction belongs to it, and one that was is somebody else's,
	// already reported at its own position.
	before := t.e.m.Err()
	t.e.sec.Emit(name, ops...)
	if after := t.e.m.Err(); after != nil && before == nil {
		return p.Errorf(at, "%s: %v", mnem, after)
	}
	return nil
}

// splitMnemonic peels a condition suffix off a mnemonic.
//
// `b.eq` is the source spelling of the table's `b.cond` with an EQ operand, and
// it is the only place on this architecture where part of an operand is
// written inside the mnemonic. Handling it here keeps the table honest: there
// is one conditional branch encoding, and the table has one row for it.
func splitMnemonic(mnem string) (string, []any) {
	lower := strings.ToLower(mnem)
	if rest, ok := strings.CutPrefix(lower, "b."); ok {
		if c, ok := operand.LookupCond(rest); ok {
			return "b.cond", []any{c}
		}
	}
	return lower, nil
}

// Options configure an assembly.
type Options struct {
	// File names the source, for diagnostics.
	File string

	// Features gates which instructions exist. The zero value is the
	// parent package's default set.
	Features arm64.FeatureSet
}

// Assemble assembles src into a finished object.
func Assemble(src string, opts Options) (*obj.Object, error) {
	m, err := AssembleInto(nil, src, opts)
	if err != nil {
		return nil, err
	}
	return m.Finalize()
}

// AssembleInto assembles src into m, creating a module when m is nil.
//
// It is the entry point for assembling more than one text into one object —
// several `.s` files, or a lowered function and the module-level asm beside
// it. The module is returned rather than finalized so the caller decides when
// there is nothing more to add.
func AssembleInto(m *arm64.Module, src string, opts Options) (*arm64.Module, error) {
	if m == nil {
		if opts.Features != (arm64.FeatureSet{}) {
			m = arm64.NewModule(arm64.WithFeatures(opts.Features))
		} else {
			m = arm64.NewModule()
		}
	}
	e := newEmitter(m)
	if err := gas.Parse(opts.File, src, &target{e: e}, e); err != nil {
		return m, err
	}
	e.declareExterns()
	return m, m.Err()
}
