package booking

import (
	"github.com/davidwilde/barnet/appointment"
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"time"
)

type bookAppointmentRequest struct {
	ClientId        client.ClientId   `json:"clientId"`
	StylistId       stylist.StylistId `json:"stylistId"`
	AppointmentTime time.Time         `json:"appointmentTime"`
}

type bookAppointmentResponse struct {
	AppointmentId appointment.AppointmentID `json:"appointmentId"`
	Err           string                    `json:"err,omitempty"`
}

func makeNewBookingEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(bookAppointmentRequest)
		v, err := svc.BookNewAppointment(req.ClientId, req.StylistId, req.AppointmentTime)
		if err != nil {
			return bookAppointmentResponse{v, err.Error()}, nil
		}
		return bookAppointmentResponse{v, ""}, nil
	}
}
