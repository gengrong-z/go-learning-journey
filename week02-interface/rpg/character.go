package rpg

// Character is all character's behavior interface
type Character interface {
	Name() string
	Attack() string
	Defend() string
	Status() string
}
