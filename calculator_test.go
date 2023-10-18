package calculator_test

import (
	"testing"

	calculator "github.com/hanan-git/go-cal"
)

type testCase struct {
	a    float64
	b    float64
	want float64
	fn   func(float64, float64) float64
}

func TestAddSubMul(t *testing.T) {
	tcs := []testCase{
		{4, 6, 10.0, calculator.Add},
		{4, 6, -2.0, calculator.Subsctract},
		{4, 6, 24, calculator.Multiply},
	}
	t.Parallel()
	for _, ts := range tcs {
		got := ts.fn(ts.a, ts.b)
		if ts.want != got {
			t.Errorf("want %.1f but got %.1f", ts.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		a        float64
		b        float64
		want     float64
		fn       func(float64, float64) (float64, error)
		errorExp bool
	}{
		{a: 6, b: 3, want: 2.0, fn: calculator.Divide, errorExp: false},
		{a: 0, b: 3, want: 0, fn: calculator.Divide, errorExp: true},
		{a: 3, b: 3, want: 1.0, fn: calculator.Divide, errorExp: false},
	}

	for _, ts := range tcs {
		got, err := ts.fn(ts.a, ts.b)
		errReceived := err != nil
		if errReceived && !ts.errorExp {
			t.Fatalf("Unknown error, check function again")
		}
		if !ts.errorExp && got != ts.want {
			t.Errorf("want %.1f, got %.1f", ts.want, got)
		}
	}
}
