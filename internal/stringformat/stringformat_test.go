package stringformat

import (
	"github.com/supercom32/dosktop/constant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsRuneCharacterWide(test *testing.T) {
	wideChineseCharacter := '读'
	wideJapaneseCharacter := 'ひ'
	wideKoreanCharacter := '밥'
	englishCharacter := 'W'
	obtainedResult := IsRuneCharacterWide(wideChineseCharacter)
	assert.Equalf(test, true, obtainedResult, "The Chinese character specified is wide, but was not detected as such.")
	obtainedResult = IsRuneCharacterWide(wideJapaneseCharacter)
	assert.Equalf(test, true, obtainedResult, "The Japanese character specified is wide, but was not detected as such.")
	obtainedResult = IsRuneCharacterWide(wideKoreanCharacter)
	assert.Equalf(test, true, obtainedResult, "The Korean character specified is wide, but was not detected as such.")
	obtainedResult = IsRuneCharacterWide(englishCharacter)
	assert.Equalf(test, false, obtainedResult, "The English character specified is not wide, but was not detected as such.")
}

func TestGetRunesFromString(test *testing.T) {
	arrayOfRunes := GetRunesFromString("This is a test string to be converted into a rune array!")
	obtainedResult := len(arrayOfRunes)
	expectedResult := 56
	assert.Equalf(test, expectedResult, obtainedResult, "The string specified did not return a rune array of proper length!")
}

func TestGetIntAsString(test *testing.T) {
	obtainedResult := GetIntAsString(123.456)
	expectedResult := "123"
	assert.Equalf(test, expectedResult, obtainedResult, "The number specified was not converted to a string correctly!")
}

func TestGetFloatAsString(test *testing.T) {
	obtainedResult := GetFloatAsString(123.456)
	expectedResult := "123.456"
	assert.Equalf(test, expectedResult, obtainedResult, "The number specified was not converted to a string correctly!")
}

func TestGetSubString(test *testing.T) {
	obtainedResult := GetSubString("This is a long string", 5, 2)
	expectedResult := "is"
	assert.Equalf(test, expectedResult, obtainedResult, "The substring requested was not correct!")
}

func TestGetStringAsBase64(test *testing.T) {
	obtainedResult := GetStringAsBase64("This is base64 encoded string")
	expectedResult := "VGhpcyBpcyBiYXNlNjQgZW5jb2RlZCBzdHJpbmc="
	assert.Equalf(test, expectedResult, obtainedResult, "The base64 encoded string requested is incorrect!")
}

func TestGetStringFromBase64(test *testing.T) {
	obtainedResult := GetStringFromBase64("VGhpcyBpcyBiYXNlNjQgZW5jb2RlZCBzdHJpbmc=")
	expectedResult := "This is base64 encoded string"
	assert.Equalf(test, expectedResult, obtainedResult, "The converted base64 string did not return the result expected!")
}

func TestGetNumberOfWideCharacters(test *testing.T) {
	arrayOfRunes := GetRunesFromString("AL ⌘读写汉字 ひらがな コンピュータワンワンローソク 보리밥보리밥 ⌘ EX")
	obtainedResult := GetNumberOfWideCharacters(arrayOfRunes)
	expectedResult := 28
	assert.Equalf(test, expectedResult, obtainedResult, "The number of wide characters detected did not match what was expected!")
}

func TestGetFormattedString(test *testing.T) {
	obtainedResult := GetFormattedString("Formatted String", 40, constants.LeftAligned)
	expectedResult := "Formatted String                        "
	assert.Equalf(test, expectedResult, obtainedResult, "The formatted string obtained was not left aligned as expected.")
	obtainedResult = GetFormattedString("Formatted String", 40, constants.RightAligned)
	expectedResult = "                        Formatted String"
	assert.Equalf(test, expectedResult, obtainedResult, "The formatted string obtained was not right aligned as expected.")
	obtainedResult = GetFormattedString("Formatted String", 40, constants.CenterAligned)
	expectedResult = "            Formatted String            "
	assert.Equalf(test, expectedResult, obtainedResult, "The formatted string obtained was not center aligned as expected.")
}