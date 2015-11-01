package types

// StringConstraint requires provided values to be a string type.
type StringConstraint struct{}

// IsFulfilled returns whether or not the provided value is a string.
func (c *StringConstraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(string)

	return ok
}
