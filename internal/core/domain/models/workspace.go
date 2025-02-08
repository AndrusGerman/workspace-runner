package models

type Workspace struct {
	*Base `bson:",inline" json:",inline"`
	Name  string `bson:"name" json:"name"`
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Base: NewBase(),
		Name: name,
	}
}
