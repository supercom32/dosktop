package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterTypeCreation(test *testing.T) {
	firstCharacterObject := NewCharacterEntry()
	firstColorObject := NewAttributeEntry()
	firstColorObject.ForegroundColor = 1234
	firstColorObject.BackgroundColor = 5678
	firstColorObject.IsBlinking = false
	firstColorObject.IsBold = true
	firstColorObject.IsReversed = false
	firstColorObject.IsUnderlined = true
	firstCharacterObject.Character = rune('A')
	firstCharacterObject.AttributeEntry = firstColorObject
	secondCharacterObject := NewCharacterEntry(&firstCharacterObject)
	assert.Equalf(test, secondCharacterObject, firstCharacterObject, "The second Character object should be the same as the first, as it was created as a copy")

	secondCharacterObject.Character = rune('Z')
	secondCharacterObject.AttributeEntry.ForegroundColor = 1234
	secondCharacterObject.AttributeEntry.BackgroundColor = 5678
	secondCharacterObject.AttributeEntry.IsBlinking = false
	secondCharacterObject.AttributeEntry.IsBold = false
	secondCharacterObject.AttributeEntry.IsReversed = true
	secondCharacterObject.AttributeEntry.IsUnderlined = false
	assert.NotEqualf(test, secondCharacterObject, firstCharacterObject,"The second Character object should not be the same as the first, as manipulating it should only effect itself.")
}
