package projetred

import (
	"time"
)

// Conclusion de l'histoire
func StoryConclusion(c *Character, progress *ScenarioProgress) {
	typeWriter("🎬 CONCLUSION - LE RETOUR À LA MAISON", 50*time.Millisecond)
	typeWriter("==================================================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🏠 De retour à la maison Simpson...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Arrivée de Marge
	typeWriter("👩‍🦱 Soudain, Marge apparaît, les mains sur les hanches...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👩‍🦱 Marge : 'HOMER J. SIMPSON !'", 50*time.Millisecond)
	typeWriter("👨‍🦲 Homer : '*sursaute* Oh oh... Bonjour chérie...'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👩‍🦱 Marge : 'Tu devais juste prendre UNE bière chez Moe !'", 40*time.Millisecond)
	typeWriter("👩‍🦱 Marge : 'Et je te retrouve dans un concours de donuts avec des ALIENS ?!'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👨‍🦲 Homer : 'Mais chérie ! C'était de la science ! Pour l'humanité !'", 40*time.Millisecond)
	typeWriter("👨‍🦲 Homer : 'Et j'ai gagné des donuts gratuits à vie !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👩‍🦱 Marge : '*soupir* Homer... tu ne changeras jamais...'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Remerciements selon le personnage
	switch c.class {
	case "bart":
		typeWriter("👩‍🦱 Marge : 'Merci Bart d'avoir retrouvé ton père.'", 40*time.Millisecond)
		typeWriter("👩‍🦱 Marge : 'Même si j'imagine que tu as fait quelques bêtises en chemin...'", 40*time.Millisecond)
		typeWriter("🎯 Bart : 'Hehe... peut-être quelques-unes...'", 40*time.Millisecond)
	case "lisa":
		typeWriter("👩‍🦱 Marge : 'Lisa, ma chérie, tu es si responsable !'", 40*time.Millisecond)
		typeWriter("👩‍🦱 Marge : 'Comment as-tu fait pour le retrouver si vite ?'", 40*time.Millisecond)
		typeWriter("🎷 Lisa : 'Méthode scientifique et déduction logique, Maman.'", 40*time.Millisecond)
	case "maggie":
		typeWriter("👩‍🦱 Marge : 'Maggie ! Comment as-tu... ?'", 40*time.Millisecond)
		typeWriter("👩‍🦱 Marge : 'Parfois je me demande si tu n'es pas plus intelligente que nous tous...'", 40*time.Millisecond)
		typeWriter("👶 Maggie : '*suce sa tétine avec un air satisfait*'", 40*time.Millisecond)
	}

	typeWriter("", 30*time.Millisecond)

	// Arrivée du Maire Quimby
	typeWriter("🎩 Soudain, une limousine s'arrête devant la maison...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎩 Le Maire Quimby en sort, suivi de journalistes !", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🎩 Maire Quimby : 'Mes chers concitoyens de Springfield !'", 40*time.Millisecond)
	typeWriter("🎩 Maire Quimby : 'Aujourd'hui, nous célébrons un héros !'", 40*time.Millisecond)

	switch c.class {
	case "bart":
		typeWriter("🎩 Maire Quimby : 'Bart Simpson a sauvé son père des griffes extraterrestres !'", 40*time.Millisecond)
		typeWriter("🎩 Maire Quimby : 'Malgré son jeune âge et ses... tendances délinquantes.'", 40*time.Millisecond)
	case "lisa":
		typeWriter("🎩 Maire Quimby : 'Lisa Simpson a résolu cette affaire avec un brillant esprit scientifique !'", 40*time.Millisecond)
		typeWriter("🎩 Maire Quimby : 'Elle est l'avenir de notre belle ville !'", 40*time.Millisecond)
	case "maggie":
		typeWriter("🎩 Maire Quimby : 'Maggie Simpson a... euh... fait quelque chose d'extraordinaire !'", 40*time.Millisecond)
		typeWriter("🎩 Maire Quimby : 'Même si nous ne comprenons pas vraiment comment...'", 40*time.Millisecond)
	}

	typeWriter("", 30*time.Millisecond)
	typeWriter("🎩 Maire Quimby : 'Springfield est sauvé ! Et... Homer a encore trouvé le moyen de faire du grabuge.'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Bonus post-crédit
	typeWriter("🌌 BONUS POST-CRÉDIT", 50*time.Millisecond)
	typeWriter("==============================", 30*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🛸 Quelque part dans l'espace, à bord d'un vaisseau alien...", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👽 Kang : 'Notre expérience était un succès, Kodos !'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'En effet ! Ces donuts contiennent une puissance insoupçonnée...'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👽 Kang : 'L'humain Homer a ingéré assez de donuts pour alimenter notre vaisseau !'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'Qui aurait cru que la gourmandise humaine serait si... utile ?'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("👽 Kang : 'Préparons la prochaine phase : l'invasion par les donuts !'", 40*time.Millisecond)
	typeWriter("👽 Kodos : 'Mwahahaha ! Les humains ne résisteront pas à nos donuts cosmiques !'", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Fin épique
	typeWriter("🎬 THE END", 80*time.Millisecond)
	typeWriter("", 30*time.Millisecond)
	typeWriter("🎵 *Musique des Simpson*", 50*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	typeWriter("🏆 FÉLICITATIONS !", 50*time.Millisecond)
	switch c.class {
	case "bart":
		typeWriter("🎯 Vous avez terminé l'aventure avec Bart Simpson !", 40*time.Millisecond)
		typeWriter("🎯 'Cowabunga mec !' - Bart", 40*time.Millisecond)
	case "lisa":
		typeWriter("🎷 Vous avez terminé l'aventure avec Lisa Simpson !", 40*time.Millisecond)
		typeWriter("🎷 'La connaissance triomphe toujours !' - Lisa", 40*time.Millisecond)
	case "maggie":
		typeWriter("👶 Vous avez terminé l'aventure avec Maggie Simpson !", 40*time.Millisecond)
		typeWriter("👶 '*suce sa tétine mystérieusement*' - Maggie", 40*time.Millisecond)
	}

	typeWriter("", 30*time.Millisecond)
	typeWriter("🎮 Merci d'avoir joué à Springfield RPG !", 40*time.Millisecond)
	typeWriter("🍩 'Mmm... RPG textuel...' - Homer Simpson", 40*time.Millisecond)
	typeWriter("", 30*time.Millisecond)

	// Retour au menu principal
	typeWriter("🔄 Retour au menu principal...", 40*time.Millisecond)
	Menu(*c)
}
