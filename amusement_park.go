package projetred

import (
	"fmt"
	"time"
)

// Arène finale : Parc d'attractions
func AmusementPark(c *Character, progress *ScenarioProgress) {
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("🎡 ÉTAPE 4 : ARÈNE FINALE - PARC D'ATTRACTIONS", 50*time.Millisecond)
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎢 Vous arrivez au parc d'attractions de Springfield...", 40*time.Millisecond)
	typeWriter("🎪 Des lumières clignotantes et une musique de fête résonnent partout.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎯 Vous apercevez une grande bannière : 'GRAND CONCOURS DU DONUT COSMIQUE'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎯 Bart : 'Cool ! Un parc d'attractions ! Bon, où est papa ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)
	typeWriter("🌴 Soudain, une voix sinistre résonne derrière vous...", 40*time.Millisecond)
	typeWriter("🔪 ??? : 'Enfin ! L'heure de ma vengeance a sonné, Bart Simpson !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)
	typeWriter("🌴 Tahiti Bob sort de derrière une attraction, ses cheveux rouges flottant au vent !", 50*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	TahitiBobFight(c, progress)
}


func TahitiBobFight(c *Character, progress *ScenarioProgress) {
	typeWriter("⚔️ COMBAT OBLIGATOIRE : TAHITI BOB", 50*time.Millisecond)
	typeWriter("========================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🌴 Tahiti Bob : 'Bart Simpson ! Depuis des années j'attends ce moment !'", 40*time.Millisecond)
	typeWriter("🌴 Tahiti Bob : 'Tes blagues stupides ont ruiné ma carrière d'acteur !'", 40*time.Millisecond)
	typeWriter("🌴 Tahiti Bob : 'Aujourd'hui, justice sera rendue !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎯 Bart : 'Oh oh... Tahiti Bob ! Euh... salut mec !'", 40*time.Millisecond)
	typeWriter("🎯 Bart : 'Écoute, je cherche juste mon père, on peut pas faire ça plus tard ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🌴 Tahiti Bob : 'JAMAIS ! En garde !'", 50*time.Millisecond)

	// Combat obligatoire
	bobMonster := InitMonster("Tahiti Bob (Vengeance)", 120, 35)
	characterTurnTahitiBob(c, &bobMonster, 1, progress)
}

// Combat spécialisé contre Tahiti Bob
func characterTurnTahitiBob(c *Character, m *Monster, t int, progress *ScenarioProgress) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter("🎉 Victoire ! Tahiti Bob s'effondre dramatiquement...", 40*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Impossible ! Vaincu par un gamin !'", 40*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Bon... ton père est dans la Maison des Donuts Magiques...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
	} else {
		typeWriter("⚔️ À votre tour ! Choisissez une action :", 30*time.Millisecond)
		typeWriter("1. 💥 Attaquer", 30*time.Millisecond)
		typeWriter("2. 🎒 Fouiller dans votre sac", 30*time.Millisecond)
		typeWriter("3. 🏃 Fuir le combat", 30*time.Millisecond)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			attackMonster(c, m)
			tahitiBobPattern(m, turn)
			characterTurnTahitiBob(c, m, turn+1, progress)
		case 2:
			AccessInventory(*c)
			characterTurnTahitiBob(c, m, turn, progress)
		case 3:
			typeWriter("🏃💨 Impossible de fuir ! Tahiti Bob bloque la sortie !", 40*time.Millisecond)
			typeWriter("🌴 Tahiti Bob : 'Tu ne m'échapperas pas cette fois !'", 40*time.Millisecond)
			characterTurnTahitiBob(c, m, turn, progress)
		default:
			typeWriter("❌ Choix invalide.", 30*time.Millisecond)
			characterTurnTahitiBob(c, m, turn, progress)
		}
	}
}

// Attaques spéciales de Tahiti Bob
func tahitiBobPattern(m *Monster, turn int) {
	if turn%4 == 0 {
		damage := m.power * 2
		typeWriter("🌴 Tahiti Bob utilise 'PIÈGE DE CHEVEUX ROUGES' !", 40*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Mes cheveux magnifiques vont t'entraver !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("💇‍♂️ Dégâts capillaires : %d points !", damage), 30*time.Millisecond)
	} else if turn%3 == 0 {
		damage := m.power + 15
		typeWriter("🔪 Tahiti Bob utilise 'COUTEAU DE CLOWN' !", 40*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Cette lame a servi dans mes meilleurs spectacles !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("🎭 Dégâts théâtraux : %d points !", damage), 30*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 8
		typeWriter("🎭 Tahiti Bob utilise 'MONOLOGUE DRAMATIQUE' !", 40*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Laissez-moi vous réciter du Shakespeare !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("📚 Dégâts culturels : %d points !", damage), 30*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("👊 Tahiti Bob utilise 'RAGE VENGERESSE' !", 40*time.Millisecond)
		typeWriter("🌴 Tahiti Bob : 'Des années de frustration dans cette attaque !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("😡 Dégâts de vengeance : %d points !", damage), 30*time.Millisecond)
	}
}
