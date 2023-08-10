package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataBaseInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func getEnvValue(key string, defaultValue string) string {
	if host, exist := os.LookupEnv(key); exist {
		return host
	}
	return defaultValue
}

func (info *DataBaseInfo) GetEnv() {
	info.Host = getEnvValue("DB_HOST", "localhost")
	info.Port = getEnvValue("DB_PORT", "5432")
	info.User = getEnvValue("DB_USER", "postgres")
	info.Password = getEnvValue("DB_PASSWORD", "")
	info.Name = getEnvValue("DB_NAME", "todo")
}

func (info *DataBaseInfo) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Name)
}

func init() {
	// loads values from .env into the system
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file found")
	}
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
