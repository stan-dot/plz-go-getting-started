package greetings_test

import (
	// importing the "production" package
	"github.com/example/module/src/greetings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreeting(t *testing.T) {
	if greetings.Greeting() == "" {
		panic("Greeting failed to produce a result")
	}

	assert.NotEqual(t, "", greetings.Greeting())
}
