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

func goblinPattern(m *Monster, turn int) {
	var textDelay = 20 * time.Millisecond
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter(fmt.Sprintf("🌟 %s utilise sa technique spéciale et inflige %d points de dégâts !", m.name, damage), textDelay)
	} else {
		damage := m.power
		typeWriter(fmt.Sprintf("👊 %s lance une attaque basique et inflige %d points de dégâts !", m.name, damage), textDelay)
	}
}

func characterTurn(c *Character, m *Monster, t int) {
	var choice int
	var textDelay = 20 * time.Millisecond
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter(fmt.Sprintf("🎉 Victoire ! Vous avez battu %s !", m.name), textDelay)
		Menu(*c)
	} else {
		typeWriter("⚔️ À votre tour ! Choisissez une action :", textDelay)
		typeWriter("1. 💥 Attaquer", textDelay)
		typeWriter("2. 🎒 Fouiller dans votre sac", textDelay)
		typeWriter("3. 🏃 Fuir le combat", textDelay)
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
			typeWriter("🏃💨 Vous fuyez le combat comme Milhouse devant Nelson !", textDelay)
			Menu(*c)
		default:
			typeWriter("❌ Choix invalide.", textDelay)
			characterTurn(c, m, turn)
		}
	}
}
