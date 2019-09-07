package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gizak/termui/widgets"
)

/* Generate random initial coordinates for the text */
func generateRandomCoords() (xCoord, yCoord int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	min := 1
	max := termHeight - 3
	yCoord = rand.Intn(max-min) + min

	max = termWidth - pTextLength - 3
	xCoord = rand.Intn(max-min) + min

	return xCoord, yCoord
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
	if py == yEdge {
		yAdd = false
		updateColor = true
	} else if py == 0 {
		yAdd = true
		updateColor = true
	}

	/* Did the text hit the right or left of term? */
	if px == xEdge {
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
