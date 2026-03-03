// Package shared holds types used by multiple domain packages (nodes, connections, sections).
// It is not part of the public API.
package shared

// EntityAttributeList is the attribute structure for nodes, connections, and sections.
// It maps attribute names to instance ids to values. See the API docs for add_attributes,
// remove_attributes, and attributes request fields.
type EntityAttributeList map[string]map[string]interface{}

// CreatedInfo contains metadata about when and how an entity was created.
type CreatedInfo struct {
	Method    string `json:"method,omitempty"`    // e.g., "desktop", "mobile"
	Timestamp int64  `json:"timestamp,omitempty"` // Unix timestamp in milliseconds
	UID       string `json:"uid,omitempty"`       // User ID who created the entity
}

// PhotoAssociation represents how a photo is associated with an entity.
// The Association field can be a boolean true or the string "main".
type PhotoAssociation struct {
	Association interface{} `json:"association,omitempty"`
}

// PhotoAssociationMap maps photo IDs to their association info.
type PhotoAssociationMap map[string]PhotoAssociation
