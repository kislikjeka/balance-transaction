package main

import (
	"TransactionTest/internal/domain"
	pgrepo "TransactionTest/internal/repository/postgres"
	"TransactionTest/internal/service/balance"
	"TransactionTest/pkg/config"
	"TransactionTest/pkg/db/postgres"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var transID int
	var resID int
	var val float64
	flag.IntVar(&transID, "t", 0, "Specify trans ID.")
	flag.IntVar(&resID, "r", 0, "Specify res ID.")
	flag.Float64Var(&val, "v", 0, "Specify Value.")

	flag.Usage = func() {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("./bin/transsaction -t 1 -r 2 -v 3.5\n")
	}
	flag.Parse()

	if transID == 0 || resID == 0 || val == 0 {
		logrus.Fatalf("Argument are requierd")
	}

	if err := config.InitConfig("config"); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := pgrepo.NewBalanceRepo(db)
	service := balance.NewService(repo)

	b1, err := service.Get(domain.ID(transID))
	if err != nil {
		logrus.Fatal(err)
	}
	b2, err := service.Get(domain.ID(resID))
	if err != nil {
		logrus.Fatal(err)
	}
	err = service.TransferFounds(b1, b2, float32(val))
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("Successful transaction")

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
