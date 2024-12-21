package tasksort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortExampleTasksCorrectly(t *testing.T) {
	tests := map[string]struct {
		input []Task
		result []Task
	} {
		"sort the given example tasks correctly":  {
			input: []Task{
				{Name: "task-1", Command: "touch /tmp/file1"},
				{Name: "task-2", Command: "cat /tmp/file1", Requires: []string{"task-3"}},
				{Name: "task-3", Command: "echo 'Hello World' > /tmp/file1", Requires: []string{"task-1"}},
				{Name: "task-4", Command: "rm /tmp/file1", Requires: []string{"task-2", "task-3"}},
			},
			result: []Task{
				{Name: "task-1", Command: "touch /tmp/file1"},
				{Name: "task-3", Command: "echo 'Hello World' > /tmp/file1", Requires: []string{"task-1"}},
				{Name: "task-2", Command: "cat /tmp/file1", Requires: []string{"task-3"}},
				{Name: "task-4", Command: "rm /tmp/file1", Requires: []string{"task-2", "task-3"}},
			},
		},
		"sorting tasks without dependencies retains the original order": {
			input: []Task{
				{Name: "task-1", Command: "touch /tmp/file1"},
				{Name: "task-2", Command: "cat /tmp/file1"},
			},
			result: []Task{
				{Name: "task-1", Command: "touch /tmp/file1"},
				{Name: "task-2", Command: "cat /tmp/file1"},
			},
		},
		"sort copmlex dependencies combined with independent tasks": {
			input: []Task{
				{Name: "task-1", Command: "echo 'Task 1'"},
				{Name: "task-2", Command: "echo 'Task 2'", Requires: []string{"task-3"}},
				{Name: "task-3", Command: "echo 'Task 3'"},
				{Name: "task-4", Command: "echo 'Task 4'", Requires: []string{"task-2", "task-3"}},
				{Name: "task-5", Command: "echo 'Task 5'"},
			},
			result: []Task{
				{Name: "task-1", Command: "echo 'Task 1'"},
				{Name: "task-3", Command: "echo 'Task 3'"},
				{Name: "task-2", Command: "echo 'Task 2'", Requires: []string{"task-3"}},
				{Name: "task-4", Command: "echo 'Task 4'", Requires: []string{"task-2", "task-3"}},
				{Name: "task-5", Command: "echo 'Task 5'"},
			},
		},
	}
	  
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			sortedTasks := Sort(test.input)

			assert.Equal(t, test.result, sortedTasks)
		})
	}
}
