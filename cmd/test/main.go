package main

import (
	"log"

	"github.com/nanoteck137/sewaddle-core/library"
)

func main() {
	lib, err := library.ReadFromDir("/Volumes/media/manga")
	if err != nil {
		log.Fatal(err)
	}

	_ = lib

	// pretty.Println(lib)
}
