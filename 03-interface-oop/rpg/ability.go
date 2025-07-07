package rpg

type Ability interface {
	Name() string
	Use(user Character) string
}
