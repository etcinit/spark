package fulfillment

// GroupResult represents the result of checking multiple constraints against
// one value.
type GroupResult struct {
	passed []Fulfillable
	failed []Fulfillable
}

// GroupContextResult represents the result of checking multiple constraints
// against a map.
type GroupContextResult struct {
	passed []ContextFulfillable
	failed []ContextFulfillable
}

// GetPassed gets all the constraints that passed.
func (r *GroupResult) GetPassed() []Fulfillable {
	return r.passed
}

// GetFailed gets all the constraints that failed.
func (r *GroupResult) GetFailed() []Fulfillable {
	return r.failed
}

// GetPassed gets all the constraints that passed.
func (r *GroupContextResult) GetPassed() []ContextFulfillable {
	return r.passed
}

// GetFailed gets all the constraints that failed.
func (r *GroupContextResult) GetFailed() []ContextFulfillable {
	return r.failed
}

// CheckGroup checks if the provided input value matches or fails the specified
// group of constraints.
func CheckGroup(
	input interface{},
	context map[string]interface{},
	constraints []Fulfillable,
) *GroupResult {
	result := GroupResult{
		passed: []Fulfillable{},
		failed: []Fulfillable{},
	}

	for _, constraint := range constraints {
		if contextConstraint, ok := constraint.(ContextFulfillable); ok == true {
			if contextConstraint.IsFulfilledByContext(input, context) {
				result.passed = append(result.passed, constraint)
			} else {
				result.failed = append(result.failed, constraint)
			}
		} else if constraint.IsFulfilled(input) {
			result.passed = append(result.passed, constraint)
		} else {
			result.failed = append(result.failed, constraint)
		}
	}

	return &result
}

// CheckContext checks if the provided input passes all the provided context
// constraints.
func CheckContext(
	context map[string]interface{},
	constraints []ContextFulfillable,
) *GroupContextResult {
	result := GroupContextResult{
		passed: []ContextFulfillable{},
		failed: []ContextFulfillable{},
	}

	for _, constraint := range constraints {
		if constraint.IsFulfilledByContext(nil, context) {
			result.passed = append(result.passed, constraint)
		} else {
			result.failed = append(result.failed, constraint)
		}
	}

	return &result
}
