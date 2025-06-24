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

var ImgWorkerStatic = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/worker.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}()

var ImgPlatform = func() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(embedFS, "assets/platform.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
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
	imgs := make([]*ebiten.Image, 6)
	for i := range 6 {
		imgs[i] = ImgObstacleSheet.SubImage(image.Rect(i*TileSize, 0, i*TileSize+TileSize, TileSize)).(*ebiten.Image)
	}
	return imgs
}()

// Font in pixelfont
var ImgFont = func() *ebiten.Image {
	img, err := LoadPFImage(embedFS, "assets/Arcade_King_of_Fighters_97_Italic_(SNK).pf", color.NRGBA{R: 51, G: 57, B: 65, A: 255}, nil)
	if err != nil {
		log.Fatal("error loading font")
	}
	return ebiten.NewImageFromImage(img)
}()
