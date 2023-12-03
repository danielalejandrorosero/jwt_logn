package main

import (
	"go-auth-service/pkg/config"
	"go-auth-service/pkg/logger"
	"go-auth-service/pkg/router"
)

func main() {
	err := config.Load()

	if err != nil {
		logger.Logger.Fatal(err)
	}

	err = router.Run()

	logger.Logger.Fatal(err)
}
