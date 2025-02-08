package services

import (
	"bytes"
	"html/template"

	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
	"github.com/AndrusGerman/workspace-runner/templates"
)

type TemplateService struct {
	homeTemplate         *template.Template
	addWorkspaceTemplate *template.Template
}

func NewTemplateService() (ports.TemplateService, error) {

	var err error
	var homeTemplate *template.Template
	homeTemplate, err = template.New("homeTemplate").Parse(string(templates.HomeTemplate))
	if err != nil {
		return nil, err
	}

	var addWorkspaceTemplate *template.Template
	addWorkspaceTemplate, err = template.New("addWorkspaceTemplate").Parse(string(templates.AddWorkspaceTemplate))
	if err != nil {
		return nil, err
	}

	return &TemplateService{
		homeTemplate:         homeTemplate,
		addWorkspaceTemplate: addWorkspaceTemplate,
	}, nil

}

func (s *TemplateService) GetHomeTemplate(data any) ([]byte, error) {
	return s.executeTemplate(s.homeTemplate, data)
}

func (s *TemplateService) GetAddWorkspaceTemplate(data any) ([]byte, error) {
	return s.executeTemplate(s.addWorkspaceTemplate, data)
}

func (s *TemplateService) executeTemplate(template *template.Template, data any) ([]byte, error) {
	var buffer bytes.Buffer

	err := template.Execute(&buffer, data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
