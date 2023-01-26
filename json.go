package spew

import (
	"encoding/json"
	"fmt"
)

func Json(a any) {
	d, _ := json.MarshalIndent(a, "", "\t")
	fmt.Println(string(d))
}
