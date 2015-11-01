package logical

import (
	"testing"

	"github.com/etcinit/spark/constraints/types"
	"github.com/stretchr/testify/assert"
)

func TestEitherConstraintIsFulfilled(t *testing.T) {
	current := EitherConstraint{
		&types.StringConstraint{},
		&types.AnyIntConstraint{},
	}

	assert.True(t, current.IsFulfilled("hello"))
	assert.True(t, current.IsFulfilled(45))
	assert.False(t, current.IsFulfilled(nil))
}
