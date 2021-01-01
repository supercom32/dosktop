package dosktop

import "github.com/supercom32/dosktop/internal/memory"

type buttonHistoryType struct {
	buttonAlias string
	layerAlias  string
}

var buttonHistory buttonHistoryType

/*
getButtonClickIdentifier allows you to obtain the layer alias and the button
alias for the text cell currently under the mouse cursor. This is useful
for determining which button button the user has clicked (if any).
 */
func getButtonClickIdentifier(mouseXLocation int, mouseYLocation int) (string, string) {
	buttonAlias := ""
	layerAlias := ""
	layerEntry := commonResource.screenLayer
	mouseYLocationOnLayer := mouseYLocation - layerEntry.ScreenYLocation
	mouseXLocationOnLayer := mouseXLocation - layerEntry.ScreenXLocation
	if mouseYLocationOnLayer >= 0 && mouseXLocationOnLayer >= 0 &&
		mouseYLocationOnLayer < len(layerEntry.CharacterMemory) && mouseXLocationOnLayer < len(layerEntry.CharacterMemory[0]) {
		characterEntry := layerEntry.CharacterMemory[mouseYLocation-layerEntry.ScreenYLocation][mouseXLocation-layerEntry.ScreenXLocation]
		buttonAlias = characterEntry.AttributeEntry.CellAlias
		layerAlias = characterEntry.LayerAlias
	}
	return layerAlias, buttonAlias
}

/*
updateButtonStates allows you to update the state of all buttons. This needs
to be called frequently so that changes in button state are reflected to
the user as quickly as possible.
*/
func updateButtonStates() {
	mouseXLocation, mouseYLocation, buttonPressed, _ := memory.MouseMemory.GetMouseStatus()
	layerAlias, buttonAlias := getButtonClickIdentifier(mouseXLocation, mouseYLocation)
	if buttonPressed != 0 {
		if buttonAlias != "" {
			memory.ButtonMemory[layerAlias][buttonAlias].IsPressed = true
			buttonHistory.layerAlias = layerAlias
			buttonHistory.buttonAlias = buttonAlias
			UpdateDisplay()
		}
	} else {
		if buttonHistory.buttonAlias != "" {
			memory.ButtonMemory[buttonHistory.layerAlias][buttonHistory.buttonAlias].IsPressed = false
			UpdateDisplay()
		}
	}
}
