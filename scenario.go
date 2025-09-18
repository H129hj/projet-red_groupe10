package projetred

import (
	"fmt"
	"time"
)

// Structure pour suivre la progression du scénario
type ScenarioProgress struct {
	Stage          int // 0=début, 1=Ned, 2=Bar Moe, 3=Magasin BD, 4=Parc
	NedCompleted   bool
	MoeCompleted   bool
	ComicCompleted bool
	HasClue1       bool // Indice de Ned
	HasClue2       bool // Indice de Barney
	HasClue3       bool // Indice du Comic Book Guy
}

// Initialise le scénario
func InitScenario() ScenarioProgress {
	return ScenarioProgress{
		Stage:          0,
		NedCompleted:   false,
		MoeCompleted:   false,
		ComicCompleted: false,
		HasClue1:       false,
		HasClue2:       false,
		HasClue3:       false,
	}
}

// Introduction du scénario
func StartHomerScenario(c *Character) ScenarioProgress {
	progress := InitScenario()

	typeWriter("==================================================", 20*time.Millisecond)
	typeWriter("🏠 SPRINGFIELD RPG - À LA RECHERCHE D'HOMER", 40*time.Millisecond)
	typeWriter("==================================================", 20*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👩‍🦱 Marge apparaît, l'air inquiet...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("💬 Marge : 'Oh mon dieu ! Homer n'est pas rentré de la taverne !'", 40*time.Millisecond)
	typeWriter("💬 Marge : 'Il devait juste prendre UNE bière chez Moe...'", 40*time.Millisecond)
	typeWriter("💬 Marge : 'Peux-tu aller demander dans Springfield si quelqu'un sait où il est ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("", 30*time.Millisecond)
	typeWriter("🎯 PREMIER OBJECTIF : Parler aux voisins", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	progress.Stage = 1
	return progress
}

// Menu du scénario principal
func ScenarioMenu(c *Character, progress *ScenarioProgress) {
	var choice int

	typeWriter("🗺️ Où voulez-vous aller ?", 30*time.Millisecond)

	switch progress.Stage {
	case 1:
		typeWriter("1. 🏡 Aller voir Ned Flanders (voisin)", 30*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 30*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 30*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 30*time.Millisecond)
	case 2:
		typeWriter("1. Allez au bar de Moe", 30*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 30*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 30*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 30*time.Millisecond)
	case 3:
		typeWriter("1. 📚 Magasin de BD", 30*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 30*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 30*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 30*time.Millisecond)
	case 4:
		typeWriter("1. 🎡 Parc d'attractions", 30*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 30*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 30*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 30*time.Millisecond)
	}

	typeWriter("0. 🏠 Retourner à la maison", 30*time.Millisecond)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		switch progress.Stage {
		case 1:
			NedFlanders(c, progress)
		case 2:
			MoeBar(c, progress)
		case 3:
			ComicBookStore(c, progress)
		case 4:
			AmusementPark(c, progress)
		}
	case 2:
		typeWriter(AccessInventory(*c), 50*time.Millisecond)
	case 3:
		typeWriter(DisplayStats(*c), 50*time.Millisecond)
	case 4:
		Shopkeeper(c)
	case 0:
		typeWriter("🏠 Vous retournez à la maison...", 50*time.Millisecond)
		Menu(*c)
		return
	default:
		typeWriter("❌ Choix invalide.", 30*time.Millisecond)
		ScenarioMenu(c, progress)
	}
}
