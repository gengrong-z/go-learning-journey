package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-learning-journey/02-cli-practice/todo"
	"log"
	"os"
	"strings"
)

const FileName = "week01-basic/todo_cli/todo.json"

func main() {
	var todos []todo.Item
	scanner := bufio.NewScanner(os.Stdin)

	todos = LoadList()
	if todos != nil {
		fmt.Println("loading todos successfully!")
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
			fmt.Println("command list: add <task> | list | done <task no> | delete <task no> | exit ")
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
			fmt.Println("👋 good bye")
			SaveList(todos)
			return
		default:
			fmt.Println("unknown command")
		}
	}
}

func LoadList() []todo.Item {
	var todos []todo.Item
	file, err := os.ReadFile(FileName)
	if err != nil {
		log.Println("history not found")
		return nil
	}
	err = json.Unmarshal(file, &todos)
	if err != nil {
		log.Fatal("failed to load todos file")
		return nil
	}
	return todos
}

func SaveList(todos []todo.Item) bool {
	data, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("failed to save todos file")
		return false
	}
	err = os.WriteFile(FileName, data, os.ModePerm)
	if err != nil {
		log.Println("failed to save todos file")
		return false
	}
	return true
}
