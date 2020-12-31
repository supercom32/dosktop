package memory

import (
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/stretchr/testify/assert"
	"github.com/supercom32/dosktop/constant"
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
	firstStyleEntry.TextForegroundColor = constant.AnsiColorByIndex[1]
	firstStyleEntry.TextBackgroundColor = constant.AnsiColorByIndex[2]
	firstStyleEntry.TextInputForegroundColor = constant.AnsiColorByIndex[3]
	firstStyleEntry.TextInputBackgroundColor = constant.AnsiColorByIndex[4]
	firstStyleEntry.CursorForegroundColor = constant.AnsiColorByIndex[5]
	firstStyleEntry.CursorBackgroundColor = constant.AnsiColorByIndex[6]
	firstStyleEntry.MenuForegroundColor = constant.AnsiColorByIndex[7]
	firstStyleEntry.MenuBackgroundColor = constant.AnsiColorByIndex[8]
	firstStyleEntry.HighlightForegroundColor = constant.AnsiColorByIndex[9]
	firstStyleEntry.HighlightBackgroundColor = constant.AnsiColorByIndex[10]
	firstStyleEntry.ButtonRaisedColor = constant.AnsiColorByIndex[11]
	firstStyleEntry.ButtonForegroundColor = constant.AnsiColorByIndex[12]
	firstStyleEntry.ButtonBackgroundColor = constant.AnsiColorByIndex[13]
	firstStyleEntry.IsSquareFont = true
	firstStyleEntry.IsWindowFooterDrawn = true
	firstStyleEntry.IsWindowHeaderDrawn = true
	firstStyleEntry.MenuTextAlignment = constant.LeftAligned
	secondStyleEntry := NewTuiStyleEntry()
	firstResult := recast.GetArrayOfInterfaces(firstStyleEntry)
	secondResult := recast.GetArrayOfInterfaces(secondStyleEntry)
	assert.NotEqualf(test, secondResult, firstResult,"The first style entry is the same as the second, even though it should be different.")

	secondStyleEntry = NewTuiStyleEntry(&firstStyleEntry)
	secondResult = recast.GetArrayOfInterfaces(secondStyleEntry)
	assert.Equalf(test, secondResult, firstResult, "The first style entry is not the same as the second, even though it should be an identical clone.")
}
