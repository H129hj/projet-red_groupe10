package projetred

import (
	"fmt"
	"strings"
	"time"
)

type Monster struct {
	name  string
	PVmax int
	PV    int
	power int
}

func attackMonster(c *Character, m *Monster) {
	var textDelay = 15 * time.Millisecond
	baseDamage := c.power
	equipmentBonus := GetTotalEquipmentBonus(c)
	damage := baseDamage + equipmentBonus

	// Vérifier si on a un donut empoisonné
	if contains(c.inventory, "Donut empoisonné") {
		damage += 30
		typeWriter("☠️ Votre attaque est empoisonnée par le donut toxique !", textDelay)
		// Retirer le donut empoisonné après usage
		for i, item := range c.inventory {
			if item == "Donut empoisonné" {
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				break
			}
		}
	}

	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}

	if equipmentBonus > 0 {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts ! (+%d équipement)", c.class, m.name, damage, equipmentBonus), textDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts !", c.class, m.name, damage), textDelay)
	}
}

func InitMonster(name string, PVmax int, power int) Monster {
	return Monster{
		name:  name,
		PVmax: PVmax,
		PV:    PVmax,
		power: power,
	}
}

func traningFight(c *Character, m *Monster) {
	CombatInterface(c, m, 1, milhousePatternInstant)
}

// Interface de combat visuelle complète avec affichage instantané
func CombatInterface(c *Character, m *Monster, turn int, enemyPattern func(*Monster, int)) {
	combatDelay := 0 * time.Millisecond // Texte instantané en combat

	for {
		// Vérifier l'état du combat
		if c.PV <= 0 {
			Wasted(c)
			return
		}
		if m.PV <= 0 {
			DisplayVictory(c, m)
			return
		}

		// Afficher l'interface de combat
		DisplayCombatScreen(c, m, turn)

		// Tour du joueur
		action := GetPlayerAction()

		switch action {
		case 1: // Attaquer
			ClearScreen()
			DisplayCombatHeader(c, m)
			attackMonsterInstant(c, m)
			if m.PV <= 0 {
				DisplayVictory(c, m)
				return
			}

		case 2: // Compétences
			ClearScreen()
			DisplayCombatHeader(c, m)
			if UseCombatSkillFromCharacterInstant(c, m) {
				if m.PV <= 0 {
					DisplayVictory(c, m)
					return
				}
			} else {
				continue // Retour au menu si annulé
			}

		case 3: // Inventaire
			ClearScreen()
			DisplayCombatHeader(c, m)
			DisplayInventoryInCombat(c)
			continue

		case 4: // Fuir
			ClearScreen()
			typeWriter("🏃💨 Vous fuyez le combat !", combatDelay)
			return

		default:
			typeWriter("❌ Choix invalide.", combatDelay)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		// Tour de l'ennemi
		if m.PV > 0 {
			time.Sleep(800 * time.Millisecond) // Pause courte pour la lisibilité
			typeWriter("\n🔄 Tour de l'ennemi...", combatDelay)
			time.Sleep(300 * time.Millisecond)
			enemyPatternInstant(m, turn, enemyPattern)

			// Appliquer les dégâts de l'ennemi
			defense := GetTotalEquipmentBonus(c) / 4 // La défense réduit les dégâts
			damage := m.power - defense
			if damage < 1 {
				damage = 1 // Dégâts minimum
			}

			c.PV -= damage
			if c.PV < 0 {
				c.PV = 0
			}

			if defense > 0 {
				typeWriter(fmt.Sprintf("🛡️ Vous subissez %d dégâts ! (-%d défense)", damage, defense), combatDelay)
			} else {
				typeWriter(fmt.Sprintf("💔 Vous subissez %d dégâts !", damage), combatDelay)
			}

			time.Sleep(800 * time.Millisecond) // Pause courte avant le prochain tour
		}

		turn++
	}
}

// Nettoyer l'écran (simulation)
func ClearScreen() {
	// Afficher des lignes vides pour simuler un nettoyage d'écran
	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}

// Afficher l'écran de combat complet
func DisplayCombatScreen(c *Character, m *Monster, turn int) {
	ClearScreen()

	// En-tête du combat
	DisplayCombatHeader(c, m)

	// Informations détaillées
	DisplayCombatDetails(c, m, turn)

	// Menu d'actions
	DisplayActionMenu()
}

// En-tête du combat
func DisplayCombatHeader(c *Character, m *Monster) {
	combatDelay := 0 * time.Millisecond
	typeWriter("═══════════════════════════════════════════════════", combatDelay)
	typeWriter("                ⚔️  COMBAT EN COURS  ⚔️", combatDelay)
	typeWriter("═══════════════════════════════════════════════════", combatDelay)
	typeWriter("", combatDelay)
}

// Détails du combat
func DisplayCombatDetails(c *Character, m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	// Barre de vie du joueur
	playerHealthBar := CreateHealthBar(c.PV, c.PVmax, 20, "❤️", "💔")
	typeWriter(fmt.Sprintf("👤 %s", strings.ToUpper(c.class)), combatDelay)
	typeWriter(fmt.Sprintf("   %s %d/%d PV", playerHealthBar, c.PV, c.PVmax), combatDelay)

	// Stats du joueur
	equipBonus := GetTotalEquipmentBonus(c)
	if equipBonus > 0 {
		typeWriter(fmt.Sprintf("   💪 Power: %d (+%d équipement)", c.power, equipBonus), combatDelay)
	} else {
		typeWriter(fmt.Sprintf("   💪 Power: %d", c.power), combatDelay)
	}

	typeWriter("", combatDelay)
	typeWriter("                      🆚", combatDelay)
	typeWriter("", combatDelay)

	// Barre de vie de l'ennemi
	enemyHealthBar := CreateHealthBar(m.PV, m.PVmax, 20, "💀", "🖤")
	typeWriter(fmt.Sprintf("👹 %s", strings.ToUpper(m.name)), combatDelay)
	typeWriter(fmt.Sprintf("   %s %d/%d PV", enemyHealthBar, m.PV, m.PVmax), combatDelay)
	typeWriter(fmt.Sprintf("   ⚡ Power: %d", m.power), combatDelay)

	typeWriter("", combatDelay)
	typeWriter(fmt.Sprintf("🔄 Tour: %d", turn), combatDelay)
	typeWriter("", combatDelay)
}

// Menu d'actions
func DisplayActionMenu() {
	combatDelay := 0 * time.Millisecond
	typeWriter("┌─────────────────────────────────────────────────┐", combatDelay)
	typeWriter("│                 ⚔️ ACTIONS ⚔️                   │", combatDelay)
	typeWriter("├─────────────────────────────────────────────────┤", combatDelay)
	typeWriter("│ 1. 💥 Attaquer                                  │", combatDelay)
	typeWriter("│ 2. 🎯 Utiliser une compétence                   │", combatDelay)
	typeWriter("│ 3. 🎒 Consulter l'inventaire                    │", combatDelay)
	typeWriter("│ 4. 🏃 Fuir le combat                           │", combatDelay)
	typeWriter("└─────────────────────────────────────────────────┘", combatDelay)
	typeWriter("", combatDelay)
}

// Créer une barre de vie visuelle
func CreateHealthBar(current, max, length int, fullChar, emptyChar string) string {
	if max == 0 {
		return strings.Repeat(emptyChar, length)
	}

	percentage := float64(current) / float64(max)
	filledLength := int(percentage * float64(length))

	if filledLength < 0 {
		filledLength = 0
	}
	if filledLength > length {
		filledLength = length
	}

	emptyLength := length - filledLength

	return strings.Repeat(fullChar, filledLength) + strings.Repeat(emptyChar, emptyLength)
}

// Obtenir l'action du joueur
func GetPlayerAction() int {
	combatDelay := 0 * time.Millisecond
	typeWriter("👉 Choisissez votre action (1-4): ", combatDelay)

	var choice int
	fmt.Scan(&choice)

	return choice
}

// Afficher l'inventaire en combat
func DisplayInventoryInCombat(c *Character) {
	combatDelay := 0 * time.Millisecond
	typeWriter("🎒 INVENTAIRE DE COMBAT", combatDelay)
	typeWriter("========================", combatDelay)
	typeWriter("", combatDelay)

	hasUsableItems := false

	for i, item := range c.inventory {
		if item == "Donut magique" || item == "Donut empoisonné" {
			typeWriter(fmt.Sprintf("%d. %s", i+1, item), combatDelay)
			hasUsableItems = true
		}
	}

	if !hasUsableItems {
		typeWriter("❌ Aucun objet utilisable en combat.", combatDelay)
	} else {
		typeWriter("", combatDelay)
		typeWriter("💡 Les donuts magiques restaurent 20 PV", combatDelay)
		typeWriter("💡 Les donuts empoisonnés augmentent vos dégâts", combatDelay)

		var choice int
		typeWriter("👉 Utiliser un objet (numéro) ou 0 pour retour: ", combatDelay)
		fmt.Scan(&choice)

		if choice > 0 && choice <= len(c.inventory) {
			item := c.inventory[choice-1]
			if item == "Donut magique" {
				TakePot(c)
				typeWriter("✨ Vous utilisez un donut magique !", combatDelay)
			}
		}
	}

	typeWriter("", combatDelay)
	typeWriter("Appuyez sur Entrée pour continuer...", combatDelay)
	fmt.Scanln()
}

// Afficher la victoire
func DisplayVictory(c *Character, m *Monster) {
	ClearScreen()

	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("                ✨ VICTOIRE ! ✨", 15*time.Millisecond)
	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter(fmt.Sprintf("🏆 %s a vaincu %s !", strings.ToUpper(c.class), m.name), 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	// Récompenses
	goldReward := 25 + (m.PVmax / 10)
	c.gold += goldReward
	typeWriter(fmt.Sprintf("💰 Vous gagnez %d dollars !", goldReward), 15*time.Millisecond)

	typeWriter("", 15*time.Millisecond)
	typeWriter("Appuyez sur Entrée pour continuer...", 15*time.Millisecond)
	fmt.Scanln()
}

// Combat de scénario avec gestion de la progression
func ScenarioCombat(c *Character, m *Monster, progress *ScenarioProgress, enemyPattern func(*Monster, int), scenarioType string) {
	for turn := 1; ; turn++ {
		// Vérifier l'état du combat
		if c.PV <= 0 {
			Wasted(c)
			return
		}
		if m.PV <= 0 {
			DisplayScenarioVictory(c, m, progress, scenarioType)
			return
		}

		// Afficher l'interface de combat
		DisplayCombatScreen(c, m, turn)

		// Tour du joueur
		action := GetPlayerAction()

		switch action {
		case 1: // Attaquer
			ClearScreen()
			DisplayCombatHeader(c, m)
			attackMonster(c, m)
			if m.PV <= 0 {
				DisplayScenarioVictory(c, m, progress, scenarioType)
				return
			}

		case 2: // Compétences
			ClearScreen()
			DisplayCombatHeader(c, m)
			if UseCombatSkillFromCharacter(c, m) {
				if m.PV <= 0 {
					DisplayScenarioVictory(c, m, progress, scenarioType)
					return
				}
			} else {
				continue // Retour au menu si annulé
			}

		case 3: // Inventaire
			ClearScreen()
			DisplayCombatHeader(c, m)
			DisplayInventoryInCombat(c)
			continue

		case 4: // Fuir
			ClearScreen()
			typeWriter("🏃💨 Vous fuyez le combat !", 15*time.Millisecond)
			// Retour au menu du scénario
			ScenarioMenu(c, progress)
			return

		default:
			typeWriter("❌ Choix invalide.", 15*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		// Tour de l'ennemi
		if m.PV > 0 {
			time.Sleep(1 * time.Second)
			typeWriter("\n🔄 Tour de l'ennemi...", 15*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			enemyPattern(m, turn)

			// Appliquer les dégâts de l'ennemi
			defense := GetTotalEquipmentBonus(c) / 4 // La défense réduit les dégâts
			damage := m.power - defense
			if damage < 1 {
				damage = 1 // Dégâts minimum
			}

			c.PV -= damage
			if c.PV < 0 {
				c.PV = 0
			}

			if defense > 0 {
				typeWriter(fmt.Sprintf("🛡️ Vous subissez %d dégâts ! (-%d défense)", damage, defense), 15*time.Millisecond)
			} else {
				typeWriter(fmt.Sprintf("💔 Vous subissez %d dégâts !", damage), 15*time.Millisecond)
			}

			time.Sleep(1 * time.Second)
		}
	}
}

// Afficher la victoire dans un scénario
func DisplayScenarioVictory(c *Character, m *Monster, progress *ScenarioProgress, scenarioType string) {
	ClearScreen()

	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("                ✨ VICTOIRE ! ✨", 15*time.Millisecond)
	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter(fmt.Sprintf("🏆 %s a vaincu %s !", strings.ToUpper(c.class), m.name), 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	// Récompenses
	goldReward := 25 + (m.PVmax / 10)
	c.gold += goldReward
	typeWriter(fmt.Sprintf("💰 Vous gagnez %d dollars !", goldReward), 15*time.Millisecond)

	// Gestion spécifique selon le type de scénario
	switch scenarioType {
	case "ned":
		typeWriter("👨‍🦳 Ned : 'Oh... désolé pour cet éclat ! Diddly-dang, que m'est-il arrivé ?'", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'Pour me faire pardonner, laisse-moi t'aider...'", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : 'J'ai vu Homer marcher vers le bar de Moe, l'air louche...'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer était louche en allant chez Moe !", 15*time.Millisecond)

		// Ajouter un ingrédient après le combat
		AddIngredient(c, "Matériau de base", "la maison de Ned")

		progress.Stage = 2

	case "barney":
		typeWriter("🍺 Barney : '*rot* D'accord... d'accord... tu gagnes...'", 15*time.Millisecond)
		typeWriter("🍺 Barney : 'Tiens... *rot* ...Homer a laissé tomber ça...'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("📋 Barney vous tend un ticket froissé : 'CONCOURS DONUT GÉANT - Comic Book Store'", 15*time.Millisecond)

		// Ajouter un ingrédient après le combat
		AddIngredient(c, "Carte Itchy & Scratchy", "le bar de Moe")

		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Ticket de concours menant au magasin de BD !", 15*time.Millisecond)
		progress.HasClue2 = true
		progress.MoeCompleted = true
		progress.Stage = 3

	case "comic":
		typeWriter("👨‍💻 Comic Book Guy : 'Impossible ! J'ai été vaincu par... un amateur !'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Bon... ton père est au parc d'attractions...'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Il participe au Grand Concours du Donut Cosmique !'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)

		// Ajouter un ingrédient rare après le combat
		AddIngredient(c, "Carte Itchy & Scratchy", "le magasin de BD")

		typeWriter("🔍 INDICE OBTENU : Homer au concours de donuts du parc d'attractions !", 15*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4

	case "bob":
		typeWriter("🌴 Tahiti Bob : 'Impossible ! Vaincu par un gamin !'", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Bon... ton père est dans la Maison des Donuts Magiques...'", 15*time.Millisecond)

		// Ajouter un ingrédient spécial après le boss final
		AddIngredient(c, "Matériau de base", "le parc d'attractions")
	}

	typeWriter("", 15*time.Millisecond)
	typeWriter("Appuyez sur Entrée pour continuer...", 15*time.Millisecond)
	fmt.Scanln()

	// Retour au menu du scénario
	ScenarioMenu(c, progress)
}

// Versions instantanées pour les combats (texte immédiat)

// Attaque instantanée pour les combats
func attackMonsterInstant(c *Character, m *Monster) {
	combatDelay := 0 * time.Millisecond
	baseDamage := c.power
	equipmentBonus := GetTotalEquipmentBonus(c)
	damage := baseDamage + equipmentBonus

	// Vérifier si on a un donut empoisonné
	if contains(c.inventory, "Donut empoisonné") {
		damage += 30
		typeWriter("☠️ Votre attaque est empoisonnée par le donut toxique !", combatDelay)
		// Retirer le donut empoisonné après usage
		for i, item := range c.inventory {
			if item == "Donut empoisonné" {
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				break
			}
		}
	}

	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}

	if equipmentBonus > 0 {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts ! (+%d équipement)", c.class, m.name, damage, equipmentBonus), combatDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts !", c.class, m.name, damage), combatDelay)
	}
}

// Compétences de combat instantanées
func UseCombatSkillFromCharacterInstant(c *Character, m *Monster) bool {
	combatDelay := 0 * time.Millisecond

	// Trouver les compétences d'attaque disponibles dans les skills du personnage
	attackSkills := []string{}
	attackDamages := map[string]int{}

	for _, skill := range c.skills {
		switch skill {
		case "Coup de fronde vicieux":
			attackSkills = append(attackSkills, skill)
			attackDamages[skill] = 45
		case "Blague empoisonnée":
			attackSkills = append(attackSkills, skill)
			attackDamages[skill] = 35
		case "Solo de jazz envoûtant":
			attackSkills = append(attackSkills, skill)
			attackDamages[skill] = 50
		case "Leçon de morale dévastatrice":
			attackSkills = append(attackSkills, skill)
			attackDamages[skill] = 60
		case "Regard hypnotique":
			attackSkills = append(attackSkills, skill)
			attackDamages[skill] = 30
		case "Cri strident":
			attackSkills = append(attackSkills, skill)
			attackDamages[skill] = 40
		}
	}

	if len(attackSkills) == 0 {
		typeWriter("❌ Aucune compétence d'attaque disponible.", combatDelay)
		return false
	}

	typeWriter("⚔️ COMPÉTENCES D'ATTAQUE DISPONIBLES", combatDelay)
	typeWriter("====================================", combatDelay)
	typeWriter("", combatDelay)

	for i, skill := range attackSkills {
		typeWriter(fmt.Sprintf("%d. %s", i+1, skill), combatDelay)
		typeWriter(fmt.Sprintf("   💥 Dégâts: %d", attackDamages[skill]), combatDelay)
		typeWriter("", combatDelay)
	}

	typeWriter("0. Retour au menu de combat", combatDelay)
	typeWriter("", combatDelay)
	typeWriter("👉 Quelle compétence voulez-vous utiliser ?", combatDelay)

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return false
	}

	if choice < 1 || choice > len(attackSkills) {
		typeWriter("❌ Choix invalide.", combatDelay)
		return false
	}

	selectedSkill := attackSkills[choice-1]
	baseDamage := attackDamages[selectedSkill]
	equipmentBonus := GetTotalEquipmentBonus(c)
	damage := baseDamage + equipmentBonus

	// Appliquer les dégâts
	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}

	typeWriter(fmt.Sprintf("✨ %s utilise '%s' !", c.class, selectedSkill), combatDelay)

	switch selectedSkill {
	case "Coup de fronde vicieux":
		typeWriter("🎯 Bart vise soigneusement avec son lance-pierre !", combatDelay)
	case "Blague empoisonnée":
		typeWriter("😂 Bart raconte une blague si nulle qu'elle fait mal !", combatDelay)
	case "Solo de jazz envoûtant":
		typeWriter("🎷 Lisa joue un solo de saxophone hypnotisant !", combatDelay)
	case "Leçon de morale dévastatrice":
		typeWriter("📚 Lisa donne une leçon de morale accablante !", combatDelay)
	case "Regard hypnotique":
		typeWriter("👁️ Maggie fixe l'ennemi avec son regard mystérieux !", combatDelay)
	case "Cri strident":
		typeWriter("😱 Maggie pousse un cri de bébé assourdissant !", combatDelay)
	}

	if equipmentBonus > 0 {
		typeWriter(fmt.Sprintf("💥 %s subit %d points de dégâts ! (+%d équipement)", m.name, damage, equipmentBonus), combatDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s subit %d points de dégâts !", m.name, damage), combatDelay)
	}

	// Effets spéciaux selon la compétence
	switch selectedSkill {
	case "Solo de jazz envoûtant":
		// Lisa récupère aussi des PV
		oldPV := c.PV
		c.PV += 15
		if c.PV > c.PVmax {
			c.PV = c.PVmax
		}
		healAmount := c.PV - oldPV
		if healAmount > 0 {
			typeWriter(fmt.Sprintf("🎵 La musique vous inspire ! +%d PV", healAmount), combatDelay)
		}
	}

	return true
}

// Pattern d'ennemi instantané
func enemyPatternInstant(m *Monster, turn int, pattern func(*Monster, int)) {
	// Appeler le pattern original mais modifier les délais dans les patterns spécifiques
	pattern(m, turn)
}

// Pattern d'attaque pour Milhouse (combat d'entraînement)
func milhousePatternInstant(m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("🕶️ Milhouse utilise 'CRISE DE NERFS PARALYSANTE' !", combatDelay)
		typeWriter("👦 Milhouse : 'Tout le monde me déteste !'", combatDelay)
		typeWriter(fmt.Sprintf("😭 Dégâts de désespoir : %d points !", damage), combatDelay)
	} else if turn%2 == 0 {
		damage := m.power + 5
		typeWriter("📚 Milhouse utilise 'RÉCITATION ENNUYEUSE' !", combatDelay)
		typeWriter("👦 Milhouse : 'Laissez-moi vous parler de mes allergies !'", combatDelay)
		typeWriter(fmt.Sprintf("⚡ Dégâts d'ennui : %d points !", damage), combatDelay)
	} else {
		damage := m.power
		typeWriter("😨 Milhouse utilise 'PLAINTES DÉSESPÉRÉES' !", combatDelay)
		typeWriter("👦 Milhouse : 'Oh non, pas encore !'", combatDelay)
		typeWriter(fmt.Sprintf("😓 Dégâts de désespoir : %d points !", damage), combatDelay)
	}
}
