package thingsboard

// ############################################################################
//  Thingsboard Error controller
// ############################################################################

type jsonError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
