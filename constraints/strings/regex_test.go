package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexConstraintIsFulfilled(t *testing.T) {
	constraint := RegexConstraint{
		Regex: "[a-zA-Z0-9]",
	}

	assert.True(t, constraint.IsFulfilled("01234"))
	assert.True(t, constraint.IsFulfilled("abcDEF01234"))
	assert.False(t, constraint.IsFulfilled("@#$%"))
	assert.False(t, constraint.IsFulfilled(nil))
}
