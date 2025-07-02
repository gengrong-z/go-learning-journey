package main

import (
	rpg2 "go-learning-journey/03-interface-oop/rpg"
)

func main() {
	w := rpg2.NewWarrior("Arthas")
	rpg2.PrintStatus(w)
	println(w.Attack())
	println(w.Defend())
}
