package booking

import (
	"testing"
	"time"

	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/repository"
	"github.com/davidwilde/barnet/stylist"
)

func TestBookNewAppointment(t *testing.T) {

	var (
		clientRepository      = repository.NewInMemClient()
		stylistRepository     = repository.NewInMemStylist()
		appointmentRepository = repository.NewInMemAppointment()
	)

	var bookingService = NewService(appointmentRepository, clientRepository, stylistRepository)
	s := stylist.New("Alice", "I love cutting hair")
	c := client.Client{ClientId: "123", EmailAddress: "bob@testing.com", FullName: "Bob McTest"}

	appointmentTime := time.Date(2016, time.November, 10, 11, 0, 0, 0, time.UTC)

	appointmentId, err := bookingService.BookNewAppointment(c.ClientId, s.StylistId, appointmentTime)

	if err != nil {
		t.Error("Could not save an appointment")
	}

	appointment, err := appointmentRepository.Find(appointmentId)

	if err != nil {
		t.Error("Saved appointment cannot be found in repository")
	}

	if appointmentTime != appointment.AppointmentTime {
		t.Errorf("BookNewAppointment appointment time expected %q, got %q", appointmentTime, appointment.AppointmentTime)
	}
}

func TestBookNewAppointmentTimeUnavailable(t *testing.T) {

	var (
		clientRepository      = repository.NewInMemClient()
		stylistRepository     = repository.NewInMemStylist()
		appointmentRepository = repository.NewInMemAppointment()
	)

	var bookingService = NewService(appointmentRepository, clientRepository, stylistRepository)
	stylist := stylist.New("Alice", "I love cutting hair")
	firstC := client.Client{ClientId: "123", EmailAddress: "bob@testing.com", FullName: "Bob McTest"}
	secondC := client.Client{ClientId: "456", EmailAddress: "charles@testing.com", FullName: "Charles Von Test"}

	appointmentTime := time.Date(2016, time.November, 10, 11, 0, 0, 0, time.UTC)

	_, err := bookingService.BookNewAppointment(firstC.ClientId, stylist.StylistId, appointmentTime)
	_, err = bookingService.BookNewAppointment(secondC.ClientId, stylist.StylistId, appointmentTime)

	if err == nil {
		t.Error("Duplicate appointment should not be saved")
	}
}
