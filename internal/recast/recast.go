package recast

import (
	"strconv"
)

const (
	StringType = iota
	IntType
	Int8Type
	Int16Type
	Int32Type
	Int64Type
	UIntType
	UInt8Type
	UInt16Type
	UInt32Type
	UInt64Type
	Float32Type
	Float64Type
)

func GetDataType(variable interface{}) int {
	switch variable.(type) {
	case string:
		return StringType
	case int:
		return IntType
	case int8:
		return Int8Type
	case int16:
		return Int16Type
	case int32:
		return Int32Type
	case int64:
		return Int64Type
	case uint:
		return UIntType
	case uint8:
		return UInt8Type
	case uint16:
		return UInt16Type
	case uint32:
		return UInt32Type
	case uint64:
		return UInt64Type
	case float32:
		return Float32Type
	case float64:
		return Float64Type
	}
	return -1
}

func GetNumberAsInt(number interface{}) int {
	detectedType := GetDataType(number)
	if detectedType == Float32Type {
		return int(number.(float32))
	}
	if detectedType == Float64Type {
		return int(number.(float64))
	}
	if detectedType == IntType {
		return number.(int)
	}
	if detectedType == Int8Type {
		return int(number.(int8))
	}
	if detectedType == Int16Type {
		return int(number.(int16))
	}
	if detectedType == Int32Type {
		return int(number.(int32))
	}
	if detectedType == Int64Type {
		return int(number.(int64))
	}
	if detectedType == UIntType {
		return int(number.(uint))
	}
	if detectedType == UInt8Type {
		return int(number.(uint8))
	}
	if detectedType == UInt16Type {
		return int(number.(uint16))
	}
	if detectedType == UInt32Type {
		return int(number.(uint32))
	}
	if detectedType == UInt64Type {
		return int(number.(uint64))
	}
	return -1
}

func GetNumberAsInt64(number interface{}) int64 {
	detectedType := GetDataType(number)
	if detectedType == Float32Type {
		return int64(number.(float32))
	}
	if detectedType == Float64Type {
		return int64(number.(float64))
	}
	if detectedType == IntType {
		return int64(number.(int))
	}
	if detectedType == Int8Type {
		return int64(number.(int8))
	}
	if detectedType == Int16Type {
		return int64(number.(int16))
	}
	if detectedType == Int32Type {
		return int64(number.(int32))
	}
	if detectedType == Int64Type {
		return number.(int64)
	}
	if detectedType == UIntType {
		return int64(number.(uint))
	}
	if detectedType == UInt8Type {
		return int64(number.(uint8))
	}
	if detectedType == UInt16Type {
		return int64(number.(uint16))
	}
	if detectedType == UInt32Type {
		return int64(number.(uint32))
	}
	if detectedType == UInt64Type {
		return int64(number.(uint64))
	}
	return -1
}

func GetStringAsInt(stringToConvert string) int {
	return int(GetStringAsInt64(stringToConvert))
}

func GetStringAsInt64(stringToConvert string) int64 {
	number, _ := strconv.Atoi(stringToConvert)
	return int64(number)
}

func GetNumberAsFloat64(number interface{}) float64 {
	detectedType := GetDataType(number)
	if detectedType == Float32Type {
		return float64(number.(float32))
	}
	if detectedType == Float64Type {
		return number.(float64)
	}
	if detectedType == IntType {
		return float64(number.(int))
	}
	if detectedType == Int8Type {
		return float64(number.(int8))
	}
	if detectedType == Int16Type {
		return float64(number.(int16))
	}
	if detectedType == Int32Type {
		return float64(number.(int32))
	}
	if detectedType == Int64Type {
		return float64(number.(int64))
	}
	if detectedType == UIntType {
		return float64(number.(uint))
	}
	if detectedType == UInt8Type {
		return float64(number.(uint8))
	}
	if detectedType == UInt16Type {
		return float64(number.(uint16))
	}
	if detectedType == UInt32Type {
		return float64(number.(uint32))
	}
	if detectedType == UInt64Type {
		return float64(number.(uint64))
	}
	return -1
}

func GetArrayOfInterfaces(variables ...interface{}) []interface{} {
	var arrayOfVariables []interface{}
	for _, currentVariable := range variables {
		arrayOfVariables = append(arrayOfVariables, currentVariable)
	}
	return arrayOfVariables
}