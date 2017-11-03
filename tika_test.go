package tika

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTikaHello(t *testing.T) {
	c := NewTestClient()
	r, err := c.Hello()

	assert.NoError(t, err)
	assert.Equal(t,"This is Tika Server (Apache Tika 1.16). Please PUT\n", r)
}

func TestTikaHtml(t *testing.T) {
	c := NewTestClient()
	c.Document = bytes.NewReader([]byte("this is english"))
	r, err := c.Tika().Html()

	assert.NoError(t, err)
	assert.Equal(t,"<html xmlns=\"http://www.w3.org/1999/xhtml\">\n<head>\n<meta name=\"X-Parsed-By\" " +
		"content=\"org.apache.tika.parser.DefaultParser\"/>\n<meta name=\"X-Parsed-By\" " +
		"content=\"org.apache.tika.parser.txt.TXTParser\"/>\n<meta name=\"Content-Encoding\" " +
		"content=\"ISO-8859-1\"/>\n<meta name=\"Content-Type\" " +
		"content=\"text/plain; " +
		"charset=ISO-8859-1\"/>\n<title>\n</title>\n</head>\n<body>\n<p>this is english</p>\n</body>\n</html>\n", r)
}

func TestTikaText(t *testing.T) {
	c := NewTestClient()
	c.Document = bytes.NewReader([]byte("this is english"))
	r, err := c.Tika().Text()

	assert.NoError(t, err)
	assert.Equal(t, "this is english\n", r)
}

func TestTikaErrNoDocument(t *testing.T) {
	c := NewTestClient()
	c.Document = nil
	_, err := c.Tika().Text()

	assert.Error(t, err)
}
