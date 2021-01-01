package memory

import (
	"github.com/supercom32/dosktop/constant"
	"encoding/json"
)

type TextStyleEntryType struct {
	ForegroundColor          int32
	BackgroundColor          int32
	IsBold                   bool
	IsUnderlined             bool
	IsReversed               bool
	IsBlinking               bool
	IsItalic                 bool
	ForegroundTransformValue float32
	BackgroundTransformValue float32
}

func (shared TextStyleEntryType) MarshalJSON() ([]byte, error) {
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
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (shared TextStyleEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewTextStyleEntry(existingAttributeEntry ...*TextStyleEntryType) TextStyleEntryType {
	var attributeEntry TextStyleEntryType
	if existingAttributeEntry != nil {
		attributeEntry.ForegroundColor = existingAttributeEntry[0].ForegroundColor
		attributeEntry.BackgroundColor = existingAttributeEntry[0].BackgroundColor
		attributeEntry.IsBold = existingAttributeEntry[0].IsBold
		attributeEntry.IsUnderlined = existingAttributeEntry[0].IsUnderlined
		attributeEntry.IsReversed = existingAttributeEntry[0].IsReversed
		attributeEntry.IsBlinking = existingAttributeEntry[0].IsBlinking
		attributeEntry.IsItalic = existingAttributeEntry[0].IsItalic
	} else {
		attributeEntry.ForegroundColor = constants.AnsiColorByIndex[15]
		attributeEntry.BackgroundColor = constants.AnsiColorByIndex[0]
	}
	return attributeEntry
}
