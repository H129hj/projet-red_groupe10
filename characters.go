package projetred

import (
	"fmt"
	"slices"
	"time"
)

type Character struct {
	name              string
	class             string
	level             int
	power             int
	PVmax             int
	PV                int
	inventory         []string
	gold              int
	skills            []string
	equipement        map[string]int
	extendedInventory int
}

func InitCharacter() Character {
	var class string
	var c Character
	var textDelay = 15 * time.Millisecond

	typeWriter("🏠 Bienvenue dans Springfield ! Choisissez votre personnage (bart, lisa, maggie): ", textDelay)
	fmt.Scan(&class)

	if class != "bart" && class != "lisa" && class != "maggie" {
		typeWriter("❌ Choix invalide ! Vous devez choisir entre Bart, Lisa ou Maggie.", textDelay)
		return InitCharacter()
	} else if class == "lisa" {
		c = Character{
			class:             class,
			level:             1,
			PVmax:             70,
			PV:                70,
			power:             150,
			inventory:         []string{"Saxophone de Lisa", "donut magique"},
			gold:              100,
			skills:            []string{"Solo de jazz envoûtant", "Leçon de morale dévastatrice", "Méditation bouddhiste"},
			equipement:        map[string]int{"Robe de première de classe": 10, "Serre-tête": 5, "Collier de perles": 7},
			extendedInventory: 0,
		}
	} else if class == "bart" {
		c = Character{
			class:             class,
			level:             1,
			PVmax:             80,
			PV:                80,
			power:             100,
			inventory:         []string{"Lance-pierre de Bart", "donut magique"},
			gold:              100,
			skills:            []string{"Coup de fronde vicieux", "Blague empoisonnée", "Échappée en skateboard"},
			equipement:        map[string]int{"T-shirt rouge": 10, "Short bleu": 5, "Chaussures de sport": 7},
			extendedInventory: 0,
		}
	} else if class == "maggie" {
		c = Character{
			class:             class,
			level:             1,
			PVmax:             100,
			PV:                100,
			power:             80,
			inventory:         []string{"Biberon de Maggie", "donut magique"},
			gold:              100,
			skills:            []string{"Regard hypnotique", "Cri strident", "Attaque surprise du berceau"},
			equipement:        map[string]int{"Grenouillère bleue": 10, "Nœud rose": 5, "Tétine magique": 7},
			extendedInventory: 0,
		}
	}
	return c
}

func DisplayStats(c Character) string {
	equipmentBonus := GetTotalEquipmentBonus(&c)
	texte := fmt.Sprintf("👤 Nom: %s\n🏆 Niveau: %d\n❤️ PV: %d/%d\n💪 Power: %d", c.class, c.level, c.PV, c.PVmax, c.power)
	if equipmentBonus > 0 {
		texte += fmt.Sprintf(" (+%d'équipement)", equipmentBonus)
	}
	texte += fmt.Sprintf("\n💰 Or: %d dollars", c.gold)
	texte += fmt.Sprintf("\n🎒 Objets: %d/%d", len(c.inventory), 10+c.extendedInventory)
	texte += fmt.Sprintf("\n⚔️ Équipements: %d", len(c.equipement))
	return texte
}

func AccessInventory(c Character) string {
	texte := "🎒 SAC À DOS :\n"
	texte += "============\n"
	if len(c.inventory) == 0 {
		texte += "❌ Votre sac à dos est vide.\n"
	} else {
		for i, item := range c.inventory {
			texte += fmt.Sprintf("%d. %s\n", i+1, item)
		}
	}
	texte += "\n⚔️ ÉQUIPEMENTS PORTÉS :\n"
	texte += "======================\n"
	if len(c.equipement) == 0 {
		texte += "❌ Aucun équipement porté.\n"
	} else {
		for equipName, value := range c.equipement {
			texte += fmt.Sprintf("✅ %s (valeur: %d)\n", equipName, value)
		}
	}
	return texte
}

func TakePot(c *Character) []string {
	for i := range c.inventory {
		if c.inventory[i] == "donut magique" && c.PV < c.PVmax {
			c.PV += 20
			c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
			if c.PV > c.PVmax {
				c.PV = c.PVmax
			}
			break
		}
	}
	return c.inventory
}

func limitedInventory(c *Character) bool {
	if len(c.inventory) >= 10+c.extendedInventory {
		typeWriter("🍩 Vos poches sont pleines de donuts ! Vous ne pouvez pas porter plus d'objets.", 15*time.Millisecond)
		return false
	}
	return true
}

func contains(slice []string, item string) bool {
	return slices.Contains(slice, item)
}


func AddIngredient(c *Character, ingredient string, source string) {
	if !limitedInventory(c) {
		return
	}

	c.inventory = append(c.inventory, ingredient)
	textDelay := 15 * time.Millisecond
	typeWriter(fmt.Sprintf("🎁 Vous trouvez : %s", ingredient), textDelay)
	typeWriter(fmt.Sprintf("📦 Trouvé près de %s", source), textDelay)
}