package projetred

import (
	"fmt"
)

type Monster struct {
	name  string
	PVmax int
	PV    int
	power int
}

func InitMonster(name string, PVmax int, power int) Monster {
	return Monster{
		name:  name,
		PVmax: PVmax,
		PV:    PVmax,
		power: power,
	}
}

func goblinPattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter(fmt.Sprintf("Le %s utilise une attaque speciale et inflige %d points de degats!", m.name, damage), textDelay)
	} else {
		damage := m.power
		typeWriter(fmt.Sprintf("Le %s utilise une attaque normale et inflige %d points de degats!", m.name, damage), textDelay)
	}
}

func characterTurn(c *Character, m *Monster, t int) {
	var choice int
	turn := t
	typeWriter("Choisissez une action:", textDelay)
	typeWriter("1. Attaquer", textDelay)
	typeWriter("2. Utiliser un objet", textDelay)
	typeWriter("3. Fuir", textDelay)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		attackMonster(c, m)
		turn++
		goblinPattern(m, turn)
		characterTurn(c, m, turn)
	case 2:
		useItem(c)
		turn++
		characterTurn(c, m, turn)
	default:
		typeWriter("Choix invalide.", textDelay)
		characterTurn(c, m, turn)
	}
}
