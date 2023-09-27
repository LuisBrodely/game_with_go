package models

import (
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
)

type Homero struct {
    posX, posY, direction float32
    running               bool
    pel                   *canvas.Image
}

func NewHomero(posx float32, posy float32, img *canvas.Image) *Homero {
    return &Homero{
        posX:      posx,
        posY:      posy,
        running:   true,
        pel:       img,
        direction: 0,
    }
}

func (t *Homero) GoRight() {
    t.direction = 1
}

func (t *Homero) GoLeft() {
    t.direction = -1
}

func (t *Homero) Run() {
    for t.running {
        var incX float32 = 50
        incX *= t.direction

        if t.posX < 50 {
            t.posX = 50
        } else if t.posX > 650 {
            t.posX = 650
        }

        t.posX += incX
        t.pel.Move(fyne.NewPos(t.posX, t.posY))
        time.Sleep(100 * time.Millisecond)
    }
}

func (t *Homero) SetRunning(pause bool) {
    t.running = pause
}

func (t *Homero) GetRunning() bool {
    return t.running
}

func (t *Homero) GetPosY() float32 {
    return t.posY
}

func (t *Homero) GetPosX() float32 {
    return t.posX
}
