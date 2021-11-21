//go:build unit || ci

package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStringSet struct {
	input    []string
	expected []string
}
type testMultiInputSet struct {
	input1   []string
	input2   []string
	expected []string
}

func TestIntersection(t *testing.T) {
	testData := []testMultiInputSet{
		{
			input1:   []string{"foo", "bar", "north"},
			input2:   []string{"foo", "say"},
			expected: []string{"foo"},
		},
		{
			input1:   []string{"foo", "bar", "north"},
			input2:   []string{"what", "say"},
			expected: []string{},
		},
		{
			input1:   []string{"foo", "bar", "north", "what"},
			input2:   []string{"bar", "what", "say"},
			expected: []string{"bar", "what"},
		},
	}
	for _, td := range testData {
		result := Intersection(td.input1, td.input2)

		assert.Equal(t, fmt.Sprintf("%s", td.expected), fmt.Sprintf("%s", result))
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testData := []testStringSet{
		{
			input:    []string{"foo", "foo"},
			expected: []string{"foo"},
		},
		{
			input:    []string{"foo", "bar"},
			expected: []string{"foo", "bar"},
		},
		{
			input:    []string{"foo", "bar", "bar"},
			expected: []string{"foo", "bar"},
		},
	}
	for _, td := range testData {
		result := RemoveDuplicates(td.input)

		assert.Equal(t, fmt.Sprintf("%s", td.expected), fmt.Sprintf("%s", result))
	}
}
