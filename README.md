# package `mogrify`

Package mogrify binds to `libgd` to perform transformations on
images.  It supports GIF/JPEG/PNG images, which it can decode and
encode to. You can decode from a format and encode to another.

The current set of transformations is limited to cropping, resizing
and resampling.

# Install

You will need [`libgd`](http://libgd.bitbucket.org/) for the C bindings, then:

```
go get github.com/tobi/mogrify-go
```

# Docs?

[Godocs](http://godoc.org/github.com/tobi/mogrify-go)!

# Usage

```go
jpg, err := mogrify.DecodeJpeg(file)
if err != nil { log.Fatalf(err) }
defer jpg.Destroy()

b := mogrify.Bounds{Width: 50, Height: 50}

resized, err := jpg.NewResized(b)
if err != nil { log.Fatalf(err) }
defer resized.Destroy()

_, _ := mogrify.EncodeGif(newfile, resized)
if err != nil { log.Fatalf(err) }
```
