package spew

import (
	"encoding/json"
	"fmt"

	"github.com/gozelle/color"
)

func Json(a any) {
	d, _ := json.MarshalIndent(a, "", "\t")
	color.Yellow("%s spew json: \n", runFuncPos())
	fmt.Println(string(d))
}
