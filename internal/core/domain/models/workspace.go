package models

type Workspace struct {
	*Base       `bson:",inline" json:",inline"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

func NewWorkspace(name string, description string) *Workspace {
	return &Workspace{
		Base:        NewBase(),
		Name:        name,
		Description: description,
	}
}
