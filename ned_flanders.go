package projetred

import (
	"fmt"
	"time"
)

// Interaction avec Ned Flanders
func NedFlanders(c *Character, progress *ScenarioProgress) {
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("🏡 DEVANT LA MAISON DE NED FLANDERS", 50*time.Millisecond)
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👨‍🦳 Vous apercevez Ned Flanders qui arrose ses fleurs parfaitement alignées...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👨‍🦳 Ned : 'Oh ! Bonjour petit voisin ! belle journée, n'est-ce pas ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	var choice int
	typeWriter("💭 Comment voulez-vous aborder Ned ?", 30*time.Millisecond)
	typeWriter("1. 'Ned, as-tu vu Homer aujourd'hui ?'", 30*time.Millisecond)
	typeWriter("2. 'Tu sais s'il est allé au bar de Moe ?'", 30*time.Millisecond)
	typeWriter("3. 'T'es un voisin ringard, Ned !'", 30*time.Millisecond)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		typeWriter("👨‍🦳 Ned : 'Oh oui ! J'ai vu Homer ce matin, diddly-dang !'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Il avait l'air... comment dire... plus louche que d'habitude !'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Il marmonnait quelque chose à propos de \"bière gratuite\" et \"plan secret\"...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer avait un plan secret impliquant de la bière gratuite !", 40*time.Millisecond)
		progress.NedCompleted = true
		progress.Stage = 2

	case 2:
		typeWriter("👨‍🦳 Ned : 'Effectivement ! Je l'ai vu marcher vers le bar de Moe...'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Mais il avait l'air différent, diddly-dong ! Presque... déterminé ?'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Ce qui est surprenant pour Homer, tu en conviendras !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer est allé chez Moe avec une détermination inhabituelle !", 40*time.Millisecond)
		progress.Stage = 2

	case 3:
		typeWriter("👨‍🦳 Ned : 'Oh non ! Cette impolitesse ne peut rester impunie !'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Au nom du Seigneur, je vais te donner une leçon de bonnes manières !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("⚔️ Ned se transforme en CHEVALIER BIBLIQUE !", 50*time.Millisecond)

		nedMonster := InitMonster("Ned Flanders (Mode Biblique)", 80, 25)
		characterTurnNed(c, &nedMonster, 1, progress)
		return

	default:
		typeWriter("❌ Choix invalide.", 30*time.Millisecond)
		NedFlanders(c, progress)
		return
	}

	typeWriter("", 30*time.Millisecond)
	ScenarioMenu(c, progress)
}

// Combat spécialisé contre Ned avec ses attaques spéciales
func characterTurnNed(c *Character, m *Monster, t int, progress *ScenarioProgress) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter("🎉 Victoire ! Ned redevient normal...", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Oh... désolé pour cet éclat ! Diddly-dang, que m'est-il arrivé ?'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Pour me faire pardonner, laisse-moi t'aider...'", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'J'ai vu Homer marcher vers le bar de Moe, l'air louche...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer était louche en allant chez Moe !", 40*time.Millisecond)
		progress.Stage = 2
		ScenarioMenu(c, progress)
	} else {
		typeWriter("⚔️ À votre tour ! Choisissez une action :", 30*time.Millisecond)
		typeWriter("1. 💥 Attaquer", 30*time.Millisecond)
		typeWriter("2. 🎒 Fouiller dans votre sac", 30*time.Millisecond)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			attackMonster(c, m)
			nedPattern(m, turn)
			characterTurnNed(c, m, turn+1, progress)
		case 2:
			typeWriter(AccessInventory(*c), 30*time.Millisecond)
			characterTurnNed(c, m, turn, progress)
		default:
			typeWriter("❌ Choix invalide.", 30*time.Millisecond)
			characterTurnNed(c, m, turn, progress)
		}
	}
}

// Attaques spéciales de Ned
func nedPattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("⛪ Ned utilise 'CITATION BIBLIQUE PARALYSANTE' !", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Tu ne tueras point... mais je peux t'étourdir un peu !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("📖 Dégâts sacrés : %d points !", damage), 30*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 10
		typeWriter("🏏 Ned utilise 'BATTE BÉNITE' !", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Cette batte a été bénie par le révérend Lovejoy !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("⚡ Dégâts divins : %d points !", damage), 30*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("✋ Ned utilise 'SERMON MORALISATEUR' !", 40*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Diddly-dang ! Tu devrais avoir honte !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("😇 Dégâts de culpabilité : %d points !", damage), 30*time.Millisecond)
	}
}
