package main

import (
	"log"
	"os/exec"
)

func main() {
	output, err := exec.Command("tailscale", "set", "--exit-node=").CombinedOutput()

	log.Println(output, err)
}
