package entity

// Label represents label object
type Label struct {
	Namespace string
	Name      string
	Active    bool
	CreatedBy string
	CreatedAt int64
	UpdatedBy string
	UpdatedAt int64
}

// LabelEntity represents relationship between label and entity
type LabelEntity struct {
	Namespace string
	Name      string
	EntityID  string
	CreatedBy string
	CreatedAt int64
	UpdatedBy string
	UpdatedAt int64
}

// CreateLabelRequest represents create label request
type CreateLabelRequest struct {
	Namespace string
	Name      string
}

// AttachLabelRequest represents attach label request
type AttachLabelRequest struct {
	Namespace string
	EntityID  string
}

// DetachLabelRequest represents detach label request
type DetachLabelRequest struct {
	Namespace string
	EntityID  string
}
