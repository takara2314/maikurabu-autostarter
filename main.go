package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// ベースソフト起動待ち
	time.Sleep(5 * time.Second)
	// tmux起動
	cmd := exec.Command("/bin/sh", "-c", "tmux new -s server -d")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func(c *exec.Cmd) {
		err := c.Run()
		if err != nil {
			fmt.Println(err)
		}
	}(cmd)

	// tmux画面分割
	tmuxSetup := `
		tmux send-keys -t server "cd ~/maikurabu" ENTER;
		tmux send-keys -t server "tmux split-window -h -p 50" ENTER;
		tmux send-keys -t server "tmux selectp -t 1" ENTER;
		tmux send-keys -t server "cd ~/maikurabu" ENTER;
		tmux send-keys -t server "ls" ENTER;
		tmux send-keys -t server "tmux split-window -v -p 50" ENTER;
		tmux send-keys -t server "tmux selectp -t 2" ENTER;
		tmux send-keys -t server "top" ENTER;
		tmux send-keys -t server "tmux selectp -t 0" ENTER;
		tmux send-keys -t server "start" ENTER;
	`

	<-time.After(1 * time.Second)

	tmuxCommands := strings.Split(tmuxSetup, "\n")

	for _, command := range tmuxCommands {
		cmd2 := exec.Command("/bin/sh", "-c", command)
		err := cmd2.Run()
		if err != nil {
			log.Println(err)
			panic(err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// シグナル設定 (Ctrl + C などのコマンドを押下したときのシグナルを受け取る)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	for {
		<-c
		fmt.Println("killing program")
		break
	}
}
