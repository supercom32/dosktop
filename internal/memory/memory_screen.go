package memory

import (
	"fmt"
	"sort"
)

var screenWidth int
var screenHeight int
var ScreenMemory map[string]*LayerEntryType

func InitializeScreenMemory() {
	ScreenMemory = make(map[string]*LayerEntryType)
}

func AddLayer(layerAlias string, xLocation int, yLocation int, width int, height int, zOrderPriority int, parentAlias string) {
	if width <= 0 {
		panic(fmt.Sprintf("The layer '%s' could not be created since a Width of '%d' was specified!", layerAlias, width))
	}
	if height <= 0 {
		panic(fmt.Sprintf("The layer '%s' could not be created since a Height of '%d' was specified!", layerAlias, height))
	}
	layerEntry := NewLayerEntry(width, height)
	layerEntry.LayerAlias = layerAlias
	layerEntry.ScreenXLocation = xLocation
	layerEntry.ScreenYLocation = yLocation
	layerEntry.ZOrder = zOrderPriority
	layerEntry.ParentAlias = parentAlias
	if parentAlias != "" {
		if IsLayerExists(parentAlias) {
			parentEntry := GetLayer(parentAlias)
			parentEntry.IsParent = true
		} else {
			panic(fmt.Sprintf("The layer '%s' could not be created since the parent alias '%s' does not exist!", layerAlias, parentAlias))
		}

	}
	//dosktop.commonResource.layerAlias = layerAlias
	ScreenMemory[layerAlias] = &layerEntry
}

func GetLayer(layerAlias string) *LayerEntryType {
	if !IsLayerExists(layerAlias) {
		panic(fmt.Sprintf("The layer '%s' could not be obtained since it does not exist!", layerAlias))
	}
	return ScreenMemory[layerAlias]
}

func DeleteLayer(layerAlias string) {
	layerEntry := GetLayer(layerAlias)
	parentAlias := layerEntry.ParentAlias
	isParent := layerEntry.IsParent
	delete(ScreenMemory, layerAlias)
	if parentAlias != "" {
		if IsLayerExists(parentAlias) && !IsAParent(parentAlias) {
			layerEntry = GetLayer(parentAlias)
			layerEntry.IsParent = false
		}
	}
	if isParent {
		DeleteAllChildrenOfParent(layerAlias)
	}
}

func DeleteAllChildrenOfParent(parentAlias string) {
	for currentKey, currentValue := range ScreenMemory {
		if currentValue.ParentAlias == parentAlias {
			DeleteLayer(currentKey)
		}
	}
}
func IsLayerExists(layerAlias string) bool {
	if _, isExist := ScreenMemory[layerAlias]; isExist {
		return true
	}
	return false
}
func IsAParent(parentAlias string) bool {
	isParent := false
	for _, currentValue := range ScreenMemory {
		if currentValue.ParentAlias == parentAlias {
			isParent = true
		}
	}
	return isParent
}

type layerAliasZOrderPair struct {
	Key   string
	Value int
}
type LayerAliasZOrderPairList []layerAliasZOrderPair

func GetSortedLayerMemoryAliasSlice() LayerAliasZOrderPairList {
	pairList := make(LayerAliasZOrderPairList, len(ScreenMemory))
	currentEntry := 0
	for currentKey, currentValue := range ScreenMemory {
		pairList[currentEntry].Key = currentKey
		pairList[currentEntry].Value = currentValue.ZOrder
		currentEntry++
	}
	sort.SliceStable(pairList, func(firstIndex, secondIndex int) bool {
		return pairList[firstIndex].Value < pairList[secondIndex].Value
	})
	return pairList
}
