package common

import (
	"errors"
	"fmt"
	"reflect"
)

// InvocationHandler .
type InvocationHandler interface {
	Invoke(proxy *Proxy, method *Method, args []any) ([]any, error)
}

// Proxy .
type Proxy struct {
	instance          any
	methods           map[string]*Method
	invocationHandler InvocationHandler
}

// NewProxy .
func NewProxy(instance any, invocationHandler InvocationHandler) *Proxy {
	typ := reflect.TypeOf(instance)
	value := reflect.ValueOf(instance)
	methods := make(map[string]*Method)

	for i := 0; i < value.NumMethod(); i++ {
		method := value.Method(i)
		methods[typ.Method(i).Name] = &Method{value: method}
	}
	return &Proxy{
		instance:          instance,
		methods:           methods,
		invocationHandler: invocationHandler,
	}
}

// InvokeMethod .
func (this *Proxy) InvokeMethod(name string, args ...any) ([]any, error) {
	return this.invocationHandler.Invoke(this, this.methods[name], args)
}

// Method .
type Method struct {
	value reflect.Value
}

// Invoke .
func (this *Method) Invoke(args ...any) (results []any, err error) {
	defer func() {
		// throw exception
		if p := recover(); p != nil {
			err = errors.New(fmt.Sprintf("%s", p))
		}
	}()

	// parameter
	params := make([]reflect.Value, 0)
	if args != nil {
		for i := 0; i < len(args); i++ {
			params = append(params, reflect.ValueOf(args[i]))
		}
	}

	// execute method
	callResults := this.value.Call(params)

	// results
	results = make([]any, 0)
	if callResults != nil && len(callResults) > 0 {
		for i := 0; i < len(callResults); i++ {
			results = append(results, callResults[i].Interface())
		}
	}
	return
}
