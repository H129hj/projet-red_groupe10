package projetred

func Wasted(c *Character) {
	if c.PV <= 0 {
		println("💀 Vous êtes K.O. ! Nelson crie 'HA-HA !' au loin...")
		c.PV = c.PVmax / 2
<<<<<<< HEAD
		println("🏥 Le Dr Hibbert vous soigne avec", c.PV, "PV. 'Ah-heh-heh-heh !'")
=======
		println("Vous avez ete ressuscite avec", c.PV, "PV.")
		Menu(*c)
>>>>>>> 7719d8322be9b313628e15d58f0bfb17f756d403
	}
}
