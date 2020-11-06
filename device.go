package thingsboard

import "fmt"

// ############################################################################
// Device Controller -> device-controller
// https://demo.thingsboard.io/swagger-ui.html#/device-controller
// ############################################################################

type jsonDevice struct {
	ID             entityID `json:"id"`
	TenantID       entityID `json:"tenantId"`
	CustomerID     entityID `json:"customerId"`
	CreatedTime    int64    `json:"createdTime"`
	AdditionalInfo struct {
		Gateway     bool   `json:"gateway"`
		Description string `json:"description"`
	} `json:"additionalInfo"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

// DELETE /api/customer/device/{deviceId}
// unassignDeviceFromCustomer

// DELETE /api/customer/device/{deviceName}/claim
// reClaimDevice

// POST /api/customer/device/{deviceName}/claim
// claimDevice

// POST /api/customer/public/device/{deviceId}
// assignDeviceToPublicCustomer

// POST /api/customer/{customerId}/device/{deviceId}
// assignDeviceToCustomer

// GET /api/customer/{customerId}/deviceInfos{?type,textSearch,sortProperty,sortOrder,pageSize,page}
// getCustomerDeviceInfos

// GET /api/customer/{customerId}/devices{?type,textSearch,sortProperty,sortOrder,pageSize,page}
// getCustomerDevices

// POST /api/device/credentials
// saveDeviceCredentials
func (tb *Thingsboard) saveDeviceCredentials() {}

// GET /api/device/info/{deviceId}
// getDeviceInfoById

// GET /api/device/types
// getDeviceTypes

// DELETE /api/device/{deviceId}
// deleteDevice

// GET /api/device/{deviceId}
// getDeviceById

// GET /api/device/{deviceId}/credentials
// getDeviceCredentialsByDeviceId

// POST /api/devices
// findByQuery

// GET /api/devices{?deviceIds}
// getDevicesByIds

// POST /api/device{?accessToken}
// saveDevice

// GET /api/tenant/deviceInfos{?type,textSearch,sortProperty,sortOrder,pageSize,page}
// getTenantDeviceInfos

// GET /api/tenant/devices{?deviceName}
// getTenantDevice (by Name!)
func (tb *Thingsboard) getTenantDevice() {}

// GetTenantDevices
// GET /api/tenant/devices{?type,textSearch,sortProperty,sortOrder,*pageSize,*page}
// 401 Unauthorized
// 403 Forbidden
// 404 Not Found
type getTenantDevices struct {
	Data          []jsonDevice `json:"data"`
	TotalPages    int          `json:"totalPages"`
	TotalElements int          `json:"totalElements"`
	HasNext       bool         `json:"hasNext"`
}

// GetTenantDevices downloads information about currently logged user devices available in his Tenant
func (tb *Thingsboard) GetTenantDevices(pageSize int, page int) {
	fmt.Println(pageSize)
	fmt.Println(page)
}

// POST /api/tenant/{tenantId}/device/{deviceId}
// assignDeviceToTenant
