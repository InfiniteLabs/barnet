package booking

import (
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/stylist"
	"github.com/go-kit/kit/endpoint"
	"github.com/rs/xid"
	"golang.org/x/net/context"
	"time"
)

type bookAppointmentRequest struct {
	Client          client.Client   `json:"client"`
	Stylist         stylist.Stylist `json:"stylistId"`
	AppointmentTime time.Time       `json:"appointmentTime"`
}

type bookAppointmentResponse struct {
	AppointmentId xid.ID `json:"appointmentId"`
	Err           string `json:"err,omitempty"`
}

func makeNewBookingEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(bookAppointmentRequest)
		v, err := svc.BookNewAppointment(req.Client, req.Stylist, req.AppointmentTime)
		if err != nil {
			return bookAppointmentResponse{v, err.Error()}, nil
		}
		return bookAppointmentResponse{v, ""}, nil
	}
}
