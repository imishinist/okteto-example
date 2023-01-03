package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("request received",
		zap.String("path", r.URL.Path),
		zap.String("remote-address", r.RemoteAddr))

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, `{"message":"Hello World"}`)
}

func main() {
	http.HandleFunc("/", Handler)

	logger, _ := zap.NewProduction()
	logger.Info("start listen http",
		zap.Int("port", 8080),
		zap.String("host", "localhost"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
