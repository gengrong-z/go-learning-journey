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
	log := battlelog.NewBattleLog()
	rpg.StartBattle(party, log)
}
