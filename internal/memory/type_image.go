package memory

import (
	"encoding/json"
	"image"
)

type ImageEntryType struct {
	ImageData  image.Image
	LayerEntry LayerEntryType
}

func (shared ImageEntryType) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ImageData  image.Image
		LayerEntry LayerEntryType
	}{
		ImageData: shared.ImageData,
		LayerEntry: shared.LayerEntry,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (shared ImageEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewImageEntry(existingImageEntry ...*ImageEntryType) ImageEntryType {
	var imageEntry ImageEntryType
	if existingImageEntry != nil {
		imageEntry.ImageData = existingImageEntry[0].ImageData
		imageEntry.LayerEntry = existingImageEntry[0].LayerEntry
	}
	return imageEntry
}

