package spec

import (
	"testing"

	"github.com/etcinit/spark/constraints/context"
	"github.com/etcinit/spark/constraints/numbers"
	"github.com/etcinit/spark/constraints/strings"
	"github.com/etcinit/spark/constraints/types"
	"github.com/etcinit/spark/fulfillment"
	"github.com/stretchr/testify/assert"
)

func TestSpecRun(t *testing.T) {
	userSpec := Spec{
		Fields: map[string][]fulfillment.Fulfillable{
			"first_name": []fulfillment.Fulfillable{
				&types.StringConstraint{},
				&strings.LengthConstraint{Min: 1, Max: 255},
			},
			"last_name": []fulfillment.Fulfillable{
				&types.StringConstraint{},
				&strings.LengthConstraint{Min: 1, Max: 255},
			},
			"zip": []fulfillment.Fulfillable{
				&types.IntConstraint{},
				&numbers.BetweenIntConstraint{Min: 10000, Max: 99999},
			},
		},
		Check: []string{"first_name", "last_name"},
		Context: []fulfillment.ContextFulfillable{
			&context.RequiredConstraint{Fields: []string{"first_name", "zip"}},
		},
	}

	input := map[string]interface{}{
		"first_name": "Bobby",
		"last_name":  "Tables",
		"zip":        12345,
	}
	input2 := map[string]interface{}{
		"first_name": "Bobby",
		"last_name":  "Tables",
	}
	input3 := map[string]interface{}{
		"first_name": "Bobby",
		"last_name":  nil,
		"zip":        12345,
	}

	result := userSpec.Run(input)
	result2 := userSpec.Run(input2)
	result3 := userSpec.Run(input3)

	assert.True(t, result.Passed())
	assert.True(t, result2.Failed())
	assert.True(t, result3.Failed())
}

func TestSpecGetFields(t *testing.T) {
	nameConstraints := []fulfillment.Fulfillable{
		&types.StringConstraint{},
		&strings.LengthConstraint{Min: 1, Max: 255},
	}
	zipConstraints := []fulfillment.Fulfillable{
		&types.IntConstraint{},
		&numbers.BetweenIntConstraint{Min: 10000, Max: 99999},
	}

	userSpec := Spec{
		Fields: map[string][]fulfillment.Fulfillable{
			"first_name": nameConstraints,
			"last_name":  nameConstraints,
			"zip":        zipConstraints,
		},
		Check: []string{"first_name", "last_name"},
		Context: []fulfillment.ContextFulfillable{
			&context.RequiredConstraint{Fields: []string{"first_name", "zip"}},
		},
	}

	assertConstraint := func(
		expected []fulfillment.Fulfillable,
		fieldName string,
	) {
		value, ok := userSpec.GetField(fieldName)

		assert.True(t, ok)
		assert.Equal(t, expected, value)
	}

	assertConstraint(nameConstraints, "first_name")
	assertConstraint(nameConstraints, "last_name")
	assertConstraint(zipConstraints, "zip")
}
