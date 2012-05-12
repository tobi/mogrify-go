package mogrify

import (
  "io/ioutil"
  "log"
  "os"
  "runtime"
  "testing"
)

func assertDimension(t *testing.T, img *Image, expected string) {
  if actual := img.Dimensions(); actual != expected {
    t.Fatalf("Got wrong dimensions expected:%s got %s", expected, actual)
  }
}

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

    t.Fatalf("%d", img.Width())
  }

  if img.Height() != 399 {
    t.Fatalf("%d", img.Height())
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
  defer img.Destroy()

  if res != nil {
    t.FailNow()
  }

  assertDimension(t, img, "600x399")
}

func TestOpenBlopSuccessPng(t *testing.T) {
  bytes, _ := ioutil.ReadFile("./assets/example.com.png")

  img := NewImage()
  res := img.OpenBlob(bytes)
  defer img.Destroy()

  if res != nil {
    t.FailNow()
  }

  if dim := img.Dimensions(); dim != "1280x500" {
    t.Fatalf("Got wrong dimensions expected:1280x500 got %s", dim)
  }
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
  defer img.Destroy()

  img2, err := img.NewTransformation("", "100x50>")

  if err != nil {
    t.FailNow()
    return
  }

  defer img2.Destroy()

  assertDimension(t, img2, "75x50")

  img3, err := img.NewTransformation("", "100x50!")
  defer img3.Destroy()

  if err != nil {
    t.FailNow()
    return
  }

  assertDimension(t, img3, "100x50")

  //img2.SaveFile("/tmp/img4.jpg")
}

func TestReadFrom(t *testing.T) {
  file, _ := os.Open("./assets/image.jpg")
  image := NewImage()
  image.ReadFrom(file)
  assertDimension(t, image, "100x50")
}

func BenchmarkAndMemoryTest(b *testing.B) {
  var before runtime.MemStats
  var after runtime.MemStats

  runtime.ReadMemStats(&before)

  work := func() {
    img := Open("./assets/image.jpg")
    img.Destroy()
  }

  for i := 0; i < 100; i++ {
    work()
  }

  runtime.ReadMemStats(&after)

  log.Printf("sys memory before: %d after %d - diff: %d", before.HeapSys, after.HeapSys, after.HeapSys-before.HeapSys)
}
