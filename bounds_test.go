package mogrify

import (
	"testing"
)

func TestBounds(t *testing.T) {

	bounds := Bounds{100, 50}
	if bounds.Width != 100 {
		t.FailNow()
	}

	if bounds.Height != 50 {
		t.FailNow()
	}

}

func TestProportionalOperation(t *testing.T) {
	// half
	bounds := Bounds{100, 50}.ScaleProportionally(50, 50)
	if bounds.Width != 50 {
		t.FailNow()
	}
	if bounds.Height != 25 {
		t.FailNow()
	}

	// no changes
	bounds = Bounds{100, 50}.ScaleProportionally(100, 100)
	if bounds.Width != 100 {
		t.FailNow()
	}
	if bounds.Height != 50 {
		t.FailNow()
	}

	// no changes
	bounds = Bounds{100, 50}.ScaleProportionally(100, 100000000)
	if bounds.Width != 100 {
		t.FailNow()
	}
	if bounds.Height != 50 {
		t.FailNow()
	}

	// scale up
	bounds = Bounds{100, 50}.ScaleProportionally(1000, 1000)
	if bounds.Width != 1000 {
		t.FailNow()
	}
	if bounds.Height != 500 {
		t.FailNow()
	}
}

func TestShrink(t *testing.T) {
	// half
	bounds := Bounds{100, 50}.ShrinkProportionally(50, 50)
	if bounds.Width != 50 {
		t.Errorf("Width is wrong: %d", bounds.Width)
	}
	if bounds.Height != 25 {
		t.Errorf("Height is wrong: %d", bounds.Height)
	}

	// no changes
	bounds = Bounds{100, 50}.ShrinkProportionally(100000, 100000)
	if bounds.Width != 100 {
		t.Errorf("Width is wrong: %d", bounds.Width)
	}
	if bounds.Height != 50 {
		t.Errorf("Height is wrong: %d", bounds.Height)
	}
}

func TestGrow(t *testing.T) {
	// no changes
	bounds := Bounds{100, 50}.GrowProportionally(50, 50)
	if bounds.Width != 100 {
		t.Errorf("Width is wrong: %d", bounds.Width)
	}
	if bounds.Height != 50 {
		t.Errorf("Height is wrong: %d", bounds.Height)
	}

	// no changes
	bounds = Bounds{100, 50}.GrowProportionally(100000, 100000)
	if bounds.Width != 100000 {
		t.Errorf("Width is wrong: %d", bounds.Width)
	}
	if bounds.Height != 50000 {
		t.Errorf("Height is wrong: %d", bounds.Height)
	}
}

var stringTests = []struct {
	name      string
	bounds    string
	want      *Bounds
	shouldErr bool
}{
	{"Two proper sizes", "100x50", &Bounds{100, 50}, false},
	{"Only height", "x50", &Bounds{0, 50}, false},
	{"Only width", "100x", &Bounds{100, 0}, false},
	{"An invalid bound", "not a bound", nil, true},
}

func TestFromString(t *testing.T) {
	for _, tt := range stringTests {
		t.Logf("Extracting bounds: %s", tt.name)
		bounds, err := BoundsFromString(tt.bounds)

		switch {

		case tt.shouldErr && err == nil:
			t.Error("want an error, got nothing")

		case !tt.shouldErr && err != nil:
			t.Errorf("want no error, got '%v'", err)

		case tt.want == nil && bounds != nil:
			t.Errorf("want nil bound, got '%#v'", bounds)

		case tt.want != nil && bounds == nil:
			t.Errorf("want '%#v', got nothing", tt.want)

		}

		if tt.shouldErr {
			continue // tt.bounds will be nil
		}

		// not an error case, bounds is not nil
		if tt.want.Width != bounds.Width {
			t.Errorf("want width %d, got %d", tt.want.Width, bounds.Width)
		}

		if tt.want.Height != bounds.Height {
			t.Errorf("want height %d, got %d", tt.want.Height, bounds.Height)
		}
	}
}
