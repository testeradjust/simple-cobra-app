package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDummyHomeDirFunc(t *testing.T) {
	str, err := DummyHomeDirFunc()
	require.Nil(t, err)
	assert.True(t, len(str) > 0)
}

func TestDummyViperFunc(t *testing.T) {
	v := DummyViperFunc()
	assert.NotNil(t, v)
}
