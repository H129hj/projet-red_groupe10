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
	typeWriter("Que souhaitez-vous faire?", 30*time.Millisecond)
	typeWriter("1. Combattre Milhouse", 30*time.Millisecond)
	typeWriter("2. Voir l'inventaire", 30*time.Millisecond)
	typeWriter("3. Voir les statistiques du personnage", 30*time.Millisecond)
	typeWriter("4. Aller chez le marchand", 30*time.Millisecond)
	typeWriter("0. Quitter le jeu", 30*time.Millisecond)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		Milhouse := InitMonster("Milhouse", 100, 20)
		characterTurn(&c1, &Milhouse, 1)
	case 2:
		typeWriter(AccessInventory(c1), 50*time.Millisecond)
		Menu(c1)
	case 3:
		typeWriter(DisplayStats(c1), 50*time.Millisecond)
		Menu(c1)
	case 4:
		Shopkeeper(&c1)
	case 0:
		typeWriter("Au revoir!", 50*time.Millisecond)
		return 
	}
}