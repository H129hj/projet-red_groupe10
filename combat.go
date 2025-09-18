package projetred

import (
	"fmt"
	"time"
)

type Monster struct {
	name  string
	PVmax int
	PV    int
	power int
}

func attackMonster(c *Character, m *Monster) {
	var textDelay = 20 * time.Millisecond
	damage := c.power
	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}
	typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts !", c.name, m.name, damage), textDelay)
}

func InitMonster(name string, PVmax int, power int) Monster {
	return Monster{
		name:  name,
		PVmax: PVmax,
		PV:    PVmax,
		power: power,
	}
}

