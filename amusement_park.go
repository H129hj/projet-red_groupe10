package projetred

import (
	"fmt"
	"time"
)

// Arène finale : Parc d'attractions
func AmusementPark(c *Character, progress *ScenarioProgress) {
	typeWriter("🎡 ÉTAPE 4 : ARÈNE FINALE - PARC D'ATTRACTIONS", 50*time.Millisecond)
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎢 Vous arrivez au parc d'attractions de Springfield...", 40*time.Millisecond)
	typeWriter("🎪 Des lumières clignotantes et une musique de fête résonnent partout.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎯 Vous apercevez une grande bannière : 'GRAND CONCOURS DU DONUT COSMIQUE'", 40*time.Millisecond)
	typeWriter("👽 Des silhouettes étranges semblent organiser l'événement...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Différence selon le personnage choisi
	switch c.class {
	case "bart":
		// Bart doit affronter Tahiti Bob
		typeWriter("🎯 Bart : 'Cool ! Un parc d'attractions ! Bon, où est papa ?'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🌴 Soudain, une voix sinistre résonne derrière vous...", 40*time.Millisecond)
		typeWriter("🔪 ??? : 'Enfin ! L'heure de ma vengeance a sonné, Bart Simpson !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🌴 Tahiti Bob sort de derrière une attraction, ses cheveux rouges flottant au vent !", 50*time.Millisecond)
		typeWriter("", 30*time.Millisecond)

		TahitiBobFight(c, progress)
		return

	case "lisa":
		// Lisa enquête méthodiquement
		typeWriter("🎷 Lisa : 'Intéressant... ce concours semble suspect.'", 40*time.Millisecond)
		typeWriter("🎷 Lisa : 'Des extraterrestres organisant un concours de donuts ? Il y a anguille sous roche.'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 Lisa mène son enquête avec méthode...", 40*time.Millisecond)

		LisaInvestigation(c, progress)
		return

	case "maggie":
		// Maggie utilise ses pouvoirs mystérieux
		typeWriter("👶 Maggie : '*suce sa tétine avec un regard déterminé*'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("👶 Maggie semble percevoir quelque chose que les autres ne voient pas...", 40*time.Millisecond)
		typeWriter("👶 Elle pointe du doigt une direction précise avec une confiance troublante.", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)

		MaggieIntuition(c, progress)
		return
	}
}

// Combat contre Tahiti Bob (spécifique à Bart)
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

// Enquête de Lisa
func LisaInvestigation(c *Character, progress *ScenarioProgress) {
	typeWriter("🔍 ENQUÊTE DE LISA", 50*time.Millisecond)
	typeWriter("==============================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎷 Lisa observe attentivement les alentours...", 40*time.Millisecond)
	typeWriter("🎷 Lisa : 'Ces \"organisateurs\" ne sont pas humains...'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👽 Soudain, deux silhouettes vertes apparaissent !", 40*time.Millisecond)
	typeWriter("👽 Kang : 'Oh non ! La petite humaine intelligente nous a repérés !'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'Peu importe ! Notre plan est presque accompli !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎷 Lisa : 'Je le savais ! Que voulez-vous à mon père ?'", 40*time.Millisecond)
	typeWriter("👽 Kang : 'Nous étudions la capacité humaine à ingérer des donuts !'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'Ton père est notre meilleur cobaye ! Il a mangé 47 donuts !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	FindHomer(c, progress)
}

// Intuition de Maggie
func MaggieIntuition(c *Character, progress *ScenarioProgress) {
	typeWriter("👶 INTUITION MYSTÉRIEUSE DE MAGGIE", 50*time.Millisecond)
	typeWriter("========================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👶 Maggie se dirige directement vers une attraction spécifique...", 40*time.Millisecond)
	typeWriter("👶 Elle semble savoir exactement où aller, comme si elle avait un sixième sens.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎪 Devant la 'Maison des Donuts Magiques', Maggie s'arrête.", 40*time.Millisecond)
	typeWriter("👶 Elle pointe du doigt l'entrée avec insistance.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👽 Deux aliens sortent, surpris :", 40*time.Millisecond)
	typeWriter("👽 Kang : 'Comment ce bébé a-t-il trouvé notre laboratoire secret ?!'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'Elle a des pouvoirs que nous ne comprenons pas !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	FindHomer(c, progress)
}

// Retrouver Homer (finale commune)
func FindHomer(c *Character, progress *ScenarioProgress) {
	typeWriter("🎉 HOMER RETROUVÉ !", 50*time.Millisecond)
	typeWriter("==============================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🍩 À l'intérieur de l'attraction, vous découvrez Homer...", 40*time.Millisecond)
	typeWriter("🍩 Il est assis devant une montagne de donuts, l'air béat.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👨‍🦲 Homer : 'Mmm... donuts cosmiques... *mâche bruyamment*'", 40*time.Millisecond)

	switch c.class {
	case "bart":
		typeWriter("🎯 Bart : 'Papa ! On te cherche partout !'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Homer : 'Bart ? Comment tu m'as trouvé ? J'avais un plan secret !'", 40*time.Millisecond)
	case "lisa":
		typeWriter("🎷 Lisa : 'Papa ! Maman s'inquiète ! Et ces aliens t'utilisent comme cobaye !'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Homer : 'Lisa ? Des aliens ? Je croyais que c'était juste un concours...'", 40*time.Millisecond)
	case "maggie":
		typeWriter("👶 Maggie : '*babille avec reproche*'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Homer : 'Maggie ?! Comment tu... ? Tu es trop intelligente pour un bébé !'", 40*time.Millisecond)
	}

	typeWriter("", 30*time.Millisecond)
	typeWriter("👽 Kang : 'Peu importe ! Notre expérience est terminée !'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'Nous avons assez de données sur la consommation humaine de donuts !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	StoryConclusion(c, progress)
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
		FindHomer(c, progress)
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
