package projetred

import "fmt"

type Character struct {
	name      string
	class     string
	level     int
	power     int
	PVmax     int
	PV        int
	inventory []string
	gold      int
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func InitCharacter() Character {
	var name string
	var class string
	var c Character
	fmt.Print("Entrez le nom du personnage: ")
	fmt.Scan(&name)
	fmt.Print("Entrez la classe du personnage (guerrier, mage, voleur): ")
	fmt.Scan(&class)
	if class != "guerrier" && class != "mage" && class != "voleur" {
		fmt.Println("Classe invalide, veuillez choisir entre guerrier, mage ou voleur.")
		return InitCharacter()
	} else if class == "mage" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     70,
			PV:        70,
			power:     150,
			inventory: []string{"Baton", "potion", "potion", "potion"},
			gold : 10,
		}
	} else if class == "voleur" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     80,
			PV:        80,
			power:     100,
			inventory: []string{"Dague", "potion", "potion", "potion"},
			gold : 10,
		}
	} else if class == "guerrier" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     100,
			PV:        100,
			power:     80,
			inventory: []string{"Epee", "potion", "potion", "potion"},
			gold : 10,
		}
	}
	return c
}

func DisplayStats(c Character) string {
	texte := fmt.Sprintf("Nom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\nPower: %v", c.name, c.class, c.level, c.PV, c.PVmax, c.power)
	return texte
}

func AccessInventory(c Character) string {
	texte := fmt.Sprintf("Inventaire: %v", c.inventory)
	return texte
}

func takePot(c *Character) []string {
	for i := range c.inventory {
		if c.inventory[i] == "potion" && c.PV < c.PVmax {
			c.PV += 20
			c.inventory = removeIndex(c.inventory, i)
		}
	}
	return c.inventory
}
