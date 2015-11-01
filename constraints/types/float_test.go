package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Float32Like float32
type Float64Like float64

func TestAnyFloatConstraintIsFulfilled(t *testing.T) {
	constraint := AnyFloatConstraint{}

	var like Float32Like = 13.37
	var like2 Float64Like = 13.37

	assert.True(t, constraint.IsFulfilled(12.45))
	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(like))
	assert.False(t, constraint.IsFulfilled(like2))
	assert.True(t, constraint.IsFulfilled(float32(like)))
	assert.True(t, constraint.IsFulfilled(float64(like2)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestFloat32ConstraintIsFulfilled(t *testing.T) {
	constraint := Float32Constraint{}

	var like Float32Like = 13.37
	var like2 Float64Like = 13.37

	assert.True(t, constraint.IsFulfilled(float32(12.45)))
	assert.False(t, constraint.IsFulfilled(12.45))
	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(like))
	assert.False(t, constraint.IsFulfilled(like2))
	assert.True(t, constraint.IsFulfilled(float32(like)))
	assert.False(t, constraint.IsFulfilled(float64(like2)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestFloat64ConstraintIsFulfilled(t *testing.T) {
	constraint := Float64Constraint{}

	var like Float32Like = 13.37
	var like2 Float64Like = 13.37

	assert.True(t, constraint.IsFulfilled(float64(12.45)))
	assert.True(t, constraint.IsFulfilled(12.45))
	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(like))
	assert.False(t, constraint.IsFulfilled(like2))
	assert.False(t, constraint.IsFulfilled(float32(like)))
	assert.True(t, constraint.IsFulfilled(float64(like2)))
	assert.False(t, constraint.IsFulfilled(nil))
}
