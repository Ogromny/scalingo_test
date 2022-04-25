package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/sclng-backend-test-v1/config"
	"github.com/Scalingo/sclng-backend-test-v1/repos"
	"github.com/Scalingo/sclng-backend-test-v1/stats"
)

func main() {
	log := logger.Default()
	log.Info("Initializing app")

	cfg, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Error("Fail to initialize configuration")
		os.Exit(-1)
	}

	log.Info("Initializing routes")

	router := handlers.NewRouter(log)
    repos.SetToken(cfg.Username, cfg.Token)
    repos.Register(router)
    stats.Register(router)

	// GET /stats

	log.WithField("port", cfg.Port).Info("Listening...")
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
