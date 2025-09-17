package projetred

import (
	"fmt"
	"math/rand"
)

func Shopkeeper(c *Character) Character {
	randomvalue := 50 + rand.Intn(101)
	for {
		fmt.Println("\n====================")
		fmt.Println("   🛒 Boutique du Marchand")
		fmt.Println("====================")
		fmt.Println("1. Acheter un objet")
		fmt.Println("2. Vendre un objet")
		fmt.Println("3. Améliorer un objet")
		fmt.Println("4. Quitter la boutique")
		fmt.Println("--------------------")
		fmt.Println("💰 Vous disposez de", c.gold, "pièces d'or.")
		fmt.Print("👉 Que voulez-vous faire ? ")

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			fmt.Println("\n--- Objets disponibles à l'achat ---")
			fmt.Println("1. Épée (50 pièces d'or)")
			fmt.Println("2. Bouclier (40 pièces d'or)")
			fmt.Println("3. Potion de soin (10 pièces d'or)")
			fmt.Println("4. Bitcoin (", randomvalue, " pièces d'or)")
			fmt.Println("5. Retour au menu précédent")
			fmt.Println("------------------------------------")
			fmt.Println("💰 Or actuel :", c.gold)
			fmt.Print("👉 Quel objet souhaitez-vous acheter ? ")

			var choice2 string
			fmt.Scan(&choice2)

			switch choice2 {
			case "1":
				if c.gold >= 50 {
					c.inventory = append(c.inventory, "Épée")
					c.gold -= 50
					fmt.Println("✅ Vous avez acheté une Épée. Il vous reste", c.gold, "or.")
				} else {
					fmt.Println("❌ Pas assez d'or !")
				}
			case "2":
				if c.gold >= 40 {
					c.inventory = append(c.inventory, "Bouclier")
					c.gold -= 40
					fmt.Println("✅ Vous avez acheté un Bouclier. Il vous reste", c.gold, "or.")
				} else {
					fmt.Println("❌ Pas assez d'or !")
				}
			case "3":
				if c.gold >= 10 {
					c.inventory = append(c.inventory, "Potion")
					c.gold -= 10
					fmt.Println("✅ Vous avez acheté une Potion. Il vous reste", c.gold, "or.")
				} else {
					fmt.Println("❌ Pas assez d'or !")
				}
			case "4":
				if c.gold >= randomvalue {
					c.inventory = append(c.inventory, "Bitcoin")
					c.gold -= randomvalue
					fmt.Println("✅ Vous avez acheté un Bitcoin. Il vous reste", c.gold, "or.")
				} else {
					fmt.Println("❌ Pas assez d'or !")
				}
			case "5":
				fmt.Println("↩ Retour au menu de la boutique.")
			default:
				fmt.Println("❌ Choix invalide.")
			}

		case "2":
			if len(c.inventory) == 0 {
				fmt.Println("⚠ Vous n'avez aucun objet à vendre.")
				continue
			}
			fmt.Println("\n--- Objets que vous pouvez vendre ---")
			for i, v := range c.inventory {
				fmt.Printf("%d. %s\n", i+1, v)
			}
			fmt.Print("👉 Entrez le numéro de l'objet à vendre : ")

			var sellChoice int
			fmt.Scan(&sellChoice)
			if sellChoice < 1 || sellChoice > len(c.inventory) {
				fmt.Println("❌ Choix invalide.")
				continue
			}

			itemToSell := c.inventory[sellChoice-1]
			var sellPrice int
			switch itemToSell {
			case "Épée":
				sellPrice = 25
			case "Bouclier":
				sellPrice = 20
			case "Potion":
				sellPrice = 5
			case "Bitcoin":
				sellPrice = randomvalue
			default:
				sellPrice = 0
			}
			c.inventory = append(c.inventory[:sellChoice-1], c.inventory[sellChoice:]...)
			c.gold += sellPrice
			fmt.Printf("💸 Vous avez vendu %s pour %d or. Nouveau solde : %d or.\n", itemToSell, sellPrice, c.gold)

		case "3":
			if len(c.inventory) == 0 {
				fmt.Println("⚠ Vous n'avez aucun objet à améliorer.")
				continue
			}
			fmt.Println("\n--- Objets disponibles à l'amélioration ---")
			for i, v := range c.inventory {
				fmt.Printf("%d. %s\n", i+1, v)
			}
			fmt.Print("👉 Entrez le numéro de l'objet à améliorer : ")

			var upgradeChoice int
			fmt.Scan(&upgradeChoice)
			if upgradeChoice < 1 || upgradeChoice > len(c.inventory) {
				fmt.Println("❌ Choix invalide.")
				continue
			}

			itemToUpgrade := c.inventory[upgradeChoice-1]
			var upgradeCost int
			switch itemToUpgrade {
			case "Épée":
				upgradeCost = 30
			case "Bouclier":
				upgradeCost = 25
			default:
				fmt.Println("⚠ Cet objet ne peut pas être amélioré.")
				continue
			}

			if c.gold >= upgradeCost {
				c.gold -= upgradeCost
				c.inventory[upgradeChoice-1] = itemToUpgrade + " +1"
				fmt.Printf("✨ Votre %s a été amélioré pour %d or !\n", itemToUpgrade, upgradeCost)
			} else {
				fmt.Println("❌ Pas assez d'or pour améliorer cet objet.")
			}

		case "4":
			fmt.Println("\n👋 Merci de votre visite, voyageur. À bientôt !")
			Menu(*c)
			return *c

		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}
