package types

// AnyIntConstraint requires provided values to be some int type.
type AnyIntConstraint struct{}

// IntConstraint requires provided values to be an int type.
type IntConstraint struct{}

// Int8Constraint requires provided values to be an int8 type.
type Int8Constraint struct{}

// Int16Constraint requires provided values to be an int16 type.
type Int16Constraint struct{}

// Int32Constraint requires provided values to be an int32 type.
type Int32Constraint struct{}

// Int64Constraint requires provided values to be an int64 type.
type Int64Constraint struct{}

// IsFulfilled returns whether or not the provided value is some int type.
func (c *AnyIntConstraint) IsFulfilled(input interface{}) bool {
	if _, ok := input.(int64); ok == true {
		return true
	}

	if _, ok := input.(int32); ok == true {
		return true
	}

	if _, ok := input.(int); ok == true {
		return true
	}

	if _, ok := input.(int8); ok == true {
		return true
	}

	if _, ok := input.(int16); ok == true {
		return true
	}

	return false
}

// IsFulfilled returns whether or not the provided value is an int.
func (c *IntConstraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(int)

	return ok
}

// IsFulfilled returns whether or not the provided value is an int8.
func (c *Int8Constraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(int8)

	return ok
}

// IsFulfilled returns whether or not the provided value is an int16.
func (c *Int16Constraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(int16)

	return ok
}

// IsFulfilled returns whether or not the provided value is an int32.
func (c *Int32Constraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(int32)

	return ok
}

// IsFulfilled returns whether or not the provided value is an int64.
func (c *Int64Constraint) IsFulfilled(input interface{}) bool {
	_, ok := input.(int64)

	return ok
}
