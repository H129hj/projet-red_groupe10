package projetred

func poisonPot(c *Character, m *Monster) {
	if contains(c.inventory, "potion empoisonnee") {
		m.PV -= 20
	}
}
