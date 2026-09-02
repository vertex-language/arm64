package operand

import "testing"

// TestSentinelsCountEverything pins the two enumeration sentinels.
//
// Both const blocks state explicit values rather than counting with iota, and
// in such a block an omitted expression repeats the previous one instead of
// continuing. Written bare, shiftCount was 3 and extendCount was 7 — one short
// each — so ROR.Valid and SXTX.Valid reported false and neither could be
// encoded through any path, typed or textual. Nothing else noticed: an
// operand that reports itself invalid is refused, and a refusal reads like an
// instruction the table does not carry.
func TestSentinelsCountEverything(t *testing.T) {
	if shiftCount != ROR+1 {
		t.Errorf("shiftCount is %d, want %d: the last shift is %v", shiftCount, ROR+1, ROR)
	}
	if extendCount != SXTX+1 {
		t.Errorf("extendCount is %d, want %d: the last extend is %v", extendCount, SXTX+1, SXTX)
	}
	for _, s := range []Shift{LSL, LSR, ASR, ROR} {
		if !s.Valid() {
			t.Errorf("%v reports itself invalid", s)
		}
	}
	for _, e := range []Extend{UXTB, UXTH, UXTW, UXTX, SXTB, SXTH, SXTW, SXTX} {
		if !e.Valid() {
			t.Errorf("%v reports itself invalid", e)
		}
	}
	if Shift(shiftCount).Valid() {
		t.Error("the sentinel itself reports valid")
	}
	if Extend(extendCount).Valid() {
		t.Error("the sentinel itself reports valid")
	}
}
