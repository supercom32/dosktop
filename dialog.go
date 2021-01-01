package dosktop

import (
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/supercom32/dosktop/internal/stringformat"
	"fmt"
)

/*
PrintDialog allows you to write text immediately to the terminal screen via a
typewriter effect. This is useful for video games or other applications that
may require printing text in a dialog box. In addition, the following
information should be noted:

- If you specify a print location outside the range of your specified text
layer, a panic will be generated to fail as fast as possible.

- If printing has reached the last line of your text layer, printing will
not advance to the next line. Instead, it will resume and overwrite
what was already printed on the same line.

- Specifying the width of your text line allows you to control when text
wrapping occurs. For example, if printing starts at location (2, 2) and you set
a line width of 10 characters, text wrapping will occur when the printing
exceeds the text layer location (12, 2). When this happens, text will continue
to print underneath the previous line at (2, 3).

- When a word is too long to be printed on a text layer line, or the width
of your line has already exceed its allowed maximum, the word will be pushed
to the line directly under it. This prevents words from being split across
two lines.

- When specifying a printing delay, the amount of time to wait is inserted
between each character printed and does not reflect the overall time to
print your specified text.

- If the dialog being printed is flagged as skipable, the user can speed up
printing by pressing the 'enter' key or right mouse button. Otherwise, they
must wait for the animation to completely finish before execution continues.

- This method supports the use of text styles during printing to add color
or styles to specific words in your string. All text styles must be enclosed
around the "{" and "}" characters. If you wish to use the default text
style, simply omit specifying any text style between your enclosing braces.
For example:

	// Add a text layer with the alias "ForegroundLayer", at location (0, 0),
	// with a width and height of 80x20 characters, z order priority of 1,
	// with no parent layer.
	dosktop.AddLayer("ForegroundLayer", 0, 0, 80, 20, 1, "")
	// Obtain a new text style entry.
	redTextStyle := dosktop.GetTextStyle()
	// Change the default foreground color of our text style to be red.
	redTextStyle.ForegroundColor = dosktop.GetRGBColor(255,0,0)
	// Register our new text style so Dosktop can use it.
	dosktop.AddTextStyle("red", redTextStyle)
	// Print some dialog text on the text layer "ForegroundLayer", at location
	// (0, 0), with a text wrapping location at 30 characters, a 10 millisecond
	// delay between each character printed, and mark the dialog as skipable.
	// Inside our string to print, we add the "{red}" tag to switch printing
	// styles on the fly to "red" and change back to the default style using
	// "{}".
	dosktop.PrintDialog("ForegroundLayer", 0, 0, 30, 10, true, "This is some dialog text in {red}red color{}. Only the words 'red color' should be colored.")
*/
func PrintDialog(layerAlias string, xLocation int, yLocation int, widthOfLineInCharacters int, printDelayInMilliseconds int, isSkipable bool, stringToPrint string) {
	layerEntry := memory.GetLayer(layerAlias)
	if xLocation < 0 || xLocation > layerEntry.Width || yLocation < 0 || yLocation > layerEntry.Height {
		panic(fmt.Sprintf("The specified location (%d, %d) is out of bounds for layer '%s' with a size of (%d, %d).", xLocation, yLocation, layerAlias, layerEntry.Width, layerEntry.Height))
	}
	printDialog(layerEntry, layerEntry.DefaultAttribute, xLocation, yLocation, widthOfLineInCharacters, printDelayInMilliseconds, isSkipable, stringToPrint)
}

/*
printDialog allows you to write text to the terminal screen via a typewriter
effect. This is useful for video games or other applications that may
require printing text in a dialog box.
*/
func printDialog(layerEntry *memory.LayerEntryType, attributeEntry memory.AttributeEntryType, xLocation int, yLocation int, widthOfLineInCharacters int, printDelayInMilliseconds int, isSkipable bool, textToPrint string) {
	if xLocation < 0 || xLocation > layerEntry.Width || yLocation < 0 || yLocation > layerEntry.Height {
		panic(fmt.Sprintf("The specified location (%d, %d) is out of bounds for the layer with a size of (%d, %d).", xLocation, yLocation, layerEntry.Width, layerEntry.Height))
	}
	var isPrintDelaySkipped bool
	arrayOfRunes := stringformat.GetRunesFromString(textToPrint)
	layerWidth := layerEntry.Width
	layerHeight := layerEntry.Height
	cursorXLocation := xLocation
	cursorYLocation := yLocation
	currentAttributeEntry := attributeEntry
	for currentCharacterIndex := 0; currentCharacterIndex < len(arrayOfRunes); currentCharacterIndex++ {
		currentCharacter := stringformat.GetSubString(textToPrint, currentCharacterIndex, 1)
		printLayer(layerEntry, currentAttributeEntry, cursorXLocation, cursorYLocation, []rune{arrayOfRunes[currentCharacterIndex]})
		cursorXLocation++
		lengthOfNextWord := 0
		if currentCharacter == " " {
			lengthOfNextWord = getLengthOfNextWord(textToPrint, currentCharacterIndex+1)
		}
		nextCharacter := stringformat.GetSubString(textToPrint, currentCharacterIndex+1, 1)
		if nextCharacter == "{" {
			attributeTag := getAttributeTag(textToPrint, currentCharacterIndex+1)
			currentAttributeEntry = getDialogAttributeEntry(attributeTag, attributeEntry)
			currentCharacterIndex += len(attributeTag)
		}
		if cursorXLocation + lengthOfNextWord - xLocation >= widthOfLineInCharacters || cursorXLocation + lengthOfNextWord >= layerWidth {
			cursorXLocation = xLocation
			cursorYLocation++
			if cursorYLocation >= layerHeight {
				cursorYLocation--
			}
		}
		if isSkipable == true {
			_,_, mouseButtonPressed, _ := memory.MouseMemory.GetMouseStatus()
			keyPressed := Inkey()
			if mouseButtonPressed != 0 || keyPressed == "enter" {
				isPrintDelaySkipped = true
			}
		}
		if isPrintDelaySkipped == false {
			SleepInMilliseconds(uint(printDelayInMilliseconds))
			UpdateDisplay()
		}

	}
	UpdateDisplay()
}

/*
getAttributeTag allows you to obtain an attribute tag from a given text string.
Attributes are always surrounded by "{" and "}" characters.  In addition, the
following information should be noted:

- If no attribute tag could be detected at the given string location, then
an empty string will be returned instead.
*/
func getAttributeTag(stringToParse string, startingCharacterIndex int) string {
	var lengthOfAttributeTag int
	for currentCharacterIndex := startingCharacterIndex; currentCharacterIndex < len(stringToParse); currentCharacterIndex++ {
		lengthOfAttributeTag++
		if stringformat.GetSubString(stringToParse, currentCharacterIndex, 1) == "}" {
			return stringformat.GetSubString(stringToParse, startingCharacterIndex, lengthOfAttributeTag)
		}
	}
	return ""
}

/*
getDialogAttributeEntry allows you to obtain an attribute entry based on the
text style detected in your attribute tag. If no text style could be found
that matches the attribute tag provided, then the default attribute entry
will be returned instead.
*/
func getDialogAttributeEntry(attributeTag string, defaultAttributeEntry memory.AttributeEntryType) memory.AttributeEntryType {
	var attributeEntry memory.AttributeEntryType
	if attributeTag != "" {
		textStyleAlias := stringformat.GetSubString(attributeTag, 1, len(attributeTag)-2)
		if memory.TextStyleMemory[textStyleAlias] != nil {
			attributeEntry = memory.GetTextStyleAsAttributeEntry(textStyleAlias)
			return attributeEntry
		}
	}
	return defaultAttributeEntry
}

/*
getLengthOfNextWord allows you to get the length of the next word at a given
position of a text string.
*/
func getLengthOfNextWord(stringToParse string, startingCharacterIndex int) int {
	var lengthOfNextWord int
	for currentCharacterIndex := startingCharacterIndex; currentCharacterIndex < len(stringToParse); currentCharacterIndex++ {
		if stringformat.GetSubString(stringToParse, currentCharacterIndex, 1) == " " {
			return lengthOfNextWord
		}
		lengthOfNextWord++
	}
	return lengthOfNextWord
}
