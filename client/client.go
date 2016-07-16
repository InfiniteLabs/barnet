package client

import (
	"errors"

	"github.com/rs/xid"
)

type ClientId xid.ID

type Client struct {
	ClientId            ClientId
	EmailAddress        string
	FullName            string
	WhatWeShouldCallYou string
	Gender              string
	TelephoneNumber     string
}

func New(emailAdrress string, fullname string, whatWeShouldCallYou string, gender string, telephoneNumber string) *Client {
	guid := xid.New()

	return &Client{
		ClientId:            ClientId(guid),
		FullName:            fullname,
		WhatWeShouldCallYou: whatWeShouldCallYou,
		Gender:              gender,
		TelephoneNumber:     telephoneNumber,
	}
}

type Repository interface {
	Store(client *Client) error
	Find(clientId ClientId) (*Client, error)
}

var ErrUnknown = errors.New("unknown client")
