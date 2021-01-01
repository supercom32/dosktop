package dosktop

import (
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerticalMenuPrompt(test *testing.T) {
	commonResource.isDebugEnabled = true
	layerAlias1 := "Layer1"
	layerWidth := 50
	layerHeight := 10
	styleEntry := memory.NewTuiStyleEntry()
	selectionEntry := memory.NewSelectionEntry()
	selectionEntry.AddSelection("MenuAlias1", "MenuValue1")
	selectionEntry.AddSelection("MenuAlias2", "MenuValue2")
	selectionEntry.AddSelection("MenuAlias3", "MenuValue3")
	selectionEntry.AddSelection("MenuAlias4", "MenuValue4")
	selectionEntry.AddSelection("MenuAlias5", "MenuValue5")
	selectionEntry.AddSelection("MenuAlias6", "MenuValue6")
	selectionEntry.AddSelection("MenuAlias7", "MenuValue7")
	selectionEntry.AddSelection("MenuAlias8", "MenuValue8")
	InitializeTerminal(layerWidth, layerHeight)
	AddLayer(layerAlias1, 0, 0, layerWidth, layerHeight, 1, "")
	Layer(layerAlias1)
	Color(4, 6)
	FillLayer(layerAlias1, "a1a2a3a4a5")
	layerEntry := memory.GetLayer(layerAlias1)
	for currentIndex := 0; currentIndex < 4; currentIndex++ {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("down")
	}
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("enter")
	selectionAlias := GetSelectionFromVerticalMenu(layerAlias1, styleEntry, selectionEntry,2,2,15, 3)
	expectedIndex := "MenuAlias5"
	obtainedIndex := selectionAlias
	assert.Equalf(test, expectedIndex, obtainedIndex, "The selected menu index returned is not correct!")
	styleEntry.MenuTextAlignment = constants.CenterAligned
	for currentIndex := 0; currentIndex < 8; currentIndex++ {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("down")
	}
	for currentIndex := 0; currentIndex < 7; currentIndex++ {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("up")
	}
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("enter")
	selectionAlias = GetSelectionFromVerticalMenu(layerAlias1, styleEntry, selectionEntry,20,2,20, 3)
	expectedIndex = "MenuAlias1"
	obtainedIndex = selectionAlias
	assert.Equalf(test, expectedIndex, obtainedIndex, "The selected menu index returned is not correct!")
	UpdateDisplay()
	obtainedValue := layerEntry.GetBasicAnsiStringAsBase64()
	expectedValue := "G1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG1hMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtChtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExG1szODsyOzI1NTsyNTU7MjU1bRtbNDg7MjswOzA7MG1NZW51VmFsdWUzICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MjU1OzI1NTsyNTVtICAgICBNZW51VmFsdWUxICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExG1szODsyOzI1NTsyNTU7MjU1bRtbNDg7MjswOzA7MG1NZW51VmFsdWU0ICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bTRhNRtbMzg7MjsyNTU7MjU1OzI1NW0bWzQ4OzI7MDswOzBtICAgICBNZW51VmFsdWUyICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExG1szODsyOzA7MDswbRtbNDg7MjsyNTU7MjU1OzI1NW1NZW51VmFsdWU1ICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bTRhNRtbMzg7MjsyNTU7MjU1OzI1NW0bWzQ4OzI7MDswOzBtICAgICBNZW51VmFsdWUzICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1G1szODsyOzA7MDswbRtbNDg7MjswOzA7MG0KG1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG1hMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtChtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1G1szODsyOzA7MDswbRtbNDg7MjswOzA7MG0KG1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG1hMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtCg=="
	DumpLayerToFile(*layerEntry)
	assert.Equalf(test, expectedValue, obtainedValue, "The updated screen does not match the master original!")
}

func TestHorizontalMenuDrawing(test *testing.T) {
	commonResource.isDebugEnabled = true
	layerAlias1 := "Layer1"
	layerWidth := 110
	layerHeight := 7
	styleEntry := memory.NewTuiStyleEntry()
	selectionEntry := memory.NewSelectionEntry()
	selectionEntry.AddSelection("MenuAlias1", "MenuValue1")
	selectionEntry.AddSelection("MenuAlias2", "MenuValueThatIsLong2")
	selectionEntry.AddSelection("MenuAlias3", "MenuValue3")
	selectionEntry.AddSelection("MenuAlias4", "MenuValueThatIsEvenLonger4")
	InitializeTerminal(layerWidth, layerHeight)
	AddLayer(layerAlias1, 0, 0, layerWidth, layerHeight, 1, "")
	Layer(layerAlias1)
	Color(4, 6)
	FillLayer(layerAlias1, "a1a2a3a4a5")
	layerEntry := memory.GetLayer(layerAlias1)
	for currentIndex := 0; currentIndex < 1; currentIndex++ {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("right")
	}
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("enter")
	selectionAlias := GetSelectionFromHorizontalMenu(layerAlias1, styleEntry, selectionEntry, 2, 1)
	expectedIndex := "MenuAlias2"
	obtainedIndex := selectionAlias
	assert.Equalf(test, expectedIndex, obtainedIndex, "The selected menu index returned is not correct!")
	for currentIndex := 0; currentIndex < 4; currentIndex++ {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("right")
	}
	for currentIndex := 0; currentIndex < 4; currentIndex++ {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("left")
	}
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("enter")
	styleEntry.MenuTextAlignment = constants.CenterAligned
	selectionAlias = GetSelectionFromProportionalHorizontalMenu(layerAlias1, styleEntry, selectionEntry, 2, 3, 100, 3, 0)
	expectedIndex = "MenuAlias1"
	obtainedIndex = selectionAlias
	assert.Equalf(test, expectedIndex, obtainedIndex, "The selected menu index returned is not correct!")
	UpdateDisplay()
	obtainedValue := layerEntry.GetBasicAnsiStringAsBase64()
	expectedValue := "G1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG1hMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtChtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTEbWzM4OzI7MjU1OzI1NTsyNTVtG1s0ODsyOzA7MDswbSBNZW51VmFsdWUxIBtbMzg7MjswOzA7MG0bWzQ4OzI7MjU1OzI1NTsyNTVtIE1lbnVWYWx1ZVRoYXRJc0xvbmcyIBtbMzg7MjsyNTU7MjU1OzI1NW0bWzQ4OzI7MDswOzBtIE1lbnVWYWx1ZTMgIE1lbnVWYWx1ZVRoYXRJc0V2ZW5Mb25nZXI0IBtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtChtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExG1szODsyOzA7MDswbRtbNDg7MjsyNTU7MjU1OzI1NW0gICAgICAgICAgIE1lbnVWYWx1ZTEgICAgICAgICAgICAbWzM4OzI7MjU1OzI1NTsyNTVtG1s0ODsyOzA7MDswbSAgICAgIE1lbnVWYWx1ZVRoYXRJc0xvbmcyICAgICAgICAgICAgICAgICAgIE1lbnVWYWx1ZTMgICAgICAgICAgICAbWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWEyYTNhNGE1G1szODsyOzA7MDswbRtbNDg7MjswOzA7MG0KG1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG1hMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtChtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1YTFhMmEzYTRhNWExYTJhM2E0YTVhMWEyYTNhNGE1G1szODsyOzA7MDswbRtbNDg7MjswOzA7MG0K"
	DumpLayerToFile(*layerEntry)
	assert.Equalf(test, expectedValue, obtainedValue, "The updated screen does not match the master original!")
}

func TestInput(test *testing.T) {
	commonResource.isDebugEnabled = true
	layerAlias1 := "Layer1"
	layerWidth := 20
	layerHeight := 3
	styleEntry := memory.NewTuiStyleEntry()
	InitializeTerminal(layerWidth, layerHeight)
	AddLayer(layerAlias1, 0, 0, layerWidth, layerHeight, 1, "")
	Layer(layerAlias1)
	Color(4, 6)
	FillLayer(layerAlias1, "a1a2a3a4a5")
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("t", "h", "i", "s", " ", "i", "s", " ", "a", " ", "t", "e", "s", "t", "!")
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("backspace", "backspace", "backspace", "backspace", "backspace")
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("e", "d", "i", "t", "!")
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("home", "end")
	memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("enter")
	GetInput(layerAlias1, styleEntry, 2, 1, 15, 20, false, "")
	UpdateDisplay()
	layerEntry := memory.GetLayer(layerAlias1)
	obtainedValue := layerEntry.GetBasicAnsiStringAsBase64()
	expectedValue := "G1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG1hMWEyYTNhNGE1YTFhMmEzYTRhNRtbMzg7MjswOzA7MG0bWzQ4OzI7MDswOzBtChtbMzg7MjswOzA7MTI4bRtbNDg7MjswOzEyODsxMjhtYTEbWzM4OzI7MjU1OzI1NTsyNTVtG1s0ODsyOzA7MDswbWhpcyBpcyBhIGVkaXQh4paIG1szODsyOzA7MDsxMjhtG1s0ODsyOzA7MTI4OzEyOG00YTUbWzM4OzI7MDswOzBtG1s0ODsyOzA7MDswbQobWzM4OzI7MDswOzEyOG0bWzQ4OzI7MDsxMjg7MTI4bWExYTJhM2E0YTVhMWEyYTNhNGE1G1szODsyOzA7MDswbRtbNDg7MjswOzA7MG0K"
	DumpLayerToFile(*layerEntry)
	assert.Equalf(test, expectedValue, obtainedValue, "The updated screen does not match the master original!")
}