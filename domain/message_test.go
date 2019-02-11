package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage(t *testing.T) {
	message := Message("this is a test")
	assert.Equal(t, M{Message: "this is a test"}, message)
}
