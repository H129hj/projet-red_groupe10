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


	switch c.class {
	case "bart":
		BartVsTahitiBob(c, progress)
	case "lisa":
		LisaVsGencivesSanglantes(c, progress)
	case "maggie":
		MaggieVsBebeFurieux(c, progress)
	default:

		BartVsTahitiBob(c, progress)
	}
}

func BartVsTahitiBob(c *Character, progress *ScenarioProgress) {
	typeWriter("🌴 Soudain, une voix sinistre résonne derrière vous...", 15*time.Millisecond)
	typeWriter("🔪 ??? : 'Enfin ! L'heure de ma vengeance a sonné, Bart Simpson !'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter("🌴 Tahiti Bob sort de derrière une attraction, ses cheveux rouges flottant au vent !", 15*time.Millisecond)
	typeWriter("🌴 Tahiti Bob : 'Cette fois, tu ne m'échapperas pas !'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("=======================", 15*time.Millisecond)
	typeWriter("⚔️ COMBAT : TAHITI BOB", 15*time.Millisecond)
	typeWriter("=======================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("🌴 Tahiti Bob : En garde !", 15*time.Millisecond)

	bobMonster := InitMonster("Tahiti Bob (Vengeance)", 120, 35)
	ScenarioCombat(c, &bobMonster, progress, tahitiBobPattern, "bob")
}

func LisaVsGencivesSanglantes(c *Character, progress *ScenarioProgress) {
	typeWriter("🎷 Soudain, une mélodie lugubre résonne dans l'air...", 15*time.Millisecond)
	typeWriter("👻 Une silhouette fantomatique apparaît avec un saxophone...", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter("🎺 Gencives Sanglantes : 'Lisa... mon ancienne élève...'", 15*time.Millisecond)
	typeWriter("👻 Gencives Sanglantes : 'Tu as abandonné le vrai jazz pour cette... pop music !'", 15*time.Millisecond)
	typeWriter("🎺 Gencives Sanglantes : 'Il est temps de te rappeler ce qu'est le VRAI jazz !'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("=======================================", 15*time.Millisecond)
	typeWriter("⚔️ COMBAT : GENCIVES SANGLANTES", 15*time.Millisecond)
	typeWriter("=======================================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("🎺 Gencives Sanglantes : 'Que la bataille musicale commence !'", 15*time.Millisecond)

	gencivesMonster := InitMonster("Gencives Sanglantes (Esprit du Jazz)", 110, 32)
	ScenarioCombat(c, &gencivesMonster, progress, gencivesSanglantesPattern, "gencives")
}

func MaggieVsBebeFurieux(c *Character, progress *ScenarioProgress) {
	typeWriter("👶 Soudain, un cri de bébé retentit avec une force surnaturelle...", 15*time.Millisecond)
	typeWriter("😡 Un autre bébé apparaît dans un landau high-tech...", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)
	typeWriter("👶 Bébé Furieux : 'WAAAAAAAHHHHH !'", 15*time.Millisecond)
	typeWriter("😡 Bébé Furieux : 'Goo goo ga ga ! (Tu n'es pas le bébé le plus fort !)'", 15*time.Millisecond)
	typeWriter("👶 Bébé Furieux : 'MAMA ! DADA ! FIGHT ! (Combat de bébés !)'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("===================================", 15*time.Millisecond)
	typeWriter("⚔️ COMBAT : BÉBÉ FURIEUX", 15*time.Millisecond)
	typeWriter("===================================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("👶 Bébé Furieux : 'GRRRRR GA GA !'", 15*time.Millisecond)

	bebeMonster := InitMonster("Bébé Furieux (Rival Mystérieux)", 100, 28)
	ScenarioCombat(c, &bebeMonster, progress, bebeFurieuxPattern, "bebe")
}

func TahitiBobFight(c *Character, progress *ScenarioProgress) {

	BartVsTahitiBob(c, progress)
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
	combatDelay := 0 * time.Millisecond
	if turn%4 == 0 {
		damage := m.power * 2
		typeWriter("🌴 Tahiti Bob utilise 'PIÈGE DE CHEVEUX ROUGES' !", combatDelay)
		typeWriter("🌴 Tahiti Bob : 'Mes cheveux magnifiques vont t'entraver !'", combatDelay)
		typeWriter(fmt.Sprintf("💇‍♂️ Dégâts capillaires : %d points !", damage), combatDelay)
	} else if turn%3 == 0 {
		damage := m.power + 15
		typeWriter("🔪 Tahiti Bob utilise 'COUTEAU DE CLOWN' !", combatDelay)
		typeWriter("🌴 Tahiti Bob : 'Cette lame a servi dans mes meilleurs spectacles !'", combatDelay)
		typeWriter(fmt.Sprintf("🎭 Dégâts théâtraux : %d points !", damage), combatDelay)
	} else if turn%2 == 0 {
		damage := m.power + 8
		typeWriter("🎭 Tahiti Bob utilise 'MONOLOGUE DRAMATIQUE' !", combatDelay)
		typeWriter("🌴 Tahiti Bob : 'Laissez-moi vous réciter du Shakespeare !'", combatDelay)
		typeWriter(fmt.Sprintf("📚 Dégâts culturels : %d points !", damage), combatDelay)
	} else {
		damage := m.power
		typeWriter("👊 Tahiti Bob utilise 'RAGE VENGERESSE' !", combatDelay)
		typeWriter("🌴 Tahiti Bob : 'Des années de frustration dans cette attaque !'", combatDelay)
		typeWriter(fmt.Sprintf("😡 Dégâts de vengeance : %d points !", damage), combatDelay)
	}
}

func gencivesSanglantesPattern(m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	if turn%4 == 0 {
		damage := m.power * 2
		typeWriter("🎺 Gencives Sanglantes utilise 'SOLO DE L'AU-DELÀ' !", combatDelay)
		typeWriter("👻 Gencives Sanglantes : 'Écoute la mélodie de l'éternité !'", combatDelay)
		typeWriter(fmt.Sprintf("🎵 Dégâts spirituels : %d points !", damage), combatDelay)
	} else if turn%3 == 0 {
		damage := m.power + 12
		typeWriter("🎷 Gencives Sanglantes utilise 'NOTES MAUDITES' !", combatDelay)
		typeWriter("👻 Gencives Sanglantes : 'Ces notes vont hanter tes rêves !'", combatDelay)
		typeWriter(fmt.Sprintf("🎶 Dégâts musicaux : %d points !", damage), combatDelay)
	} else if turn%2 == 0 {
		damage := m.power + 8
		typeWriter("💀 Gencives Sanglantes utilise 'SOUFFLE FANTOMATIQUE' !", combatDelay)
		typeWriter("👻 Gencives Sanglantes : 'Le jazz véritable ne meurt jamais !'", combatDelay)
		typeWriter(fmt.Sprintf("❄️ Dégâts glacés : %d points !", damage), combatDelay)
	} else {
		damage := m.power
		typeWriter("😢 Gencives Sanglantes utilise 'MÉLANCOLIE ÉTERNELLE' !", combatDelay)
		typeWriter("👻 Gencives Sanglantes : 'Tu ressens ma tristesse infinie...'", combatDelay)
		typeWriter(fmt.Sprintf("💔 Dégâts émotionnels : %d points !", damage), combatDelay)
	}
}

func bebeFurieuxPattern(m *Monster, turn int) {
	combatDelay := 0 * time.Millisecond
	if turn%4 == 0 {
		damage := m.power * 2
		typeWriter("👶 Bébé Furieux utilise 'CRI SUPERSONIQUE' !", combatDelay)
		typeWriter("😡 Bébé Furieux : 'WAAAAAAAHHHHHHHHH !'", combatDelay)
		typeWriter(fmt.Sprintf("🔊 Dégâts soniques : %d points !", damage), combatDelay)
	} else if turn%3 == 0 {
		damage := m.power + 10
		typeWriter("🍼 Bébé Furieux utilise 'BIBERON PROJECTILE' !", combatDelay)
		typeWriter("👶 Bébé Furieux : 'MAMA NO LIKE ! (Maman n'aime pas !)'", combatDelay)
		typeWriter(fmt.Sprintf("🎯 Dégâts de lait : %d points !", damage), combatDelay)
	} else if turn%2 == 0 {
		damage := m.power + 6
		typeWriter("🧸 Bébé Furieux utilise 'PELUCHE VOLANTE' !", combatDelay)
		typeWriter("😡 Bébé Furieux : 'TEDDY ATTACK ! (Attaque de nounours !)'", combatDelay)
		typeWriter(fmt.Sprintf("🐻 Dégâts câlins : %d points !", damage), combatDelay)
	} else {
		damage := m.power
		typeWriter("😭 Bébé Furieux utilise 'CAPRICE DESTRUCTEUR' !", combatDelay)
		typeWriter("👶 Bébé Furieux : 'NO NO NO NO !'", combatDelay)
		typeWriter(fmt.Sprintf("💢 Dégâts de colère : %d points !", damage), combatDelay)
	}
}