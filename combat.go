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
	var textDelay = 15 * time.Millisecond
	baseDamage := c.power
	equipmentBonus := GetTotalEquipmentBonus(c)
	damage := baseDamage + equipmentBonus

	// Vérifier si on a un donut empoisonné
	if contains(c.inventory, "Donut empoisonné") {
		damage += 30
		typeWriter("☠️ Votre attaque est empoisonnée par le donut toxique !", textDelay)
		// Retirer le donut empoisonné après usage
		for i, item := range c.inventory {
			if item == "Donut empoisonné" {
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				break
			}
		}
	}

	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}

	if equipmentBonus > 0 {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts ! (+%d équipement)", c.class, m.name, damage, equipmentBonus), textDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts !", c.class, m.name, damage), textDelay)
	}
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
	typeWriter(fmt.Sprintf("👾 Milhouse (pv restant) :  %d PV !", m.PV), 15*time.Millisecond)
	typeWriter("⚔️ À votre tour ! Choisissez une action :", 15*time.Millisecond)
	typeWriter("1. 💥 Attaquer", 15*time.Millisecond)
	typeWriter("2. 🎯 Utiliser une compétence", 15*time.Millisecond)
	typeWriter("3. 🎒 Fouiller dans votre sac", 15*time.Millisecond)
	typeWriter("4. 🏃 Fuir le combat", 15*time.Millisecond)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		attackMonster(c, m)
		if m.PV <= 0 {
			typeWriter(fmt.Sprintf("🎉 Victoire ! Vous avez vaincu %s !", m.name), 15*time.Millisecond)
			return
		}
		milhousePattern(m, 3)
		traningFight(c, m)
	case 2:
		if UseCombatSkillFromCharacter(c, m) {
			if m.PV <= 0 {
				typeWriter(fmt.Sprintf("🎉 Victoire ! Vous avez vaincu %s !", m.name), 15*time.Millisecond)
				return
			}
			milhousePattern(m, 3)
		}
		traningFight(c, m)
	case 3:
		typeWriter(AccessInventory(*c), 15*time.Millisecond)
		traningFight(c, m)
	case 4:
		typeWriter("🏃💨 Vous fuyez le combat !", 15*time.Millisecond)
		return
	default:
		typeWriter("❌ Choix invalide.", 15*time.Millisecond)
		traningFight(c, m)
	}
}

func milhousePattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("🕶️ Milhouse utilise 'CRISE DE NERFS PARALYSANTE' !", 15*time.Millisecond)
		typeWriter("👦 Milhouse : 'Bart... pourquoi moi toujours ?!'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("📚 Dégâts embarrassants : %d points !", damage), 15*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 10
		typeWriter("🖍️ Milhouse utilise 'CRAYON MAGIQUE' !", 15*time.Millisecond)
		typeWriter("👦 Milhouse : 'Ce crayon m'a été donné par Lisa !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("⚡ Dégâts artistiques : %d points !", damage), 15*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("😨 Milhouse utilise 'PLAINTES DÉSESPÉRÉES' !", 15*time.Millisecond)
		typeWriter("👦 Milhouse : 'Oh non, pas encore !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("😓 Dégâts de désespoir : %d points !", damage), 15*time.Millisecond)
	}
}
