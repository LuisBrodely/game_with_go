package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var t *Homero

type Dona struct {
	posX, posY, direction float32
	running bool	
	pel *canvas.Image
}

func NewDona(posx float32, posy float32, img *canvas.Image, Homero *Homero) *Dona {
	t = Homero
	return &Dona{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
	}
}

func (w *Dona) Run() {
	for w.running {
		var inc float32 = 50

		if w.posY > 450 {
			w.posY = -50
			w.posX = float32((rand.Intn(12) + 1) * 50)
		}
		
		w.posY += inc
		w.pel.Move(fyne.NewPos(w.posX,w.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (w *Dona) SetRunning(pause bool) {
	w.running = pause
}
func (w *Dona) GetRunning() bool {
	return w.running
}
func (w *Dona) GetPosY() float32 {
	return w.posY
}
func (w *Dona) GetPosX() float32 {
	return w.posX
}

