package thingsboard

// ############################################################################
// User Controller -> user-controller
// http://demo.thingsboard.io/swagger-ui.html#/user-controller
// ############################################################################

// User structure
// 	authority (string, optional) = ['SYS_ADMIN',
// 					'TENANT_ADMIN', 'CUSTOMER_USER', 'REFRESH_TOKEN']
type jsonUser struct {
	ID             entityID `json:"id"`
	TenantID       entityID `json:"tenantId"`
	CustomerID     entityID `json:"customerId"`
	CreatedTime    int64    `json:"createdTime"`
	Email          string   `json:"email"`
	Authority      string   `json:"authority"` // TODO: ENUM
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Name           string   `json:"name"`
	AdditionalInfo struct {
		Description                string           `json:"description"`
		DefaultDashboardID         string           `json:"defaultDashboardId"`
		DefaultDashboardFullscreen bool             `json:"defaultDashboardFullscreen"`
		LastLoginTs                int64            `json:"lastLoginTs"`
		UserPasswordHistory        map[int64]string `json:"userPasswordHistory"`
		FailedLoginAttempts        int              `json:"failedLoginAttempts"`
	} `json:"additionalInfo"`
}

// GET /api/customer/{customerId}/users{?textSearch,sortProperty,sortOrder,pageSize,page}
// getCustomerUsers

// GET /api/tenant/{tenantId}/users{?textSearch,sortProperty,sortOrder,pageSize,page}
// getTenantAdmins

// POST /api/user/sendActivationMail{?email}
// sendActivationEmail

// GET /api/user/tokenAccessEnabled
// isUserTokenAccessEnabled

// DELETE /api/user/{userId}
// deleteUser

// GET /api/user/{userId}
// getUserById

// GET /api/user/{userId}/activationLink
// getActivationLink

// GET /api/user/{userId}/token
// getUserToken

// POST /api/user/{userId}/userCredentialsEnabled{?userCredentialsEnabled}
// setUserCredentialsEnabled

// GET /api/users{?textSearch,sortProperty,sortOrder,pageSize,page}
// getUsers

// POST /api/user{?sendActivationMail}
// saveUser
