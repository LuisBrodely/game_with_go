package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
	window fyne.Window
}

func NewGameOverScene(fyneWindow fyne.Window) *GameOverScene {
	scene := &GameOverScene{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *GameOverScene) Render() {

	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/over.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )


	btnRestart := widget.NewButton("Volver a jugar", s.restart)
	btnRestart.Resize(fyne.NewSize(160, 40))
	btnRestart.Move(fyne.NewPos(420, 460))
	

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, btnRestart)) 
}

func (s *GameOverScene) goMenu() {
	NewMenuScene(s.window)
}
func (s *GameOverScene) restart() {
	NewGameScene(s.window)
}