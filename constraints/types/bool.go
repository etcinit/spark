package types

// BoolConstraint requires provided values to be a bool type.
type BoolConstraint struct{}

// IsFulfilled returns whether or not the provided value is a bool.
func (c *BoolConstraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(bool)

	return ok
}
