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
	var choice int
	typeWriter("🏠 Vous êtes dans la maison des Simpson. Que voulez-vous faire ?", 30*time.Millisecond)
	typeWriter("1. 🎒 Regarder dans votre sac à dos", 30*time.Millisecond)
	typeWriter("2. 📊 Voir vos statistiques", 30*time.Millisecond)
	typeWriter("3. 🏪 Aller chez Apu au Kwik-E-Mart", 30*time.Millisecond)
	typeWriter("0. 🚪 Rentrer à la maison", 30*time.Millisecond)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		typeWriter(AccessInventory(c1), 50*time.Millisecond)
		Menu(c1)
	case 2:
		typeWriter(DisplayStats(c1), 50*time.Millisecond)
		Menu(c1)
	case 3:
		Shopkeeper(&c1)
	case 0:
		typeWriter("🏠 Marge vous appelle pour le dîner. À bientôt !", 50*time.Millisecond)
		return
	}
}
