package thingsboard

import "fmt"

// ############################################################################
//  Thingsboard REST error controller
// ############################################################################

// HTTP Authentication and error codes
// We will use access token device credentials in this article and they will be referred to later as $ACCESS_TOKEN. The application needs to include $ACCESS_TOKEN as a path parameter in each HTTP request. Possible error codes and their reasons:

// 400 Bad Request - Invalid URL, request parameters or body.
// 401 Unauthorized - Invalid $ACCESS_TOKEN.
// 404 Not Found - Resource not found.

// TBError holds and parses Thingboard JSON error response
// * - Field is always present in JSON error response
type TBError struct {
	Timestamp string `json:"timestamp"` // *
	Status    int    `json:"status"`    // *
	Message   string `json:"message"`   // *
	Code      int    `json:"errorCode"`
	Type      string `json:"error"`
	Path      string `json:"path"`
}

func (e *TBError) Error() string {

	// TODO: Good Error parsing

	s := fmt.Sprintf("REST Error: [%v] [%v]",
		e.Timestamp, e.Status)

	if e.Code > 0 {
		s += fmt.Sprintf(" Code: %v",
			e.Code)
	}

	if e.Path != "" {
		s += fmt.Sprintf(" [%v]",
			e.Path)
	}

	s += fmt.Sprintf(" %v %v",
		e.Message, e.Type)

	return s

}

// 1st type of error reponses
// {
// 	"status": 400,
// 	"message": "Invalid UUID string: xxx",
// 	"errorCode": 31,
// 	"timestamp": "2020-11-07T17:44:59.544+0000"
//   }

// 2nd type
// {
// "timestamp": "2020-11-07T17:59:23.691+0000",
// "status": 400,
// "error": "Bad Request",
// "message": "JSON parse error: Unrecognized token 'fff': was expecting (JSON String, Number, Array, Object or token 'null', 'true' or 'false'); nested exception is com.fasterxml.jackson.core.JsonParseException: Unrecognized token 'fff': was expecting (JSON String, Number, Array, Object or token 'null', 'true' or 'false')\n at [Source: (PushbackInputStream); line: 1, column: 4]",
// "path": "/api/devices"
//   }

// {
// 	"status": 401,
// 	"message": "Token has expired",
// 	"errorCode": 11,
// 	"timestamp": "2020-11-09T20:36:22.306+0000"
//   }
