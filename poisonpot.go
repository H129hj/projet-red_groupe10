package projetred

import (
	"time"
)


func poisonPot(c *Character, m *Monster) {
	if contains(c.inventory, "Donut empoisonné") {
		m.PV -= 20
		typeWriter("☠️ Le donut empoisonné inflige des dégâts supplémentaires !", 15*time.Millisecond)
	}
}


func CheckPoisonEffects(c *Character, m *Monster) {

	if contains(c.inventory, "Donut empoisonné") {
		poisonPot(c, m)
	}
}