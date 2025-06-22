package game

import (
	"embed"
	"image"
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
	imgs := make([]*ebiten.Image, 5)
	for i := range 5 {
		imgs[i] = ImgObstacleSheet.SubImage(image.Rect(i*TileSize, 0, i*TileSize+TileSize, TileSize)).(*ebiten.Image)
	}
	return imgs
}()
