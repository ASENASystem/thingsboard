// Package thingsboard ...
// Author: Paweł 'felixd' Wojciechowski - Outsourcing IT Konopnickiej.Com
// (c) 2020 ASENA System
package thingsboard

import (
	"github.com/go-resty/resty/v2"
)

// ‘X-Authorization’ to “Bearer $YOUR_JWT_TOKEN”

// Thingsboard ...
type Thingsboard struct {
	user    string
	pass    string
	host    string
	apiHost string
	auth    Auth
	resty   *resty.Client
}

// New returns Thingsboard instance
func New(host string, user string, pass string) (*Thingsboard, error) {

	tb := Thingsboard{
		user:    user,
		pass:    pass,
		host:    host,
		apiHost: host + "/api",
	}

	// RESTy v2 client
	// https://github.com/go-resty/resty/v2
	tb.resty = resty.New()

	err := tb.AuthLogin()
	if err != nil {
		return nil, err
	}
	return &tb, err
}
