package projetred

import (
	"fmt"
	"time"
)

// Arène 2 : Magasin de BD
func ComicBookStore(c *Character, progress *ScenarioProgress) {
	typeWriter("📚 ÉTAPE 3 : ARÈNE 2 - MAGASIN DE BD", 50*time.Millisecond)
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🏪 Vous entrez dans l'antre du Comic Book Guy...", 40*time.Millisecond)
	typeWriter("📚 Des milliers de BD, figurines et objets geek s'entassent partout.", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👨‍💻 Comic Book Guy : 'Ah ! Un client potentiel dans mon sanctuaire de la culture pop !'", 40*time.Millisecond)
	typeWriter("👨‍💻 Comic Book Guy : 'Que puis-je faire pour... attends, tu es un Simpson !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	switch c.class {
	case "bart":
		typeWriter("🎯 Bart : 'Salut mec ! Je cherche mon père, Homer Simpson.'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ah, le célèbre Homer ! Collectionneur de... donuts.'", 40*time.Millisecond)
	case "lisa":
		typeWriter("🎷 Lisa : 'Bonjour, je mène une enquête méthodique sur mon père disparu.'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Enfin quelqu'un avec un cerveau ! Contrairement à ton père...'", 40*time.Millisecond)
	case "maggie":
		typeWriter("👶 Maggie : '*regarde fixement une figurine de superhéros*'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ce bébé a du goût ! Elle regarde ma figurine rare de...'", 40*time.Millisecond)
	}

	typeWriter("", 30*time.Millisecond)

	var choice int
	typeWriter("💭 Comment aborder le Comic Book Guy ?", 30*time.Millisecond)
	typeWriter("1. 😊 'Pouvez-vous m'aider poliment, s'il vous plaît ?'", 30*time.Millisecond)
	typeWriter("2. 😤 'Dépêche-toi, j'ai pas toute la journée !'", 30*time.Millisecond)
	typeWriter("3. 🤓 'J'ai besoin de vos connaissances encyclopédiques.'", 30*time.Millisecond)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Réponse polie - Indice direct
		typeWriter("👨‍💻 Comic Book Guy : 'Ah ! Enfin de la politesse ! Très rafraîchissant.'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ton père est venu ce matin avec un ticket étrange...'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Il parlait d'un \"Grand Concours du Donut Cosmique\" au parc d'attractions.'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Apparemment, le gagnant obtient des donuts gratuits à vie !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer participe au Grand Concours du Donut Cosmique !", 40*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4

	case 2:
		// Réponse impolie - Combat optionnel !
		typeWriter("👨‍💻 Comic Book Guy : 'QUOI ?! L'audace ! L'impudence !'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Personne ne parle ainsi au maître de la culture geek !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)

		var fightChoice int
		typeWriter("💭 Le Comic Book Guy semble énervé...", 30*time.Millisecond)
		typeWriter("1. ⚔️ 'Alors on se bat ?'", 30*time.Millisecond)
		typeWriter("2. 😅 'Désolé, je me suis mal exprimé...'", 30*time.Millisecond)
		fmt.Scan(&fightChoice)

		if fightChoice == 1 {
			typeWriter("⚔️ Comic Book Guy entre en mode CRITIQUE ULTIME !", 50*time.Millisecond)
			comicGuyMonster := InitMonster("Comic Book Guy (Critique Ultime)", 85, 28)
			characterTurnComicGuy(c, &comicGuyMonster, 1, progress)
			return
		} else {
			typeWriter("👨‍💻 Comic Book Guy : 'Hmph ! Bon... ton père parlait du parc d'attractions.'", 40*time.Millisecond)
			typeWriter("👨‍💻 Comic Book Guy : 'Quelque chose à propos d'un concours de donuts...'", 40*time.Millisecond)
			typeWriter("", 30*time.Millisecond)
			typeWriter("🔍 INDICE OBTENU : Homer est au parc d'attractions pour un concours !", 40*time.Millisecond)
			progress.HasClue3 = true
			progress.ComicCompleted = true
			progress.Stage = 4
		}

	case 3:
		// Réponse flatteuse - Indice détaillé
		typeWriter("👨‍💻 Comic Book Guy : '*rougit* Mes connaissances encyclopédiques ! Enfin quelqu'un qui apprécie !'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ton père, Homer Simpson, est venu avec un ticket mystérieux...'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Il s'agit du légendaire \"Grand Concours du Donut Cosmique\"...'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Organisé par des forces... extraterrestres... au parc d'attractions !'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Le gagnant obtient des donuts magiques pour l'éternité !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Concours extraterrestre de donuts au parc d'attractions !", 40*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4

	default:
		typeWriter("❌ Choix invalide.", 30*time.Millisecond)
		ComicBookStore(c, progress)
		return
	}

	typeWriter("", 30*time.Millisecond)
	typeWriter("👨‍💻 Comic Book Guy : 'Maintenant va ! Et que la force soit avec toi... ou pas.'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	ScenarioMenu(c, progress)
}

// Combat optionnel contre Comic Book Guy
func characterTurnComicGuy(c *Character, m *Monster, t int, progress *ScenarioProgress) {
	var choice int
	turn := t
	if c.PV <= 0 {
		Wasted(c)
	} else if m.PV <= 0 {
		typeWriter("🎉 Victoire ! Comic Book Guy ajuste ses lunettes...", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Impossible ! J'ai été vaincu par... un amateur !'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Bon... ton père est au parc d'attractions...'", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Il participe au Grand Concours du Donut Cosmique !'", 40*time.Millisecond)
		typeWriter("", 30*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer au concours de donuts du parc d'attractions !", 40*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4
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
			comicGuyPattern(m, turn)
			characterTurnComicGuy(c, m, turn+1, progress)
		case 2:
			AccessInventory(*c)
			characterTurnComicGuy(c, m, turn, progress)
		case 3:
			typeWriter("🏃💨 Vous fuyez en évitant les comics qui volent !", 40*time.Millisecond)
			// Même si on fuit, on obtient l'indice
			typeWriter("👨‍💻 Comic Book Guy : 'Fuyez ! Votre père est au parc d'attractions !'", 40*time.Millisecond)
			progress.HasClue3 = true
			progress.ComicCompleted = true
			progress.Stage = 4
			ScenarioMenu(c, progress)
		default:
			typeWriter("❌ Choix invalide.", 30*time.Millisecond)
			characterTurnComicGuy(c, m, turn, progress)
		}
	}
}

// Attaques spéciales du Comic Book Guy
func comicGuyPattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("📖 Comic Book Guy utilise 'CRITIQUE PARALYSANTE' !", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Pire joueur... de tous les temps !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("😱 Dégâts psychologiques : %d points !", damage), 30*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 12
		typeWriter("💥 Comic Book Guy utilise 'LANCER DE COMICS ÉDITION LIMITÉE' !", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Mes précieux comics ! Ils valent une fortune !'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("📚 Dégâts collectionnables : %d points !", damage), 30*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("🤓 Comic Book Guy utilise 'LEÇON DE CULTURE GEEK' !", 40*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Laissez-moi vous expliquer pourquoi vous avez tort...'", 40*time.Millisecond)
		typeWriter(fmt.Sprintf("📖 Dégâts intellectuels : %d points !", damage), 30*time.Millisecond)
	}
}
