package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/theckman/yacspin"
)

func createSpinner() (*yacspin.Spinner, error) {
	// build the configuration, each field is documented
	cfg := yacspin.Config{
		Frequency: 100 * time.Millisecond,
		CharSet:   yacspin.CharSets[11],
		Suffix:    " ", // puts a least one space between the animating spinner and the Message
		// Message:           "collecting files",
		// SuffixAutoColon: true,
		// ColorAll:        true,
		Colors: []string{"fgYellow"},
		// StopCharacter: " ",
		StopColors: []string{"fgGreen"},
		// StopMessage:       "done",
		StopFailCharacter: "✗",
		StopFailColors:    []string{"fgRed"},
		StopFailMessage:   "failed",
	}

	s, err := yacspin.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to make spinner from struct: %w", err)
	}
	if err := s.Start(); err != nil {
		return nil, err
	}
	stopSpinnerOnSignal(s)

	return s, nil
}

func stopSpinnerOnSignal(spinner *yacspin.Spinner) {
	// ensure we stop the spinner before exiting, otherwise cursor will remain
	// hidden and terminal will require a `reset`
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh

		spinner.StopFailMessage("interrupted")

		// ignoring error intentionally
		_ = spinner.StopFail()

		os.Exit(0)
	}()
}

func cleanupSpinner(s *yacspin.Spinner) {
	s.StopMessage("done")
	s.StopColors("fgGreen")
	if err := s.Stop(); err != nil {
		s.StopFail()
	}
}
