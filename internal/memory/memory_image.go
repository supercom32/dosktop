package memory

import "fmt"

var ImageMemory map[string]*ImageEntryType

func InitializeImageMemory() {
	ImageMemory = make(map[string]*ImageEntryType)
}

func AddImage(imageAlias string, imageEntry ImageEntryType) {
	// verify if any errors occurred?
	ImageMemory[imageAlias] = &imageEntry
}

func GetImage(imageAlias string) ImageEntryType {
	if ImageMemory[imageAlias] == nil {
		panic(fmt.Sprintf("The requested image with alias '%s' could not be returned since it does not exist.", imageAlias))
	}
	return *ImageMemory[imageAlias]
}
func DeleteImage(imageAlias string) {
	delete(ImageMemory, imageAlias)
}