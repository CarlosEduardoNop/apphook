package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func OpenConnection() (*sql.DB, error) {
	godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	CreateDatabase(port)

	sc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), port, os.Getenv("DATABASE_NAME"))

	conn, err := sql.Open("mysql", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

func CreateDatabase(port int) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), port))

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS sistema")

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
