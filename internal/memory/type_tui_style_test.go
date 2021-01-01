package memory

import (
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/stretchr/testify/assert"
	"github.com/supercom32/dosktop/constants"
	"testing"
)

func TestStyleTypeCreation(test *testing.T) {
	firstStyleEntry := NewTuiStyleEntry()
	firstStyleEntry.UpperLeftCorner = 'a'
	firstStyleEntry.UpperRightCorner = 'b'
	firstStyleEntry.HorizontalLine = 'c'
	firstStyleEntry.LeftSideTConnector = 'd'
	firstStyleEntry.RightSideTConnector = 'e'
	firstStyleEntry.UpSideTConnector = 'f'
	firstStyleEntry.DownSideTConnector = 'g'
	firstStyleEntry.VerticalLine = 'g'
	firstStyleEntry.LowerRightCorner = 'h'
	firstStyleEntry.LowerLeftCorner = 'i'
	firstStyleEntry.CrossConnector = 'j'
	firstStyleEntry.DesktopPattern = 'k'
	firstStyleEntry.ProgressBarBackgroundPattern = 'l'
	firstStyleEntry.ProgressBarForegroundPattern = 'm'
	firstStyleEntry.TextForegroundColor = constants.AnsiColorByIndex[1]
	firstStyleEntry.TextBackgroundColor = constants.AnsiColorByIndex[2]
	firstStyleEntry.TextInputForegroundColor = constants.AnsiColorByIndex[3]
	firstStyleEntry.TextInputBackgroundColor = constants.AnsiColorByIndex[4]
	firstStyleEntry.CursorForegroundColor = constants.AnsiColorByIndex[5]
	firstStyleEntry.CursorBackgroundColor = constants.AnsiColorByIndex[6]
	firstStyleEntry.MenuForegroundColor = constants.AnsiColorByIndex[7]
	firstStyleEntry.MenuBackgroundColor = constants.AnsiColorByIndex[8]
	firstStyleEntry.HighlightForegroundColor = constants.AnsiColorByIndex[9]
	firstStyleEntry.HighlightBackgroundColor = constants.AnsiColorByIndex[10]
	firstStyleEntry.ButtonRaisedColor = constants.AnsiColorByIndex[11]
	firstStyleEntry.ButtonForegroundColor = constants.AnsiColorByIndex[12]
	firstStyleEntry.ButtonBackgroundColor = constants.AnsiColorByIndex[13]
	firstStyleEntry.IsSquareFont = true
	firstStyleEntry.IsWindowFooterDrawn = true
	firstStyleEntry.IsWindowHeaderDrawn = true
	firstStyleEntry.MenuTextAlignment = constants.LeftAligned
	secondStyleEntry := NewTuiStyleEntry()
	firstResult := recast.GetArrayOfInterfaces(firstStyleEntry)
	secondResult := recast.GetArrayOfInterfaces(secondStyleEntry)
	assert.NotEqualf(test, secondResult, firstResult,"The first style entry is the same as the second, even though it should be different.")

	secondStyleEntry = NewTuiStyleEntry(&firstStyleEntry)
	secondResult = recast.GetArrayOfInterfaces(secondStyleEntry)
	assert.Equalf(test, secondResult, firstResult, "The first style entry is not the same as the second, even though it should be an identical clone.")
}
