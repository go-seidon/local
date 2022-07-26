package hashing

type Hasher interface {
	Generator
	Verificator
}

type Generator interface {
	Generate(src string) ([]byte, error)
}

type Verificator interface {
	Verify(hash string, text string) error
}
