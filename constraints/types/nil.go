package types

// NilConstraint requires provided values to be nil.
type NilConstraint struct{}

// IsFulfilled returns whether or not the provided value is nil.
func (c *NilConstraint) IsFulfilled(input interface{}) bool {
	return input == nil
}
