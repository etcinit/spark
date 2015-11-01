package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilConstraintIsFulfilled(t *testing.T) {
	constraint := NilConstraint{}

	assert.True(t, constraint.IsFulfilled(nil))
	assert.False(t, constraint.IsFulfilled(""))
	assert.False(t, constraint.IsFulfilled("hello"))
}
