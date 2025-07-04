package rpg

import "fmt"

func PrintStatus(c Character) {
	fmt.Println(c.Status())
}

func PrintRoleDetail(c Character) {
	switch v := c.(type) {
	case *Warrior:
		fmt.Println("💪 This is a brave Warrior with strength", v.Strength)
	case *Mage:
		fmt.Println("🧙 This is a wise Mage with mana", v.Mana)
	default:
		fmt.Println("Unknown character type")
	}
}

func SimulateTurn(characters []Character) {
	for _, c := range characters {
		fmt.Println("🎭 " + c.Name())
		fmt.Println("🗡️  " + c.Attack())
		fmt.Println("🛡️  " + c.Defend())

		if caster, ok := c.(Caster); ok {
			fmt.Println("✨", caster.CastSpell())
		}

		fmt.Println()
	}
}
