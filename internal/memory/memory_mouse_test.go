package memory

import "testing"

func TestMouseMemory(test *testing.T) {
	MouseMemory.SetMouseStatus(1,2,3,"down")
	xLocation, yLocation, buttonPressed, wheelState := MouseMemory.GetMouseStatus()
	if (xLocation != 1 || yLocation != 2 || buttonPressed != 3 || wheelState != "down") {
		test.Errorf("Mouse state was saved, but not read back with the same values.")
	}
}

func TestMouseMemoryClear(test *testing.T) {
	MouseMemory.SetMouseStatus(1,2,3,"down")
	MouseMemory.ClearMouseMemory()
	xLocation, yLocation, buttonPressed, wheelState := MouseMemory.GetMouseStatus()
	if (xLocation != -1 || yLocation != -1 || buttonPressed != 0 || wheelState != "") {
		test.Errorf("Mouse state was saved, but not read back with the same values.")
	}
}

func TestIsMouseInBoundingBox(test *testing.T) {
	MouseMemory.SetMouseStatus(2,2,3,"down")
	if MouseMemory.IsMouseInBoundingBox(1, 1, 10, 10) != true {
		test.Errorf("Mouse was in the bounding box, but was not detected as such.")
	}
	MouseMemory.SetMouseStatus(0,0,3,"down")
	if MouseMemory.IsMouseInBoundingBox(1, 1, 10, 10) != false {
		test.Errorf("Mouse was out of the bounding box, but was not detected as such.")
	}
}