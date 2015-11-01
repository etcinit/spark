package logical

import (
	"testing"

	"github.com/etcinit/spark/constraints/types"
	"github.com/stretchr/testify/assert"
)

func TestMaybeConstraintIsFulfilled(t *testing.T) {
	current := MaybeConstraint{
		&types.StringConstraint{},
	}

	assert.True(t, current.IsFulfilled("hello"))
	assert.False(t, current.IsFulfilled(45))
	assert.True(t, current.IsFulfilled(nil))
}
