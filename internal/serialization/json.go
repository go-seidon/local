package serialization

import "encoding/json"

type jsonSerializer struct {
}

func (s *jsonSerializer) Marshal(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func (s *jsonSerializer) Unmarshal(i []byte, o interface{}) error {
	return json.Unmarshal(i, o)
}

func NewJsonSerializer() *jsonSerializer {
	return &jsonSerializer{}
}
