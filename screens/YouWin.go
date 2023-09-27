package scenes

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)

type YouWin struct {
    window fyne.Window
}

func NewYouWin(fyneWindow fyne.Window) *YouWin {
    scene := &YouWin{window: fyneWindow}
    scene.Render()
    return scene
}

func (s *YouWin) Render() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/ganaste.png"))
    backgroundImage.Resize(fyne.NewSize(800, 600))
    backgroundImage.Move(fyne.NewPos(0, 0))

    btnExit := widget.NewButton("Salir", s.exitGame) 
    btnExit.Resize(fyne.NewSize(160, 40))
    btnExit.Move(fyne.NewPos(70, 470))

    s.window.SetContent(container.NewWithoutLayout(backgroundImage, btnExit))
}

func (s *YouWin) exitGame() {
    s.window.Close() 
}
