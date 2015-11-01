package logical

import (
	"github.com/etcinit/spec/core"
	"github.com/etcinit/spec/fulfillment"
)

// EitherConstraint is a constraint that can be satisfied when either one of
// its two inner constraints are met.
type EitherConstraint struct {
	Left  fulfillment.Fulfillable
	Right fulfillment.Fulfillable
}

// IsFulfilled returns whether or not the provided value is matches any of the
// two internal constraints.
func (c *EitherConstraint) IsFulfilled(input interface{}) bool {
	return core.Truthy([]func() bool{
		func() bool { return c.Left.IsFulfilled(input) },
		func() bool { return c.Right.IsFulfilled(input) },
	})
}
