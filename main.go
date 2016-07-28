package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/davidwilde/barnet/appointment"
	"github.com/davidwilde/barnet/booking"
	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/repository"
	"github.com/davidwilde/barnet/stylist"
	"github.com/go-kit/kit/log"
)

func main() {
	var (
		addr = ":8080"

		ctx = context.Background()
	)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = &serializedLogger{Logger: logger}
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	var (
		appointments appointment.Repository
		clients      client.Repository
		stylists     stylist.Repository
	)

	appointments = repository.NewInMemAppointment()
	clients = repository.NewInMemClient()
	stylists = repository.NewInMemStylist()

	var bs booking.Service
	bs = booking.NewService(appointments, clients, stylists)
	bs = booking.NewLoggingService(log.NewContext(logger).With("component", "booking"), bs)

	httpLogger := log.NewContext(logger).With("component", "http")

	mux := http.NewServeMux()

	mux.Handle("/booking/v1/", booking.MakeHandler(ctx, bs, httpLogger))

	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		fmt.Println("listening and serving")
		errs <- http.ListenAndServe(addr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

type serializedLogger struct {
	mtx sync.Mutex
	log.Logger
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (l *serializedLogger) Log(keyvals ...interface{}) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	return l.Logger.Log(keyvals...)
}
