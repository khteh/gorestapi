package greetings

import (
	"regexp"
	"testing"
)

func TestGreetingShouldPass(t *testing.T) {
	name := "Test"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting(name)
	t.Logf(msg)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Test") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
func TestEmptyNameGreetingShouldPass(t *testing.T) {
	name := ""
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Test") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
