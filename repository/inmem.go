package repository

import (
	"sync"
	"time"

	"github.com/davidwilde/barnet/appointment"
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
)

type stylistRepository struct {
	mtx      sync.RWMutex
	stylists map[stylist.StylistId]*stylist.Stylist
}

func (r *stylistRepository) Store(s *stylist.Stylist) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.stylists[s.StylistId] = s
	return nil
}

func (r *stylistRepository) Find(stylistID stylist.StylistId) (*stylist.Stylist, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.stylists[stylistID]; ok {
		return val, nil
	}
	return nil, stylist.ErrUnknown
}

func (r *stylistRepository) FindAll() []*stylist.Stylist {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*stylist.Stylist, 0, len(r.stylists))
	for _, val := range r.stylists {
		c = append(c, val)
	}
	return c
}

func NewInMemStylist() stylist.Repository {
	return &stylistRepository{
		stylists: make(map[stylist.StylistId]*stylist.Stylist),
	}
}

type clientRepository struct {
	mtx     sync.RWMutex
	clients map[client.ClientId]*client.Client
}

func (r *clientRepository) Store(c *client.Client) error {
	r.mtx.RLock()
	defer r.mtx.Unlock()
	r.clients[c.ClientId] = c
	return nil
}

func (r *clientRepository) Find(clientId client.ClientId) (*client.Client, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.clients[clientId]; ok {
		return val, nil
	}
	return nil, client.ErrUnknown
}

func NewInMemClient() client.Repository {
	return &clientRepository{
		clients: make(map[client.ClientId]*client.Client),
	}
}

type appointmentRepository struct {
	mtx          sync.RWMutex
	appointments map[appointment.AppointmentID]*appointment.Appointment
}

func (r *appointmentRepository) Store(a *appointment.Appointment) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.appointments[a.AppointmentId] = a
	return nil
}

func (r *appointmentRepository) Find(appointmentId appointment.AppointmentID) (*appointment.Appointment, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.appointments[appointmentId]; ok {
		return val, nil
	}
	return nil, appointment.ErrUnknown
}

func (r *appointmentRepository) FindStylistAtTime(stylistId stylist.StylistId, appointmentTime time.Time) (*appointment.Appointment, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	// TODO: There is probably a better data structure to this which would involve mapping to a stylist
	// and sorting by appointment date. This way a binary search could occur on an appointment date.
	for _, v := range r.appointments {
		if v.StylistId == stylistId && v.AppointmentTime == appointmentTime {
			return v, nil
		}
	}
	return nil, nil
}

func NewInMemAppointment() appointment.Repository {
	return &appointmentRepository{
		appointments: make(map[appointment.AppointmentID]*appointment.Appointment),
	}
}
