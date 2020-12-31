package memory

import "fmt"

var TextStyleMemory map[string]*TextStyleEntryType

func InitializeTextStyleMemory() {
	TextStyleMemory = make(map[string]*TextStyleEntryType)
}

// AddTextStyle asdadsdas
func AddTextStyle(textStyleAlias string, attributeEntry TextStyleEntryType) {
	TextStyleMemory[textStyleAlias] = &attributeEntry
}

func GetTextStyle(textStyleAlias string) TextStyleEntryType {
	if TextStyleMemory[textStyleAlias] == nil {
		panic(fmt.Sprintf("The requested text style with alias '%s' could not be returned since it does not exist.", textStyleAlias))
	}
	return *TextStyleMemory[textStyleAlias]
}

func GetTextStyleAsAttributeEntry(textStyleAlias string) AttributeEntryType {
	textStyleEntry := TextStyleMemory[textStyleAlias]
	attributeEntry := NewAttributeEntry()
	attributeEntry.ForegroundColor = textStyleEntry.ForegroundColor
	attributeEntry.BackgroundColor = textStyleEntry.BackgroundColor
	attributeEntry.IsBlinking = textStyleEntry.IsBlinking
	attributeEntry.IsItalic = textStyleEntry.IsItalic
	attributeEntry.IsReversed = textStyleEntry.IsReversed
	attributeEntry.IsUnderlined = textStyleEntry.IsUnderlined
	attributeEntry.IsBold = textStyleEntry.IsBold
	return attributeEntry
}

// DeleteTextStyle asdasds
func DeleteTextStyle(textStyleAlias string) {
	delete(TextStyleMemory, textStyleAlias)
}
