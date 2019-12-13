package main

import (
	"fmt"
	"os"

	tl "github.com/JoelOtter/termloop"
)

var maze [][]rune
var width int = 10
var height int = 10
var score = 0
var scoreText *tl.Text

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(30)
	scoreText = tl.NewText(0, 0, " Score: 0 ", tl.ColorBlack, tl.ColorWhite)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	maze = generateMaze(width, height)
	level.AddEntity(scoreText)
	game.Screen().SetLevel(level)

	for i, row := range maze {
		for j, path := range row {
			j += 1
			if path == '*' {
				level.AddEntity(tl.NewRectangle(i, j, 1, 1, tl.ColorBlack))
			} else if path == 'S' {
				player := Player{
					Entity: tl.NewEntity(i, j, 1, 1),
					level:  level,
				}
				player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'X'})
				level.AddEntity(&player)
			} else if path == 'L' {
				gold := Gold{
					Entity: tl.NewEntity(i, j, 1, 1),
					level:  level,
				}
				gold.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '$'})
				level.AddEntity(&gold)
			} else if path == 'E' {
				enemy := Enemy{
					Entity: tl.NewEntity(i, j, 1, 1),
					level:  level,
				}
				enemy.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: 'E'})
				level.AddEntity(&enemy)
			}
		}
	}

	game.Start()
}

func logging(message string) {
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := message
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// IncreaseScore increases the score by the given amount and updates the
// score text.
func IncreaseScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score))
}
