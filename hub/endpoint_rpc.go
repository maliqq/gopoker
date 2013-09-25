package hub

import (
	"reflect"
)

import (
	"gopoker/event"
	"gopoker/message"
)

type Receiver struct {
	Instance        *event.Service
	instanceType    reflect.Type
	instanceMethods map[string]reflect.Method
}

func NewReceiver(recv *event.Service) *Receiver {
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

	return &Receiver{
		Instance:        recv,
		instanceType:    instanceType,
		instanceMethods: methods,
	}
}

func (r *Receiver) Send(message interface{}) {
	req, ok := message.(*event.Call)

	if !ok {
		return
	}

	method, found := r.instanceMethods[req.Method]
	if !found {
		return
	}

	go method.Func.Call([]reflect.Value{
		reflect.ValueOf(req.Message),
	})
}
