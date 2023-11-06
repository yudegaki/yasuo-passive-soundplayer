package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/yudegaki/yasuo-passive-soundplayer/lolapi"
)

func checkYasuoPassiveStatus(summoner *lolapi.Summoner, prevResourceValue *float32) bool {
	// Check if the champion is Yasuo
	if summoner.ChampionStats.ResourceType != "Wind" {
		return false
	}

	currentResourceValue := summoner.ChampionStats.ResourceValue
	r := false

	if *prevResourceValue == summoner.ChampionStats.ResourceMax && currentResourceValue < *prevResourceValue {
		r = true
	}

	*prevResourceValue = currentResourceValue

	return r
}

func play() error{	
	// Sound player initialization
	f, err := os.Open("play.mp3")
	if err != nil {
		return errors.New("Failed to open mp3 file. Please check if 'play.mp3' exists.")
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return errors.New("Failed to decode mp3 file")
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return errors.New("Failed to initialize oto sound player")
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return errors.New("Failed to play mp3 file")
	}
	return nil
}

func main() {

	var prevResourceValue float32 = -1

	t := time.NewTicker(100 * time.Millisecond)

	defer t.Stop()
	for {
		select {
		case <-t.C:
			summoner, err := lolapi.GetAllGameData()
			if err != nil {
				panic(err)
			}

			if checkYasuoPassiveStatus(summoner, &prevResourceValue) {
				// Play sound
				if err := play(); err != nil {
					panic(err)
				}
			}
		}
	}
}
