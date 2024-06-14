package main

import (
	"bina/internal/es"
	"bina/internal/es/kafka"
	service2 "bina/internal/service"
	ser "bina/internal/service/webapi"
	storage2 "bina/internal/storage"
	"bina/internal/storage/psql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
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

		err := godotenv.Load()
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
		Addres: viper.GetString("kafka.addres"),
		Topic:  viper.GetString("kafka.topic"),
	})

	massageBroker := es.NewKafkaMessageBroker(kafkaWriter)

	// ::: STORAGE setup

	storage:= storage2.NewStorage(db)


	// ::: SERVICE setup

	service := service2.NewService(storage,binannceWEBapi,massageBroker)

	// ::: HANDLER setup




}

func initCfg() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
