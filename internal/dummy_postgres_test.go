package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDummyPostgres(t *testing.T) {
	dp, err := NewDummyPostgres("localhost", "dummydb", "dummydb", 5432)
	require.NoError(t, err)
	assert.NotNil(t, dp)
}

func TestDummyPostgres_TestConnection(t *testing.T) {
	dp, err := NewDummyPostgres("localhost", "dummydb", "dummydb", 5432)
	require.NoError(t, err)
	assert.NotNil(t, dp)

	err = dp.TestConnection()
	assert.NoError(t, err)
}
