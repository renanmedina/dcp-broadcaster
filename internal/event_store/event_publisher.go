package event_store

import (
	"fmt"
	"reflect"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type PublishableEvent interface {
	Name() string
	ObjectId() string
	ObjectType() string
	EventData() map[string]interface{}
}

type EventHandler interface {
	Handle(event PublishableEvent)
}

type EventPublisher struct {
	handlers map[string][]EventHandler
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		handlers: make(map[string][]EventHandler),
	}
}

func NewEventPublisherWith(handlersSetup map[string][]EventHandler) *EventPublisher {
	return &EventPublisher{
		handlers: handlersSetup,
	}
}

func (p *EventPublisher) Publish(event PublishableEvent) bool {
	eventHandlers, exists := p.handlers[event.Name()]

	if exists {
		go (func(evt PublishableEvent, handlers []EventHandler) {
			for _, handler := range handlers {
				logMsg := fmt.Sprintf("Calling handler %s for event %s", reflect.TypeOf(handler), event.Name())
				utils.GetApplicationLogger().Info(logMsg)
				handler.Handle(event)
			}
		})(event, eventHandlers)

		return true
	}

	return false
}

func (p *EventPublisher) Subscribe(eventName string, handler *EventHandler) *EventPublisher {
	eventHandlers, exists := p.handlers[eventName]

	if !exists {
		eventHandlers = make([]EventHandler, 0)
	}

	eventHandlers = append(eventHandlers, *handler)
	p.handlers[eventName] = eventHandlers
	return p
}
