package hashing

import "golang.org/x/crypto/bcrypt"

type bcryptHasher struct {
}

func (h *bcryptHasher) Generate(src string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
}

func (h *bcryptHasher) Verify(hash string, text string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
}

func NewBcryptHasher() *bcryptHasher {
	h := &bcryptHasher{}
	return h
}
