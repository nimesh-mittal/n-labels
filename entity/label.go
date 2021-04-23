package entity

type Label struct {
  Namespace string
  Name string
  Active bool
  CreatedBy string
  CreatedAt int64
  UpdatedBy string
  UpdatedAt int64
}

type LabelEntity struct{
  Namespace string
  Name string
  EntityID string
  CreatedBy string
  CreatedAt int64
  UpdatedBy string
  UpdatedAt int64
}

type CreateLabelRequest struct{
  Namespace string
  Name string
}

type AttachLabelRequest struct{
  Namespace string
  EntityId string
}