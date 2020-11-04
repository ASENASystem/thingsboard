// Package thingsboard ...
// Author: Paweł 'felixd' Wojciechowski - Outsourcing IT Konopnickiej.Com
// (c) 2020 ASENA System
package thingsboard

import (
	"strings"

	"github.com/go-resty/resty/v2"
)

// ‘X-Authorization’ to “Bearer $YOUR_JWT_TOKEN”

// Thingsboard ...
type Thingsboard struct {
	user    string
	pass    string
	host    string
	apiHost string
	Auth    *Auth
	resty   *resty.Client
}

// New returns Thingsboard instance
func Connect(host string, user string, pass string) (*Thingsboard, error) {

	tb := Thingsboard{
		user:    user,
		pass:    pass,
		host:    host,
		apiHost: host + "/api",
	}

	// RESTy v2: Client - https://github.com/go-resty/resty/v2
	tb.resty = resty.New()

	// RESTy v2: Append middleware function to set proper X-Authorization header.
	// https://github.com/go-resty/resty/blob/master/client.go#L359
	tb.resty.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		c.Header.Set("X-Authorization", "Bearer "+c.Token)
		return nil
	})

	err := tb.AuthLogin()
	if err != nil {
		return nil, err
	}
	return &tb, err
}

// IsStringEmpty method tells whether given string is empty or not
func IsStringEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
