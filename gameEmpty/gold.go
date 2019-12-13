package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type Gold struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (gold *Gold) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Player:
		// gold.level.RemoveEntity(gold)
		RespawnGold(gold)
		IncreaseScore(1)
		logging("Gold Collected")
	case *Enemy:
		RespawnGold(gold)
		IncreaseScore(-1)
		logging("Gold Stolen")
	}
}

func RespawnGold(gold *Gold) {
	rand.Seed(time.Now().UnixNano())
	var x int = rand.Intn(width)
	var y int = rand.Intn(height)
	for maze[x][y] == '*' {
		x = rand.Intn(width)
		y = rand.Intn(height)
	}
	gold.SetPosition(x, y+1)
}

func (gold *Gold) Draw(screen *tl.Screen) {
	gold.Entity.Draw(screen)
}
