package projetred

import (
	"fmt"
	"time"
)

func AmusementPark(c *Character, progress *ScenarioProgress) {
	typeWriter("======================", 15*time.Millisecond)
	typeWriter("🎡 PARC D'ATTRACTIONS", 15*time.Millisecond)
	typeWriter("======================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("🎢 Vous arrivez au parc d'attractions de Springfield...", 15*time.Millisecond)
	typeWriter("🎪 Des lumières clignotantes et une musique de fête résonnent partout.", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("🌴 Soudain, une voix sinistre résonne derrière vous...", 15*time.Millisecond)
	typeWriter("🔪 ??? : 'Enfin ! L'heure de ma vengeance a sonné, Bart Simpson !'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter("🌴 Tahiti Bob sort de derrière une attraction, ses cheveux rouges flottant au vent !", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	TahitiBobFight(c, progress)
}

func TahitiBobFight(c *Character, progress *ScenarioProgress) {
	typeWriter("=======================", 15*time.Millisecond)
	typeWriter("⚔️ COMBAT : TAHITI BOB", 15*time.Millisecond)
	typeWriter("=======================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("🌴 Tahiti Bob : En garde !", 15*time.Millisecond)

	bobMonster := InitMonster("Tahiti Bob (Vengeance)", 120, 35)
	ScenarioCombat(c, &bobMonster, progress, tahitiBobPattern, "bob")
}

func characterTurnTahitiBob(c *Character, m *Monster, t int, progress *ScenarioProgress) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter("🎉 Victoire ! Tahiti Bob s'effondre dramatiquement...", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Impossible ! Vaincu par un gamin !'", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Bon... ton père est dans la Maison des Donuts Magiques...'", 15*time.Millisecond)

		// Ajouter un ingrédient spécial après le boss final
		AddIngredient(c, "Matériau de base", "le parc d'attractions")

		typeWriter("", 15*time.Millisecond)
	} else {
		typeWriter("⚔️ À votre tour ! Choisissez une action :", 15*time.Millisecond)
		typeWriter("1. 💥 Attaquer", 15*time.Millisecond)
		typeWriter("2. 🎒 Fouiller dans votre sac", 15*time.Millisecond)
		typeWriter("3. 🏃 Fuir le combat", 15*time.Millisecond)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			attackMonster(c, m)
			tahitiBobPattern(m, turn)
			characterTurnTahitiBob(c, m, turn+1, progress)
		case 2:
			typeWriter(AccessInventory(*c), 15*time.Millisecond)
			characterTurnTahitiBob(c, m, turn, progress)
		case 3:
			typeWriter("🏃💨 Impossible de fuir ! Tahiti Bob bloque la sortie !", 15*time.Millisecond)
			typeWriter("🌴 Tahiti Bob : 'Tu ne m'échapperas pas cette fois !'", 15*time.Millisecond)
			characterTurnTahitiBob(c, m, turn, progress)
		default:
			typeWriter("❌ Choix invalide.", 15*time.Millisecond)
			characterTurnTahitiBob(c, m, turn, progress)
		}
	}
}

func tahitiBobPattern(m *Monster, turn int) {
	if turn%4 == 0 {
		damage := m.power * 2
		typeWriter("🌴 Tahiti Bob utilise 'PIÈGE DE CHEVEUX ROUGES' !", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Mes cheveux magnifiques vont t'entraver !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("💇‍♂️ Dégâts capillaires : %d points !", damage), 15*time.Millisecond)
	} else if turn%3 == 0 {
		damage := m.power + 15
		typeWriter("🔪 Tahiti Bob utilise 'COUTEAU DE CLOWN' !", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Cette lame a servi dans mes meilleurs spectacles !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("🎭 Dégâts théâtraux : %d points !", damage), 15*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 8
		typeWriter("🎭 Tahiti Bob utilise 'MONOLOGUE DRAMATIQUE' !", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Laissez-moi vous réciter du Shakespeare !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("📚 Dégâts culturels : %d points !", damage), 15*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("👊 Tahiti Bob utilise 'RAGE VENGERESSE' !", 15*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Des années de frustration dans cette attaque !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("😡 Dégâts de vengeance : %d points !", damage), 15*time.Millisecond)
	}
}
