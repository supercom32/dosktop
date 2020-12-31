package ansi

import (
	"fmt"
	"image/color"
	"github.com/mbndr/figlet4go"
)

func GetColorComponentsFromAnsi256Color(colorIndex int) (int32, int32, int32) {
	ascii := figlet4go.NewAsciiRender()
	ascii.LoadFont("")
	return GetColorComponentsFromHexColor(ansi265Color[colorIndex])
}

func GetColorComponentsFromHexColor(s string) (int32, int32, int32) {
	var c color.RGBA
	c.A = 0xff
	switch len(s) {
	case 4:
		_, _ = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	case 7:
		_, _ = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	default:
		fmt.Sprint("Bad HEX VALUE")
	}
	redColorIndex := c.R
	greenColorIndex := c.G
	blueColorIndex := c.B
	return int32(redColorIndex), int32(greenColorIndex), int32(blueColorIndex)
}

/*
func LoadAnsiImage(fileName string) {
	var characterMemory [][]memory.CharacterEntryType
	var row []memory.CharacterEntryType
	data, _ := getFileFromFileSystem("typescript.ans")
	arrayOfRunes := stringformat.GetRunesFromString(data)
	attributeEntry := memory.NewAttributeEntry()
	for currentIndex := 0; currentIndex < len(arrayOfRunes); currentIndex++ {
		if arrayOfRunes[currentIndex] == '\n' {
			if len(row) < 190 {
				for i := 0; i < 190; i++ {
					characterEntry := memory.NewCharacterEntry()
					row = append(row, characterEntry)
				}
			}
			characterMemory = append(characterMemory, row)
			row = nil
		}
		ansiCodeLength := getAnsiCodeAsString(arrayOfRunes, currentIndex)
		if ansiCodeLength != 0 {
			fmt.Sprintf(string(arrayOfRunes[currentIndex:currentIndex+ansiCodeLength]))
			attributeEntry = getAnsiColor(arrayOfRunes[currentIndex:currentIndex+ansiCodeLength], attributeEntry)
			currentIndex += ansiCodeLength
		} else {
			characterEntry := memory.NewCharacterEntry()
			characterEntry.Character = arrayOfRunes[currentIndex]
			characterEntry.AttributeEntry = memory.NewAttributeEntry(&attributeEntry)
			row = append(row, characterEntry)
		}
	}
	layerEntry := memory.NewLayerEntry(len(characterMemory[0]), len(characterMemory))
	layerEntry.CharacterMemory = characterMemory
	DumpLayerToFile(layerEntry)
	fmt.Sprint("")
}



func getAnsiColor(ansiCode []rune, attributeEntry memory.AttributeEntryType) memory.AttributeEntryType {
	if ansiCode[0] ==  constant.AnsiEsc && ansiCode[1] == '[' && ansiCode[len(ansiCode)-1] == 'm' {
		delimitedAnsiString := strings.Split(string(ansiCode[2:len(ansiCode) - 1]), ";")
		for currentIndex := 0; currentIndex < len(delimitedAnsiString); currentIndex++ {
			if delimitedAnsiString[currentIndex] == "38" && delimitedAnsiString[currentIndex+1] == "5" {
				redColorIndex, greenColorIndex, blueColorIndex := ansi.GetColorComponentsFromAnsi256Color(recast.GetStringAsInt(delimitedAnsiString[currentIndex+2]))
				attributeEntry.ForegroundColor = GetRGBColor(redColorIndex, greenColorIndex, blueColorIndex)
				currentIndex += 3
			}
			if delimitedAnsiString[currentIndex] == "48" && delimitedAnsiString[currentIndex+1] == "5" {
				redColorIndex, greenColorIndex, blueColorIndex := ansi.GetColorComponentsFromAnsi256Color(recast.GetStringAsInt(delimitedAnsiString[currentIndex+2]))
				attributeEntry.BackgroundColor = GetRGBColor(redColorIndex, greenColorIndex, blueColorIndex)
				currentIndex += 3
			}
		}
		fmt.Sprint(delimitedAnsiString)
	}
	return attributeEntry
}

func getAnsiCodeAsString(arrayOfRunes []rune, startingIndex int) int {
	lengthOfCode := 0
	if arrayOfRunes[startingIndex] == constant.AnsiEsc && arrayOfRunes[startingIndex+1] == '[' {
		lengthOfCode += 2
		for currentIndex := 2; currentIndex < len(arrayOfRunes); currentIndex++ {
			if isRuneValidAnsi(arrayOfRunes[startingIndex+currentIndex]) {
				lengthOfCode++
			} else {
				return lengthOfCode
			}
		}
	}
	return lengthOfCode
}

func isRuneValidAnsi(character rune) bool {
	isValid := false
	if character == '[' ||
		unicode.IsDigit(character) ||
		character == ';' ||
		character == 'm' {
		isValid = true
	}
	return isValid
}
*/

/*
package main

import (
"github.com/mbndr/figlet4go"
"fmt"
)

func main() {
ascii := figlet4go.NewAsciiRender()
options := figlet4go.NewRenderOptions()
options.FontName = "smbraille"

ascii.LoadFont("/path/to/fonts/")

renderStr, _ := ascii.RenderOpts("Hello Fonts", options)
fmt.Print(renderStr)

}
 */