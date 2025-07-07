package rpg

type Stealth struct{}

func (s *Stealth) Name() string {
	return "Stealth"
}

func (s *Stealth) Use(user Character) string {
	return user.Name() + " vanishes into the shadows!"
}
