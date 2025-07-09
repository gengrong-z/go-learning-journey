package rpg

import (
	"go-learning-journey/03-interface-oop/ability"
	"go-learning-journey/03-interface-oop/core"
)

func NewState() *core.Status {
	return &core.Status{HP: 100, MP: 100}
}

func NewMage(name string) *Mage {
	return &Mage{
		HeroName:  name,
		Mana:      0,
		Defense:   0,
		State:     *NewState(),
		Abilities: []core.Ability{&ability.Fireball{}},
	}
}

func NewRogue(name string) *Rogue {
	return &Rogue{
		HeroName:  name,
		Mana:      100,
		Defense:   100,
		State:     *NewState(),
		Abilities: []core.Ability{&ability.Stealth{}},
	}
}

func NewWarrior(name string) *Warrior {
	return &Warrior{
		HeroName: name,
		Strength: 0,
		Defense:  0,
		State:    *NewState(),
		// Warrior has no special abilities
	}
}
