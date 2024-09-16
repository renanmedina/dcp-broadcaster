package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
)

func newEventPublisher() *event_store.EventPublisher {
	return event_store.NewEventPublisherWith(
		configEventHandlers(),
	)
}

func configEventHandlers() map[string][]event_store.EventHandler {
	return map[string][]event_store.EventHandler{
		QUESTION_CREATED_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
			NewSendQuestionToUsersHandler(),
		},
		QUESTION_SENT_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
		},
	}
}
