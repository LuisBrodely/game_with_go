package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type Menu struct {
	window fyne.Window
}

func NewMenu(fyneWindow fyne.Window) *Menu {
	scene := &Menu{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *Menu) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/background.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )
	
	botonIniciar := widget.NewButton("Comenzar a jugar", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(160, 40))
	botonIniciar.Move(fyne.NewPos(325,500))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, botonIniciar)) 
}

func (s *Menu) StartGame() {
	NewGameScene( s.window )
}

