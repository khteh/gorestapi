package fibonacci
import "testing"
func TestFibonacci(t *testing.T) {
	cases := []struct {
		in uint32;
		expect uint64;
	} {
		{0, 1},
		{1, 1},
		{10, 55},
		{20, 6765},
	}
	for _, c := range cases {
		got := Fibonacci(c.in)
		if got != c.expect {
			t.Errorf("Fibonacci(%q) == %q, expect %q", c.in, got, c.expect)
		}
	}
}
