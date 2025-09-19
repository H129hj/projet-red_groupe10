package projetred

import (
	"fmt"
	"time"
)

type ScenarioProgress struct {
	Stage          int
	NedCompleted   bool
	MoeCompleted   bool
	ComicCompleted bool
	HasClue1       bool
	HasClue2       bool
	HasClue3       bool
}

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

func StartHomerScenario(c *Character) ScenarioProgress {
	progress := InitScenario()

	typeWriter("==================================================", 15*time.Millisecond)
	typeWriter("🏠 SPRINGFIELD RPG - À LA RECHERCHE D'HOMER", 15*time.Millisecond)
	typeWriter("==================================================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("👩‍🦱 Marge apparaît, l'air inquiet...", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("💬 Marge : 'Oh mon dieu ! Homer n'est pas rentré de la taverne !'", 15*time.Millisecond)
	typeWriter("💬 Marge : 'Il devait juste prendre UNE bière chez Moe...'", 15*time.Millisecond)
	typeWriter("💬 Marge : 'Peux-tu aller demander dans Springfield si quelqu'un sait où il est ?'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("", 15*time.Millisecond)
	typeWriter("🎯 PREMIER OBJECTIF : Parler aux voisins", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	progress.Stage = 1
	return progress
}

func ScenarioMenu(c *Character, progress *ScenarioProgress) {
	var choice int

	typeWriter("🗺️ Où voulez-vous aller ?", 15*time.Millisecond)

	switch progress.Stage {
	case 1:
		typeWriter("1. 🏡 Aller voir Ned Flanders (voisin)", 15*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
		typeWriter("5. 🥊\u200b Casse la gueule à Milhouse pour passer le temps", 15*time.Millisecond)
	case 2:
		typeWriter("1. 🍻 Allez au bar de Moe", 15*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
	case 3:
		typeWriter("1. 📚 Magasin de BD", 15*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
	case 4:
		typeWriter("1. 🎡 Parc d'attractions", 15*time.Millisecond)
		typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
		typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
		typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
	}

	typeWriter("0. 🏠 Retourner à la maison", 15*time.Millisecond)
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
		typeWriter(AccessInventory(*c), 15*time.Millisecond)
	case 3:
		typeWriter(DisplayStats(*c), 15*time.Millisecond)
	case 4:
		Shopkeeper(c)
	case 5:
		if progress.Stage == 1 {
			typeWriter("👦 Milhouse : 'Hé ! Tu veux t'entraîner au combat ?'", 15*time.Millisecond)
			typeWriter("👦 Milhouse : 'Allez, montre-moi ce que tu sais faire !'", 15*time.Millisecond)
			milhouse := InitMonster("Milhouse Van Houten", 600, 12)
			CombatInterface(c, &milhouse, 1, milhousePatternInstant)
			ScenarioMenu(c, progress)
		}
	case 0:
		typeWriter("🏠 Vous retournez à la maison...", 15*time.Millisecond)
		Menu(*c)
		return
	default:
		typeWriter("❌ Choix invalide.", 15*time.Millisecond)
		ScenarioMenu(c, progress)
	}
}
