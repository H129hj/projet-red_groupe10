package projetred

func Wasted(c *Character) {
	if c.PV <= 0 {
		println("💀 Vous êtes K.O. ! Nelson crie 'HA-HA !' au loin...")
		c.PV = c.PVmax / 2
		println("Vous avez ete ressuscite avec", c.PV, "PV.")
		Menu(*c)
	}
}
