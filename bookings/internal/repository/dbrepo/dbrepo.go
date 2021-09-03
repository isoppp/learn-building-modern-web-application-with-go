package dbrepo

import (
	"database/sql"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/repository"

	"github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
