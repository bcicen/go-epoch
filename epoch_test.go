package epoch

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func runTest(t *testing.T, n int64, tt time.Time, p time.Duration) {
	et := Time(n).UTC()
	margs := []any{"n=%d et=%s t=%s", n, et, tt}
	assert.Equal(t, tt, et, margs...)
	assert.Equal(t, tt.UnixMilli(), et.UnixMilli(), margs...)
}

func TestNewInputSeconds(t *testing.T) {
	tt := mustParseTime("1970-04-18T12:34:56.000000000Z")
	for i := 0; i < 2000; i++ {
		runTest(t, tt.Unix(), tt, time.Second)
		tt = tt.Add(32 * 24 * time.Hour)
	}
}

func TestNewInputMillis(t *testing.T) {
	tt := mustParseTime("1970-04-18T12:34:56.789000000Z")
	for i := 0; i < 2000; i++ {
		runTest(t, tt.UnixMilli(), tt, time.Millisecond)
		tt = tt.Add(32 * 24 * time.Hour)
	}
}

func TestNewInputNanos(t *testing.T) {
	tt := mustParseTime("1970-04-18T12:34:56.789123456Z")
	for i := 0; i < 2000; i++ {
		runTest(t, tt.UnixNano(), tt, time.Nanosecond)
		tt = tt.Add(32 * 24 * time.Hour)
	}
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return t
}
