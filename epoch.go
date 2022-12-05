package epoch

import (
	"math"
	"time"
)

const (
	msNanos = 1000000 // 1 ms in ns
	sMillis = 1000    // 1 s in ms

	maxNanos   = math.MaxInt64
	maxMillis  = maxNanos / msNanos
	maxSeconds = maxMillis / sMillis
)

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64
}

// Time interprets a provided s/ms/ns unix timestamp,
// returning a time.Time
// if n <= 0, returns zero time
func Time[T Numeric](ts T) time.Time {
	n := int64(ts)
	if n <= 0 {
		return time.Time{}
	}
	switch {
	case n > maxMillis:
		// input already ns
		return time.Unix(0, n)
	case n > maxSeconds:
		// input in ms
		ns := int64(time.Duration(n) * time.Millisecond)
		return time.Unix(0, ns)
	default:
		// input in s
		return time.Unix(n, 0)
	}
}
