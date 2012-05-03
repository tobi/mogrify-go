package main

import (
  "fmt"
	"github.com/tobi/mogrify-go"
)

func main() {

  img := mogrify.NewImage()
  defer img.Destroy()
  
  if err := img.OpenFile("../assets/example.com.png"); err != nil {
    panic(err.Error())
  }

  fmt.Printf("Image dimensions: %s\n", img.Dimensions())
  
  resized, err := img.NewTransformation("", "100x75>")
  
  if err != nil {
    panic(err.Error())
  }
  defer resized.Destroy()

  fmt.Printf("Resized image dimensions: %s\n", resized.Dimensions())

  resized.SaveFile("/tmp/image.png")

}
