package main

import (
	"fmt"
	"go-learning-journey/03-interface-oop/battlelog"
	"go-learning-journey/03-interface-oop/core"
	"go-learning-journey/03-interface-oop/rpg"
)

func main() {
	party := []core.Character{
		rpg.NewWarrior("Arthas"),
		rpg.NewMage("Jaina"),
		rpg.NewRogue("Loki"),
	}

	fmt.Println("🎮 Battle Start!")

	round := 1
	log := battlelog.NewBattleLog()

	fmt.Println("SimulateAbilities Start")
	rpg.SimulateAbilities(party, round, log)

	fmt.Println("\nSimulateTurn Start")
	rpg.SimulateTurn(party, round, log)

	//for _, c := range party {
	//	fmt.Println()
	//	rpg.PrintStatus(c)
	//	rpg.PrintRoleDetail(c)
	//}
}
