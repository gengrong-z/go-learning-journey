package todo

import "fmt"

type Item struct {
	Content string
	Done    bool
}

func New(content string) Item {
	return Item{Content: content, Done: false}
}

func PrintAll(todos []Item) {
	if len(todos) == 0 {
		fmt.Println("no todos")
		return
	}

	for i, todo := range todos {
		status := "❌"
		if todo.Done {
			status = "✅"
		}

		fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Content)
	}
}
