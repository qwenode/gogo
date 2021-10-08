package main

import (
	"github.com/qwenode/gogo/cmdline"
	"log"
)

func main() {
	stdout := cmdline.CommandRealtimeStdout(func(output string, errCode int, done bool) bool {
		if done {
			log.Println(output, errCode, done)
			return true
		}
		log.Println(output)
		return true
	}, "ping", "127.0.0.1", "-c", "10")
	log.Println(stdout)
	realtimeStdout := cmdline.CommandRealtimeStdout(func(output string, errCode int, done bool) bool {
		if done {
			log.Println(output, errCode, done)
			return true
		}
		log.Println(output)
		return true
	}, "ping", "127.0.0.1", "-c", "10")
	log.Println(realtimeStdout)
}
