package spew

import (
	"fmt"
	"runtime"
)

func runFuncPos() string {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(f.Entry())
	return fmt.Sprintf("%s:%d", file, line)
}
