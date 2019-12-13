package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()
	game.Start()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	level.AddEntity(tl.NewRectangle(5, 10, 30, 1, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(5, 9, 2, 1, tl.ColorBlack))
	game.Screen().SetLevel(level)

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'X'})
	level.AddEntity(&player)

	// gold
	gold := Gold{
		Entity: tl.NewEntity(2, 2, 1, 1),
		level:  level,
	}
	gold.SetCell(0, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '$'})
	level.AddEntity(&gold)

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
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

// func (player *Player) Size() (int, int) {
// 	return player.Size()
// }

// func (player *Player) Position() (int, int) {
// 	return player.Position()
// }

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}

type Gold struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (gold *Gold) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Player:
		gold.level.RemoveEntity(gold)
	}
}
