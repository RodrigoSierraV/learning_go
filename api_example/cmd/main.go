package main

import (
	"log"
	"runtime/trace"

	"github.com/RodrigoSierraV/learnin_go/api_example/cmd/api"
	"github.com/RodrigoSierraV/learnin_go/api_example/config"
	"github.com/RodrigoSierraV/learnin_go/api_example/db"
	"github.com/go-sql-driver/mysql"
)


func main() {
	db, dbErr := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DB_USER,
		Passwd: config.Envs.DB_PASSWORD,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DB_NAME,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if dbErr != nil {
		log.Fatal(dbErr)
	}
	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}