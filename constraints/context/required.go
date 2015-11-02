package context

import "strings"

// RequiredConstraint is a constraint that requires all the specified fields to
// be present in the input/context.
type RequiredConstraint struct {
	Fields []string
}

// IsFulfilled always returns true since this constraint is effective on the
// context, not the value.
func (r *RequiredConstraint) IsFulfilled(value interface{}) bool {
	return true
}

// IsFulfilledByContext returns whether or not the provided context passes the
// constraint.
func (r *RequiredConstraint) IsFulfilledByContext(
	value interface{},
	context map[string]interface{},
) bool {
	for _, field := range r.Fields {
		if _, exists := context[field]; exists == false {
			return false
		}
	}

	return true
}

// DescribeContextFailure describes why the constraint failed.
func (r *RequiredConstraint) DescribeContextFailure(
	context map[string]interface{},
) string {
	missing := []string{}

	for _, field := range r.Fields {
		if _, exists := context[field]; exists == false {
			missing = append(missing, field)
		}
	}

	return "The following fields are missing from the input: " +
		strings.Join(missing, ", ")
}
