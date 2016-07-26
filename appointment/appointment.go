package appointment

import (
	"errors"
	"time"

	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/rs/xid"
)

type Appointment struct {
	AppointmentId   xid.ID
	AppointmentTime time.Time
	Client          client.Client
	Stylist         stylist.Stylist
}

func New(appointmentTime time.Time, client client.Client, stylist stylist.Stylist) *Appointment {
	guid := xid.New()

	return &Appointment{
		AppointmentId:   guid,
		AppointmentTime: appointmentTime,
		Client:          client,
		Stylist:         stylist,
	}
}

type Repository interface {
	Store(appointment *Appointment) error
	Find(appointmentId xid.ID) (*Appointment, error)
	FindStylistAtTime(stylist stylist.Stylist, time time.Time) (*Appointment, error)
}

var ErrUnknown = errors.New("unknown appointment")
