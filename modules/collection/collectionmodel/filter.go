package collectionmodel

type Filter struct {
	FakeCreatedBy string `json:"created_by" form:"created_by"`
	CreatedBy     int
	FakeParentId  string `json:"parent_id" form:"parent_id"`
	ParentId      int
}
