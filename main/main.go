package main

import (
	"projetred"
)

func main() {
	go projetred.MusiqueJouer()
	c1 := projetred.InitCharacter()
	projetred.Menu(c1)
	//projetred.playSoundAsync()
	//projetred.MusiqueJouer()
	//projetred.MusiqueArreter()
}
