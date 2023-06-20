package spew

import (
	"fmt"
	"github.com/gozelle/color"
	"reflect"
)

func Type(a interface{}) {
	color.Yellow("%s spew type: \n", runFuncPos())
	fmt.Println(reflect.TypeOf(a))
}
