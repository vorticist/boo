package assets

import (
	"image"
	"os"
	"path/filepath"
	"strings"

	"gioui.org/font/opentype"
	"gioui.org/text"
	"gitlab.com/vorticist/logger"
)

var (
	Icons = map[string]image.Image{}
	Fonts = []text.FontFace{}
)

func Load() {
	loadFonts()
}

func loadFonts() {
	fontPath := "/home/vorticist/boo/assets/fonts/"
	files := []string{}
	err := filepath.Walk(fontPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".ttf" {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		logger.Errorf("error walking the path %s: %s", fontPath, err)
	}
	for _, file := range files {
		name := filepath.Base(file)
		name = strings.Replace(name, filepath.Ext(file), "", -1)
		ttBytes, err := os.ReadFile(file)
		if err != nil {
			logger.Errorf("could not load font: %v", err)
			continue
		}
		face, err := opentype.Parse(ttBytes)
		if err != nil {
			logger.Errorf("could not parse font: %v", err)
			continue
		}

		Fonts = append(Fonts, text.FontFace{
			Font: text.Font{Typeface: text.Typeface(name)},
			Face: face,
		})
	}
}
