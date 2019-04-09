package models

import (
	"testing"

	"github.com/gobuffalo/nulls"
	"github.com/stretchr/testify/assert"
)

func Test_Transaction(t *testing.T) {
	roll := RollReward{
		GroupID:    1,
		Receiver:   1,
		Multiplier: nulls.NewUInt32(0),
	}
	// Dubs
	assert.Equal(t, 20, roll.GetAmount())

	// Trips
	roll.Multiplier = nulls.NewUInt32(1)
	assert.Equal(t, 200, roll.GetAmount())

	// Quads
	roll.Multiplier = nulls.NewUInt32(2)
	assert.Equal(t, 2000, roll.GetAmount())
}
