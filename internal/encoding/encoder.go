package encoding

type Encoder interface {
	Encodeable
	Decodeable
}

type Encodeable interface {
	Encode(d []byte) (string, error)
}

type Decodeable interface {
	Decode(d string) ([]byte, error)
}
