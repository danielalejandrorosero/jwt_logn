package router

import (
	"fmt"
	"go-auth-service/pkg/config"
	"go-auth-service/pkg/controller"
	"go-auth-service/pkg/logger"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", controller.Health)

	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/check", controller.Check)

	logger.Logger.Println("Start http server on " + fmt.Sprintf("%s:%d", config.Instance.Application.Url, config.Instance.Application.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Instance.Application.Port), mux)
	return err
}
