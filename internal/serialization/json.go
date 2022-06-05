package serialization

import "encoding/json"

type jsonSerializer struct {
}

func (s *jsonSerializer) Encode(i interface{}) ([]byte, error) {
	data, err := json.Marshal(i)
	return data, err
}

func (s *jsonSerializer) Decode(i []byte, o interface{}) error {
	return json.Unmarshal(i, o)
}

func NewJsonSerializer() *jsonSerializer {
	return &jsonSerializer{}
}
