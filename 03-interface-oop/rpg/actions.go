package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/battlelog"
	"go-learning-journey/03-interface-oop/core"
)

func StartBattle(characters []core.Character, log *battlelog.BattleLog) {
	player1 := characters[0]
	player2 := characters[1]
	for round := 1; ; round++ {
		if round/2 == 0 {
			SimulateTurn(player1, player2, round, log)
		} else {
			SimulateTurn(player2, player1, round, log)
		}

		if !player1.GetStatus().IsAlive() {
			fmt.Println("Battle End")
			log.Print()
			fmt.Printf(" 🏆 Winner is %s ", player2.Name())
			break
		} else if !player2.GetStatus().IsAlive() {
			fmt.Println("Battle End")
			log.Print()
			fmt.Printf(" 🏆 Winner is %s ", player1.Name())
			break
		}
	}
}

func SimulateTurn(player, target core.Character, round int, log *battlelog.BattleLog) {

	UseAbilities(log, player, target, round)

}

func UseAbilities(log *battlelog.BattleLog, c core.Character, target core.Character, round int) {
	if caster, ok := c.(interface{ GetAbilities() []core.Ability }); ok {
		for _, ability := range caster.GetAbilities() {
			if ability.IsNotAvailable(c) {
				continue
			}
			log.Add(battlelog.LogEntry{
				Round:  round,
				Actor:  c.Name(),
				Action: ability.Name(),
				Target: target.Name(),
				Effect: ability.Use(c, target),
			})
			return
		}
	}

	effect := target.TakeDamage(c.AttackerPower())
	log.Add(battlelog.LogEntry{
		Round:  round,
		Actor:  c.Name(),
		Action: c.Attack(),
		Target: target.Name(),
		Effect: effect,
	})
}
