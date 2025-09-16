package projetred

import (
	"fmt"
	"math/rand"
)

func Shopkeeper(c *Character) Character {
	var choice string
	var choice2 string
	var randomvalue int = 50 + rand.Intn(101)
	fmt.Println("Bienvenue dans ma boutique! Que souhaitez-vous faire?")
	fmt.Println("1. Acheter un objet")
	fmt.Println("2. Vendre un objet")
	fmt.Println("3. Ameliorer un objet")
	fmt.Println("4. Quitter la boutique")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		fmt.Println("Voici les objets disponibles a l'achat:")
		fmt.Println("1. - Epee (50 pieces d'or)")
		fmt.Println("2. - Bouclier (40 pieces d'or)")
		fmt.Println("3. - Potion de soin (10 pieces d'or)")
		fmt.Println("4. - Bitcoin", randomvalue, "pieces d'or)")
		fmt.Println("Quel objet souhaitez-vous acheter?")
		fmt.Println("Vous disposez de ", c.gold, "pieces d'or.")
		fmt.Scan(&choice2)
		switch choice2 {
		case "1":
			if c.gold >= 50 {
				c.inventory = append(c.inventory, "Epee")
				c.gold -= 50
				fmt.Println("Vous avez achete une Epee.")
			} else {
				fmt.Println("Vous n'avez pas assez de pieces d'or.")
				return Shopkeeper(c)
			}
		case "2":
			if c.gold >= 40 {
				c.inventory = append(c.inventory, "Bouclier")
				c.gold -= 40
				fmt.Println("Vous avez achete un Bouclier.")
			} else {
				fmt.Println("Vous n'avez pas assez de pieces d'or.")
			}
		case "3":
			if c.gold >= 10 {
				c.inventory = append(c.inventory, "Potion")
				c.gold -= 10
				fmt.Println("Vous avez achete une Potion de soin.")
			} else {
				fmt.Println("Vous n'avez pas assez de pieces d'or.")
			}
		case "4":
			if c.gold >= randomvalue {
				c.inventory = append(c.inventory, "Bitcoin")
				c.gold -= randomvalue
				fmt.Println("Vous avez achete un Bitcoin.")
			} else {
				fmt.Println("Vous n'avez pas assez de pieces d'or.")
			}
		default:
			fmt.Println("Objet invalide.")
		}
	case "2":
		if len(c.inventory) == 0 {
			fmt.Println("Vous n'avez aucun objet a vendre.")
			return Shopkeeper(c)
		}
		fmt.Println("Voici les objets que vous pouvez vendre:")
		for i, v := range c.inventory {
			fmt.Printf("%d. %s\n", i+1, v)
		}
		fmt.Println("Quel objet souhaitez-vous vendre? (entrez le numero)")
		var sellChoice int
		fmt.Scan(&sellChoice)
		if sellChoice < 1 || sellChoice > len(c.inventory) {
			fmt.Println("Choix invalide.")
			return Shopkeeper(c)
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
		fmt.Printf("Vous avez vendu %s pour %d pieces d'or.\n", itemToSell, sellPrice)
	case "3":
		if len(c.inventory) == 0 {
			fmt.Println("Vous n'avez aucun objet a ameliorer.")
			return Shopkeeper(c)
		}
		fmt.Println("Voici les objets que vous pouvez ameliorer:")
		for i, v := range c.inventory {
			fmt.Printf("%d. %s\n", i+1, v)
		}
		fmt.Println("Quel objet souhaitez-vous ameliorer? (entrez le numero)")
		var upgradeChoice int
		fmt.Scan(&upgradeChoice)
		if upgradeChoice < 1 || upgradeChoice > len(c.inventory) {
			fmt.Println("Choix invalide.")
			return Shopkeeper(c)
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
		if c.gold >= upgradeCost {
			c.gold -= upgradeCost
			fmt.Printf("Vous avez ameliore %s pour %d pieces d'or.\n", itemToUpgrade, upgradeCost)
		} else {
			fmt.Println("Vous n'avez pas assez de pieces d'or.")
		}
	case "4":
		fmt.Println("Merci de votre visite! A bientot.")
		return *c
	default:
		fmt.Println("Choix invalide.")
	}
	return *c
}
