package math

import (
	"github.com/supercom32/dosktop/internal/recast"
	"fmt"
	"math"
	"strconv"
)

func GetAbsoluteValueAsFloat64(number interface{}) float64 {
	numberAsFloat64 := recast.GetNumberAsFloat64(number)
	if numberAsFloat64 < 0 {
		return -numberAsFloat64
	}
	return numberAsFloat64
}

func GetAbsoluteValueAsInt(number interface{}) int {
	numberAsInt := recast.GetNumberAsInt(number)
	if numberAsInt < 0 {
		return -numberAsInt
	}
	return numberAsInt
}

func RoundToWholeNumber(number interface{}) float64 {
	numberAsFloat64 := recast.GetNumberAsFloat64(number)
	return math.Round(numberAsFloat64)
}

func RoundToDecimal(number interface{}, numberOfPlaces int) float64 {
	numberAsFloat64 := recast.GetNumberAsFloat64(number)
	numberFormat := "%." + strconv.Itoa(numberOfPlaces) + "f"
	numberAsString := fmt.Sprintf(numberFormat, numberAsFloat64)
	roundedNumber, _ := strconv.ParseFloat(numberAsString, 64)
	return roundedNumber
}

func IsNumberEven(number interface{}) bool {
	numberAsInt64 := recast.GetNumberAsInt64(number)
	remainder := numberAsInt64 % 2
	if remainder == 0 {
		return true
	}
	return false
}

// IsFloatEffectivelyEqual This method allows you to check if two floating point numbers are effectively equal to each other.
// Since floating point operations perform approximate arithmetic, it is normal that there will be an accumulation
// of rounding errors in floating-point operations. By using this method, you can check if your numbers are
// for most practical purposes, equal or not by automatically rounding numbers down to 7 places.
func IsFloatEffectivelyEqual(firstNumber, secondNumber float64) bool {
	firstNumberRounded := RoundToDecimal(firstNumber, 7)
	secondNumberRounded := RoundToDecimal(secondNumber, 7)
	return firstNumberRounded == secondNumberRounded
}
