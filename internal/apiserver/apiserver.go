package apiserver

import (
	"RPG/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// API server
type APIserver struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Store
}

// new...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}

// start...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("start api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIserver) configureStore() error {
	st := storage.New(s.config.Storage)

	if err := st.Open(); err != nil {
		return err
	}
	s.storage = st
	return nil
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
