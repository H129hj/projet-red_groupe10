package projetred

func Equipementprotection(c *Character) int {
	var protection int
	for _, item := range c.inventory {
		switch item {
		case "Armure de bronzer":
			protection += 5
		case "Armure de fer":
			protection += 10
		case "Armure de or":
			protection += 15
		case "casque en bronze":
			protection += 3
		case "casque en fer":
			protection += 6
		case "casque en or":
			protection += 9
		case "Botte en bronze":
			protection += 2
		case "Botte en fer":
			protection += 4
		case "Botte en or":
			protection += 6
		}
	}
	return protection
}
