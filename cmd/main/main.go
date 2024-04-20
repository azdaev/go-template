package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/azdaev/template/internal/audiolectures/api"
	"github.com/azdaev/template/internal/audiolectures/repo"
	"github.com/azdaev/template/internal/audiolectures/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	if err := godotenv.Load(); err != nil {
		panic("no .env file found")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	audioLecturesRepo := repo.New(db)
	audioLecturesService := service.New(audioLecturesRepo)
	audioLecturesHandler := api.New(audioLecturesService, log)

	gin.SetMode(os.Getenv("GIN_MODE"))
	if err = audioLecturesHandler.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(errors.Wrap(err, "Run"))
	}
}
