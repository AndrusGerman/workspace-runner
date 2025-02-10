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
	"github.com/gorilla/mux"
)

type server struct {
	workspaceService ports.WorkspaceService
	projectService   ports.ProjectService
	router           *mux.Router
}

func NewServer(
	workspaceService ports.WorkspaceService,
	projectService ports.ProjectService,
) *server {
	return &server{
		workspaceService: workspaceService,
		projectService:   projectService,
		router:           mux.NewRouter(),
	}
}

func (s *server) Start() {
	var port = "8000"

	log.Println("Starting server on port", port)

	log.Println("Server http://localhost:" + port)

	s.Get("/api/workspaces/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]

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

		s.json(w, http.StatusOK, workspace)
	})

	s.Get("/api/workspaces/{id}/projects", func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]

		if idString == "" {
			s.error(w, http.StatusBadRequest, "ID is required")
			return
		}

		id, err := types.NewIdByString(idString)
		if err != nil {
			s.error(w, http.StatusBadRequest, err.Error())
			return
		}

		projects, err := s.projectService.GetByWorkspaceId(context.Background(), id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.json(w, http.StatusOK, projects)
	})

	s.Delete("/api/projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]

		if idString == "" {
			s.error(w, http.StatusBadRequest, "ID is required")
			return
		}

		id, err := types.NewIdByString(idString)
		if err != nil {
			s.error(w, http.StatusBadRequest, err.Error())
			return
		}

		err = s.projectService.Delete(context.Background(), id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.json(w, http.StatusOK, map[string]string{"message": "Project deleted"})
	})

	s.Post("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			s.error(w, http.StatusBadRequest, err.Error())
			return
		}
		project.Id = types.NewId().GetPrimitive()
		err = s.projectService.Create(context.Background(), &project)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.json(w, http.StatusOK, project)
	})

	s.Post("/api/workspace/add", func(w http.ResponseWriter, r *http.Request) {

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

	s.Get("/api/workspaces", func(w http.ResponseWriter, r *http.Request) {
		workspaces, err := s.workspaceService.Search(context.Background(), criteria.EmptyCriteria())
		if err != nil {
			s.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		s.json(w, http.StatusOK, workspaces)
	})

	http.Handle("/", s.router)
	http.ListenAndServe(":"+port, nil)

}

func (s *server) Get(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, s.generic(handler)).Methods("GET", "OPTIONS")
}

func (s *server) Post(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, s.generic(handler)).Methods("POST", "OPTIONS")
}

func (s *server) Put(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, s.generic(handler)).Methods("PUT", "OPTIONS")
}

func (s *server) Delete(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, s.generic(handler)).Methods("DELETE", "OPTIONS")
}

func (s *server) generic(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}

}

func (s *server) json(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *server) html(w http.ResponseWriter, statusCode int, data []byte) {
	w.WriteHeader(statusCode)
	w.Write(data)

	w.Header().Set("Content-Type", "text/html")
}

func (s *server) error(w http.ResponseWriter, statusCode int, message string) {
	s.json(w, statusCode, map[string]string{"error": message})
}
