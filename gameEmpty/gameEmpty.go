package main

import (
	"fmt"
	"os"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(30)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	maze := generateMaze(30, 30)
	// level.AddEntity(tl.NewRectangle(5, 10, 30, 1, tl.ColorBlue))
	// level.AddEntity(tl.NewRectangle(5, 9, 2, 1, tl.ColorBlack))
	game.Screen().SetLevel(level)

	for i, row := range maze {
		for j, path := range row {
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
			}
		}
	}

	game.Start()
}

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
			logging("player walking to Right")
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
			logging("player walking to Left")
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
			logging("player walking to Up")
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
			logging("player walking to Down")
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	// screenWidth, screenHeight := screen.Size()
	// x, y := player.Position()
	// player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// // We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
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
