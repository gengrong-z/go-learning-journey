package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/battlelog"
	"go-learning-journey/03-interface-oop/core"
	"math/rand"
)

func PrintStatus(c core.Character) {
	fmt.Println(c.Status())
}

func PrintRoleDetail(c core.Character) {
	switch v := c.(type) {
	case *Warrior:
		fmt.Println("💪 This is a brave Warrior with strength", v.Strength)
	case *Mage:
		fmt.Println("🧙 This is a wise Mage with mana", v.Mana)
	default:
		fmt.Println("Unknown character type")
	}
}

func SimulateAbilities(characters []core.Character, round int, log *battlelog.BattleLog) {
	for _, c := range characters {
		fmt.Println("🎭 " + c.Name())
		fmt.Println("🗡️ " + c.Attack())
		fmt.Println("🛡️ " + c.Defend())

		//if caster, ok := c.(Caster); ok {
		//	fmt.Println("✨ ", caster.CastSpell())
		//}

		UseAbilities(log, c, round)

		fmt.Println("🗡️ " + c.Attack())
		log.Add(battlelog.LogEntry{
			Round:  round,
			Actor:  c.Name(),
			Action: c.Attack(),
			Target: "Monster",
			Effect: "normal attack",
		})

		fmt.Println()
	}

	log.Print()
}

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

		//fmt.Println("🗡️ " + c.Attack())
		effect := target.TakeDamage(c.AttackerPower())
		//if !target.GetStatus().IsAlive() {
		//	fmt.Printf("  %s is dead\n", c.Name())
		//} else {
		//	fmt.Printf("  %s ---- HP:%d\n", c.Name(), c.GetStatus().HP)
		//}
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
			log.Add(battlelog.LogEntry{
				Round:  round,
				Actor:  c.Name(),
				Action: ability.Name(),
				Target: c.Name(),
				Effect: ability.Use(c),
			})
			fmt.Println("✨", ability.Use(c))
		}
	}
}
