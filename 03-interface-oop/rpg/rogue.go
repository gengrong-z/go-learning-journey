package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Rogue struct {
	HeroName  string
	Mana      int
	Defense   int
	State     core.Status
	Abilities []core.Ability
}

func (r *Rogue) AttackerPower() int {
	return 10
}

func (r *Rogue) Name() string {
	return r.HeroName
}

func (r *Rogue) Attack() string {
	return r.Name() + " stabs swiftly!"
}

func (r *Rogue) Defend() string {
	return r.Name() + " conjures a magic barrier!"
}

//func (m *Rogue) Status() string {
//	return m.Name() + " [Type: Rogue] Mana:" +
//		strconv.Itoa(m.Mana) + " DEF:" + strconv.Itoa(m.Defense)
//}

func (r *Rogue) CastSpell() string {
	return r.Name() + " casts Arcane Blast!"
}

func (r *Rogue) GetAbilities() []core.Ability {
	return r.Abilities
}

func (r *Rogue) TakeDamage(dmg int) string {
	return r.State.TakeDamage(dmg)
}

func (r *Rogue) GetStatus() *core.Status {
	return &r.State
}

func (r *Rogue) Status() string {
	return fmt.Sprintf("%s (HP: %d)", r.HeroName, r.State.HP)
}
