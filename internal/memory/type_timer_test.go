package memory

import (
	"github.com/supercom32/dosktop/internal/recast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTimerEntry(test *testing.T) {
	firstTimerEntry := NewTimerEntry()
	secondTimerEntry := NewTimerEntry()
	secondTimerEntry.IsTimerEnabled = true
	secondTimerEntry.TimerLength = 100
	secondTimerEntry.StartTime = 200

	obtainedResult := recast.GetArrayOfInterfaces(firstTimerEntry)
	expectedResult := recast.GetArrayOfInterfaces(secondTimerEntry)
	assert.NotEqualf(test, expectedResult, obtainedResult, "The first image entry is the same as the second, even though it should be different.")

	firstTimerEntry = NewTimerEntry(&secondTimerEntry)
	obtainedResult = recast.GetArrayOfInterfaces(firstTimerEntry)
	assert.Equalf(test, expectedResult, obtainedResult, "The first image entry is not the same as the second, even though it should be an identical clone.")


}
