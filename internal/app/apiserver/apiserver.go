package apiserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Irishery/golang_webserv.git/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver ...
type APIserver struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
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

	if err := s.configureStore(); err != nil {
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

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/", s.MainPart())
	s.router.HandleFunc("/get_info", s.getCurInfo())
}

func (s *APIserver) MainPart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/get_info", http.StatusFound)
	}
}

func (s *APIserver) getCurInfo() http.HandlerFunc {
	curArray, err := s.store.Currency().GetAll()
	if err != nil {
		s.logger.Error(err)

		return nil
	}

	jsonstring, err := json.Marshal(curArray)
	if err != nil {
		s.logger.Error(err)

		return nil
	}

	s.logger.Info("Server has sent currency info")

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, string(jsonstring))
		if err != nil {
			s.logger.Error(err)
		}
	}
}

func (s *APIserver) grabeData(d time.Duration) {
	for range time.Tick(d) {
		data := MakeRequest()
		log.Print("Data has been grabed")
		go s.store.Currency().CreateMany(data)
	}
}
