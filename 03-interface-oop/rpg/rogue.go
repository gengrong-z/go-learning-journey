package rpg

import (
	"go-learning-journey/03-interface-oop/core"
	"strconv"
)

type Rogue struct {
	HeroName  string
	Mana      int
	Defense   int
	Abilities []core.Ability
}

func (m *Rogue) Name() string {
	return m.HeroName
}

func (m *Rogue) Attack() string {
	return m.Name() + " stabs swiftly!"
}

func (m *Rogue) Defend() string {
	return m.Name() + " conjures a magic barrier!"
}

func (m *Rogue) Status() string {
	return m.Name() + " [Type: Rogue] Mana:" +
		strconv.Itoa(m.Mana) + " DEF:" + strconv.Itoa(m.Defense)
}

func (m *Rogue) CastSpell() string {
	return m.Name() + " casts Arcane Blast!"
}

func (m *Rogue) GetAbilities() []core.Ability {
	return m.Abilities
}
