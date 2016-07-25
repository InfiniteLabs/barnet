package booking

import (
	"fmt"
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
	stylist := stylist.New("Alice", "I love cutting hair")
	client := client.New("bob@testing.com", "Bob McTest", "Bob", "Male", "555-1234")

	appointmentTime := time.Date(2016, time.November, 10, 11, 0, 0, 0, time.UTC)

	appointmentId, err := bookingService.BookNewAppointment(*client, *stylist, appointmentTime)

	if err != nil {
		t.Error("Could not save an appointment")
	}

	fmt.Println(appointmentId.String())

	appointment, err := appointmentRepository.Find(appointmentId)

	if err != nil {
		t.Error("Saved appointment cannot be found in repository")
	}

	if appointmentTime == appointment.AppointmentTime {
		t.Errorf("BookNewAppointment appointment time expected %q, got %q", appointmentTime, appointment.AppointmentTime)
	}

}
