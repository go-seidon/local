package serialization

type Serializer interface {
	Encoder
	Decoder
}

type Encoder interface {
	Encode(i interface{}) ([]byte, error)
}

type Decoder interface {
	Decode(i []byte, o interface{}) error
}
