package main

import (
	"errors"
	"os"
	"swe"
	"swe/model"
	"swe/pkg/handler"
	"swe/pkg/repository"
	"swe/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found, please set")
	}
}

func main() {
	logrus.Print("Startup server")
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initzializing env: %s", err.Error())
	}

	db, err := repository.NewPostgreDB(
		os.Getenv("DSN"),
	)

	if err != nil {
		logrus.Fatalf(err.Error())
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	err = migrateGorm(gormDB)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db, gormDB)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	staticHandler := handler.NewStaticHandler(service)
	/*if os.Getenv("Seed") != "" {
		err := repos.Seeder.SeedCategory()
		if err != nil {
			log.Print(err)
		}
	}*/
	srv := new(swe.Server)
	staticSrv := new(swe.Server)

	logrus.Print("Server Runing on Dev mode")

	go func() {
		if err := srv.Run(os.Getenv("APIPortHTTP"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	if err := staticSrv.Run(os.Getenv("StaticPortHTTP"), staticHandler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func migrateGorm(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.Doctor{})
	if err != nil {
		return err
	}
	return nil
}

func initEnv() error {

	reqs := []string{
		"StaticPortHTTP",
		"APIPortHTTP",
	}

	for i := 0; i < len(reqs); i++ {
		_, exists := os.LookupEnv(reqs[i])

		if !exists {
			return errors.New(".env variables not set")
		}
	}

	return nil
}
