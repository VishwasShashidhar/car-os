package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	actual := SayHello()
	assert.Equal(t, "Hello, cars!", actual)
}

func TestGetAppName(t *testing.T) {
	actual := GetAppName()
	assert.Equal(t, "CarOS", actual)
}
