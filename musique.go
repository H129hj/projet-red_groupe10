package projetred

import (
	"log"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// MusiqueJouer lit et joue le fichier son.mp3 en boucle infinie
func MusiqueJouer() {
	f, err := os.Open("son.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(1024))

	// Boucle infinie
	loop := beep.Loop(-1, streamer)
	speaker.Play(loop)

	select {} // bloque pour garder la musique active
}
