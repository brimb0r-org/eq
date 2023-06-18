package eq_math

import (
	"fmt"
	"reflect"
)

type iOperander interface {
	val() float64
}

func Sub(operanders ...interface{}) float64 {
	if len(operanders) < 1 {
		return 0
	}

	result := val(operanders[0])

	for _, v := range operanders[1:] {

		result -= val(v)
	}

	return result
}

func Add(operanders ...interface{}) float64 {
	if len(operanders) < 1 {
		return 0
	}

	result := val(operanders[0])

	for _, v := range operanders[1:] {

		result += val(v)
	}

	return result
}

func val(operander interface{}) float64 {
	if x, ok := operander.(iOperander); ok {
		return x.val()
	}

	x := reflect.ValueOf(operander)

	// nolint:exhaustive
	switch x.Kind() {
	case reflect.Bool:
		if x.Bool() {
			return 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return float64(x.Uint())
	case reflect.Float32, reflect.Float64:
		return x.Float()
	case reflect.Complex64, reflect.Complex128:
		y := x.Complex()
		return real(y)
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		fmt.Println("im array")
		return float64(x.Len())
	case reflect.Struct:
		return float64(x.NumField())
	}

	return 0
}
