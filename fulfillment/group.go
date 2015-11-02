package fulfillment

import "github.com/etcinit/spark/fulfillment"

// GroupResult represents the result of checking multiple constraints against
// one value.
type GroupResult struct {
	passed []fulfillment.Fulfillable
	failed []fulfillment.Fulfillable
}

// GroupContextResult represents the result of checking multiple constraints
// against a map.
type GroupContextResult struct {
	passed []fulfillment.ContextFulfillable
	failed []fulfillment.ContextFulfillable
}

// GetPassed gets all the constraints that passed.
func (r *GroupResult) GetPassed() []fulfillment.Fulfillable {
	return r.passed
}

// GetFailed gets all the constraints that failed.
func (r *GroupResult) GetFailed() []fulfillment.Fulfillable {
	return r.failed
}

// GetPassed gets all the constraints that passed.
func (r *GroupContextResult) GetPassed() []fulfillment.ContextFulfillable {
	return r.passed
}

// GetFailed gets all the constraints that failed.
func (r *GroupContextResult) GetFailed() []fulfillment.ContextFulfillable {
	return r.failed
}

// CheckGroup checks if the provided input value matches or fails the specified
// group of constraints.
func CheckGroup(
	input interface{},
	context map[string]interface{},
	constraints []fulfillment.Fulfillable,
) *GroupResult {
	result := GroupResult{
		passed: []fulfillment.Fulfillable{},
		failed: []fulfillment.Fulfillable{},
	}

	for _, constraint := range constraints {
		if contextConstraint, ok := constraint.(ContextFulfillable); ok == true {
			if contextConstraint.IsFulfilledByContext(input, context) {
				result.passed = append(result.passed, constraint)
			} else {
				result.failed = append(result.passed, constraint)
			}
		} else if constraint.IsFulfilled(input) {
			result.passed = append(result.passed, constraint)
		} else {
			result.failed = append(result.passed, constraint)
		}
	}

	return &result
}

// CheckContext checks if the provided input passes all the provided context
// constraints.
func CheckContext(
	context map[string]interface{},
	constraints []fulfillment.ContextFulfillable,
) *GroupContextResult {
	result := GroupContextResult{
		passed: []fulfillment.ContextFulfillable{},
		failed: []fulfillment.ContextFulfillable{},
	}

	for _, constraint := range constraints {
		if contextConstraint.IsFulfilledByContext(input, context) {
			result.passed = append(result.passed, constraint)
		} else {
			result.failed = append(result.passed, constraint)
		}
	}

	return &result
}
