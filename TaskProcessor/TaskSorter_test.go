package TaskProcessor

import (
	"github.com/stretchr/testify/assert"
	"task-handler/Models"
	"testing"
)

func TestSortExampleTasksCorrectly(t *testing.T) {
	tasks := []Models.Task{
		{Name: "task-1", Command: "touch /tmp/file1"},
		{Name: "task-2", Command: "cat /tmp/file1", Requires: []string{"task-3"}},
		{Name: "task-3", Command: "echo 'Hello World' > /tmp/file1", Requires: []string{"task-1"}},
		{Name: "task-4", Command: "rm /tmp/file1", Requires: []string{"task-2", "task-3"}},
	}
	// slices are passed by reference (?), 
	// so we derefence them before passing to the sorting handler
	sut := TaskSorter{}
	sortedTasks := sut.Sort(append([]Models.Task(nil), tasks...))
	// correct order should be :
	tasks[1], tasks[2] = tasks[2], tasks[1]

	assert.Equal(t, tasks, sortedTasks)
}

func TestSortCorreclyWhenRequiredIsMissing(t *testing.T) {
	tasks := []Models.Task{
		{Name: "task-1", Command: "touch /tmp/file1"},
		{Name: "task-2", Command: "cat /tmp/file1"},
	}

	sut := TaskSorter{}
	sortedTasks := sut.Sort(append([]Models.Task(nil), tasks...))

	assert.Equal(t, tasks, sortedTasks)
}

func TestComplexDependenciesWithIndependentTasks(t *testing.T) {
	tasks := []Models.Task{
		{Name: "task-1", Command: "echo 'Task 1'"},
		{Name: "task-2", Command: "echo 'Task 2'", Requires: []string{"task-3"}},
		{Name: "task-3", Command: "echo 'Task 3'"},
		{Name: "task-4", Command: "echo 'Task 4'", Requires: []string{"task-2", "task-3"}},
		{Name: "task-5", Command: "echo 'Task 5'"},
	}

	sut := TaskSorter{}
	sortedTasks := sut.Sort(append([]Models.Task(nil), tasks...))
	tasks[1], tasks[2] = tasks[2], tasks[1]
	assert.Equal(t, tasks, sortedTasks)
}
