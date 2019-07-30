package sortutil

import (
	"fmt"
	"reflect"
)

type Order uint

const (
	ASC Order = iota + 1
	DESC
)

type Sort struct {
	slice interface{}
}

func New(slice interface{}) *Sort {
	rv := reflect.ValueOf(slice)
	t := rv.Type()
	fmt.Printf("type: %v\n", t.Kind())

	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		panic(fmt.Sprintf("disable type: %v", t.Kind()))
	}

	return &Sort{slice}
}

func (s *Sort) Order(name string, order Order) (interface{}, error) {
	rv := reflect.ValueOf(s.slice)
	n := rv.Len()
	t := rv.Type()

	fmt.Printf("len(n): %d\n", n)
	fmt.Printf("rv: %+v\n", rv)

	fmt.Printf("rv.Type(): %v", rv.Type())

	sf, exists := t.FieldByName(name)
	fmt.Printf("sf: %v", sf)
	fmt.Printf("exists: %v\n", exists)

	return nil, nil
}
