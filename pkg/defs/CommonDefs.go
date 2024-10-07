package defs

import (
	"fmt"
	"strconv"
)
type TDecimal string
type TTime float64

type TNumber int32

// multiplicator for money and game multiplicator
const X_100 = 100


func Number2Float(num TNumber) float64{

	return (float64(num))/X_100;
}
func Float2Number(num float64) TNumber{

	return TNumber(num * X_100);
}

func Number2Dec(num TNumber) TDecimal{

	var val = fmt.Sprintf("%.2f", Number2Float(num) );
	return TDecimal(val);
}
func Dec2Number(num TDecimal) TNumber{

	var fNum, err = strconv.ParseFloat( string(num), 64 );
	if( err != nil){
		return 0;
	}
	return TNumber(fNum * X_100);
}