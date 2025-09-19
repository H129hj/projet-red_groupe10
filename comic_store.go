package projetred

import (
	"fmt"
	"time"
)

func ComicBookStore(c *Character, progress *ScenarioProgress) {
	typeWriter("==================================================", 15*time.Millisecond)
	typeWriter("             📚 MAGASIN DE BD", 15*time.Millisecond)
	typeWriter("==================================================", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("🏪 Vous entrez dans l'antre du Comic Book Guy...", 15*time.Millisecond)
	typeWriter("📚 Des milliers de BD, figurines et objets geek s'entassent partout.", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	typeWriter("👨‍💻 Comic Book Guy : 'Ah ! Un client potentiel dans mon sanctuaire de la culture pop !'", 15*time.Millisecond)
	typeWriter("👨‍💻 Comic Book Guy : 'Que puis-je faire pour... attends, tu es un Simpson !'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	switch c.class {
	case "bart":
		typeWriter("🎯 Bart : 'Salut mec ! Je cherche mon père, Homer Simpson.'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ah, le célèbre Homer ! Collectionneur de... donuts.'", 15*time.Millisecond)
	case "lisa":
		typeWriter("🎷 Lisa : 'Bonjour, je mène une enquête méthodique sur mon père disparu.'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Enfin quelqu'un avec un cerveau ! Contrairement à ton père...'", 15*time.Millisecond)
	case "maggie":
		typeWriter("👶 Maggie : '*regarde fixement une figurine de superhéros*'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ce bébé a du goût ! Elle regarde ma figurine rare de...'", 15*time.Millisecond)
	}

	typeWriter("", 15*time.Millisecond)

	var choice int
	typeWriter("💭 Comment aborder le Comic Book Guy ?", 15*time.Millisecond)
	typeWriter("1. 😊 'Pouvez-vous m'aider poliment, s'il vous plaît ?'", 15*time.Millisecond)
	typeWriter("2. 😤 'Dépêche-toi, j'ai pas toute la journée !'", 15*time.Millisecond)
	typeWriter("3. 🤓 'J'ai besoin de vos connaissances encyclopédiques.'", 15*time.Millisecond)
	ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
	fmt.Scan(&choice)

	switch choice {
	case 1:
		typeWriter("👨‍💻 Comic Book Guy : 'Ah ! Enfin de la politesse ! Très rafraîchissant.'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Ton père est venu ce matin avec un ticket étrange...'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Il parlait d'un \"Grand Concours du Donut Cosmique\" au parc d'attractions.'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Apparemment, le gagnant obtient des donuts gratuits à vie !'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Homer participe au Grand Concours du Donut Cosmique !", 15*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4

	case 2:
		typeWriter("👨‍💻 Comic Book Guy : 'QUOI ?! L'audace ! L'impudence !'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Personne ne parle ainsi au maître de la culture geek !'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)

		var fightChoice int
		typeWriter("💭 Le Comic Book Guy semble énervé...", 15*time.Millisecond)
		typeWriter("1. ⚔️ 'Alors on se bat ?'", 15*time.Millisecond)
		typeWriter("2. 😅 'Désolé, je me suis mal exprimé...'", 15*time.Millisecond)
		ColoredTypeWriter("➤ Votre choix : ", 15*time.Millisecond, BrightCyan+Bold)
		fmt.Scan(&fightChoice)

		if fightChoice == 1 {
			typeWriter("⚔️ Comic Book Guy entre en mode CRITIQUE ULTIME !", 15*time.Millisecond)
			comicGuyMonster := InitMonster("Comic Book Guy (Critique Ultime)", 350, 28)
			ScenarioCombat(c, &comicGuyMonster, progress, comicGuyPattern, "comic")
			return
		} else {
			typeWriter("👨‍💻 Comic Book Guy : 'Hmph ! Bon... ton père parlait du parc d'attractions.'", 15*time.Millisecond)
			typeWriter("👨‍💻 Comic Book Guy : 'Quelque chose à propos d'un concours de donuts...'", 15*time.Millisecond)
			typeWriter("", 15*time.Millisecond)
			typeWriter("🔍 INDICE OBTENU : Homer est au parc d'attractions pour un concours !", 15*time.Millisecond)
			progress.HasClue3 = true
			progress.ComicCompleted = true
			progress.Stage = 4
		}

	case 3:
		typeWriter("👨‍💻 Comic Book Guy : '*rougit* Mes connaissances encyclopédiques ! Enfin quelqu'un qui apprécie !'", 15*time.Millisecond)

		typeWriter("👨‍💻 Comic Book Guy : 'Ton père, Homer Simpson, est venu avec un ticket mystérieux...'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Il s'agit du légendaire \"Grand Concours du Donut Cosmique\"...'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Organisé par des forces... extraterrestres... au parc d'attractions !'", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Le gagnant obtient des donuts magiques pour l'éternité !'", 15*time.Millisecond)
		typeWriter("", 15*time.Millisecond)
		typeWriter("🔍 INDICE OBTENU : Concours extraterrestre de donuts au parc d'attractions !", 15*time.Millisecond)
		progress.HasClue3 = true
		progress.ComicCompleted = true
		progress.Stage = 4

	default:
		typeWriter("❌ Choix invalide.", 15*time.Millisecond)
		ComicBookStore(c, progress)
		return
	}

	typeWriter("", 15*time.Millisecond)
	typeWriter("👨‍💻 Comic Book Guy : 'Maintenant va ! Et que la force soit avec toi... ou pas.'", 15*time.Millisecond)
	typeWriter("", 15*time.Millisecond)

	ScenarioMenu(c, progress)
}

func comicGuyPattern(m *Monster, turn int) {
	if turn%3 == 0 {
		damage := m.power * 2
		typeWriter("📖 Comic Book Guy utilise 'CRITIQUE PARALYSANTE' !", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Pire joueur... de tous les temps !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("😱 Dégâts psychologiques : %d points !", damage), 15*time.Millisecond)
	} else if turn%2 == 0 {
		damage := m.power + 12
		typeWriter("💥 Comic Book Guy utilise 'LANCER DE COMICS ÉDITION LIMITÉE' !", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Mes précieux comics ! Ils valent une fortune !'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("📚 Dégâts collectionnables : %d points !", damage), 15*time.Millisecond)
	} else {
		damage := m.power
		typeWriter("🤓 Comic Book Guy utilise 'LEÇON DE CULTURE GEEK' !", 15*time.Millisecond)
		typeWriter("👨‍💻 Comic Book Guy : 'Laissez-moi vous expliquer pourquoi vous avez tort...'", 15*time.Millisecond)
		typeWriter(fmt.Sprintf("📖 Dégâts intellectuels : %d points !", damage), 15*time.Millisecond)
	}
}
