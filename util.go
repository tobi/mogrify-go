package mogrify

import (
  "bytes"
  "io"
)

func drain(reader io.Reader) []byte {
  var buffer bytes.Buffer

  _, err := buffer.ReadFrom(reader)
  if err != nil {
    return []byte{}
  }

  return buffer.Bytes()
}
