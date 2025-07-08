package ability

import "go-learning-journey/03-interface-oop/core"

type Fireball struct{}

func (f *Fireball) Name() string {
	return "Fireball"
}

func (f *Fireball) Use(user core.Character) string {
	return user.Name() + " hurls a blazing fireball!"
}
