package spew

import (
	"encoding/json"
	"fmt"
	
	"github.com/gozelle/color"
)

func Json(a ...any) {
	color.Yellow("%s spew json: \n", runFuncPos())
	for _, v := range a {
		d, _ := json.MarshalIndent(v, "", "\t")
		fmt.Println(string(d))
	}
}

func PrintJson(a ...any) {
	color.Yellow("%s spew json: \n", runFuncPos())
	for _, v := range a {
		d, _ := json.Marshal(v)
		fmt.Println(string(d))
	}
}
