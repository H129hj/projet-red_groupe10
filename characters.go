package projetred

import (
	"fmt"
	"slices"
	"strings"
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

	MenuHeader("SÉLECTION DE PERSONNAGE", SystemTheme)
	ThemedTypeWriter("🏠 Bienvenue dans Springfield ! Choisissez votre héros Simpson :", textDelay, SystemTheme, "primary")
	fmt.Println()
	ThemedTypeWriter("🎭 bart  - Le garnement rebelle", textDelay, BartTheme, "primary")
	ThemedTypeWriter("🎓 lisa  - Le génie de la famille", textDelay, LisaTheme, "primary")
	ThemedTypeWriter("👶 maggie - La petite mystérieuse", textDelay, MaggieTheme, "primary")
	fmt.Println()
	ColoredTypeWriter("➤ Votre choix : ", textDelay, BrightCyan+Bold)
	fmt.Scan(&class)

	if class != "bart" && class != "lisa" && class != "maggie" {
		ColoredTypeWriter("❌ Choix invalide ! Vous devez choisir entre Bart, Lisa ou Maggie.", textDelay, BrightRed+Bold)
		return InitCharacter()
	} else if class == "lisa" {
		c = Character{
			class:             class,
			level:             1,
			PVmax:             130,
			PV:                130,
			power:             80,
			inventory:         []string{"Saxophone de Lisa", "donut magique", "donut magique", "donut magique"},
			gold:              100,
			skills:            []string{"Solo de jazz envoûtant", "Leçon de morale dévastatrice", "Méditation bouddhiste"},
			equipement:        map[string]int{"Robe de première de classe": 10, "Serre-tête": 5, "Collier de perles": 7},
			extendedInventory: 0,
		}
	} else if class == "bart" {
		c = Character{
			class:             class,
			level:             1,
			PVmax:             120,
			PV:                120,
			power:             70,
			inventory:         []string{"Lance-pierre de Bart", "donut magique", "donut magique", "donut magique"},
			gold:              100,
			skills:            []string{"Coup de fronde vicieux", "Blague empoisonnée", "Échappée en skateboard"},
			equipement:        map[string]int{"T-shirt rouge": 10, "Short bleu": 5, "Chaussures de sport": 7},
			extendedInventory: 0,
		}
	} else if class == "maggie" {
		c = Character{
			class:             class,
			level:             1,
			PVmax:             110,
			PV:                110,
			power:             60,
			inventory:         []string{"Biberon de Maggie", "donut magique", "donut magique", "donut magique"},
			gold:              100,
			skills:            []string{"Regard hypnotique", "Cri strident", "Attaque surprise du berceau"},
			equipement:        map[string]int{"Grenouillère bleue": 10, "Nœud rose": 5, "Tétine magique": 7},
			extendedInventory: 0,
		}
	}

	theme := GetCharacterTheme(c.class)
	fmt.Println()
	ThemedTypeWriter(fmt.Sprintf("🎉 Personnage %s créé avec succès !", strings.ToUpper(c.class)), textDelay, theme, "primary")

	switch c.class {
	case "lisa":
		ThemedTypeWriter("🎷 Lisa Simpson - L'intellectuelle de Springfield", textDelay, theme, "text")
		ThemedTypeWriter("📚 Spécialité : Intelligence supérieure et attaques sonores", textDelay, theme, "secondary")
	case "bart":
		ThemedTypeWriter("🛹 Bart Simpson - Le roi des bêtises", textDelay, theme, "text")
		ThemedTypeWriter("🎯 Spécialité : Attaques sournoises et évasion", textDelay, theme, "secondary")
	case "maggie":
		ThemedTypeWriter("👶 Maggie Simpson - Le mystère en couches", textDelay, theme, "text")
		ThemedTypeWriter("🍼 Spécialité : Attaques surprises et résistance", textDelay, theme, "secondary")
	}

	fmt.Println()
	return c
}

func DisplayStats(c Character) string {
	theme := GetCharacterTheme(c.class)
	equipmentBonus := GetTotalEquipmentBonus(&c)

	var texte string
	texte += ColorizeThemed(fmt.Sprintf("👤 Personnage: %s", strings.ToUpper(c.class)), theme, "primary") + "\n"
	texte += ColorizeThemed(fmt.Sprintf("🏆 Niveau: %d", c.level), theme, "accent") + "\n"
	texte += HealthBar(c.PV, c.PVmax, 20) + "\n"
	texte += ColorizeThemed(fmt.Sprintf("💪 Puissance: %d", c.power), theme, "text")
	if equipmentBonus > 0 {
		texte += Colorize(fmt.Sprintf(" (+%d équipement)", equipmentBonus), BrightGreen)
	}
	texte += "\n" + CurrencyDisplay(c.gold) + "\n"
	texte += ProgressBar(len(c.inventory), 10+c.extendedInventory, 15, theme) + " 🎒 Inventaire\n"
	texte += ColorizeThemed(fmt.Sprintf("⚔️ Équipements: %d objets portés", len(c.equipement)), theme, "secondary")

	return texte
}

func AccessInventory(c Character) string {
	theme := GetCharacterTheme(c.class)
	var texte string

	texte += ColorizeThemed("🎒 SAC À DOS", theme, "primary") + "\n"
	texte += ColorizeThemed("============", theme, "accent") + "\n"
	if len(c.inventory) == 0 {
		texte += Colorize("❌ Votre sac à dos est vide.", BrightRed) + "\n"
	} else {
		for i, item := range c.inventory {
			rarity := "common"
			if strings.Contains(item, "magique") || strings.Contains(item, "Spirituel") || strings.Contains(item, "Super") {
				rarity = "legendary"
			} else if strings.Contains(item, "rare") || strings.Contains(item, "Itchy") {
				rarity = "rare"
			}
			texte += fmt.Sprintf("%s. %s\n", ColorizeThemed(fmt.Sprintf("%d", i+1), theme, "accent"), ItemDisplay(item, rarity))
		}
	}

	texte += "\n" + ColorizeThemed("⚔️ ÉQUIPEMENTS PORTÉS", theme, "primary") + "\n"
	texte += ColorizeThemed("======================", theme, "accent") + "\n"
	if len(c.equipement) == 0 {
		texte += Colorize("❌ Aucun équipement porté.", BrightRed) + "\n"
	} else {
		for equipName, value := range c.equipement {
			powerDisplay := Colorize(fmt.Sprintf("(+%d)", value), BrightGreen+Bold)
			texte += fmt.Sprintf("✅ %s %s\n", ItemDisplay(equipName, "rare"), powerDisplay)
		}
	}
	return texte
}

func TakePot(c *Character) []string {
	for i := range c.inventory {
		if c.inventory[i] == "donut magique" && c.PV < c.PVmax {
			c.PV += 50
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
