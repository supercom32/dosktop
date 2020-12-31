package memory

import (
	"encoding/json"
)

type TimerEntryType struct {
	IsTimerEnabled bool
	StartTime      int64
	TimerLength    int64
}

func (shared TimerEntryType) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		IsTimerEnabled bool
		StartTime int64
		TimerLength int64
	}{
		IsTimerEnabled: shared.IsTimerEnabled,
		StartTime: shared.StartTime,
		TimerLength: shared.TimerLength,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (shared TimerEntryType) GetEntryAsJsonDump() string {
	j, err := json.Marshal(shared)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewTimerEntry(existingTimerEntry ...*TimerEntryType) TimerEntryType {
	var timerEntry TimerEntryType
	if existingTimerEntry != nil {
		timerEntry.IsTimerEnabled = existingTimerEntry[0].IsTimerEnabled
		timerEntry.StartTime = existingTimerEntry[0].StartTime
		timerEntry.TimerLength = existingTimerEntry[0].TimerLength
	}
	return timerEntry
}
