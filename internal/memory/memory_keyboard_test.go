package memory

import "testing"

func TestKeyboardMemory(test *testing.T) {
	KeyboardMemory.AddKeystrokeToKeyboardBuffer("a")
	KeyboardMemory.AddKeystrokeToKeyboardBuffer("b")
	KeyboardMemory.AddKeystrokeToKeyboardBuffer("c")
	if KeyboardMemory.GetKeystrokeFromKeyboardBuffer() != "a" {
		test.Errorf("The first keyboard Character was not returned when it should be next in queue.")
	}
	if KeyboardMemory.GetKeystrokeFromKeyboardBuffer() != "b" {
		test.Errorf("The second keyboard Character was not returned when it should be next in queue.")
	}
	if KeyboardMemory.GetKeystrokeFromKeyboardBuffer() != "c" {
		test.Errorf("The second keyboard Character was not returned when it should be next in queue.")
	}
	if KeyboardMemory.GetKeystrokeFromKeyboardBuffer() != "" {
		test.Errorf("No keyboard keystrokes should have been returned, but one was given.")
	}
}
