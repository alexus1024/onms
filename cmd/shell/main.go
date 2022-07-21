package main

import (
	"net/http"

	"github.com/alexus1024/onms/internal/api/server"
	"github.com/alexus1024/onms/internal/storage/memory"
	"github.com/sirupsen/logrus"
)

const serverAddr = ":4000"

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{}) // better for Kibana
	logger.WithField("addr", serverAddr).Info("App started")

	repo := memory.NewMemoryRepo()
	appContext := &server.AppContext{
		Log:  logrus.NewEntry(logger),
		Repo: repo,
	}
	mainHandler := server.GetMux(appContext)

	server := http.Server{
		Addr:    serverAddr,
		Handler: mainHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.WithError(err).Error("http server exited with error")
	}
}
