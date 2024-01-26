package main

import (
	"context"

	"github.com/Kartochnik010/test-effectivemobile-jan/internal/config"
	repo "github.com/Kartochnik010/test-effectivemobile-jan/internal/repository"
	"github.com/Kartochnik010/test-effectivemobile-jan/internal/repository/postgres"
	"github.com/Kartochnik010/test-effectivemobile-jan/internal/service"
	transport "github.com/Kartochnik010/test-effectivemobile-jan/internal/transport/http"
	"github.com/Kartochnik010/test-effectivemobile-jan/pkg/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	l := logger.NewLogger("")
	db, err := postgres.InitDB(context.Background(), cfg.DSN)
	if err != nil {
		panic(err)
	}
	l.Info().Msg("Successful connection to database")
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Err(err).Msg("failed to close db connection")
		}
	}()
	// *"github.com/Kartochnik010/test-effectivemobile-jan/internal/repository".Repository)
	// *"github.com/Kartochnik010/test-effectivemobile-jan/internal/repository".Repository
	repos := repo.NewPostgresRepo(db)
	service := service.NewService(repos)
	t := transport.NewHandler(service, l)
	var srv transport.HttpServer
	l.Info().Msg("Starting server on port :" + cfg.Port)
	if err := srv.Run(":"+cfg.Port, t.Routes()); err != nil {
		l.Err(err).Msg(err.Error())
	}
}
