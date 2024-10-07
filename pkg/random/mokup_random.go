package random

import (
	"math/rand"
	"time"
)

type Mokup_Random struct {
	rand float64
}

func (This *Mokup_Random) Constr() IRandom{

	rand.Seed(time.Now().UnixNano())
	return This;	
}

func (This *Mokup_Random) GetRandom() float64 {

	var randomNumber = rand.Float64() 
	return randomNumber;
}

func (This *Mokup_Random) GetRandom2(limitStart float64, limitEnd float64) float64 {

	var randomNumber = rand.Float64() 
	return randomNumber * ( limitEnd - limitStart) + limitStart;
}


