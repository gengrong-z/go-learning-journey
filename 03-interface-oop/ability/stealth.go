package ability

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Stealth struct{}

func (s *Stealth) Damage() int {
	return 10
}

func (s *Stealth) Name() string {
	return "Stealth"
}

func (s *Stealth) Use(user core.Character, target core.Character) string {
	user.GetStatus().MP -= s.Cost()
	effect := target.TakeDamage(s.Damage())
	return user.Name() + " vanishes into the shadows!\n" + effect
}

func (s *Stealth) Cost() int {
	return 100
}

func (s *Stealth) IsNotAvailable(user core.Character) bool {
	if user.GetStatus().MP < s.Cost() {
		fmt.Printf("%s Cost is not enough!", user.Name())
		return true
	}
	return false
}
