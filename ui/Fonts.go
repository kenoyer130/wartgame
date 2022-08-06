package ui

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var initilized = false
var fontFaces map[string]font.Face

func GetFontNormalFace() font.Face {
	if !initilized {
		initFonts()
	}

	return fontFaces["faceNormal"]
}

func GetFontBold() font.Face {
	if !initilized {
		initFonts()
	}

	return fontFaces["faceBold"]
}

func GetFontItalic() font.Face {
	if !initilized {
		initFonts()
	}

	return fontFaces["faceItalic"]
}

func initFonts() {

	fontFaces = make(map[string]font.Face)

	fontFaces["faceNormal"] = createFaceNormal()
	fontFaces["faceBold"] = createFaceBold()
	fontFaces["faceItalic"] = createFaceItalic()
	fontFaces["faceTiny"] = createFaceTiny()

	initilized = true
}

func createFaceNormal() font.Face {
	f, err := opentype.Parse(goregular.TTF)

	if err != nil {
		log.Fatal(err)
	}

	return createFontFace(f)
}

func createFaceBold() font.Face {

	f, err := opentype.Parse(gobold.TTF)

	if err != nil {
		log.Fatal(err)
	}

	return createFontFace(f)
}

func GetFontTiny() font.Face {
	if !initilized {
		initFonts()
	}

	return fontFaces["faceTiny"]
}

func createFaceItalic() font.Face {

	f, err := opentype.Parse(goitalic.TTF)

	if err != nil {
		log.Fatal(err)
	}

	return createFontFace(f)
}

func createFontFace(f *sfnt.Font) font.Face {
	face, _ := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	return face
}

func createFaceTiny() font.Face {
	f, err := opentype.Parse(goregular.TTF)

	if err != nil {
		log.Fatal(err)
	}

	return createTinyFace(f)
}

func createTinyFace(f *sfnt.Font) font.Face {
	face, _ := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    8,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	return face
}
