package memory

import (
	"sync"
)

type keyboardMemoryType struct {
	keyboardMemory []string
	mutex          sync.Mutex
}

var KeyboardMemory keyboardMemoryType

func (shared *keyboardMemoryType) AddKeystrokeToKeyboardBuffer(keystroke ...string) {
	shared.mutex.Lock()
	for _, currentKeystroke := range keystroke {
		shared.keyboardMemory = append(shared.keyboardMemory, currentKeystroke)
	}
	shared.mutex.Unlock()
}

func (shared *keyboardMemoryType) GetKeystrokeFromKeyboardBuffer() string {
	if shared.keyboardMemory == nil || len(shared.keyboardMemory) == 0 {
		return ""
	}
	keystroke := ""
	shared.mutex.Lock()
	keystroke = shared.keyboardMemory[0]
	shared.keyboardMemory = shared.keyboardMemory[1:]
	shared.mutex.Unlock()
	return keystroke
}
