package projetred

import (
	"fmt"
	"os"
	"time"
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

	theme := GetCharacterTheme(c.class)

	for {
		var stageChoice int
		fmt.Println()
		MenuHeader("MENU DE NAVIGATION", SystemTheme)
		ThemedTypeWriter("🗺️ Où voulez-vous aller ?", 15*time.Millisecond, theme, "primary")

		switch progressLocal.Stage {
		case 1:
			typeWriter("1. 🏡 Aller voir Ned Flanders (voisin)", 15*time.Millisecond)
			typeWriter("2. 🎒 Regarder dans votre sac à dos", 15*time.Millisecond)
			typeWriter("3. 📊 Voir vos statistiques", 15*time.Millisecond)
			typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", 15*time.Millisecond)
			typeWriter("5. ⚔️ Gérer les équipements", 15*time.Millisecond)
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

		fmt.Println()
		typeWriter("0. 🏠 Retourner à la maison", 15*time.Millisecond)
		ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
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
			// Offer to use a donut magique outside combat
			var invChoice string
			ColoredTypeWriter("➤ Utiliser un donut magique maintenant ? (o/n)", 15*time.Millisecond, BrightYellow)
			fmt.Scan(&invChoice)
			if invChoice == "o" || invChoice == "O" || invChoice == "y" || invChoice == "Y" {
				oldPV := c.PV
				TakePot(c)
				healed := c.PV - oldPV
				if healed > 0 {
					ColoredTypeWriter(fmt.Sprintf("✨ Donut magique utilisé ! +%d PV (%d/%d)", healed, c.PV, c.PVmax), 15*time.Millisecond, BrightGreen)
				} else {
					ColoredTypeWriter("ℹ️ Aucun soin (PV déjà au maximum ou pas de donut).", 15*time.Millisecond, BrightYellow)
				}
			}
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
				traningFight(c, &Monster{name: "Milhouse", PVmax: 150, PV: 150, power: 2})
				continue
			}
		case 0:
			typeWriter("🏠 Vous retournez à la maison...", 15*time.Millisecond)
			os.Exit(0)
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

	fmt.Println()
	MenuHeader("À LA RECHERCHE D'HOMER", BossTheme)

	ColoredTypeWriter("👩‍🦱 Marge apparaît, l'air inquiet...", 15*time.Millisecond, BrightBlue+Bold)
	fmt.Println()

	typeWriter("💬 Marge : Oh mon dieu ! Homer n'est pas rentré de la taverne !", 15*time.Millisecond)
	typeWriter("💬 Marge : Il devait juste prendre UNE bière chez Moe...", 15*time.Millisecond)
	typeWriter("💬 Marge : Peux-tu aller demander dans Springfield si quelqu'un sait où il est ?", 15*time.Millisecond)
	fmt.Println()

	BoxedText("🎯 PREMIER OBJECTIF : Parler aux voisins", SystemTheme)
	fmt.Println()

	progress.Stage = 1
	return progress
}
