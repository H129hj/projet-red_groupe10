package projetred

import (
	"fmt"
	"time"
)

// Arène 1 : Bar de Moe
func MoeBar(c *Character, progress *ScenarioProgress) {
	typeWriter("🍺 ÉTAPE 2 : ARÈNE 1 - BAR DE MOE", 50*time.Millisecond)
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🏪 Vous poussez les portes battantes du bar le plus miteux de Springfield...", 40*time.Millisecond)
	typeWriter("💨 Une odeur de bière rance et de désespoir vous accueille.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👨‍🦲 Moe : 'Hé ! Qu'est-ce qu'un gamin fait dans mon établissement ?'", 40*time.Millisecond)
	typeWriter("👨‍🦲 Moe : 'Attends... tu es un des gosses Simpson, pas vrai ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	switch c.class {
	case "bart":
		typeWriter("🎯 Bart : 'Salut Moe ! Je cherche mon père, il a disparu !'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Moe : 'Homer ? Ouais, il était là ce matin... bizarre d'ailleurs.'", 40*time.Millisecond)
	case "lisa":
		typeWriter("🎷 Lisa : 'Excusez-moi Monsieur Moe, je mène une enquête sur la disparition de mon père.'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Moe : 'Waouh, la petite Einstein ! Ouais, Homer était là...'", 40*time.Millisecond)
	case "maggie":
		typeWriter("👶 Maggie : '*suce sa tétine en regardant Moe intensément*'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Moe : 'Euh... le bébé me fait flipper... Bon, Homer était là ce matin.'", 40*time.Millisecond)
	}

	typeWriter("", 30*time.Millisecond)
	typeWriter("👨‍🦲 Moe : 'Il parlait d'un plan pour avoir de la bière gratuite à vie...'", 40*time.Millisecond)
	typeWriter("👨‍🦲 Moe : 'Quelque chose à propos d'un \"concours secret\" au parc d'attractions.'", 40*time.Millisecond)
	typeWriter("👨‍🦲 Moe : 'Si tu veux retrouver ton père, suis la trace des verres de bière vides...'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Déclencheur du combat selon le scénario
	typeWriter("🍺 Soudain, Barney Gumble se lève de son tabouret...", 40*time.Millisecond)
	typeWriter("🍺 Barney : '*rot* Hé toi ! Tu cherches Homer ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	var choice int
	typeWriter("💭 Comment répondre à Barney ?", 30*time.Millisecond)
	typeWriter("1. 😊 'Oui, peux-tu m'aider s'il te plaît ?'", 30*time.Millisecond)
	typeWriter("2. 😤 'J'aime pas ta tête, Barney !'", 30*time.Millisecond)
	typeWriter("3. 🤔 'Tu as des informations utiles ?'", 30*time.Millisecond)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Réponse polie
		typeWriter("🍺 Barney : '*rot* Ouais, Homer avait un ticket... quelque chose avec des donuts...'", 40*time.Millisecond)
		typeWriter("🍺 Barney : 'Il parlait du Comic Book Guy... *rot* ...et d'un concours...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer avait un ticket lié aux donuts et au Comic Book Guy !", 40*time.Millisecond)
		progress.HasClue2 = true
		progress.MoeCompleted = true
		progress.Stage = 3

	case 2:
		// Combat déclenché !
		typeWriter("🍺 Barney : '*rot* QUOI ?! Personne insulte Barney Gumble !'", 40*time.Millisecond)
		typeWriter("👨‍🦲 Moe : 'Oh non... voilà que ça recommence...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("⚔️ Barney entre en mode IVROGNE FURIEUX !", 50*time.Millisecond)

		// Combat contre Barney
		barneyMonster := InitMonster("Barney Gumble (Ivrogne Furieux)", 90, 22)
		characterTurnBarney(c, &barneyMonster, 1, progress)
		return

	case 3:
		// Réponse neutre
		typeWriter("🍺 Barney : '*rot* Homer... oui... il avait un papier bizarre...'", 40*time.Millisecond)
		typeWriter("🍺 Barney : 'Quelque chose sur un concours de donuts... *rot*'", 40*time.Millisecond)
		typeWriter("🍺 Barney : 'Va voir le Comic Book Guy, il saura peut-être...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer avait un papier sur un concours de donuts !", 40*time.Millisecond)
		progress.HasClue2 = true
		progress.MoeCompleted = true
		progress.Stage = 3

	default:
		typeWriter("❌ Choix invalide.", 30*time.Millisecond)
		MoeBar(c, progress)
		return
	}

	typeWriter("", 30*time.Millisecond)
	typeWriter("👨‍🦲 Moe : 'Bon, maintenant sortez de mon bar ! C'est pas un lieu pour les gosses !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	ScenarioMenu(c, progress)
}

// Combat spécialisé contre Barney
func characterTurnBarney(c *Character, m *Monster, t int, progress *ScenarioProgress) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter("🎉 Victoire ! Barney s'effondre sur son tabouret...", 40*time.Millisecond)
		typeWriter("🍺 Barney : '*rot* D'accord... d'accord... tu gagnes...'", 40*time.Millisecond)
		typeWriter("🍺 Barney : 'Tiens... *rot* ...Homer a laissé tomber ça...'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("📋 Barney vous tend un ticket froissé : 'CONCOURS DONUT GÉANT - Comic Book Store'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Ticket de concours menant au magasin de BD !", 40*time.Millisecond)
		progress.HasClue2 = true
		progress.MoeCompleted = true
		progress.Stage = 3
		ScenarioMenu(c, progress)
	} else {
		typeWriter("⚔️ À votre tour ! Choisissez une action :", 30*time.Millisecond)
		typeWriter("1. 💥 Attaquer", 30*time.Millisecond)
		typeWriter("2. 🎒 Fouiller dans votre sac", 30*time.Millisecond)
		typeWriter("3. 🏃 Fuir le combat", 30*time.Millisecond)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			attackMonster(c, m)
			barneyPattern(m, turn)
			characterTurnBarney(c, m, turn+1, progress)
		case 2:
			AccessInventory(*c)
			characterTurnBarney(c, m, turn, progress)
		case 3:
			typeWriter("🏃💨 Vous fuyez le bar en évitant les chopes volantes !", 40*time.Millisecond)
			ScenarioMenu(c, progress)
		default:
			typeWriter("❌ Choix invalide.", 30*time.Millisecond)
			characterTurnBarney(c, m, turn, progress)
		}
	}
}

// Attaques spéciales de Barney
func barneyPattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("💨 Barney utilise 'ROT DESTRUCTEUR' !", 40*time.Millisecond)
		typeWriter("🍺 Barney : '*ROOOOOOOT* Ça c'est de la Duff premium !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("🤢 Dégâts toxiques : %d points !", damage), 30*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 8
		typeWriter("🍺 Barney utilise 'JET DE CHOPE' !", 40*time.Millisecond)
		typeWriter("🍺 Barney : 'Prends ça ! *rot*'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("💥 Dégâts alcoolisés : %d points !", damage), 30*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("🥴 Barney utilise 'CHARGE TITUBANTE' !", 40*time.Millisecond)
		typeWriter("🍺 Barney : '*rot* Viens te battre comme un homme !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("🍻 Dégâts d'ivrogne : %d points !", damage), 30*time.Millisecond)
	}
}
