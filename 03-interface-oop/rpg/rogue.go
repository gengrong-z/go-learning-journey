package rpg

import "strconv"

type Rogue struct {
	HeroName  string
	Mana      int
	Defense   int
	Abilities []Ability
}

func NewRogue() *Rogue {
	return &Rogue{
		HeroName:  "Hero",
		Mana:      100,
		Defense:   100,
		Abilities: []Ability{&Stealth{}},
	}
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

func (m *Rogue) GetAbilities() []Ability {
	return m.Abilities
}
