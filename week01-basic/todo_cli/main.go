package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-learning-journey/week01-basic/todo_cli/todo"
	"os"
	"strings"
)

const FileName = "week01-basic/todo_cli/todo.json"

func main() {
	var todos []todo.Item
	scanner := bufio.NewScanner(os.Stdin)

	if true {
		todos = LoadList()
		fmt.Println("loading history...")
	}
	fmt.Println("📝 TODO CLI（input help to view command）")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			fmt.Println("command list: add <task> | list | done <task no> | exit ")
		case "add":
			if len(args) < 2 {
				fmt.Println("Please enter task")
				continue
			}
			content := strings.Join(args[1:], " ")
			todos = append(todos, todo.New(content))
			fmt.Println("✅ task added")
		case "list":
			todo.PrintAll(todos)
		case "done":
			if len(args) != 2 {
				fmt.Println("Please enter task no")
				continue
			}
			todo.MarkDone(&todos, args[1])
			SaveList(todos)
		case "delete":
			if len(args) != 2 {
				fmt.Println("Please enter task no")
				continue
			}
			todo.Delete(&todos, args[1])
			SaveList(todos)
		case "exit":
			SaveList(todos)
			fmt.Println("👋 good bye")
			return
		default:
			fmt.Println("unknown command")
		}
	}
}

func LoadList() []todo.Item {
	var todos []todo.Item
	file, _ := os.ReadFile(FileName)
	json.Unmarshal(file, &todos)
	// todo error handler
	return todos
}

func SaveList(todos []todo.Item) {
	data, _ := json.Marshal(todos)
	os.WriteFile(FileName, data, os.ModePerm)
	// todo error handler
}
