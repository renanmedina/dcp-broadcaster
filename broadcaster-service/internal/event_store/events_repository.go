package event_store

import (
	"encoding/json"
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/utils"
	"gorm.io/gorm"
)

const TABLE_NAME = "events"

type EventsRepository struct {
	db *gorm.DB
}

func NewEventsRepository() *EventsRepository {
	return &EventsRepository{
		db: utils.GetDatabaseConnection(),
	}
}

func (r *EventsRepository) Save(event PublishableEvent) error {
	eventData, err := json.Marshal(event.EventData())

	if err != nil {
		fmt.Println("Failed marshal event data")
		return err
	}

	dbValues := map[string]interface{}{
		"event_name":  event.Name(),
		"object_id":   event.ObjectId(),
		"object_type": event.ObjectType(),
		"event_data":  eventData,
	}

	result := r.db.Table(TABLE_NAME).Create(dbValues)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
