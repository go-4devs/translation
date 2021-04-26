package arg

import "fmt"

// Arg defaults argument.
type Arg struct {
	Key   string
	Value interface{}
}

// String arg to string.
func (a Arg) String() string {
	return fmt.Sprintf("key:%s, value:%v", a.Key, a.Value)
}

// Val gets valuet argument.
func (a Arg) Val() interface{} {
	return a.Value
}
