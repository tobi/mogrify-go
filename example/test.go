package main

import (
  "io/ioutil"
	"github.com/tobi/mogrify-go"
)

// func Server(w http.ResponseWriter, r *http.Request) {
// }

func init() {
  mogrify.Init()
}

func main() {
  

	img := mogrify.NewImage()
	defer img.Destroy()
	bytes, _ := ioutil.ReadFile("image.png")

  defer print("\n")

	err := img.OpenBlob(bytes)

  if err != nil {
    print("error ")
    print(err.Error())
    return
  }

  print(img.Dimensions())

  
  
}
