package main

import (
	"net/http"
	"os"

	"github.com/alexus1024/onms/internal/api/server"
	"github.com/alexus1024/onms/internal/config"
	"github.com/alexus1024/onms/internal/storage/memory"
	"github.com/sirupsen/logrus"
)

func main() {

	if len(os.Args) >= 2 && os.Args[1] == "--help" {
		config.PrintHelp()
		return
	}

	appConfig, err := config.ReadConfig()
	if err != nil {
		config.PrintHelp()
		panic("can not read environment variables: " + err.Error())
	}

	logger := logrus.New()
	logger.SetLevel(appConfig.LogLevel)

	if appConfig.JsonLog {
		logger.SetFormatter(&logrus.JSONFormatter{}) // better for Kibana
	}

	logger.WithField("addr", appConfig.ServerAddr).WithField("min_log_level", logger.Level).Info("App started")

	repo := memory.NewMemoryRepo()
	appContext := &server.AppContext{
		Log:  logrus.NewEntry(logger),
		Repo: repo,
	}
	mainHandler := server.GetMux(appContext)

	server := http.Server{
		Addr:    appConfig.ServerAddr,
		Handler: mainHandler,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.WithError(err).Error("http server exited with error")
	}
}
