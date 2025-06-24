package game

import (
	"embed"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PFData []byte

func OpenPFFile(fs embed.FS, p string) (PFData, error) {
	buf, err := fs.ReadFile(p)
	if err != nil {
		return make(PFData, 0), err
	}
	return buf, nil
}

func PFToPalettedImage(pf PFData, fontColor color.Color, bgColor color.Color) image.Image {
	if fontColor == nil {
		fontColor = color.Black
	}
	if bgColor == nil {
		bgColor = color.Transparent
	}
	// 128x128 (16*8)x(16*8)
	// Every 8 bytes is a character
	// Use pallated colors
	p := color.Palette{
		bgColor,
		fontColor,
	}
	img := image.NewPaletted(image.Rect(0, 0, 128, 128), p)

	// Ugh. This should be doable with Pix[] offset but my brain fried out on the math
	for byi, byt := range pf {
		charInd := (byi / 8)
		charRow := (byi / 8) / 16
		charCol := (charInd % 16)
		py := (byi % 8) + (charRow * 8)
		px := charCol * 8
		for bti, btt := range []uint8{7, 6, 5, 4, 3, 2, 1, 0} {
			var c color.Color = bgColor
			if (byt>>btt)&1 == 1 {
				c = fontColor
			}

			img.Set(px+bti, py, c)
		}
	}
	return img
}

func LoadPFImage(fs embed.FS, path string, fontColor color.Color, bgColor color.Color) (image.Image, error) {
	pfd, err := OpenPFFile(fs, path)
	if err != nil {
		return nil, err
	}
	return PFToPalettedImage(pfd, fontColor, bgColor), nil
}

const AsciiOffset = 32

type CharacterMap struct {
	CharImgs []*ebiten.Image // Indexed from ASCII offset
}

func (cm *CharacterMap) GetFromIndex(i uint8) *ebiten.Image {
	if int(i) >= len(cm.CharImgs) {
		return nil
	}
	return cm.CharImgs[i]
}
func (cm *CharacterMap) GetFromRune(r rune) *ebiten.Image {
	if r > 255+AsciiOffset || r < AsciiOffset {
		r = 0
	}
	return cm.GetFromIndex(uint8(r - AsciiOffset))
}
func (cm *CharacterMap) GetFromString(s string) *ebiten.Image {
	if len(s) >= 1 {
		for _, r := range s[0:1] {
			return cm.GetFromRune(r)
		}
	}
	return cm.GetFromRune(0)
}
func MakeCharacterMap(fontImg *ebiten.Image) *CharacterMap {
	imgs := make([]*ebiten.Image, 256)
	for i := range 256 {
		charRow := i / 16
		charCol := (i % 16)
		oy := charRow * 8
		ox := charCol * 8
		imgs[i] = fontImg.SubImage(image.Rect(ox, oy, ox+8, oy+8)).(*ebiten.Image)
	}
	return &CharacterMap{
		CharImgs: imgs,
	}
}

type FontText struct {
	X         int16
	Y         int16
	LineSpace uint8
	CharImgs  [][]*ebiten.Image
	CharMap   *CharacterMap
}

func (ft *FontText) SetText(txts []string) {
	rowMax := len(txts)
	ft.CharImgs = make([][]*ebiten.Image, rowMax)
	lineMax := 0
	for li, ln := range txts {
		lnLen := len(ln)
		ft.CharImgs[li] = make([]*ebiten.Image, lnLen)
		if lnLen > lineMax {
			lineMax = lnLen
		}
		for si, r := range ln {
			ft.CharImgs[li][si] = ft.CharMap.GetFromRune(r)
		}
	}
}
func (ft *FontText) Draw(dstImg *ebiten.Image) {
	for li, ln := range ft.CharImgs {
		for ci, cim := range ln {
			ob := ebiten.DrawImageOptions{}
			ls := 0
			if li > 0 {
				ls = int(ft.LineSpace) * li
			}
			ob.GeoM.Translate(
				float64(int(ft.X)+ci*8),
				float64(int(ft.Y)+li*8+ls),
			)
			dstImg.DrawImage(cim, &ob)
		}
	}
}

func MakeFontText(cm *CharacterMap, txts []string) *FontText {
	ft := &FontText{
		CharMap: cm,
	}
	ft.SetText(txts)
	return ft
}
