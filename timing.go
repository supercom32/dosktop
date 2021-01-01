package dosktop

import (
	"github.com/supercom32/dosktop/internal/memory"
	"time"
)

/*
Sleep allows you to pause execution for a given amount of milliseconds.
This method is simply a convenient wrapper for the method
'SleepInMilliseconds'.
*/
func Sleep(timeInMilliseconds uint) {
	SleepInMilliseconds(timeInMilliseconds)
}

/*
SleepInSeconds allows you to pause execution for a given amount of seconds.
*/
func SleepInSeconds(timeInSeconds uint) {
	SleepInMilliseconds(timeInSeconds * 1000)
}

/*
SleepInMilliseconds allows you to pause execution for a given amount of
milliseconds.
*/
func SleepInMilliseconds(timeInMilliseconds uint) {
	timeDuration := time.Duration(timeInMilliseconds)
	time.Sleep(timeDuration * time.Millisecond)
}

/*
IsTimerExpired allows you to check if a created timer has expired or not.
If the specified timer has expired, then it will automatically be disabled.
In order to activate the timer again, simply call 'StartTimer'.
*/
func IsTimerExpired(timerAlias string) bool{
	timerEntry := memory.TimerMemory[timerAlias]
	if timerEntry.IsTimerEnabled {
		timeElapsedInMilliseconds := GetCurrentTimeInMilliseconds() - timerEntry.StartTime
		if timeElapsedInMilliseconds > timerEntry.TimerLength {
			timerEntry.IsTimerEnabled = false
			return true
		}
	}
	return false
}

/*
SetTimer allows you to create a new timer to measure time with. If the timer
is not enabled by default, you must call 'StartTimer' when you wish for it
to begin.
*/
func SetTimer(timerAlias string, durationInMilliseconds int64, isEnabled bool) {
	timerEntry := memory.TimerMemory[timerAlias]
	timerEntry.StartTime = GetCurrentTimeInMilliseconds()
	timerEntry.TimerLength = durationInMilliseconds
	timerEntry.IsTimerEnabled = isEnabled
}

/*
StartTimer allows you to start a timer that has already been previously
created. In addition, the following information should be noted:

- If you specify a timer that does not exist, then a panic will be
generated to fail as fast as possible.
*/
func StartTimer(timerAlias string) {
	timerEntry := memory.TimerMemory[timerAlias]
	timerEntry.StartTime = GetCurrentTimeInMilliseconds()
	timerEntry.IsTimerEnabled = true
}

/*
GetCurrentTimeInMilliseconds allows you to get the current epoch
time in milliseconds.
*/
func GetCurrentTimeInMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}