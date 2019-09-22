package iteration

import "testing"

func TestRepeatWithTimes(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected '%q' but got '%q'", expected, repeated)
	}
}

func TestRepeatWithZeroTime(t *testing.T) {
	repeated := Repeat("a", 0)
	expected := ""

	if repeated != expected {
		t.Errorf("Expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
