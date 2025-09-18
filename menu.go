package projetred

import (
	"fmt"
	"time"
)

const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
)

func typeWriter(str string, delay time.Duration) {
	for _, r := range str {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
	fmt.Println()
}

func Menu(c Character) {
	var choice int
	var textDelay = 15 * time.Millisecond
	typeWriter("🏠 Vous êtes dans la maison des Simpson. Que voulez-vous faire ?", textDelay)
	typeWriter("1. 💼 Sortir de la maison (Scénario)", textDelay)
	typeWriter("2. 🎒 Sac à dos et équipements", textDelay)
	typeWriter("3. 📊 Voir vos statistiques", textDelay)
	typeWriter("4. ⚔️ Gérer les équipements", textDelay)
	typeWriter("5. 🎯 Menu des compétences", textDelay)
	typeWriter("6. 🏪 Kwik-E-Mart (Shop + Craft)", textDelay)
	typeWriter("0. 🚪 Rentrer à la maison", textDelay)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		progress := StartHomerScenario(&c)
		ScenarioMenu(&c, &progress)
	case 2:
		typeWriter(AccessInventory(c), textDelay)
		Menu(c)
	case 3:
		typeWriter(DisplayStats(c), textDelay)
		Menu(c)
	case 4:
		EquipmentMenu(&c)
		Menu(c)
	case 5:
		SkillsMenuSimple(&c)
		Menu(c)
	case 6:
		Shopkeeper(&c)
	case 0:
		typeWriter("🏠 Marge vous appelle pour le dîner. À bientôt !", 15*time.Millisecond)
		return
	default:
		typeWriter("❌ Choix invalide.", 15*time.Millisecond)
		Menu(c)
	}
}
