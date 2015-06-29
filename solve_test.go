package rpn

import "testing"

func TestSolveBadInputs(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"foo"}, "foo is neither an operator nor an operand."},
		{[]string{"42", "foo", "-"}, "foo is neither an operator nor an operand."},
		{[]string{""}, "empty expression"},
	}
	for _, c := range cases {
		got := Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSolveAddition(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"4", "5", "+"}, "9"},
		{[]string{"9", "3", "-"}, "6"},
		{[]string{"9", "3", "-", "3", "-"}, "3"},
		{[]string{"4", "5", "6", "+", "-"}, "-7"},
		{[]string{"42"}, "42"},
		{[]string{"42", "-"}, "-42"},
	}
	for _, c := range cases {
		got := Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSolveMultiplication(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"4", "5", "*"}, "20"},
		{[]string{"9", "3", "/"}, "3"},
		{[]string{"9", "3", "/", "3", "*"}, "9"},
		{[]string{"12", "4", "6", "*", "/"}, "0.5"},
	}
	for _, c := range cases {
		got := Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
