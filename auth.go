package thingsboard

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

// Auth controller structure
type Auth struct {
	Token        string
	RefreshToken string
}

// AuthLogin [POST] /api/auth/login
// https://thingsboard.io/docs/reference/rest-api/
func (tb *Thingsboard) authLogin() error {

	a := Auth{}

	_, err := tb.resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"` + tb.user + `", "password":"` + tb.pass + `"}`).
		SetResult(&a). // or SetResult(&AuthSuccess{}).
		Post(tb.apiHost + "/auth/login")

	tb.resty.SetAuthToken(a.Token)

	// Get User details
	err = tb.authUser()
	if err != nil {
		return err
	}

	tb.Auth = &a

	return err
}

// auth-controller : Auth ControllerShow/HideList OperationsExpand Operations
// POST /api/auth/changePassword
// changePassword

// POST /api/auth/logout
// logout

// GET /api/auth/user
// getUser
func (tb *Thingsboard) authUser() error {

	u := User{}
	_, err := tb.resty.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&u).
		Get(tb.apiHost + "/auth/user")
	return err
}

// GET /api/noauth/activate{?activateToken}
// checkActivateToken

// POST /api/noauth/activate{?sendActivationMail}
// activateUser

// POST /api/noauth/oauth2Clients
// getOAuth2Clients

// POST /api/noauth/resetPassword
// resetPassword

// POST /api/noauth/resetPasswordByEmail
// requestResetPasswordByEmail

// GET /api/noauth/resetPassword{?resetToken}
// checkResetToken

// GET /api/noauth/userPasswordPolicy
// getUserPasswordPolicy

// AuthLogout [POST] /api/auth/logout
// https://demo.thingsboard.io/swagger-ui.html#!/auth-controller/logoutUsingPOST
// How it works? Does it work at all? :D
func (tb *Thingsboard) authLogout() error {

	_, err := tb.resty.R().
		Post(tb.apiHost + "/auth/logout")

	return err

}
