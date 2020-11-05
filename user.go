package thingsboard

type userID entityID
type User jsonUser

// User structure
// 	authority (string, optional) = ['SYS_ADMIN',
// 					'TENANT_ADMIN', 'CUSTOMER_USER', 'REFRESH_TOKEN']
type jsonUser struct {
	ID             userID     `json:"id"`
	TenantID       tenantID   `json:"tenantId"`
	CustomerID     customerID `json:"customerId"`
	CreatedTime    int64      `json:"createdTime"`
	Email          string     `json:"email"`
	Authority      string     `json:"authority"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Name           string     `json:"name"`
	AdditionalInfo struct {
		Description                string           `json:"description"`
		DefaultDashboardID         string           `json:"defaultDashboardId"`
		DefaultDashboardFullscreen bool             `json:"defaultDashboardFullscreen"`
		LastLoginTs                int64            `json:"lastLoginTs"`
		UserPasswordHistory        map[int64]string `json:"userPasswordHistory"`
		FailedLoginAttempts        int              `json:"failedLoginAttempts"`
	} `json:"additionalInfo"`
}
