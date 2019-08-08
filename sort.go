package sortutil

import (
	"fmt"
	"reflect"
	"sort"
)

type OrderType uint

const (
	ASC OrderType = iota + 1
	DESC
)

type Sort struct {
	slice interface{}
}

func Order(slice interface{}) *Sort {
	rv := reflect.ValueOf(slice)
	kind := rv.Type().Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		panic(fmt.Sprintf("disable type: %v", kind))
	}

	return &Sort{slice: slice}
}

func (s *Sort) Asc(name string) {
	s.sort(name, ASC)
}

func (s *Sort) Desc(name string) {
	s.sort(name, DESC)
}

func (s *Sort) sort(name string, orderType OrderType) {
	rv := reflect.ValueOf(s.slice)
	t := rv.Index(0).FieldByName(name).Type()

	var sortFunc func(i, j int) bool

	switch t.Kind() {
	// case reflect.Bool:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if orderType == ASC {
			sortFunc = func(i, j int) bool { return rv.Index(i).FieldByName(name).Int() < rv.Index(j).FieldByName(name).Int() }
		} else {
			sortFunc = func(i, j int) bool { return rv.Index(i).FieldByName(name).Int() > rv.Index(j).FieldByName(name).Int() }
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if orderType == ASC {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Uint() < rv.Index(j).FieldByName(name).Uint()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Uint() > rv.Index(j).FieldByName(name).Uint()
			}
		}
	case reflect.Float32, reflect.Float64:
		if orderType == ASC {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Float() < rv.Index(j).FieldByName(name).Float()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return rv.Index(i).FieldByName(name).Float() > rv.Index(j).FieldByName(name).Float()
			}
		}
	case reflect.String:
		if orderType == ASC {
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
