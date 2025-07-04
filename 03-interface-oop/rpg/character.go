package rpg

type Attacker interface {
	Attack() string
}

type Defender interface {
	Defend() string
}

type Caster interface {
	CastSpell() string
}

// Character is all character's behavior interface
type Character interface {
	Name() string
	Status() string
	Attacker
	Defender
}
