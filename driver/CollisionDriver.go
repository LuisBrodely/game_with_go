package driver

import (
	"SimpsonsGame/models"
)

var t *models.Homero
var w *models.Dona

type CollisionDriver struct {
	gameOver bool
}

func NewCollisionDriver(Homero *models.Homero, Dona *models.Dona) *CollisionDriver {
	t = Homero
	w = Dona
	return &CollisionDriver{
		gameOver: false,
	}
}

func (c *CollisionDriver) Run() {
	for !c.gameOver{
		if w.GetPosY() >= 400 {
			if w.GetPosX() >= t.GetPosX()-50 && w.GetPosX() <= t.GetPosX()+50 {
				w.SetRunning(false)
				t.SetRunning(false)
				c.gameOver = true
			}
		} 
	}
}

func (c *CollisionDriver) GetGameOver() bool {
	return c.gameOver
}