// Package thingsboard ...
// Author: Paweł 'felixd' Wojciechowski - Outsourcing IT Konopnickiej.Com
// (c) 2020 ASENA System

// https://thingsboard.io/docs/user-guide/entities-and-relations/
// Entities Overview
// Thingsboard provides the user interface and REST APIs to provision and manage multiple entity types and their relations in your IoT application. Supported entities are:

// Tenants - you can treat tenant as a separate business-entity: individual or organization who owns or produce devices and assets; Tenant may have multiple tenant administrator users and millions of customers;
// Customers - customer is also a separate business-entity: individual or organization who purchase or uses tenant devices and/or assets; Customer may have multiple users and millions of devices and/or assets;
// Users - users are able to browse dashboards and manage entities;
// Devices - basic IoT entities that may produce telemetry data and handle RPC commands. For example sensors, actuators, switches;
// Assets - abstract IoT entities that may be related to other devices and assets. For example factory, field, vehicle;
// Alarms - events that identify issues with your assets, devices or other entities;
// Dashboards - visualization of your IoT data and ability to control particular devices through user interface;
// Rule Node - processing units for incoming messages, entity lifecycle events, etc;
// Rule Chain - logic unit of related Rule Nodes;
// Each entity supports:

// Attributes - static and semi-static key-value pairs associated with entities. For example serial number, model, firmware version;
// Telemetry data - time-series data points available for storage, querying and visualization. For example temperature, humidity, battery level;
// Relations - directed connections to other entities. For example contains, manages, owns, produces.
// Additionally, devices and assets also have a type. This allows distinguising them and process data from them in a different way.

// This guide provides the overview of the features listed above, some useful links to get more details and real-life examples of their usage.

package thingsboard

import (
	"github.com/go-resty/resty/v2"
)

// Thingsboard ...
type Thingsboard struct {
	user    string
	pass    string
	host    string
	apiHost string
	Auth    *Auth
	User    *User
	resty   *resty.Client
}

// New returns new Thingsboard instance
func New(host string, user string, pass string) (*Thingsboard, error) {

	tb := Thingsboard{
		user:    user,
		pass:    pass,
		host:    host,
		apiHost: host + "/api",
	}

	err := tb.Connect()

	return &tb, err
}

// Connect is used by New() method to connect to Thingsboard server.
func (tb *Thingsboard) Connect() error {
	// RESTy v2: Client - https://github.com/go-resty/resty/v2
	tb.resty = resty.New()

	// RESTy v2: Append middleware function to set proper X-Authorization header.
	// https://github.com/go-resty/resty/blob/master/client.go#L359
	tb.resty.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		c.Header.Set("X-Authorization", "Bearer "+c.Token)
		return nil
	})

	err := tb.login()
	if err != nil {
		return err
	}

	return nil
}

// Disconnect should logout, remove currently authenticated user (authorization for token should be removed).
// But it dosn't work at all on Thingsboard.
func (tb *Thingsboard) Disconnect() error {
	return tb.logout()
}
