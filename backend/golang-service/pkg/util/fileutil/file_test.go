package fileutil

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRemoveContents(t *testing.T) {
	const filename = "tmp/abc/123/def/hello.txt"
	err := CreateRecursively(filename)
	assert.NoError(t, err)

	err = RemoveContents("tmp/abc/123/def")
	assert.NoError(t, err)

	exists, err := IsFileExists(filename)
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestRemoveDir(t *testing.T) {
	err := os.RemoveAll("tmp")
	assert.NoError(t, err)
}
