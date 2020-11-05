package thingsboard

// Device Controller -> device-controller
// https://demo.thingsboard.io/swagger-ui.html#/device-controller

// device-controller : Device ControllerShow/HideList OperationsExpand Operations
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

// GET /api/tenant/devices{?type,textSearch,sortProperty,sortOrder,*pageSize,*page}
// getTenantDevices

// POST /api/tenant/{tenantId}/device/{deviceId}
// assignDeviceToTenant
