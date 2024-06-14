package main

import (
	"bina/internal/es"
	"bina/internal/es/kafka"
	service2 "bina/internal/service"
	ser "bina/internal/service/webapi"
	storage2 "bina/internal/storage"
	"bina/internal/storage/psql"
	"bina/internal/transport/rest"
	handler2 "bina/internal/transport/rest/handler"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	//"github.com/spf13/viper"
)

func main() {

	//	::: LOGGER setup

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Println("Helo")

	// ::: CONFIG setup

	if err := initCfg(); err != nil {
		logrus.Fatal("reading cfg error ::: %s", err)
	}

	// ::: ENV VARIABLES setup

	err:= godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("loading env variables ::: ", err)
	}

	// ::: POSTGRES DB setup

	cfg := &psql.Config{
		Host:     viper.GetString("db.host"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     viper.GetString("db.port"),
		DB:       viper.GetString("db.dbname"),
		User:     viper.GetString("db.user"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	db, err := psql.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatal("connecting to DB  ::: ", err)
	}

	// ::: BINANCE WEPAPI setup

	binannceWEBapi := ser.NewBinanceWebApi(&ser.BinanceWebApiCFG{
		APIKey:    "LALALL",
		APISecret: "LALALL",
	})

	// ::: KAFKA setup
	kafkaWriter := kafka.NewKafkaWriter(&kafka.Config{
		Addres: viper.GetString("kafka.address"),
		Topic:  viper.GetString("kafka.topic"),
	})

	massageBroker := es.NewKafkaMessageBroker(kafkaWriter)

	// ::: STORAGE setup

	storage := storage2.NewStorage(db)

	// ::: SERVICE setup

	service := service2.NewService(storage, binannceWEBapi, massageBroker)

	// ::: HANDLER setup

	handler := handler2.NewHandler(service)

	// ::: SERVER setup
	server := rest.NewServer(viper.GetString("port"), handler.InitRoutes())

	//	:::SERVER ON

	go func() {

		if err := server.Run(); err != nil {
			logrus.Errorf("Error while running server :::: %w", err)
		}
	}()
	logrus.Printf("App started on localhost:%s \n", viper.GetString("port"))

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	//	:::SERVER OFF

	logrus.Println("App shutting down")

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Fatalf("Error while shutting down ::: %w\n", err)
		os.Exit(1)
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("Error while shutting down ::: %w\n", err)
		os.Exit(1)
	}
logrus.Println("!!! GOODBYE !!!")
}

func initCfg() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
