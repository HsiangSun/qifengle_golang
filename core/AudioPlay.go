package core

import (
	"strings"
	time "time"
)

const (
	MAIN = iota
	ACCOMPANIMENTS
)

type Mode int

type AudioPlay struct {
	Notes []string
	Times int
	Mode  Mode
}

func (a *AudioPlay) LoadNotes(notes string) {
	a.Notes = strings.Split(notes, " ")
}

func (a *AudioPlay) Play() {

	audio := Audio{Path: "resource/test.mp3"}
	audio.Play()
	time.Sleep(1000 * time.Millisecond)

	for _, note := range a.Notes {
		if len(note) == 0 {
			continue
		}
		switch note {
		case "1--":
			audio := Audio{Path: "resource/ll1.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "2--":
			audio := Audio{Path: "resource/ll2.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "3--":
			audio := Audio{Path: "resource/ll3.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "4--":
			audio := Audio{Path: "resource/ll4.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "5--":
			audio := Audio{Path: "resource/ll5.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "6--":
			audio := Audio{Path: "resource/ll6.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "7--":
			audio := Audio{Path: "resource/ll7.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "1-":
			audio := Audio{Path: "resource/l1.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "2-":
			audio := Audio{Path: "resource/l2.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "3-":
			audio := Audio{Path: "resource/l3.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "4-":
			audio := Audio{Path: "resource/l4.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "5-":
			audio := Audio{Path: "resource/l5.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "6-":
			audio := Audio{Path: "resource/l6.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "7-":
			audio := Audio{Path: "resource/l7.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "1":
			audio := Audio{Path: "resource/m1.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "2":
			audio := Audio{Path: "resource/m2.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "3":
			audio := Audio{Path: "resource/m3.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "4":
			audio := Audio{Path: "resource/m4.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "5":
			audio := Audio{Path: "resource/m5.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "6":
			audio := Audio{Path: "resource/m6.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "7":
			audio := Audio{Path: "resource/m7.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "1+":
			audio := Audio{Path: "resource/h1.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "2+":
			audio := Audio{Path: "resource/h2.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "3+":
			audio := Audio{Path: "resource/h3.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "4+":
			audio := Audio{Path: "resource/h4.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "5+":
			audio := Audio{Path: "resource/h5.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "6+":
			audio := Audio{Path: "resource/h6.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "7+":
			audio := Audio{Path: "resource/h7.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "1++":
			audio := Audio{Path: "resource/hh1.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "2++":
			audio := Audio{Path: "resource/hh2.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "3++":
			audio := Audio{Path: "resource/hh3.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "4++":
			audio := Audio{Path: "resource/hh4.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "5++":
			audio := Audio{Path: "resource/hh5.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "6++":
			audio := Audio{Path: "resource/hh6.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "7++":
			audio := Audio{Path: "resource/hh7.mp3"}
			audio.Play()
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		case "0":
			time.Sleep(time.Duration(a.Times/2) * time.Millisecond)
			break
		default:
			continue
		}

	}
}
