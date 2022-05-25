package apiserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver ...
type APIserver struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	go s.grabeData(10 * time.Second)
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
	s.router.HandleFunc("/get_info", s.getCurInfo())
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *APIserver) getCurInfo() http.HandlerFunc {
	jsonstring, err := json.Marshal(MakeRequest())
	if err != nil {
		s.logger.Error(err)
		return nil

	} else {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, string(jsonstring))
		}
	}
}

func (s *APIserver) grabeData(d time.Duration) {
	for range time.Tick(d) {
		log.Print("Grabber has done some stuff")
	}
}
