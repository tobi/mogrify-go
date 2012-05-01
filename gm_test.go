package mogrify

import (
  "io/ioutil"
  "log"
  "os"
  "testing"
)

func TestOpenExisting(t *testing.T) {
  img := Open("./assets/image.jpg")
  if img == nil {
    t.Fail()
  }
  img.Destroy()
}

func TestOpenNonExisting(t *testing.T) {
  if Open("./assets/image_does_not_exist.jpg") != nil {
    t.Fail()
  }
}

func TestHeightWidth(t *testing.T) {
  img := Open("./assets/image.jpg")
  if img == nil {
    t.Fail()
  }
  if img.Width() != 600 {
    log.Printf("%d", img.Width())
    t.Fail()
  }

  if img.Height() != 399 {
    log.Printf("%d", img.Height())
    t.Fail()
  }

}

func TestResizeSuccess(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  status := img.Resize(50, 50)
  if status != nil {
    log.Printf("resize failed %s", status)
    t.Fail()
  }

  if img.Width() != 50 || img.Height() != 50 {
    log.Printf("size was %dx%d", img.Width(), img.Height())
    t.Fail()
  }
}

func TestResizeFailure(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  status := img.Resize(0, 50)
  if status == nil {
    t.Fail()
  }
}

func TestSaveToSuccess(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  res := img.SaveFile("/tmp/img.jpg")
  if res != nil {
    t.Fail()
  }
}

func TestSaveToFailure(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  res := img.SaveFile("/dgksjogdsksdgsdkgsd;lfsd-does-not-exist/img.jpg")
  if res == nil {
    t.Fail()
  }
}

func TestOpenBlopSuccess(t *testing.T) {
  bytes, _ := ioutil.ReadFile("./assets/image.jpg")

  img := NewImage()
  res := img.OpenBlob(bytes)

  if res != nil {
    t.Fail()
  }

  img.Destroy()
}

func TestOpenBlopFailure(t *testing.T) {

  img := NewImage()
  res := img.OpenBlob([]byte{'a'})
  defer img.Destroy()

  if res == nil {
    t.Fail()
  }

  res = img.OpenBlob([]byte{})

  if res == nil {
    t.Fail()
  }
}

func TestSaveToBlob(t *testing.T) {
  img := Open("./assets/image.jpg")
  defer img.Destroy()

  fp, err := os.Create("/tmp/img3.jpg")
  if err != nil {
    t.Fail()
  }

  defer fp.Close()

  _, err = img.Write(fp)

  if err != nil {
    t.Fail()
  }

}

func TestTransformation(t *testing.T) {
  img := Open("./assets/image.jpg")
  img2 := img.NewTransformation("", "100x50>")
  defer img.Destroy()
  defer img2.Destroy()

  if img2.Dimensions() != "75x50" {
    t.Fail()
  }

  //img2.SaveFile("/tmp/img4.jpg")
}
