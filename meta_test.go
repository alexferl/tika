package tika

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	c := NewTestClient()
	c.Document = bytes.NewReader([]byte("this is english"))
	r, err := c.Meta().Csv()

	assert.NoError(t, err)
	assert.Equal(t, "\"X-Parsed-By\"," +
		"\"org.apache.tika.parser.DefaultParser\"," +
		"\"org.apache.tika.parser.txt.TXTParser\"\n\"Content-Encoding\"," +
		"\"ISO-8859-1\"\n\"language\"," +
		"\"en\"\n\"Content-Type\"," +
		"\"text/plain; charset=ISO-8859-1\"\n", r)
}