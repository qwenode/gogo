package system

import (
	"runtime"
	"strconv"
	"strings"
)

// GoroutineID get Goroutine id
// But it should be noted that get Stack information will affect performance,
// so it is recommended that you only use it when debugging.
func GoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.Atoi(idField)
	//if err != nil {
	//	panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	//}
	return id
}
