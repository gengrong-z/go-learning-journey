package rpg

type Fireball struct{}

func (f *Fireball) Name() string {
	return "Fireball"
}

func (f *Fireball) Use(user Character) string {
	return user.Name() + " hurls a blazing fireball!"
}
