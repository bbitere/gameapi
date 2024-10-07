package random

type IRandom interface {
	Constr() *IRandom
	GetRandom() float64
}