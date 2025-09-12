package projetred

import "fmt"

type Character struct {
	name string
	class string
	level int
	PVmax int
	PV int
	inventory []string
}

func InitCharacter() Character {
	var name string
	fmt.Print("Entrez le nom du personnage: ")
	fmt.Scan(&name)

	c := Character{
		name:      name,
		class:     "guerrier",
		level:     1,
		PVmax:     100,
		PV:        40,
		inventory: []string{"potion", "potion", "potion"},
	}

	return c
}
