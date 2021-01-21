package ffmt

import (
	"fmt"
	"strings"
)

// SpinPrint spin print for linux
func SpinPrint(a string) {
	fmt.Printf("\r%s", strings.TrimSpace(a))
}
