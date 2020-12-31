package memory

import (
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSelectionEntry(test *testing.T) {
	selectionEntry := NewSelectionEntry()
	selectionEntry.AddSelection("selectionAlias1", "selectionValue1")
	obtainedValue := recast.GetArrayOfInterfaces(selectionEntry.SelectionAlias[0], selectionEntry.SelectionValue[0])
	expectedValue := recast.GetArrayOfInterfaces("selectionAlias1", "selectionValue1")
	assert.Equalf(test, expectedValue, obtainedValue, "The selection entry obtained does not match what was set!")
	selectionEntry.ClearSelectionEntry()
	obtainedValue = recast.GetArrayOfInterfaces(len(selectionEntry.SelectionAlias), len(selectionEntry.SelectionValue))
	expectedValue = recast.GetArrayOfInterfaces(0, 0)
	assert.Equalf(test, expectedValue, obtainedValue, "The number of selection entries does not what was expected!")
}