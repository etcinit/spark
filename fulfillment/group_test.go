package fulfillment

import (
	"testing"

	"github.com/etcinit/spark/constraints/context"
	"github.com/etcinit/spark/constraints/types"
	"github.com/stretchr/testify/assert"
)

func TestNewGroupResult(t *testing.T) {
	result := NewGroupResult(
		[]Fulfillable{&types.StringConstraint{}},
		[]Fulfillable{&types.IntConstraint{}, &types.IntConstraint{}},
	)

	assert.Equal(t, 1, len(result.GetPassed()))
	assert.Equal(t, 2, len(result.GetFailed()))
}

func TestNewGroupContextResult(t *testing.T) {
	result := NewGroupContextResult(
		[]ContextFulfillable{&context.RequiredConstraint{}},
		[]ContextFulfillable{
			&context.RequiredConstraint{},
			&context.RequiredConstraint{},
		},
	)

	assert.Equal(t, 1, len(result.GetPassed()))
	assert.Equal(t, 2, len(result.GetFailed()))
}

func TestGroupResultGetPassed(t *testing.T) {
	result := GroupResult{
		passed: []Fulfillable{
			&types.StringConstraint{},
			&types.IntConstraint{},
		},
	}

	assert.Equal(t, 2, len(result.GetPassed()))
}

func TestGroupResultGetFailed(t *testing.T) {
	result := GroupResult{
		failed: []Fulfillable{
			&types.StringConstraint{},
			&types.IntConstraint{},
		},
	}

	assert.Equal(t, 2, len(result.GetFailed()))
}

func TestGroupContextResultGetPassed(t *testing.T) {
	result := GroupContextResult{
		passed: []ContextFulfillable{
			&context.RequiredConstraint{},
		},
	}

	assert.Equal(t, 1, len(result.GetPassed()))
}

func TestGroupContextResultGetFailed(t *testing.T) {
	result := GroupContextResult{
		failed: []ContextFulfillable{
			&context.RequiredConstraint{},
		},
	}

	assert.Equal(t, 1, len(result.GetFailed()))
}

func TestCheckGroup(t *testing.T) {
	result := CheckGroup(
		"omg",
		map[string]interface{}{"something": "omg"},
		[]Fulfillable{
			&types.StringConstraint{},
			&types.NilConstraint{},
			&types.NilConstraint{},
		},
	)

	assert.Equal(t, 1, len(result.GetPassed()))
	assert.Equal(t, 2, len(result.GetFailed()))
}

func TestCheckGroup_withContext(t *testing.T) {
	result := CheckGroup(
		"omg",
		map[string]interface{}{"something": "omg"},
		[]Fulfillable{
			&context.RequiredConstraint{Fields: []string{"something"}},
			&context.RequiredConstraint{Fields: []string{"nope"}},
			&context.RequiredConstraint{Fields: []string{"nope2"}},
		},
	)

	assert.Equal(t, 1, len(result.GetPassed()))
	assert.Equal(t, 2, len(result.GetFailed()))
}

func TestCheckContext(t *testing.T) {
	result := CheckContext(
		map[string]interface{}{"something": "omg"},
		[]ContextFulfillable{
			&context.RequiredConstraint{Fields: []string{"something"}},
			&context.RequiredConstraint{Fields: []string{"nope"}},
			&context.RequiredConstraint{Fields: []string{"nope2"}},
		},
	)

	assert.Equal(t, 1, len(result.GetPassed()))
	assert.Equal(t, 2, len(result.GetFailed()))
}
