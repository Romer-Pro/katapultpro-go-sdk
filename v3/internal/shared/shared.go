// Package shared holds types used by multiple domain packages (nodes, connections, sections).
// It is not part of the public API.
package shared

// EntityAttributeList is the attribute structure for nodes, connections, and sections.
// It maps attribute names to instance ids to values. See the API docs for add_attributes,
// remove_attributes, and attributes request fields.
type EntityAttributeList map[string]map[string]interface{}
