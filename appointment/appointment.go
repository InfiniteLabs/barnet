package appointment

import (
	"errors"
	"time"

	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/rs/xid"
)

type AppointmentId xid.ID

type Appointment struct {
	AppointmentId   AppointmentId
	AppointmentTime time.Time
	Client          client.Client
	Stylist         stylist.Stylist
}

func New(appointmentTime time.Time, client client.Client, stylist stylist.Stylist) *Appointment {
	guid := xid.New()

	return &Appointment{
		AppointmentId: AppointmentId(guid),
		Client:        client,
		Stylist:       stylist,
	}
}

type Repository interface {
	Store(appointment *Appointment) error
	Find(appointmentId AppointmentId) (*Appointment, error)
}

var ErrUnknown = errors.New("unknown appointment")
