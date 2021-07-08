package domain

// Attribute can be part of either a subject or (later) group.
type Attribute struct {
	IID       int64  `json:"-"`
	OwnerID   int64  `json:"-"`
	OwnerType int8   `json:"-"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}
