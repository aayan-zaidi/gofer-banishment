// I COULDNT WAIT FOR YOU TO CALL AND CLEAR THE CUPBOARDS
// BUT NOW YOURE GOING TO LEAVE WIHT NATINK BUT A SIGN
// ANOTHER EVENIN ILL BE SITTIN READIN INBETWEEN YOUR LINES
// BECAUSE I MISS YOU ALL THE TIME
// SO GET AWAY ANOTHER WAY TO FEEL WHAT YOU DIDNT WANT YOURELF TO KNOW AND LET YOURSELF GO
// YOU KNOW YOU DIDNT LOSE YOUR SELF CONTROL
// LETS START AT THe RAINBOW
// sorry. wish we could go to the moon together
// heres some trauma for you muheheehhehe


// im sorry for making this btw, first game tho so still pretty proud

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
	introDone bool
	gameOber bool
	
	goferSprite *ebiten.Image
	gameOberOverlay *ebiten.Image
	goferX float64
	goferY float64
	direction int
	
	bg1 *ebiten.Image
	bg2	*ebiten.Image
	bgY float64
	
	loreThingy *ebiten.Image
	
	gameState int
	score int
	
	currentSpeed float64
	currentAngle float64

	rockSprites [3]*ebiten.Image
	rocks [3]RockItem
}

func (g *Game) resetGame() error {
		g.goferX = 400
		g.goferY = 300
		g.direction = 0
		g.bgY = 0
		g.currentSpeed = 10.0
		g.gameState = 0
		g.currentAngle = 0
		g.score = 0 
		g.gameOber = false
		for i := 0; i < 3; i++ {
			g.rocks[i].y = float64(rand.Intn(400)) + 850 + float64(i * 250)
			g.rocks[i].x = float64(rand.Intn(700))
		}
		
		return nil
}

func (g *Game) UpdateBackground() error {
		if g.direction != 0 {
		g.bgY += g.currentSpeed
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
		g.rocks[i].y -= g.currentSpeed

	if g.rocks[i].y < -50 {
		// DO THIS TOMMOROW CHAGNE IT
		g.rocks[i].y = float64(rand.Intn(400)) + 850 + float64(i * 250)
		g.rocks[i].x = float64(rand.Intn(700))
	}

	padding := 10.0

	// collision
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
		g.gameOber = true
	}}
	
	return nil
}

// runs 60 times a secon, main loop
func (g *Game) Update() error {
	// again quick checks and sh
	if !g.introDone {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.introDone = true
		}
		return nil 
	}
	if g.gameOber {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.gameOber = false
		}
		return nil 
	}

	if g.gameState == 1 {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.resetGame()
		}
		return nil
	}

	// i feel so proud not having everything in one function bro
	g.UpdatePlayer()
	g.UpdateBackground()
	g.UpdateRocks()	
	return nil
}

// draws za sproits
func (g *Game) Draw(screen *ebiten.Image) {
	// quick returns and checks
	if g.introDone == false {
		screen.DrawImage(g.loreThingy, nil)
		return
	}

	if g.gameOber == true {
		screen.DrawImage(g.gameOberOverlay, nil)
		return
	}


	g.currentSpeed = 10.0 + float64(g.score)/500.0

	if g.currentSpeed > 30.0 {
		g.currentSpeed = 30.0
	}

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

	scoreText := strconv.Itoa(g.score) + " MILES"
		
		
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

	// loadin za gofer sproit
	gameOberOverlay, _, err := ebitenutil.NewImageFromFile("assets/gameober.png")
	if err != nil {
		log.Println("Failed to load game oberlay:", err)
	}


	// uhh ooga booga
	loreThingy, _, err := ebitenutil.NewImageFromFile("assets/lore.png")
	if err != nil {
		log.Println("Failed to load lore sprite:", err)
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


	// loadin bg images
	bg1, _, err := ebitenutil.NewImageFromFile("assets/bg.png")
	if err != nil {
		log.Println("Failed to load background image:", err)
	}
	bg2, _, err := ebitenutil.NewImageFromFile("assets/bg.png")
	if err != nil {
		log.Println("Failed to load background image:", err)
	}

	// gaem
	game := &Game{
		goferSprite: goferSprite,
		goferX: 400,
		goferY: 300,
		direction: 0,
		gameOber: false,
		gameOberOverlay: gameOberOverlay,
		bg1: bg1,
		bg2: bg2,
		bgY: 0,
		score: 0,
		currentSpeed: 10.0,
		loreThingy: loreThingy,
		gameState: 0,
		introDone: false,
		rockSprites: loadedRocks,
		rocks: [3]RockItem{
			{ x: float64(rand.Intn(700)), y: float64(rand.Intn(400)) + 850 },
			{ x: float64(rand.Intn(700)), y: float64(rand.Intn(400)) + 1300 },
			{ x: float64(rand.Intn(700)), y: float64(rand.Intn(400)) + 1550 },
		},
	}

	// run da gaem
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// ts pmod me so much