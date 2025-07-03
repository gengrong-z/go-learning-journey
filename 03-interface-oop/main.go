package main

import (
	rpg "go-learning-journey/03-interface-oop/rpg"
)

func main() {
	party := []rpg.Character{
		rpg.NewWarrior("Arthas"),
		rpg.NewMage("Jaina"),
	}

	rpg.SimulateTurn(party)

	for _, c := range party {
		rpg.PrintStatus(c)
		rpg.PrintRoleDetail(c)
	}
}
