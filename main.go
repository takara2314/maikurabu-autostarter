package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "tmux new -t server")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
