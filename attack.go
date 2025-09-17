package projetred

import "fmt"

type Skill struct {
	name_skill  string
	power_skill int
	cost_skill  int
}

func Attack(c1 *Character, c2 *Character) {
	var choice int
	switch choice {
	case 1:
		fmt.Println(c1.name, "attaque", c2.name, "avec une attaque basique!")
		c2.PV -= c1.power
		if c2.PV < 0 {
			c2.PV = 0
		}
		fmt.Println(c2.name, "a maintenant", c2.PV, "PV.")
	case 2:
		fmt.Println(c1.name, "utilise une compétence!")
		// Ici, vous pouvez ajouter la logique pour utiliser une compétence
	default:
		fmt.Println("Choix invalide, attaque basique par defaut.")
		c2.PV -= c1.power
		if c2.PV < 0 {
			c2.PV = 0
		}
		fmt.Println(c2.name, "a maintenant", c2.PV, "PV.")
	}
	if isdead(*c2) {
		Wasted(c2)
	}
}
