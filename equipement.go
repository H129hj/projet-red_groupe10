package projetred

import (
	"fmt"
	"time"
)


func (c *Character) AddEquipement(nom string, valeur int) {
	if c.equipement == nil {
		c.equipement = make(map[string]int)
	}
	c.equipement[nom] = valeur
}

func (c *Character) RemoveEquipement(nom string) {
	delete(c.equipement, nom)
}


func EquipmentMenu(c *Character) {
	textDelay := 15 * time.Millisecond

	for {
		typeWriter("⚔️ GESTION DES ÉQUIPEMENTS", textDelay)
		typeWriter("==========================", textDelay)
		typeWriter("", textDelay)

		typeWriter("1. 👕 Équiper un objet du sac à dos", textDelay)
		typeWriter("2. 🔄 Déséquiper un objet", textDelay)
		typeWriter("3. 📊 Voir équipements portés", textDelay)
		typeWriter("4. 🎒 Voir le sac à dos complet", textDelay)
		typeWriter("0. 🚪 Retour au menu principal", textDelay)
		typeWriter("", textDelay)

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			EquipFromInventory(c)
		case "2":
			UnequipToInventory(c)
		case "3":
			ShowEquippedItems(c)
		case "4":
			typeWriter(AccessInventory(*c), textDelay)
			typeWriter("", textDelay)
			typeWriter("Appuyez sur Entrée pour continuer...", textDelay)
			fmt.Scanln()
		case "0":
			return
		default:
			typeWriter("❌ Choix invalide.", textDelay)
		}
	}
}


func EquipFromInventory(c *Character) {
	textDelay := 15 * time.Millisecond

	if len(c.inventory) == 0 {
		typeWriter("❌ Votre sac à dos est vide.", textDelay)
		return
	}

	typeWriter("🎒 OBJETS DANS VOTRE SAC À DOS", textDelay)
	typeWriter("==============================", textDelay)
	typeWriter("", textDelay)


	equipableItems := []string{}
	for _, item := range c.inventory {
		if IsEquipable(item) {
			equipableItems = append(equipableItems, item)
		}
	}

	if len(equipableItems) == 0 {
		typeWriter("❌ Aucun objet équipable dans votre sac à dos.", textDelay)
		return
	}

	for i, item := range equipableItems {
		typeWriter(fmt.Sprintf("%d. %s", i+1, item), textDelay)
	}

	typeWriter("0. Annuler", textDelay)
	typeWriter("", textDelay)
	typeWriter("👉 Quel objet voulez-vous équiper ?", textDelay)

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(equipableItems) {
		typeWriter("❌ Choix invalide.", textDelay)
		return
	}

	selectedItem := equipableItems[choice-1]


	if _, equipped := c.equipement[selectedItem]; equipped {
		typeWriter("❌ Cet objet est déjà équipé !", textDelay)
		return
	}


	equipValue := GetEquipmentValue(selectedItem)
	c.AddEquipement(selectedItem, equipValue)

	typeWriter(fmt.Sprintf("✅ %s équipé avec succès !", selectedItem), textDelay)
	typeWriter(fmt.Sprintf("⚡ Valeur déquipement : %d", equipValue), textDelay)
}


func UnequipToInventory(c *Character) {
	textDelay := 15 * time.Millisecond

	if len(c.equipement) == 0 {
		typeWriter("❌ Aucun équipement porté.", textDelay)
		return
	}

	typeWriter("⚔️ ÉQUIPEMENTS PORTÉS", textDelay)
	typeWriter("====================", textDelay)
	typeWriter("", textDelay)

	equippedItems := []string{}
	for equipName := range c.equipement {
		equippedItems = append(equippedItems, equipName)
	}

	for i, item := range equippedItems {
		value := c.equipement[item]
		typeWriter(fmt.Sprintf("%d. %s (valeur: %d)", i+1, item, value), textDelay)
	}

	typeWriter("0. Annuler", textDelay)
	typeWriter("", textDelay)
	typeWriter("👉 Quel objet voulez-vous déséquiper ?", textDelay)

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(equippedItems) {
		typeWriter("❌ Choix invalide.", textDelay)
		return
	}

	selectedItem := equippedItems[choice-1]


	if !limitedInventory(c) {
		return
	}


	c.RemoveEquipement(selectedItem)

	typeWriter(fmt.Sprintf("✅ %s déséquipé avec succès !", selectedItem), textDelay)
	typeWriter("📦 Lobjet reste dans votre inventaire.", textDelay)
}


func ShowEquippedItems(c *Character) {
	textDelay := 15 * time.Millisecond

	typeWriter("⚔️ ÉQUIPEMENTS PORTÉS", textDelay)
	typeWriter("====================", textDelay)
	typeWriter("", textDelay)

	if len(c.equipement) == 0 {
		typeWriter("❌ Aucun équipement porté.", textDelay)
	} else {
		totalValue := 0
		for equipName, value := range c.equipement {
			typeWriter(fmt.Sprintf("✅ %s (valeur: %d)", equipName, value), textDelay)
			totalValue += value
		}
		typeWriter("", textDelay)
		typeWriter(fmt.Sprintf("💪 Valeur totale d'équipement : %d", totalValue), textDelay)
	}

	typeWriter("", textDelay)
	typeWriter("Appuyez sur Entrée pour continuer...", textDelay)
	fmt.Scanln()
}


func IsEquipable(itemName string) bool {
	equipableItems := []string{
		"Lance-pierre de Bart", "Saxophone de Lisa", "Biberon de Maggie",
		"Batte de baseball", "Bouclier fait maison", "T-shirt rouge",
		"Robe de première de classe", "Grenouillère bleue",
		"Super Batte de Bart", "Saxophone Enchanté", "Bouclier de Springfield",
		"Biberon Cosmique", "Saxophone Spirituel", "Super Biberon",
		"Serre-tête", "Collier de perles", "Short bleu",
		"Chaussures de sport", "Nœud rose", "Tétine magique",
	}

	for _, item := range equipableItems {
		if item == itemName {
			return true
		}
	}
	return false
}


func GetEquipmentValue(itemName string) int {
	equipmentValues := map[string]int{
		"Lance-pierre de Bart":       25,
		"Saxophone de Lisa":          30,
		"Biberon de Maggie":          20,
		"Batte de baseball":          35,
		"Bouclier fait maison":       25,
		"T-shirt rouge":              10,
		"Robe de première de classe": 15,
		"Grenouillère bleue":         12,
		"Super Batte de Bart":        75,
		"Saxophone Enchanté":         85,
		"Bouclier de Springfield":    60,
		"Biberon Cosmique":           120,
		"Saxophone Spirituel":        95,
		"Super Biberon":              140,
		"Serre-tête":                 5,
		"Collier de perles":          7,
		"Short bleu":                 5,
		"Chaussures de sport":        7,
		"Nœud rose":                  5,
		"Tétine magique":             7,
	}

	if value, exists := equipmentValues[itemName]; exists {
		return value
	}
	return 10
}


func GetTotalEquipmentBonus(c *Character) int {
	total := 0
	for _, value := range c.equipement {
		total += value
	}
	return total
}