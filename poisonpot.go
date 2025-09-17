package projetred

func poisonPot(c *Character, m *Monster) {
	if contains(c.inventory, "donut empoisonné") {
		m.PV -= 20
	}
}
