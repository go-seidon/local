package text

import (
	"github.com/segmentio/ksuid"
)

type Identifier interface {
	GenerateId() (string, error)
}

type ksuIdentifier struct {
}

func (i *ksuIdentifier) GenerateId() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func NewKsuid() *ksuIdentifier {
	return &ksuIdentifier{}
}
