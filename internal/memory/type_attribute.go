package memory

import (
	"github.com/supercom32/dosktop/constant"
	"encoding/json"
)

type AttributeEntryType struct {
	ForegroundColor          int32
	BackgroundColor          int32
	IsBold                   bool
	IsUnderlined             bool
	IsReversed               bool
	IsBlinking               bool
	IsItalic                 bool
	ForegroundTransformValue float32
	BackgroundTransformValue float32
	CellId                   int
	CellType                 int
	CellAlias                string
}

func (shared AttributeEntryType) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ForegroundColor int32
		BackgroundColor int32
		IsBold bool
		IsUnderlined bool
		IsReversed bool
		IsBlinking bool
		IsItalic bool
		ForegroundTransformValue float32
		BackgroundTransformValue float32
		CellId int
		CellType int
		CellAlias string
	}{
		ForegroundColor: shared.ForegroundColor,
		BackgroundColor: shared.BackgroundColor,
		IsBold: shared.IsBold,
		IsUnderlined: shared.IsUnderlined,
		IsReversed: shared.IsReversed,
		IsBlinking: shared.IsBlinking,
		IsItalic: shared.IsItalic,
		ForegroundTransformValue: shared.ForegroundTransformValue,
		BackgroundTransformValue: shared.ForegroundTransformValue,
		CellId: shared.CellId,
		CellType: shared.CellType,
		CellAlias: shared.CellAlias,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (shared AttributeEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewAttributeEntry(existingAttributeEntry ...*AttributeEntryType) AttributeEntryType {
	var attributeEntry AttributeEntryType
	if existingAttributeEntry != nil {
		attributeEntry.ForegroundColor = existingAttributeEntry[0].ForegroundColor
		attributeEntry.BackgroundColor = existingAttributeEntry[0].BackgroundColor
		attributeEntry.IsBold = existingAttributeEntry[0].IsBold
		attributeEntry.IsUnderlined = existingAttributeEntry[0].IsUnderlined
		attributeEntry.IsReversed = existingAttributeEntry[0].IsReversed
		attributeEntry.IsBlinking = existingAttributeEntry[0].IsBlinking
		attributeEntry.IsItalic = existingAttributeEntry[0].IsItalic
		attributeEntry.ForegroundTransformValue = existingAttributeEntry[0].ForegroundTransformValue
		attributeEntry.BackgroundTransformValue = existingAttributeEntry[0].BackgroundTransformValue
		attributeEntry.CellId = existingAttributeEntry[0].CellId
		attributeEntry.CellType = existingAttributeEntry[0].CellType
		attributeEntry.CellAlias = existingAttributeEntry[0].CellAlias
	} else {
		attributeEntry.ForegroundTransformValue = 1
		attributeEntry.BackgroundTransformValue = 1
		attributeEntry.ForegroundColor = constants.AnsiColorByIndex[15]
		attributeEntry.BackgroundColor = constants.AnsiColorByIndex[0]
		attributeEntry.CellId = constants.NullCellId
		attributeEntry.CellType = constants.NullCellType
	}
	return attributeEntry
}
