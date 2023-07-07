package spew

import (
	"fmt"
	"github.com/gozelle/color"
	"reflect"
)

func Type(a ...any) {
	color.Yellow("%s spew type: \n", runFuncPos())
	for _, v := range a {
		fmt.Println(reflect.TypeOf(v))
	}
}
