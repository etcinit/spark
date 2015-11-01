package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyIntConstraintIsFulfilled(t *testing.T) {
	constraint := AnyIntConstraint{}

	assert.True(t, constraint.IsFulfilled(12))
	assert.True(t, constraint.IsFulfilled(int(12)))
	assert.True(t, constraint.IsFulfilled(int8(12)))
	assert.True(t, constraint.IsFulfilled(int16(12)))
	assert.True(t, constraint.IsFulfilled(int32(12)))
	assert.True(t, constraint.IsFulfilled(int64(12)))
	assert.False(t, constraint.IsFulfilled(13.37))
	assert.False(t, constraint.IsFulfilled(float32(13.37)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestIntConstraintIsFulfilled(t *testing.T) {
	constraint := IntConstraint{}

	assert.True(t, constraint.IsFulfilled(12))
	assert.True(t, constraint.IsFulfilled(int(12)))
	assert.False(t, constraint.IsFulfilled(int8(12)))
	assert.False(t, constraint.IsFulfilled(int16(12)))
	assert.False(t, constraint.IsFulfilled(int32(12)))
	assert.False(t, constraint.IsFulfilled(int64(12)))
	assert.False(t, constraint.IsFulfilled(13.37))
	assert.False(t, constraint.IsFulfilled(float32(13.37)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestInt8ConstraintIsFulfilled(t *testing.T) {
	constraint := Int8Constraint{}

	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(int(12)))
	assert.True(t, constraint.IsFulfilled(int8(12)))
	assert.False(t, constraint.IsFulfilled(int16(12)))
	assert.False(t, constraint.IsFulfilled(int32(12)))
	assert.False(t, constraint.IsFulfilled(int64(12)))
	assert.False(t, constraint.IsFulfilled(13.37))
	assert.False(t, constraint.IsFulfilled(float32(13.37)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestInt16ConstraintIsFulfilled(t *testing.T) {
	constraint := Int16Constraint{}

	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(int(12)))
	assert.False(t, constraint.IsFulfilled(int8(12)))
	assert.True(t, constraint.IsFulfilled(int16(12)))
	assert.False(t, constraint.IsFulfilled(int32(12)))
	assert.False(t, constraint.IsFulfilled(int64(12)))
	assert.False(t, constraint.IsFulfilled(13.37))
	assert.False(t, constraint.IsFulfilled(float32(13.37)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestInt32ConstraintIsFulfilled(t *testing.T) {
	constraint := Int32Constraint{}

	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(int(12)))
	assert.False(t, constraint.IsFulfilled(int8(12)))
	assert.False(t, constraint.IsFulfilled(int16(12)))
	assert.True(t, constraint.IsFulfilled(int32(12)))
	assert.False(t, constraint.IsFulfilled(int64(12)))
	assert.False(t, constraint.IsFulfilled(13.37))
	assert.False(t, constraint.IsFulfilled(float32(13.37)))
	assert.False(t, constraint.IsFulfilled(nil))
}

func TestInt64ConstraintIsFulfilled(t *testing.T) {
	constraint := Int64Constraint{}

	assert.False(t, constraint.IsFulfilled(12))
	assert.False(t, constraint.IsFulfilled(int(12)))
	assert.False(t, constraint.IsFulfilled(int8(12)))
	assert.False(t, constraint.IsFulfilled(int16(12)))
	assert.False(t, constraint.IsFulfilled(int32(12)))
	assert.True(t, constraint.IsFulfilled(int64(12)))
	assert.False(t, constraint.IsFulfilled(13.37))
	assert.False(t, constraint.IsFulfilled(float32(13.37)))
	assert.False(t, constraint.IsFulfilled(nil))
}
