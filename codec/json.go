package codec

import (
  "encoding/json"
)

type jsonCodec struct{
}

var JSON = jsonCodec{}

func (c jsonCodec) Dump(message interface{}) []byte {
  data, err := json.Marshal(message)
  if err != nil {
  }
  return data
}
