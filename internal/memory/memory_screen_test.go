package memory

import (
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScreenLayerCreation(test *testing.T) {
	layerAlias := "MyAlias"
	layerWidth := 20
	layerHeight := 10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	layerParentAlias := ""

	InitializeScreenMemory()
	AddLayer(layerAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
	layer := GetLayer(layerAlias)
	obtainedResult := recast.GetArrayOfInterfaces(layer.Width, layer.Height, layer.ScreenXLocation, layer.ScreenYLocation, layer.ZOrder, layer.ParentAlias)
	expectedResult := recast.GetArrayOfInterfaces(20, 10, 1, 2, 1, "")
	assert.Equalf(test, expectedResult, obtainedResult, "The created layer was not added correctly!")
}

func TestScreenLayerParentIsCorrectlyLinked(test *testing.T) {
	layerWidth := 20
	layerHeight := 10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	parentAlias := "ParentAlias"
	childAlias := "ChildAlias"
	InitializeScreenMemory()
	AddLayer(parentAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, "")
	parentLayer := GetLayer(parentAlias)
	AddLayer(childAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, "ParentAlias")
	childLayer := GetLayer(childAlias)
	if parentLayer.IsParent != true {
		test.Errorf("Creating a child layer failed to update the 'IsParent' flag on the parent layer!")
	}
	if childLayer.ParentAlias != parentAlias {
		test.Errorf("Creating a child layer did not update itself with the correct parent alias!")
	}
}

func TestScreenLayerInvalidParent(test *testing.T) {
	layerAlias := "MyAlias"
	layerWidth := 20
	layerHeight := 10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	layerParentAlias := "BadParent"
	defer func() {
		if r := recover(); r == nil {
			test.Errorf("Creating a layer with a bad parent should have thrown a panic!")
		}
	}()
	InitializeScreenMemory()
	AddLayer(layerAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
}

func TestScreenLayerInvalidWidth(test *testing.T) {
	layerAlias := "MyAlias"
	layerWidth := -20
	layerHeight := 10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	layerParentAlias := ""
	defer func() {
		if r := recover(); r == nil {
			test.Errorf("Creating a layer with an invalid Width should have thrown a panic!")
		}
	}()
	InitializeScreenMemory()
	AddLayer(layerAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
}

func TestScreenLayerInvalidHeight(test *testing.T) {
	layerAlias := "MyAlias"
	layerWidth := 20
	layerHeight := -10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	layerParentAlias := ""
	defer func() {
		if r := recover(); r == nil {
			test.Errorf("Creating a layer with an invalid Height should have thrown a panic!")
		}
	}()
	InitializeScreenMemory()
	AddLayer(layerAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
}

func TestScreenLayerSimpleDelete(test *testing.T) {
	layerAlias := "MyAlias"
	layerWidth := 20
	layerHeight := 10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	layerParentAlias := ""
	InitializeScreenMemory()
	AddLayer(layerAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
	DeleteLayer(layerAlias)
	defer func() {
		if r := recover(); r == nil {
			test.Errorf("Obtaining a layer that has already been deleted should throw a panic!")
		}
	}()
	_ = GetLayer(layerAlias)
}

func TestScreenLayerChildrenDelete(test *testing.T) {
	layerWidth := 20
	layerHeight := 10
	layerXLocation := 1
	layerYLocation := 2
	layerZOrderPriority := 1
	layerParentAlias := "ParentAlias"

	childAlias1 := "ChildAlias1"
	childAlias2 := "ChildAlias2"
	childAlias3 := "ChildAlias3"

	InitializeScreenMemory()
	AddLayer(layerParentAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, "")
	AddLayer(childAlias1, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
	AddLayer(childAlias2, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
	AddLayer(childAlias3, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, layerParentAlias)
	var layerEntry = GetLayer(layerParentAlias)

	DeleteLayer(childAlias2)
	if layerEntry.IsParent != true {
		test.Errorf("The parent layer is no longer marked as a parent when it should have two children remaining!")
	}
	DeleteLayer(childAlias1)
	if layerEntry.IsParent != true {
		test.Errorf("The parent layer is no longer marked as a parent when it should have one children remaining!")
	}
	DeleteLayer(childAlias3)
	if layerEntry.IsParent == true {
		test.Errorf("The parent layer is no longer a parent, but is still marked as one!")
	}
}

func TestScreenLayerParentDelete(test *testing.T) {
	var layerWidth = 20
	var layerHeight = 10
	var layerXLocation = 1
	var layerYLocation = 2
	var layerZOrderPriority = 1

	var parentAlias = "ParentAlias"
	var childAlias1 = "ChildAlias1"
	var childAlias2 = "ChildAlias2"
	var childAlias3 = "ChildAlias3"

	InitializeScreenMemory()
	AddLayer(parentAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, "")
	AddLayer(childAlias1, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, parentAlias)
	AddLayer(childAlias2, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, parentAlias)
	AddLayer(childAlias3, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, parentAlias)
	DeleteLayer(parentAlias)
	if IsLayerExists(parentAlias) {
		test.Errorf("The parent layer exists when it should have been deleted!")
	}
	if IsLayerExists(childAlias1) {
		test.Errorf("The first child layer should have been deleted since the parent no longer exists!")
	}
	if IsLayerExists(childAlias2) {
		test.Errorf("The second child layer should have been deleted since the parent no longer exists!")
	}
	if IsLayerExists(childAlias3) {
		test.Errorf("The third child layer should have been deleted since the parent no longer exists!")
	}
}

func TestScreenLayerSubParentDelete(test *testing.T) {
	var layerWidth = 20
	var layerHeight = 10
	var layerXLocation = 1
	var layerYLocation = 2
	var layerZOrderPriority = 1

	var parentAlias = "ParentAlias"
	var childAlias1 = "ChildAlias1"
	var childAlias3 = "ChildAlias3"
	var subParent1 = "SubParent1"
	var subChild1 = "SubChild1"
	var subChild2 = "SubChild2"

	InitializeScreenMemory()
	AddLayer(parentAlias, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, "")
	AddLayer(childAlias1, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, parentAlias)
	AddLayer(subParent1, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, parentAlias)
	AddLayer(childAlias3, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, parentAlias)
	AddLayer(subChild1, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, subParent1)
	AddLayer(subChild2, layerXLocation, layerYLocation, layerWidth, layerHeight, layerZOrderPriority, subParent1)
	DeleteLayer(subParent1)
	if IsLayerExists(subChild1) {
		test.Errorf("Deleting the sub-parent layer did not delete the first sub-child layer as expected!")
	}
	if IsLayerExists(subChild2) {
		test.Errorf("Deleting the sub-parent layer did not delete the second sub-child layer as expected!")
	}
	if IsLayerExists(subParent1) {
		test.Errorf("The sub-parent was supposed to be deleted, but it still exists!")
	}
}

func TestScreenLayerSorting(test *testing.T) {
	var layerWidth = 20
	var layerHeight = 10
	var layerXLocation = 1
	var layerYLocation = 2
	var layerParentAlias = ""
	InitializeScreenMemory()
	AddLayer("Alias1", layerXLocation, layerYLocation, layerWidth, layerHeight, 6, layerParentAlias)
	AddLayer("Alias2", layerXLocation, layerYLocation, layerWidth, layerHeight, 4, layerParentAlias)
	AddLayer("Alias3", layerXLocation, layerYLocation, layerWidth, layerHeight, 1, layerParentAlias)
	AddLayer("Alias4", layerXLocation, layerYLocation, layerWidth, layerHeight, 3, layerParentAlias)
	AddLayer("Alias5", layerXLocation, layerYLocation, layerWidth, layerHeight, 9, layerParentAlias)
	AddLayer("Alias6", layerXLocation, layerYLocation, layerWidth, layerHeight, 8, layerParentAlias)
	var pairList LayerAliasZOrderPairList = GetSortedLayerMemoryAliasSlice()
	if pairList[0].Key != "Alias3" || pairList[1].Key != "Alias4" || pairList[2].Key != "Alias2" ||
		pairList[3].Key != "Alias1" || pairList[4].Key != "Alias6" || pairList[5].Key != "Alias5" {
		test.Errorf("The sorted screen layer pair list is not correct")
	}
}
