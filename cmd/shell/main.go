package main

import (
	"net/http"

	"github.com/alexus1024/onms/internal/api/server"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.Info("App started")

	h := server.GetMux(logrus.NewEntry(logger))

	http.Handle("/", h)

}
