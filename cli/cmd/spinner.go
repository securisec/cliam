package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/theckman/yacspin"
)

func getSpinner(ctx context.Context, count chan int, total int) {
	cfg := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		Colors:            []string{"fgYellow"},
		CharSet:           yacspin.CharSets[11],
		Suffix:            " ",
		SuffixAutoColon:   true,
		StopCharacter:     "✓",
		StopColors:        []string{"fgGreen"},
		StopMessage:       "done",
		StopFailCharacter: "✗",
		StopFailColors:    []string{"fgRed"},
		StopFailMessage:   "failed",
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		exitf("failed to make spinner from struct: %v", err)
	}
	// handle spinner cleanup on interrupts
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	defer signal.Stop(sigCh)

	go func(ctx context.Context) { // this is just an example signal handler, should be more robust
		select {
		case <-ctx.Done():
			spinner.StopMessage("done")
		case <-sigCh:
			spinner.StopFailMessage("interrupted")
		}

		os.Exit(0)
	}(ctx)
}

func exitf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
