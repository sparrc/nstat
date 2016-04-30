package nstat

import (
	"testing"
)

var result map[string]int64

func BenchmarkNstatGet(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{}
	for n := 0; n < b.N; n++ {
		counters = c.Get()
	}
	result = counters
}

func BenchmarkNstatGetDumpZeros(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{
		DumpZeros: true,
	}
	for n := 0; n < b.N; n++ {
		counters = c.Get()
	}
	result = counters
}
