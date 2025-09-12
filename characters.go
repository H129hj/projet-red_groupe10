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
		}
	}
	return c
}
