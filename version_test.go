package tika

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	c := NewTestClient()
	r, _ := c.Version()

	assert.Equal(t,"Apache Tika 1.16", r)
}
