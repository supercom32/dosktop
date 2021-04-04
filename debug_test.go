package dosktop

import (
	"github.com/stretchr/testify/assert"
	"github.com/supercom32/filesystem"
	"testing"
)

func TestPrintDebugLog(test *testing.T) {
	fileName := "debug.log"
	expectedValue := "This is a test of the debug logger."
	filesystem.DeleteFile(fileName)
	printDebugLog(fileName, expectedValue)
	fileContentsAsBytes, err := filesystem.GetFileContentsAsBytes(fileName)
	assert.NoErrorf(test, err, "Failed to read the file '%s': ", fileName)
	obtainedValue := string(fileContentsAsBytes)
	assert.Equalf(test, expectedValue + "\n", obtainedValue, "The written log message does not match the original!")
	err = filesystem.DeleteFile(fileName)
	assert.NoErrorf(test, err, "Failed to delete the file '%s': ", fileName)
}

func TestDumpScreen(test *testing.T) {
	commonResource.isDebugEnabled = true
	layerAlias1 := "Layer1"
	layerWidth := 100
	layerHeight := 35
	InitializeTerminal(layerWidth, layerHeight)
	AddLayer(layerAlias1, 0, 0, layerWidth, layerHeight, 1, "")
	Layer(layerAlias1)
	LoadBase64Image(sampleBase64Image, "sampleImage")
	DrawImageToLayer(layerAlias1, "sampleImage", 1, 1, 100, 35, 0)
	UpdateDisplay()
	dumpScreenToFile()
	assert.FileExistsf(test, commonResource.debugDirectory + "screenDump.b64", "The base64 screen dump image could not be found!")
	assert.FileExistsf(test, commonResource.debugDirectory + "screenDump.ans", "The ansi screen dump image could not be found!")
}