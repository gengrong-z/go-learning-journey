package todo

import (
	"testing"
)

func TestNew(t *testing.T) {
	item := New("a new test task")
	if item.Content != "a new test task" {
		t.Errorf("Expected 'a new test task', got %s", item.Content)
	}
	if item.Done != false {
		t.Error("Expected Done to be false by default")
	}
}

func TestDelete(t *testing.T) {
	var todos []Item
	todos = append(todos, New("test task 1"), New("test task 2"), New("test task 3"))
	Delete(&todos, "2")
	if len(todos) != 2 {
		t.Error("Expected 2 todos, got ", len(todos))
	}
}

func TestMarkDone(t *testing.T) {
	var todos []Item
	todos = append(todos, New("test task 1"))
	MarkDone(&todos, "1")
	if todos[0].Done != true {
		t.Error("Expected Done to be true, got ", todos[0].Done)
	}
}

// question: how to test this
func TestPrintAll(t *testing.T) {
	var todos []Item
	todos = append(todos, New("test task 1"))

	PrintAll(todos)
}
