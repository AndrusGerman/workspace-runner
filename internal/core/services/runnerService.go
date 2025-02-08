package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
)

func NewRunnerService(logger ports.RunnerLogger) ports.RunnerService {
	return &RunnerService{
		logger: logger,
	}
}

type RunnerService struct {
	logger ports.RunnerLogger
}

func (s *RunnerService) Run(ctx context.Context, workspace *models.Workspace, projects []*models.Project) error {
	fmt.Println("Workspace: ", workspace.Name)

	var waitClose = new(sync.WaitGroup)
	waitClose.Add(1)

	for _, project := range projects {
		go func(project *models.Project) {
			err := s.runProject(ctx, project, waitClose)
			if err != nil {
				fmt.Println("Error: ", err)
				waitClose.Done()
			}
		}(project)
	}

	waitClose.Wait()

	return nil
}

func (s *RunnerService) runProject(_ context.Context, project *models.Project, waitClose *sync.WaitGroup) error {
	fmt.Println("Project: ", project.Name)

	cmd := exec.Command(project.Cmd.Command, project.Cmd.Args...)
	cmd.Dir = project.WorkDirectory

	cmd.Stdout = s.logger.GetStdout(project.Name)
	cmd.Stderr = s.logger.GetStderr(project.Name)

	cmd.Env = os.Environ()
	for _, env := range project.Cmd.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", env.Key, env.Value))
	}

	go func() {
		waitClose.Wait()

		log.Println("Killing project: ", project.Name)
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}()

	return cmd.Run()

}
