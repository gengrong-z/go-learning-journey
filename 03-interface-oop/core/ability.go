package core

type Ability interface {
	Name() string
	MP() int
	Use(user Character) string
	IsNotAvailable(user Character) bool
	//isAvailable(user Character) bool
}
