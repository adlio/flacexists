package main

// Test
import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var flacDir = "/Volumes/MEDIA/FLAC/"

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Printf("root: %s\n", root)

	dirs, _ := ioutil.ReadDir(root)
	for _, dir := range dirs {
		artistName := dir.Name()
		artistDir := path.Join(root, artistName)
		albumDirs, _ := ioutil.ReadDir(artistDir)

		missingAlbums := make([]string, 0)

		for _, albumDir := range albumDirs {
			if !albumDir.IsDir() {
				continue
			}
			albumName := albumDir.Name()

			// Check if FLAC Directory exists
			_, err := os.Stat(path.Join(flacDir, artistName, albumName))
			if err != nil {
				missingAlbums = append(missingAlbums, albumName)
			}
		}

		if len(missingAlbums) > 0 {
			fmt.Printf("%s\n", artistName)
			for _, missingAlbum := range missingAlbums {
				fmt.Printf("  : %s\n", missingAlbum)
			}
		}
	}
}
