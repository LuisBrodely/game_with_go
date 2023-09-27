package driver

import (
	"SimpsonsGame/models"
)

var t *models.Homero
var w *models.Dona

type CollisionMechanic struct {
	gameOver bool
}

func NewCollisionMechanic(Homero *models.Homero, Dona *models.Dona) *CollisionMechanic {
	t = Homero
	w = Dona
	return &CollisionMechanic{
		gameOver: false,
	}
}

func (c *CollisionMechanic) Run() {
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

func (c *CollisionMechanic) GetGameOver() bool {
	return c.gameOver
}