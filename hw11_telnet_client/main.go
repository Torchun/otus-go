package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var timeout time.Duration

// Optional flag, will be executed on import, can be overwritten by "--timeout=5s" flag.
func init() {
	flag.DurationVar(&timeout, "timeout", 10, "default timeout duration")
}

func main() {
	flag.Parse()
	// ["go-telnet", "--timeout=5s", "localhost", "4242"]
	port := os.Args[len(os.Args)-1]
	host := os.Args[len(os.Args)-2]

	in := &bytes.Buffer{}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

	client := NewTelnetClient(host+":"+port, timeout, io.NopCloser(in), os.Stdout)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		fmt.Printf("connection failed: %s\n", err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go tcpsend(ctx, stop, &wg, in, client)
	go conncontrol(ctx, &wg, client)

	wg.Wait()
}

func tcpsend(ctx context.Context, stop context.CancelFunc, wg *sync.WaitGroup, in *bytes.Buffer, t TelnetClient) {
	// wait until context stopped
	go func() {
		<-ctx.Done()
		// fmt.Println("ctx.Done at tcpsend")
		// close connection to server manually
		t.Close()
		wg.Done()
	}()

	reader := bufio.NewReader(os.Stdin)
	// infinite read from stdin and pass to server line by line
	for {
		resp, err := reader.ReadString('\n')
		// fmt.Println(resp, err)
		if err != nil {
			// fmt.Println(err)
			// pass signal to close ctx
			stop()
			return
		}

		in.WriteString(resp)
		err = t.Send()
		if err != nil {
			fmt.Println(err)
			// stop()
			return
		}
	}
}

func conncontrol(ctx context.Context, wg *sync.WaitGroup, t TelnetClient) {
	// to perform select case and wait for interrupt/error
	errCh := make(chan error)
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			// fmt.Println("ctx.Done() poped")
			// main interrupted by syscall (Ctrl+C)
			// deferred wg.Done() executed (wg delta -1)
			return
		case errCh <- t.Receive():
			err := <-errCh
			if err != nil {
				fmt.Println(err)
				// io.Copy completed due to EOF/Error
				// deferred wg.Done() executed (wg delta -1)
				return
			}
		}
	}
}
