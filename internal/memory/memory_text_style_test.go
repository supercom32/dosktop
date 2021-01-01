package memory

import (
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTextStyle(test *testing.T) {
	expectedAlias := "MyCustomAttribute"
	expectedForegroundColor := int32(constants.ColorRed)
	expectedBackgroundColor := int32(constants.ColorBrightGreen)
	expectedIsBlinking := true
	InitializeTextStyleMemory()
	attributeEntry := NewTextStyleEntry()
	attributeEntry.ForegroundColor = expectedForegroundColor
	attributeEntry.BackgroundColor = expectedBackgroundColor
	attributeEntry.IsBlinking = expectedIsBlinking
	AddTextStyle(expectedAlias, attributeEntry)
	expectedResult := recast.GetArrayOfInterfaces(expectedForegroundColor, expectedBackgroundColor, expectedIsBlinking)
	obtainedAttributeEntry := GetTextStyle(expectedAlias)
	obtainedResult := recast.GetArrayOfInterfaces(obtainedAttributeEntry.ForegroundColor, obtainedAttributeEntry.BackgroundColor, obtainedAttributeEntry.IsBlinking)
	assert.Equalf(test, expectedResult, obtainedResult, "The created dialog attribute style did not match what was supposed to be created!")
}

func TestDeleteTextStyle(test *testing.T) {
	expectedAlias := "MyCustomAttribute"
	expectedForegroundColor := int32(constants.ColorRed)
	expectedBackgroundColor := int32(constants.ColorBrightGreen)
	expectedIsBlinking := true
	InitializeTextStyleMemory()
	attributeEntry := NewTextStyleEntry()
	attributeEntry.ForegroundColor = expectedForegroundColor
	attributeEntry.BackgroundColor = expectedBackgroundColor
	attributeEntry.IsBlinking = expectedIsBlinking
	AddTextStyle(expectedAlias, attributeEntry)
	DeleteTextStyle(expectedAlias)
	assert.Panics(test, func() {GetTextStyle("expectedAlias")}, "The created dialog attribute style did not match what was supposed to be created!")
	//assert.Equalf(test, (*AttributeEntryType)(nil), TextStyleMemory[expectedAlias], "The created dialog attribute style did not match what was supposed to be created!")
}