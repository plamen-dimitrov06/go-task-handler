package tasksort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBashFormatTasksCorrectly(t *testing.T) {
	tests := map[string]struct {
		input  []Task
		expected string
	}{
		"formating only includes the command": {
			input: []Task{
				{Name: "task-1", Command: "echo 'Task 1'"},
				{Name: "task-2", Command: "echo 'Task 2'", Requires: []string{"task-3"}},
			},
			expected: "#!/usr/bin/env bash\n\necho 'Task 1'\necho 'Task 2'",
		},
		"formating empty list of tasks return only bash shell path": {
			input: []Task{},
			expected: "#!/usr/bin/env bash\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			sut := NewBashFormater()
			actual := sut.Format(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}
