package serialization

type Serializer interface {
	Marshaller
	Unmarshaller
}

type Marshaller interface {
	Marshal(i interface{}) ([]byte, error)
}

type Unmarshaller interface {
	Unmarshal(i []byte, o interface{}) error
}
