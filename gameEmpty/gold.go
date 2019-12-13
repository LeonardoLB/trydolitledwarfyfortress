package main

import tl "github.com/JoelOtter/termloop"

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
	}
}

func RespawnGold(gold *Gold) {

}
