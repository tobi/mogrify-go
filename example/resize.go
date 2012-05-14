package main

import (
  "os"
  "fmt"
	"github.com/tobi/mogrify-go"
)

func main() {

  file, err := os.Open("../assets/image.jpg")

  if err != nil {
    fmt.Errorf("Error loading file: %s ", err)
    os.Exit(1)
    return
  }

  jpg := mogrify.NewJpg(file)
  
  if jpg == nil {
    fmt.Println("could not load image")
    os.Exit(1)
    return
  }

  defer jpg.Destroy()

  resized, err := jpg.NewResized(50, 50)

  if err != nil {
    fmt.Errorf("failed to resize: %s", err)
    os.Exit(1)
    return
  } 

  defer resized.Destroy()

  fmt.Printf("Resized image from %s => %s \n", mogrify.Dimensions(jpg), mogrify.Dimensions(resized))
}
