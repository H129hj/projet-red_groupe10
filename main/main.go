package main

import (
	"projetred"
)

func main() {
	go projetred.MusiqueJouer()
	c1 := projetred.InitCharacter()
	projetred.ScenarioMenu(&c1, nil)



}