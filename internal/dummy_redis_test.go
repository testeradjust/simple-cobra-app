package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDummyRedis(t *testing.T) {
	dummyRedis, err := NewDummyRedis("localhost", 6379)
	assert.NoError(t, err)
	assert.NotNil(t, dummyRedis)
}

func TestDummyRedis_TestConnection(t *testing.T) {
	dummyRedis, err := NewDummyRedis("localhost", 6379)
	assert.NoError(t, err)
	require.NotNil(t, dummyRedis)
	assert.NoError(t, dummyRedis.TestConnection())
}
