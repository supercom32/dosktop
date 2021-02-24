package dosktop

import (
	"github.com/supercom32/dosktop/internal/memory"
	"log"
	"os"
	"github.com/supercom32/dosktop/internal/filesystem"
)

/*
DumpScreenToFile allows you to dump the current visible screen layer to files
on disk. In addition, the following information should be noted:

- This is a method just for internal testing and normally is not used unless
required for debugging.
*/
func dumpScreenToFile() {
	dumpLayerToFile(commonResource.screenLayer)
}

/*
DumpScreenToTerminal allows you to dump the current visible screen layer to the
terminal. In addition, the following information should be noted:

- This is a method just for internal testing and normally is not used unless
required for debugging.
*/
func dumpScreenToTerminal() {
	log.Println(commonResource.screenLayer.GetBasicAnsiStringAsBase64())
}

/*
DumpLayerToFile allows you to dump a display layer to files on a disk. One
file is a limited ansi representation of the screen layer. The other is
a base64 string representation of the same layer. In addition, the
following information should be noted:

- This is a method just for internal testing and normally is not used unless
required for debugging.
*/
func dumpLayerToFile(layerEntry memory.LayerEntryType) {
	filesystem.WriteBytesToFile(commonResource.debugDirectory + "screenDump.ans", []byte(layerEntry.GetBasicAnsiString()), 0644)
	filesystem.WriteBytesToFile(commonResource.debugDirectory + "screenDump.b64", []byte(layerEntry.GetBasicAnsiStringAsBase64()),0644)
}

/*
printDebugLog allows you to write debug logs to a file. In addition, the
following information should be noted:

- This is a method just for internal testing and normally is not used
unless required for debugging.
*/
func printDebugLog(fileName string, textToPrint string) {
	f, err := os.OpenFile(commonResource.debugDirectory + fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(textToPrint + "\n"); err != nil {
		log.Println(err)
	}
}
