package memory

import "sync"

type mouseMemoryType struct {
	xLocation     int
	yLocation     int
	buttonPressed uint
	wheelState    string
	mutex         sync.Mutex
}

var MouseMemory mouseMemoryType

func (shared *mouseMemoryType) ClearMouseMemory() {
	shared.mutex.Lock()
	shared.xLocation = -1
	shared.yLocation = -1
	shared.buttonPressed = 0
	shared.wheelState = ""
	shared.mutex.Unlock()
}

func (shared *mouseMemoryType) SetMouseStatus(xLocation int, yLocation int, buttonPressed uint, wheelState string) {
	shared.mutex.Lock()
	shared.xLocation = xLocation
	shared.yLocation = yLocation
	shared.buttonPressed = buttonPressed
	shared.wheelState = wheelState
	shared.mutex.Unlock()
}

func (shared *mouseMemoryType) GetMouseStatus() (int, int, uint, string) {
	shared.mutex.Lock()
	currentXLocation := shared.xLocation
	currentYLocation := shared.yLocation
	currentButtonPressed := shared.buttonPressed
	currentWheelState := shared.wheelState
	shared.mutex.Unlock()
	return currentXLocation, currentYLocation, currentButtonPressed, currentWheelState
}

func (shared *mouseMemoryType) IsMouseInBoundingBox(xLocation int, yLocation int, width int, height int) bool {
	mouseXLocation, mouseYLocation, _, _ := shared.GetMouseStatus()
	if mouseXLocation >= xLocation && mouseXLocation <= xLocation+width {
		if mouseYLocation >= yLocation && mouseYLocation <= yLocation+height {
			return true
		}
	}
	return false
}
