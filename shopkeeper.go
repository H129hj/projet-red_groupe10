package projetred

import (
	"fmt"
	"math/rand"
	"time"
)

func Shopkeeper(c *Character) Character {
	randomvalue := 50 + rand.Intn(101)
	textDelay := 15 * time.Millisecond

	for {
		fmt.Println()
		MenuHeader("KWIK-E-MART D'APU", ShopTheme)
		typeWriter("🏪 Apu : Bonjour mon ami ! Que puis-je faire pour vous ?", textDelay)
		fmt.Println()

		typeWriter("1. 🛍️ Acheter un objet", textDelay)
		typeWriter("2. 💰 Vendre un objet", textDelay)
		typeWriter("3. 🔧 Atelier de bricolage d'Apu", textDelay)
		typeWriter("0. 🚪 Quitter le magasin", textDelay)

		fmt.Println()
		ColoredTypeWriter(CurrencyDisplay(c.gold), textDelay, "")
		ColoredTypeWriter("👉 Que voulez-vous faire ? ", textDelay, BrightCyan+Bold)

		var choice string
		ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
		fmt.Scan(&choice)

		switch choice {
		case "1":
			fmt.Println()
			MenuHeader("BOUTIQUE D'APU", ShopTheme)

			typeWriter("1. "+ItemDisplay("Batte de baseball de Bart", "rare")+" (50 dollars) - Arme", textDelay)
			typeWriter("2. "+ItemDisplay("Bouclier fait maison", "common")+" (40 dollars) - Armure", textDelay)
			typeWriter("3. "+ItemDisplay("Donut magique de chez Homer", "legendary")+" (10 dollars) - Consommable", textDelay)
			typeWriter(fmt.Sprintf("4. "+ItemDisplay("Carte rare Itchy & Scratchy", "rare")+" (%d dollars) - Matériau", randomvalue), textDelay)
			typeWriter("5. "+ItemDisplay("Kit de craft basique", "common")+" (75 dollars) - Matériaux", textDelay)
			typeWriter("6. ⬆️ Agrandir l'inventaire (25 dollars) - Amélioration", textDelay)
			typeWriter("7. 🔙 Retour au menu d'Apu", textDelay)

			fmt.Println()
			ColoredTypeWriter(CurrencyDisplay(c.gold), textDelay, "")
			ColoredTypeWriter("👉 Quel objet souhaitez-vous acheter ? ", textDelay, BrightCyan+Bold)

			var choice2 string
			ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
			fmt.Scan(&choice2)

			switch choice2 {
			case "1":
				if c.gold >= 50 && !contains(c.inventory, "Batte de baseball") && limitedInventory(c) {
					c.inventory = append(c.inventory, "Batte de baseball")
					c.gold -= 50
					typeWriter(fmt.Sprintf("✅ Apu : Excellente arme ! Il vous reste %d dollars.", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : Désolé, pas assez dargent mon ami !", textDelay)
				}
			case "2":
				if c.gold >= 40 && !contains(c.inventory, "Batte de baseball") && limitedInventory(c) {
					c.inventory = append(c.inventory, "Bouclier fait maison")
					c.gold -= 40
					typeWriter(fmt.Sprintf("✅ Apu : Protection garantie ! Il vous reste %d dollars.", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : Vos poches sont vides comme le frigo des Simpson !", textDelay)
				}
			case "3":
				if c.gold >= 10 && limitedInventory(c) {
					c.inventory = append(c.inventory, "Donut magique")
					c.gold -= 10
					typeWriter(fmt.Sprintf("✅ Apu : Mmm... donut magique ! Il vous reste %d dollars.", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : Même pas 10 dollars ? Allez voir Homer !", textDelay)
				}
			case "4":
				if c.gold >= randomvalue && limitedInventory(c) {
					c.inventory = append(c.inventory, "Carte Itchy & Scratchy")
					c.gold -= randomvalue
					typeWriter(fmt.Sprintf("✅ Apu : Carte très rare ! Il vous reste %d dollars.", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : Cette carte coûte plus cher que le sal'aire de Homer !", textDelay)
				}
			case "5":
				if c.gold >= 75 && limitedInventory(c) {

					c.inventory = append(c.inventory, "Matériau de base", "Matériau de base")
					c.gold -= 75
					typeWriter(fmt.Sprintf("✅ Apu : Kit de craft acheté ! Il vous reste %d dollars.", c.gold), textDelay)
				} else if !limitedInventory(c) {
					typeWriter("❌ Apu : Inventaire plein !", textDelay)
				} else {
					typeWriter("❌ Apu : Pas assez dargent pour le kit de craft !", textDelay)
				}
			case "6":
				if c.gold >= 25 {
					c.extendedInventory += 1
					c.gold -= 25
					typeWriter(fmt.Sprintf("✅ Apu : Votre inventaire a été agrandi ! Il vous reste %d dollars.", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : Pas assez dargent pour agrandir votre inventaire !", textDelay)
				}
			case "7":
				typeWriter("↩ Apu : Très bien, très bien...", textDelay)
			default:
				typeWriter("❌ Apu : Je ne comprends pas, parlez plus cl'airement !", textDelay)
			}

		case "2":
			if len(c.inventory) == 0 {
				typeWriter("⚠ Apu : Vous navez rien à vendre, revenez plus tard !", textDelay)
				continue
			}
			typeWriter("\n--- 💸 Apu rachète vos objets ---", textDelay)
			for i, v := range c.inventory {
				typeWriter(fmt.Sprintf("%d. %s", i+1, v), textDelay)
			}
			typeWriter("👉 Apu : Quel objet voulez-vous vendre ? ", textDelay)

			var sellChoice int
			ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
			fmt.Scan(&sellChoice)
			if sellChoice < 1 || sellChoice > len(c.inventory) {
				typeWriter("❌ Apu : Numéro invalide, comptez mieux que ça !", textDelay)
				continue
			}

			itemToSell := c.inventory[sellChoice-1]
			var sellPrice int
			switch itemToSell {
			case "Batte de baseball":
				sellPrice = 25
			case "Bouclier fait maison":
				sellPrice = 20
			case "Donut magique":
				sellPrice = 5
			case "Carte Itchy & Scratchy":
				sellPrice = randomvalue
			default:
				sellPrice = 1
			}
			c.inventory = append(c.inventory[:sellChoice-1], c.inventory[sellChoice:]...)
			c.gold += sellPrice
			typeWriter(fmt.Sprintf("💸 Apu : J'ai acheté votre %s pour %d dollars. Solde : %d dollars.", itemToSell, sellPrice, c.gold), textDelay)

		case "3":
			CraftMenuInShop(c, textDelay)

		case "0":
			typeWriter("\n👋 Apu : Merci pour votre visite ! Revenez vite au Kwik-E-Mart !", textDelay)
			ScenarioMenu(c, &ScenarioProgress{})
			return *c

		default:
			typeWriter("❌ Apu : Je ne comprends pas votre choix, mon ami !", textDelay)
		}
	}
}

type CraftRecipe struct {
	Name        string
	Ingredients []string
	Result      string
	Description string
}

var craftRecipes = []CraftRecipe{
	{
		Name:        "Donut Empoisonné",
		Ingredients: []string{"Donut magique", "Carte Itchy & Scratchy"},
		Result:      "Donut empoisonné",
		Description: "Un donut magique imprégné de la violence d'Itchy & Scratchy",
	},
	{
		Name:        "Super Batte",
		Ingredients: []string{"Batte de baseball", "Lance-pierre de Bart"},
		Result:      "Super Batte de Bart",
		Description: "Une batte renforcée avec la technologie du lance-pierre",
	},
	{
		Name:        "Saxophone Magique",
		Ingredients: []string{"Saxophone de Lisa", "Donut magique"},
		Result:      "Saxophone Enchanté",
		Description: "Un saxophone imprégné de magie de donut",
	},
	{
		Name:        "Bouclier Renforcé",
		Ingredients: []string{"Bouclier fait maison", "Batte de baseball"},
		Result:      "Bouclier de Springfield",
		Description: "Protection ultime fabriquée avec les matériaux de Springfield",
	},
	{
		Name:        "Biberon Ultime",
		Ingredients: []string{"Biberon de Maggie", "Donut magique", "Carte Itchy & Scratchy"},
		Result:      "Biberon Cosmique",
		Description: "Le biberon le plus puissant de l'univers",
	},
	{
		Name:        "Saxophone Spirituel",
		Ingredients: []string{"Saxophone de Lisa", "Note de Jazz Éternelle"},
		Result:      "Saxophone Spirituel",
		Description: "Un saxophone béni par l'esprit du jazz",
	},
	{
		Name:        "Super Biberon",
		Ingredients: []string{"Biberon de Maggie", "Biberon de Respect"},
		Result:      "Super Biberon",
		Description: "Le biberon ultime des bébés champions",
	},
}

func CraftMenuInShop(c *Character, textDelay time.Duration) {
	for {
		typeWriter("\n🔧 ATELIER DE BRICOLAGE DAPU", textDelay)
		typeWriter("==============================", textDelay)
		typeWriter("🏪 Apu : Bienvenue dans mon petit atelier ! J'ai appris quelques trucs...", textDelay)
		typeWriter("🏪 Apu : Avec les bons ingrédients, je peux créer des objets uniques !", textDelay)
		typeWriter("", textDelay)

		typeWriter("1. 📋 Voir les recettes disponibles", textDelay)
		typeWriter("2. 🔨 Créer un objet", textDelay)
		typeWriter("0. 🚪 Retour au Kwik-E-Mart", textDelay)
		typeWriter("", textDelay)

		var choice string
		ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
		fmt.Scan(&choice)

		switch choice {
		case "1":
			ShowCraftRecipesInShop(c, textDelay)
		case "2":
			CraftItemInShop(c, textDelay)
		case "0":
			return
		default:
			typeWriter("❌ Choix invalide.", textDelay)
		}
	}
}

func ShowCraftRecipesInShop(c *Character, textDelay time.Duration) {
	typeWriter("📋 RECETTES DE CRAFT DISPONIBLES", textDelay)
	typeWriter("=================================", textDelay)
	typeWriter("", textDelay)

	for i, recipe := range craftRecipes {
		typeWriter(fmt.Sprintf("%d. %s", i+1, recipe.Name), textDelay)
		typeWriter(fmt.Sprintf("   📝 %s", recipe.Description), textDelay)
		typeWriter("   🧪 Ingrédients :", textDelay)

		for _, ingredient := range recipe.Ingredients {
			hasIngredient := contains(c.inventory, ingredient)
			if hasIngredient {
				typeWriter(fmt.Sprintf("     ✅ %s", ingredient), textDelay)
			} else {
				typeWriter(fmt.Sprintf("     ❌ %s", ingredient), textDelay)
			}
		}

		typeWriter(fmt.Sprintf("   🎁 Résultat : %s", recipe.Result), textDelay)
		typeWriter("", textDelay)
	}
}

func CraftItemInShop(c *Character, textDelay time.Duration) {
	typeWriter("🔨 CRÉATION DOBJET", textDelay)
	typeWriter("===================", textDelay)
	typeWriter("", textDelay)

	availableRecipes := []CraftRecipe{}
	for _, recipe := range craftRecipes {
		canCraft := true
		for _, ingredient := range recipe.Ingredients {
			if !contains(c.inventory, ingredient) {
				canCraft = false
				break
			}
		}

		if canCraft {
			availableRecipes = append(availableRecipes, recipe)
			typeWriter(fmt.Sprintf("%d. %s", len(availableRecipes), recipe.Name), textDelay)
			typeWriter(fmt.Sprintf("   📝 %s", recipe.Description), textDelay)
		}
	}

	if len(availableRecipes) == 0 {
		typeWriter("❌ Aucune recette disponible avec vos ingrédients actuels.", textDelay)
		typeWriter("🏪 Apu : Il vous faut plus dingrédients, mon ami !", textDelay)
		typeWriter("🏪 Apu : Battez-vous, parlez aux gens, explorez Springfield !", textDelay)
		return
	}

	typeWriter("0. Annuler", textDelay)
	typeWriter("", textDelay)
	typeWriter("👉 Quelle recette voulez-vous utiliser ?", textDelay)

	var choice int
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(availableRecipes) {
		typeWriter("❌ Choix invalide.", textDelay)
		return
	}

	selectedRecipe := availableRecipes[choice-1]

	if !limitedInventory(c) {
		return
	}

	typeWriter(fmt.Sprintf("🔨 Voulez-vous vraiment créer : %s ?", selectedRecipe.Name), textDelay)
	typeWriter("1. Oui", textDelay)
	typeWriter("2. Non", textDelay)

	var confirm int
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&confirm)

	if confirm != 1 {
		typeWriter("❌ Création annulée.", textDelay)
		return
	}

	for _, ingredient := range selectedRecipe.Ingredients {
		for i, item := range c.inventory {
			if item == ingredient {
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				break
			}
		}
	}

	c.inventory = append(c.inventory, selectedRecipe.Result)

	typeWriter("✨ BRICOLAGE RÉUSSI !", textDelay)
	typeWriter(fmt.Sprintf("🏪 Apu : Magnifique ! Vous avez créé : %s !", selectedRecipe.Result), textDelay)
	typeWriter(fmt.Sprintf("📝 %s", selectedRecipe.Description), textDelay)
	typeWriter("🏪 Apu : Mes talents de bricoleur ne cessent de métonner !'", textDelay)
	typeWriter("", textDelay)
}
