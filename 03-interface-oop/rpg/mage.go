package rpg

import "strconv"

type Mage struct {
	HeroName string
	Mana     int
	Defense  int
}

func NewMage(name string) *Mage {
	return &Mage{
		HeroName: name,
		Mana:     0,
		Defense:  0,
	}
}

func (m *Mage) Name() string {
	return m.HeroName
}

func (m *Mage) Attack() string {
	return m.Name() + " casts a fireball!"
}

func (m *Mage) Defend() string {
	return m.Name() + " conjures a magic barrier!"
}

func (m *Mage) Status() string {
	return m.Name() + " [Type: Mage] Mana:" +
		strconv.Itoa(m.Mana) + " DEF:" + strconv.Itoa(m.Defense)
}

func (m *Mage) CastSpell() string {
	return m.Name() + " casts Arcane Blast!"
}
