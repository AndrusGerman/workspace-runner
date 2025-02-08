package ports

type TemplateService interface {
	GetHomeTemplate(data any) ([]byte, error)
	GetAddWorkspaceTemplate(data any) ([]byte, error)
	GetEditWorkspaceTemplate(data any) ([]byte, error)
}
