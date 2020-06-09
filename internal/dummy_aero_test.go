package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDummyAero(t *testing.T) {
	dummyAero, err := NewDummyAero("172.28.128.3", 3000)
	assert.NoError(t, err)
	assert.NotNil(t, dummyAero, "aero client returned should not be nil")
}

func TestDummyAero_TestConnection(t *testing.T) {
	dummyAero, err := NewDummyAero("172.28.128.3", 3000)
	assert.NoError(t, err)
	require.NotNil(t, dummyAero, "aero client returned should not be nil")

	err = dummyAero.TestConnection()
	assert.Nil(t, err)
}
