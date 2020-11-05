package thingsboard

import "fmt"

// Auth Controller -> auth-controller
// https://demo.thingsboard.io/swagger-ui.html#/auth-controller
// https://thingsboard.io/docs/reference/rest-api/
// REST API Auth
// Thingsboard uses JWT for request auth. You will need to populate “X-Authorization”
// header using “Authorize” button in the top-right corner of the Swagger UI.
//
// Request:
// curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json'
// -d '{"username":"tenant@thingsboard.org", "password":"tenant"}' 'http://THINGSBOARD_URL/api/auth/login'
//
// Response:
// { "token":		 "$YOUR_JWT_TOKEN",
//   "refreshToken": "$YOUR_JWT_REFRESH_TOKEN" }

// {
// 	"status": 401,
// 	"message": "Token has expired",
// 	"errorCode": 11,
// 	"timestamp": "2020-11-05T09:55:30.235+0000"
// }

type Auth struct {
	Token        string
	RefreshToken string
}

// AuthLogin [POST] /api/auth/login
// https://thingsboard.io/docs/reference/rest-api/
func (tb *Thingsboard) AuthLogin() error {

	a := Auth{}

	_, err := tb.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"` + tb.user + `", "password":"` + tb.pass + `"}`).
		SetResult(&a). // or SetResult(&AuthSuccess{}).
		Post(tb.apiHost + "/auth/login")

	tb.resty.SetAuthToken(a.Token)
	tb.Auth = &a

	return err
}

// AuthUser Response model:
// User {
// 	additionalInfo (string, optional),
// 	authority (string, optional) = ['SYS_ADMIN', 'TENANT_ADMIN', 'CUSTOMER_USER', 'REFRESH_TOKEN']stringEnum:"SYS_ADMIN", "TENANT_ADMIN", "CUSTOMER_USER", "REFRESH_TOKEN",
// 	createdTime (integer, optional),
// 	customerId (CustomerId, optional),
// 	email (string, optional),
// 	firstName (string, optional),
// 	id (UserId, optional),
// 	lastName (string, optional),
// 	name (string, optional),
// 	tenantId (TenantId, optional)
// 	}
// CustomerId {
// 	id (string, optional)
// 	}
// UserId {
// 	id (string, optional)
// 	}
// TenantId {
// 	id (string, optional)
// 	}
func (tb *Thingsboard) AuthUser() {}

// AuthLogout [POST] /api/auth/logout
// https://demo.thingsboard.io/swagger-ui.html#!/auth-controller/logoutUsingPOST
// How it works? Does it work at all? :D
func (tb *Thingsboard) AuthLogout() error {

	r, err := tb.resty.R().
		Post(tb.apiHost + "/auth/logout")
	fmt.Println(r)
	return err

}
