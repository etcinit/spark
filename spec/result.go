package spec

import "github.com/etcinit/spark/fulfillment"

// Result represents the result of a check performed on some input against
// a spec.
type Result struct {
	met           bool
	passed        map[string][]fulfillment.Fulfillable
	failed        map[string][]fulfillment.Fulfillable
	contextPassed []fulfillment.ContextFulfillable
	contextFailed []fulfillment.ContextFulfillable
	checked       []string
}

// NewResult creates a new empty instance of a Result.
func NewResult() *Result {
	return &Result{
		met:           true,
		passed:        map[string][]fulfillment.Fulfillable{},
		failed:        map[string][]fulfillment.Fulfillable{},
		contextPassed: []fulfillment.ContextFulfillable{},
		contextFailed: []fulfillment.ContextFulfillable{},
		checked:       []string{},
	}
}

// SetFieldResults sets the constraints that passed and failed for single
// field.
func (r *Result) SetFieldResults(
	field string,
	group *fulfillment.GroupResult,
) {
	r.passed[field] = group.GetPassed()
	r.failed[field] = group.GetFailed()

	if len(r.failed[field]) > 0 {
		r.met = false
	}
}

// SetContextResults sets the constraints that passed and failed for the
// context.
func (r *Result) SetContextResults(group *fulfillment.GroupContextResult) {
	r.contextPassed = group.GetPassed()
	r.contextFailed = group.GetFailed()

	if len(r.contextFailed) > 0 {
		r.met = false
	}
}

// SetChecked sets the names of the fields that were checked.
func (r *Result) SetChecked(checked []string) {
	r.checked = checked
}

// GetChecked gets the name of the fields that were checked.
func (r *Result) GetChecked() []string {
	return r.checked
}

// Passed returns true if the input passed all the constraints in the spec.
func (r *Result) Passed() bool {
	return r.met
}

// Failed returns true if the input failed one or more of the constraints in
// the spec.
func (r *Result) Failed() bool {
	return !r.met
}

// GetPassed returns a map of fields with the constraints that passed.
func (r *Result) GetPassed() map[string][]fulfillment.Fulfillable {
	return r.passed
}

// GetFailed returns a map of fields with the constraints that failed.
func (r *Result) GetFailed() map[string][]fulfillment.Fulfillable {
	return r.failed
}

// GetContextPassed returns a slice of all the context constraints that passed.
func (r *Result) GetContextPassed() []fulfillment.ContextFulfillable {
	return r.contextPassed
}

// GetContextFailed returns a slice of all the context constraints that failed.
func (r *Result) GetContextFailed() []fulfillment.ContextFulfillable {
	return r.contextFailed
}
