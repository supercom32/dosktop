package dosktop

import (
	"fmt"
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/supercom32/dosktop/internal/stringformat"
)

/*
drawProportionalHorizontalMenu allows you to draw a horizontal menu with
proportional spacing between menu options. In addition, the following
information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The viewport position represents the first menu item that needs to be
drawn. This is useful for menus sizes which are significantly larger than
the visible display area allocated for the menu.
 */
func drawProportionalHorizontalMenu(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, menuWidth int, numberOfItemsOnRow int, viewportPosition int, itemSelected int) {
	layerEntry := memory.GetLayer(layerAlias)
	menuAttributeEntry := memory.NewAttributeEntry()
	menuAttributeEntry.ForegroundColor = styleEntry.MenuForegroundColor
	menuAttributeEntry.BackgroundColor = styleEntry.MenuBackgroundColor
	highlightAttributeEntry := memory.NewAttributeEntry()
	highlightAttributeEntry.ForegroundColor = styleEntry.HighlightForegroundColor
	highlightAttributeEntry.BackgroundColor = styleEntry.HighlightBackgroundColor
	menuItemWidth := int(float64(menuWidth) / float64(numberOfItemsOnRow))
	menuWidthRemainder := menuWidth % menuItemWidth
	for currentMenuItemIndex := 0; currentMenuItemIndex < numberOfItemsOnRow; currentMenuItemIndex++ {
		if currentMenuItemIndex >= len(selectionEntry.SelectionValue) {
			menuItem := stringformat.GetFilledString(menuWidth, " ")
			arrayOfRunes := stringformat.GetRunesFromString(menuItem)
			printLayer(layerEntry, menuAttributeEntry, xLocation+(currentMenuItemIndex*menuItemWidth), yLocation, arrayOfRunes)
			continue
		}
		attributeEntry := menuAttributeEntry
		if currentMenuItemIndex == itemSelected {
			attributeEntry = highlightAttributeEntry
		}
		menuItemName := stringformat.GetFormattedString(" "+selectionEntry.SelectionValue[viewportPosition+currentMenuItemIndex]+" ", menuItemWidth, styleEntry.MenuTextAlignment)
		if currentMenuItemIndex == numberOfItemsOnRow-1 {
			menuItemName = stringformat.GetFormattedString(" "+selectionEntry.SelectionValue[viewportPosition+currentMenuItemIndex]+" ", menuItemWidth + menuWidthRemainder, styleEntry.MenuTextAlignment)
		}
		arrayOfRunes := stringformat.GetRunesFromString(menuItemName)
		attributeEntry.CellId = currentMenuItemIndex
		printLayer(layerEntry, attributeEntry, xLocation+(currentMenuItemIndex*menuItemWidth), yLocation, arrayOfRunes)
	}
}


/*
drawHorizontalMenu allows you to draw a horizontal menu with no padded spacing
between menu options. In addition, the following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The viewport position represents the first menu item that needs to be
drawn. This is useful for menus sizes which are significantly larger than
the visible display area allocated for the menu.
*/
func drawHorizontalMenu(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, itemSelected int) {
	layerEntry := memory.GetLayer(layerAlias)
	menuAttributeEntry := memory.NewAttributeEntry()
	menuAttributeEntry.ForegroundColor = styleEntry.MenuForegroundColor
	menuAttributeEntry.BackgroundColor = styleEntry.MenuBackgroundColor
	highlightAttributeEntry := memory.NewAttributeEntry()
	highlightAttributeEntry.ForegroundColor = styleEntry.HighlightForegroundColor
	highlightAttributeEntry.BackgroundColor = styleEntry.HighlightBackgroundColor
	currentXLocationOffset := 0
	for currentMenuItemIndex := 0; currentMenuItemIndex < len(selectionEntry.SelectionValue); currentMenuItemIndex++ {
		attributeEntry := menuAttributeEntry
		if currentMenuItemIndex == itemSelected {
			attributeEntry = highlightAttributeEntry
		}
		menuItemName := " " + selectionEntry.SelectionValue[currentMenuItemIndex] + " "
		arrayOfRunes := stringformat.GetRunesFromString(menuItemName)
		if xLocation+currentXLocationOffset+len(arrayOfRunes) > layerEntry.Width {
			return // Stop rendering if the menu item is larger than the layer supports.
			//panic(fmt.Sprintf("The menu item '%s' on layer '%s' cannot be drawn since it would be longer than the Width of the layer!", selectionEntry.selectionValue[currentMenuItemIndex], LayerAlias))
		}
		attributeEntry.CellId = currentMenuItemIndex
		printLayer(layerEntry, attributeEntry, xLocation+currentXLocationOffset, yLocation, arrayOfRunes)
		// Here we need to check how many wide characters we have printed so we can add double the amount
		currentXLocationOffset += len(arrayOfRunes) + stringformat.GetNumberOfWideCharacters(arrayOfRunes)
	}
}

/*
clearMouseMenuItemIdentifiers allows you to remove all mouse menu item
identifiers from a given text layer area. Mouse menu item identifiers are
attribute markers that are tagged to text cells which enable the easy
identification of which menu item is currently being selected.
*/
func clearMouseMenuItemIdentifiers(layerAlias string, xLocation int, yLocation int, width int, height int) {
	layerEntry := memory.GetLayer(layerAlias)
	for currentYLocation := 0; currentYLocation < yLocation+height; currentYLocation++ {
		for currentXLocation := 0; currentXLocation < xLocation+width; currentXLocation++ {
			if yLocation+currentYLocation < len(layerEntry.CharacterMemory) && xLocation+currentXLocation < len(layerEntry.CharacterMemory[0]) {
				characterEntry := &layerEntry.CharacterMemory[yLocation+currentYLocation][xLocation+currentXLocation]
				characterEntry.AttributeEntry.CellId = constants.NullCellId
			}
		}
	}
}

/*
GetSelectionFromProportionalHorizontalMenu allows you to obtain a user selection
from a proportional horizontal menu. This differs from a regular menu since all
menu selections are evenly spaced out. If you would like to have menu items
match the size of the selection, use 'GetSelectionFromHorizontalMenu' instead.
In addition, the following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The returned value is the selection alias for the item that was
chosen.
*/
func GetSelectionFromProportionalHorizontalMenu(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, menuWidth int, numberOfItemsOnRow int, defaultItemSelected int) string {
	selectionIndex := GetSelectionFromProportionalHorizontalMenuByIndex(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, menuWidth, numberOfItemsOnRow, defaultItemSelected)
	if selectionIndex == constants.NullSelectionIndex {
		return ""
	}
	return selectionEntry.SelectionAlias[selectionIndex]
}

/*
GetSelectionFromProportionalHorizontalMenuByIndex allows you to obtain a user
selection from a proportional horizontal menu. If you would like to have menu
items match the size of the selection, use
'GetSelectionFromHorizontalMenuByIndex' instead. In addition, the following
information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The returned value is the index number of your selection, where 0 is the
first item on your selection list.
 */
func GetSelectionFromProportionalHorizontalMenuByIndex(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, menuWidth int, numberOfItemsOnRow int, defaultItemSelected int) int {
	if menuWidth <= 0 {
		panic(fmt.Sprintf("The specified menu width of '%d' is invalid!", menuWidth))
	}
	if numberOfItemsOnRow <= 0 {
		panic(fmt.Sprintf("The specified number of menu items on row of '%d' is invalid!", numberOfItemsOnRow))
	}
	if defaultItemSelected < 0 || defaultItemSelected >= len(selectionEntry.SelectionValue) {
		panic(fmt.Sprintf("The specified default item selected of '%d' is invalid for a selection range of 0 to %d!", defaultItemSelected, len(selectionEntry.SelectionValue)))
	}
	layerEntry := memory.GetLayer(layerAlias)
	isItemSelected := false
	selectedItem := defaultItemSelected
	viewportPosition := 0
	// If your default item selected is greater than the number of items allowed on a line, set the viewport
	// and default selected item to be your selected item.
	if defaultItemSelected + 1 > numberOfItemsOnRow {
		viewportPosition = defaultItemSelected
		selectedItem = 0
		// If your selected item is greater than the last viewport position possible, set your viewport back
		// to the last possible position and adjust the selected item to match the default chosen.
		if len(selectionEntry.SelectionValue) - viewportPosition < numberOfItemsOnRow {
			selectedItem = len(selectionEntry.SelectionValue) - viewportPosition
			viewportPosition = len(selectionEntry.SelectionValue) - numberOfItemsOnRow
		}
	}
	returnValue := 0
	previouslySelectedItem := selectedItem
	drawProportionalHorizontalMenu(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, menuWidth, numberOfItemsOnRow, viewportPosition, selectedItem)
	UpdateDisplay()
	for isItemSelected == false {
		currentKeyPressed := Inkey()
		mouseXLocation, mouseYLocation, mouseButtonPressed, _ := memory.MouseMemory.GetMouseStatus()
		mouseCellIdentifier := getCellIdByLayerAlias(layerAlias, mouseXLocation, mouseYLocation)
		if mouseCellIdentifier != constants.NullCellId {
			selectedItem = mouseCellIdentifier
			if mouseButtonPressed == 1 {
				returnValue = selectedItem
				isItemSelected = true
			}
		}
		if currentKeyPressed == "left" {
			selectedItem--
			if selectedItem < 0 {
				selectedItem = 0
				viewportPosition--
				if viewportPosition < 0 {
					viewportPosition = 0
				}
			}
			previouslySelectedItem = -1
			memory.MouseMemory.ClearMouseMemory()
		}
		if currentKeyPressed == "right" {
			selectedItem++
			if selectedItem >= numberOfItemsOnRow {
				selectedItem = numberOfItemsOnRow - 1
				viewportPosition++
				if viewportPosition + numberOfItemsOnRow > len(selectionEntry.SelectionValue) {
					viewportPosition = len(selectionEntry.SelectionValue) - numberOfItemsOnRow
				}
			}
			previouslySelectedItem = -1
			memory.MouseMemory.ClearMouseMemory()
		}
		if currentKeyPressed == "enter" {
			returnValue = selectedItem
			isItemSelected = true
		}
		if currentKeyPressed == "esc" {
			returnValue = constants.NullSelectionIndex
			isItemSelected = true
		}
		if previouslySelectedItem != selectedItem {
			if selectedItem >= len(selectionEntry.SelectionValue) {
				selectedItem = len(selectionEntry.SelectionValue) - 1
			}
			drawProportionalHorizontalMenu(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, menuWidth, numberOfItemsOnRow, viewportPosition, selectedItem)
			UpdateDisplay()
			previouslySelectedItem = selectedItem
		}

	}
	clearMouseMenuItemIdentifiers(layerAlias, xLocation, yLocation, layerEntry.Width, 1)
	return returnValue
}

/*
GetSelectionFromHorizontalMenu allows you to obtain a user selection from a
horizontal menu. This is different from a proportional menu since all menu
items are drawn to the size of the current selection. If you want each
menu item to be of equal width, consider using
'GetSelectionFromProportionalHorizontalMenu' instead. In addition, the
following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The returned value is the selection alias for the item that was
chosen.
*/
func GetSelectionFromHorizontalMenu(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, defaultItemSelected int) string {
	selectionIndex := GetSelectionFromHorizontalMenuByIndex(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, defaultItemSelected)
	if selectionIndex == constants.NullSelectionIndex {
		return ""
	}
	return selectionEntry.SelectionAlias[selectionIndex]
}

/*
GetSelectionFromHorizontalMenuByIndex allows you to obtain a user selection
from a horizontal menu. This is different from a proportional menu since all
menu items are drawn to the size of the current selection. If you want each
menu item to be of equal width, consider using
'GetSelectionFromProportionalHorizontalMenuByIndex' instead. In addition, the
following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The returned value is the index number of your selection, where 0 is the
first item on your selection list.
*/
func GetSelectionFromHorizontalMenuByIndex(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, defaultItemSelected int) int {
	if defaultItemSelected < 0 || defaultItemSelected >= len(selectionEntry.SelectionValue) {
		panic(fmt.Sprintf("The specified default item selected of '%d' is invalid for a selection range of 0 to %d!", defaultItemSelected, len(selectionEntry.SelectionValue)))
	}
	layerEntry := memory.GetLayer(layerAlias)
	isItemSelected := false
	selectedItem :=  defaultItemSelected
	returnValue := 0
	previouslySelectedItem := 0
	numberOfMenuItems := len(selectionEntry.SelectionValue)
	drawHorizontalMenu(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, selectedItem)
	UpdateDisplay()
	for isItemSelected == false {
		currentKeyPressed := Inkey()
		mouseXLocation, mouseYLocation, mouseButtonPressed, _ := memory.MouseMemory.GetMouseStatus()
		mouseCellIdentifier := getCellIdByLayerAlias(layerAlias, mouseXLocation, mouseYLocation)
		if mouseCellIdentifier != constants.NullCellId {
			selectedItem = mouseCellIdentifier
			if mouseButtonPressed == 1 {
				returnValue = selectedItem
				isItemSelected = true
			}
		}
		if currentKeyPressed == "left" {
			selectedItem--
			if selectedItem < 0 {
				selectedItem = 0
			}
			previouslySelectedItem = -1
			memory.MouseMemory.ClearMouseMemory()
		}
		if currentKeyPressed == "right" {
			selectedItem++
			if selectedItem >= numberOfMenuItems {
				selectedItem = numberOfMenuItems - 1
			}
			previouslySelectedItem = -1
			memory.MouseMemory.ClearMouseMemory()
		}
		if currentKeyPressed == "enter" {
			returnValue = selectedItem
			isItemSelected = true
		}
		if currentKeyPressed == "esc" {
			returnValue = constants.NullSelectionIndex
			isItemSelected = true
		}
		if previouslySelectedItem != selectedItem {
			if selectedItem >= len(selectionEntry.SelectionValue) {
				selectedItem = len(selectionEntry.SelectionValue) - 1
			}
			drawHorizontalMenu(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, selectedItem)
			UpdateDisplay()
			previouslySelectedItem = selectedItem
		}

	}
	clearMouseMenuItemIdentifiers(layerAlias, xLocation, yLocation, layerEntry.Width, 1)
	return returnValue
}

/*
DrawVerticalMenu allows you to obtain a user selection from a horizontal menu.
In addition, the following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The viewport position represents the first menu item that needs to be
drawn. This is useful for menus sizes which are significantly larger than
the visible display area allocated for the menu.
*/
func DrawVerticalMenu(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, menuWidth int, menuHeight int, viewportPosition int, itemSelected int) {
	layerEntry := memory.GetLayer(layerAlias)
	menuAttributeEntry := memory.NewAttributeEntry()
	menuAttributeEntry.ForegroundColor = styleEntry.MenuForegroundColor
	menuAttributeEntry.BackgroundColor = styleEntry.MenuBackgroundColor
	highlightAttributeEntry := memory.NewAttributeEntry()
	highlightAttributeEntry.ForegroundColor = styleEntry.HighlightForegroundColor
	highlightAttributeEntry.BackgroundColor = styleEntry.HighlightBackgroundColor
	for currentMenuItemIndex := 0; currentMenuItemIndex < menuHeight; currentMenuItemIndex++ {
		if currentMenuItemIndex >= len(selectionEntry.SelectionValue) {
			menuItem := stringformat.GetFilledString(menuWidth, " ")
			arrayOfRunes := stringformat.GetRunesFromString(menuItem)
			printLayer(layerEntry, menuAttributeEntry, xLocation, yLocation+currentMenuItemIndex, arrayOfRunes)
			continue
		}
		attributeEntry := menuAttributeEntry
		if currentMenuItemIndex == itemSelected {
			attributeEntry = highlightAttributeEntry
		}
		menuItemName := stringformat.GetFormattedString(selectionEntry.SelectionValue[viewportPosition+currentMenuItemIndex], menuWidth, styleEntry.MenuTextAlignment)
		arrayOfRunes := stringformat.GetRunesFromString(menuItemName)
		attributeEntry.CellId = currentMenuItemIndex
		printLayer(layerEntry, attributeEntry, xLocation, yLocation+currentMenuItemIndex, arrayOfRunes)
	}
}

/*
GetSelectionFromVerticalMenu allows you to obtain a user selection
from a vertical menu. In addition, the following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The returned value is the selection alias for the item that was
chosen.
*/
func GetSelectionFromVerticalMenu (layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, menuWidth int, menuHeight int, defaultItemSelected int) string {
	selectionIndex := GetSelectionFromVerticalMenuByIndex(layerAlias, styleEntry , selectionEntry, xLocation, yLocation, menuWidth, menuHeight, defaultItemSelected)
	if selectionIndex == constants.NullSelectionIndex {
		return ""
	}
	return selectionEntry.SelectionAlias[selectionIndex]
}

/*
GetSelectionFromVerticalMenuByIndex allows you to obtain a user selection
from a vertical menu. In addition, the following information should be noted:

- If the location to draw a menu item falls outside of the range of the text
layer, then only the visible portion of your menu item will be drawn.

- The returned value is the index number of your selection, where 0 is the
first item on your selection list.
*/
func GetSelectionFromVerticalMenuByIndex(layerAlias string, styleEntry memory.TuiStyleEntryType, selectionEntry memory.SelectionEntryType, xLocation int, yLocation int, menuWidth int, menuHeight int, defaultItemSelected int ) int {
	isItemSelected := false
	selectedItem := defaultItemSelected
	viewportPosition := 0
	returnValue := 0
	previouslySelectedItem := 0
	if defaultItemSelected < 0 || defaultItemSelected >= len(selectionEntry.SelectionValue) {
		panic(fmt.Sprintf("The specified default item selected of '%d' is invalid for a selection range of 0 to %d!", defaultItemSelected, len(selectionEntry.SelectionValue)))
	}
	DrawVerticalMenu(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, menuWidth, menuHeight, viewportPosition, selectedItem)
	UpdateDisplay()
	for isItemSelected == false {
		currentKeyPressed := Inkey()
		mouseXLocation, mouseYLocation, mouseButtonPressed, _ := memory.MouseMemory.GetMouseStatus()
		mouseCellIdentifier := getCellIdByLayerAlias(layerAlias, mouseXLocation, mouseYLocation)
		if mouseCellIdentifier != constants.NullCellId {
			selectedItem = mouseCellIdentifier
			if mouseButtonPressed == 1 {
				returnValue = selectedItem
				isItemSelected = true
			}
		}
		if currentKeyPressed == "up" {
			selectedItem--
			if selectedItem < 0 {
				selectedItem = 0
				viewportPosition--
				if viewportPosition < 0 {
					viewportPosition = 0
				}
			}
			previouslySelectedItem = -1
			// When the mouse is off screen, the last known location becomes static. This can effect
			// Keyboard commands if the mouse's last position appears to be over menu items. So
			// We clear the mouse memory here so that it appears to be off screen until it is used again.
			memory.MouseMemory.ClearMouseMemory()
		}
		if currentKeyPressed == "down" {
			selectedItem++
			if selectedItem >= menuHeight {
				selectedItem = menuHeight - 1
				viewportPosition++
				if viewportPosition+menuHeight > len(selectionEntry.SelectionValue) {
					viewportPosition = len(selectionEntry.SelectionValue) - menuHeight
				}
			}
			previouslySelectedItem = -1
			memory.MouseMemory.ClearMouseMemory()
		}
		if currentKeyPressed == "enter" {
			returnValue = viewportPosition + selectedItem
			isItemSelected = true
		}
		if currentKeyPressed == "esc" {
			returnValue = constants.NullSelectionIndex
			isItemSelected = true
		}
		if previouslySelectedItem != selectedItem {
			if selectedItem >= len(selectionEntry.SelectionValue) {
				selectedItem = len(selectionEntry.SelectionValue) - 1
			}
			DrawVerticalMenu(layerAlias, styleEntry, selectionEntry, xLocation, yLocation, menuWidth, menuHeight, viewportPosition, selectedItem)
			UpdateDisplay()
			previouslySelectedItem = selectedItem
		}
	}
	clearMouseMenuItemIdentifiers(layerAlias, xLocation, yLocation, menuWidth, menuHeight)
	return returnValue
}

/*
GetInput allows you to obtain keyboard input from the user. This is useful for
letting applications accept configurations, settings, or other options in
an interactive way at runtime. In addition, the following information
should be noted:

- If the location specified for the input field  falls outside of the range
of the text layer, then only the visible portion of your input field will be
drawn.

- If the max length of your input field is less than or equal to 0, a panic
will be generated to fail as fast as possible.

- Password protection will echo back '*' characters to the terminal instead
of the actual characters entered.

- Specifying a default value will simply pre-populate the input field with
the value specified.

- If the cursor position moves outside of the visible display area of the
field, then the entire input field will shift to ensure the cursor is always
visible.
*/
func GetInput(layerAlias string, styleEntry memory.TuiStyleEntryType, xLocation int, yLocation int, width int, maxLengthAllowed int, IsPasswordProtected bool, defaultValue string) string {
	layerEntry := memory.GetLayer(layerAlias)
	inputString := defaultValue
	isScreenUpdateRequired := true
	var currentKeyPressed string
	var cursorPosition int
	var viewportPosition int
	if maxLengthAllowed <= 0 {
		panic(fmt.Sprintf("The specified maximum input length of '%d' is invalid!", maxLengthAllowed))
	}
	widthOfViewport := width - 1
	if len(inputString) > 0 {
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer("end")
	}
	for currentKeyPressed != "enter" {
		mouseXLocation, mouseYLocation, mouseButtonPressed, _ := memory.MouseMemory.GetMouseStatus()
		mouseCellIdentifier := getCellIdByLayerAlias(layerAlias, mouseXLocation, mouseYLocation)
		if mouseCellIdentifier != constants.NullCellId {
			if mouseButtonPressed == 1 {
				cursorPosition = mouseCellIdentifier
				isScreenUpdateRequired = true
			}
		}
		currentKeyPressed = Inkey()
		if currentKeyPressed != "" {
			// If character is pressed.
			if len(currentKeyPressed) == 1 {
				if len(inputString) < maxLengthAllowed {
					inputString = inputString[:viewportPosition+cursorPosition] + currentKeyPressed + inputString[viewportPosition+cursorPosition:]
					if cursorPosition < widthOfViewport {
						cursorPosition++
					} else {
						viewportPosition++
					}
					isScreenUpdateRequired = true
				}
			}
			if currentKeyPressed == "delete" {
				if inputString != "" {
					// Protect if nothing else to delete left of string
					if viewportPosition+cursorPosition+1 <= len(inputString) {
						inputString = inputString[:viewportPosition+cursorPosition] + inputString[viewportPosition+cursorPosition+1:]
						isScreenUpdateRequired = true
					}
				}
			}
			if currentKeyPressed == "home" {
				cursorPosition = 0
				viewportPosition = 0
				isScreenUpdateRequired = true
			}
			if currentKeyPressed == "end" {
				// If your current viewport shows the end of the input string, just move the cursor to the end of the string.
				if viewportPosition > len(inputString)-widthOfViewport {
					cursorPosition = len(inputString) - viewportPosition
				} else {
					// Otherwise advance viewport to end of input string.
					viewportPosition = len(inputString) - widthOfViewport
					if viewportPosition < 0 {
						// If input string is smaller than even one viewport block, just set cursor to end.
						viewportPosition = 0
						cursorPosition = len(inputString)
					} else {
						// Otherwise place cursor at end of viewport / string
						cursorPosition = widthOfViewport
					}
				}
				isScreenUpdateRequired = true
			}
			if currentKeyPressed == "backspace" || currentKeyPressed == "backspace2" {
				if inputString != "" {
					// Protect if nothing else to delete left of string
					if viewportPosition+cursorPosition-1 >= 0 {
						inputString = inputString[:viewportPosition+cursorPosition-1] + inputString[viewportPosition+cursorPosition:]
						cursorPosition--
						if cursorPosition < 1 {
							if len(inputString) < widthOfViewport {
								cursorPosition = viewportPosition + cursorPosition
								viewportPosition = 0
							} else {
								if viewportPosition != 0 { // If your not at the start of the input string
									if cursorPosition == 0 {
										viewportPosition = viewportPosition - widthOfViewport + 1
									} else {
										viewportPosition = viewportPosition - widthOfViewport
									}
									if viewportPosition < 0 {
										viewportPosition = 0
									}
									cursorPosition = widthOfViewport - 1
								}
							}
						}
						isScreenUpdateRequired = true
					}
				}
			}
			if currentKeyPressed == "left" {
				cursorPosition--
				if cursorPosition < 0 {
					if viewportPosition == 0 {
						cursorPosition = 0
					} else {
						viewportPosition = viewportPosition - widthOfViewport
						if viewportPosition < 0 {
							viewportPosition = 0
						}
						cursorPosition = widthOfViewport
					}
				}
				isScreenUpdateRequired = true
			}
			if currentKeyPressed == "right" {
				cursorPosition++
				if viewportPosition+cursorPosition >= len(inputString)+1 {
					cursorPosition--
				} else {
					if cursorPosition > widthOfViewport {
						viewportPosition += widthOfViewport
						cursorPosition = 0
					}
				}
				isScreenUpdateRequired = true
			}
		}
		if isScreenUpdateRequired {
			drawInputString(layerEntry, styleEntry, xLocation, yLocation, width, 0, false, stringformat.GetFilledString(width, " "))
			if IsPasswordProtected {
				passwordProtectedString := stringformat.GetFilledString(len(inputString), "*")
				drawInputString(layerEntry, styleEntry, xLocation, yLocation, widthOfViewport, viewportPosition, true, passwordProtectedString)
			} else {
				drawInputString(layerEntry, styleEntry, xLocation, yLocation, widthOfViewport, viewportPosition, true, inputString)
			}
			drawCursor(layerEntry, styleEntry, xLocation, yLocation, cursorPosition, false)
			UpdateDisplay()
			isScreenUpdateRequired = false
		}
	}
	return inputString
}

/*
drawCursor allows you to draw a cursor at the appropriate location for a
text field. In addition, the following information should be noted:

- If the location specified for the cursor range falls outside of the text
layer, then the cursor will only be rendered on the visible portion.

- The cursor position indicates how many spaces to the right of the starting x
and y location your cursor should be drawn at.

- If it is indicated that your cursor is moving backwards, then the space in
which the cursor was previously located will be automatically cleared.
*/
func drawCursor(layerEntry *memory.LayerEntryType, styleEntry memory.TuiStyleEntryType, xLocation int, yLocation int, cursorPosition int, isMovementBackwards bool) {
	attributeEntry := memory.NewAttributeEntry()
	attributeEntry.ForegroundColor = styleEntry.CursorForegroundColor
	attributeEntry.BackgroundColor = styleEntry.CursorBackgroundColor
	arrayOfRunes := stringformat.GetRunesFromString(string(styleEntry.CursorCharacter))
	printLayer(layerEntry, attributeEntry, xLocation+cursorPosition, yLocation, arrayOfRunes)
	if isMovementBackwards {
		arrayOfRunes = stringformat.GetRunesFromString(" ")
		printLayer(layerEntry, attributeEntry, xLocation+cursorPosition+1, yLocation, arrayOfRunes)
	}
}

/*
drawInputString allows you to draw a string for an input field. This is
different than regular printing, since input fields are usually
constrained for space and have the possibility of not being able to
show the entire string. In addition, the following information should be
noted:

- If the location specified for the input string falls outside the range of
the text layer, then only the visible portion will be displayed.

- Width indicates how large the visible area of your input string should be.

- String position indicates the location in your string where printing should
start. If the remainder of your string is too long for the specified width,
then only the visible portion will be displayed.
*/
func drawInputString(layerEntry *memory.LayerEntryType, styleEntry memory.TuiStyleEntryType, xLocation int, yLocation int, width int, stringPosition int, isCellIdsRequired bool, inputString string) {
	attributeEntry := memory.NewAttributeEntry()
	attributeEntry.ForegroundColor = styleEntry.TextInputForegroundColor
	attributeEntry.BackgroundColor = styleEntry.TextInputBackgroundColor
	attributeEntry.CellType = constants.CellTypeTextInput
	runeSlice := []rune(inputString)
	var safeSubstring string
	if stringPosition+width <= len(inputString) {
		safeSubstring = string(runeSlice[stringPosition : stringPosition+width])
	} else {
		safeSubstring = string(runeSlice[stringPosition : stringPosition+len(inputString)-stringPosition])
	}
	arrayOfRunes := stringformat.GetRunesFromString(safeSubstring)
	// Here we loop over each character to draw since we need to accommodate for unique
	// cell IDs (if required for mouse location detection).
	for currentRuneIndex := 0; currentRuneIndex < len(arrayOfRunes); currentRuneIndex ++ {
		if isCellIdsRequired {
			attributeEntry.CellId = currentRuneIndex
		}
		printLayer(layerEntry, attributeEntry, xLocation + currentRuneIndex, yLocation, []rune{arrayOfRunes[currentRuneIndex]})
	}
}
