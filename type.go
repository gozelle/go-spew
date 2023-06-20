package spew

import "reflect"

func Type(a interface{}) string {
	return reflect.TypeOf(a).String()
}
