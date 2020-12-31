package memory

import (
	"testing"
)

func TestCreateDeleteButton(test *testing.T) {
	InitializeButtonMemory()
	styleEntry := NewTuiStyleEntry()
	AddButton("MyLayer","MyButtonAlias", "ButtonLabel", styleEntry, 0, 0, 10, 11)
	if ButtonMemory["MyLayer"]["MyButtonAlias"] == nil {
		test.Errorf("A button was requested to be created, but could not be found in memory!")
	}
	DeleteButton("MyLayer", "MyButtonAlias")
	if ButtonMemory["MyLayer"]["MyButtonAlias"] != nil {
		test.Errorf("A button was requested to be delete, but it could still be found in memory!")
	}
}
