package ability

import (
	"go-learning-journey/03-interface-oop/core"
)

type Stealth struct{}

func (s *Stealth) Name() string {
	return "Stealth"
}

func (s *Stealth) Use(user core.Character) string {
	return user.Name() + " vanishes into the shadows!"
}
