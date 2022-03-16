package greetings

import "fmt"

func Greet() string {
	return "Hello, Go!"
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
