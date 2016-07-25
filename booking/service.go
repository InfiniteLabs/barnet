package booking

import (
	"github.com/davidwilde/barnet/appointment"
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/rs/xid"
	"time"
)

type Service interface {
	BookNewAppointment(client client.Client, stylist stylist.Stylist, appointmentTime time.Time) (xid.ID, error)
}

type service struct {
	appointments appointment.Repository
	clients      client.Repository
	stylists     stylist.Repository
}

func (s *service) BookNewAppointment(client client.Client, stylist stylist.Stylist, appointmentTime time.Time) (xid.ID, error) {
	appointment := appointment.New(appointmentTime, client, stylist)
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
