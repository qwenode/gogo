package cmdline

// CommandLineRealtimeStdout run exec.command(name,arg...) and realtime output ,just one command,actually exec /bin/bash -c "your command"
func CommandLineRealtimeStdout(fn commandStdoutFunc, c string) bool {
	return CommandRealtimeStdout(fn, "/bin/bash", "-c", c)
}

// CommandLineFn run exec.command(name,arg...).CombinedOutput(), just one command,actually exec /bin/bash -c "your command"
func CommandLineFn(fn commandFunc, c string) bool {
	return CommandFn(fn, "/bin/bash", "-c", c)
}
