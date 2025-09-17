package projetred

import (
	"fmt"
	"math/rand"
	"time"
)

func Shopkeeper(c *Character) Character {
	randomvalue := 50 + rand.Intn(101)
	textDelay := 20 * time.Millisecond

	for {
		typeWriter("\n====================", textDelay)
		typeWriter("   🏪 Kwik-E-Mart d'Apu", textDelay)
		typeWriter("====================", textDelay)
		typeWriter("🇮🇳 Apu : 'Bonjour mon ami ! Que puis-je faire pour vous ?'", textDelay)
		typeWriter("1. Acheter un objet", textDelay)
		typeWriter("2. Vendre un objet", textDelay)
		typeWriter("3. Améliorer un objet", textDelay)
		typeWriter("0. Quitter le magasin", textDelay)
		typeWriter("--------------------", textDelay)
		typeWriter(fmt.Sprintf("💰 Vous avez %d dollars dans votre tirelire.", c.gold), textDelay)
		typeWriter("👉 Que voulez-vous faire ? ", textDelay)

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			typeWriter("\n--- 🛍️ Objets en vente chez Apu ---", textDelay)
			typeWriter("1. Batte de baseball de Bart (50 dollars)", textDelay)
			typeWriter("2. Bouclier fait maison (40 dollars)", textDelay)
			typeWriter("3. Donut magique de chez Homer (10 dollars)", textDelay)
			typeWriter(fmt.Sprintf("4. Carte rare Itchy & Scratchy (%d dollars)", randomvalue), textDelay)
			typeWriter("5. Agrandir l'inventaire (25 dollars)", textDelay)
			typeWriter("6. Retour au menu d'Apu", textDelay)
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
				if c.gold >= 40 && !contains(c.inventory, "Batte de baseball") && limitedInventory(c){
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
				if c.gold >= 25 {
					c.extendedInventory += 1
					c.gold -= 25
					typeWriter(fmt.Sprintf("✅ Apu : 'Votre inventaire a été agrandi ! Il vous reste %d dollars.'", c.gold), textDelay)
				} else {
					typeWriter("❌ Apu : 'Pas assez d'argent pour agrandir votre inventaire !'", textDelay)
				}
			case "6":
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

		case "0":
			typeWriter("\n👋 Apu : 'Merci pour votre visite ! Revenez vite au Kwik-E-Mart !'", textDelay)
			Menu(*c)
			return *c

		default:
			typeWriter("❌ Apu : 'Je ne comprends pas votre choix, mon ami !'", textDelay)
		}
	}
}
