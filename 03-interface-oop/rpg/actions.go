package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/battlelog"
	"go-learning-journey/03-interface-oop/core"
	"math/rand"
)

func selectTarget(c core.Character, party []core.Character) core.Character {
	for _, character := range party {
		if c.Name() == character.Name() {
			return character
		}
	}
	return nil
}

func selectTargetRandom(party []core.Character) core.Character {
	return party[rand.Intn(len(party))]
}

func StartBattle(characters []core.Character, log *battlelog.BattleLog) {
	for round := 1; ; round++ {
		SimulateTurn(characters, round, log)

		alive := len(characters)
		for _, c := range characters {
			if !c.GetStatus().IsAlive() {
				alive--
			}
		}
		if alive == 1 {
			fmt.Println("Battle End")
			for _, c := range characters {
				if c.GetStatus().IsAlive() {
					fmt.Printf(" 🏆 Winner is %s ", c.Name())
				}
			}
			break
		} else if alive == 0 {
			fmt.Println(" 🤝Not Winner")
		}
	}
}

func SimulateTurn(characters []core.Character, round int, log *battlelog.BattleLog) {
	for _, c := range characters {
		if !c.GetStatus().IsAlive() {
			continue
		}

		//target := selectTarget(c, characters)
		target := selectTargetRandom(characters)
		if target == nil {
			continue
		}

		UseAbilities(log, c, round)

		effect := target.TakeDamage(c.AttackerPower())

		log.Add(battlelog.LogEntry{
			Round:  round,
			Actor:  c.Name(),
			Action: c.Attack(),
			Target: target.Name(),
			Effect: effect,
		})
	}

	log.Print()
}

func UseAbilities(log *battlelog.BattleLog, c core.Character, round int) {
	if caster, ok := c.(interface{ GetAbilities() []core.Ability }); ok {
		for _, ability := range caster.GetAbilities() {
			if ability.IsNotAvailable(c) {
				continue
			}
			log.Add(battlelog.LogEntry{
				Round:  round,
				Actor:  c.Name(),
				Action: ability.Name(),
				Target: c.Name(),
				Effect: ability.Use(c),
			})
		}
	}
}
