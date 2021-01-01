package dosktop

import (
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimerFunctionality(test *testing.T) {
	memory.InitializeTimerMemory()
	memory.AddTimer("TestTimer", 3000, true)
	SleepInMilliseconds(1000)
	assert.Equalf(test, IsTimerExpired("TestTimer"), false, "The timer was flagged as expired when not enough time has elapsed.")
	SleepInMilliseconds(2500)
	assert.Equalf(test, IsTimerExpired("TestTimer"), true, "The timer was not flagged as expired when more time has elapsed.")
}

func TestSetTimerFunctionality(test *testing.T) {
	memory.InitializeTimerMemory()
	memory.AddTimer("TestTimer", 9000, false)
	SetTimer("TestTimer", 3000, true)
	SleepInSeconds(1)
	assert.Equalf(test, IsTimerExpired("TestTimer"), false, "The timer was flagged as expired when not enough time has elapsed.")

	Sleep(2500)
	assert.Equalf(test, IsTimerExpired("TestTimer"), true, "The timer was not flagged as expired when more time has elapsed.")
}

func TestResetTimerFunctionality(test *testing.T) {
	memory.InitializeTimerMemory()
	memory.AddTimer("TestTimer", 1000, true)
	SleepInMilliseconds(1500)
	assert.Equalf(test, true, IsTimerExpired("TestTimer"), "The initial timer was not flagged as expired when more time has elapsed.")

	StartTimer("TestTimer")
	SleepInMilliseconds(500)
	assert.Equalf(test, false, IsTimerExpired("TestTimer"), "The reset timer was flagged as expired when not enough time has elapsed.")

	SleepInMilliseconds(1000)
	assert.Equalf(test, true, IsTimerExpired("TestTimer"), "The reset timer was not flagged as expired when more time has elapsed.")
}
