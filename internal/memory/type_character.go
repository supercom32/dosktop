package memory

import (
	"encoding/json"
)

type CharacterEntryType struct {
	Character      rune
	AttributeEntry AttributeEntryType
	LayerAlias     string
}

func (shared CharacterEntryType) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Character      rune
		AttributeEntry AttributeEntryType
		LayerAlias     string
	}{
		Character: shared.Character,
		AttributeEntry: shared.AttributeEntry,
		LayerAlias: shared.LayerAlias,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (shared CharacterEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewCharacterEntry(existingCharacterEntry ...*CharacterEntryType) CharacterEntryType {
	var characterEntry CharacterEntryType
	if existingCharacterEntry != nil {
		characterEntry.Character = existingCharacterEntry[0].Character
		characterEntry.AttributeEntry = NewAttributeEntry(&existingCharacterEntry[0].AttributeEntry)
		characterEntry.LayerAlias = existingCharacterEntry[0].LayerAlias
	} else {
		characterEntry.AttributeEntry = NewAttributeEntry()
	}
	return characterEntry
}