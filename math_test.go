package main

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {

	res := Soma(15, 15)
	// assert.Equal(t, 10, soma1, "Shoud be the same.")

	if res != 30 {
		t.Error("Invalid Result")
	}
}
