package ui

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

var initilized = false
var fontFaces map[string]font.Face

func GetFontNormalFace() font.Face {
	if !initilized {
		initFonts()
	}

	return fontFaces["faceNormal"]
}

func initFonts() {

	fontFaces = make(map[string]font.Face)

	f, err := opentype.Parse(goregular.TTF)

	if err != nil {
		log.Fatal(err)
	}

	faceNormal, _ := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	fontFaces["faceNormal"] = faceNormal

	initilized = true

}
