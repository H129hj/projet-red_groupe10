package projetred

import (
	"fmt"
	"math/rand"
)

func Shopkeeper(c *Character) Character {
	randomvalue := 50 + rand.Intn(101)
	for {
		fmt.Println("\n====================")
		fmt.Println("   🏪 Kwik-E-Mart d'Apu")
		fmt.Println("====================")
		fmt.Println("🇮🇳 Apu : 'Bonjour mon ami ! Que puis-je faire pour vous ?'")
		fmt.Println("1. Acheter un objet")
		fmt.Println("2. Vendre un objet")
		fmt.Println("3. Améliorer un objet")
		fmt.Println("4. Quitter le magasin")
		fmt.Println("--------------------")
		fmt.Println("💰 Vous avez", c.gold, "dollars dans votre tirelire.")
		fmt.Print("👉 Que voulez-vous faire ? ")

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			fmt.Println("\n--- 🛍️ Objets en vente chez Apu ---")
			fmt.Println("1. Batte de baseball de Bart (50 dollars)")
			fmt.Println("2. Bouclier fait maison (40 dollars)")
			fmt.Println("3. Donut magique de chez Homer (10 dollars)")
			fmt.Println("4. Carte rare Itchy & Scratchy (", randomvalue, " dollars)")
			fmt.Println("5. Retour au menu d'Apu")
			fmt.Println("------------------------------------")
			fmt.Println("💰 Argent actuel :", c.gold, "dollars")
			fmt.Print("👉 Quel objet souhaitez-vous acheter ? ")

			var choice2 string
			fmt.Scan(&choice2)

			switch choice2 {
			case "1":
				if c.gold >= 50 {
					c.inventory = append(c.inventory, "Batte de baseball")
					c.gold -= 50
					fmt.Println("✅ Apu : 'Excellente arme ! Il vous reste", c.gold, "dollars.'")
				} else {
					fmt.Println("❌ Apu : 'Désolé, pas assez d'argent mon ami !'")
				}
			case "2":
				if c.gold >= 40 {
					c.inventory = append(c.inventory, "Bouclier fait maison")
					c.gold -= 40
					fmt.Println("✅ Apu : 'Protection garantie ! Il vous reste", c.gold, "dollars.'")
				} else {
					fmt.Println("❌ Apu : 'Vos poches sont vides comme le frigo des Simpson !'")
				}
			case "3":
				if c.gold >= 10 {
					c.inventory = append(c.inventory, "Donut magique")
					c.gold -= 10
					fmt.Println("✅ Apu : 'Mmm... donut magique ! Il vous reste", c.gold, "dollars.'")
				} else {
					fmt.Println("❌ Apu : 'Même pas 10 dollars ? Allez voir Homer !'")
				}
			case "4":
				if c.gold >= randomvalue {
					c.inventory = append(c.inventory, "Carte Itchy & Scratchy")
					c.gold -= randomvalue
					fmt.Println("✅ Apu : 'Carte très rare ! Il vous reste", c.gold, "dollars.'")
				} else {
					fmt.Println("❌ Apu : 'Cette carte coûte plus cher que le salaire de Homer !'")
				}
			case "5":
				fmt.Println("↩ Apu : 'Très bien, très bien...'")
			default:
				fmt.Println("❌ Apu : 'Je ne comprends pas, parlez plus clairement !'")
			}

		case "2":
			if len(c.inventory) == 0 {
				fmt.Println("⚠ Apu : 'Vous n'avez rien à vendre, revenez plus tard !'")
				continue
			}
			fmt.Println("\n--- 💸 Apu rachète vos objets ---")
			for i, v := range c.inventory {
				fmt.Printf("%d. %s\n", i+1, v)
			}
			fmt.Print("👉 Apu : 'Quel objet voulez-vous vendre ?' ")

			var sellChoice int
			fmt.Scan(&sellChoice)
			if sellChoice < 1 || sellChoice > len(c.inventory) {
				fmt.Println("❌ Apu : 'Numéro invalide, comptez mieux que ça !'")
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
			fmt.Printf("💸 Apu : 'J'ai acheté votre %s pour %d dollars. Solde : %d dollars.'\n", itemToSell, sellPrice, c.gold)

		case "3":
			if len(c.inventory) == 0 {
				fmt.Println("⚠ Apu : 'Aucun objet à améliorer dans vos poches !'")
				continue
			}
			fmt.Println("\n--- 🔧 Atelier d'amélioration d'Apu ---")
			for i, v := range c.inventory {
				fmt.Printf("%d. %s\n", i+1, v)
			}
			fmt.Print("👉 Apu : 'Quel objet voulez-vous que j'améliore ?' ")

			var upgradeChoice int
			fmt.Scan(&upgradeChoice)
			if upgradeChoice < 1 || upgradeChoice > len(c.inventory) {
				fmt.Println("❌ Apu : 'Mauvais numéro, essayez encore !'")
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
				fmt.Println("⚠ Apu : 'Cet objet ne peut pas être amélioré, désolé !'")
				continue
			}

			if c.gold >= upgradeCost {
				c.gold -= upgradeCost
				c.inventory[upgradeChoice-1] = itemToUpgrade + " +1"
				fmt.Printf("✨ Apu : 'Votre %s est maintenant amélioré pour %d dollars !'\n", itemToUpgrade, upgradeCost)
			} else {
				fmt.Println("❌ Apu : 'Pas assez d'argent pour cette amélioration !'")
			}

		case "4":
			fmt.Println("\n👋 Apu : 'Merci pour votre visite ! Revenez vite au Kwik-E-Mart !'")
			Menu(*c)
			return *c

		default:
			fmt.Println("❌ Apu : 'Je ne comprends pas votre choix, mon ami !'")
		}
	}
}
