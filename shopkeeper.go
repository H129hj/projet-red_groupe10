package projetred

import (
	"fmt"
	"math/rand"
	"time"
)

const textDelay = 25 * time.Millisecond

func Shopkeeper(c *Character) {
	for {
		var choice string
		var choice2 string
		var randomvalue int = 50 + rand.Intn(101)

		typeWriter("Bienvenue dans ma boutique! Que souhaitez-vous faire?", textDelay)
		typeWriter("1. Acheter un objet", textDelay)
		typeWriter("2. Vendre un objet", textDelay)
		typeWriter("3. Ameliorer un objet", textDelay)
		typeWriter("4. Quitter la boutique", textDelay)
		fmt.Scan(&choice)

		switch choice {
		case "1":
			typeWriter("Voici les objets disponibles a l'achat:", textDelay)
			typeWriter("1. - Epee (50 pieces d'or)", textDelay)
			typeWriter("2. - Bouclier (40 pieces d'or)", textDelay)
			typeWriter("3. - Potion de soin (10 pieces d'or)", textDelay)
			typeWriter(fmt.Sprintf("4. - Bitcoin (%d pieces d'or)", randomvalue), textDelay)
			typeWriter(fmt.Sprintf("Vous disposez de %d pieces d'or.", c.gold), textDelay)
			typeWriter("Quel objet souhaitez-vous acheter?", textDelay)
			fmt.Scan(&choice2)

			switch choice2 {
			case "1":
				if c.gold >= 50 && !limitedInventory(c) {
					c.inventory = append(c.inventory, "Epee")
					c.gold -= 50
					typeWriter("Vous avez achete une Epee.", textDelay)
				} else {
					typeWriter("Vous n'avez pas assez de pieces d'or.", textDelay)
				}
			case "2":
				if c.gold >= 40 && !limitedInventory(c){
					c.inventory = append(c.inventory, "Bouclier")
					c.gold -= 40
					typeWriter("Vous avez achete un Bouclier.", textDelay)
				} else {
					typeWriter("Vous n'avez pas assez de pieces d'or.", textDelay)
				}
			case "3":
				if c.gold >= 10 && !limitedInventory(c){
					c.inventory = append(c.inventory, "Potion")
					c.gold -= 10
					typeWriter("Vous avez achete une Potion de soin.", textDelay)
				} else {
					typeWriter("Vous n'avez pas assez de pieces d'or.", textDelay)
				}
			case "4":
				if c.gold >= randomvalue && !limitedInventory(c){
					c.inventory = append(c.inventory, "Bitcoin")
					c.gold -= randomvalue
					typeWriter("Vous avez achete un Bitcoin.", textDelay)
				} else {
					typeWriter("Vous n'avez pas assez de pieces d'or.", textDelay)
				}
			default:
				typeWriter("Objet invalide.", textDelay)
			}

		case "2":
			if len(c.inventory) == 0 {
				typeWriter("Vous n'avez aucun objet a vendre.", textDelay)
				continue
			}
			typeWriter("Voici les objets que vous pouvez vendre:", textDelay)
			for i, v := range c.inventory {
				typeWriter(fmt.Sprintf("%d. %s", i+1, v), textDelay)
			}
			typeWriter("Quel objet souhaitez-vous vendre? (entrez le numero)", textDelay)
			var sellChoice int
			fmt.Scan(&sellChoice)
			if sellChoice < 1 || sellChoice > len(c.inventory) {
				typeWriter("Choix invalide.", textDelay)
				continue
			}
			itemToSell := c.inventory[sellChoice-1]
			var sellPrice int
			switch itemToSell {
			case "Epee":
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
			typeWriter(fmt.Sprintf("Vous avez vendu %s pour %d pieces d'or.", itemToSell, sellPrice), textDelay)

		case "3":
			if len(c.inventory) == 0 {
				typeWriter("Vous n'avez aucun objet a ameliorer.", textDelay)
				continue
			}
			typeWriter("Voici les objets que vous pouvez ameliorer:", textDelay)
			for i, v := range c.inventory {
				typeWriter(fmt.Sprintf("%d. %s", i+1, v), textDelay)
			}
			typeWriter("Quel objet souhaitez-vous ameliorer? (entrez le numero)", textDelay)
			var upgradeChoice int
			fmt.Scan(&upgradeChoice)
			if upgradeChoice < 1 || upgradeChoice > len(c.inventory) {
				typeWriter("Choix invalide.", textDelay)
				continue
			}
			itemToUpgrade := c.inventory[upgradeChoice-1]
			var upgradeCost int
			switch itemToUpgrade {
			case "Epee":
				upgradeCost = 30
			case "Bouclier":
				upgradeCost = 25
			default:
				upgradeCost = 0
			}
			if upgradeCost == 0 {
				typeWriter("Cet objet ne peut pas etre ameliore.", textDelay)
				continue
			}
			if c.gold >= upgradeCost {
				c.gold -= upgradeCost
				typeWriter(fmt.Sprintf("Vous avez ameliore %s pour %d pieces d'or.", itemToUpgrade, upgradeCost), textDelay)
			} else {
				typeWriter("Vous n'avez pas assez de pieces d'or.", textDelay)
			}

		case "4":
			typeWriter("Merci de votre visite! A bientot.", textDelay)
			Menu(*c)
			return

		default:
			typeWriter("Choix invalide.", textDelay)
		}
	}
}
