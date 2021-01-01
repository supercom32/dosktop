package dosktop

import (
	"github.com/supercom32/dosktop/internal/memory"
	"fmt"
	"github.com/gdamore/tcell"
	"strings"
)

/*
updateEventQueues allows you to update all event queues so that information
such as mouse clicks, keystrokes, and other events are properly registered.
*/
func updateEventQueues() {
	event := commonResource.screen.PollEvent()
	switch event := event.(type) {
	case *tcell.EventResize:
		commonResource.screen.Sync()
	case *tcell.EventKey:
		keystroke := ""
		if strings.Contains(event.Name(), "Rune") {
			keystroke = fmt.Sprintf("%c", event.Rune())
		} else {
			keystroke = strings.ToLower(event.Name())
		}
		memory.KeyboardMemory.AddKeystrokeToKeyboardBuffer(keystroke)
	case *tcell.EventMouse:
		mouseXLocation, mouseYLocation := event.Position()
		var mouseButtonNumber uint
		mouseButton := event.Buttons()
		for index := uint(0); index < 8; index++ {
			if int(mouseButton)&(1<<index) != 0 {
				mouseButtonNumber = index + 1
			}
		}
		wheelState := ""
		if mouseButton&tcell.WheelUp != 0 {
			wheelState = "Up"
		} else if mouseButton&tcell.WheelDown != 0 {
			wheelState = "Down"
		} else if mouseButton&tcell.WheelLeft != 0 {
			wheelState = "Left"
		} else if mouseButton&tcell.WheelRight != 0 {
			wheelState = "Right"
		}
		memory.MouseMemory.SetMouseStatus(mouseXLocation, mouseYLocation, mouseButtonNumber, wheelState)
		updateButtonStates()
	}
}
