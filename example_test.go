package mogrify_test

import (
	"fmt"
	"github.com/tobi/mogrify-go"
	"log"
	"os"
)

func ExampleDecodeJpeg() {

	file, err := os.Open("testdata/image.jpg")
	if err != nil {
		log.Fatalf("loading file: %s ", err)
	}

	jpg, err := mogrify.DecodeJpeg(file)
	if err != nil {
		log.Fatalf("decoding JPEG: %v", err)
	}
	defer jpg.Destroy()

	b := mogrify.Bounds{Width: 50, Height: 50}

	resized, err := jpg.NewResized(b)
	if err != nil {
		log.Fatalf("resizing JPEG: %s", err)
	}
	defer resized.Destroy()

	fmt.Printf("Resized image from %s => %s \n", mogrify.Dimensions(jpg), mogrify.Dimensions(resized))

	// Output:
	// Resized image from 600x399 => 50x50
}

func ExampleDecodeGif() {

	file, err := os.Open("testdata/image.gif")
	if err != nil {
		log.Fatalf("loading file: %s ", err)
	}

	jpg, err := mogrify.DecodeGif(file)
	if err != nil {
		log.Fatalf("decoding GIF: %v", err)
	}
	defer jpg.Destroy()

	b := mogrify.Bounds{Width: 50, Height: 50}

	resized, err := jpg.NewResized(b)
	if err != nil {
		log.Fatalf("resizing GIF: %s", err)
	}
	defer resized.Destroy()

	fmt.Printf("Resized image from %s => %s \n", mogrify.Dimensions(jpg), mogrify.Dimensions(resized))

	// Output:
	// Resized image from 458x399 => 50x50
}

func ExampleDecodePng() {

	file, err := os.Open("testdata/image.png")
	if err != nil {
		log.Fatalf("loading file: %s ", err)
	}

	jpg, err := mogrify.DecodePng(file)
	if err != nil {
		log.Fatalf("decoding PNG: %v", err)
	}
	defer jpg.Destroy()

	b := mogrify.Bounds{Width: 50, Height: 50}

	resized, err := jpg.NewResized(b)
	if err != nil {
		log.Fatalf("resizing PNG: %s", err)
	}
	defer resized.Destroy()

	fmt.Printf("Resized image from %s => %s \n", mogrify.Dimensions(jpg), mogrify.Dimensions(resized))

	// Output:
	// Resized image from 1280x500 => 50x50
}
