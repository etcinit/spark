package spec

import (
	"testing"

	"github.com/etcinit/spark/constraints/context"
	"github.com/etcinit/spark/constraints/types"
	"github.com/etcinit/spark/fulfillment"
	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {
	result := NewResult()

	assert.NotNil(t, result)
}

func TestResultSetFieldResults(t *testing.T) {
	groupResult := fulfillment.NewGroupResult(
		[]fulfillment.Fulfillable{&types.StringConstraint{}},
		[]fulfillment.Fulfillable{&types.IntConstraint{}, &types.IntConstraint{}},
	)

	result := Result{
		passed: map[string][]fulfillment.Fulfillable{},
		failed: map[string][]fulfillment.Fulfillable{},
	}

	result.SetFieldResults("hello", groupResult)

	assert.Equal(t, 1, len(result.passed["hello"]))
	assert.Equal(t, 2, len(result.failed["hello"]))
}

func TestResultSetContextResult(t *testing.T) {
	groupContextResult := fulfillment.NewGroupContextResult(
		[]fulfillment.ContextFulfillable{&context.RequiredConstraint{}},
		[]fulfillment.ContextFulfillable{
			&context.RequiredConstraint{},
			&context.RequiredConstraint{},
		},
	)

	result := Result{
		contextPassed: []fulfillment.ContextFulfillable{},
		contextFailed: []fulfillment.ContextFulfillable{},
	}

	result.SetContextResults(groupContextResult)

	assert.Equal(t, 1, len(result.contextPassed))
	assert.Equal(t, 2, len(result.contextFailed))
}

func TestResultSetChecked(t *testing.T) {
	result := NewResult()

	result.SetChecked([]string{"hello", "world"})

	assert.Equal(t, []string{"hello", "world"}, result.checked)
}

func TestResultGetChecked(t *testing.T) {
	result := NewResult()

	result.checked = []string{"hello", "world"}

	assert.Equal(t, []string{"hello", "world"}, result.GetChecked())
}

func TestResultPassed(t *testing.T) {
	result := NewResult()

	assert.True(t, result.Passed())
}

func TestResultFailed(t *testing.T) {
	result := NewResult()

	assert.False(t, result.Failed())
}

func TestResultGetPassed(t *testing.T) {
	passed := map[string][]fulfillment.Fulfillable{
		"hello": []fulfillment.Fulfillable{&types.StringConstraint{}},
	}

	result := Result{passed: passed}

	assert.Equal(t, passed, result.GetPassed())
}

func TestResultGetFailed(t *testing.T) {
	failed := map[string][]fulfillment.Fulfillable{
		"hello": []fulfillment.Fulfillable{&types.StringConstraint{}},
	}

	result := Result{failed: failed}

	assert.Equal(t, failed, result.GetFailed())
}

func TestResultGetContextPassed(t *testing.T) {
	passed := []fulfillment.ContextFulfillable{
		&context.RequiredConstraint{},
	}

	result := Result{contextPassed: passed}

	assert.Equal(t, passed, result.GetContextPassed())
}

func TestResultGetContextFailed(t *testing.T) {
	failed := []fulfillment.ContextFulfillable{
		&context.RequiredConstraint{},
	}

	result := Result{contextFailed: failed}

	assert.Equal(t, failed, result.GetContextFailed())
}
