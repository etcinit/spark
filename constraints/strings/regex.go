package strings

import "regexp"

// RegexConstraint requires provided values to match a regex.
type RegexConstraint struct {
	Regex string
}

// IsFulfilled returns whether or not a string matches a regex.
func (c *RegexConstraint) IsFulfilled(input interface{}) bool {
	if str, ok := input.(string); ok == true {
		matched, _ := regexp.MatchString(c.Regex, str)

		return matched
	}

	return false
}
