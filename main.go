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
)

func main() {

	/* Get text color */
	textColor := getTextColor()

	/* Initialize termui */
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	/* Get OS/distro name */
	osName := getOsName()

	/* Make the text widget */
	p := widgets.NewParagraph()
	p.Border = false
	p.Text = fmt.Sprintf("[%s](fg:%s,mod:bold)", osName, textColor)
	pTextLength = len(osName)
	termWidth, termHeight = ui.TerminalDimensions()
	drawFunction(&p)

	ui.Render(p)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(100 * time.Millisecond).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			drawFunction(&p)
			ui.Render(p)
		}
	}
}

func getTextColor() string {
	textColor := flag.String("c", "blue", "color for the bouncing text")
	flag.Parse()

	colors := []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}

	for _, c := range colors {
		if *textColor == c {
			return *textColor
		}
	}

	fmt.Printf("Colors available:\n%s\n", strings.Join(colors, " "))
	os.Exit(1)

	return "blue"
}

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

func drawFunction(p **widgets.Paragraph) {
	if py == termHeight-1 {
		yAdd = false
	} else if py == 0 {
		yAdd = true
	}

	if px == termWidth-pTextLength-2 {
		xAdd = false
	} else if px == 0 {
		xAdd = true
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
