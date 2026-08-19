package utils

import (
	"database/sql"
	"log"
	"fmt"
	"os"
	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var DB *sql.DB

func SetDB(){
	// Capture connection properties.
	cfg := mysql.NewConfig();
	cfg.User = os.Getenv("DB_USER");
	cfg.Passwd = os.Getenv("DB_PASSWORD");
	cfg.Net = "tcp";
	cfg.Addr = os.Getenv("DB_ADDRESS");
	cfg.DBName = os.Getenv("DB_NAME");

	var err error
	DB, err = sql.Open("mysql",cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database Connected!!")
}