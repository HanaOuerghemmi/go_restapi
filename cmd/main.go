package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HanaOuerghemmi/go_restapi/cmd/api"
	"github.com/HanaOuerghemmi/go_restapi/config"
	"github.com/HanaOuerghemmi/go_restapi/db"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
	//? initialize db
	dbConn, err := db.NewPostgreSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPORT,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName))

	if err != nil {
		log.Fatal("error connecting to postgres", err)
	}
	if err := initDB(dbConn); err != nil {
		log.Fatal("connection with db error :", err)
	}
	//? start api server

	apiServer := api.NewAPIServe(":8080")
	if err := apiServer.Run(); err != nil {
		log.Fatal("error runing api server")
	}

}

func initDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	log.Println("Connected to database : ", config.Envs.DBName)
	return nil
}
