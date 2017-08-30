package work

import "testing"

func newTaskorFatal(t *testing.T, title string) *Task {
	task, err := NewTask(title)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	return task
}
func TestNewTask(t *testing.T) {
	task, err := NewTask("learn Go")
	if err != nil {
		t.Fatalf("unexpected error")
	}
	if task.Title != "learn Go" {
		t.Errorf("expected learn Go, got %v", task.Title)
	}
}

func TestNewTaskWithEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Errorf("expected 'empty title' error, got nil")
	}
	// if task != nil {
	// t.Errorf("expected nil, got %v", task)
	// }
}

func TestSaveTask(t *testing.T) {
	task := newTaskorFatal(t, "learn Go")

	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func TestSaveTaskAndRetrieve(t *testing.T) {
	task := newTaskorFatal(t, "learn Go")

	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	all := m.All()
	if len(all) != 1 {
		t.Fatalf("expected one task, got %v", len(all))
	}

	if all[0].Title != task.Title {
		t.Errorf("expected title %q, got %q", task.Title, all[0].Title)
	}

}
