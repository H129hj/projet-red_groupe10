package projetred

func Wasted(c *Character) {
	if c.PV <= 0 {
		println("Vous etes mort !")
		c.PV = c.PVmax / 2
		println("Vous avez ete ressuscite avec", c.PV, "PV.")
	}
}
