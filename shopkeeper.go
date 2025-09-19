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
		typeWriter("\n====================", textDelay)
		typeWriter("   🏪 Kwik-E-Mart d'Apu", textDelay)
		typeWriter("====================", textDelay)
		typeWriter("Apu : 'Bonjour mon ami ! Que puis-je faire pour vous ?'", textDelay)
		typeWriter("1. Acheter un objet", textDelay)
		typeWriter("2. Vendre un objet", textDelay)
		typeWriter("3. Améliorer un objet", textDelay)
		typeWriter("4. Atelier de bricolage d'Apu", textDelay)
		typeWriter("0. Quitter le magasin", textDelay)
		typeWriter("--------------------", textDelay)
		typeWriter(fmt.Sprintf("💰 Vous avez %d dollars dans votre tirelire.", c.gold), textDelay)
		typeWriter("👉 Que voulez-vous faire ? ", textDelay)

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			typeWriter("\n--- 🛍️ Objets en vente chez Apu ---", textDelay)
			typeWriter("1. Batte de baseball de Bart (50 dollars) - Arme", textDelay)
			typeWriter("2. Bouclier fait maison (40 dollars) - Armure", textDelay)
			typeWriter("3. Donut magique de chez Homer (10 dollars) - Consommable", textDelay)
			typeWriter(fmt.Sprintf("4. Carte rare Itchy & Scratchy (%d dollars) - Matériau", randomvalue), textDelay)
			typeWriter("5. Kit de craft basique (75 dollars) - Matériaux", textDelay)
			typeWriter("6. Agrandir l'inventaire (25 dollars) - Amélioration", textDelay)
			typeWriter("7. Retour au menu d'Apu", textDelay)
			typeWriter("------------------------------------", textDelay)
			typeWriter(fmt.Sprintf("💰 Argent actuel : %d dollars", c.gold), textDelay)
			typeWriter("👉 Quel objet souhaitez-vous acheter ? ", textDelay)

			var choice2 string
			fmt.Scan(&choice2)

			switch choice2 {
			case "1":
				if c.gold >= 50 && !contains(c.inventory, "Batte de baseball") && limitedInventory(c) {
					c.inventory = append(c.inventory, "Batte de baseball")
					c.gold -= 50
					typeWriter(fmt.Sprintf("✅ Apu : 'Excellente arme ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : 'Désolé, pas assez d'argent mon ami !'", textDelay)
				}
			case "2":
				if c.gold >= 40 && !contains(c.inventory, "Batte de baseball") && limitedInventory(c) {
					c.inventory = append(c.inventory, "Bouclier fait maison")
					c.gold -= 40
					typeWriter(fmt.Sprintf("✅ Apu : 'Protection garantie ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : 'Vos poches sont vides comme le frigo des Simpson !'", textDelay)
				}
			case "3":
				if c.gold >= 10 && limitedInventory(c) {
					c.inventory = append(c.inventory, "Donut magique")
					c.gold -= 10
					typeWriter(fmt.Sprintf("✅ Apu : 'Mmm... donut magique ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : 'Même pas 10 dollars ? Allez voir Homer !'", textDelay)
				}
			case "4":
				if c.gold >= randomvalue && limitedInventory(c) {
					c.inventory = append(c.inventory, "Carte Itchy & Scratchy")
					c.gold -= randomvalue
					typeWriter(fmt.Sprintf("✅ Apu : 'Carte très rare ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : 'Cette carte coûte plus cher que le salaire de Homer !'", textDelay)
				}
			case "5":
				if c.gold >= 75 && limitedInventory(c) {
					// Ajouter des matériaux de craft
					c.inventory = append(c.inventory, "Matériau de base", "Matériau de base")
					c.gold -= 75
					typeWriter(fmt.Sprintf("✅ Apu : 'Kit de craft acheté ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else if !limitedInventory(c) {
					typeWriter("❌ Apu : 'Inventaire plein !'", textDelay)
				} else {
					typeWriter("❌ Apu : 'Pas assez d'argent pour le kit de craft !'", textDelay)
				}
			case "6":
				if c.gold >= 25 {
					c.extendedInventory += 1
					c.gold -= 25
					typeWriter(fmt.Sprintf("✅ Apu : 'Votre inventaire a été agrandi ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : 'Pas assez d'argent pour agrandir votre inventaire !'", textDelay)
				}
			case "7":
				typeWriter("↩ Apu : 'Très bien, très bien...'", textDelay)
			default:
				typeWriter("❌ Apu : 'Je ne comprends pas, parlez plus clairement !'", textDelay)
			}

		case "2":
			if len(c.inventory) == 0 {
				typeWriter("⚠ Apu : 'Vous n'avez rien à vendre, revenez plus tard !'", textDelay)
				continue
			}
			typeWriter("\n--- 💸 Apu rachète vos objets ---", textDelay)
			for i, v := range c.inventory {
				typeWriter(fmt.Sprintf("%d. %s", i+1, v), textDelay)
			}
			typeWriter("👉 Apu : 'Quel objet voulez-vous vendre ?' ", textDelay)

			var sellChoice int
			fmt.Scan(&sellChoice)
			if sellChoice < 1 || sellChoice > len(c.inventory) {
				typeWriter("❌ Apu : 'Numéro invalide, comptez mieux que ça !'", textDelay)
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
			typeWriter(fmt.Sprintf("💸 Apu : 'J'ai acheté votre %s pour %d dollars. Solde : %d dollars.'", itemToSell, sellPrice, c.gold), textDelay)

		case "3":
			if len(c.inventory) == 0 {
				typeWriter("⚠ Apu : 'Aucun objet à améliorer dans vos poches !'", textDelay)
				continue
			}
			typeWriter("\n--- 🔧 Atelier d'amélioration d'Apu ---", textDelay)
			for i, v := range c.inventory {
				typeWriter(fmt.Sprintf("%d. %s", i+1, v), textDelay)
			}
			typeWriter("👉 Apu : 'Quel objet voulez-vous que j'améliore ?' ", textDelay)

			var upgradeChoice int
			fmt.Scan(&upgradeChoice)
			if upgradeChoice < 1 || upgradeChoice > len(c.inventory) {
				typeWriter("❌ Apu : 'Mauvais numéro, essayez encore !'", textDelay)
				continue
			}

			itemToUpgrade := c.inventory[upgradeChoice-1]
			var upgradeCost int
			switch itemToUpgrade {
			case "Batte de baseball":
				upgradeCost = 30
			case "Bouclier fait maison":
				upgradeCost = 25
			default:
				typeWriter("⚠ Apu : 'Cet objet ne peut pas être amélioré, désolé !'", textDelay)
				continue
			}

			if c.gold >= upgradeCost {
				c.gold -= upgradeCost
				c.inventory[upgradeChoice-1] = itemToUpgrade + " +1"
				typeWriter(fmt.Sprintf("✨ Apu : 'Votre %s est maintenant amélioré pour %d dollars !'", itemToUpgrade, upgradeCost), textDelay)
			} else {
				typeWriter("❌ Apu : 'Pas assez d'argent pour cette amélioration !'", textDelay)
			}

		case "4":
			CraftMenuInShop(c, textDelay)

		case "0":
			typeWriter("\n👋 Apu : 'Merci pour votre visite ! Revenez vite au Kwik-E-Mart !'", textDelay)
			ScenarioMenu(c, &ScenarioProgress{})
			return *c

		default:
			typeWriter("❌ Apu : 'Je ne comprends pas votre choix, mon ami !'", textDelay)
		}
	}
}

// Recettes de craft intégrées
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
}

// Menu de craft intégré dans le shop
func CraftMenuInShop(c *Character, textDelay time.Duration) {
	for {
		typeWriter("\n🔧 ATELIER DE BRICOLAGE D'APU", textDelay)
		typeWriter("==============================", textDelay)
		typeWriter("🏪 Apu : 'Bienvenue dans mon petit atelier ! J'ai appris quelques trucs...'", textDelay)
		typeWriter("🏪 Apu : 'Avec les bons ingrédients, je peux créer des objets uniques !'", textDelay)
		typeWriter("", textDelay)

		typeWriter("1. 📋 Voir les recettes disponibles", textDelay)
		typeWriter("2. 🔨 Créer un objet", textDelay)
		typeWriter("0. 🚪 Retour au Kwik-E-Mart", textDelay)
		typeWriter("", textDelay)

		var choice string
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

// Afficher les recettes dans le shop
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

	typeWriter("Appuyez sur Entrée pour continuer...", textDelay)
	fmt.Scanln()
}

// Créer un objet dans le shop
func CraftItemInShop(c *Character, textDelay time.Duration) {
	typeWriter("🔨 CRÉATION D'OBJET", textDelay)
	typeWriter("===================", textDelay)
	typeWriter("", textDelay)

	// Trouver les recettes disponibles
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
		typeWriter("🏪 Apu : 'Il vous faut plus d'ingrédients, mon ami !'", textDelay)
		typeWriter("🏪 Apu : 'Battez-vous, parlez aux gens, explorez Springfield !'", textDelay)
		return
	}

	typeWriter("0. Annuler", textDelay)
	typeWriter("", textDelay)
	typeWriter("👉 Quelle recette voulez-vous utiliser ?", textDelay)

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(availableRecipes) {
		typeWriter("❌ Choix invalide.", textDelay)
		return
	}

	selectedRecipe := availableRecipes[choice-1]

	// Vérifier si l'inventaire a de la place
	if !limitedInventory(c) {
		return
	}

	// Confirmer la création
	typeWriter(fmt.Sprintf("🔨 Voulez-vous vraiment créer : %s ?", selectedRecipe.Name), textDelay)
	typeWriter("1. Oui", textDelay)
	typeWriter("2. Non", textDelay)

	var confirm int
	fmt.Scan(&confirm)

	if confirm != 1 {
		typeWriter("❌ Création annulée.", textDelay)
		return
	}

	// Retirer les ingrédients
	for _, ingredient := range selectedRecipe.Ingredients {
		for i, item := range c.inventory {
			if item == ingredient {
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				break
			}
		}
	}

	// Ajouter le résultat
	c.inventory = append(c.inventory, selectedRecipe.Result)

	typeWriter("✨ BRICOLAGE RÉUSSI !", textDelay)
	typeWriter(fmt.Sprintf("🏪 Apu : 'Magnifique ! Vous avez créé : %s !'", selectedRecipe.Result), textDelay)
	typeWriter(fmt.Sprintf("📝 %s", selectedRecipe.Description), textDelay)
	typeWriter("🏪 Apu : 'Mes talents de bricoleur ne cessent de m'étonner !'", textDelay)
	typeWriter("", textDelay)
}
