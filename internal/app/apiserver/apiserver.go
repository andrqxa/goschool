package apiserver

import (
	"net/http"
	"text/template"

	"github.com/andrqxa/goschool/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	tpl      *template.Template
	allUsers model.Users
	user     *model.User
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	allUsers = model.NewUsers()
}

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(newconfig *Config) *APIServer {
	return &APIServer{
		config: newconfig,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.handleRoot())
	s.router.HandleFunc("/add", s.handleAdd())
}

func (s *APIServer) handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := struct {
			Users []*model.User
		}{
			allUsers.ToArray(),
		}
		tpl.ExecuteTemplate(w, "index.gohtml", d)
	}
}

func (s *APIServer) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		newUser, err := model.NewUser(name, email)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		err = allUsers.AddUser(newUser)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
