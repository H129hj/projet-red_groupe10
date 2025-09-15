package main

import (
	"fmt"
	"projetred"
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

<<<<<<< HEAD
func main() {
=======
func main(){
    var choice int
>>>>>>> 68d4e5c84ceab752c276bfdb514094afc53aec20
	c1 := projetred.InitCharacter()
	c1.PV = 0
	projetred.Wasted(&c1)
	typeWriter(projetred.DisplayInfo(c1), 50*time.Millisecond)
	typeWriter(projetred.AccessInventory(c1), 50*time.Millisecond)
<<<<<<< HEAD
	typeWriter(projetred.AccessInventory(c1), 50*time.Millisecond)
=======
    switch choice {
    case 1:
        typeWriter("Vous avez choisi l'option 1", 50*time.Millisecond)
    case 2:
        typeWriter("Vous avez choisi l'option 2", 50*time.Millisecond)
    default:
        typeWriter("Choix invalide", 50*time.Millisecond)
    }
>>>>>>> 68d4e5c84ceab752c276bfdb514094afc53aec20
}
