package cmdline

import (
	"bufio"
	"github.com/qwenode/gogo/sanitize"
	"io/ioutil"
	"os/exec"
)

// commandStdoutFunc call by CommandRealtimeStdout
type commandStdoutFunc func(output string, errCode int, done bool) bool

// CommandRealtimeStdout run exec.command(name,arg...) and realtime output
func CommandRealtimeStdout(fn commandStdoutFunc, name string, arg ...string) bool {
	command := exec.Command(name, arg...)
	pipe, _ := command.StdoutPipe()
	stderrPipe, _ := command.StderrPipe()
	defer pipe.Close()
	defer stderrPipe.Close()
	if err := command.Start(); err != nil {
		all, _ := ioutil.ReadAll(stderrPipe)
		return fn(string(all), sanitize.Int(string(all)), true)
	}
	reader := bufio.NewReader(pipe)
	r := true
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			all, _ := ioutil.ReadAll(stderrPipe)
			r = fn(string(all), sanitize.Int(string(all)), true)
			break
		}
		r = fn(readString, 0, false)
		if r == false {
			break
		}
	}
	return r
}

// CommandLineRealtimeStdout run exec.command(name,arg...) and realtime output ,just one command,actually exec /bin/bash -c "your command"
func CommandLineRealtimeStdout(fn commandStdoutFunc, c string) bool {
	return CommandRealtimeStdout(fn, "/bin/bash", "-c", c)
}

// CommandLineFn run exec.command(name,arg...).CombinedOutput(), just one command,actually exec /bin/bash -c "your command"
func CommandLineFn(fn commandFunc, c string) bool {
	return CommandFn(fn, "/bin/bash", "-c", c)
}
