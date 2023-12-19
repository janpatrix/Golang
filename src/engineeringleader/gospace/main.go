package main

import (
	"embed"
	"fmt"
	b "gospace/basic"
	"os"
	"path/filepath"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImagesFromFolder(folder embed.FS) []*ebiten.Image {
	var images []*ebiten.Image

	// Walk through the folder and get all files with the .png extension
	err := filepath.WalkDir("basic/assets/meteors", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == "*.png" {
			fmt.Println("Loading image:", path)
			images = append(images, b.MustLoadImage(path))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return images
}

func main() {

	g := &Game{
		player:           NewPlayer(),
		meteorSpawnTimer: b.NewTimer(time.Second * 5),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
