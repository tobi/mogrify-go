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
  if bounds.Width != 50 {  t.FailNow()  }
  if bounds.Height != 25 { t.FailNow()  }

  // no changes
  bounds = Bounds{100, 50}.ScaleProportionally(100, 100)
  if bounds.Width != 100 { t.FailNow()  }
  if bounds.Height != 50 { t.FailNow()  }

  // no changes
  bounds = Bounds{100, 50}.ScaleProportionally(100, 100000000)
  if bounds.Width != 100 { t.FailNow()  }
  if bounds.Height != 50 { t.FailNow()  }

  // scale up
  bounds = Bounds{100, 50}.ScaleProportionally(1000, 1000)
  if bounds.Width != 1000 { t.FailNow()  }
  if bounds.Height != 500 { t.FailNow()  }
}

func TestShrink(t *testing.T) {
  // half
  bounds := Bounds{100, 50}.ShrinkProportionally(50, 50)
  if bounds.Width != 50 { t.Errorf("Width is wrong: %d", bounds.Width) }
  if bounds.Height != 25 { t.Errorf("Height is wrong: %d", bounds.Height) }

  // no changes
  bounds = Bounds{100, 50}.ShrinkProportionally(100000, 100000)
  if bounds.Width != 100 { t.Errorf("Width is wrong: %d", bounds.Width) }
  if bounds.Height != 50 { t.Errorf("Height is wrong: %d", bounds.Height) }
}

func TestGrow(t *testing.T) {
  // no changes
  bounds := Bounds{100, 50}.GrowProportionally(50, 50)
  if bounds.Width != 100 { t.Errorf("Width is wrong: %d", bounds.Width) }
  if bounds.Height != 50 { t.Errorf("Height is wrong: %d", bounds.Height) }

  // no changes
  bounds = Bounds{100, 50}.GrowProportionally(100000, 100000)
  if bounds.Width != 100000 { t.Errorf("Width is wrong: %d", bounds.Width) }
  if bounds.Height != 50000 { t.Errorf("Height is wrong: %d", bounds.Height) }
}

func TestFromString(t *testing.T) {
  bounds := BoundsFromString("100x50")
  if bounds.Width != 100 { t.Errorf("Width is wrong: %d", bounds.Width) }
  if bounds.Height != 50 { t.Errorf("Height is wrong: %d", bounds.Height) }
}

