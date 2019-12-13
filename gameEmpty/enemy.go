package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type Enemy struct {
	*tl.Entity
	prevX    int
	prevY    int
	level    *tl.BaseLevel
	cooldown int
}

func (enemy *Enemy) Tick(event tl.Event) {
	if enemy.cooldown > 0 {
		enemy.cooldown -= 1
	} else {
		enemy.cooldown = cd
		rand.Seed(time.Now().UnixNano())
		list := rand.Perm(4)
		var x, y = enemy.Position()
		for _, element := range list {
			switch element {
			case 0:
				if maze[x][y-2] != '*' {
					enemy.SetPosition(x, y-1)
					break
				}
			case 1:
				if maze[x-1][y-1] != '*' {
					enemy.SetPosition(x-1, y)
					break
				}
			case 2:
				if maze[x][y] != '*' {
					enemy.SetPosition(x, y+1)
					break
				}
			case 3:
				if maze[x+1][y-1] != '*' {
					enemy.SetPosition(x+1, y)
					break
				}
			}
		}
	}
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
