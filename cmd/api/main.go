package main

import (
	"bina/internal/es"
	"bina/internal/es/kafka"
	"bina/internal/service/webapi"
	"bina/internal/storage/psql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	//"github.com/spf13/viper"
)

func main() {

	//	::: LOGGER setup
	{
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.Println("Helo")
	}
	// ::: CONFIG setup
	{
		if err := initCfg(); err != nil {
			logrus.Fatal("reading cfg error ::: %s", err)
		}
	}
	// ::: ENV VARIABLES setup
	{
		err := godotenv.Load()
		if err != nil {
			logrus.Fatal("loading env variables ::: ", err)
		}
	}
	// ::: POSTGRES DB setup
	{
		cfg := &psql.Config{
			Host:     viper.GetString("db.host"),
			Password: viper.GetString("db.password"),
			Port:     viper.GetString("db.port"),
			DB:       viper.GetString("db.db"),
			User:     viper.GetString("db.user"),
			SSLMode:  viper.GetString("db.sslmode"),
		}
		db, err := psql.NewPostgresDB(cfg)
		if err != nil {
			logrus.Fatal("connecting to DB  ::: ", err)
		}
	}

	// ::: BINANCE WEPAPI setup

	{
		binannceWEBapi := ser.NewBinanceWebApi(&ser.BinanceWebApiCFG{
			APIKey:    "LALALL",
			APISecret: "LALALL",
		})
	}

	// ::: KAFKA setup


	kafkaWriter:= kafka.NewKafkaWriter(&kafka.Config{
		Addres: viper.GetString("kafka.addres"),
		Topic:  viper.GetString("kafka.topic"),
	})

massageBroker:= es.NewKafkaMessageBroker(kafkaWriter)



}

func initCfg() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigFile("cfg")
	viper.AddConfigPath("yaml")
	return viper.ReadInConfig()
}
