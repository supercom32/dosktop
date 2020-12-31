package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetImageEntry(test *testing.T) {
	firstImageEntry := NewImageEntry()
	secondImageEntry := NewImageEntry()
	layerEntry := NewLayerEntry(20,20)
	layerEntry.IsParent = false
	layerEntry.LayerAlias = "MyAlias"
	layerEntry.ParentAlias = "MyParentAlias"
	layerEntry.ScreenXLocation = 1
	layerEntry.ScreenYLocation = 2
	layerEntry.CursorXLocation = 3
	layerEntry.CursorYLocation = 4
	layerEntry.ZOrder = 1
	layerEntry.IsVisible = true
	secondImageEntry.LayerEntry = layerEntry
	assert.NotEqualf(test, secondImageEntry, firstImageEntry,"The first image entry is the same as the second, even though it should be different.")

	firstImageEntry = NewImageEntry(&secondImageEntry)
	assert.Equalf(test, secondImageEntry, firstImageEntry, "The first image entry is not the same as the second, even though it should be an identical clone.")

}
