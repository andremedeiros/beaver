package store

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*Store, func()) {
	t.Helper()

	root, err := ioutil.TempDir("", "beaver")
	assert.NoError(t, err)

	s, err := NewStore(root)
	assert.NoError(t, err)

	teardown := func() {
		s.Close()
		os.RemoveAll(root)
	}

	return s, teardown
}

func TestNewStore(t *testing.T) {
	root, err := ioutil.TempDir("", "beaver")
	assert.NoError(t, err)

	_, err = NewStore(root)
	assert.NoError(t, err)
}
