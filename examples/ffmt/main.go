package main

import (
	"fmt"
	"github.com/qwenode/gogo/ffmt"
	"time"
)

func main() {
	for i := 0; i <= 30; i++ {
		ffmt.SpinPrint(fmt.Sprintf("Process:%d/30", i))
		time.Sleep(time.Millisecond * 100)
	}
}
