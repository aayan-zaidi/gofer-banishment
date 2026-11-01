package main

import (
	"log"
	"strconv"

	_ "image/jpeg"
	_ "image/png"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RockItem struct {
	x float64
	y float64
}

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
	score int

	rockSprites [3]*ebiten.Image
	rocks [3]RockItem
}

func (g *Game) resetGame() error {
		g.goferX = 400
		g.goferY = 300
		g.direction = 0
		g.bgY = 0
		g.gameState = 0
		g.currentAngle = 0
		g.score = 0 
		for i := 0; i < 3; i++ {
			g.rocks[i].y = float64(rand.Intn(400)) + 850 + float64(i * 250)
			g.rocks[i].x = float64(rand.Intn(700))
		}
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

	speed := 5.0
	g.goferX += float64(g.direction) * speed

		
	if g.direction != 0 {
		g.score++
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

func (g *Game) UpdateRocks() error {
  if g.direction == 0 {
		return nil
	}

	for i := 0; i < 3; i++ {
		g.rocks[i].y -= 6.7

	if g.rocks[i].y < -50 {
		// DO THIS TOMMOROW CHAGNE IT
		g.rocks[i].y = float64(rand.Intn(400)) + 850 + float64(i * 250)
		g.rocks[i].x = float64(rand.Intn(700))
	}

	padding := 10.0

	goferW, goferH := g.goferSprite.Bounds().Dx(), g.goferSprite.Bounds().Dy()
	rockW, rockH := g.rockSprites[i].Bounds().Dx(), g.rockSprites[i].Bounds().Dy()

	playerLeft   := g.goferX - (float64(goferW)/2) + padding
	playerRight  := g.goferX + (float64(goferW)/2) - padding
	playerTop    := g.goferY - (float64(goferH)/2) + padding
	playerBottom := g.goferY + (float64(goferH)/2) - padding

	rockLeft   := g.rocks[i].x + padding
	rockRight  := g.rocks[i].x + float64(rockW) - padding
	rockTop    := g.rocks[i].y + padding
	rockBottom := g.rocks[i].y + float64(rockH) - padding

	if playerLeft < rockRight && playerRight > rockLeft &&
		playerTop < rockBottom && playerBottom > rockTop {
		g.gameState = 1 
	}}
	
	return nil

}

// runs 60 times a secon, main loop
func (g *Game) Update() error {
	if g.gameState == 1 {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.resetGame()
		}
		return nil
	}

	g.UpdatePlayer()
	g.UpdateBackground()
	g.UpdateRocks()	
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

	for i := 0; i < len(g.rocks); i++ {
		rockOptions := &ebiten.DrawImageOptions{}
		rockOptions.GeoM.Translate(g.rocks[i].x, g.rocks[i].y)
		screen.DrawImage(g.rockSprites[i], rockOptions)
	}


	geom := ebiten.GeoM{}

	width, height := g.goferSprite.Bounds().Dx(), g.goferSprite.Bounds().Dy()

	geom.Translate(-float64(width)/2, -float64(height)/2)

	geom.Rotate(g.currentAngle)
	geom.Translate(g.goferX, g.goferY)


	screen.DrawImage(g.goferSprite, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	scoreText := strconv.Itoa(g.score) + " MILES "
		
		
	textWidth := len(scoreText) * 6
	rightX := 800 - textWidth - 15

	ebitenutil.DebugPrintAt(screen, scoreText, rightX, 15)

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
	goferSprite, _, err := ebitenutil.NewImageFromFile("assets/gofer2.png")
	if err != nil {
		log.Println("Failed to load gofer sprite:", err)
	}

	var loadedRocks [3]*ebiten.Image
	for i := 0; i < 3; i++ {
		path := "assets/rock" + strconv.Itoa(i+1) + ".png"
		rockSprite, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Println("Failed to load rock sprite:", err)
		} 
		loadedRocks[i] = rockSprite	
		}


	// loadin background images
	bg1, _, err := ebitenutil.NewImageFromFile("assets/bg2.png")
	if err != nil {
		log.Println("Failed to load background image 1:", err)
	}
	bg2, _, err := ebitenutil.NewImageFromFile("assets/bg2.png")
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
		score: 0,
		rockSprites: loadedRocks,
		rocks: [3]RockItem{
			{ x: float64(rand.Intn(700)), y: float64(rand.Intn(400)) + 850 },
			{ x: float64(rand.Intn(700)), y: float64(rand.Intn(400)) + 1200 },
			{ x: float64(rand.Intn(700)), y: float64(rand.Intn(400)) + 1450 },
		},
	}

	// run da gaem
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}