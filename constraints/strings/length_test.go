package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthConstraintIsFulfilled(t *testing.T) {
	constraint := LengthConstraint{Min: 8, Max: 16}

	assert.True(t, constraint.IsFulfilled("omg-this-is-long"))
	assert.False(t, constraint.IsFulfilled("omg"))
	assert.False(t, constraint.IsFulfilled("omg-this-is-longomg-this-is-long"))
	assert.False(t, constraint.IsFulfilled(45))
	assert.False(t, constraint.IsFulfilled(nil))
}
