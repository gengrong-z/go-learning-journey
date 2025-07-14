package ability

import (
	"fmt"
	"go-learning-journey/03-interface-oop/core"
)

type Stealth struct{}

func (s *Stealth) Name() string {
	return "Stealth"
}

func (s *Stealth) Use(user core.Character) string {
	return user.Name() + " vanishes into the shadows!"
}

func (s *Stealth) MP() int {
	return 100
}

func (s *Stealth) IsNotAvailable(user core.Character) bool {
	if user.GetStatus().MP < s.MP() {
		fmt.Printf("%s MP is not enough!", user.Name())
		return true
	}
	return false
}
