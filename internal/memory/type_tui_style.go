package memory

import (
	"github.com/supercom32/dosktop/constant"
	"encoding/json"
)

// TuiStyleEntryType adsas
type TuiStyleEntryType struct {
	UpperLeftCorner              rune
	UpperRightCorner             rune
	HorizontalLine               rune
	LeftSideTConnector           rune
	RightSideTConnector          rune
	UpSideTConnector             rune
	DownSideTConnector           rune
	VerticalLine                 rune
	LowerRightCorner             rune
	LowerLeftCorner              rune
	CrossConnector               rune
	DesktopPattern               rune
	ProgressBarBackgroundPattern rune
	ProgressBarForegroundPattern rune
	IsSquareFont                 bool
	IsWindowHeaderDrawn          bool
	IsWindowFooterDrawn          bool
	TextForegroundColor          int32
	TextBackgroundColor          int32
	TextInputForegroundColor     int32
	TextInputBackgroundColor int32
	CursorCharacter    			rune
	CursorForegroundColor    int32
	CursorBackgroundColor    int32
	MenuForegroundColor      int32
	MenuBackgroundColor      int32
	HighlightForegroundColor int32
	HighlightBackgroundColor int32
	ButtonRaisedColor        int32
	ButtonForegroundColor    int32
	ButtonBackgroundColor    int32
	MenuTextAlignment        int
}

func (shared TuiStyleEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// NewTuiStyleEntry asdasd existingCharacterObject ...*CharacterEntryType) CharacterEntryType
func NewTuiStyleEntry(existingStyleEntry ...*TuiStyleEntryType) TuiStyleEntryType {
	var styleEntry TuiStyleEntryType
	if existingStyleEntry != nil {
		styleEntry.UpperLeftCorner = existingStyleEntry[0].UpperLeftCorner
		styleEntry.UpperRightCorner = existingStyleEntry[0].UpperRightCorner
		styleEntry.HorizontalLine = existingStyleEntry[0].HorizontalLine
		styleEntry.LeftSideTConnector = existingStyleEntry[0].LeftSideTConnector
		styleEntry.RightSideTConnector = existingStyleEntry[0].RightSideTConnector
		styleEntry.UpSideTConnector = existingStyleEntry[0].UpSideTConnector
		styleEntry.DownSideTConnector = existingStyleEntry[0].DownSideTConnector
		styleEntry.VerticalLine = existingStyleEntry[0].VerticalLine
		styleEntry.LowerRightCorner = existingStyleEntry[0].LowerRightCorner
		styleEntry.LowerLeftCorner = existingStyleEntry[0].LowerLeftCorner
		styleEntry.CrossConnector = existingStyleEntry[0].CrossConnector
		styleEntry.DesktopPattern = existingStyleEntry[0].DesktopPattern
		styleEntry.ProgressBarBackgroundPattern = existingStyleEntry[0].ProgressBarBackgroundPattern
		styleEntry.ProgressBarForegroundPattern = existingStyleEntry[0].ProgressBarForegroundPattern
		styleEntry.TextForegroundColor = existingStyleEntry[0].TextForegroundColor
		styleEntry.TextBackgroundColor = existingStyleEntry[0].TextBackgroundColor
		styleEntry.TextInputForegroundColor = existingStyleEntry[0].TextInputForegroundColor
		styleEntry.TextInputBackgroundColor = existingStyleEntry[0].TextInputBackgroundColor
		styleEntry.CursorCharacter = existingStyleEntry[0].CursorCharacter
		styleEntry.CursorForegroundColor = existingStyleEntry[0].CursorForegroundColor
		styleEntry.CursorBackgroundColor = existingStyleEntry[0].CursorBackgroundColor
		styleEntry.MenuForegroundColor = existingStyleEntry[0].MenuForegroundColor
		styleEntry.MenuBackgroundColor = existingStyleEntry[0].MenuBackgroundColor
		styleEntry.HighlightForegroundColor = existingStyleEntry[0].HighlightForegroundColor
		styleEntry.HighlightBackgroundColor = existingStyleEntry[0].HighlightBackgroundColor
		styleEntry.ButtonRaisedColor = existingStyleEntry[0].ButtonRaisedColor
		styleEntry.ButtonForegroundColor = existingStyleEntry[0].ButtonForegroundColor
		styleEntry.ButtonBackgroundColor = existingStyleEntry[0].ButtonBackgroundColor
		styleEntry.IsSquareFont = existingStyleEntry[0].IsSquareFont
		styleEntry.IsWindowFooterDrawn = existingStyleEntry[0].IsWindowFooterDrawn
		styleEntry.IsWindowHeaderDrawn = existingStyleEntry[0].IsWindowHeaderDrawn
		styleEntry.MenuTextAlignment = existingStyleEntry[0].MenuTextAlignment
	} else {
		styleEntry.UpperLeftCorner = constants.CharULCorner
		styleEntry.UpperRightCorner = constants.CharURCorner
		styleEntry.HorizontalLine = constants.CharHLine
		styleEntry.LeftSideTConnector = constants.CharSingleLineTLeft
		styleEntry.RightSideTConnector = constants.CharSingleLineTRight
		styleEntry.UpSideTConnector = constants.CharSingleLineTUp
		styleEntry.DownSideTConnector = constants.CharSingleLineTDown
		styleEntry.VerticalLine = constants.CharSingleLineVertical
		styleEntry.LowerRightCorner = constants.CharSingleLineLowerRightCorner
		styleEntry.LowerLeftCorner = constants.CharSingleLineLowerLeftCorner
		styleEntry.CrossConnector = constants.CharSingleLineCross
		styleEntry.DesktopPattern = constants.CharBlockSparce
		styleEntry.ProgressBarBackgroundPattern = constants.CharBlockSparce
		styleEntry.ProgressBarForegroundPattern = constants.CharBlockSolid
		styleEntry.TextForegroundColor = constants.AnsiColorByIndex[15]
		styleEntry.TextBackgroundColor = constants.AnsiColorByIndex[0]
		styleEntry.TextInputForegroundColor = constants.AnsiColorByIndex[15]
		styleEntry.TextInputBackgroundColor = constants.AnsiColorByIndex[0]
		styleEntry.CursorCharacter = constants.CharBlockSolid
		styleEntry.CursorForegroundColor = constants.AnsiColorByIndex[15]
		styleEntry.CursorBackgroundColor = constants.AnsiColorByIndex[0]
		styleEntry.MenuForegroundColor = constants.AnsiColorByIndex[15]
		styleEntry.MenuBackgroundColor = constants.AnsiColorByIndex[0]
		styleEntry.HighlightForegroundColor = constants.AnsiColorByIndex[0]
		styleEntry.HighlightBackgroundColor = constants.AnsiColorByIndex[15]
		styleEntry.ButtonRaisedColor = constants.AnsiColorByIndex[15]
		styleEntry.ButtonForegroundColor = constants.AnsiColorByIndex[0]
		styleEntry.ButtonBackgroundColor = constants.AnsiColorByIndex[7]
		styleEntry.IsSquareFont = false
		styleEntry.IsWindowFooterDrawn = false
		styleEntry.IsWindowHeaderDrawn = false
		styleEntry.MenuTextAlignment = constants.LeftAligned
	}
	return styleEntry
}
