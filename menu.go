package projetred

import "fmt"
import "time"

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
	fmt.Println("Que souhaitez-vous faire?")
	fmt.Println("1. Aller chez le marchand")
	fmt.Println("2. Voir l'inventaire")
	fmt.Println("3. Voir les statistiques du personnage")
	fmt.Println("4. Quitter le jeu")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		Shopkeeper(&c1)
	case 2:
		typeWriter(AccessInventory(c1), 50*time.Millisecond)
		Menu(c1)
	case 3:
		typeWriter(DisplayStats(c1), 50*time.Millisecond)
		Menu(c1)
	case 4:
		typeWriter("Au revoir!", 50*time.Millisecond)
		return 
	}
}