package main

import (
	_ "embed"
	"flag"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	SpeedOfSoundInFeetPerSecond float64 = 1125
	FeetPerMile                 float64 = 5280
)

//go:embed app.png
var icon []byte

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
	w := a.NewWindow("Lightning")
	w.SetIcon(fyne.NewStaticResource("icon", icon))
	seconds := widget.NewLabel(fmt.Sprintf("%f seconds", float64(1)))
	miles := widget.NewLabel(fmt.Sprintf("%f miles", distanceCalc(1)))
	cb = widget.NewButton("Flash", func() {
		if waiting {
			tt := time.Now().Sub(start).Seconds()
			seconds.SetText(fmt.Sprintf("%f seconds", tt))
			miles.SetText(fmt.Sprintf("%f miles", distanceCalc(tt)))
			cb.SetText("Flash")
			waiting = false
			return
		}
		start = time.Now()
		waiting = true
		cb.SetText("Thunder!")
	})
	vbox := container.NewVBox(
		seconds,
		miles,
		cb,
	)
	w.SetContent(vbox)
	w.ShowAndRun()
}

func main() {
	d := flag.Bool("d", false, "disable gui")
	flag.Parse()
	if *d {
		for {
			fmt.Print("press enter on flash")
			fmt.Scanln()
			t0 := time.Now()
			fmt.Print("press enter on thunder")
			fmt.Scanln()
			tt := time.Now().Sub(t0).Seconds()
			fmt.Printf("strike occured %f miles away, sound took %f seconds\n", distanceCalc(tt), tt)
		}
	}
	gui()
}
