package greetings_test

import (
	// importing the "production" package
	"github.com/example/module/src/greetings"
	"testing"
)

func TestGreeting(t *testing.T) {
	if greetings.Greeting() == "" {
		panic("Greeting failed to produce a result")
	}
}
