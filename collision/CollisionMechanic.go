package driver

import (
	"SimpsonsGame/models"
)

var h *models.Homero
var d *models.Dona

type CollisionMechanic struct {
	gameOver bool
}

func NewCollisionMechanic(Homero *models.Homero, Dona *models.Dona) *CollisionMechanic {
	h = Homero
	d = Dona
	return &CollisionMechanic{
		gameOver: false,
	}
}

func (c *CollisionMechanic) Run() {
	for !c.gameOver{
		if d.GetPosY() >= 400 {
			if d.GetPosX() >= h.GetPosX()-50 && d.GetPosX() <= h.GetPosX()+50 {
				h.SetRunning(false)
				d.SetRunning(false)
				c.gameOver = true
			}
		} 
	}
}

func (c *CollisionMechanic) GetGameOver() bool {
	return c.gameOver
}