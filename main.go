package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"tomcze15/internal/database"
)

func createRouter(apiCfg *apiConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(getCorsConfig()))

	routerV1 := chi.NewRouter()

	routerV1.Get("/healthz", handlerReadiness)
	routerV1.Get("/err", handlerErr)
	routerV1.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", routerV1)

	return router
}

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(ENV_PATH)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, errCon := sql.Open("postgres", dbUrl)
	defer conn.Close()

	if errCon != nil {
		log.Fatal("Can't connect to database: ", errCon)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := createRouter(&apiCfg)

	srv := &http.Server{Handler: router, Addr: ":" + port}

	log.Printf("Server starting on port %v", port)
	errServer := srv.ListenAndServe()

	if errServer != nil {
		log.Fatal(errServer)
	}

	fmt.Println("My port is ", port)
}
