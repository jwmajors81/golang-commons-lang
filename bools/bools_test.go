package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	var tests = []struct {
		input           []bool
		expectedOutcome bool
		testId          string
	}{
		{expectedOutcome: false, testId: "single false", input: []bool{false}},
		{expectedOutcome: false, testId: "single false and multiple true", input: []bool{false, true, true}},
		{expectedOutcome: false, testId: "multiple false and single true", input: []bool{false, true, false}},
		{expectedOutcome: false, testId: "empty", input: []bool{}},
		{expectedOutcome: true, testId: "all true", input: []bool{true, true, true}},
		{expectedOutcome: true, testId: "single true", input: []bool{true}},
	}

	for _, test := range tests {
		t.Run(test.testId, func(t *testing.T) {
			assert.Equal(t, test.expectedOutcome, And(test.input...))
		})
	}
}

func TestOr(t *testing.T) {
	var tests = []struct {
		input           []bool
		expectedOutcome bool
		testId          string
	}{
		{expectedOutcome: false, testId: "single false", input: []bool{false}},
		{expectedOutcome: true, testId: "single false and multiple true", input: []bool{false, true, true}},
		{expectedOutcome: true, testId: "multiple false and single true", input: []bool{false, true, false}},
		{expectedOutcome: false, testId: "empty", input: []bool{}},
		{expectedOutcome: true, testId: "all true", input: []bool{true, true, true}},
		{expectedOutcome: true, testId: "single true", input: []bool{true}},
	}

	for _, test := range tests {
		t.Run(test.testId, func(t *testing.T) {
			assert.Equal(t, test.expectedOutcome, Or(test.input...), test.testId)
		})

	}
}
