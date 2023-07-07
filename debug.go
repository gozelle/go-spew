package spew

import (
	"fmt"
	"runtime"
)

func runFuncPos() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}
