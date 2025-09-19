package projetred

import (
	"fmt"
	"time"
)

func SkillsMenuSimple(c *Character) {
	textDelay := 15 * time.Millisecond

	for {
		typeWriter("🎯 MENU DES COMPÉTENCES", textDelay)
		typeWriter("=======================", textDelay)
		typeWriter("", textDelay)

		typeWriter("1. 📋 Voir mes compétences", textDelay)
		typeWriter("2. ⚡ Utiliser une compétence de soin", textDelay)
		typeWriter("0. 🚪 Retour", textDelay)
		typeWriter("", textDelay)

		var choice string
		ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
		fmt.Scan(&choice)

		switch choice {
		case "1":
			ShowSkillsFromCharacter(c)
		case "2":
			UseHealingSkillFromCharacter(c)
		case "0":
			return
		default:
			typeWriter("❌ Choix invalide.", textDelay)
		}
	}
}

func ShowSkillsFromCharacter(c *Character) {
	textDelay := 15 * time.Millisecond

	typeWriter("📋 VOS COMPÉTENCES", textDelay)
	typeWriter("==================", textDelay)
	typeWriter("", textDelay)

	typeWriter(fmt.Sprintf("👤 Personnage : %s", c.class), textDelay)
	typeWriter("", textDelay)

	for i, skill := range c.skills {
		typeWriter(fmt.Sprintf("%d. ✅ %s", i+1, skill), textDelay)

		switch skill {
		case "Coup de fronde vicieux":
			typeWriter("   📝 Attaque précise avec le lance-pierre (+45 dégâts)", textDelay)
		case "Blague empoisonnée":
			typeWriter("   📝 Une blague si mauvaise quelle fait mal (+35 dégâts)", textDelay)
		case "Échappée en skateboard":
			typeWriter("   📝 Récupère de lénergie en faisant du skate (+25 PV)", textDelay)
		case "Solo de jazz envoûtant":
			typeWriter("   📝 Musique qui inspire et blesse (+50 dégâts, +15 PV)", textDelay)
		case "Leçon de morale dévastatrice":
			typeWriter("   📝 Sermon qui démoralise lennemi (+60 dégâts)", textDelay)
		case "Méditation bouddhiste":
			typeWriter("   📝 Restaure lénergie spirituelle (+40 PV)", textDelay)
		case "Regard hypnotique":
			typeWriter("   📝 Fixe lennemi avec intensité (+30 dégâts)", textDelay)
		case "Cri strident":
			typeWriter("   📝 Cri de bébé assourdissant (+40 dégâts)", textDelay)
		case "Attaque surprise du berceau":
			typeWriter("   📝 Récupère en faisant la sieste (+35 PV)", textDelay)
		default:
			typeWriter("   📝 Compétence spéciale", textDelay)
		}
		typeWriter("", textDelay)
	}
}

func UseHealingSkillFromCharacter(c *Character) {
	textDelay := 15 * time.Millisecond

	if c.PV == c.PVmax {
		typeWriter("❤️ Vous êtes déjà en pleine forme !", textDelay)
		return
	}

	healingSkills := []string{}
	healingAmounts := map[string]int{}

	for _, skill := range c.skills {
		switch skill {
		case "Échappée en skateboard":
			healingSkills = append(healingSkills, skill)
			healingAmounts[skill] = 25
		case "Méditation bouddhiste":
			healingSkills = append(healingSkills, skill)
			healingAmounts[skill] = 40
		case "Attaque surprise du berceau":
			healingSkills = append(healingSkills, skill)
			healingAmounts[skill] = 35
		}
	}

	if len(healingSkills) == 0 {
		typeWriter("❌ Aucune compétence de soin disponible.", textDelay)
		typeWriter(fmt.Sprintf("📋 Vos compétences actuelles : %v", c.skills), textDelay)
		return
	}

	typeWriter("⚡ COMPÉTENCES DE SOIN DISPONIBLES", textDelay)
	typeWriter("==================================", textDelay)
	typeWriter("", textDelay)

	for i, skill := range healingSkills {
		typeWriter(fmt.Sprintf("%d. %s", i+1, skill), textDelay)
		typeWriter(fmt.Sprintf("   ❤️ Soins: %d PV", healingAmounts[skill]), textDelay)
		typeWriter("", textDelay)
	}

	typeWriter("0. Annuler", textDelay)
	typeWriter("", textDelay)
	typeWriter(fmt.Sprintf("❤️ PV actuels: %d/%d", c.PV, c.PVmax), textDelay)
	typeWriter("👉 Quelle compétence voulez-vous utiliser ?", textDelay)

	var choice int
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(healingSkills) {
		typeWriter("❌ Choix invalide.", textDelay)
		return
	}

	selectedSkill := healingSkills[choice-1]
	healAmount := healingAmounts[selectedSkill]

	oldPV := c.PV
	c.PV += healAmount
	if c.PV > c.PVmax {
		c.PV = c.PVmax
	}

	actualHeal := c.PV - oldPV

	typeWriter(fmt.Sprintf("✨ Vous utilisez %s' !", selectedSkill), textDelay)

	switch selectedSkill {
	case "Échappée en skateboard":
		typeWriter("🛹 Bart fait quelques figures en skateboard pour récupérer !", textDelay)
	case "Méditation bouddhiste":
		typeWriter("🧘 Lisa médite profondément pour restaurer son énergie...", textDelay)
	case "Attaque surprise du berceau":
		typeWriter("😴 Maggie fait une petite sieste réparatrice...", textDelay)
	}

	typeWriter(fmt.Sprintf("❤️ Vous récupérez %d PV ! (%d/%d)", actualHeal, c.PV, c.PVmax), textDelay)
}

func UseCombatSkillFromCharacter(c *Character, m *Monster) bool {
	textDelay := 15 * time.Millisecond

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
		typeWriter("❌ Aucune compétence dattaque disponible.", textDelay)
		return false
	}

	typeWriter("⚔️ COMPÉTENCES DATTAQUE DISPONIBLES", textDelay)
	typeWriter("====================================", textDelay)
	typeWriter("", textDelay)

	for i, skill := range attackSkills {
		typeWriter(fmt.Sprintf("%d. %s", i+1, skill), textDelay)
		typeWriter(fmt.Sprintf("   💥 Dégâts: %d", attackDamages[skill]), textDelay)
		typeWriter("", textDelay)
	}

	typeWriter("0. Retour au menu de combat", textDelay)
	typeWriter("", textDelay)
	typeWriter("👉 Quelle compétence voulez-vous utiliser ?", textDelay)

	var choice int
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&choice)

	if choice == 0 {
		return false
	}

	if choice < 1 || choice > len(attackSkills) {
		typeWriter("❌ Choix invalide.", textDelay)
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

	typeWriter(fmt.Sprintf("✨ %s utilise '%s !", c.class, selectedSkill), textDelay)

	switch selectedSkill {
	case "Coup de fronde vicieux":
		typeWriter("🎯 Bart vise soigneusement avec son lance-pierre !", textDelay)
	case "Blague empoisonnée":
		typeWriter("😂 Bart raconte une blague si nulle quelle fait mal !", textDelay)
	case "Solo de jazz envoûtant":
		typeWriter("🎷 Lisa joue un solo de saxophone hypnotisant !", textDelay)
	case "Leçon de morale dévastatrice":
		typeWriter("📚 Lisa donne une leçon de morale accablante !", textDelay)
	case "Regard hypnotique":
		typeWriter("👁️ Maggie fixe l'ennemi avec son regard mystérieux !", textDelay)
	case "Cri strident":
		typeWriter("😱 Maggie pousse un cri de bébé assourdissant !", textDelay)
	}

	if equipmentBonus > 0 {
		typeWriter(fmt.Sprintf("💥 %s subit %d points de dégâts ! (+%d'équipement)", m.name, damage, equipmentBonus), textDelay)
	} else {
		typeWriter(fmt.Sprintf("💥 %s subit %d points de dégâts !", m.name, damage), textDelay)
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
			typeWriter(fmt.Sprintf("🎵 La musique vous inspire ! +%d PV", healAmount), textDelay)
		}
	}

	return true
}
