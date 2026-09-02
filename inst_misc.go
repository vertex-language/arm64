package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Exceptions ---------------------------------------------------------------

var (
	svc = form("Svc")
	hvc = form("Hvc")
	smc = form("Smc")
	brk = form("Brk")
	hlt = form("Hlt")
)

// Svc emits SVC #imm: a supervisor call.
func (s *Section) Svc(imm uint16) { s.inst(svc, uint64(imm)) }

// Hvc emits HVC #imm: a hypervisor call.
func (s *Section) Hvc(imm uint16) { s.inst(hvc, uint64(imm)) }

// Smc emits SMC #imm: a secure monitor call.
func (s *Section) Smc(imm uint16) { s.inst(smc, uint64(imm)) }

// Brk emits BRK #imm: a breakpoint, caught by a debugger rather than the OS.
func (s *Section) Brk(imm uint16) { s.inst(brk, uint64(imm)) }

// Hlt emits HLT #imm: a halt, caught by external debug hardware.
func (s *Section) Hlt(imm uint16) { s.inst(hlt, uint64(imm)) }

// ---- Hints and no-ops -----------------------------------------------------

var (
	hint    = form("Hint")
	paciasp = form("Paciasp")
	autiasp = form("Autiasp")

	nop   = form("Nop")
	yield = form("Yield")
	wfe   = form("Wfe")
	wfi   = form("Wfi")
	sev   = form("Sev")
	sevl  = form("Sevl")
)

// Nop emits NOP: D503201F, the architecture's one canonical no-op.
func (s *Section) Nop() { s.inst(nop) }

// Hint emits HINT #imm, the hint space itself. Every named hint below is one
// of these, and so is every hint a processor does not implement — which is
// the point of the encoding: an unimplemented hint is a NOP rather than an
// undefined instruction, so a barrier written for a later processor runs on
// an earlier one.
func (s *Section) Hint(imm uint8) { s.inst(hint, uint64(imm)) }

// Paciasp signs the return address in X30 against SP; Autiasp authenticates
// it. They are hints 25 and 29, which is why they are NOPs on a processor
// without pointer authentication and why a function using them still runs
// there.
func (s *Section) Paciasp() { s.inst(paciasp) }
func (s *Section) Autiasp() { s.inst(autiasp) }

func (s *Section) Yield() { s.inst(yield) }
func (s *Section) Wfe()   { s.inst(wfe) }
func (s *Section) Wfi()   { s.inst(wfi) }
func (s *Section) Sev()   { s.inst(sev) }
func (s *Section) Sevl()  { s.inst(sevl) }

// ---- Barriers -----------------------------------------------------------------

var (
	dsb = form("Dsb")
	dmb = form("Dmb")
	isb = form("Isb")
)

// Dsb emits DSB option, or DSB SY with no operand — SY is the field's own
// default, the same reasoning RET's omitted operand defaults to X30.
func (s *Section) Dsb(option ...Barrier) { s.inst(dsb, opt(option)...) }
func (s *Section) Dmb(option ...Barrier) { s.inst(dmb, opt(option)...) }
func (s *Section) Isb(option ...Barrier) { s.inst(isb, opt(option)...) }

// ---- System register move ------------------------------------------------------

var (
	mrs    = form("Mrs")
	msrReg = form("MsrReg")
)

// Mrs emits MRS Xt, sysreg: reads a system register accessible to the
// current exception level into Xt.
func (s *Section) Mrs(rt reg.X, sr Sys) { s.inst(mrs, rt, sr) }

// MsrReg emits MSR sysreg, Xt: writes Xt into a system register.
func (s *Section) MsrReg(sr Sys, rt reg.X) { s.inst(msrReg, sr, rt) }

// ---- MSR (immediate) --------------------------------------------------------

var msrImm = form("MsrImm")

// MsrImm emits MSR <pstatefield>, #imm: four bits written into a named piece
// of process state. `MsrImm(operand.DAIFSet, 2)` is how interrupts are
// masked, and it is a different instruction from MSR with a register — there
// is no register anywhere in the word.
func (s *Section) MsrImm(field PState, imm uint8) { s.inst(msrImm, field, uint64(imm)) }
