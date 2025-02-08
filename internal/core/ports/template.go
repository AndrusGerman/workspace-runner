package ports

type TemplateService interface {
	GetHomeTemplate(data any) ([]byte, error)
	GetAddWorkspaceTemplate(data any) ([]byte, error)
}
