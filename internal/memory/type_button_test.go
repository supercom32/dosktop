package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestButtonTypeCreation(test *testing.T) {
	firstButtonEntry := NewButtonEntry()
	firstButtonEntry.Height = 10
	firstButtonEntry.Width = 11
	firstButtonEntry.ButtonAlias = "MyButton"
	firstButtonEntry.IsPressed = true
	firstButtonEntry.IsSelected = true
	firstButtonEntry.StyleEntry = NewTuiStyleEntry()
	firstButtonEntry.XLocation = 1
	firstButtonEntry.YLocation = 2
	secondButtonEntry := NewButtonEntry()
	assert.NotEqualf(test, secondButtonEntry, firstButtonEntry, "The second Character object should not be the same as the first, as manipulating it should only effect itself.")

	secondButtonEntry = NewButtonEntry(&firstButtonEntry)
	assert.Equalf(test, secondButtonEntry, firstButtonEntry, "The first layer is not the same as the second, even though it should be an identical clone.")
}
