package event_store

type SaveEventToStoreHandler struct {
	allEvents *EventsRepository
}

func NewSaveEventToStoreHandler() *SaveEventToStoreHandler {
	return &SaveEventToStoreHandler{
		NewEventsRepository(),
	}
}

func (h *SaveEventToStoreHandler) Handle(event PublishableEvent) {
	h.allEvents.Save(event)
}
