package main

import (
	"flag"
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	SpeedOfSoundInFeetPerSecond float64 = 1125
	FeetPerMile                 float64 = 5280
)

var (
	cb      *widget.Button
	start   time.Time
	waiting bool
)

// distanceCalc : ( TimeTook * SpeedOfSoundInFeetPerSecond ) / FeetPerMile
func distanceCalc(t float64) float64 {
	return (t * SpeedOfSoundInFeetPerSecond) / FeetPerMile
}

func gui() {
	a := app.New()
	w := a.NewWindow("lightning")
	info := widget.NewLabel(fmt.Sprintf("%f seconds, %f miles", float64(1), distanceCalc(1)))
	calFunc := func() {
		if waiting {
			tt := time.Now().Sub(start).Seconds()
			info.SetText(fmt.Sprintf("%f seconds, %f miles", tt, distanceCalc(tt)))
			cb.SetText("flash")
			waiting = false
			return
		}
		start = time.Now()
		waiting = true
		cb.SetText("Thunder!")
	}
	cb = widget.NewButton("Flash", calFunc)
	w.SetContent(container.NewVBox(
		info,
		cb,
	))
	w.ShowAndRun()
}

func main() {
	d := flag.Bool("d", false, "disable gui")
	flag.Parse()
	fmt.Println(*d)
	if *d {
		for {
			fmt.Println("press enter on flash")
			fmt.Scanln()
			t0 := time.Now()
			fmt.Println("press enter on thunder")
			fmt.Scanln()
			tt := time.Now().Sub(t0).Seconds()
			fmt.Printf("strike occured %f miles away, sound took %f seconds\n", distanceCalc(1), tt)
		}
	}
	gui()
}
