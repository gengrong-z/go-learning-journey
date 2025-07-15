package ability

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Fireball struct{}

func (f *Fireball) Damage() int {
	return 30
}

func (f *Fireball) Name() string {
	return "Fireball"
}

func (f *Fireball) Use(user core.Character, target core.Character) string {
	user.GetStatus().MP -= f.Cost()
	effect := target.TakeDamage(f.Damage())
	return user.Name() + " hurls a blazing fireball!\n" + effect
}

func (f *Fireball) Cost() int {
	return 100
}

func (f *Fireball) IsNotAvailable(user core.Character) bool {
	if user.GetStatus().MP < f.Cost() {
		fmt.Printf("%s Cost is not enough!", user.Name())
		return true
	}
	return false
}
