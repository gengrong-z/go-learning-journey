package main

import (
	"bufio"
	"fmt"
	"go-learning-journey/week01-basic/todo_cli/todo"
	"os"
	"strings"
)

func main() {
	todos := []todo.Item{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("📝 TODO CLI（输入 help 查看命令）")

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
			fmt.Println("command list: add <task> | list | exit")
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
		case "exit":
			fmt.Println("👋 good bye")
		default:
			fmt.Println("unknown command")
		}
	}
}
