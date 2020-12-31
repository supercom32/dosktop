package memory

import "encoding/json"

// func DrawButton(LayerAlias string, ButtonLabel string, StyleEntry TuiStyleEntryType, IsPressed bool, IsSelected bool, XLocation int, YLocation int, Width int, Height int) {
type ButtonEntryType struct {
	StyleEntry  TuiStyleEntryType
	ButtonAlias string
	ButtonLabel string
	IsPressed   bool
	IsSelected  bool
	XLocation   int
	YLocation   int
	Width       int
	Height      int
}

func (shared ButtonEntryType) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		StyleEntry  TuiStyleEntryType
		ButtonAlias string
		ButtonLabel string
		IsPressed   bool
		IsSelected  bool
		XLocation   int
		YLocation   int
		Width       int
		Height      int
	}{
		StyleEntry: shared.StyleEntry,
		ButtonAlias: shared.ButtonAlias,
		ButtonLabel: shared.ButtonLabel,
		IsPressed: shared.IsPressed,
		IsSelected: shared.IsSelected,
		XLocation: shared.XLocation,
		YLocation: shared.YLocation,
		Width: shared.Width,
		Height: shared.Height,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (shared ButtonEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewButtonEntry(existingButtonEntry ...*ButtonEntryType) ButtonEntryType {
	var buttonEntry ButtonEntryType
	if existingButtonEntry != nil {
		buttonEntry.StyleEntry = NewTuiStyleEntry(&existingButtonEntry[0].StyleEntry)
		buttonEntry.ButtonAlias = existingButtonEntry[0].ButtonAlias
		buttonEntry.ButtonLabel = existingButtonEntry[0].ButtonLabel
		buttonEntry.IsPressed = existingButtonEntry[0].IsPressed
		buttonEntry.IsSelected = existingButtonEntry[0].IsSelected
		buttonEntry.XLocation = existingButtonEntry[0].XLocation
		buttonEntry.YLocation = existingButtonEntry[0].YLocation
		buttonEntry.Width = existingButtonEntry[0].Width
		buttonEntry.Height = existingButtonEntry[0].Height
	}
	return buttonEntry
}

func IsButtonEntryEqual(sourceButtonEntry *ButtonEntryType, targetButtonEntry *ButtonEntryType) bool {
	if sourceButtonEntry.StyleEntry == targetButtonEntry.StyleEntry &&
		sourceButtonEntry.ButtonAlias == targetButtonEntry.ButtonAlias &&
		sourceButtonEntry.ButtonLabel == targetButtonEntry.ButtonLabel &&
		sourceButtonEntry.IsPressed == targetButtonEntry.IsPressed &&
		sourceButtonEntry.IsSelected == targetButtonEntry.IsSelected &&
		sourceButtonEntry.XLocation == targetButtonEntry.XLocation &&
		sourceButtonEntry.YLocation == targetButtonEntry.YLocation &&
		sourceButtonEntry.Width == targetButtonEntry.Width &&
		sourceButtonEntry.Height == targetButtonEntry.Height {
		return true
	}
	return false
}
