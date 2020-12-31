package memory

import (
	"fmt"
	"time"
)

var TimerMemory map[string]*TimerEntryType

func InitializeTimerMemory() {
	TimerMemory = make(map[string]*TimerEntryType)
}

func AddTimer(timerAlias string, lengthOfTimerInMilliseconds int64, isTimerEnabled bool) {
	timerEntry := NewTimerEntry()
	timerEntry.IsTimerEnabled = isTimerEnabled
	timerEntry.StartTime = GetCurrentTimeInMilliseconds()
	timerEntry.TimerLength = lengthOfTimerInMilliseconds
	TimerMemory[timerAlias] = &timerEntry
}

func GetTimer(timerAlias string) TimerEntryType {
	if TimerMemory[timerAlias] == nil {
		panic(fmt.Sprintf("The requested timer with alias '%s' could not be returned since it does not exist.", timerAlias))
	}
	return *TimerMemory[timerAlias]
}

func DeleteTimer(timerAlias string) {
	delete(TimerMemory, timerAlias)
}

func GetCurrentTimeInMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
