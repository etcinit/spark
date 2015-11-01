package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BoolLike bool

func TestBoolConstraintIsFulfilled(t *testing.T) {
	constraint := BoolConstraint{}

	var like BoolLike = true

	assert.True(t, constraint.IsFulfilled(true))
	assert.True(t, constraint.IsFulfilled(false))
	assert.False(t, constraint.IsFulfilled(like))
	assert.True(t, constraint.IsFulfilled(bool(like)))
	assert.False(t, constraint.IsFulfilled(nil))
}
