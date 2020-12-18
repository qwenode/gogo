package cmd

import (
	"github.com/qwenode/gogo/sanitize"
	"os/exec"
)

// CommandFunc call by CommandFn
type CommandFunc func(output string, errCode int) bool

// CommandFn run exec.command(name,arg...).CombinedOutput()
func CommandFn(fn CommandFunc, name string, arg ...string) bool {
	output, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		return fn(string(output), sanitize.Int(err.Error()))
	}
	return fn(string(output), 0)
}
