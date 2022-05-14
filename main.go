package main

import (
	"fmt"
	"log"
	"time"
)

type TimeEpoch struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

type Event struct {
	Name string `json:"name"`
	*TimeEpoch
}

func (e *Event) IsValid() bool {
	return e.StartTime != 0 && e.EndTime != 0
}

func (t *TimeEpoch) GetStartTime() (unix time.Time) {
	return time.Unix(t.StartTime, 0).AddDate(0, 0, 1)
}

func (t *TimeEpoch) GetEndTime() (unix time.Time) {
	return time.Unix(t.EndTime, 0)
}

func (e *Event) IsExpired() bool {
	return e.EndTime < time.Now().Unix()
}

func (e *Event) IsOverlapping(other *Event) bool {
	return e.StartTime < other.EndTime && e.EndTime > other.StartTime
}

func (e *Event) IsCurrent() bool {
	return e.StartTime <= time.Now().Unix() && e.EndTime > time.Now().Unix()
}

func (e *Event) IsPast() bool {
	return e.EndTime < time.Now().Unix()
}

func (e *Event) IsFuture() bool {
	return e.StartTime > time.Now().Unix()
}

type Events []*Event

func (e Events) IsValid() bool {
	for i := 0; i < len(e); i++ {
		for j := i + 1; j < len(e); j++ {
			if e[i].IsOverlapping(e[j]) {
				log.Println("Overlapping events:", e[i], e[j])
				return false
			}
		}
	}
	return true
}

// if event has passed, schedule it for next year
func (e *Event) ScheduleForNextYear() {
	e.StartTime = time.Now().AddDate(1, 0, 0).Unix()
	e.EndTime = time.Now().AddDate(1, 0, 0).Unix()
}

func (e *Event) ScheduleForNextMonth() {
	e.StartTime = time.Now().AddDate(0, 1, 0).Unix()
	e.EndTime = time.Now().AddDate(0, 1, 0).Unix()
}

func (e *Event) ScheduleForNextWeek() {
	e.StartTime = time.Now().AddDate(0, 0, 7).Unix()
	e.EndTime = time.Now().AddDate(0, 0, 7).Unix()
}

func (e *Event) ScheduleForNextDay() {
	e.StartTime = time.Now().AddDate(0, 0, 1).Unix()
	e.EndTime = time.Now().AddDate(0, 0, 1).Unix()
}

func (e *Event) ScheduleForNextHour() {
	e.StartTime = time.Now().Add(time.Hour).Unix()
	e.EndTime = time.Now().Add(time.Hour).Unix()
}

func (e *Event) ScheduleForNextYearIfExpired() {
	if e.IsExpired() {
		e.ScheduleForNextYear()
	}
}

func (e *Event) String() string {
	return fmt.Sprintf("%s S: %s - E: %s", e.Name, e.GetStartTime().Format(time.RFC1123), e.GetEndTime().Format(time.RFC1123))
}

func main() {
	current := time.Now().Unix()
	end := time.Now().AddDate(0, 0, 0).Unix()
	event := &Event{Name: "Test", TimeEpoch: &TimeEpoch{current, end}}
	fmt.Println(event.String())
}
