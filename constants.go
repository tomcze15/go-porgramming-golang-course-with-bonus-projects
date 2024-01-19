package main

import (
	"github.com/go-chi/cors"
)

func getCorsConfig() cors.Options {
	return cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
}

const (
	ENV_PATH = ".env"
)
