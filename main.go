package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jammy-dodgers/gophenslide/openslide"
)

const SLIDE_PAHT = "./assets/slide1.svs"

func main() {

	//Detecting Vendor
	vendor, err := openslide.DetectVendor(SLIDE_PAHT)
	if err != nil {
		panic("Failed to load image")
	} else if err == nil && vendor == "" {
		panic("Err nil but vendor blank")
	}
	fmt.Println("Vendor:", vendor)

	//OPEN
	slide, err := openslide.Open(SLIDE_PAHT)
	defer slide.Close()
	if err != nil {
		panic("Failed to open image")
	}
	fmt.Println(slide)

	//Properties
	props := slide.PropertyNames()
	for i := 0; i < len(props); i++ {
		fmt.Println(props[i], "=", slide.PropertyValue(props[i]))
	}

	//Read Region
	bytes, err := slide.ReadRegion(10, 10, 6, 400, 400)
	if err != nil {
		panic(err.Error())
	}
	const testRawFilename = "testdata/raw_region.data"
	if info, e := os.Stat(testRawFilename); os.IsExist(e) && !info.IsDir() {
		if remErr := os.Remove(testRawFilename); remErr != nil {
			fmt.Println("Could not remove file ", testRawFilename)
		}
	}
	writeErr := ioutil.WriteFile(testRawFilename, bytes, 0660)
	if writeErr != nil {
		panic(writeErr.Error())
	}

	//LEVEL
	levels := slide.LevelCount()
	if levels == -1 {
		panic("An error has occured")
	}
	w, h := slide.LargestLevelDimensions()
	fmt.Println("Base lvl0 (", w, ", ", h, "): ", slide.LevelDownsample(0))
	for i := int32(1); i < levels; i++ {
		w, h = slide.LevelDimensions(i)
		fmt.Println("Level ", i, " (", w, ", ", h, "): ", slide.LevelDownsample(i))
	}
	//PRINT Version
	version := openslide.Version()
	fmt.Println("Openslide Version: ", version)
}
