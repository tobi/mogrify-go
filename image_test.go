package mogrify

import (
	"bytes"
	"os"
	"testing"
)

func assertDimension(t *testing.T, img Image, expected string) {
	if actual := Dimensions(img); actual != expected {
		t.Errorf("Got wrong dimensions expected:%s got %s", expected, actual)
	}
}

func asset(asset string) Image {
	file, _ := os.Open("./assets/image.jpg")
	defer file.Close()

	image := DecodeJpg(file)

	if image == nil {
		panic("Image didnt load")
	}

	return image
}

func TestOpenExisting(t *testing.T) {
	image := asset("./assets/image.jpg")
	defer image.Destroy()

	assertDimension(t, image, "600x399")
}

func TestHeightWidth(t *testing.T) {
	img := asset("./assets/image.jpg")
	defer img.Destroy()

	if img.Width() != 600 {
		t.Fatalf("%d", img.Width())
	}

	if img.Height() != 399 {
		t.Fatalf("%d", img.Height())
	}
}

func TestResizeSuccess(t *testing.T) {
	img := asset("./assets/image.jpg")
	defer img.Destroy()

	resized, err := img.NewResized(50, 50)
	if err != nil {
		t.Error(err)
	}
	defer resized.Destroy()

	assertDimension(t, resized, "50x50")
}

func TestResampleSuccess(t *testing.T) {
	img := asset("./assets/image.jpg")
	defer img.Destroy()

	resized, err := img.NewResampled(50, 50)
	if err != nil {
		t.Error(err)
	}
	defer resized.Destroy()

	assertDimension(t, resized, "50x50")
}

func TestCrateFailure(t *testing.T) {
	image := NewImage(-1, -1)
	if image != nil {
		t.Fatalf("This should have failed...")
	}
}

func TestResampleFailure(t *testing.T) {
	img := asset("./assets/image.jpg")
	defer img.Destroy()

	_, err := img.NewResampled(-1, 50)
	if err == nil {
		t.Fatalf("This should have failed...")
	}
}

func TestDecodeEncode(t *testing.T) {
	img := asset("./assets/image.jpg")
	defer img.Destroy()

	resized, err := img.NewResampled(100, 100)

	if err != nil {
		t.Error(err)
		return
	}

	assertDimension(t, resized, "100x100")

	dest, _ := os.Create("/tmp/dest.jpg")
	defer dest.Close()

	var buffer bytes.Buffer

	_, err = EncodeJpg(&buffer, resized)

	if err != nil {
		t.Error(err)
		return
	}

	roundtrip := DecodeJpg(&buffer)

	assertDimension(t, roundtrip, "100x100")

}
