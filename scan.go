package gosha

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
)

type (
	IScanner interface {
		Scan(io.ReadCloser, Cmd) error
	}

	Scanner struct{}
)

func (s *Scanner) Scan(stdout io.ReadCloser, cmd Cmd) error {
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		msg := scanner.Text()
		slog.Info(fmt.Sprintf("[%s]: %s", cmd.Path, msg))
	}
	return nil
}

func GetScanner() IScanner {
	return &Scanner{}
}
