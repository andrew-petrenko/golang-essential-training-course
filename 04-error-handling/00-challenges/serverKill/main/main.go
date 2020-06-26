package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)

	if err != nil {
		return errors.Wrap(err, "can't open pid file")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid file - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	fmt.Printf("killing server with pid=%d\n", pid)

	return nil
}
