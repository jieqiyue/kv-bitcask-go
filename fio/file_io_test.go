package fio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewFileIOManager(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/tmp", "a.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/tmp", "a.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)
	t.Log(n)

	n, err = fio.Write([]byte("hello"))
	assert.Equal(t, 5, n)
	assert.Nil(t, err)
	t.Log(n)

	n, err = fio.Write([]byte("world"))
	assert.Equal(t, 5, n)
	assert.Nil(t, err)
	t.Log(n)

	_, _ = fio.Write([]byte("bitcask kv\n"))
}
