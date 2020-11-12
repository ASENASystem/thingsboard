# Thingsboard

Thingsboard Go Library

https://gitlab.com/asenasystem-opensource/go/thingsboard

## Thingsboard Entity structure

* System admin (**ONE** global TB administrator account)
    * Tenant Adminsrators
        * Customers 
            * Users
        * Assets
        * Devices 
        * Dashboards
        * Entity Views
        * Rule Chains

```
Tenant Administrator is able to do following actions:
* Provision and Manage Devices.
* Provision and Manage Assets.
* Create and Manage Customers.
* Create and Manage Dashboards.
* Configure Rule Engine
* Add or modify default widgets using Widget Library.

Customer
Customer may be a separate business-entity: individual or organization who purchase or uses tenant devices and/or assets. Customer may also be a division within Tenant organization. Customer may have multiple users, inner customers and millions of devices and/or assets.

User
Users are able to login to ThingsBoard web interface, execute REST API calls, access devices and assets if allowed. User is also an Entity in ThingsBoard.
```

## Go

* Resty v2: https://github.com/go-resty/resty/v2

## UML Diagram

Before every commit Git `pre-commit` script is run in order to generate below UML diagram for this project.
Please check https://github.com/jfeliu007/goplantuml project for tools needed.

[UML Diagram](DIAGRAM.md)