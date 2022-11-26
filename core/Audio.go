package core

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

type Audio struct {
	Path  string
	Times int
}

//func (a *Audio) Play() {
//	go func() {
//		a.Times = 400
//
//		// Read the mp3 file into memory
//		fileBytes, err := os.ReadFile(a.Path)
//		if err != nil {
//			panic(fmt.Sprintf("reading %s.mp3 failed: ", a.Path) + err.Error())
//		}
//
//		// Convert the pure bytes into a reader object that can be used with the mp3 decoder
//		fileBytesReader := bytes.NewReader(fileBytes)
//
//		// Decode file
//		decodedMp3, err := mp3.NewDecoder(fileBytesReader)
//		if err != nil {
//			panic("mp3.NewDecoder failed: " + err.Error())
//		}
//
//		// Prepare an Oto context (this will use your default audio device) that will
//		// play all our sounds. Its configuration can't be changed later.
//
//		// Usually 44100 or 48000. Other values might cause distortions in Oto
//		samplingRate := 44100
//
//		// Number of channels (aka locations) to play sounds from. Either 1 or 2.
//		// 1 is mono sound, and 2 is stereo (most speakers are stereo).
//		numOfChannels := 2
//
//		// Bytes used by a channel to represent one sample. Either 1 or 2 (usually 2).
//		audioBitDepth := 2
//
//		// Remember that you should **not** create more than one context
//		otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
//		if err != nil {
//			panic("oto.NewContext failed: " + err.Error())
//		}
//		// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
//		<-readyChan
//
//		var player oto.Player
//
//		player = otoCtx.NewPlayer(decodedMp3)
//		player.Play()
//
//		timer := time.NewTimer(time.Duration(a.Times) * time.Millisecond)
//
//		<-timer.C
//		err = player.Close()
//		if err != nil {
//			log.Println(err.Error())
//		}
//
//		//time.Sleep(time.Duration(a.Times) * time.Millisecond)
//		//player.Close()
//	}()
//}

var speakerInit bool = false

func (a *Audio) Play() {

	go func() {

		f, err := os.Open(a.Path)
		if err != nil {
			log.Fatal(err)
		}

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		defer streamer.Close()

		if !speakerInit {
			err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			if err != nil {
				log.Fatalln(err)
			}
			speakerInit = true
		}

		speaker.Play(streamer)

		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		<-done
	}()
}
