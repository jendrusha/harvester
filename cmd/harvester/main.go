package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"golang.org/x/exp/slog"

	"github.com/jendrusha/harvester/internal/app"
	"github.com/jendrusha/harvester/internal/usecase/createharvest"
	"github.com/jendrusha/harvester/pkg/storage"
)

func main() {
	// app dependencies
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mysqlStorage, err := storage.NewMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	// prepare app configuration
	cfg := buildConfig()

	// usecase handlers
	createHarvestHandler := createharvest.NewHandler(mysqlStorage)

	// new app instance, usecase handlers are registred too
	harvester := app.New(
		app.WithConfig(cfg),
		app.WithLogger(logger),

		app.RegisterUsecase[createharvest.CreateHarvestRequest](
			http.MethodPost,
			"v1",
			"harvest",
			createHarvestHandler,
		),
	)

	if err := harvester.Run(); err != nil {
		log.Fatalln(err)
	}
}

func buildConfig() app.Config {
	port := flag.Int("port", 3000, "Application port")
	env := flag.String("env", "dev", "Application environment")

	flag.Parse()

	return app.Config{
		Env:  *env,
		Port: *port,
	}
}
