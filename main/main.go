package main

import (
	"projetred"
)

func main(){
	c1 := projetred.InitCharacter()
	c1.PV = 0
	projetred.Wasted(&c1)
	projetred.DisplayInfo(c1)
}
