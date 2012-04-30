package mogrify

import (
  "log"
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
