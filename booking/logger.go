package booking

import (
	"time"

	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/go-kit/kit/log"

	"github.com/davidwilde/barnet/appointment"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) BookNewAppointment(clientId client.ClientId, stylistId stylist.StylistId, appointmentTime time.Time) (appointmentId appointment.AppointmentID, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "book",
			"client", clientId,
			"stylist", stylistId,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.BookNewAppointment(clientId, stylistId, appointmentTime)
}
