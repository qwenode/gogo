package cmdline

import (
	"github.com/qwenode/gogo/sanitize"
	"os/exec"
)

//windows command: cmdline.exe /c "command"
//linux command: /bin/bash -c "command"
//TODO add CommandRealtimeStdoutTimeout run with timeout
//TODO add CommandFnTimeout run with timeout
// commandFunc call by CommandFn
type commandFunc func(output string, errCode int) bool

// CommandFn run exec.command(name,arg...).CombinedOutput()
func CommandFn(fn commandFunc, name string, arg ...string) bool {
	output, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		return fn(string(output), sanitize.Int(err.Error()))
	}
	return fn(string(output), 0)
}
