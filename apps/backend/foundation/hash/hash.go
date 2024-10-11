package hash

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(val []byte) (string, error)
	Compare(val []byte, hash []byte) bool
}

type bcryptHasher struct{}

func NewBcryptHasher() *bcryptHasher {
	return &bcryptHasher{}
}

func (h *bcryptHasher) Hash(val []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(val, 13)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *bcryptHasher) Compare(val []byte, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, val) == nil
}
