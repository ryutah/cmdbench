package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const usage = `Usage of cmdbench:
    cmdbench [COMMAND]
    e.g.:
      cmdbench "echo HELLO"`

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	procTime, err := bench(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Average: %v\n", procTime)
}

func createCommand(command string) *exec.Cmd {
	cmds := strings.Split(command, " ")
	cmd, args := cmds[0], cmds[1:]
	c := exec.Command(cmd, args...)
	return c
}

func bench(cmd string) (time.Duration, error) {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		cmd := createCommand(os.Args[1])
		stderr := new(bytes.Buffer)
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return -1, fmt.Errorf("%v, failed to exec command: %v", err, stderr.String())
		}
		stderr.Reset()
	}
	end := time.Now()

	return end.Sub(start) / 1000, nil
}

func printUsage() {
	fmt.Println(usage)
}
