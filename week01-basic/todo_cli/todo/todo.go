package todo

import (
	"fmt"
	"strconv"
	"time"
)

// Item It`s one task on the todo list
type Item struct {
	Content    string `json:"content"`
	Done       bool   `json:"done"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// New Create a new todo item
func New(content string) Item {
	return Item{
		Content:    content,
		Done:       false,
		CreateTime: time.Now().Format(time.DateOnly),
		UpdateTime: time.Now().Format(time.DateOnly),
	}
}

// PrintAll Print all item from the todo list
func PrintAll(todos []Item) {
	if len(todos) == 0 {
		fmt.Println("no todos")
		return
	}

	fmt.Println("No Create Update Status Content")
	for i, todo := range todos {
		status := "❌"
		if todo.Done {
			status = "✅"
		}

		fmt.Printf("%d. %s %s [%s] %s\n", i+1, todo.CreateTime, todo.UpdateTime, status, todo.Content)
	}
}

// MarkDone Mark todo item to complete by number
func MarkDone(todos *[]Item, indexStr string) {
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 1 || index > len(*todos) {
		fmt.Println("⚠️ invalid index")
		return
	}

	(*todos)[index-1].Done = true
	(*todos)[index-1].UpdateTime = time.Now().Format(time.DateOnly)
	fmt.Println("✅ task is marked done")
}

// Delete delete todo item from todo list by number
func Delete(todos *[]Item, indexStr string) {
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 1 || index > len(*todos) {
		fmt.Println("⚠️ invalid index")
		return
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	fmt.Println("task has been deleted")
}
