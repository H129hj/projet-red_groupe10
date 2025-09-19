package projetred

import (
	"fmt"
	"time"
)


func NedFlanders(c *Character, progress *ScenarioProgress) {
	typeWriter("==================================================", 15*time.Millisecond)
	typeWriter("🏡 DEVANT LA MAISON DE NED FLANDERS", 15*time.Millisecond)
	typeWriter("==================================================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("👨‍🦳 Ned : Oh ! Bonjour petit voisin ! belle journée, n'est-ce pas ?", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	var choice int
	typeWriter("💭 Comment voulez-vous aborder Ned ?", 15*time.Millisecond)
	typeWriter("1. Ned, as-tu vu Homer aujourd'hui ?", 15*time.Millisecond)
	typeWriter("2. Tu sais s'il est allé au bar de Moe ?", 15*time.Millisecond)
	typeWriter("3. Tes un voisin ringard, Ned !", 15*time.Millisecond)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		typeWriter("👨‍🦳 Ned : Oh oui ! J'ai vu Homer ce matin, diddly-dang !", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Il avait l'air... comment dire... plus louche que d'habitude !'", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Il marmonnait quelque chose à propos de \"bière gratuite\" et \"plan secret\"...", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer avait un plan secret impliquant de la bière gratuite !", 15*time.Millisecond)


		AddIngredient(c, "Matériau de base", "l'aide de Ned")

		progress.NedCompleted = true
		progress.Stage = 2

	case 2:
		typeWriter("👨‍🦳 Ned : Effectivement ! Je l'ai vu marcher vers le bar de Moe...", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Mais'il avait l'air différent, diddly-dong ! Presque... déterminé ?", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Ce qui est surprenant pour Homer, tu en conviendras !", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer est allé chez Moe avec une détermination inhabituelle !", 15*time.Millisecond)


		AddIngredient(c, "Matériau de base", "l'aide de Ned")

		progress.Stage = 2

	case 3:
		typeWriter("👨‍🦳 Ned : Oh non ! Cette impolitesse ne peut rester impunie !", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Au nom du Seigneur, je vais te donner une leçon de bonnes manières !", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("⚔️ Ned se transforme en CHEVALIER BIBLIQUE !", 15*time.Millisecond)

		nedMonster := InitMonster("Ned Flanders (Mode Biblique)", 80, 25)
		ScenarioCombat(c, &nedMonster, progress, nedPattern, "ned")
		return

	default:
		typeWriter("❌ Choix invalide.", 15*time.Millisecond)
		NedFlanders(c, progress)
		return
	}

	typeWriter("", 15*time.Millisecond)
	ScenarioMenu(c, progress)
}


func characterTurnNed(c *Character, m *Monster, t int, progress *ScenarioProgress) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter("🎉 Victoire ! Ned redevient normal...", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Oh... désolé pour cet éclat ! Diddly-dang, que m'est-il arrivé ?", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Pour me faire pardonner, l'aisse-moi t'aider...", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : J'ai vu Homer marcher vers le bar de Moe, l'air louche...", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer était louche en allant chez Moe !", 15*time.Millisecond)


		AddIngredient(c, "Matériau de base", "la maison de Ned")

		progress.Stage = 2
		ScenarioMenu(c, progress)
	} else {
		typeWriter("⚔️ À votre tour ! Choisissez une action :", 15*time.Millisecond)
		typeWriter("1. 💥 Attaquer", 15*time.Millisecond)
		typeWriter("2. 🎒 Fouiller dans votre sac", 15*time.Millisecond)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			attackMonster(c, m)
			nedPattern(m, turn)
			characterTurnNed(c, m, turn+1, progress)
		case 2:
			typeWriter(AccessInventory(*c), 15*time.Millisecond)
			characterTurnNed(c, m, turn, progress)
		default:
			typeWriter("❌ Choix invalide.", 15*time.Millisecond)
			characterTurnNed(c, m, turn, progress)
		}
	}
}


func nedPattern(m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("⛪ Ned utilise CITATION BIBLIQUE PARALYSANTE' !", combatDelay)
		typeWriter("👨‍🦳 Ned : Tu ne tueras point... mais je peux t'étourdir un peu !", combatDelay)
		typeWriter(fmt.Sprintf("📖 Dégâts sacrés : %d points !", damage), combatDelay)
	} else if turn%2 == 0 {
		damage := m.power + 10
		typeWriter("🏏 Ned utilise BATTE BÉNITE' !", combatDelay)
		typeWriter("👨‍🦳 Ned : Cette batte a été bénie par le révérend Lovejoy !", combatDelay)
		typeWriter(fmt.Sprintf("⚡ Dégâts divins : %d points !", damage), combatDelay)
	} else {
		damage := m.power
		typeWriter("✋ Ned utilise SERMON MORALISATEUR !", combatDelay)
		typeWriter("👨‍🦳 Ned : Diddly-dang ! Tu devrais avoir honte !", combatDelay)
		typeWriter(fmt.Sprintf("😇 Dégâts de culpabilité : %d points !", damage), combatDelay)
	}
}