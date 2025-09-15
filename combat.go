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