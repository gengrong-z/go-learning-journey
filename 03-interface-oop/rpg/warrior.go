package rpg

import (
	"go-learning-journey/03-interface-oop/core"
	"strconv"
)

type Warrior struct {
	HeroName string
	Strength int
	Defense  int
	State    core.Status
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
	return w.Name() + " [Type: Warrior] STR:" +
		strconv.Itoa(w.Strength) + " DEF:" + strconv.Itoa(w.Defense)
}
