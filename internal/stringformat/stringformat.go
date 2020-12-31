package stringformat

import (
	"github.com/supercom32/dosktop/internal/recast"
	"encoding/base64"
	"fmt"
	"golang.org/x/text/width"
)

const maxLen = 4096
const nullRune = '\x00'
const leftAligned = 0
const rightAligned = 1
const centerAligned = 2

func IsRuneCharacterWide(character rune) bool {
	isCharacterWide := false
	properties := width.LookupRune(character)
	if properties.Kind() == width.EastAsianWide || properties.Kind() == width.EastAsianFullwidth {
		isCharacterWide = true
	}
	return isCharacterWide
}

func GetRunesFromString(stringToConvert string) []rune {
	//narrow := width.Narrow.String(stringToConvert)
	runes := []rune(stringToConvert)
	if len(runes) == 0 {
		runes = append(runes, nullRune)
	}
	return runes
}

func GetIntAsString(number interface{}) string {
	numberAsFloatint64 := recast.GetNumberAsInt64(number)
	return fmt.Sprintf("%d", numberAsFloatint64)
}
func GetFloatAsString(number interface{}) string {
	numberAsFloat64 := recast.GetNumberAsFloat64(number)
	return fmt.Sprintf("%g", numberAsFloat64)
}


func GetSubString(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}

func GetStringAsBase64(inputString string) string {
	base64String := base64.StdEncoding.EncodeToString([]byte(inputString))
	return base64String
}

func GetStringFromBase64(inputString string) string {
	decodedString, err := base64.StdEncoding.DecodeString(inputString)
	if err != nil {
		panic(err)
	}
	return string(decodedString)
}

func GetNumberOfWideCharacters(arrayOfRunes []rune) int {
	numberOfWideCharacters := 0
	for _, currentRune := range arrayOfRunes {
		if IsRuneCharacterWide((currentRune)) {
			numberOfWideCharacters++
		}
	}
	return numberOfWideCharacters
}

func GetFormattedString(stringToFormat string, lengthOfString int, position int) string {
	arrayOfRunes := GetRunesFromString(stringToFormat)
	formattedArrayOfRunes := []rune{}
	paddingSize := lengthOfString - len(arrayOfRunes)
	if paddingSize < 0 {
		paddingSize = 0
	}
	// Since wide runes take up two spaces, we need to subtract that amount of space from our padding so that
	// when everything is drawn, we don't draw more spaces than required (since printing automatically advances wide
	// characters by one space).
	paddingSize = paddingSize - (GetNumberOfWideCharacters(arrayOfRunes))
	if paddingSize < 0 {
		paddingSize = 0
	}
	fullStringPadding := GetFilledString(paddingSize, " ")
	halfStringPadding := GetFilledString(paddingSize/2, " ")
	if position == rightAligned {
		formattedArrayOfRunes = append(GetRunesFromString(fullStringPadding), arrayOfRunes...)
	} else if position == centerAligned {
		formattedArrayOfRunes = append(GetRunesFromString(halfStringPadding), arrayOfRunes...)
		formattedArrayOfRunes = append(formattedArrayOfRunes, GetRunesFromString(halfStringPadding)...)
		if len(formattedArrayOfRunes) < lengthOfString {
			formattedArrayOfRunes = append(formattedArrayOfRunes, ' ')
		}
	} else {
		formattedArrayOfRunes = append(arrayOfRunes, GetRunesFromString(fullStringPadding)...)
	}
	return string(formattedArrayOfRunes)
}

func GetFilledString(lengthOfString int, character string) string {
	newString := ""
	for currentIndex := 0; currentIndex < lengthOfString; currentIndex++ {
		newString += character
	}
	return newString
}