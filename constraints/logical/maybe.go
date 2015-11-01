package logical

import (
	"github.com/etcinit/spark/constraints/types"
	"github.com/etcinit/spark/core"
	"github.com/etcinit/spark/fulfillment"
)

// MaybeConstraint is a constraint that can be satisfied when either its inner
// constraint is met or the passed value is nil.
type MaybeConstraint struct {
	Left fulfillment.Fulfillable
}

// IsFulfilled returns whether or not its inner constraint is met or the Value
// is nil.
func (c *MaybeConstraint) IsFulfilled(input interface{}) bool {
	nilConstraint := types.NilConstraint{}

	return core.Truthy([]func() bool{
		func() bool { return c.Left.IsFulfilled(input) },
		func() bool { return nilConstraint.IsFulfilled(input) },
	})
}
