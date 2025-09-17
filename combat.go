package projetred

import (
	"fmt"
	"time"
)

type Monster struct {
	name      string
	PVmax     int
	PV        int
	power     int
	textDelay int
}

func attackMonster(c *Character, m *Monster) {
	damage := c.power
	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}
	typeWriter(fmt.Sprintf("%s attaque le %s et inflige %d points de degats!", c.name, m.name, damage), time.Duration(m.textDelay)*time.Millisecond)
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
		typeWriter(fmt.Sprintf("Le %s utilise une attaque speciale et inflige %d points de degats!", m.name, damage), time.Duration(m.textDelay)*time.Millisecond)
	} else {
		damage := m.power
		typeWriter(fmt.Sprintf("Le %s utilise une attaque normale et inflige %d points de degats!", m.name, damage), time.Duration(m.textDelay)*time.Millisecond)
	}
}

func characterTurn(c *Character, m *Monster, t int) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter(fmt.Sprintf("Vous avez vaincu le %s!", m.name), time.Duration(m.textDelay)*time.Millisecond)
		Menu(*c)
	} else {
		typeWriter("Choisissez une action:", time.Duration(m.textDelay)*time.Millisecond)
	typeWriter("1. Attaquer", time.Duration(m.textDelay)*time.Millisecond)
	typeWriter("2. Acceder à l'inventaire", time.Duration(m.textDelay)*time.Millisecond)
	typeWriter("3. Fuir", time.Duration(m.textDelay)*time.Millisecond)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		attackMonster(c, m)
		goblinPattern(m, turn)
		characterTurn(c, m, turn)
	case 2:
		AccessInventory(*c)
		characterTurn(c, m, turn)
	case 3:
		typeWriter("Vous avez fui le combat!", time.Duration(m.textDelay)*time.Millisecond)
		Menu(*c)
	default:
		typeWriter("Choix invalide.", time.Duration(m.textDelay)*time.Millisecond)
		characterTurn(c, m, turn)
	}
}
}