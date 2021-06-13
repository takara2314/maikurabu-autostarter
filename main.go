package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// tmux起動
	cmd := exec.Command("/bin/sh", "-c", "tmux new -t server")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// tmux画面分割
	tmuxSetup := `
		tmux send-keys -t server.0 "tmux split-window -v -p 50" ENTER;
		tmux send-keys -t server.0 "tmux selectp -t 1" ENTER;
		tmux send-keys -t server-0 "tmux split-window -h -p 50" ENTER;
	`

	<-time.After(1 * time.Second)

	cmd = exec.Command("/bin/sh", "-c", tmuxSetup)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
