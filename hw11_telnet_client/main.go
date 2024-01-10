package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const TIMEOUT = 10 * time.Second

var (
	ErrWrongArgs = errors.New("hostname and port required")
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	timeout := flag.Duration("timeout", TIMEOUT, "connection timeout")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		return ErrWrongArgs
	}

	address := net.JoinHostPort(args[0], args[1])
	telnet := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)
	if err := telnet.Connect(); err != nil {
		return err
	}
	defer telnet.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	go send(&telnet, cancel)
	go receive(&telnet, cancel)

	<-ctx.Done()

	err := telnet.Close()
	return err
}

func send(tc *TelnetClient, cancelFunc context.CancelFunc) {
	if err := (*tc).Send(); err != nil {
		_, _1 := fmt.Fprintf(os.Stderr, "%v\n", err)
		_ = _1
	} else {
		_, _1 := fmt.Fprintf(os.Stderr, "diconnect.\n")
		_ = _1
	}
	cancelFunc()
}

func receive(tc *TelnetClient, cancelFunc context.CancelFunc) {
	if err := (*tc).Receive(); err != nil {
		_, _1 := fmt.Fprintf(os.Stderr, "%v\n", err)
		_ = _1
	} else {
		_, _1 := fmt.Fprintf(os.Stderr, "Connection have been closed by peer.\n")
		_ = _1
	}
	cancelFunc()
}
