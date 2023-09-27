package scenes

import (
    "SimpsonsGame/driver"
    "SimpsonsGame/models"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)

type GameScene struct {
    window fyne.Window
}

var t *models.Homero
var w *models.Dona
var c *driver.CollisionDriver

func NewGameScene(window fyne.Window) *GameScene {
    scene := &GameScene{window: window}
    scene.Render()
    scene.StartGame()
    return scene
}

func (s *GameScene) Render() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/moe.png"))
    backgroundImage.Resize(fyne.NewSize(800, 600))
    backgroundImage.Move(fyne.NewPos(0, 0))
    // Models Render
    HomeroPeel := createPeel("./assets/homero.png", 100, 100, 100, 450)
    t = models.NewHomero(350, 450, HomeroPeel)

    DonaPeel := createPeel("./assets/dona.png", 100, 100, 100, 50)
    w = models.NewDona(350, 600, DonaPeel, t)

    //CollisionDriver
    c = driver.NewCollisionDriver(t, w)

    // Buttons Render
    btnLeft := widget.NewButton("<", t.GoLeft)
    btnLeft.Resize(fyne.NewSize(50, 50))
    btnLeft.Move(fyne.NewPos(350, 550))

    btnRight := widget.NewButton(">", t.GoRight)
    btnRight.Resize(fyne.NewSize(50, 50))
    btnRight.Move(fyne.NewPos(400, 550))

    // Registrar el manejador de eventos de teclado
    s.window.Canvas().SetOnTypedKey(s.handleArrowKeys)

    s.window.SetContent(container.NewWithoutLayout(backgroundImage, HomeroPeel, DonaPeel, btnLeft, btnRight))
}

// handleArrowKeys maneja eventos de teclado para mover al personaje.
func (s *GameScene) handleArrowKeys(ev *fyne.KeyEvent) {
    switch ev.Name {
    case "Left":
        t.GoLeft() // Llama a la función GoLeft() en tu objeto de jugador (t)
    case "Right":
        t.GoRight() // Llama a la función GoRight() en tu objeto de jugador (t)
    }
}

func (s *GameScene) StartGame() {
    go t.Run()
    go w.Run()
    go c.Run()
    go s.checkGameOver() 
}

func (s *GameScene) StopGame() {
    t.SetRunning(!t.GetRunning())
    w.SetRunning(!w.GetRunning())
}

func (s *GameScene) checkGameOver() {
    running := true
    for running {
        if c.GetGameOver() {
            running = false
            time.Sleep(2000 * time.Millisecond)
            NewGameOverScene(s.window)

        }
    }
}

func createPeel(fileURI string, sizeX float32, sizeY float32, posX float32, posY float32) *canvas.Image {
    image := canvas.NewImageFromURI(storage.NewFileURI(fileURI))
    image.Resize(fyne.NewSize(sizeX, sizeY))
    image.Move(fyne.NewPos(posX, posY))
    return image
}