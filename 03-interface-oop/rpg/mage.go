package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Mage struct {
	HeroName  string
	Mana      int
	Defense   int
	State     core.Status
	Abilities []core.Ability
}

func (m *Mage) AttackerPower() int {
	return m.Mana
}

func (m *Mage) GetStatus() *core.Status {
	return &m.State
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

//func (m *Mage) Status() string {
//	return m.Name() + " [Type: Mage] Mana:" +
//		strconv.Itoa(m.Mana) + " DEF:" + strconv.Itoa(m.Defense)
//}

func (m *Mage) CastSpell() string {
	return m.Name() + " casts Arcane Blast!"
}

func (m *Mage) GetAbilities() []core.Ability {
	return m.Abilities
}

func (m *Mage) TakeDamage(dmg int) string {
	return m.State.TakeDamage(dmg)
}

func (m *Mage) Status() string {
	return fmt.Sprintf("%s (HP: %d)", m.HeroName, m.State.HP)
}
