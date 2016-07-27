package booking

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"golang.org/x/net/context"

	"github.com/davidwilde/barnet/appointment"
)

func MakeHandler(ctx context.Context, bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	bookAppointmentHandler := kithttp.NewServer(
		ctx,
		makeNewBookingEndpoint(bs),
		decodeBookAppointmentRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/booking/v1/appointments", bookAppointmentHandler).Methods("POST")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeBookAppointmentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Client          string    `json:"origin"`
		Stylist         string    `json:"stylist"`
		AppointmentTime time.Time `json:"appointment_time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	clientId, err := xid.FromString(body.Client)
	if err != nil {
		return nil, err
	}
	stylistId, err := xid.FromString(body.Stylist)

	if err != nil {
		return nil, err
	}

	return bookAppointmentRequest{
		Client:          clientId,
		Stylist:         stylistId,
		AppointmentTime: body.AppointmentTime,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; chareset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case appointment.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
