package main

import tl "github.com/JoelOtter/termloop"

type Enemy struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (enemy *Enemy) Tick(event tl.Event) {
	// if event.Type == tl.EventKey { // Is it a keyboard event?
	// 	player.prevX, player.prevY = player.Position()
	// 	switch event.Key { // If so, switch on the pressed key.
	// 	case tl.KeyArrowRight:
	// 		player.SetPosition(player.prevX+1, player.prevY)
	// 		logging("player walking to Right")
	// 	case tl.KeyArrowLeft:
	// 		player.SetPosition(player.prevX-1, player.prevY)
	// 		logging("player walking to Left")
	// 	case tl.KeyArrowUp:
	// 		player.SetPosition(player.prevX, player.prevY-1)
	// 		logging("player walking to Up")
	// 	case tl.KeyArrowDown:
	// 		player.SetPosition(player.prevX, player.prevY+1)
	// 		logging("player walking to Down")
	// 	}
	// }
}

func (enemy *Enemy) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		enemy.SetPosition(enemy.prevX, enemy.prevY)
	}
}

func (enemy *Enemy) Draw(screen *tl.Screen) {
	// screenWidth, screenHeight := screen.Size()
	// x, y := player.Position()
	// player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// // We need to make sure and call Draw on the underlying Entity.
	enemy.Entity.Draw(screen)
}
