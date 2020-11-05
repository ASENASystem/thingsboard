package thingsboard

type userID entityID

// User structure
// 	authority (string, optional) = ['SYS_ADMIN', 'TENANT_ADMIN', 'CUSTOMER_USER', 'REFRESH_TOKEN']
// 						 stringEnum:"SYS_ADMIN", "TENANT_ADMIN", "CUSTOMER_USER", "REFRESH_TOKEN",
type User struct {
	id             userID
	createdTime    string
	additionalInfo string
	tenantID       tenantID
	customerID     customerID
	email          string
	authority      string
	firstName      string
	lastName       string
	name           string
}

type jsonUser struct {
	ID             entityID   `json:"id"`
	TenantID       tenantID   `json:"tenantId"`
	CustomerID     customerID `json:"customerId"`
	CreatedTime    int64      `json:"createdTime"`
	Email          string     `json:"email"`
	Authority      string     `json:"authority"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Name           string     `json:"name"`
	AdditionalInfo struct {
		Description                string `json:"description"`
		DefaultDashboardID         string `json:"defaultDashboardId"`
		DefaultDashboardFullscreen bool   `json:"defaultDashboardFullscreen"`
		LastLoginTs                int64  `json:"lastLoginTs"`
		FailedLoginAttempts        int    `json:"failedLoginAttempts"`
	} `json:"additionalInfo"`
}
