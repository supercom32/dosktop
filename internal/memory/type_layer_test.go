package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLayerTypeCreation(test *testing.T) {
	firstLayerEntry := NewLayerEntry(20, 20)
	firstLayerEntry.IsParent = false
	firstLayerEntry.LayerAlias = "MyAlias"
	firstLayerEntry.ParentAlias = "MyParentAlias"
	firstLayerEntry.ScreenXLocation = 1
	firstLayerEntry.ScreenYLocation = 2
	firstLayerEntry.CursorXLocation = 3
	firstLayerEntry.CursorYLocation = 4
	firstLayerEntry.ZOrder = 1
	firstLayerEntry.IsVisible = true
	secondLayerEntry := NewLayerEntry(20, 20)
	assert.NotEqualf(test, secondLayerEntry, firstLayerEntry,"The first layer entry is the same as the second, even though it should be different.")

	secondLayerEntry = NewLayerEntry(0,0, &firstLayerEntry)
	assert.Equalf(test, secondLayerEntry, firstLayerEntry, "The first layer is not the same as the second, even though it should be an identical clone.")
}
