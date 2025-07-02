package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"

	if got := Hello(); got != expected {
		t.Errorf("Hello() = %v, want %v", got, expected)
	}
}

func TestHelloLength(t *testing.T) {
	result := Hello()

	if len(result) == 0 {
		t.Error("Hello() returned an empty string")
	}
}
