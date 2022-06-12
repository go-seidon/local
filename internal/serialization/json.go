package serialization

import "encoding/json"

type jsonSerializer struct {
}

func (s *jsonSerializer) Encode(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func (s *jsonSerializer) Decode(i []byte, o interface{}) error {
	return json.Unmarshal(i, o)
}

func NewJsonSerializer() *jsonSerializer {
	return &jsonSerializer{}
}
