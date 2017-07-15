package model

type Hasher interface {
	Hash() (int, error)
}
