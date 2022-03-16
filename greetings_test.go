package greetings

import "testing"

func TestGreeting(t *testing.T) {
	t.Run("veryfying correct greeting value", func(t *testing.T) {
		gs := Greet()
		expected := "Hello, Go!"

		if gs != expected {
			t.Errorf("Expected %v; but got %v", expected, gs)
		}
	})
}

func TestHelloWithName(t *testing.T) {
	t.Run("veryfying correct hello value", func(t *testing.T) {
		hello := Hello("Marek")
		expected := "Hello Marek!"

		if hello != expected {
			t.Errorf("Expected %v; but got %v", expected, hello)
		}
	})
}

func TestHelloEmptyName(t *testing.T) {
	t.Run("veryfying correct hello value", func(t *testing.T) {
		hello := Hello("")
		expected := "Hello !"

		if hello != expected {
			t.Errorf("Expected %v; but got %v", expected, hello)
		}
	})
}
