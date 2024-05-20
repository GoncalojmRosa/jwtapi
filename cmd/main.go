package main

import (
	"database/sql"
	"goncalojrmosa/jwtapi/cmd/api"
	"goncalojrmosa/jwtapi/config"
	"goncalojrmosa/jwtapi/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	sv := api.NewAPIServer(":8080", nil)
	if err := sv.Run(); err != nil{
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB)  {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
}
