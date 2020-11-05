package thingsboard

type Error struct {
	// Status: "status": 401,
	Status int8
	// Message: "message": "Token has expired",
	Message string
	// ErrorCode: "errorCode": 11,
	ErrorCode int8
	// Timestamp: "timestamp": "2020-11-05T09:55:30.235+0000"
	Timestamp string
}
