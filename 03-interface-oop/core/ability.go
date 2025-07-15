package core

type Ability interface {
	Name() string
	Cost() int
	Use(user Character, target Character) string
	IsNotAvailable(user Character) bool
	Damage() int
	//isAvailable(user Character) bool
}
