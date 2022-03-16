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
