package memory

import (
	"testing"
)

func TestCreateDeleteImage(test *testing.T) {
	InitializeImageMemory()
	imageEntry := NewImageEntry()
	AddImage("MyImageAlias", imageEntry)
	if ImageMemory["MyImageAlias"] == nil {
		test.Errorf("An image entry was requested to be created, but could not be found in memory!")
	}
	DeleteImage("MyImageAlias")
	if ImageMemory["MyImageAlias"] != nil {
		test.Errorf("An image was requested to be delete, but it could still be found in memory!")
	}
}
