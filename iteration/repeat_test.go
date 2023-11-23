package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("b", 3)
	expected := "bbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// benchmark test
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

//To run the benchmarks do go test -bench=.
