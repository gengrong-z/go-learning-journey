package rpg

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Warrior struct {
	HeroName string
	Strength int
	Defense  int
	State    core.Status
}

func (w *Warrior) AttackerPower() int {
	return w.Strength
}

func (w *Warrior) Name() string {
	return w.HeroName
}

func (w *Warrior) Attack() string {
	return w.Name() + " swings a sword fiercely!"
}

func (w *Warrior) Defend() string {
	return w.Name() + " raises a shield to block!"
}

func (w *Warrior) Status() string {
	return fmt.Sprintf("%s (HP: %d)", w.HeroName, w.State.HP)
}

func (w *Warrior) TakeDamage(dmg int) string {
	return w.State.TakeDamage(dmg)
}

func (w *Warrior) GetStatus() *core.Status {
	return &w.State
}
