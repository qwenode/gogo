package cmd

import (
	"bufio"
	"github.com/qwenode/gogo/sanitize"
	"io/ioutil"
	"os/exec"
	"strings"
)

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

// CommandFn run exec.command(name,arg...).CombinedOutput(), just one command
func CommandLineFn(fn commandFunc, c string) bool {
	if len(c) == 0 {
		return false
	}
	split := strings.Split(c, " ")
	args := split[1:]
	return CommandFn(fn, split[0], args...)
}

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

// CommandRealtimeStdout run exec.command(name,arg...) and realtime output ,just one command
func CommandLineRealtimeStdout(fn commandStdoutFunc, c string) bool {
	if len(c) == 0 {
		return false
	}
	split := strings.Split(c, " ")
	args := split[1:]
	return CommandRealtimeStdout(fn, split[0], args...)
}

//TODO add CommandRealtimeStdoutTimeout run with timeout
//TODO add CommandFnTimeout run with timeout
