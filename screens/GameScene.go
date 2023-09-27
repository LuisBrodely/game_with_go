package scenes

import (
    "SimpsonsGame/driver"
    "SimpsonsGame/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
    "time"
    "strconv" // Importar el paquete strconv para convertir el puntaje a cadena
)

type GameScene struct {
    window fyne.Window
    score  int // Variable para llevar el conteo de puntos
    scoreLabel *widget.Label // Label para mostrar el puntaje
}

var t *models.Homero
var w *models.Dona
var c *driver.CollisionDriver

func NewGameScene(window fyne.Window) *GameScene {
    scene := &GameScene{window: window}
    scene.score = 0 // Inicializa el conteo de puntos
    scene.Render()
    scene.StartGame()
    return scene
}

func (s *GameScene) Render() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/moe.png"))
    backgroundImage.Resize(fyne.NewSize(800, 600))
    backgroundImage.Move(fyne.NewPos(0, 0))

    HomeroPeel := createPeel("./assets/homero.png", 100, 100, 100, 450)
    t = models.NewHomero(350, 450, HomeroPeel)

    DonaPeel := createPeel("./assets/dona.png", 100, 100, 100, 50)
    w = models.NewDona(350, 600, DonaPeel, t)

    c = driver.NewCollisionDriver(t, w)

    btnLeft := widget.NewButton("<", t.GoLeft)
    btnLeft.Resize(fyne.NewSize(80, 80))
    btnLeft.Move(fyne.NewPos(0, 520))

    btnRight := widget.NewButton(">", t.GoRight)
    btnRight.Resize(fyne.NewSize(80, 80))
    btnRight.Move(fyne.NewPos(720, 520))

    // Agregar el marcador de puntuación
    s.scoreLabel = widget.NewLabel("Score: 0")
    s.scoreLabel.Move(fyne.NewPos(10, 10))

    s.window.Canvas().SetOnTypedKey(s.handleArrowKeys)

    s.window.SetContent(container.NewWithoutLayout(
        backgroundImage,
        HomeroPeel,
        DonaPeel,
        btnLeft,
        btnRight,
        s.scoreLabel, // Agregar el marcador al contenido
    ))
}

func (s *GameScene) handleArrowKeys(ev *fyne.KeyEvent) {
    switch ev.Name {
    case "Left":
        t.GoLeft()
    case "Right":
        t.GoRight()
    }
}

func (s *GameScene) StartGame() {
    go s.countScore()
    go t.Run()
    go w.Run()
    go c.Run()
    go s.checkGameOver()
}

func (s *GameScene) incrementScore() {
    s.score++
    s.updateScoreLabel() // Actualizar el texto del marcador
}

func (s *GameScene) getScore() int {
    return s.score
}

func (s *GameScene) countScore() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            s.incrementScore()
        }
    }
}

func (s *GameScene) checkGameOver() {
    running := true
    for running {
        if c.GetGameOver() || s.getScore() >= 50 { // Cambia 10 a la cantidad deseada de puntos para mostrar Game Over
            running = false
            time.Sleep(2000 * time.Millisecond)
            NewGameOverScene(s.window)
        }
    }
}

func (s *GameScene) updateScoreLabel() {
    scoreStr := strconv.Itoa(s.getScore())
    s.scoreLabel.SetText("Score: " + scoreStr)
}

func createPeel(fileURI string, sizeX float32, sizeY float32, posX float32, posY float32) *canvas.Image {
    image := canvas.NewImageFromURI(storage.NewFileURI(fileURI))
    image.Resize(fyne.NewSize(sizeX, sizeY))
    image.Move(fyne.NewPos(posX, posY))
    return image
}
