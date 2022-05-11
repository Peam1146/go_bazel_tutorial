package greet

import "testing"

func TestGreet(t *testing.T) {
	t.Run("Test Greet", func(t *testing.T) {
		if got := Greet(); got != "Hello World" {
			t.Errorf("Greet() = %q, want %q", got, "Hello World")
		}
	})
}
