package databases

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connectionString := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString Error...")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid ...")
	}

}

func CreateCon() *sql.DB {
	return db
}
