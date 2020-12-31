package math

import (
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/supercom32/dosktop/internal/stringformat"
	"reflect"
	"testing"
)

func TestGetAbsoluteValue(test *testing.T) {
	var intValues = recast.GetArrayOfInterfaces(int(-3), int8(-3), int16(-3), int32(-3), int64(-3))
	var uIntValues = recast.GetArrayOfInterfaces(uint(3), uint8(3), uint16(3), uint32(3), uint64(3))
	var floatValues = recast.GetArrayOfInterfaces(float32(-3.3), float64(-3.3))
	for _, currentValue := range intValues {
		absoluteValue := GetAbsoluteValueAsFloat64(currentValue)
		if absoluteValue != 3 {
			test.Errorf("The absolute value of '-3' given as type '" + reflect.TypeOf(currentValue).String() + "' should have been '3', but '" + stringformat.GetIntAsString(absoluteValue) + "' was received instead.")
		}
	}
	for _, currentValue := range uIntValues {
		absoluteValue := GetAbsoluteValueAsFloat64(currentValue)
		if absoluteValue != 3 {
			test.Errorf("The absolute value of '3' given as type '" + reflect.TypeOf(currentValue).String() + "' should have been '3', but '" + stringformat.GetIntAsString(absoluteValue) + "' was received instead.")
		}
	}
	for _, currentValue := range floatValues {
		absoluteValue := GetAbsoluteValueAsFloat64(currentValue)
		if !IsFloatEffectivelyEqual(absoluteValue, 3.3){
			test.Errorf("The absolute value of '-3.3' given as type '" + reflect.TypeOf(currentValue).String() + "' should have been '3.3', but '" + stringformat.GetIntAsString(absoluteValue) + "' was received instead.")
		}
	}
}

func TestIsNumberEven(test *testing.T) {
	var intValues = recast.GetArrayOfInterfaces(int(12), int8(12), int16(12), int32(12), int64(12))
	var uIntValues = recast.GetArrayOfInterfaces(uint(13), uint8(13), uint16(13), uint32(13), uint64(13))
	var floatValues = recast.GetArrayOfInterfaces(float32(12.3), float64(12.3))
	for _, currentValue := range intValues {
		isEven := IsNumberEven(currentValue)
		if isEven != true {
			test.Errorf("The value of '12' given as type '" + reflect.TypeOf(currentValue).String() + "' should return 'true' if checked as an even number.")
		}
	}
	for _, currentValue := range uIntValues {
		isEven := IsNumberEven(currentValue)
		if isEven != false {
			test.Errorf("The value of '13' given as type '" + reflect.TypeOf(currentValue).String() + "' should return 'false' if checked as an even number.")
		}
	}
	for _, currentValue := range floatValues {
		isEven := IsNumberEven(currentValue)
		if isEven != true {
			test.Errorf("The value of '12.3' given as type '" + reflect.TypeOf(currentValue).String() + "' should return 'false' if checked as an even number.")
		}
	}
}

func TestRoundToWholeNumber(test *testing.T) {
	var intValues = recast.GetArrayOfInterfaces(int(5), int8(5), int16(5), int32(5), int64(5))
	var uIntValues = recast.GetArrayOfInterfaces(uint(5), uint8(5), uint16(5), uint32(5), uint64(5))
	var floatValues = recast.GetArrayOfInterfaces(float32(5.5), float64(5.5))
	for _, currentValue := range intValues {
		result := RoundToWholeNumber(currentValue)
		if result != 5 {
			test.Errorf("The value of '5' given as type '" + reflect.TypeOf(currentValue).String() + "' should be rounded to '5', but was rounded to '" + stringformat.GetIntAsString(result) +"' instead.")
		}
	}
	for _, currentValue := range uIntValues {
		result := RoundToWholeNumber(currentValue)
		if result != 5 {
			test.Errorf("The value of '5' given as type '" + reflect.TypeOf(currentValue).String() + "' should be rounded to '5', but was rounded to '" + stringformat.GetIntAsString(result) +"' instead.")
		}
	}
	for _, currentValue := range floatValues {
		result := RoundToWholeNumber(currentValue)
		if !IsFloatEffectivelyEqual(result, 6) {
			test.Errorf("The value of '5.5' given as type '" + reflect.TypeOf(currentValue).String() + "' should be rounded to '6', but was rounded to '" + stringformat.GetIntAsString(result) +"' instead.")
		}
	}
}

func TestRoundToDecimal(test *testing.T) {
	var intValues = recast.GetArrayOfInterfaces(int(3), int8(3), int16(3), int32(3), int64(3))
	var uIntValues = recast.GetArrayOfInterfaces(uint(3), uint8(3), uint16(3), uint32(3), uint64(3))
	var floatValues = recast.GetArrayOfInterfaces(float32(5.1234567), float64(5.1234567))
	for _, currentValue := range intValues {
		result := RoundToDecimal(currentValue, 3)
		if !IsFloatEffectivelyEqual(result, 3) {
			test.Errorf("The value of '3' given as type '" + reflect.TypeOf(currentValue).String() + "' should be rounded to '3' decimal places, but returned '" + stringformat.GetIntAsString(result) +"' instead.")
		}
	}
	for _, currentValue := range uIntValues {
		result := RoundToDecimal(currentValue, 3)
		if !IsFloatEffectivelyEqual(result, 3) {
			test.Errorf("The value of '3' given as type '" + reflect.TypeOf(currentValue).String() + "' should be rounded to '3' decimal places, but returned '" + stringformat.GetIntAsString(result) +"' instead.")
		}
	}
	for _, currentValue := range floatValues {
		result := RoundToDecimal(currentValue, 4)
		if !IsFloatEffectivelyEqual(result, 5.1235) {
			test.Errorf("The value of '5.1234567' given as type '" + reflect.TypeOf(currentValue).String() + "' should be rounded to '4' decimal places, but returned '" + stringformat.GetIntAsString(result) +"' instead.")
		}
	}
}