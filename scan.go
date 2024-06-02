package gosha

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
)

const defaultScannerLogPattern = "[%s]: Path: %s Message: %s"

type (
	IScanner interface {
		Scan(io.ReadCloser, *Cmd) error
	}

	DefaultScanner struct{}
)

func (s *DefaultScanner) Scan(stdout io.ReadCloser, cmd *Cmd) error {
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		msg := scanner.Text()
		slog.Info(fmt.Sprintf(defaultScannerLogPattern, cmd.Title, cmd.Path, msg))
	}
	return nil
}

func GetDefaultScanner() IScanner {
	return &DefaultScanner{}
}
