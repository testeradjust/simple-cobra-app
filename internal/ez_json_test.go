package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEzJsonPerson(t *testing.T) {
	p1 := EzJsonPerson{
		Age:  100,
		Name: "P1",
	}
	p1b, err := p1.MarshalJSON()
	require.NoError(t, err)
	assert.NotEmpty(t, p1b)

	var p1ret EzJsonPerson
	err = json.Unmarshal(p1b, &p1ret)
	require.NoError(t, err)
	require.NotNil(t, p1ret)
	assert.Equal(t, 100, p1ret.Age)
	assert.Equal(t, "P1", p1ret.Name)
}

func TestEzJsonCity(t *testing.T) {
	c1 := NewEzJsonCity("city1", 2500)
	assert.NotNil(t, c1)

	p1 := EzJsonPerson{
		Age:  100,
		Name: "P1",
	}
	c1.People = append(c1.People, p1)

	c1b, err := c1.MarshalJSON()
	require.NoError(t, err)
	assert.NotEmpty(t, c1b)

	var c1ret EzJsonCity
	err = json.Unmarshal(c1b, &c1ret)
	require.NoError(t, err)
	require.NotNil(t, c1ret)
	assert.Equal(t, 2500, c1ret.Population)
	assert.Equal(t, "city1", c1ret.Name)
	assert.Len(t, c1.People, 1)
}
