package booking

import (
	"errors"
	"time"

	"github.com/davidwilde/barnet/appointment"
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/rs/xid"
)

var ErrInvalidArgument = errors.New("Invalid Argument")

type Service interface {
	BookNewAppointment(client client.Client, stylist stylist.Stylist, appointmentTime time.Time) (xid.ID, error)
}

type service struct {
	appointments appointment.Repository
	clients      client.Repository
	stylists     stylist.Repository
}

func (s *service) BookNewAppointment(client client.Client, stylist stylist.Stylist, appointmentTime time.Time) (xid.ID, error) {
	if appointmentTime.IsZero() {
		return nil, ErrInvalidArgument
	}
	val, err := s.appointments.FindStylistAtTime(stylist, appointmentTime)
	if err != nil {
		panic(err)
	}
	appointment := appointment.New(appointmentTime, client, stylist)
	if val != nil {
		return appointment.AppointmentId, errors.New("Cannot make duplicate appointment")
	}
	s.appointments.Store(appointment)
	return appointment.AppointmentId, nil
}

func NewService(ar appointment.Repository, cr client.Repository, sr stylist.Repository) Service {
	return &service{
		appointments: ar,
		clients:      cr,
		stylists:     sr,
	}
}
