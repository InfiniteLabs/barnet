package client

import (
	"errors"
)

type ClientId string

type Client struct {
	ClientId            ClientId // May come from JWT token
	EmailAddress        string
	FullName            string
	WhatWeShouldCallYou string
	Gender              string
	TelephoneNumber     string
}

type Repository interface {
	Store(client *Client) error
	Find(clientId ClientId) (*Client, error)
}

var ErrUnknown = errors.New("unknown client")
