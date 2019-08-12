package sortutil

import (
	"fmt"
	"reflect"
	"sort"
)

// OrderType kind of sort, ASC or DESC.
type OrderType uint

const (
	// ASC in ascending order in any field.
	ASC OrderType = iota + 1
	// DESC in descending order in any field.
	DESC
)

// Sort struct has list of targets.
type Sort struct {
	slice            interface{}
	sortedFieldNames []string
}

// Order is initialize Sort.
// `slice` type should be Slice or Array.
func Order(slice interface{}) *Sort {
	rv := reflect.ValueOf(slice)
	kind := rv.Type().Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		panic(fmt.Sprintf("disable type: %v", kind))
	}

	return &Sort{slice: slice}
}

// Asc in ascending order in any field.
func (s *Sort) Asc(name string) *Sort {
	switch len(s.sortedFieldNames) {
	case 0:
		s.first(name, ASC)
	case 1:
		s.second(name, ASC)
	case 2:
		s.third(name, ASC)
	default:
		fmt.Printf("No more can be sorted: by %s", name)
		return s
	}

	return s
}

// Desc in descending order in any field.
func (s *Sort) Desc(name string) *Sort {
	s.first(name, DESC)
	return s
}

func (s *Sort) first(name string, orderType OrderType) {
	if s.sorted(name) {
		fmt.Printf("No more can be sorted: by %s", name)
		return
	}

	rv := reflect.ValueOf(s.slice)
	t := rv.Index(0).FieldByName(name).Type()

	// TODO Add sort func to list(max size 3)

	var sortFunc func(i, j int) bool

	switch t.Kind() {
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
	s.sortedFieldNames = append(s.sortedFieldNames, name)
}

func (s *Sort) second(name string, orderType OrderType) {

}

func (s *Sort) third(name string, orderType OrderType) {

}

func (s *Sort) sorted(name string) bool {
	for _, sortedFieldName := range s.sortedFieldNames {
		if sortedFieldName == name {
			return true
		}
	}
	return false
}
