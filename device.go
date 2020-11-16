package thingsboard

import (
	"errors"
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

// DeviceCredentials struct to hold getDeviceCredentialsByDeviceId() repsponse
type DeviceCredentials struct {
	CreatedTime   int    `json:"createdTime"`
	CredentialsID string `json:"credentialsId"`
	// ['ACCESS_TOKEN', 'X509_CERTIFICATE']stringEnum:"ACCESS_TOKEN", "X509_CERTIFICATE",
	CredentialsType  string   `json:"credentialsType"`
	CredentialsValue string   `json:"credentialsValue"`
	DeviceID         entityID `json:"deviceId"`
	ID               entityID `json:"id"` // DeviceCredentialsId
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
		Get(tb.api + "/device/" + deviceID)

	return &d, err
}

// GetDeviceCredentialsByDeviceID returns device credentials structure
// GET /api/device/{deviceId}/credentials
func (tb *Thingsboard) GetDeviceCredentialsByDeviceID(deviceID string) (*DeviceCredentials, error) {
	return &DeviceCredentials{}, nil
}

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
		Get(tb.api + "/devices")

	return d, err
}

// POST /api/device{?accessToken}
// saveDevice

// GET /api/tenant/deviceInfos{?type,textSearch,sortProperty,sortOrder,pageSize,page}
// getTenantDeviceInfos

// GetDeviceAccessTokenByName returns Acccess Token for Device with specified name
func (tb *Thingsboard) GetDeviceAccessTokenByName(deviceName string) (string, error) {
	dc, err := tb.GetDeviceCredentialsByDeviceName(deviceName)
	if err != nil {
		return "", err
	}

	if dc.CredentialsType == "ACCESS_TOKEN" {
		return dc.CredentialsValue, nil
	}

	return "", errors.New("GetDeviceAccessTokenByName: ACCESS_TOKEN was not provided in CredentialsType")
}

// GetDeviceAccessTokenByID returns Acccess Token for Device with specified name
func (tb *Thingsboard) GetDeviceAccessTokenByID(deviceID string) (string, error) {
	dc, err := tb.GetDeviceCredentialsByDeviceID(deviceID)
	if err != nil {
		return "", err
	}

	if dc.CredentialsType == "ACCESS_TOKEN" {
		return dc.CredentialsValue, nil
	}

	return "", errors.New("GetDeviceAccessTokenByID: ACCESS_TOKEN was not provided in CredentialsType")
}

// GetDeviceCredentialsByDeviceName returns Access Token for Device (by name)
func (tb *Thingsboard) GetDeviceCredentialsByDeviceName(name string) (*DeviceCredentials, error) {

	d, err := tb.GetDeviceByName(name)
	if err != nil {
		return &DeviceCredentials{}, err
	}

	dc, err := tb.GetDeviceCredentialsByDeviceID(d.ID.ID)
	if err != nil {
		return &DeviceCredentials{}, err
	}

	return dc, nil
}

// GetTenantDevice is an alias of GetDeviceByName() - compatibility with TB Swagger
func (tb *Thingsboard) GetTenantDevice(name string) (*Device, error) {
	return tb.GetDeviceByName(name)
}

// GetDeviceByName (by Name!)
// GET /api/tenant/devices{?deviceName}
func (tb *Thingsboard) GetDeviceByName(name string) (*Device, error) {

	d := Device{}

	_, err := tb.resty.R().
		SetQueryParams(map[string]string{
			"deviceName": string(name),
		}).
		SetResult(&d).
		Get(tb.api + "/tenant/devices")

	return &d, err
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
		Get(tb.api + "/tenant/devices")

	// TODO: Check requqest response codes

	if err != nil {
		return d, err
	}

	return d, nil
}

// POST /api/tenant/{tenantId}/device/{deviceId}
// assignDeviceToTenant
