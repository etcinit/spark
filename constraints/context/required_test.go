package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredConstraintIsFulfilled(t *testing.T) {
	constraint := RequiredConstraint{}

	assert.True(t, constraint.IsFulfilled("hello"))
}

func TestRequiredConstraintIsFulfilledByContext(t *testing.T) {
	constraint := RequiredConstraint{Fields: []string{"hello", "world"}}

	input := map[string]interface{}{
		"hello": "how are you?",
		"world": "its very small",
	}
	input2 := map[string]interface{}{
		"hello": "how are you?",
	}

	assert.True(t, constraint.IsFulfilledByContext(nil, input))
	assert.False(t, constraint.IsFulfilledByContext(nil, input2))
}

func TestRequiredConstraintDescribeContextFailure(t *testing.T) {
	constraint := RequiredConstraint{Fields: []string{"hello", "world"}}

	input := map[string]interface{}{
		"hello": "how are you?",
		"world": "its very small",
	}
	input2 := map[string]interface{}{
		"hello": "how are you?",
	}
	input3 := map[string]interface{}{}

	assert.Equal(
		t,
		"The following fields are missing from the input: ",
		constraint.DescribeContextFailure(input),
	)
	assert.Equal(
		t,
		"The following fields are missing from the input: world",
		constraint.DescribeContextFailure(input2),
	)
	assert.Equal(
		t,
		"The following fields are missing from the input: hello, world",
		constraint.DescribeContextFailure(input3),
	)
}
