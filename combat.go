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

func traningFight(c *Character, m *Monster) {
	var choice int
	typeWriter(fmt.Sprintf("👾 Milhouse (pv restant) :  %d PV !", m.PV), 30*time.Millisecond)
	typeWriter("⚔️ À votre tour ! Choisissez une action :", 30*time.Millisecond)
	typeWriter("1. 💥 Attaquer", 30*time.Millisecond)
	typeWriter("2. 🎒 Fouiller dans votre sac", 30*time.Millisecond)
	typeWriter("3. 🏃 Fuir le combat", 30*time.Millisecond)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		attackMonster(c, m)
		if m.PV <= 0 {
			typeWriter(fmt.Sprintf("🎉 Victoire ! Vous avez vaincu %s !", m.name), 40*time.Millisecond)
			return
		}
		milhousePattern(m, 3)
		traningFight(c, m)
	case 2:
		typeWriter(AccessInventory(*c), 30*time.Millisecond)
		traningFight(c, m)
	case 3:
		typeWriter("🏃💨 Vous fuyez le combat !", 40*time.Millisecond)
		return
	default:
		typeWriter("❌ Choix invalide.", 30*time.Millisecond)
		traningFight(c, m)
	}
}

func milhousePattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("🕶️ Milhouse utilise 'CRISE DE NERFS PARALYSANTE' !", 40*time.Millisecond)
		typeWriter("👦 Milhouse : 'Bart... pourquoi moi toujours ?!'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("📚 Dégâts embarrassants : %d points !", damage), 30*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 10
		typeWriter("🖍️ Milhouse utilise 'CRAYON MAGIQUE' !", 40*time.Millisecond)
		typeWriter("👦 Milhouse : 'Ce crayon m'a été donné par Lisa !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("⚡ Dégâts artistiques : %d points !", damage), 30*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("😨 Milhouse utilise 'PLAINTES DÉSESPÉRÉES' !", 40*time.Millisecond)
		typeWriter("👦 Milhouse : 'Oh non, pas encore !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("😓 Dégâts de désespoir : %d points !", damage), 30*time.Millisecond)
	}
}
