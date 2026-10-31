package main

import (
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type Game struct{
	goferSprite *ebiten.Image
	bg1 *ebiten.Image
	bg2	*ebiten.Image
	goferX float64
	goferY float64
	direction int
	bgY float64
	gameState int
	currentAngle float64
}

func (g *Game) resetGame() error {
		g.goferX = 400
		g.goferY = 300
		g.direction = 0
		g.bgY = 0
		g.gameState = 0
		g.currentAngle = 0
		return nil
}

func (g *Game) UpdateBackground() error {
		if g.direction != 0 {
		g.bgY += 9
		if g.bgY > 800 {
			g.bgY = 0
		}
	}
	return nil
}

func (g *Game) UpdatePlayer() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.direction = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.direction = -1
	}

	speed := 3.0
	g.goferX += float64(g.direction) * speed

		
	if g.direction != 0 {
		if g.goferX < 38 || g.goferX > 762 {
			g.gameState = 1
		}
	}

	targetAngle := 0.0
	if g.direction == 1{
		targetAngle = -0.6
	}
	if g.direction == -1 {
		targetAngle = 0.6
	}

	rotationSpeed := 0.1
	g.currentAngle += (targetAngle - g.currentAngle) * rotationSpeed

	return nil
}

// runs 60 times a secon, main loop
func (g *Game) Update() error {
	if gameState := 1; g.gameState == gameState {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.resetGame()
		}
		return nil
	}

	g.UpdatePlayer()
	g.UpdateBackground()
	return nil
}

// draws za sproits
func (g *Game) Draw(screen *ebiten.Image) {
	if g.bg1 != nil && g.bg2 != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, -g.bgY)
		screen.DrawImage(g.bg1, options)

		options2 := &ebiten.DrawImageOptions{}
		options2.GeoM.Translate(0, -g.bgY+800)
		screen.DrawImage(g.bg2, options2)
	}


	geom := ebiten.GeoM{}

	width, height := g.goferSprite.Bounds().Dx(), g.goferSprite.Bounds().Dy()

	geom.Translate(-float64(width)/2, -float64(height)/2)

	geom.Rotate(g.currentAngle)
	geom.Translate(g.goferX, g.goferY)


	screen.DrawImage(g.goferSprite, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

// idk wat ts is but coul 
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 800
}

// za ain func
func main() {
	// basic sh
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Gofer: Banishment")
	
	// loadin za gofer sproit
	goferSprite, _, err := ebitenutil.NewImageFromFile("gofer2.png")
	if err != nil {
		log.Println("Failed to load gofer sprite:", err)
	}

	// loadin background images
	bg1, _, err := ebitenutil.NewImageFromFile("bg1.png")
	if err != nil {
		log.Println("Failed to load background image 1:", err)
	}
	bg2, _, err := ebitenutil.NewImageFromFile("bg2.png")
	if err != nil {
		log.Println("Failed to load background image 2:", err)
	}

	// gaem
	game := &Game{
		goferSprite: goferSprite,
		goferX: 400,
		goferY: 300,
		direction: 0,
		bg1: bg1,
		bg2: bg2,
		bgY: 0,
	}

	// run da gaem
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}