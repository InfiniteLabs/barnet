package booking

import (
	"testing"
	"time"

	"github.com/davidwilde/barnet/repository"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestBookNewAppointment(c *C) {

	var (
		clientRepository      = repository.NewInMemclient()
		stylistRepository     = repository.NewInMemStylist()
		appointmentRepository = repository.NewInMemApointment()
	)

	var bookingService = NewService(clientRepository, stylistRepository)

	appointmentTime := time.Date(2016, time.November, 10, 11, 0, 0, 0, time.UTC)

	appointmentID, err := bookingService.BookNewAppointment(client, stylist, appointmentTime)

	c.Assert(err, IsNil)

	appointment, err := appointmentRepository.Find(appointmentId)

	c.Assert(err, IsNil)

	c.Check(appointmentId, Equals, appointment.AppointmentID)
	c.Check(appointmentTime, Equals, appointment.AppointmentTime)
}
