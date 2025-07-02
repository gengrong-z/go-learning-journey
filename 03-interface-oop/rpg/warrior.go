package rpg

import "strconv"

type Warrior struct {
	HeroName string
	Strength int
	Defense  int
}

func NewWarrior(name string) *Warrior {
	return &Warrior{
		HeroName: name,
		Strength: 0,
		Defense:  0,
	}
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
