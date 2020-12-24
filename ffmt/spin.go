package ffmt

import (
	"fmt"
	"strings"
)

func SpinPrint(a string) {
	fmt.Printf("\r%s", strings.TrimSpace(a))
}
