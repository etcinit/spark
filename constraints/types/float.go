package types

// AnyFloatConstraint requires provided values to be any float type
// (float32, float64).
type AnyFloatConstraint struct{}

// Float32Constraint requires provided values to be a float32 type
// (float32, float64).
type Float32Constraint struct{}

// Float64Constraint requires provided values to be a float64 type.
type Float64Constraint struct{}

// IsFulfilled returns whether or not the provided value is a float type
func (c *AnyFloatConstraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(float32)

	_, ok2 := input.(float64)

	return ok || ok2
}

// IsFulfilled returns whether or not the provided value is a float type
func (c *Float32Constraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(float32)

	return ok
}

// IsFulfilled returns whether or not the provided value is a float type
func (c *Float64Constraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(float64)

	return ok
}
