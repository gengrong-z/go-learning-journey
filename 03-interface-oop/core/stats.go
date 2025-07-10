package core

import "fmt"

type Status struct {
	HP int
	MP int
}

func (s *Status) TakeDamage(dmg int) string {
	s.HP -= dmg
	if s.HP < 0 {
		s.HP = 0
	}

	return fmt.Sprintf("lost %d HP (now %d)", dmg, s.HP)
}

func (s *Status) IsAlive() bool {
	return s.HP > 0
}
