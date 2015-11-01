package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringLike string

func TestStringConstraintIsFulfilled(t *testing.T) {
	constraint := StringConstraint{}

	var like StringLike = "hello again"

	assert.True(t, constraint.IsFulfilled("hello world"))
	assert.False(t, constraint.IsFulfilled(like))
	assert.True(t, constraint.IsFulfilled(string(like)))
	assert.False(t, constraint.IsFulfilled(nil))
}
