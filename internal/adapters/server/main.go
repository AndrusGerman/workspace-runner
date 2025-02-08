package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/AndrusGerman/go-criteria"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
)

type server struct {
	workspaceService ports.WorkspaceService
	templateService  ports.TemplateService
	projectService   ports.ProjectService
}

func NewServer(
	workspaceService ports.WorkspaceService,
	templateService ports.TemplateService,
	projectService ports.ProjectService,
) *server {
	return &server{
		workspaceService: workspaceService,
		templateService:  templateService,
		projectService:   projectService,
	}
}

func (s *server) Start() {
	var port = "8000"

	log.Println("Starting server on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		workspaces, err := s.workspaceService.Search(context.Background(), criteria.EmptyCriteria())
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		template, err := s.templateService.GetHomeTemplate(workspaces)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.html(w, http.StatusOK, template)
	})

	http.HandleFunc("/add-workspace", func(w http.ResponseWriter, r *http.Request) {
		template, err := s.templateService.GetAddWorkspaceTemplate(nil)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.html(w, http.StatusOK, template)
	})

	http.HandleFunc("/edit-workspace", func(w http.ResponseWriter, r *http.Request) {
		idString := r.URL.Query().Get("id")

		if idString == "" {
			s.error(w, http.StatusBadRequest, "ID is required")
			return
		}

		id, err := types.NewIdByString(idString)
		if err != nil {
			s.error(w, http.StatusBadRequest, err.Error())
			return
		}

		workspace, err := s.workspaceService.GetById(context.Background(), id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		projects, err := s.projectService.GetByWorkspaceId(context.Background(), id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		type editWorkspaceTemplateData struct {
			*models.Workspace
			Projects []*models.Project
		}

		var data = editWorkspaceTemplateData{
			Workspace: workspace,
			Projects:  projects,
		}

		template, err := s.templateService.GetEditWorkspaceTemplate(data)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.html(w, http.StatusOK, template)

	})

	http.HandleFunc("/api/workspace/add", func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()
		if err != nil {
			s.error(w, http.StatusBadRequest, err.Error())
			return
		}

		var name = r.FormValue("name")
		var description = r.FormValue("description")

		if name == "" {
			s.error(w, http.StatusBadRequest, "Name is required")
			return
		}

		if description == "" {
			s.error(w, http.StatusBadRequest, "Description is required")
			return
		}

		err = s.workspaceService.Create(context.Background(), models.NewWorkspace(name, description))

		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/api/workspaces", func(w http.ResponseWriter, r *http.Request) {
		workspaces, err := s.workspaceService.Search(context.Background(), criteria.EmptyCriteria())
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.json(w, http.StatusOK, workspaces)
	})

	http.ListenAndServe(":"+port, nil)

}

func (s *server) json(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
	w.Header().Set("Content-Type", "application/json")
}

func (s *server) html(w http.ResponseWriter, statusCode int, data []byte) {
	w.WriteHeader(statusCode)
	w.Write(data)

	w.Header().Set("Content-Type", "text/html")
}

func (s *server) error(w http.ResponseWriter, statusCode int, message string) {
	s.json(w, statusCode, map[string]string{"error": message})
}
