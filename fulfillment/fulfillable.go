package fulfillment

// Fulfillable represents an object that is capable of asserting whether or not
// a certain value meets its own constraints or rules.
type Fulfillable interface {
	IsFulfilled(interface{}) bool
}
