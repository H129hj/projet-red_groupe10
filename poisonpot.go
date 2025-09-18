package projetred

import (
	"time"
)

// Utiliser un donut empoisonné en combat
func poisonPot(c *Character, m *Monster) {
	if contains(c.inventory, "Donut empoisonné") {
		m.PV -= 20
		typeWriter("☠️ Le donut empoisonné inflige des dégâts supplémentaires !", 15*time.Millisecond)
	}
}

// Vérifier les effets de poison après chaque tour
func CheckPoisonEffects(c *Character, m *Monster) {
	// Si le joueur a utilisé un donut empoisonné, l'ennemi peut subir des dégâts de poison
	if contains(c.inventory, "Donut empoisonné") {
		poisonPot(c, m)
	}
}
