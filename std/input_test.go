package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetInput(t *testing.T) {
	testInput := strings.NewReader("hello world")
	result, err := capitalize(testInput)
	if err != nil {
		t.Fatal(err)
	}
	expectedResult := "HELLO WORLD"
	assert.Equal(t, expectedResult, result, "expected but got different")
	t.Log(result, err)
}
