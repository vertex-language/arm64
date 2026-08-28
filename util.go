package arm64

// The small shared helpers. They are here rather than beside their callers
// because each has more than one, and a function copied into two files is two
// functions that can drift.

// hex renders an offset the way a diagnostic prints one, with no 0x prefix —
// the caller supplies that, because "at .text+0x11" and "0x11 bytes" want
// different framing around the same digits.
func hex(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	u := uint64(n)
	if neg {
		u = uint64(-n)
	}
	const digits = "0123456789abcdef"
	var b [17]byte
	i := len(b)
	for u > 0 {
		i--
		b[i] = digits[u&0xf]
		u >>= 4
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

// decimal is strconv.FormatInt without the import.
func decimal(v int64) string {
	if v == 0 {
		return "0"
	}
	neg := v < 0
	u := uint64(v)
	if neg {
		u = uint64(-v)
	}
	var b [21]byte
	i := len(b)
	for u > 0 {
		i--
		b[i] = byte('0' + u%10)
		u /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

// itoa is decimal for the counts and widths a diagnostic quotes.
func itoa(n int) string { return decimal(int64(n)) }
