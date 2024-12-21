package tasksort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONFormatTasksCorrectly(t *testing.T) {
	tests := map[string]struct {
		input  []Task
		expected string
	}{
		"formating excludes the reuires property": {
			input: []Task{
				{Name: "t-1", Command: "touch /tmp/file1"},
				{Name: "t-2", Command: "cat /tmp/file1", Requires: []string{"t-1"}},
			},
			expected: `[{"name":"t-1","command":"touch /tmp/file1"},{"name":"t-2","command":"cat /tmp/file1"}]`,
		},
		"formating empty list of tasks returns null": {
			input: []Task{},
			expected: "null",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			sut := NewJSONFormater()
			actual := sut.Format(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}
