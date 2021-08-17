package data

import (
	"meal_api/xer"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Event struct {
	ID       uint   `gorm:"primaryKey"`
	TeamUUID string `gorm:"size:36"`
	Team     Team
	Sort     int
	Name     string
}

type EventList []Event

func (events EventList) Len() int           { return len(events) }
func (events EventList) Swap(i, j int)      { events[i], events[j] = events[j], events[i] }
func (events EventList) Less(i, j int) bool { return events[i].Sort < events[j].Sort }

func FetchEventById(id int) (event Event, err error) {
	Result := Db.First(&event, id)
	if errors.Is(Result.Error, gorm.ErrRecordNotFound) {
		err = xer.Err4xx{ErrType: xer.EventNotFound}
	}
	return event, errors.WithStack(err)
}

func (event Event) setIDandName2Json(eventResponses *EventResponseJson) {
	eventResponses.ID = event.ID
	eventResponses.Name = event.Name
}
