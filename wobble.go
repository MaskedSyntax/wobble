package main

import (
    "os"
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {
    text := "hello"
    if len(os.Args) > 1 {
        text = strings.Join(os.Args[1:], " ")
    }
	fig := figure.NewFigure(text, "", true)
	rainbowPrint(fig.String())
}

func rainbowPrint(text string) {
	colors := [][]int{
		{255, 153, 204},
		{204, 255, 255},
		{204, 255, 204},
		{255, 204, 255},
		{255, 255, 204},
	}
	colorIndex := 0

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		for _, char := range line {
			r, g, b := colors[colorIndex%len(colors)][0], colors[colorIndex%len(colors)][1], colors[colorIndex%len(colors)][2]
			fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, char)
			colorIndex++
		}
		fmt.Println()
	}
}
