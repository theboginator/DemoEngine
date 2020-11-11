package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"testing"
)

func TestGot_gold(t *testing.T) {
	picture, _, err := ebitenutil.NewImageFromFile("galleon.png")
	if err != nil {
		log.Fatal("Failed to load image ", err)
	}
	player := Sprite{xLoc: 100, yLoc: 100, pict: picture}
	coins, _, err := ebitenutil.NewImageFromFile("gold-coins.png")
	if err != nil {
		log.Fatal("Failed to load image ", err)
	}
	gold := Sprite{xLoc: 500, yLoc: 500, pict: coins}
	result := gotGold(player, gold)
	if result != false {
		t.Error("Wrongly found a collision")
	}
}
