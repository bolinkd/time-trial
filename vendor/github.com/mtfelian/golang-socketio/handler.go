package gosocketio

import (
	"errors"
	"reflect"
)

// handler is an event handler representation
type handler struct {
	function reflect.Value
	args     reflect.Type
	hasArgs  bool
	out      bool
}

var (
	ErrorHandlerIsNotFunc   = errors.New("f is not a function")
	ErrorHandlerHasNot2Args = errors.New("f should have 1 or 2 arguments")
	ErrorHandlerWrongResult = errors.New("f should return no more than one value")
)

// newHandler parses function f (event handler) using reflection, and stores it's representation
func newHandler(f interface{}) (*handler, error) {
	fVal := reflect.ValueOf(f)
	if fVal.Kind() != reflect.Func {
		return nil, ErrorHandlerIsNotFunc
	}

	fType := fVal.Type()
	if fType.NumOut() > 1 {
		return nil, ErrorHandlerWrongResult
	}

	curCaller := &handler{
		function: fVal,
		out:      fType.NumOut() == 1,
	}

	switch fType.NumIn() {
	case 1:
		curCaller.args = nil
		curCaller.hasArgs = false
	case 2:
		curCaller.args = fType.In(1)
		curCaller.hasArgs = true
	default:
		return nil, ErrorHandlerHasNot2Args
	}

	return curCaller, nil
}

// arguments returns function parameter as it is present in it using reflection
func (h *handler) arguments() interface{} { return reflect.New(h.args).Interface() }

// call func with given arguments from its representation using reflection
func (h *handler) call(c *Channel, arguments interface{}) []reflect.Value {
	// nil is untyped, so use the default empty value of correct type
	if arguments == nil {
		arguments = h.arguments()
	}

	a := []reflect.Value{reflect.ValueOf(c), reflect.ValueOf(arguments).Elem()}
	if !h.hasArgs {
		a = a[0:1]
	}

	return h.function.Call(a)
}
