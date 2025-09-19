package projetred

import (
	"fmt"
	"time"
)

const (
	Reset     = "\u001b[0m"
	Red       = "\u001b[31m"
	Green     = "\u001b[32m"
	Yellow    = "\u001b[33m"
	Blue      = "\u001b[34m"
	Magenta   = "\u001b[35m"
	Cyan      = "\u001b[36m"
	Bold      = "\u001b[1m"
	Underline = "\u001b[4m"
)

func typeWriter(str string, delay time.Duration) {
	for _, r := range str {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
	fmt.Println()
}

func ScenarioMenu(c *Character, progress *ScenarioProgress) {
	var progressLocal ScenarioProgress
	if progress != nil {
		progressLocal = *progress
	} else {
		progressLocal = StartHomerScenario(c)
	}

	for {
		var stageChoice int
		typeWriter("🗺️  Où voulez-vous aller ?", 15*time.Millisecond)

		switch progressLocal.Stage {
		case 1:
			typeWriter("1. 🏡 Aller voir Ned Flanders (voisin)", 15*time.Millisecond)
			typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
			typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
			typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
			typeWriter("5. ⚔️  Gérer les équipements", 15*time.Millisecond)
			typeWriter("6. 🥊 Casse la gueule à Milhouse pour passer le temps (entrainement)", 15*time.Millisecond)
		case 2:
			typeWriter("1. 🍻 Allez au bar de Moe", 15*time.Millisecond)
			typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
			typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
			typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
			typeWriter("5. ⚔️ Gérer les équipements", 15*time.Millisecond)
		case 3:
			typeWriter("1. 📚 Magasin de BD", 15*time.Millisecond)
			typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
			typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
			typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
			typeWriter("5. ⚔️ Gérer les équipements", 15*time.Millisecond)
		case 4:
			typeWriter("1. 🎡 Parc d'attractions", 15*time.Millisecond)
			typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
			typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
			typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
			typeWriter("5. ⚔️ Gérer les équipements", 15*time.Millisecond)
		}

		typeWriter("0. 🏠 Retourner à la maison", 15*time.Millisecond)
		fmt.Scan(&stageChoice)

		switch stageChoice {
		case 1:
			switch progressLocal.Stage {
			case 1:
				NedFlanders(c, &progressLocal)
			case 2:
				MoeBar(c, &progressLocal)
			case 3:
				ComicBookStore(c, &progressLocal)
			case 4:
				AmusementPark(c, &progressLocal)
			}
		case 2:
			typeWriter(AccessInventory(*c), 15*time.Millisecond)
			continue
		case 3:
			typeWriter(DisplayStats(*c), 15*time.Millisecond)
			continue
		case 4:
			Shopkeeper(c)
			continue
		case 5:
			EquipmentMenu(c)
			continue
		case 6:
			if progressLocal.Stage == 1 {
				traningFight(c, &Monster{name: "Milhouse", PVmax: 1000000, PV: 1000000, power: 2})
				continue
			}
		case 0:
			typeWriter("🏠 Vous retournez à la maison...", 15*time.Millisecond)
			return
		default:
			typeWriter("❌ Choix invalide.", 15*time.Millisecond)
			continue
		}
	}
}

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
