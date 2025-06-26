package game

import (
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var embedFS embed.FS

var ImgGameBg = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/gamebg.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()
var ImgGameBgDrawOp = func() *ebiten.DrawImageOptions {
	return &ebiten.DrawImageOptions{}
}()

const WorkerTileCount = 10

var ImgWorkerSheet = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/workersheet.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()
var ImgsWorker = func() []*ebiten.Image {
	imgs := make([]*ebiten.Image, WorkerTileCount)
	for i := range WorkerTileCount {
		imgs[i] = ImgWorkerSheet.SubImage(image.Rect(i*TileSize, 0, i*TileSize+TileSize, TileSize*2)).(*ebiten.Image)
	}
	return imgs
}()

var ImgForeman = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/foreman.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()
var ImgFormanDrawOp = func() *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(375, 141)
	return &op
}()

var ImgObstacleSheet = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/obstacles.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()

var ImgsObstacles = func() []*ebiten.Image {
	imgs := make([]*ebiten.Image, ObstacleTileCount)
	for i := range ObstacleTileCount {
		imgs[i] = ImgObstacleSheet.SubImage(image.Rect(i*TileSize, 0, i*TileSize+TileSize, TileSize)).(*ebiten.Image)
	}
	return imgs
}()

var ImgFaceSheet = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/facesheet.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()

var ImgsFaceSheet = func() []*ebiten.Image {
	// 7 mouth shapes and one blink
	imgs := make([]*ebiten.Image, 8)
	for i := range 7 {
		imgs[i] = ImgFaceSheet.SubImage(image.Rect(i*(TileSize*2), 0, i*(TileSize*2)+(TileSize*2), TileSize)).(*ebiten.Image)
	}
	imgs[7] = ImgFaceSheet.SubImage(image.Rect(224, 0, 224+(TileSize*3), TileSize)).(*ebiten.Image)
	return imgs
}()
var ImgBlinkOp = func() *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(384, 186)
	return &op
}()
var ImgMouthOp = func() *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(397, 210)
	return &op
}()

var ImgPlatform = ImgsObstacles[TilePlatform]

// Font in pixelfont
var ImgFont = func() *ebiten.Image {
	img, err := LoadPFImage(embedFS, "assets/Arcade_King_of_Fighters_97_Italic_(SNK).pf", color.NRGBA{R: 51, G: 57, B: 65, A: 255}, nil)
	if err != nil {
		log.Fatal("error loading font")
	}
	return ebiten.NewImageFromImage(img)
}()

var CM = MakeCharacterMap(ImgFont)
