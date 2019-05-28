/*
 * DVDTS - DVD like screensaver for term
 * Github repo: https://github.com/ameyrk99/dvdts
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

var (
	termWidth  = 0
	termHeight = 0

	/* Top left co-ordinates of the text */
	px = 1
	py = 1

	/* Text size */
	pTextLength = 0

	xAdd = true
	yAdd = true

	colors    = []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	colorsPos = 0

	osName = "Linux"

	allColors = true
)

func main() {

	/* Get custom text color */
	tempTextColor := flag.String("c", "blue", "color for the bouncing text")

	/* Get text speed */
	textSpeed := flag.Int("s", 10, "speed of text [more is slower]")

	/* Get whether to cycle through colors */
	flag.BoolVar(&allColors, "a", false, "cycle through terminal colors")

	flag.Parse()

	/* Get text color */
	textColor := getTextColor(tempTextColor)

	/* Initialize termui */
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	/* Get OS/distro name */
	osName = getOsName()

	/* Make the text widget */
	p := widgets.NewParagraph()
	p.Border = false
	p.Text = fmt.Sprintf("[%s](fg:%s,mod:bold)", osName, textColor)
	pTextLength = len(osName)
	termWidth, termHeight = ui.TerminalDimensions()
	drawText(&p)

	ui.Render(p)

	uiEvents := ui.PollEvents()
	/* ticker to update the position/color of the text after fixed interval */
	ticker := time.NewTicker(time.Duration(*textSpeed) * 10 * time.Millisecond).C
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

/* Get the name of OS/Distro to display as text */
func getOsName() (osName string) {
	out, err := exec.Command("lsb_release", "-a").Output()
	if err != nil {
		fmt.Println("Error getting distro name:", err)
		os.Exit(1)
	}

	osInfo := fmt.Sprintf("%s", out)
	osStrings := strings.Split(osInfo, "\n")

	return strings.Split(osStrings[2], "\t")[1]
}

/* Get starting/only text color from user on use of c flag */
func getTextColor(textColor *string) string {
	for i, c := range colors {
		if *textColor == c {
			colorsPos = i
			return *textColor
		}
	}

	fmt.Printf("Colors available:\n%s\n", strings.Join(colors, " "))
	os.Exit(1)

	return "blue"
}

/* Update the color of text when it hits the corner if all colors is used */
func updateTextColor(p **widgets.Paragraph) {
	colorsPos++
	(*p).Text = fmt.Sprintf("[%s](fg:%s,mod:bold)", osName, colors[colorsPos])

	/* If the last we're at the last element in the list, begin from start */
	if colorsPos == len(colors)-1 {
		colorsPos = 0
	}
}

/* Draw the text on termui */
func drawText(p **widgets.Paragraph) {
	updateColor := false

	/* Did text hit the bottom or top of the term? */
	if py == termHeight-2 {
		yAdd = false
		updateColor = true
	} else if py == 0 {
		yAdd = true
		updateColor = true
	}

	/* Did the text hit the right or left of term? */
	if px == termWidth-pTextLength-2 {
		xAdd = false
		updateColor = true
	} else if px == 0 {
		xAdd = true
		updateColor = true
	}

	/* Update color on hit and when all a flag is used */
	if updateColor && allColors {
		updateTextColor(p)
	}

	if yAdd {
		py++
	} else {
		py--
	}

	if xAdd {
		px++
	} else {
		px--
	}
	(*p).SetRect(px, py, termWidth, termHeight)
}
