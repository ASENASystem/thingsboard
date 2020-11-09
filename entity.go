package thingsboard

type entityID struct {
	EntityType string `json:"entityType"`
	ID         string `json:"id"`
}

// EntityTypes holds types of all Thingsboard entities ;]
type EntityTypes string

// const (
// 	TENANT EntityType = iota
// 	CUSTOMER
// 	USER
// 	DASHBOARD
// 	ASSET
// 	DEVICE
// 	ALARM
// 	RULE_CHAIN
// 	RULE_NODE
// 	ENTITY_VIEW
// 	WIDGETS_BUNDLE
// 	WIDGET_TYPE
// 	TENANT_PROFILE
// 	DEVICE_PROFILE
// )
