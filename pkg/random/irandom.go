package random

type IRandom interface {
	Constr() IRandom
	GetRandom() float64
	GetRandom2(limitStart float64, limitEnd float64) float64
}