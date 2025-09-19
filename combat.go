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
	theme := GetCharacterTheme(c.class)
	baseDamage := c.power
	equipmentBonus := GetTotalEquipmentBonus(c)
	damage := baseDamage + equipmentBonus

	if contains(c.inventory, "Donut empoisonné") {
		damage += 30
		ColoredTypeWriter("☠️ Votre attaque est empoisonnée par le donut toxique !", textDelay, BrightMagenta+Bold)

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

	attackText := fmt.Sprintf("💥 %s attaque %s", strings.ToUpper(c.class), m.name)
	damageText := DamageNumber(damage, damage > baseDamage+equipmentBonus)

	if equipmentBonus > 0 {
		bonusText := Colorize(fmt.Sprintf("(+%d équipement)", equipmentBonus), BrightGreen+Bold)
		ThemedTypeWriter(attackText+" "+damageText+" "+bonusText, textDelay, theme, "primary")
	} else {
		ThemedTypeWriter(attackText+" "+damageText, textDelay, theme, "primary")
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

func CombatInterface(c *Character, m *Monster, turn int, enemyPattern func(*Monster, int)) {
	combatDelay := 0 * time.Millisecond

	for {

		if c.PV <= 0 {
			Wasted(c)
			return
		}
		if m.PV <= 0 {
			DisplayVictory(c, m)
			return
		}

		DisplayCombatScreen(c, m, turn)

		action := GetPlayerAction()

		switch action {
		case 1:
			ClearScreen()
			DisplayCombatHeader(c, m)
			attackMonsterInstant(c, m)
			if m.PV <= 0 {
				DisplayVictory(c, m)
				return
			}

		case 2:
			ClearScreen()
			DisplayCombatHeader(c, m)
			if UseCombatSkillFromCharacterInstant(c, m) {
				if m.PV <= 0 {
					DisplayVictory(c, m)
					return
				}
			} else {
				continue
			}

		case 3:
			ClearScreen()
			DisplayCombatHeader(c, m)
			DisplayInventoryInCombat(c)
			continue

		case 4:
			ClearScreen()
			typeWriter("🏃💨 Vous fuyez le combat !", combatDelay)
			return

		default:
			typeWriter("❌ Choix invalide.", combatDelay)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if m.PV > 0 {
			time.Sleep(800 * time.Millisecond)
			typeWriter("\n🔄 Tour de lennemi...", combatDelay)
			time.Sleep(300 * time.Millisecond)
			enemyPatternInstant(m, turn, enemyPattern)

			defense := GetTotalEquipmentBonus(c) / 4
			damage := m.power - defense
			if damage < 1 {
				damage = 1
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

			time.Sleep(800 * time.Millisecond)
		}

		turn++
	}
}

func ClearScreen() {

	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}

func DisplayCombatScreen(c *Character, m *Monster, turn int) {
	ClearScreen()

	DisplayCombatHeader(c, m)

	DisplayCombatDetails(c, m, turn)

	DisplayActionMenu()
}

func DisplayCombatHeader(c *Character, m *Monster) {
	combatDelay := 0 * time.Millisecond

	ColoredTypeWriter("═══════════════════════════════════════════════════", combatDelay, BrightRed+Bold)
	ColoredTypeWriter("                ⚔️  COMBAT EN COURS  ⚔️", combatDelay, BrightYellow+Bold)
	ColoredTypeWriter("═══════════════════════════════════════════════════", combatDelay, BrightRed+Bold)
	typeWriter("", combatDelay)
}

func DisplayCombatDetails(c *Character, m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	theme := GetCharacterTheme(c.class)

	playerHealthBar := HealthBar(c.PV, c.PVmax, 20)
	ThemedTypeWriter(fmt.Sprintf("👤 %s", strings.ToUpper(c.class)), combatDelay, theme, "primary")
	ColoredTypeWriter(fmt.Sprintf("   %s", playerHealthBar), combatDelay, "")

	equipBonus := GetTotalEquipmentBonus(c)
	if equipBonus > 0 {
		powerText := fmt.Sprintf("💪 Power: %d %s", c.power, Colorize(fmt.Sprintf("(+%d équipement)", equipBonus), BrightGreen+Bold))
		ColoredTypeWriter(fmt.Sprintf("   %s", powerText), combatDelay, BrightWhite)
	} else {
		ThemedTypeWriter(fmt.Sprintf("   💪 Power: %d", c.power), combatDelay, theme, "text")
	}

	typeWriter("", combatDelay)
	ColoredTypeWriter("                      🆚", combatDelay, BrightMagenta+Bold)
	typeWriter("", combatDelay)

	enemyHealthBar := HealthBar(m.PV, m.PVmax, 20)
	ColoredTypeWriter(fmt.Sprintf("👹 %s", strings.ToUpper(m.name)), combatDelay, BrightRed+Bold)
	ColoredTypeWriter(fmt.Sprintf("   %s", enemyHealthBar), combatDelay, "")
	ColoredTypeWriter(fmt.Sprintf("   ⚡ Power: %d", m.power), combatDelay, BrightYellow)

	typeWriter("", combatDelay)
	ColoredTypeWriter(fmt.Sprintf("🔄 Tour: %d", turn), combatDelay, BrightCyan+Bold)
	typeWriter("", combatDelay)
}

func DisplayActionMenu() {
	combatDelay := 0 * time.Millisecond
	ColoredTypeWriter("┌─────────────────────────────────────────────────┐", combatDelay, BrightCyan)
	ColoredTypeWriter("│                 ⚔️ ACTIONS ⚔️                   │", combatDelay, BrightCyan+Bold)
	ColoredTypeWriter("├─────────────────────────────────────────────────┤", combatDelay, BrightCyan)
	ColoredTypeWriter("│ 1. 💥 Attaquer                                  │", combatDelay, BrightWhite)
	ColoredTypeWriter("│ 2. 🎯 Utiliser une compétence                   │", combatDelay, BrightWhite)
	ColoredTypeWriter("│ 3. 🎒 Consulter l'inventaire                   │", combatDelay, BrightWhite)
	ColoredTypeWriter("│ 4. 🏃 Fuir le combat                           │", combatDelay, BrightWhite)
	ColoredTypeWriter("└─────────────────────────────────────────────────┘", combatDelay, BrightCyan)
	typeWriter("", combatDelay)
}

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

func GetPlayerAction() int {
	combatDelay := 0 * time.Millisecond
	typeWriter("👉 Choisissez votre action (1-4): ", combatDelay)

	var choice int
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&choice)

	return choice
}

func DisplayInventoryInCombat(c *Character) {
	combatDelay := 0 * time.Millisecond
	typeWriter("🎒 INVENTAIRE DE COMBAT", combatDelay)
	typeWriter("========================", combatDelay)
	typeWriter("", combatDelay)

	hasUsableItems := false

	for i, item := range c.inventory {
		if item == "donut magique" || item == "Donut empoisonné" {
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
		ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
		fmt.Scan(&choice)

		if choice > 0 && choice <= len(c.inventory) {
			item := c.inventory[choice-1]
			if item == "donut magique" {
				oldPV := c.PV
				TakePot(c)
				healed := c.PV - oldPV
				if healed > 0 {
					typeWriter("✨ Vous utilisez un donut magique !", combatDelay)
					typeWriter(fmt.Sprintf("❤️ Vous récupérez %d PV (%d/%d)", healed, c.PV, c.PVmax), combatDelay)
				} else {
					typeWriter("ℹ️ Vous êtes déjà au maximum de PV.", combatDelay)
				}
			}
		}
	}

}

func DisplayVictory(c *Character, m *Monster) {
	ClearScreen()

	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("                ✨ VICTOIRE ! ✨", 15*time.Millisecond)
	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter(fmt.Sprintf("🏆 %s a vaincu %s !", strings.ToUpper(c.class), m.name), 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	goldReward := 25 + (m.PVmax / 10)
	c.gold += goldReward
	typeWriter(fmt.Sprintf("💰 Vous gagnez %d dollars !", goldReward), 15*time.Millisecond)

}

func ScenarioCombat(c *Character, m *Monster, progress *ScenarioProgress, enemyPattern func(*Monster, int), scenarioType string) {
	for turn := 1; ; turn++ {

		if c.PV <= 0 {
			Wasted(c)
			return
		}
		if m.PV <= 0 {
			DisplayScenarioVictory(c, m, progress, scenarioType)
			return
		}

		DisplayCombatScreen(c, m, turn)

		action := GetPlayerAction()

		switch action {
		case 1:
			ClearScreen()
			DisplayCombatHeader(c, m)
			attackMonster(c, m)
			if m.PV <= 0 {
				DisplayScenarioVictory(c, m, progress, scenarioType)
				return
			}

		case 2:
			ClearScreen()
			DisplayCombatHeader(c, m)
			if UseCombatSkillFromCharacter(c, m) {
				if m.PV <= 0 {
					DisplayScenarioVictory(c, m, progress, scenarioType)
					return
				}
			} else {
				continue
			}

		case 3:
			ClearScreen()
			DisplayCombatHeader(c, m)
			DisplayInventoryInCombat(c)
			continue

		case 4:
			ClearScreen()
			typeWriter("🏃💨 Vous fuyez le combat !", 15*time.Millisecond)

			ScenarioMenu(c, progress)
			return

		default:
			typeWriter("❌ Choix invalide.", 15*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if m.PV > 0 {
			time.Sleep(1 * time.Second)
			typeWriter("\n🔄 Tour de lennemi...", 15*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			enemyPattern(m, turn)

			defense := GetTotalEquipmentBonus(c) / 4
			damage := m.power - defense
			if damage < 1 {
				damage = 1
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

func DisplayScenarioVictory(c *Character, m *Monster, progress *ScenarioProgress, scenarioType string) {
	ClearScreen()

	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("                ✨ VICTOIRE ! ✨", 15*time.Millisecond)
	typeWriter("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter(fmt.Sprintf("🏆 %s a vaincu %s !", strings.ToUpper(c.class), m.name), 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	goldReward := 25 + (m.PVmax / 10)
	c.gold += goldReward
	typeWriter(fmt.Sprintf("💰 Vous gagnez %d dollars !", goldReward), 15*time.Millisecond)

	switch scenarioType {
	case "ned":
		typeWriter("👨‍🦳 Ned : Oh... désolé pour cet éclat ! Diddly-dang, que m'est-il arrivé ?", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : Pour me faire pardonner, l'aisse-moi t'aider...", 15*time.Millisecond)
		typeWriter("👨‍🦳 Ned : J'ai vu Homer marcher vers le bar de Moe, l'air louche...'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer était louche en allant chez Moe !", 15*time.Millisecond)

		AddIngredient(c, "Matériau de base", "la maison de Ned")

		progress.Stage = 2

	case "barney":
		typeWriter("🍺 Barney : *rot* Daccord... d'accord... tu gagnes...", 15*time.Millisecond)
		typeWriter("🍺 Barney : Tiens... *rot* ...Homer a l'aissé tomber ça...", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("📋 Barney vous tend un ticket froissé : CONCOURS DONUT GÉANT - Comic Book Store", 15*time.Millisecond)

		AddIngredient(c, "Carte Itchy & Scratchy", "le bar de Moe")

		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Ticket de concours menant au magasin de BD !", 15*time.Millisecond)
		progress.HasClue2 = true
		progress.MoeCompleted = true
		progress.Stage = 3

	case "comic":
		typeWriter("👨‍💻 Comic Book Guy : Impossible ! J'ai été vaincu par... un amateur !", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : Bon... ton père est au parc dattractions...", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : Il participe au Grand Concours du Donut Cosmique !", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)

		AddIngredient(c, "Carte Itchy & Scratchy", "le magasin de BD")

		typeWriter("🔍 INDICE OBTENU : Homer au concours de donuts du parc dattractions !", 15*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4

	case "bob":
		typeWriter("🌴 Tahiti Bob : Impossible ! Vaincu par un gamin !", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : Bon... ton père est dans la Maison des Donuts Magiques...", 15*time.Millisecond)

		AddIngredient(c, "Matériau de base", "le parc d'attractions")

	case "gencives":
		typeWriter("👻 Gencives Sanglantes : Impossible... vaincu par mon ancienne élève...", 15*time.Millisecond)
		typeWriter("🎺 Gencives Sanglantes : Tu as... évolué, Lisa... Le jazz vit en toi...", 15*time.Millisecond)
		typeWriter("👻 Gencives Sanglantes : Ton père... il est dans la Maison des Donuts... Joue pour lui...", 15*time.Millisecond)
		typeWriter("🎷 *La s'ilhouette fantomatique disparaît dans une mélodie apaisante*", 15*time.Millisecond)

		AddIngredient(c, "Note de Jazz Éternelle", "l'esprit de Gencives Sanglantes")

	case "bebe":
		typeWriter("😭 Bébé Furieux : WAAAAHHH... NO WIN... (Je ne gagne pas...)", 15*time.Millisecond)
		typeWriter("👶 Bébé Furieux : YOU STRONG BABY... (Tu es un bébé fort...)", 15*time.Millisecond)
		typeWriter("😊 Bébé Furieux : PAPA HOMER... DONUT HOUSE... (Papa Homer... maison donuts...)", 15*time.Millisecond)
		typeWriter("👶 *Le Bébé Furieux tend un biberon en signe de respect*", 15*time.Millisecond)

		AddIngredient(c, "Biberon de Respect", "le Bébé Furieux")
	}

	ScenarioMenu(c, progress)
}

func attackMonsterInstant(c *Character, m *Monster) {
	combatDelay := 0 * time.Millisecond
	baseDamage := c.power
	equipmentBonus := GetTotalEquipmentBonus(c)
	damage := baseDamage + equipmentBonus

	if contains(c.inventory, "Donut empoisonné") {
		damage += 30
		typeWriter("☠️ Votre attaque est empoisonnée par le donut toxique !", combatDelay)

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
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts ! (+%d'équipement)", c.class, m.name, damage, equipmentBonus), combatDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s attaque %s et inflige %d points de dégâts !", c.class, m.name, damage), combatDelay)
	}
}

func UseCombatSkillFromCharacterInstant(c *Character, m *Monster) bool {
	combatDelay := 0 * time.Millisecond

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
		typeWriter("❌ Aucune compétence dattaque disponible.", combatDelay)
		return false
	}

	typeWriter("⚔️ COMPÉTENCES DATTAQUE DISPONIBLES", combatDelay)
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
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
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

	m.PV -= damage
	if m.PV < 0 {
		m.PV = 0
	}

	typeWriter(fmt.Sprintf("✨ %s utilise '%s !", c.class, selectedSkill), combatDelay)

	switch selectedSkill {
	case "Coup de fronde vicieux":
		typeWriter("🎯 Bart vise soigneusement avec son lance-pierre !", combatDelay)
	case "Blague empoisonnée":
		typeWriter("😂 Bart raconte une blague si nulle quelle fait mal !", combatDelay)
	case "Solo de jazz envoûtant":
		typeWriter("🎷 Lisa joue un solo de saxophone hypnotisant !", combatDelay)
	case "Leçon de morale dévastatrice":
		typeWriter("📚 Lisa donne une leçon de morale accablante !", combatDelay)
	case "Regard hypnotique":
		typeWriter("👁️ Maggie fixe lennemi avec son regard mystérieux !", combatDelay)
	case "Cri strident":
		typeWriter("😱 Maggie pousse un cri de bébé assourdissant !", combatDelay)
	}

	if equipmentBonus > 0 {
		typeWriter(fmt.Sprintf("💥 %s subit %d points de dégâts ! (+%d'équipement)", m.name, damage, equipmentBonus), combatDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s subit %d points de dégâts !", m.name, damage), combatDelay)
	}

	switch selectedSkill {
	case "Solo de jazz envoûtant":

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

func enemyPatternInstant(m *Monster, turn int, pattern func(*Monster, int)) {

	pattern(m, turn)
}

func milhousePatternInstant(m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("🕶️ Milhouse utilise CRISE DE NERFS PARALYSANTE' !", combatDelay)
		typeWriter("👦 Milhouse : Tout le monde me déteste !", combatDelay)
		typeWriter(fmt.Sprintf("😭 Dégâts de désespoir : %d points !", damage), combatDelay)
	} else if turn%2 == 0 {
		damage := m.power + 5
		typeWriter("📚 Milhouse utilise 'RÉCITATION ENNUYEUSE !", combatDelay)
		typeWriter("👦 Milhouse : Laissez-moi vous parler de mes allergies !", combatDelay)
		typeWriter(fmt.Sprintf("⚡ Dégâts dennui : %d points !", damage), combatDelay)
	} else {
		damage := m.power
		typeWriter("😨 Milhouse utilise PLAINTES DÉSESPÉRÉES !", combatDelay)
		typeWriter("👦 Milhouse : Oh non, pas encore !", combatDelay)
		typeWriter(fmt.Sprintf("😓 Dégâts de désespoir : %d points !", damage), combatDelay)
	}
}
