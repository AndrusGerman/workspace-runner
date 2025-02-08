package services

import (
	"bufio"
	"fmt"
	"io"

	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
)

func NewRunnerLogger() ports.RunnerLogger {
	return &RunnerLogger{}
}

type RunnerLogger struct{}

func (s *RunnerLogger) GetStdout(name string) io.Writer {
	return s.GetPipeLogger(name)
}

func (s *RunnerLogger) GetStderr(name string) io.Writer {
	return s.GetPipeLogger(name)
}

func (s *RunnerLogger) GetPipeLogger(name string) io.Writer {

	var reader, writer = io.Pipe()

	go func() {
		var buffer = bufio.NewReader(reader)
		for {
			var line, err = buffer.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Printf("%s: %s", name, line)
		}
	}()
	return writer
}
