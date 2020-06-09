package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	aerospikeHostname = "localhost"
	aerospikePort     = 3000
)

func TestNewDummyAero(t *testing.T) {
	dummyAero, err := NewDummyAero(aerospikeHostname, aerospikePort)
	assert.NoError(t, err)
	assert.NotNil(t, dummyAero, "aero client returned should not be nil")
}

func TestDummyAero_TestConnection(t *testing.T) {
	dummyAero, err := NewDummyAero(aerospikeHostname, aerospikePort)
	assert.NoError(t, err)
	require.NotNil(t, dummyAero, "aero client returned should not be nil")

	err = dummyAero.TestConnection()
	assert.Nil(t, err)
}
