/*
 * DVDTS - DVD like screensaver for term
 * Github repo: https://github.com/ameyrk99/dvdts
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

const (
	termOffset = 2
)

var (
	termWidth  = 0
	termHeight = 0
	xEdge      = 0
	yEdge      = 0

	/* Top left co-ordinates of the text */
	px = 1
	py = 1

	/* Text size */
	pTextLength = 0

	xAdd = true
	yAdd = true

	colors    = []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	colorsPos = -1

	osName string

	allColors = true

	textSpeed int
	textColor string
)

func main() {
	log.SetPrefix("")

	/* Get custom text color */
	flag.StringVar(&textColor, "c", "blue", "color for the bouncing text")

	/* Get text speed */
	flag.IntVar(&textSpeed, "s", 10, "speed of text [more is slower]")

	/* Get whether to cycle through colors */
	flag.BoolVar(&allColors, "a", false, "cycle through terminal colors")

	/* Get OS/distro name */
	flag.StringVar(&osName, "t", getOsName(), "custom text to display")

	flag.Parse()

	/* Get text color */
	for i, c := range colors {
		if textColor == c {
			colorsPos = i
			break
		}
	}

	if colorsPos == -1 {
		log.Fatalf("Colors available:\n%s\n", strings.Join(colors, " "))
	}

	/* Initialize termui */
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	/* Make the text widget */
	p := widgets.NewParagraph()
	p.Border = false
	p.Text = fmt.Sprintf("[%s](fg:%s,mod:bold)", osName, textColor)

	pTextLength = len(osName)

	termWidth, termHeight = ui.TerminalDimensions()
	/* Top left doesn't go all the way to the end of terminal due to text wrap, hence new edges */
	xEdge = termWidth - pTextLength - termOffset
	yEdge = termHeight - termOffset

	px, py = generateRandomCoords()

	gcd := float64(getGCD(xEdge, yEdge))

	/* Start loop only if GCD for terminal dimensions isn't 1.0 */
	if gcd != 1.0 {
		for {
			/* The below condition needs to be met for thing to hit the corner */
			/* Math from http://prgreen.github.io/blog/2013/09/30/the-bouncing-dvd-logo-explained/ */
			if math.Mod(math.Abs(float64(px-py)), gcd) == 0 {
				/* If it is met, recalculate starting coordinates */
				px, py = generateRandomCoords()
			} else {
				break
			}
		}
	}

	drawText(&p)

	ui.Render(p)

	uiEvents := ui.PollEvents()
	/* ticker to update the position/color of the text after fixed interval */
	ticker := time.NewTicker(time.Duration(textSpeed) * 10 * time.Millisecond).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				/* Quit the program on q or Ctrl+c */
				return
			case "a":
				/* Switch all colors on/off in the program */
				allColors = !allColors
			}

		case <-ticker:
			drawText(&p)
			ui.Render(p)
		}
	}
}
