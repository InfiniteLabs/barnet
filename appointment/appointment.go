package appointment

import (
	"errors"
	"time"

	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/rs/xid"
)

type AppointmentID string

type Appointment struct {
	AppointmentId   AppointmentID
	AppointmentTime time.Time
	ClientId        client.ClientId
	StylistId       stylist.StylistId
}

func New(appointmentTime time.Time, clientId client.ClientId, stylistId stylist.StylistId) *Appointment {
	return &Appointment{
		AppointmentTime: appointmentTime,
		ClientId:        clientId,
		StylistId:       stylistId,
	}
}

type Repository interface {
	Store(appointment *Appointment) error
	Find(appointmentId AppointmentID) (*Appointment, error)
	FindStylistAtTime(stylistId stylist.StylistId, time time.Time) (*Appointment, error)
}

func NextAppointmentID() AppointmentID {
	guid := xid.New()
	return AppointmentID(guid.String())
}

var ErrUnknown = errors.New("unknown appointment")
