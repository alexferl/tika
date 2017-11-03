package tika

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguageEn(t *testing.T) {
	c := NewTestClient()
	c.Document = bytes.NewReader([]byte("this is english"))
	r, err := c.Language().Text()

	assert.NoError(t, err)
	assert.Equal(t,"en", r)
}

func TestLanguageFr(t *testing.T) {
	c := NewTestClient()
	c.Document = bytes.NewReader([]byte("comme ci comme ça"))
	r, err := c.Language().Text()

	assert.NoError(t, err)
	assert.Equal(t,"fr", r)
}

func TestLanguageErrNoDocument(t *testing.T) {
	c := NewTestClient()
	_, err := c.Language().Text()

	assert.Error(t, err)
}
