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

var BGColor = color.NRGBA{R: 185, G: 191, B: 251, A: 255}

func LoadEbitenImgFatal(p string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, p)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func MakeTranslateOp(x float64, y float64) *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	return &op
}

var ImgGameBg = LoadEbitenImgFatal("assets/gamebg.png")
var ImgGameBgDrawOp = func() *ebiten.DrawImageOptions {
	return &ebiten.DrawImageOptions{}
}()

const WorkerTileCount = 10

var ImgWorkerSheet = LoadEbitenImgFatal("assets/workersheet.png")
var ImgsWorker = func() []*ebiten.Image {
	imgs := make([]*ebiten.Image, WorkerTileCount)
	for i := range WorkerTileCount {
		imgs[i] = ImgWorkerSheet.SubImage(image.Rect(i*TileSize, 0, i*TileSize+TileSize, TileSize*2)).(*ebiten.Image)
	}
	return imgs
}()

var ImgForeman = LoadEbitenImgFatal("assets/foreman.png")
var ImgFormanDrawOp = MakeTranslateOp(375, 141)

var ImgObstacleSheet = LoadEbitenImgFatal("assets/obstacles.png")
var ImgsObstacles = func() []*ebiten.Image {
	imgs := make([]*ebiten.Image, ObstacleTileCount)
	for i := range ObstacleTileCount {
		imgs[i] = ImgObstacleSheet.SubImage(image.Rect(i*TileSize, 0, i*TileSize+TileSize, TileSize)).(*ebiten.Image)
	}
	return imgs
}()

var ImgFaceSheet = LoadEbitenImgFatal("assets/facesheet.png")
var ImgsFaceSheet = func() []*ebiten.Image {
	// 7 mouth shapes and one blink
	imgs := make([]*ebiten.Image, 8)
	for i := range 7 {
		imgs[i] = ImgFaceSheet.SubImage(image.Rect(i*(TileSize*2), 0, i*(TileSize*2)+(TileSize*2), TileSize)).(*ebiten.Image)
	}
	imgs[7] = ImgFaceSheet.SubImage(image.Rect(224, 0, 224+(TileSize*3), TileSize)).(*ebiten.Image)
	return imgs
}()
var ImgBlinkOp = MakeTranslateOp(384, 186)
var ImgMouthOp = MakeTranslateOp(397, 210)

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

var ImgTitlecard = LoadEbitenImgFatal("assets/titlecard.png")
var ImgTitlecardDrawOp = MakeTranslateOp(128, 32)

var ImgNameHightlight = LoadEbitenImgFatal("assets/namehighlight.png")
