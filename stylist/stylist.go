package stylist

import (
	"errors"

	"github.com/rs/xid"
)

type StylistId xid.ID

type Stylist struct {
	StylistId StylistId
	Name      string
	Pitch     string
}

func New(name string, pitch string) *Stylist {
	guid := xid.New()

	return &Stylist{
		StylistId: StylistId(guid),
		Name:      name,
		Pitch:     pitch,
	}
}

type Repository interface {
	Store(stylist *Stylist) error
	Find(stylistId StylistId) (*Stylist, error)
	FindAll() []*Stylist
}

var ErrUnknown = errors.New("unknown stylist")
