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
		{expectedOutcome: false, testId: "1", input: []bool{false}},
		{expectedOutcome: false, testId: "2", input: []bool{false, true, true}},
		{expectedOutcome: false, testId: "3", input: []bool{false, true, true}},
		{expectedOutcome: false, testId: "4", input: []bool{}},
		{expectedOutcome: true, testId: "5", input: []bool{true, true, true}},
		{expectedOutcome: true, testId: "6", input: []bool{true}},
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
		{expectedOutcome: false, testId: "1", input: []bool{false}},
		{expectedOutcome: true, testId: "2", input: []bool{false, true, true}},
		{expectedOutcome: true, testId: "3", input: []bool{false, true, true}},
		{expectedOutcome: false, testId: "4", input: []bool{}},
		{expectedOutcome: true, testId: "5", input: []bool{true, true, true}},
		{expectedOutcome: true, testId: "6", input: []bool{true}},
	}

	for _, test := range tests {
		t.Run(test.testId, func(t *testing.T) {
			assert.Equal(t, test.expectedOutcome, Or(test.input...), test.testId)
		})

	}
}
