package spec

import "github.com/etcinit/spark/fulfillment"

// Validator is a simple wrapper around a Spec which provides enough
// functionality to automatically generate error messages or responses from the
// result of running user input through a Spec.
type Validator struct {
	Spec     Spec
	Defaults map[string]interface{}
}

// ValidatorResult constaints the result of the spec with some additional
// messages explaining why some part of the spec did not pass (if any).
type ValidatorResult struct {
	SpecResult      Result
	Messages        map[string][]string
	ContextMessages []string
}

// Run runs the internal Spec against the provided input and then processes the
// result by extracting description messages on constraints that failed (if
// any).
func (v *Validator) Run(input map[string]interface{}) *ValidatorResult {
	// Add default value to the input array.
	for index, value := range v.Defaults {
		if _, exists := input[index]; exists == false {
			input[index] = value
		}
	}

	// Run the input through the spec.
	specResult := v.Spec.Run(input)

	return &ValidatorResult{
		SpecResult:      *specResult,
		Messages:        extractFailedMessagesFromResult(specResult, input),
		ContextMessages: extractFailedContextMessagesFromResult(specResult, input),
	}
}

func extractFailedMessagesFromResult(
	result *Result,
	context map[string]interface{},
) map[string][]string {
	messages := map[string][]string{}

	for field, constraints := range result.GetFailed() {
		messages[field] = []string{}

		for _, constraint := range constraints {
			messages[field] = append(
				messages[field],
				getConstraintDescription(constraint, field, context[field], context),
			)
		}
	}

	return messages
}

func extractFailedContextMessagesFromResult(
	result *Result,
	context map[string]interface{},
) []string {
	messages := []string{}

	for _, constraint := range result.GetContextFailed() {
		messages = append(
			messages,
			getConstraintDescription(constraint, "", nil, context),
		)
	}

	return messages
}

func getConstraintDescription(
	current fulfillment.Fulfillable,
	field string,
	value interface{},
	context map[string]interface{},
) string {
	if describable, ok := current.(fulfillment.FailureDescribable); ok == true {
		return describable.DescribeFailure(field, value, context)
	}

	contextDescribable, ok := current.(fulfillment.ContextFailureDescribable)
	if ok == true {
		return contextDescribable.DescribeContextFailure(context)
	}

	return "A constraint did not pass for the value provided on this field. " +
		"The constraint did not provide a description on why it did not pass."
}
