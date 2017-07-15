package models

type Hasher interface {
	Hash() (int, error)
}
