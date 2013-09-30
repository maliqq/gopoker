package event

import (
  "strings"
  "reflect"
)

import (
  "gopoker/message"
)

type Observer struct {
  handlerMethods map[string]reflect.Method
}

const (
  handleMethodPrefix = "Handle"
)

func NewObserver(handler interface{}) *Observer {
  handlerType := reflect.TypeOf(handler)
  msgType := reflect.TypeOf((*message.Message)(nil)).Elem()

  methods := map[string]reflect.Method{}
  for i := 0; i < handlerType.NumMethod(); i++ {
    method := handlerType.Method(i)

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

    methods[method.Name] = method
  }

  return &Observer{
    handlerMethods: methods,
  }
}

func (observer *Observer) handlerFor(msgType string) (reflect.Method, bool) {
  methodName := handleMethodPrefix + msgType

  handler, ok := observer.handlerMethods[methodName]
  if !ok {
    handler, ok = observer.handlerMethods[handleMethodPrefix]
  }

  return handler, ok
}

func (observer *Observer) Send(message interface{}) {
  n, ok := message.(*Notification)
  if !ok {
    return
  }

  method, found := observer.handlerFor(n.Type)
  if !found {
    return
  }

  go method.Func.Call([]reflect.Value{
    reflect.ValueOf(n),
  })
}
