package memory

type SelectionEntryType struct {
	SelectionAlias []string
	SelectionValue []string
}

func NewSelectionEntry() SelectionEntryType {
	var selectionEntry SelectionEntryType
	return selectionEntry
}

func (shared *SelectionEntryType) AddSelection(selectionAlias string, selectionValue string) {
	shared.SelectionAlias = append(shared.SelectionAlias, selectionAlias)
	shared.SelectionValue = append(shared.SelectionValue, selectionValue)
}

// ClearSelectionEntry asdasd
func (shared *SelectionEntryType) ClearSelectionEntry() {
	shared.SelectionAlias = nil
	shared.SelectionValue = nil
}
