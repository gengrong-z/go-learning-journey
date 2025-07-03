package main

import (
	"fmt"
	rpg "go-learning-journey/03-interface-oop/rpg"
)

func main() {
	party := []rpg.Character{
		rpg.NewWarrior("Arthas"),
		rpg.NewMage("Jaina"),
	}

	for _, c := range party {
		rpg.PrintStatus(c)
		rpg.PrintRoleDetail(c)
		fmt.Println(c.Attack())
		fmt.Println(c.Defend())
		fmt.Println()
	}
}
