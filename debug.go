package dosktop

import (
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/math"
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/supercom32/dosktop/internal/recast"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

// PrintDebugLog(sourceLayerEntry.LayerAlias + " " + strconv.FormatFloat(float64(sourceAttributeEntry.foregroundTransformValue), 'f', 6, 32))

// Debug SDSDDS
func Debug(textToPrint string) {
	printDebugLog(textToPrint)
}

// PrintDebugLog adsada
func PrintDebugLog(textToPrint string) {
	if commonResource.isDebugEnabled {
		printDebugLog(textToPrint)
	}
}

func getTimeAndDate() string {
	currentTime := time.Now()
	return currentTime.String()
}

func DumpScreenToFile() {
	DumpLayerToFile(commonResource.screenLayer)
}

func DumpLayerToFile(layerEntry memory.LayerEntryType) {
	writeStringToFile("/home/administrator/Documents/Workspaces/golang/src/screenDump.ans", layerEntry.GetBasicAnsiString())
	writeStringToFile("/home/administrator/Documents/Workspaces/golang/src/screenDump.b64", layerEntry.GetBasicAnsiStringAsBase64())
}

func writeStringToFile(fileName string, stringToWrite string) {
	f, err := os.OpenFile(fileName,
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	err = f.Truncate(0)
	defer f.Close()
	if _, err := f.WriteString(stringToWrite); err != nil {
		log.Println(err)
	}
}
func printDebugLog(textToPrint string) {
	f, err := os.OpenFile("/home/administrator/Documents/Workspaces/golang/src/debug.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(textToPrint + "\n"); err != nil {
		log.Println(err)
	}
}
func printLog(fileName string, textToPrint string) {
	f, err := os.OpenFile("/home/administrator/Documents/Workspaces/golang/src/"+fileName,
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(textToPrint); err != nil {
		log.Println(err)
	}
}

func Test() {
	// base64String := getBase64PngFromImage(memory.GetImage("street").ImageData)
	// PrintDebugLog("IMG: " + base64String)
}

func expectEquals(obtainedValues []interface{}, expectedValues []interface{}, errorMessage string) string {
	return expect(obtainedValues, expectedValues, true, errorMessage)
}

func expectNotEquals(obtainedValues []interface{}, expectedValues []interface{}, errorMessage string) string {
	return expect(obtainedValues, expectedValues, false, errorMessage)
}

func expect(obtainedValues []interface{}, expectedValues []interface{}, isEqualTo bool, errorMessage string) string {
	var returnedError string
	if len(expectedValues) != len(obtainedValues) {
		return fmt.Sprintf("The number of obtained values (%d) did not match the number of expected values (%d).", len(obtainedValues), len(expectedValues))
	}
	for currentIndex :=0; currentIndex < len(expectedValues); currentIndex++ {
		obtainedValueDataType := recast.GetDataType(obtainedValues[currentIndex])
		expectedValueDataType := recast.GetDataType(expectedValues[currentIndex])
		if obtainedValueDataType != expectedValueDataType || obtainedValueDataType == constants.NullDataType || expectedValueDataType == constants.NullDataType {
			return fmt.Sprintf("The obtained value type '" + reflect.TypeOf(obtainedValues[currentIndex]).String() + "' does not match the expected type of '" + reflect.TypeOf(expectedValues[currentIndex]).String() + "'.")
		}
		if isEqualTo {
			if obtainedValueDataType == recast.StringType {
				obtainedValue := obtainedValues[currentIndex].(string)
				expectedValue := expectedValues[currentIndex].(string)
				if obtainedValue != expectedValue {
					returnedError += errorMessage + "\n"
					returnedError += fmt.Sprintf("Obtained value: '%s'\n", obtainedValue)
					returnedError += fmt.Sprintf("Expected value: '%s'\n", expectedValue)
					return returnedError
				}
			} else {
				obtainedValue := recast.GetNumberAsFloat64(obtainedValues[currentIndex])
				expectedValue := recast.GetNumberAsFloat64(expectedValues[currentIndex])
				if !math.IsFloatEffectivelyEqual(obtainedValue, expectedValue) {
					returnedError += errorMessage + "\n"
					returnedError += fmt.Sprintf("Obtained value: '%g'\n", obtainedValue)
					returnedError += fmt.Sprintf("Expected value: '%g'\n", expectedValue)
					return returnedError
				}
			}
		} else {
			if obtainedValueDataType == recast.StringType {
				obtainedValue := obtainedValues[currentIndex].(string)
				notExpectedValue := expectedValues[currentIndex].(string)
				if obtainedValue == notExpectedValue {
					returnedError += errorMessage + "\n"
					returnedError += fmt.Sprintf("Obtained value: '%s'\n", obtainedValue)
					returnedError += fmt.Sprintf("Expected value: '%s'\n", notExpectedValue)
					return returnedError
				}
			} else {
				obtainedValue := recast.GetNumberAsFloat64(obtainedValues[currentIndex])
				notExpectedValue := recast.GetNumberAsFloat64(expectedValues[currentIndex])
				if math.IsFloatEffectivelyEqual(obtainedValue, notExpectedValue){
					returnedError += errorMessage + "\n"
					returnedError += fmt.Sprintf("Obtained value: '%g'\n", obtainedValue)
					returnedError += fmt.Sprintf("Expected value: '%g'\n", notExpectedValue)
					return returnedError
				}
			}
		}
	}
	return ""
}