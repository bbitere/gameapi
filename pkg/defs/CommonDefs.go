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
func Float2Number(num float64) int32 { //TNumber{

	return int32(num * X_100);
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


func Float2Dec(num float64) TDecimal{

	var val = fmt.Sprintf("%.2f", num );
	return TDecimal(val);
}
func Dec2Float(num TDecimal) float64{

	var fNum, err = strconv.ParseFloat( string(num), 64 );
	if( err != nil){
		return 0;
	}
	return fNum;
}