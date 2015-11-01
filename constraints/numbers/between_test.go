package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetweenIntConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenIntConstraint{
		Min: -23,
		Max: 45,
	}

	assert.True(t, constraint.IsFulfilled(0))
	assert.False(t, constraint.IsFulfilled(-100))
	assert.False(t, constraint.IsFulfilled(300))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestBetweenInt8ConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenInt8Constraint{
		Min: -23,
		Max: 45,
	}

	assert.True(t, constraint.IsFulfilled(int8(0)))
	assert.False(t, constraint.IsFulfilled(int8(-100)))
	assert.False(t, constraint.IsFulfilled(int8(47)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestBetweenInt16ConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenInt16Constraint{
		Min: -23,
		Max: 45,
	}

	assert.True(t, constraint.IsFulfilled(int16(0)))
	assert.False(t, constraint.IsFulfilled(int16(-100)))
	assert.False(t, constraint.IsFulfilled(int16(47)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestBetweenInt32ConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenInt32Constraint{
		Min: -23,
		Max: 45,
	}

	assert.True(t, constraint.IsFulfilled(int32(0)))
	assert.False(t, constraint.IsFulfilled(int32(-100)))
	assert.False(t, constraint.IsFulfilled(int32(47)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestBetweenInt64ConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenInt64Constraint{
		Min: -23,
		Max: 45,
	}

	assert.True(t, constraint.IsFulfilled(int64(0)))
	assert.False(t, constraint.IsFulfilled(int64(-100)))
	assert.False(t, constraint.IsFulfilled(int64(47)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestBetweenFloat32ConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenFloat32Constraint{
		Min: -23.32,
		Max: 45.87,
	}

	assert.True(t, constraint.IsFulfilled(float32(0.0)))
	assert.False(t, constraint.IsFulfilled(float32(-100.12)))
	assert.False(t, constraint.IsFulfilled(float32(47.23)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestBetweenFloat64ConstraintIsFulfilled(t *testing.T) {
	constraint := BetweenFloat64Constraint{
		Min: -23.67,
		Max: 45.89,
	}

	assert.True(t, constraint.IsFulfilled(float64(0.0)))
	assert.False(t, constraint.IsFulfilled(float64(-100.45)))
	assert.False(t, constraint.IsFulfilled(float64(47)))
	assert.False(t, constraint.IsFulfilled(nil))
}
