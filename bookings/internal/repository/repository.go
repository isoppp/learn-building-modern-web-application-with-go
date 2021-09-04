package repository

import "github.com/isoppp/learn-building-modern-web-application-with-go/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
