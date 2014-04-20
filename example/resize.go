package main

import (
	"fmt"
	"github.com/tobi/mogrify-go"
	"log"
	"os"
)

func main() {

	file, err := os.Open("../assets/image.jpg")

	if err != nil {
		log.Fatalf("Error loading file: %s ", err)
	}

	jpg := mogrify.DecodeJpeg(file)
	if jpg == nil {
		log.Fatalln("could not load image")
	}
	defer jpg.Destroy()

	b := mogrify.Bounds{Width: 50, Height: 50}

	resized, err := jpg.NewResized(b)
	if err != nil {
		log.Fatalf("failed to resize: %s", err)
	}
	defer resized.Destroy()

	fmt.Printf("Resized image from %s => %s \n", mogrify.Dimensions(jpg), mogrify.Dimensions(resized))
}
