package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOverScreen struct {
	window fyne.Window
}

func NewGameOverScreen(fyneWindow fyne.Window) *GameOverScreen {
	scene := &GameOverScreen{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *GameOverScreen) Render() {

	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/over.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )


	btnRestart := widget.NewButton("Volver a jugar", s.restart)
	btnRestart.Resize(fyne.NewSize(160, 40))
	btnRestart.Move(fyne.NewPos(420, 460))
	

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, btnRestart)) 
}

func (s *GameOverScreen) goMenu() {
	NewMenu(s.window)
}
func (s *GameOverScreen) restart() {
	NewGameScene(s.window)
}