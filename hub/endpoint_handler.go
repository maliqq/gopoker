package hub

import (
	"reflect"
	"strings"
)

import (
	"gopoker/event"
	"gopoker/message"
)

type Handler struct {
	Observer *event.Observer
	observerType reflect.Type
	handlers map[string]reflect.Method
}

const (
	handleMethodPrefix = "Handle"
)

func NewHandler(recv *event.Observer) *Handler {
	observerType := reflect.TypeOf(recv)
	msgType := reflect.TypeOf((*message.Message)(nil)).Elem()

	handlers := map[string]reflect.Method{}
	for i := 0; i < observerType.NumMethod(); i++ {
		method := observerType.Method(i)

		if strings.Index(method.Name, handleMethodPrefix) < 0 {
			continue
		}

		methodType := method.Type
		if methodType.NumIn() != 2 {
			continue
		}

		arg := methodType.In(1)
		if !arg.Implements(msgType) {
			continue
		}

		handlers[method.Name] = method
	}

	return &Handler{
		Observer:        recv,
		observerType:    observerType,
		handlers: handlers,
	}
}

func (h *Handler) Send(message interface{}) {
	n, ok := message.(*event.Notification)
	if !ok {
		return
	}
	
	methodName := handleMethodPrefix + n.Type

	handler, found := h.handlers[methodName]
	if !found {
		handler, found = h.handlers[handleMethodPrefix]
		if !found {
			return
		}
	}

	go handler.Func.Call([]reflect.Value{
		reflect.ValueOf(n),
	})
}
