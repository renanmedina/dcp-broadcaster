package event_store

import (
	"encoding/json"
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

const TABLE_NAME = "events"

type EventsRepository struct {
	db *utils.DatabaseAdapdater
}

func NewEventsRepository() *EventsRepository {
	return &EventsRepository{
		db: utils.GetDatabase(),
	}
}

func (r *EventsRepository) Save(event PublishableEvent) error {
	eventData, err := json.Marshal(event.EventData())

	if err != nil {
		fmt.Println("Failed marshal event data")
		return err
	}

	_, err = r.db.Insert(TABLE_NAME, map[string]interface{}{
		"event_name":  event.Name(),
		"object_id":   event.ObjectId(),
		"object_type": event.ObjectType(),
		"event_data":  eventData,
	})

	if err != nil {
		return err
	}

	return nil
}
