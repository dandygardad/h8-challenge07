package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type Postgres struct {
	Username string
	Password string
	Port     int
	Address  string
	Database string

	DB *sql.DB
}

type PsqlDB struct {
	*Postgres
}

var (
	PSQL *PsqlDB
)

func InitPostgres() error {
	PSQL = new(PsqlDB)
	portPostgres, errCvt := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if errCvt != nil {
		return errCvt
	}
	PSQL.Postgres = &Postgres{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     portPostgres,
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	err := PSQL.Postgres.OpenConnection()
	if err != nil {
		return err
	}

	// Ping
	err = PSQL.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) OpenConnection() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.Address, p.Port, p.Username, p.Password, p.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.DB = db

	// test connection
	errPing := p.DB.Ping()
	if errPing != nil {
		return errPing
	}

	fmt.Println("Success connect to db")
	return nil
}
