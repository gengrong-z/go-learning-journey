package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
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

func SimulateTurn(characters []core.Character) {
	for _, c := range characters {
		fmt.Println("🎭 " + c.Name())
		fmt.Println("🗡️ " + c.Attack())
		fmt.Println("🛡️ " + c.Defend())

		//if caster, ok := c.(Caster); ok {
		//	fmt.Println("✨ ", caster.CastSpell())
		//}

		UseAbilities(c)
		fmt.Println()
	}
}

func UseAbilities(c core.Character) {
	if caster, ok := c.(interface{ GetAbilities() []core.Ability }); ok {
		for _, ability := range caster.GetAbilities() {
			fmt.Println("✨", ability.Use(c))
		}
	}
}
