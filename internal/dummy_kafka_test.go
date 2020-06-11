package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDummyKafka(t *testing.T) {
	brokers := []string{"localhost:9092"}
	dummyKafka, err := NewDummyKafka(brokers)
	assert.NoError(t, err)
	assert.NotNil(t, dummyKafka)
}

func TestDummyKafka_ProduceTestMessage(t *testing.T) {
	brokers := []string{"localhost:9092"}
	dummyKafka, err := NewDummyKafka(brokers)
	assert.NoError(t, err)
	require.NotNil(t, dummyKafka)

	assert.NotPanics(t, func() {
		dummyKafka.ProduceTestMessage("test message")
	})
}
