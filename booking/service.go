package booking

import (
	"errors"
	"time"

	"github.com/davidwilde/barnet/appointment"
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
)

var ErrInvalidArgument = errors.New("Invalid Argument")

type Service interface {
	BookNewAppointment(client client.ClientId, stylist stylist.StylistId, appointmentTime time.Time) (appointment.AppointmentID, error)
}

type service struct {
	appointments appointment.Repository
	clients      client.Repository
	stylists     stylist.Repository
}

func (s *service) BookNewAppointment(clientId client.ClientId, stylistId stylist.StylistId, appointmentTime time.Time) (appointment.AppointmentID, error) {
	if appointmentTime.IsZero() {
		return "", ErrInvalidArgument
	}
	val, err := s.appointments.FindStylistAtTime(stylistId, appointmentTime)
	if err != nil {
		panic(err)
	}
	a := appointment.New(appointmentTime, clientId, stylistId)
	a.AppointmentId = appointment.NextAppointmentID()

	if val != nil {
		return "", errors.New("Cannot make duplicate appointment")
	}
	if err := s.appointments.Store(a); err != nil {
		return "", err
	}
	return a.AppointmentId, nil
}

func NewService(ar appointment.Repository, cr client.Repository, sr stylist.Repository) Service {
	return &service{
		appointments: ar,
		clients:      cr,
		stylists:     sr,
	}
}
