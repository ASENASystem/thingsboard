package thingsboard

// ############################################################################
//  Thingsboard RESTy controller
// ############################################################################

// TBError holds Thingboard error response
type TBError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
