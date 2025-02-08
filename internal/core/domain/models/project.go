package models

import (
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Project struct {
	*Base         `bson:",inline" json:",inline"`
	Name          string        `bson:"name" json:"name"`
	WorkspaceId   bson.ObjectID `bson:"workspace_id" json:"workspace_id"`
	WorkDirectory string        `bson:"work_directory" json:"work_directory"`
	Cmd           *Cmd          `bson:"cmd" json:"cmd"`
}

func NewProject(name string, workspaceId types.Id, workDirectory string, cmd *Cmd) *Project {
	return &Project{
		Base:          NewBase(),
		Name:          name,
		WorkspaceId:   bson.ObjectID(workspaceId),
		WorkDirectory: workDirectory,
		Cmd:           cmd,
	}
}

type Cmd struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func NewCmd(command string, args []string) *Cmd {
	return &Cmd{
		Command: command,
		Args:    args,
	}
}
