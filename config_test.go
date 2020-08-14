package llbl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	var err error
	resolver, err = ioutil.TempDir("", "resolver")
	assert.NoError(t, err)

	restore, err := Configure(1234)
	assert.NoError(t, err)

	path := filepath.Join(resolver, "localhost")
	assert.FileExists(t, path)

	b, err := ioutil.ReadFile(path)
	assert.NoError(t, err)
	assert.Equal(t, "nameserver 127.0.0.1\nport 1234", string(b))

	assert.NoError(t, restore())

	_, err = os.Stat(path)
	assert.True(t, os.IsNotExist(err))
}
