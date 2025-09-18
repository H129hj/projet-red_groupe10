package projetred

func Wasted(c *Character) {
	if c.PV <= 0 {
		println("💀 Vous êtes K.O. ! Nelson crie 'HA-HA !' au loin...")
		c.PV = c.PVmax / 2
<<<<<<< HEAD
		println("🏥 Le Dr Hibbert vous soigne avec", c.PV, "PV. 'Ah-heh-heh-heh !'")
=======
		println("Vous avez ete ressuscite avec", c.PV, "PV.")
>>>>>>> 6c0957eb8c47f03ac1280581fe4bc9ab941c9af4
		Menu(*c)
	}
}
