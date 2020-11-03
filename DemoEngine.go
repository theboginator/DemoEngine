package DemoEngine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const(
	ScreenWidth = 800
	ScreenHeight = 700
)

type Sprite struct {
	pict *ebiten.Image
	xLoc 	int
	yLoc 	int
	dx		int
	dy 		int
}

type Game struct {
	playerSprite Sprite
	drawOps ebiten.DrawImageOptions
}

func main(){
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("APEX Legends except way worse")
	gameObject := Game{}
	loadImage(&gameObject)
}

func loadImage(g *Game){
	pict, _, err := ebitenutil.NewImageFromFile("galleon.png")
}