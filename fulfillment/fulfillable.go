package fulfillment

// Fulfillable represents an object that is capable of asserting whether or not
// a certain value meets its own constraints or rules.
type Fulfillable interface {
	IsFulfilled(interface{}) bool
}

// ContextFulfillable represents an object that is capable of asserting whether
// or not a certain value meets its own constraints or rules.
//
// Unlike regular Fulfillables, ContextFulfillable also consider the context of
// a value. For user input, context could be the other fields in the form.
//
// Additionally, ContextFulfillables can also just ignore the provided value
// completely and act using only information from the context.
type ContextFulfillable interface {
	IsFulfilled(interface{}) bool
	IsFulfilledByContext(interface{}, map[string]interface{}) bool
}
