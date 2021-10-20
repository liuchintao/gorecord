package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		panic("please input path")
	}
	p := os.Args[1]
	p, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		log.Println("visit", path, "isDir", d.IsDir(), "error", err)
		if d.IsDir() && path != p {
			return filepath.SkipDir
		}
		if err != nil {
			return err
		}
		return nil
	})
}
