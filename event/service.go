package event

import (
  "reflect"
)

import (
  "gopoker/message"
)

type Service struct {
  instanceMethods map[string]reflect.Method
}

type ServiceRequest struct {
  Method  string
  Type string
  Message *message.Message
}

type ServiceResponse struct {
  Error error
  Type string
  Message *message.Message
}

func NewReceiver(recv interface{}) *Service {
  instanceType := reflect.TypeOf(recv)
  msgType := reflect.TypeOf((*message.Message)(nil)).Elem()

  methods := map[string]reflect.Method{}
  for i := 0; i < instanceType.NumMethod(); i++ {
    method := instanceType.Method(i)

    methodType := method.Type
    // in/outs
    if methodType.NumIn() != 2 || methodType.NumOut() != 1 {
      continue
    }

    // one arg
    arg := methodType.In(1)
    if !arg.Implements(msgType) {
      continue
    }

    // return
    ret := methodType.Out(0)
    if !ret.Implements(msgType) {
      continue
    }

    methods[method.Name] = method
  }

  return &Service{
    instanceMethods: methods,
  }
}

func (service *Service) Send(req *ServiceRequest) {
  method, ok := service.instanceMethods[req.Method]
  
  if !ok {
    return
  }
  
  go method.Func.Call([]reflect.Value{
    reflect.ValueOf(req.Message),
  })
}
