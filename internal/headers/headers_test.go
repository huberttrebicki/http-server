package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeadersParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\nFooFoo: barbar\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	host, _ := headers.Get("Host")
	foo, _ := headers.Get("FooFoo")
	missingKey, _ := headers.Get("MissingKey")
	assert.Equal(t, "localhost:42069", host)
	assert.Equal(t, "barbar", foo)
	assert.Equal(t, "", missingKey)
	assert.Equal(t, 41, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nHost: localhost:42069\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069,localhost:42069", host)
	assert.False(t, done)
}
