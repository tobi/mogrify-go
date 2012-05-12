package main

import (
	"github.com/tobi/mogrify-go"
)

func main() {
  img := mogrify.NewImage()
  defer img.Destroy()

  if err := img.OpenFile("../assets/example.com.png"); err != nil {
    panic(err.Error())
  }
}
