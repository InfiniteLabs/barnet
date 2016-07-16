package repository

import (
	"sync"

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
