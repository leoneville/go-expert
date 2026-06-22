package events

import (
	"errors"
	"slices"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}

	return nil
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if ed.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		if slices.Contains(ed.handlers[eventName], handler) {
			return true
		}
	}
	return false
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	handlers, ok := ed.handlers[eventName]
	if !ok {
		return nil
	}

	for i, h := range handlers {
		if h == handler {
			ed.handlers[eventName] = slices.Delete(handlers, i, i+1)
			return nil

			// oldLen := len(handlers)

			// handlers = append(handlers[:i], handlers[i+1:]...)
			// clear(handlers[len(handlers):oldLen])

			// ed.handlers[eventName] = handlers
			// return nil
		}
	}
	return nil
}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}
