package main

import (
	"fmt"

	"github.com/jammy-dodgers/gophenslide/openslide"
)

const SLIDE_PAHT = "./assets/slide1.svs"

func main() {
	slide, err := openslide.Open(SLIDE_PAHT)
	if err != nil {
		panic("hello")
	}
	fmt.Println(slide)
}
