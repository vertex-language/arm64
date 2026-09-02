package arm64

import "github.com/vertex-language/arm64/reg"

// ---- Data processing, one source ------------------------------------------

var (
	rbit32    = form("Rbit32")
	rbit64    = form("Rbit64")
	rev16_32  = form("Rev16_32")
	rev16_64  = form("Rev16_64")
	rev32     = form("Rev32")
	rev64     = form("Rev64")
	rev64Alt  = form("Rev64Alias")
	rev32In64 = form("Rev32In64")
	clz32     = form("Clz32")
	clz64     = form("Clz64")
	cls32     = form("Cls32")
	cls64     = form("Cls64")
)

// Rbit32 emits RBIT Wd, Wn: Wd's bits in reverse order.
func (s *Section) Rbit32(rd, rn reg.W) { s.inst(rbit32, rd, rn) }
func (s *Section) Rbit64(rd, rn reg.X) { s.inst(rbit64, rd, rn) }

// Rev16_32 emits REV16 Wd, Wn: the bytes of each halfword of Wn reversed.
func (s *Section) Rev16_32(rd, rn reg.W) { s.inst(rev16_32, rd, rn) }
func (s *Section) Rev16_64(rd, rn reg.X) { s.inst(rev16_64, rd, rn) }

// Rev32 emits REV Wd, Wn: the bytes of Wn reversed.
func (s *Section) Rev32(rd, rn reg.W) { s.inst(rev32, rd, rn) }

// Rev64 emits REV Xd, Xn: the bytes of Xn reversed.
func (s *Section) Rev64(rd, rn reg.X) { s.inst(rev64, rd, rn) }

// Rev64Alias is the same word under the name the ARM ARM prefers when the
// register is an X. Two spellings, one instruction, and assembly written by
// hand uses both.
func (s *Section) Rev64Alias(rd, rn reg.X) { s.inst(rev64Alt, rd, rn) }

// Rev32In64 emits REV32 Xd, Xn: the bytes within each word of Xn reversed,
// which is a different instruction from REV and not a spelling of it.
func (s *Section) Rev32In64(rd, rn reg.X) { s.inst(rev32In64, rd, rn) }

// Clz32 emits CLZ Wd, Wn: the count of leading zero bits.
func (s *Section) Clz32(rd, rn reg.W) { s.inst(clz32, rd, rn) }
func (s *Section) Clz64(rd, rn reg.X) { s.inst(clz64, rd, rn) }

// Cls32 emits CLS Wd, Wn: the count of leading bits that match the sign bit,
// excluding the sign bit itself.
func (s *Section) Cls32(rd, rn reg.W) { s.inst(cls32, rd, rn) }
func (s *Section) Cls64(rd, rn reg.X) { s.inst(cls64, rd, rn) }
