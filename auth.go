package thingsboard

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Auth Controller -> auth-controller
// https://demo.thingsboard.io/swagger-ui.html#/auth-controller
// https://thingsboard.io/docs/reference/rest-api/
// REST API Auth
// ThingsBoard uses JWT for request auth. You will need to populate “X-Authorization”
// header using “Authorize” button in the top-right corner of the Swagger UI.
//
// Request:
// curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json'
// -d '{"username":"tenant@thingsboard.org", "password":"tenant"}' 'http://THINGSBOARD_URL/api/auth/login'
//
// Response:
// { "token":		 "$YOUR_JWT_TOKEN",
//   "refreshToken": "$YOUR_JWT_REFRESH_TOKEN" }
type Auth struct {
	Token        string
	RefreshToken string
}

// AuthLogin /api/auth/login
// https://thingsboard.io/docs/reference/rest-api/
func (tb *Thingsboard) AuthLogin() error {

	r := resty.New()
	resp, err := r.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"` + tb.user + `", "password":"` + tb.pass + `"}`).
		SetResult(Auth{}). // or SetResult(&AuthSuccess{}).
		Post(tb.apiHost + "/auth/login")

	fmt.Println(resp)

	return err
}

func logout() {}
