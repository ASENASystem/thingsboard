package thingsboard

type entityID struct {
	EntityType string `json:"entityType"`
	ID         string `json:"id"`
}

// TODO: Correct ENUM type for EntityType
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
