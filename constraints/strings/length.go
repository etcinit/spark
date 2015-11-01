package strings

// LengthConstraint requires provided values to be within a range in order to
// be satisfied.
type LengthConstraint struct {
	Min int
	Max int
}

// IsFulfilled returns whether or not the string is within the required range.
func (c *LengthConstraint) IsFulfilled(input interface{}) bool {
	if str, ok := input.(string); ok == true {
		stringLength := len(str)

		if c.Min > stringLength {
			return false
		} else if c.Max < stringLength {
			return false
		}

		return true
	}

	return false
}
