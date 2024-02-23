package greetings

import (
	"math/rand"
)

var greetings = []string{
	"hello",
	"bonjour",
	"Marhaaban",
}

func Greeting()string {
	return greetings[rand.Intn(len(greetings))]
}	
