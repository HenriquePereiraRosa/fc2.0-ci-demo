package main

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {

	soma := Soma(5, 5)
	// assert.Equal(t, 10, soma1, "Shoud be the same.")

	if soma != 30 {
		t.Error("Inavlid Result")
	}
}
