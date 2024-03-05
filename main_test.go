package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T){
	result := add(2,3)
	assert.Equal(t, result,5)
		
}
