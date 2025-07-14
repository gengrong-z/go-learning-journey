package ability

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Fireball struct{}

func (f *Fireball) Name() string {
	return "Fireball"
}

func (f *Fireball) Use(user core.Character) string {
	return user.Name() + " hurls a blazing fireball!"
}

func (f *Fireball) MP() int {
	return 100
}

func (f *Fireball) IsNotAvailable(user core.Character) bool {
	if user.GetStatus().MP < f.MP() {
		fmt.Printf("%s MP is not enough!", user.Name())
		return true
	}
	return false
}
