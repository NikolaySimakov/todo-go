package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DataBaseInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (info *DataBaseInfo) GetEnv() {
	info.Host = os.Getenv("DB_HOST")
	info.Port = os.Getenv("DB_PORT")
	info.User = os.Getenv("DB_USER")
	info.Password = os.Getenv("DB_PASSWORD")
	info.Name = os.Getenv("DB_NAME")
}

func (info *DataBaseInfo) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Name)
}

func Database() *sql.DB {
	// get connection info
	connInfo := DataBaseInfo{}
	connInfo.GetEnv()

	// open DataBase
	db, err := sql.Open("postgres", connInfo.ToString())

	if err != nil {
		log.Fatal(err)
	}

	return db
}
