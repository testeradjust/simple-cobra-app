package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer(1, "s1")
	require.NotNil(t, server)
	assert.Equal(t, 1, server.Version)
	assert.Equal(t, "s1", server.Name)
}