package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruthy(t *testing.T) {
	var cases = []struct {
		input  []func() bool
		output bool
	}{
		{
			[]func() bool{},
			false,
		},
		{
			[]func() bool{
				func() bool { return false },
				func() bool { return true },
			},
			true,
		},
		{
			[]func() bool{
				func() bool { return true },
				func() bool { panic("I should not be called") },
			},
			true,
		},
	}

	for _, current := range cases {
		assert.Equal(t, Truthy(current.input), current.output)
	}
}

func TestFalsy(t *testing.T) {
	var cases = []struct {
		input  []func() bool
		output bool
	}{
		{
			[]func() bool{},
			true,
		},
		{
			[]func() bool{
				func() bool { return true },
				func() bool { return false },
			},
			false,
		},
		{
			[]func() bool{
				func() bool { return false },
				func() bool { panic("I should not be called") },
			},
			false,
		},
	}

	for _, current := range cases {
		assert.Equal(t, Falsy(current.input), current.output)
	}
}
