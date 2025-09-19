package projetred

import (
	"fmt"
	"time"
)

func CharacterCreationPersonalized() Character {
	go MusiqueJouer()
	var name string
	var class string
	var c Character
	fmt.Print("🏠 Entrez le nom de votre enfant Simpson: ")
	fmt.Scan(&name)
	fmt.Print("🎭 Choisissez le style (bart, lisa, maggie): ")
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&class)
	if class != "bart" && class != "lisa" && class != "maggie" {
		fmt.Println("❌ Choix invalide ! Vous devez choisir entre Bart, Lisa ou Maggie.")
		return CharacterCreationPersonalized()
	} else if class == "lisa" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     70,
			PV:        70,
			power:     150,
			inventory: []string{"Saxophone de Lisa", "donut magique", "donut magique", "donut magique"},
			gold:      10,
		}
	} else if class == "bart" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     80,
			PV:        80,
			power:     100,
			inventory: []string{"Lance-pierre de Bart", "donut magique", "donut magique", "donut magique"},
			gold:      10,
		}
	} else if class == "maggie" {
		c = Character{
			name:      name,
			class:     class,
			level:     1,
			PVmax:     100,
			PV:        100,
			power:     80,
			inventory: []string{"Biberon de Maggie", "donut magique", "donut magique", "donut magique"},
			gold:      10,
		}
	}
	fmt.Println("🎉 Enfant Simpson créé avec succès !")
	fmt.Println("👤 Nom:", c.name)
	fmt.Println("🎭 Style:", c.class)
	fmt.Println("🏆 Niveau:", c.level)
	fmt.Println("❤️ PV:", c.PV, "/", c.PVmax)
	fmt.Println("💪 Puissance:", c.power)
	fmt.Println("🎒 Inventaire:", c.inventory)
	fmt.Println("💰 Argent de poche:", c.gold, "dollars")
	return c
}

func DisplayInfo(c Character) string {
	texte := fmt.Sprintf("Nom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\nPower: %v", c.name, c.class, c.level, c.PV, c.PVmax, c.power)
	return texte
}
