package api

import (
	"log"
	"os"

	"github.com/rwiteshbera/microservices_demo/authenticationService/config"
)

var AuthHandlerInstance *AuthHandler

type AuthHandler struct {
	Logger *log.Logger
	Env    *config.EnvConfig
}

func init() {
	envConfig, err := config.LoadEnvConfig()
	if err != nil {
		log.Panic(err)
	}

	logger := log.New(os.Stdout, "auth_api\t", log.LstdFlags)

	AuthHandlerInstance = &AuthHandler{
		Logger: logger,
		Env:    envConfig,
	}

}
