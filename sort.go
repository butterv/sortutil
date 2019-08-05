package sortutil

import (
	"fmt"
	"reflect"
	"sort"
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
	kind := rv.Type().Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		panic(fmt.Sprintf("disable type: %v", kind))
	}

	return &Sort{slice: slice}
}

func (s *Sort) Order(name string, order Order) {
	rv := reflect.ValueOf(s.slice)
	t := rv.Index(0).FieldByName(name).Type()

	var sortFunc func(i, j int) bool

	switch t.Kind() {
	// case reflect.Bool:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if order == ASC {
			sortFunc = func(i, j int) bool { return rv.Index(i).FieldByName(name).Int() < rv.Index(j).FieldByName(name).Int() }
		} else {
			sortFunc = func(i, j int) bool { return rv.Index(i).FieldByName(name).Int() > rv.Index(j).FieldByName(name).Int() }
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if order == ASC {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Uint() < rv.Index(j).FieldByName(name).Uint()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Uint() > rv.Index(j).FieldByName(name).Uint()
			}
		}
	case reflect.Float32, reflect.Float64:
		if order == ASC {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Float() < rv.Index(j).FieldByName(name).Float()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Float() > rv.Index(j).FieldByName(name).Float()
			}
		}
	case reflect.String:
		if order == ASC {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).String() < rv.Index(j).FieldByName(name).String()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).String() > rv.Index(j).FieldByName(name).String()
			}
		}
	default:
		panic(fmt.Sprintf("unsupported type: %s", rv.Index(0).Type().Kind().String()))
	}

	sort.Slice(s.slice, sortFunc)
}

// method chain
func (s *Sort) Asc() {

}

// method chain
func (s *Sort) Desc() {

}
