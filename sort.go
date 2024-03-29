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

const (
	// MaxConditions is maximum number of conditions.
	MaxConditions int = 5
)

// Sort struct has list of targets.
type Sort struct {
	slice          interface{}
	value          reflect.Value
	sortConditions []*sortCondition
}

type sortCondition struct {
	fieldName string
	orderType OrderType
}

// Order is initialize Sort.
// `slice` type should be Slice or Array.
func Order(slice interface{}) *Sort {
	rv := reflect.ValueOf(slice)
	kind := rv.Type().Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		panic(fmt.Sprintf("disable type: %v", kind))
	}

	return &Sort{
		slice: slice,
		value: rv,
	}
}

// Asc in ascending order in any field.
func (s *Sort) Asc(name string) *Sort {
	if s.sorted(name) {
		fmt.Printf("%s already selected", name)
		return s
	}
	if len(s.sortConditions) > MaxConditions {
		fmt.Printf("No more can be sorted: by %s", name)
		return s
	}

	s.sortConditions = append(s.sortConditions, &sortCondition{
		fieldName: name,
		orderType: ASC,
	})

	return s
}

// Desc in descending order in any field.
func (s *Sort) Desc(name string) *Sort {
	if s.sorted(name) {
		fmt.Printf("%s already selected", name)
		return s
	}
	if len(s.sortConditions) > MaxConditions {
		fmt.Printf("No more can be sorted: by %s", name)
		return s
	}

	s.sortConditions = append(s.sortConditions, &sortCondition{
		fieldName: name,
		orderType: DESC,
	})

	return s
}

func (s *Sort) sorted(name string) bool {
	for _, sortCondition := range s.sortConditions {
		if sortCondition.fieldName == name {
			return true
		}
	}
	return false
}

// Exec performs a sort
func (s *Sort) Exec() {
	sort.SliceStable(s.slice, s.makeSortFuncs(len(s.sortConditions)))
}

func (s *Sort) makeSortFuncs(count int) func(i, j int) bool {
	var funcs []func(i, j int) bool
	for i := 0; i < count; i++ {
		funcs = append(funcs, s.sort(i))
	}

	return func(i, j int) bool {
		for _, ff := range funcs[:count-1] {
			if ff(i, j) {
				return true
			}
			if ff(j, i) {
				return false
			}
		}

		return funcs[count-1](i, j)
	}
}

func (s *Sort) sort(index int) func(i, j int) bool {
	sortCondition := s.sortConditions[index]
	t1 := s.value.Index(0).FieldByName(sortCondition.fieldName).Type()

	var sortFunc func(i, j int) bool

	switch t1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if sortCondition.orderType == ASC {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).Int() < s.value.Index(j).FieldByName(sortCondition.fieldName).Int()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).Int() > s.value.Index(j).FieldByName(sortCondition.fieldName).Int()
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if sortCondition.orderType == ASC {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).Uint() < s.value.Index(j).FieldByName(sortCondition.fieldName).Uint()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).Uint() > s.value.Index(j).FieldByName(sortCondition.fieldName).Uint()
			}
		}
	case reflect.Float32, reflect.Float64:
		if sortCondition.orderType == ASC {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).Float() < s.value.Index(j).FieldByName(sortCondition.fieldName).Float()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).Float() > s.value.Index(j).FieldByName(sortCondition.fieldName).Float()
			}
		}
	case reflect.String:
		if sortCondition.orderType == ASC {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).String() < s.value.Index(j).FieldByName(sortCondition.fieldName).String()
			}
		} else {
			sortFunc = func(i, j int) bool {
				return s.value.Index(i).FieldByName(sortCondition.fieldName).String() > s.value.Index(j).FieldByName(sortCondition.fieldName).String()
			}
		}
	default:
		panic(fmt.Sprintf("unsupported type: %s", s.value.Index(0).Type().Kind().String()))
	}

	return sortFunc
}
