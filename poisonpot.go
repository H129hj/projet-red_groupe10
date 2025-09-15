package projetred

import "time"

func poisonPot(c *Character) {
	println("Vous avez bu une potion empoisonnee ! Vous perdez 20 PV.")
	c.PV -= 20
	if c.PV < 0 {
		c.PV = 0
	}
	time.Sleep(1 * time.Second)
	println("Il vous reste", c.PV, "PV.")
	if isdead(*c) {
		Wasted(c)
	}
}
