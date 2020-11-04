package thingsboard

import "fmt"

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

// AuthLogout [POST] /api/auth/logout
// https://demo.thingsboard.io/swagger-ui.html#!/auth-controller/logoutUsingPOST
// How it works? Does it work at all? :D
func (tb *Thingsboard) AuthLogout() error {

	r, err := tb.resty.R().
		Post(tb.apiHost + "/auth/logout")
	fmt.Println(r)
	return err

}
