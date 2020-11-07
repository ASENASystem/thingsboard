package thingsboard

import (
	"strconv"
)

// ############################################################################
// Device Controller -> device-controller
// https://demo.thingsboard.io/swagger-ui.html#/device-controller
// ############################################################################

// Device structure to hold information about devices
type Device struct {
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

// GetTenantDevices represents structure for GetTenantDevices() method
type GetTenantDevices struct {
	Data          []Device `json:"data"`
	TotalPages    int      `json:"totalPages"`
	TotalElements int      `json:"totalElements"`
	HasNext       bool     `json:"hasNext"`
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

// GetDeviceByID GET /api/device/{deviceId}
// TODO: HTTP Request codes
func (tb *Thingsboard) GetDeviceByID(deviceID string) (*Device, error) {

	d := Device{}

	_, err := tb.resty.R().
		SetResult(&d).
		Get(tb.apiHost + "/device/x" + deviceID)

	return &d, err

}

// GET /api/device/{deviceId}/credentials
// getDeviceCredentialsByDeviceId

// POST /api/devices
// findByQuery

// GetDevicesByIds GET /api/devices{?deviceIds}
func (tb *Thingsboard) GetDevicesByIds(deviceIDs string) ([]Device, error) {

	d := []Device{}

	_, err := tb.resty.R().
		SetQueryParams(map[string]string{
			"deviceIds": string(deviceIDs),
		}).
		SetResult(&d).
		Get(tb.apiHost + "/devices")

	return d, err
}

// POST /api/device{?accessToken}
// saveDevice

// GET /api/tenant/deviceInfos{?type,textSearch,sortProperty,sortOrder,pageSize,page}
// getTenantDeviceInfos

// GetTenantDevice (by Name!)
// GET /api/tenant/devices{?deviceName}
func (tb *Thingsboard) GetTenantDevice(name string) (Device, error) {

	d := Device{}

	_, err := tb.resty.R().
		SetQueryParams(map[string]string{
			"deviceName": string(name),
		}).
		SetResult(&d).
		Get(tb.apiHost + "/tenant/devices")

	return d, err
}

// GetTenantDevices downloads information about currently logged user devices available in his Tenant
// GET /api/tenant/devices{?type,textSearch,sortProperty,sortOrder,*pageSize,*page}
// 401 Unauthorized
// 403 Forbidden
// 404 Not Found
func (tb *Thingsboard) GetTenantDevices(pageSize int, page int) (GetTenantDevices, error) {

	d := GetTenantDevices{}

	_, err := tb.resty.R().
		SetQueryParams(map[string]string{
			"pageSize": strconv.Itoa(pageSize),
			"page":     strconv.Itoa(page),
		}).
		SetResult(&d).
		Get(tb.apiHost + "/tenant/devices")

	// TODO: Check requqest response codes

	if err != nil {
		return d, err
	}

	return d, nil
}

// POST /api/tenant/{tenantId}/device/{deviceId}
// assignDeviceToTenant
