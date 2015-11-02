package spec

import "github.com/etcinit/spark/fulfillment"

// Spec represents a grouping of constraints that losely or stricly define how
// an object stored in a map of strings to values should look like.
//
// The purpose of Spec is to facilitate user input validation and generation of
// repsonse messages with explanations of why validation passed or failed.
type Spec struct {
	Fields  map[string][]fulfillment.Fulfillable
	Check   []string
	Context []fulfillment.ContextFulfillable
}

// Run runs the provided input through a group of constraints and return the
// result.
func (s *Spec) Run(input map[string]interface{}) *Result {
	checked := map[string]bool{}
	result := NewResult()

	// Fields in Check will always be checked.
	for _, field := range s.Check {
		if constraints, ok := s.GetField(field); ok == true {
			groupResult := fulfillment.CheckGroup(input[field], constraints)

			result.SetFieldResults(field, groupResult)
			checked[field] = true
		}
	}

	for field, constraints := range s.Fields {
		if _, exists := checked[field]; exists == true {
			continue
		}

		groupResult := fulfillment.CheckGroup(input[field], constraints)

		result.SetFieldResults(field, groupResult)
		checked[field] = true
	}

	result.SetContextResults(fulfillment.CheckContext(
		input,
		s.Context,
	))

	return &result
}

// GetField attempts to get the constraints for a specific field.
func (s *Spec) GetField(key string) ([]fulfillment.Fulfillable, bool) {
	return s.Fields[key]
}
