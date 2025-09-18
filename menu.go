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

func Menu(c1 Character) {
	checkpoint := 0
	var choice int
<<<<<<< HEAD
	typeWriter("🏠 Vous êtes dans la maison des Simpson. Que voulez-vous faire ?", 30*time.Millisecond)
	typeWriter("1. 🎬 NOUVEAU : À la recherche d'Homer (Scénario principal)", 30*time.Millisecond)
	typeWriter("2. 🥊 Combat d'entraînement contre Ralph Wiggum", 30*time.Millisecond)
	typeWriter("3. 🎒 Regarder dans votre sac à dos", 30*time.Millisecond)
	typeWriter("4. 📊 Voir vos statistiques", 30*time.Millisecond)
	typeWriter("5. 🏪 Aller chez Apu au Kwik-E-Mart", 30*time.Millisecond)
	typeWriter("0. 🚪 Rentrer à la maison", 30*time.Millisecond)
=======
	var textDelay = 20 * time.Millisecond
	if checkpoint == 0 {
		typeWriter("🏠 Vous êtes dans la maison des Simpson. Que voulez-vous faire ?", textDelay)
		typeWriter("1. 💼 Sortir de la maison", textDelay)
	}
	typeWriter("2. 🎒 Regarder dans votre sac à dos", textDelay)
	typeWriter("3. 📊 Voir vos statistiques", textDelay)
	typeWriter("4. 🏪 Aller chez Apu au Kwik-E-Mart", textDelay)
	typeWriter("0. 🚪 Rentrer à la maison", textDelay)
>>>>>>> 6c0957eb8c47f03ac1280581fe4bc9ab941c9af4
	fmt.Scan(&choice)
	switch choice {
	case 1:
		// Nouveau scénario principal
		progress := StartHomerScenario(&c1)
		ScenarioMenu(&c1, &progress)
	case 2:
		// Combat d'entraînement
		Ralph := InitMonster("Ralph Wiggum", 100, 20)
		characterTurn(&c1, &Ralph, 1)
	case 3:
		typeWriter(AccessInventory(c1), 50*time.Millisecond)
		Menu(c1)
	case 4:
		typeWriter(DisplayStats(c1), 50*time.Millisecond)
		Menu(c1)
	case 5:
		Shopkeeper(&c1)
	case 0:
		typeWriter("🏠 Marge vous appelle pour le dîner. À bientôt !", 50*time.Millisecond)
		return
	default:
		typeWriter("❌ Choix invalide.", 30*time.Millisecond)
		Menu(c1)
	}
}
