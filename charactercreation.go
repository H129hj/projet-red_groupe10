package projetred

import "fmt"

func CharacterCreationPersonalized() Character {
	var name string
	var class string
	var c Character
	fmt.Print("Entrez le nom du personnage: ")
	fmt.Scan(&name)
	fmt.Print("Entrez la classe du personnage (guerrier, mage, voleur): ")
	fmt.Scan(&class)
	if class != "guerrier" && class != "mage" && class != "voleur" {
		fmt.Println("Classe invalide, veuillez choisir entre guerrier, mage ou voleur.")
		return CharacterCreationPersonalized()
	} else if class == "mage" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     70,
			PV:        70,
			power:     150,
			inventory: []string{"Baton", "potion", "potion", "potion"},
			gold:      10,
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
			gold:      10,
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
			gold:      10,
		}
	}
	fmt.Println("Personnage cree avec succes!")
	fmt.Println("Nom:", c.name)
	fmt.Println("Classe:", c.class)
	fmt.Println("Niveau:", c.level)
	fmt.Println("PV:", c.PV, "/", c.PVmax)
	fmt.Println("Puissance:", c.power)
	fmt.Println("Inventaire:", c.inventory)
	fmt.Println("Or:", c.gold, "pieces d'or")
	return c
}

func DisplayInfo(c Character) string {
	texte := fmt.Sprintf("Nom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\nPower: %v", c.name, c.class, c.level, c.PV, c.PVmax, c.power)
	return texte
}
