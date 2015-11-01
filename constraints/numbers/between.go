package numbers

// BetweenIntConstraint requires provided values to be within a range in order
// to be satisfied.
type BetweenIntConstraint struct {
	Min int
	Max int
}

// BetweenInt8Constraint requires provided values to be within a range in order
// to be satisfied.
type BetweenInt8Constraint struct {
	Min int8
	Max int8
}

// BetweenInt16Constraint requires provided values to be within a range in
// order to be satisfied.
type BetweenInt16Constraint struct {
	Min int16
	Max int16
}

// BetweenInt32Constraint requires provided values to be within a range in
// order to be satisfied.
type BetweenInt32Constraint struct {
	Min int32
	Max int32
}

// BetweenInt64Constraint requires provided values to be within a range in
// order to be satisfied.
type BetweenInt64Constraint struct {
	Min int64
	Max int64
}

// BetweenFloat32Constraint requires provided values to be within a range in
// order to be satisfied.
type BetweenFloat32Constraint struct {
	Min float32
	Max float32
}

// BetweenFloat64Constraint requires provided values to be within a range in
// order to be satisfied.
type BetweenFloat64Constraint struct {
	Min float64
	Max float64
}

// IsFulfilled returns whether or not the int is within the required range.
func (c *BetweenIntConstraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(int); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}

// IsFulfilled returns whether or not the int8 is within the required range.
func (c *BetweenInt8Constraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(int8); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}

// IsFulfilled returns whether or not the int16 is within the required range.
func (c *BetweenInt16Constraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(int16); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}

// IsFulfilled returns whether or not the int32 is within the required range.
func (c *BetweenInt32Constraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(int32); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}

// IsFulfilled returns whether or not the int64 is within the required range.
func (c *BetweenInt64Constraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(int64); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}

// IsFulfilled returns whether or not the float32 is within the required range.
func (c *BetweenFloat32Constraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(float32); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}

// IsFulfilled returns whether or not the float64 is within the required range.
func (c *BetweenFloat64Constraint) IsFulfilled(input interface{}) bool {
	if number, ok := input.(float64); ok == true {
		if c.Min > number {
			return false
		} else if c.Max < number {
			return false
		}

		return true
	}

	return false
}
